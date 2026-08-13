package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[python go c]\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestByLength(t *testing.T) {
	got := byLength([]string{"ab", "cd", "e", "xyz"})
	want := []string{"xyz", "ab", "cd", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("byLength() = %v, want %v", got, want)
	}
}
