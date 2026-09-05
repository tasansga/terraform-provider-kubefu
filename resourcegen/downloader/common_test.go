package downloader

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestApplyAuthHeader_WithGHToken(t *testing.T) {
	t.Setenv("GH_TOKEN", "test-token-123")

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/kubernetes/kubernetes/branches", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	applyAuthHeader(req)

	got := req.Header.Get("Authorization")
	want := "Bearer test-token-123"
	if got != want {
		t.Errorf("expected Authorization header %q, got %q", want, got)
	}
}

func TestApplyAuthHeader_WithoutGHToken(t *testing.T) {
	t.Setenv("GH_TOKEN", "")

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/kubernetes/kubernetes/branches", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	applyAuthHeader(req)

	if got := req.Header.Get("Authorization"); got != "" {
		t.Errorf("expected empty Authorization header, got %q", got)
	}
}

func TestApplyAuthHeader_NonAPIHost(t *testing.T) {
	t.Setenv("GH_TOKEN", "test-token-123")

	req, err := http.NewRequest(http.MethodGet, "https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.30/swagger.json", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	applyAuthHeader(req)

	if got := req.Header.Get("Authorization"); got != "" {
		t.Errorf("expected empty Authorization header for non-api host, got %q", got)
	}
}

func TestApplyAuthHeader_PreserveExistingHeader(t *testing.T) {
	t.Setenv("GH_TOKEN", "test-token-123")

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/repos/kubernetes/kubernetes/branches", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer existing-token")

	applyAuthHeader(req)

	got := req.Header.Get("Authorization")
	want := "Bearer existing-token"
	if got != want {
		t.Errorf("expected Authorization header %q, got %q", want, got)
	}
}

func TestApplyAuthHeader_NilSafety(t *testing.T) {
	t.Setenv("GH_TOKEN", "test-token-123")

	// Should not panic on nil request or nil URL
	applyAuthHeader(nil)
	applyAuthHeader(&http.Request{})
}

func TestDoRequest_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := doRequest(server.Client(), req)
	if err != nil {
		t.Fatalf("unexpected error from doRequest: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestIsRateLimited(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		headers    map[string]string
		want       bool
	}{
		{
			name:       "too many requests status",
			statusCode: http.StatusTooManyRequests,
			want:       true,
		},
		{
			name:       "forbidden with remaining 0",
			statusCode: http.StatusForbidden,
			headers:    map[string]string{"X-RateLimit-Remaining": "0"},
			want:       true,
		},
		{
			name:       "forbidden with remaining > 0",
			statusCode: http.StatusForbidden,
			headers:    map[string]string{"X-RateLimit-Remaining": "10"},
			want:       false,
		},
		{
			name:       "status ok",
			statusCode: http.StatusOK,
			want:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tc.statusCode,
				Header:     make(http.Header),
			}
			for k, v := range tc.headers {
				resp.Header.Set(k, v)
			}
			if got := isRateLimited(resp); got != tc.want {
				t.Errorf("isRateLimited() = %v, want %v", got, tc.want)
			}
		})
	}

	if isRateLimited(nil) {
		t.Errorf("isRateLimited(nil) = true, want false")
	}
}

func TestRateLimitWait(t *testing.T) {
	t.Run("nil response", func(t *testing.T) {
		if got := rateLimitWait(nil); got != 0 {
			t.Errorf("rateLimitWait(nil) = %v, want 0", got)
		}
	})

	t.Run("Retry-After header", func(t *testing.T) {
		resp := &http.Response{
			Header: http.Header{
				"Retry-After": []string{"5"},
			},
		}
		got := rateLimitWait(resp)
		want := 6 * time.Second
		if got != want {
			t.Errorf("rateLimitWait() = %v, want %v", got, want)
		}
	})

	t.Run("X-RateLimit-Reset header", func(t *testing.T) {
		resetEpoch := time.Now().Add(10 * time.Second).Unix()
		resp := &http.Response{
			Header: http.Header{
				"X-RateLimit-Reset": []string{strconv.FormatInt(resetEpoch, 10)},
			},
		}
		got := rateLimitWait(resp)
		if got <= 0 || got > 15*time.Second {
			t.Errorf("rateLimitWait() returned unexpected duration: %v", got)
		}
	})
}
