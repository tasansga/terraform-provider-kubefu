package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	karpenterRepo       = "aws/karpenter-provider-aws"
	karpenterTagsURL    = "https://api.github.com/repos/" + karpenterRepo + "/tags?per_page=100"
	karpenterCRDsPath   = "pkg/apis/crds"
	minKarpenterVersion = "0.37.0"
)

// KarpenterDownloader downloads Karpenter CRD bundles from GitHub.
type KarpenterDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewKarpenterDownloader builds a downloader for Karpenter releases.
func NewKarpenterDownloader(logger logrus.FieldLogger) *KarpenterDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &KarpenterDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions lists Karpenter release tags.
func (d *KarpenterDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := karpenterTagsURL

	for nextURL != "" {
		req, err := http.NewRequest(http.MethodGet, nextURL, nil)
		if err != nil {
			return nil, fmt.Errorf("build request: %w", err)
		}
		req.Header.Set("User-Agent", userAgent)

		resp, err := doRequest(d.client, req)
		if err != nil {
			return nil, fmt.Errorf("list tags: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("list tags: %s%s", resp.Status, githubAPIError(resp))
		}

		var tags []struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("decode tags response: %w", err)
		}
		resp.Body.Close()

		for _, tag := range tags {
			if canonicalSemver(tag.Name) == "v0.0.0" {
				continue
			}
			versions[tag.Name] = struct{}{}
		}

		nextURL = parseNextLink(resp.Header.Get("Link"))
	}

	list := make([]string, 0, len(versions))
	for version := range versions {
		list = append(list, version)
	}
	sort.Slice(list, func(i, j int) bool {
		return compareSemver(list[i], list[j]) < 0
	})
	return list, nil
}

// DownloadMissingSchemas downloads missing Karpenter bundles into schemaDir.
func (d *KarpenterDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		return 0, fmt.Errorf("create schema dir: %w", err)
	}

	versions, err := d.FetchReleaseVersions()
	if err != nil {
		return 0, err
	}

	var (
		count          int
		downloadErrors []error
	)

	for _, version := range versions {
		if compareSemver(version, minKarpenterVersion) < 0 {
			continue
		}
		releaseCount, errs := d.downloadReleaseCRDs(version, schemaDir)
		count += releaseCount
		downloadErrors = append(downloadErrors, errs...)
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

func (d *KarpenterDownloader) downloadReleaseCRDs(version, schemaDir string) (int, []error) {
	targetPath := filepath.Join(schemaDir, fmt.Sprintf("karpenter-aws-%s-crds.yaml", strings.TrimPrefix(version, "v")))
	fields := logrus.Fields{
		"version":    version,
		"schemaPath": targetPath,
		"schemaDir":  schemaDir,
	}

	if exists(targetPath) {
		d.logger.WithFields(fields).WithField("status", "skipped").Info("skipping existing karpenter schema")
		return 0, nil
	}

	files, err := d.listCRDFiles(version)
	if err != nil {
		return 0, []error{err}
	}
	if len(files) == 0 {
		return 0, []error{fmt.Errorf("no CRD files found for %s", version)}
	}

	contents := make(map[string][]byte, len(files))
	for _, file := range files {
		data, err := d.downloadFile(file.DownloadURL)
		if err != nil {
			return 0, []error{err}
		}
		contents[file.Name] = data
	}

	if err := writeCRDBundle(targetPath, contents); err != nil {
		return 0, []error{err}
	}

	d.logger.WithFields(fields).Info("downloaded karpenter crds")
	return 1, nil
}

type karpenterCRDFile struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

func (d *KarpenterDownloader) listCRDFiles(version string) ([]karpenterCRDFile, error) {
	return d.fetchCRDPath(version, karpenterCRDsPath)
}

func (d *KarpenterDownloader) fetchCRDPath(version, path string) ([]karpenterCRDFile, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s?ref=%s", karpenterRepo, path, version)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request for %s: %w", version, err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return nil, fmt.Errorf("list crds for %s: %w", version, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list crds for %s: %s%s", version, resp.Status, githubAPIError(resp))
	}

	var entries []karpenterCRDFile
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, fmt.Errorf("decode crd listing for %s: %w", version, err)
	}

	var files []karpenterCRDFile
	for _, entry := range entries {
		if entry.Type != "file" {
			continue
		}
		if !strings.HasSuffix(entry.Name, ".yaml") && !strings.HasSuffix(entry.Name, ".yml") {
			continue
		}
		if entry.DownloadURL == "" {
			continue
		}
		files = append(files, entry)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
	return files, nil
}

func (d *KarpenterDownloader) downloadFile(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build file request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return nil, fmt.Errorf("fetch file %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch file %s: %s%s", url, resp.Status, githubAPIError(resp))
	}
	return io.ReadAll(resp.Body)
}
