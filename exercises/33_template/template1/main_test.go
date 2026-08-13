package main

import (
	"gostlings/internal/testutil"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "Hello, Ada!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRender(t *testing.T) {
	got, err := render("Hello, {{.Name}}!", person{Name: "Ada"})
	if err != nil {
		t.Fatal(err)
	}
	if got != "Hello, Ada!" {
		t.Fatalf("render() = %q, want %q", got, "Hello, Ada!")
	}
}

func TestRenderRejectsBrokenTemplate(t *testing.T) {
	if _, err := render("{{.Name", person{Name: "Ada"}); err == nil {
		t.Fatal("render() accepted a broken template")
	}
}
