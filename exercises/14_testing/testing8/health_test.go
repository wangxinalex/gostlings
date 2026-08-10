// Concept: testing HTTP handlers without opening a real port
// Task: use httptest.NewRequest and httptest.NewRecorder to verify status and body
// Expected output: PASS (run with `go test ./exercises/14_testing/testing8`)
// Hint: create a GET request for /health, create a recorder, call Handler(recorder, request),
//       then assert Code is http.StatusOK and Body.String() is "ok\n".

package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()
	Handler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if got := recorder.Body.String(); got != "ok\n" {
		t.Errorf("body = %q, want %q", got, "ok\n")
	}
	// TODO: Remove this guard after the status and body assertions are complete.
	t.Fatal("TODO: finish the httptest assertions")
}
