package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "go,rust,python\ngo rust python\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitParts(t *testing.T) {
	got := splitParts("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("splitParts() = %v, want %v", got, want)
	}
}
