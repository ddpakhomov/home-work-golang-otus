package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		method         string
		target         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{"GET", "/", "", http.StatusOK, "Hello, this is a GET response!"},
		{"POST", "/", "test data", http.StatusOK, "Hello, this is a POST response!"},
		{"PUT", "/", "", http.StatusMethodNotAllowed, "Method not allowed\n"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.target, strings.NewReader(tt.body))
		w := httptest.NewRecorder()
		handler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tt.expectedStatus {
			t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("could not read response body: %v", err)
		}

		if string(body) != tt.expectedBody {
			t.Errorf("expected body %q, got %q", tt.expectedBody, body)
		}
	}
}
