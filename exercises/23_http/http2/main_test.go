package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithRequestIDPropagatesToHandlerAndResponse(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(requestIDKey).(string)
		if !ok {
			t.Fatal("request ID missing from context")
		}
		fmt.Fprint(w, id)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Request-ID", "req-123")
	rec := httptest.NewRecorder()

	withRequestID(next).ServeHTTP(rec, req)

	if got := rec.Header().Get("X-Request-ID"); got != "req-123" {
		t.Fatalf("response X-Request-ID = %q, want %q", got, "req-123")
	}
	if got := rec.Body.String(); got != "req-123" {
		t.Fatalf("body = %q, want %q", got, "req-123")
	}
}
