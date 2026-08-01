// Concept: HTTP handler contracts — method, status, headers, and JSON body
// Task: make GET /health return the JSON response and reject every other method
// Expected output: focused httptest checks pass (run `go test ./exercises/23_http/http1`)
// Hint: check r.Method before writing; set Content-Type before the first write

package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Return 405 for non-GET requests. For GET, set the JSON content type
	//       and write {"status":"ok"} followed by a newline.
	fmt.Fprintln(w, "ok")
}
