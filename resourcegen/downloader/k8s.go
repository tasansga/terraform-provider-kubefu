package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	kubernetesRepo       = "kubernetes/kubernetes"
	branchesURL          = "https://api.github.com/repos/" + kubernetesRepo + "/branches?per_page=100"
	rawSchemaURLPattern  = "https://raw.githubusercontent.com/" + kubernetesRepo + "/release-%s/api/openapi-spec/swagger.json"
	minKubernetesVersion = "1.5.0"
)

var releaseBranchRE = regexp.MustCompile(`^release-(\d+\.\d+)$`)

type KubernetesDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewKubernetesDownloader creates a downloader with the default HTTP client and
// the provided logger. If logger is nil we fall back to a standalone logrus logger.
func NewKubernetesDownloader(logger logrus.FieldLogger) *KubernetesDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &KubernetesDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions returns sorted release versions (e.g., "1.24") for
// release branches that match release-1.* in the upstream repository.
func (d *KubernetesDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := branchesURL

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
			if !strings.HasPrefix(branch.Name, "release-1.") {
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

func sortedVersions(versions map[string]struct{}) []string {
	list := make([]string, 0, len(versions))
	for version := range versions {
		list = append(list, version)
	}
	sort.Slice(list, func(i, j int) bool {
		return compareSemver(list[i], list[j]) < 0
	})
	return list
}

// DownloadMissingSchemas writes any missing release swagger files to schemaDir.
// It returns the number of downloaded files.
func (d *KubernetesDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		return 0, fmt.Errorf("create schema dir: %w", err)
	}

	versions, err := d.FetchReleaseVersions()
	if err != nil {
		return 0, err
	}

	d.logger.WithField("schemaDir", schemaDir).Info("scanning schema directory")

	var (
		count          int
		downloadErrors []error
	)

	for _, version := range versions {
		if compareSemver(version, minKubernetesVersion) < 0 {
			continue
		}
		targetPath := filepath.Join(schemaDir, fmt.Sprintf("kubernetes-api-%s.json", version))
		if exists(targetPath) {
			d.logger.WithFields(logrus.Fields{
				"version":     version,
				"schemaPath":  targetPath,
				"schemaDir":   schemaDir,
				"status":      "skipped",
				"description": "schema already exists",
			}).Info("skipping existing schema")
			continue
		}

		ok, err := d.downloadSchema(version, targetPath)
		if err != nil {
			d.logger.WithFields(logrus.Fields{
				"version":    version,
				"schemaPath": targetPath,
				"schemaDir":  schemaDir,
			}).WithError(err).Error("failed to download schema")
			downloadErrors = append(downloadErrors, fmt.Errorf("version %s: %w", version, err))
			continue
		}
		if ok {
			d.logger.WithFields(logrus.Fields{
				"version":    version,
				"schemaPath": targetPath,
				"schemaDir":  schemaDir,
				"status":     "downloaded",
			}).Info("downloaded schema")
			count++
		}
	}

	if len(downloadErrors) > 0 {
		var errStrings []string
		for _, err := range downloadErrors {
			errStrings = append(errStrings, err.Error())
		}
		return count, fmt.Errorf("download errors: %s", strings.Join(errStrings, "; "))
	}
	return count, nil
}

func (d *KubernetesDownloader) downloadSchema(version, targetPath string) (bool, error) {
	url := fmt.Sprintf(rawSchemaURLPattern, version)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		d.logger.WithFields(logrus.Fields{
			"version":    version,
			"schemaPath": targetPath,
			"schemaDir":  filepath.Dir(targetPath),
		}).Warn("schema not found (404); skipping version")
		return false, nil
	default:
		return false, fmt.Errorf("http %s%s", resp.Status, githubAPIError(resp))
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(targetPath), "schema-*.tmp")
	if err != nil {
		return false, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return false, err
	}
	if err := tmpFile.Close(); err != nil {
		return false, err
	}

	if err := os.Rename(tmpFile.Name(), targetPath); err != nil {
		return false, err
	}
	return true, nil
}
