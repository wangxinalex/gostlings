package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "rect 6.00\ncircle 12.56\ntotal 18.56\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTotalAreaSumsAHeterogeneousSlice(t *testing.T) {
	got := totalArea([]Shape{Rect{W: 4, H: 5}, Circle{R: 1}, Rect{W: 1, H: 1}})
	const want = 24.14
	if diff := got - want; diff > 1e-9 || diff < -1e-9 {
		t.Fatalf("totalArea() = %v, want %v", got, want)
	}
}
