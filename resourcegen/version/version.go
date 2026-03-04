package version

import (
	"regexp"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

var versionPattern = regexp.MustCompile(`\d+(?:\.\d+){0,2}`)

// FromFilename extracts the first version-looking substring from the provided
// file name and returns its canonical semver representation.
func FromFilename(name string) string {
	return Canonical(versionPattern.FindString(name))
}

// Canonical converts the input into a semver string (like v1.20.0). Inputs such
// as "1.20" become "v1.20.0". Empty or invalid strings yield "".
func Canonical(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}
	input = strings.TrimPrefix(strings.TrimPrefix(input, "v"), "V")
	if input == "" {
		return ""
	}
	segments := strings.Split(input, ".")
	var cleaned []string
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
	return ""
}

// NormalizeList cleans, deduplicates, and sorts all versions in the slice.
func NormalizeList(versions []string) []string {
	set := make(map[string]struct{}, len(versions))
	for _, version := range versions {
		if canonical := Canonical(version); canonical != "" {
			set[canonical] = struct{}{}
		}
	}
	result := make([]string, 0, len(set))
	for version := range set {
		result = append(result, version)
	}
	sort.Slice(result, func(i, j int) bool {
		return semver.Compare(result[i], result[j]) < 0
	})
	return result
}

// IsCompatibleWith returns true if the configured version is empty, if there
// are no supported versions, or if the trimmed version matches one of the known
// canonical versions.
func IsCompatibleWith(configured string, supported []string) bool {
	if configured == "" || len(supported) == 0 {
		return true
	}
	canonical := Canonical(configured)
	if canonical == "" {
		return true
	}
	for _, version := range supported {
		if canonical == version {
			return true
		}
	}
	return false
}
