// Concept: HTTP client lifecycle and response status handling
// Task: return a 2xx response body, return an error for non-2xx responses, and always close the body
// Expected output: focused httptest checks pass (run `go test ./exercises/23_http/http3`)
// Hint: close the response body once the request succeeds, and decide on the
//       status before reading the body; the error must name the status code

package main

import "net/http"

func fetchGreeting(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	// TODO: Close resp.Body, reject non-2xx statuses with an error containing
	//       the status code, then read and return the response body.
	return "", nil
}
