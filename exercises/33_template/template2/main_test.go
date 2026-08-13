package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "go,rust,python\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRender(t *testing.T) {
	got, err := render([]string{"go", "rust", "python"})
	if err != nil {
		t.Fatal(err)
	}
	if got != "go,rust,python" {
		t.Fatalf("render() = %q, want %q", got, "go,rust,python")
	}
}

func TestRenderEmptyList(t *testing.T) {
	got, err := render(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("render(nil) = %q, want %q", got, "")
	}
}
