// Concept: HTTP client lifecycle and response status handling
// Task: return a 2xx response body, return an error for non-2xx responses, and always close the body
// Expected output: focused httptest checks pass (run `go test ./solutions/23_http/http3`)
// Hint: defer resp.Body.Close() immediately after a successful client.Do/Get; check StatusCode before accepting the body

package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchGreeting(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
