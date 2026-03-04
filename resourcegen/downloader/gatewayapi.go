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
	gatewayAPIRepo        = "kubernetes-sigs/gateway-api"
	gatewayAPITagsURL     = "https://api.github.com/repos/" + gatewayAPIRepo + "/tags?per_page=100"
	gatewayAPIStandardURL = "https://github.com/" + gatewayAPIRepo + "/releases/download/%s/standard-install.yaml"
	minGatewayAPIVersion  = "0.5.0"
)

var gatewayAPIReleaseRE = regexp.MustCompile(`^v(\d+\.\d+\.\d+)$`)

// GatewayAPIDownloader downloads Gateway API CRD bundles.
type GatewayAPIDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewGatewayAPIDownloader creates a downloader for Gateway API releases.
func NewGatewayAPIDownloader(logger logrus.FieldLogger) *GatewayAPIDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &GatewayAPIDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions lists Gateway API release tags.
func (d *GatewayAPIDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := gatewayAPITagsURL

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
			if matches := gatewayAPIReleaseRE.FindStringSubmatch(tag.Name); matches != nil {
				versions[matches[0]] = struct{}{}
			}
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

// DownloadMissingSchemas downloads missing Gateway API bundles into schemaDir.
func (d *GatewayAPIDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
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
		if compareSemver(version, minGatewayAPIVersion) < 0 {
			continue
		}
		releaseCount, errs := d.downloadReleaseBundle(version, schemaDir)
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

func (d *GatewayAPIDownloader) downloadReleaseBundle(version, schemaDir string) (int, []error) {
	url := fmt.Sprintf(gatewayAPIStandardURL, version)
	targetPath := filepath.Join(schemaDir, fmt.Sprintf("gateway-api-%s-standard.yaml", strings.TrimPrefix(version, "v")))

	fields := logrus.Fields{
		"version":    version,
		"schemaPath": targetPath,
		"schemaDir":  schemaDir,
		"url":        url,
	}

	if exists(targetPath) {
		d.logger.WithFields(fields).WithField("status", "skipped").Info("skipping existing Gateway API schema")
		return 0, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, []error{fmt.Errorf("build bundle request for %s: %w", version, err)}
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return 0, []error{fmt.Errorf("fetch bundle for %s: %w", version, err)}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		d.logger.WithFields(fields).Warn("bundle not found for version, skipping")
		return 0, nil
	default:
		return 0, []error{fmt.Errorf("bundle %s: %s%s", version, resp.Status, githubAPIError(resp))}
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(targetPath), "schema-*.tmp")
	if err != nil {
		return 0, []error{err}
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return 0, []error{err}
	}
	if err := tmpFile.Close(); err != nil {
		return 0, []error{err}
	}

	if err := os.Rename(tmpFile.Name(), targetPath); err != nil {
		return 0, []error{err}
	}

	d.logger.WithFields(fields).WithField("status", "downloaded").Info("downloaded Gateway API schema")
	return 1, nil
}
