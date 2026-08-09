package testutil

import (
	"fmt"
	"os"
	"testing"
)

func TestCaptureStdoutCapturesOutputAndRestoresStdout(t *testing.T) {
	old := os.Stdout

	got := CaptureStdout(t, func() {
		fmt.Print("captured")
	})

	if got != "captured" {
		t.Fatalf("CaptureStdout() = %q, want %q", got, "captured")
	}
	if os.Stdout != old {
		t.Fatal("CaptureStdout() did not restore os.Stdout")
	}
}
