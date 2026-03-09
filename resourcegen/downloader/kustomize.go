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
	kustomizeRepo       = "kubernetes-sigs/kustomize"
	kustomizeTagsURL    = "https://api.github.com/repos/" + kustomizeRepo + "/tags?per_page=100"
	kustomizeSchemaPath = "kyaml/openapi/kustomizationapi/swagger.json"
	minKustomizeVersion = "0.3.1"
)

var kustomizeReleaseRE = regexp.MustCompile(`^kyaml/v(\d+\.\d+\.\d+)$`)

// KustomizeDownloader downloads local Kustomization OpenAPI schemas from tagged releases.
type KustomizeDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewKustomizeDownloader builds a downloader for Kustomize releases.
func NewKustomizeDownloader(logger logrus.FieldLogger) *KustomizeDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &KustomizeDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions lists valid Kustomize release tags.
func (d *KustomizeDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := kustomizeTagsURL

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
			if matches := kustomizeReleaseRE.FindStringSubmatch(tag.Name); matches != nil {
				versions[matches[1]] = struct{}{}
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

// DownloadMissingSchemas fetches Kustomize schemas into schemaDir/kustomize-<version>.json.
func (d *KustomizeDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
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
		if compareSemver(version, minKustomizeVersion) < 0 {
			continue
		}
		targetPath := filepath.Join(schemaDir, fmt.Sprintf("kustomize-%s.json", version))
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

func (d *KustomizeDownloader) downloadSchema(version, targetPath string) (bool, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/kyaml/v%s/%s", kustomizeRepo, version, kustomizeSchemaPath)
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
