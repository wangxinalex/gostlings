// Concept: HTTP middleware and request-scoped context values
// Task: copy X-Request-ID into a private context key and response header before calling the next handler
// Expected output: focused httptest checks pass (run `go test ./exercises/23_http/http2`)
// Hint: derive a new request that carries the value, set the response header,
//       then hand the derived request to the next handler
// Stuck?: http.Request's WithContext plus context.WithValue, keyed by a private type.

package main

import (
	"context"
	"net/http"
)

type contextKey string

const requestIDKey contextKey = "request-id"

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Read X-Request-ID, store it under requestIDKey in a derived
		//       context, set the same response header, and call next.
		ctx := context.WithValue(r.Context(), requestIDKey, "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
