package resourcegen

import (
	"strings"
	"unicode"
)

func toSnakeCase(input string) string {
	runes := []rune(input)
	var b []rune
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 && (shouldInsertUnderscore(runes, i)) {
				b = append(b, '_')
			}
			b = append(b, unicode.ToLower(r))
			continue
		}
		b = append(b, r)
	}
	return string(b)
}

func shouldInsertUnderscore(runes []rune, i int) bool {
	if i == 0 {
		return false
	}
	prev := runes[i-1]
	nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
	return unicode.IsLower(prev) || nextLower
}

func versionSegment(version string) string {
	if version == "" {
		return "v1"
	}
	return toSnakeCase(version)
}

func versionCamel(version string) string {
	if version == "" {
		return "V1"
	}
	runes := []rune(version)
	var segments []string
	var current []rune
	prevType := 0
	for _, r := range runes {
		t := 0
		switch {
		case unicode.IsDigit(r):
			t = 2
		case unicode.IsLetter(r):
			t = 1
		default:
			t = 3
		}
		if t == 3 {
			if len(current) > 0 {
				segments = append(segments, string(current))
				current = current[:0]
			}
			prevType = 0
			continue
		}
		if prevType != 0 && t != prevType {
			segments = append(segments, string(current))
			current = current[:0]
		}
		current = append(current, unicode.ToLower(r))
		prevType = t
	}
	if len(current) > 0 {
		segments = append(segments, string(current))
	}
	var b strings.Builder
	for _, seg := range segments {
		if seg == "" {
			continue
		}
		b.WriteString(strings.Title(seg))
	}
	if b.Len() == 0 {
		return "V1"
	}
	return b.String()
}

func providerSegment(provider string) string {
	name := toSnakeCase(sanitizeProviderInput(provider))
	if name == "" {
		return "kubefu"
	}
	return name
}

func providerCamel(provider string) string {
	name := providerSegment(provider)
	if name == "" {
		return "Kubefu"
	}
	segments := strings.Split(name, "_")
	var b strings.Builder
	for _, seg := range segments {
		if seg == "" {
			continue
		}
		b.WriteString(capitalizeSegment(seg))
	}
	if b.Len() == 0 {
		return "Kubefu"
	}
	return b.String()
}

func capitalizeSegment(segment string) string {
	if segment == "" {
		return ""
	}
	runes := []rune(segment)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func groupSegment(group string) string {
	if group == "" {
		return "core"
	}
	name := toSnakeCase(sanitizeProviderInput(group))
	if name == "" {
		return "core"
	}
	return name
}

func groupCamel(group string) string {
	name := groupSegment(group)
	if name == "" {
		return "Core"
	}
	segments := strings.Split(name, "_")
	var b strings.Builder
	for _, seg := range segments {
		if seg == "" {
			continue
		}
		b.WriteString(capitalizeSegment(seg))
	}
	if b.Len() == 0 {
		return "Core"
	}
	return b.String()
}

func sanitizeProviderInput(input string) string {
	if input == "" {
		return ""
	}
	var b strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			continue
		}
		b.WriteRune('_')
	}
	return b.String()
}
