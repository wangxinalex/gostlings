package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunPipelineReturnsResultsAndCloses(t *testing.T) {
	in := make(chan job, 2)
	in <- job{value: 2}
	in <- job{value: 4}
	close(in)
	results, failures := runPipeline(context.Background(), in)
	var got []int
	for len(got) < 2 {
		select {
		case value, ok := <-results:
			if !ok {
				t.Fatalf("results closed after %d values", len(got))
			}
			got = append(got, value.value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("results did not complete")
		}
	}
	select {
	case _, ok := <-results:
		if ok {
			t.Fatal("results had extra values")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("results did not close")
	}
	if len(got) != 2 || got[0] != 4 || got[1] != 8 {
		t.Fatalf("results = %v, want [4 8]", got)
	}
	select {
	case err := <-failures:
		t.Fatalf("unexpected failure: %v", err)
	default:
	}
}

func TestRunPipelineReportsFirstFailure(t *testing.T) {
	in := make(chan job, 1)
	in <- job{fail: true}
	close(in)
	results, failures := runPipeline(context.Background(), in)
	select {
	case err := <-failures:
		if !errors.Is(err, errJobFailed) {
			t.Fatalf("error = %v, want job failure", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("pipeline did not report failure")
	}
	select {
	case _, ok := <-results:
		if ok {
			for range results {
			}
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("results did not close after failure")
	}
}
