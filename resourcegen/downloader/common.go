package downloader

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

const userAgent = "resourcegen-download-schema"

func parseNextLink(linkHeader string) string {
	if linkHeader == "" {
		return ""
	}

	for _, part := range strings.Split(linkHeader, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		segments := strings.Split(part, ";")
		if len(segments) < 2 {
			continue
		}
		rel := strings.TrimSpace(segments[len(segments)-1])
		if rel != `rel="next"` {
			continue
		}
		url := strings.TrimSpace(segments[0])
		if len(url) >= 2 && url[0] == '<' && url[len(url)-1] == '>' {
			return url[1 : len(url)-1]
		}
	}
	return ""
}

// SchemaDownloader defines the subset of functionality we expect from each
// downloader implementation.
type SchemaDownloader interface {
	DownloadMissingSchemas(schemaDir string) (int, error)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func githubAPIError(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	limit := resp.Header.Get("X-RateLimit-Limit")
	remaining := resp.Header.Get("X-RateLimit-Remaining")
	reset := resp.Header.Get("X-RateLimit-Reset")
	resource := resp.Header.Get("X-RateLimit-Resource")
	retryAfter := resp.Header.Get("Retry-After")

	var details []string
	if limit != "" {
		details = append(details, "limit="+limit)
	}
	if remaining != "" {
		details = append(details, "remaining="+remaining)
	}
	if reset != "" {
		details = append(details, "reset="+reset)
	}
	if resource != "" {
		details = append(details, "resource="+resource)
	}
	if retryAfter != "" {
		details = append(details, "retry_after="+retryAfter)
	}
	if len(details) == 0 {
		return ""
	}
	return fmt.Sprintf(" (github rate limit: %s)", strings.Join(details, ", "))
}

func doRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	for {
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if !isRateLimited(resp) {
			return resp, nil
		}
		wait := rateLimitWait(resp)
		resp.Body.Close()
		if wait <= 0 {
			return nil, fmt.Errorf("github rate limit exceeded%s", githubAPIError(resp))
		}
		if wait > 10*time.Minute {
			return nil, fmt.Errorf("github rate limit exceeded; wait %s exceeds 10m%s", wait, githubAPIError(resp))
		}
		time.Sleep(wait)
	}
}

func isRateLimited(resp *http.Response) bool {
	if resp == nil {
		return false
	}
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		return true
	case http.StatusForbidden:
		remaining := strings.TrimSpace(resp.Header.Get("X-RateLimit-Remaining"))
		return remaining == "0"
	default:
		return false
	}
}

func rateLimitWait(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}
	if retryAfter := strings.TrimSpace(resp.Header.Get("Retry-After")); retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil && seconds > 0 {
			return time.Duration(seconds)*time.Second + time.Second
		}
	}
	if reset := strings.TrimSpace(resp.Header.Get("X-RateLimit-Reset")); reset != "" {
		if epoch, err := strconv.ParseInt(reset, 10, 64); err == nil {
			wait := time.Until(time.Unix(epoch, 0)) + time.Second
			if wait < 0 {
				return 0
			}
			return wait
		}
	}
	return 0
}

func compareSemver(a, b string) int {
	return semver.Compare(canonicalSemver(a), canonicalSemver(b))
}

func canonicalSemver(version string) string {
	v := strings.TrimSpace(version)
	if v == "" {
		return "v0.0.0"
	}
	v = strings.TrimPrefix(strings.TrimPrefix(v, "v"), "V")
	if v == "" {
		return "v0.0.0"
	}
	segments := strings.Split(v, ".")
	cleaned := make([]string, 0, 3)
	for _, seg := range segments {
		if seg == "" {
			continue
		}
		cleaned = append(cleaned, seg)
		if len(cleaned) == 3 {
			break
		}
	}
	for len(cleaned) < 3 {
		cleaned = append(cleaned, "0")
	}
	candidate := "v" + strings.Join(cleaned[:3], ".")
	if semver.IsValid(candidate) {
		return semver.Canonical(candidate)
	}
	return "v0.0.0"
}

func writeCRDBundle(targetPath string, contents map[string][]byte) error {
	if len(contents) == 0 {
		return fmt.Errorf("no CRD content to write")
	}

	keys := make([]string, 0, len(contents))
	for name := range contents {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	tmpFile, err := os.CreateTemp(filepath.Dir(targetPath), "schema-*.tmp")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	var buf bytes.Buffer
	for i, name := range keys {
		if i > 0 {
			buf.WriteString("\n---\n")
		}
		buf.WriteString("# Source: ")
		buf.WriteString(name)
		buf.WriteString("\n")
		data := bytes.TrimRight(contents[name], "\n")
		buf.Write(data)
		buf.WriteByte('\n')
	}

	if _, err := tmpFile.Write(buf.Bytes()); err != nil {
		tmpFile.Close()
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}
	return os.Rename(tmpFile.Name(), targetPath)
}
