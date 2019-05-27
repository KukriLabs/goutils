package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/kukrilabs/goutils/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRequestIDMiddleware(t *testing.T) {
	tt := []struct {
		name           string
		inputRequestID string
		randomOutput   bool
	}{
		{name: "empty header", inputRequestID: "", randomOutput: true},
		{"uuid header", uuid.Must(uuid.NewV4()).String(), false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				reqID := r.Header.Get(middleware.RequestIDHeader)
				if tc.randomOutput {
					assert.NotEmpty(t, reqID, "Request ID should not be empty")
				} else {
					assert.Equal(t, tc.inputRequestID, reqID)
				}
			})

			req := httptest.NewRequest("GET", "http://example.com", nil)
			if tc.inputRequestID != "" {
				req.Header.Set(middleware.RequestIDHeader, tc.inputRequestID)
			}

			middleware.RequestID(nextHandler).ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}

// TestMultipleRequestIDMiddleware checks what happens when there is an array of requestIDs passed to the middleware. It should just take the first value for future requests
func TestMultipleRequestIDMiddleware(t *testing.T) {
	tt := []struct {
		name            string
		inputRequestIDs []string
		randomOutput    bool
	}{
		{name: "empty header", inputRequestIDs: []string{""}, randomOutput: true},
		{"uuid header", []string{uuid.Must(uuid.NewV4()).String()}, false},
		{"multi ids", []string{uuid.Must(uuid.NewV4()).String(), uuid.Must(uuid.NewV4()).String(), uuid.Must(uuid.NewV4()).String()}, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				reqID := r.Header.Get(middleware.RequestIDHeader)
				if tc.randomOutput {
					assert.NotEmpty(t, reqID, "Request ID should not be empty")
				} else {
					assert.Equal(t, tc.inputRequestIDs[0], reqID)
				}
			})

			req := httptest.NewRequest("GET", "http://example.com", nil)
			if len(tc.inputRequestIDs) > 0 {
				for _, reqID := range tc.inputRequestIDs {
					req.Header.Add(middleware.RequestIDHeader, reqID)
				}
			}

			middleware.RequestID(nextHandler).ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}
