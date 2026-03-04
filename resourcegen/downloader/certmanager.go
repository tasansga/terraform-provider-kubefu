package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	certManagerRepo          = "cert-manager/cert-manager"
	certManagerBranchesURL   = "https://api.github.com/repos/" + certManagerRepo + "/branches?per_page=100"
	certManagerTreeURL       = "https://api.github.com/repos/" + certManagerRepo + "/git/trees/%s?recursive=1"
	certManagerRawFileURL    = "https://raw.githubusercontent.com/" + certManagerRepo + "/release-%s/%s"
	certManagerCRDsDirectory = "deploy/crds/"
	minCertManagerVersion    = "0.0.0"
)

// CertManagerDownloader downloads cert-manager CRDs for every release branch.
type CertManagerDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewCertManagerDownloader creates a downloader for cert-manager releases.
func NewCertManagerDownloader(logger logrus.FieldLogger) *CertManagerDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &CertManagerDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions returns cert-manager release versions from GitHub.
func (d *CertManagerDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := certManagerBranchesURL

	for nextURL != "" {
		req, err := http.NewRequest(http.MethodGet, nextURL, nil)
		if err != nil {
			return nil, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", userAgent)

		resp, err := doRequest(d.client, req)
		if err != nil {
			return nil, fmt.Errorf("list release branches: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("list release branches: %s%s", resp.Status, githubAPIError(resp))
		}

		var branches []struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&branches); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("decode branches response: %w", err)
		}
		resp.Body.Close()

		for _, branch := range branches {
			if !strings.HasPrefix(branch.Name, "release-") {
				continue
			}
			if matches := releaseBranchRE.FindStringSubmatch(branch.Name); matches != nil {
				versions[matches[1]] = struct{}{}
			}
		}

		nextURL = parseNextLink(resp.Header.Get("Link"))
	}

	return sortedVersions(versions), nil
}

// DownloadMissingSchemas downloads any missing cert-manager CRDs into schemaDir.
func (d *CertManagerDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		return 0, fmt.Errorf("create schema dir: %w", err)
	}

	versions, err := d.FetchReleaseVersions()
	if err != nil {
		return 0, err
	}

	total := 0
	var downloadErrors []error
	for _, version := range versions {
		if compareSemver(version, minCertManagerVersion) < 0 {
			continue
		}
		count, errs := d.downloadReleaseCRDs(version, schemaDir)
		total += count
		downloadErrors = append(downloadErrors, errs...)
	}

	if len(downloadErrors) > 0 {
		var errStrings []string
		for _, err := range downloadErrors {
			errStrings = append(errStrings, err.Error())
		}
		return total, fmt.Errorf("download errors: %s", strings.Join(errStrings, "; "))
	}
	return total, nil
}

func (d *CertManagerDownloader) downloadReleaseCRDs(version, schemaDir string) (int, []error) {
	treeURL := fmt.Sprintf(certManagerTreeURL, "release-"+version)
	req, err := http.NewRequest(http.MethodGet, treeURL, nil)
	if err != nil {
		return 0, []error{fmt.Errorf("build tree request for %s: %w", version, err)}
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return 0, []error{fmt.Errorf("list tree for %s: %w", version, err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, []error{fmt.Errorf("tree %s: %s%s", version, resp.Status, githubAPIError(resp))}
	}

	var tree struct {
		Tree []struct {
			Path string `json:"path"`
			Type string `json:"type"`
		} `json:"tree"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tree); err != nil {
		return 0, []error{fmt.Errorf("decode tree response for %s: %w", version, err)}
	}

	var (
		downloadErrors []error
		count          int
	)

	for _, entry := range tree.Tree {
		if entry.Type != "blob" {
			continue
		}
		if !strings.HasPrefix(entry.Path, certManagerCRDsDirectory) {
			continue
		}
		if !strings.HasSuffix(entry.Path, ".yaml") {
			continue
		}

		targetName := fmt.Sprintf("cert-manager-%s-%s", version, filepath.Base(entry.Path))
		targetPath := filepath.Join(schemaDir, targetName)

		fields := logrus.Fields{
			"version":    version,
			"schemaPath": targetPath,
			"schemaDir":  schemaDir,
			"crd":        entry.Path,
		}

		if exists(targetPath) {
			d.logger.WithFields(fields).WithField("status", "skipped").Info("skipping existing cert-manager schema")
			continue
		}

		url := fmt.Sprintf(certManagerRawFileURL, version, entry.Path)
		if err := d.downloadCRD(url, targetPath); err != nil {
			d.logger.WithFields(fields).WithError(err).Error("failed to download cert-manager schema")
			downloadErrors = append(downloadErrors, fmt.Errorf("version %s file %s: %w", version, entry.Path, err))
			continue
		}

		d.logger.WithFields(fields).WithField("status", "downloaded").Info("downloaded cert-manager schema")
		count++
	}

	return count, downloadErrors
}

func (d *CertManagerDownloader) downloadCRD(url, targetPath string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http %s%s", resp.Status, githubAPIError(resp))
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(targetPath), "schema-*.tmp")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}

	if err := os.Rename(tmpFile.Name(), targetPath); err != nil {
		return err
	}

	return nil
}
