package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, this is a GET response!"))
			t.Logf("Received GET request for %s\n", r.URL.Path)
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Can't read body", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, this is a POST response!"))
			t.Logf("Received POST request for %s with data: %s\n", r.URL.Path, string(body))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	tests := []struct {
		method         string
		target         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			method:         "GET",
			target:         "/",
			body:           "",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, this is a GET response!",
		},
		{
			method:         "POST",
			target:         "/",
			body:           "Sample data",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, this is a POST response!",
		},
		{
			method:         "PUT",
			target:         "/",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.target, strings.NewReader(test.body))
		w := httptest.NewRecorder()

		handler(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != test.expectedStatus {
			t.Errorf("Expected status %d, got %d", test.expectedStatus, resp.StatusCode)
		}
		if string(body) != test.expectedBody {
			t.Errorf("Expected body %q, got %q", test.expectedBody, string(body))
		}
	}
}
