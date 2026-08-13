package main

import (
	"gostlings/internal/testutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "exit 3\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExitCode(t *testing.T) {
	if code, err := exitCode(exec.Command("echo", "ok")); err != nil || code != 0 {
		t.Fatalf("exitCode(success) = (%d, %v), want (0, nil)", code, err)
	}
	if code, err := exitCode(exec.Command("sh", "-c", "exit 3")); err != nil || code != 3 {
		t.Fatalf("exitCode(exit 3) = (%d, %v), want (3, nil)", code, err)
	}
	if _, err := exitCode(exec.Command("definitely-not-a-real-command")); err == nil {
		t.Fatal("exitCode() returned nil for a command that could not start")
	}
}
