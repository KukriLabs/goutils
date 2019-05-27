package middleware

import (
	"net/http"

	"github.com/gofrs/uuid"
)

const (
	RequestIDHeader = "X-Request-ID"
)

// RequestID adds a X-Request-ID header if not present
func RequestID() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get(RequestIDHeader) == "" {
				r.Header.Set(RequestIDHeader, uuid.Must(uuid.NewV4()).String())
			}

			next.ServeHTTP(w, r)
		})
	}
}
