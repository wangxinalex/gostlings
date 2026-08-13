package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "true false\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestIsHex(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"3f", true},
		{"A0", true},
		{"xyz", false},
		{"abc", false},
		{"3", false},
	}
	for _, c := range cases {
		if got := isHex(c.in); got != c.want {
			t.Errorf("isHex(%q) = %t, want %t", c.in, got, c.want)
		}
	}
}
