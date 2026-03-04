package downloader

import (
	"bufio"
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
	fluxRepo              = "fluxcd/flux2"
	fluxTagsURL           = "https://api.github.com/repos/" + fluxRepo + "/tags?per_page=100"
	fluxKustomizationPath = "manifests/crds/kustomization.yaml"
	minFluxVersion        = "0.12.2"
)

var fluxReleaseRE = regexp.MustCompile(`^v(\d+\.\d+\.\d+)$`)

// FluxDownloader downloads Flux v2 CRDs referenced by each release.
type FluxDownloader struct {
	client *http.Client
	logger logrus.FieldLogger
}

// NewFluxDownloader builds a downloader for Flux releases.
func NewFluxDownloader(logger logrus.FieldLogger) *FluxDownloader {
	if logger == nil {
		logger = logrus.New()
	}
	return &FluxDownloader{
		client: http.DefaultClient,
		logger: logger,
	}
}

// FetchReleaseVersions lists valid Flux release tags.
func (d *FluxDownloader) FetchReleaseVersions() ([]string, error) {
	versions := map[string]struct{}{}
	nextURL := fluxTagsURL

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
			if matches := fluxReleaseRE.FindStringSubmatch(tag.Name); matches != nil {
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

// DownloadMissingSchemas fetches Flux CRD YAMLs into schemaDir/flux/<release>.
func (d *FluxDownloader) DownloadMissingSchemas(schemaDir string) (int, error) {
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
		if compareSemver(version, minFluxVersion) < 0 {
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

func (d *FluxDownloader) downloadReleaseCRDs(version, schemaDir string) (int, []error) {
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		return 0, []error{fmt.Errorf("create schema dir %s: %w", schemaDir, err)}
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", fluxRepo, version, fluxKustomizationPath)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, []error{fmt.Errorf("build kustomization request for %s: %w", version, err)}
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doRequest(d.client, req)
	if err != nil {
		return 0, []error{fmt.Errorf("fetch kustomization for %s: %w", version, err)}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		d.logger.WithFields(logrus.Fields{
			"version": version,
			"url":     url,
		}).Warn("kustomization file not found for version, skipping")
		return 0, nil
	default:
		return 0, []error{fmt.Errorf("kustomization %s: %s%s", version, resp.Status, githubAPIError(resp))}
	}

	scanner := bufio.NewScanner(resp.Body)
	var (
		urls []string
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.HasPrefix(line, "-") {
			continue
		}
		resource := strings.TrimSpace(strings.TrimPrefix(line, "-"))
		if resource == "" {
			continue
		}
		urls = append(urls, resource)
	}
	if err := scanner.Err(); err != nil {
		return 0, []error{fmt.Errorf("read kustomization for %s: %w", version, err)}
	}

	var (
		count int
		errs  []error
	)
	seen := map[string]struct{}{}
	for _, resourceURL := range urls {
		if _, ok := seen[resourceURL]; ok {
			continue
		}
		seen[resourceURL] = struct{}{}

		filename := filepath.Base(resourceURL)
		targetPath := filepath.Join(schemaDir, fmt.Sprintf("flux-%s-%s", strings.TrimPrefix(version, "v"), filename))
		fields := logrus.Fields{
			"version":    version,
			"schemaPath": targetPath,
			"schemaDir":  schemaDir,
			"resource":   resourceURL,
		}

		if exists(targetPath) {
			d.logger.WithFields(fields).WithField("status", "skipped").Info("skipping existing flux schema")
			continue
		}

		if err := d.downloadFile(resourceURL, targetPath); err != nil {
			d.logger.WithFields(fields).WithError(err).Error("failed to download flux schema")
			errs = append(errs, fmt.Errorf("version %s resource %s: %w", version, resourceURL, err))
			continue
		}

		d.logger.WithFields(fields).WithField("status", "downloaded").Info("downloaded flux schema")
		count++
	}

	return count, errs
}

func (d *FluxDownloader) downloadFile(url, targetPath string) error {
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
