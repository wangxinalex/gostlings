package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[123 45]\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindNumbers(t *testing.T) {
	if got := findNumbers("a=123 b=45"); !reflect.DeepEqual(got, []string{"123", "45"}) {
		t.Fatalf("findNumbers() = %v, want [123 45]", got)
	}
	if got := findNumbers("no digits"); len(got) != 0 {
		t.Fatalf("findNumbers() = %v, want no matches", got)
	}
}
