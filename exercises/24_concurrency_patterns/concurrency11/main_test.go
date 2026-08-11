package main

import (
	"context"
	"errors"
	"testing"
)

func TestOrderedReturnsInputOrder(t *testing.T) {
	jobs := []orderedJob{{value: 3}, {value: 1}, {value: 2}}
	got, err := ordered(context.Background(), 2, jobs)
	if err != nil || len(got) != 3 || got[0] != 6 || got[1] != 2 || got[2] != 4 {
		t.Fatalf("ordered() = (%v, %v)", got, err)
	}
}

func TestOrderedStopsOnFailure(t *testing.T) {
	_, err := ordered(context.Background(), 2, []orderedJob{{value: 1}, {fail: true}})
	if !errors.Is(err, errOrderedJob) {
		t.Fatalf("error = %v, want ordered job failure", err)
	}
}
