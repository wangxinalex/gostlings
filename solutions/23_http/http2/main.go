// Concept: HTTP middleware and request-scoped context values
// Task: copy X-Request-ID into a private context key and response header before calling the next handler
// Expected output: focused httptest checks pass (run `go test ./solutions/23_http/http2`)
// Hint: use r.WithContext(context.WithValue(...)) and call next.ServeHTTP with the new request

package main

import (
	"context"
	"net/http"
)

type contextKey string

const requestIDKey contextKey = "request-id"

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
