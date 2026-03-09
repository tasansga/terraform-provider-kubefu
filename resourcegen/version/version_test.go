package version

import (
	"reflect"
	"testing"
)

func TestCanonical(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "whitespace", input: "  ", want: ""},
		{name: "major", input: "1", want: "v1.0.0"},
		{name: "majorMinor", input: "1.2", want: "v1.2.0"},
		{name: "majorMinorPatch", input: "1.2.3", want: "v1.2.3"},
		{name: "leadingV", input: "v1.2.3", want: "v1.2.3"},
		{name: "leadingUpperV", input: "V1.2.3", want: "v1.2.3"},
		{name: "extraSegments", input: "1.2.3.4", want: "v1.2.3"},
		{name: "invalid", input: "nope", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Canonical(tt.input); got != tt.want {
				t.Fatalf("Canonical(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeList(t *testing.T) {
	input := []string{"1.2", "v1.2.0", "1.3.0", "  ", "v1.2.0", "1"}
	want := []string{"v1.0.0", "v1.2.0", "v1.3.0"}
	if got := NormalizeList(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeList(%v) = %v, want %v", input, got, want)
	}
}

func TestFilterIncompatible(t *testing.T) {
	tests := []struct {
		name       string
		configured []string
		supported  []string
		want       []string
	}{
		{name: "emptyConfigured", configured: nil, supported: []string{"v1.2.0"}, want: nil},
		{name: "emptySupported", configured: []string{"1.2.0"}, supported: nil, want: nil},
		{name: "allCompatible", configured: []string{"1.2", "1.3.0"}, supported: []string{"v1.2.0", "v1.3.0"}, want: nil},
		{name: "someIncompatible", configured: []string{"1.2", "1.4.0"}, supported: []string{"v1.2.0", "v1.3.0"}, want: []string{"v1.4.0"}},
		{name: "invalidConfiguredIgnored", configured: []string{"nope", "1.2.0"}, supported: []string{"v1.2.0"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterIncompatible(tt.configured, tt.supported); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("FilterIncompatible(%v, %v) = %v, want %v", tt.configured, tt.supported, got, tt.want)
			}
		})
	}
}
