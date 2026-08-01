// Concept: HTTP handler contracts — method, status, headers, and JSON body
// Task: make GET /health return the JSON response and reject every other method
// Expected output: focused httptest checks pass (run `go test ./solutions/23_http/http1`)
// Hint: check r.Method before writing; set Content-Type before the first write

package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"ok"}`)
}
