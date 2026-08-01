package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestFetchGreetingReturnsBody(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("hello")),
			Header:     make(http.Header),
			Request:    r,
		}, nil
	})}

	got, err := fetchGreeting(client, "http://example.test")
	if err != nil {
		t.Fatalf("fetchGreeting() error = %v", err)
	}
	if got != "hello" {
		t.Fatalf("fetchGreeting() = %q, want %q", got, "hello")
	}
}

func TestFetchGreetingReportsNon2xx(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(strings.NewReader("unavailable")),
			Header:     make(http.Header),
			Request:    r,
		}, nil
	})}

	_, err := fetchGreeting(client, "http://example.test")
	if err == nil {
		t.Fatal("fetchGreeting() error = nil, want non-2xx error")
	}
	if !strings.Contains(err.Error(), "503") {
		t.Fatalf("error = %q, want it to contain status 503", err)
	}
}

func TestFetchGreetingClosesResponseBody(t *testing.T) {
	body := &trackingBody{Reader: strings.NewReader("hello")}
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
			Header:     make(http.Header),
			Request:    r,
		}, nil
	})}

	if _, err := fetchGreeting(client, "http://example.test"); err != nil {
		t.Fatalf("fetchGreeting() error = %v", err)
	}
	if !body.closed {
		t.Fatal("fetchGreeting() did not close response body")
	}
}

type trackingBody struct {
	io.Reader
	closed bool
}

func (b *trackingBody) Close() error {
	b.closed = true
	return nil
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
