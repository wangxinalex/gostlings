package main

import (
	"bytes"
	"testing"
)

func TestReuseBufferReturnsWrittenValueAndPutsItBack(t *testing.T) {
	previousGet, previousPut := getBuffer, putBuffer
	buffer := new(bytes.Buffer)
	buffer.WriteString("stale")
	gotPut := false
	getBuffer = func() *bytes.Buffer { return buffer }
	putBuffer = func(got *bytes.Buffer) {
		if got != buffer {
			t.Errorf("put a different buffer")
		}
		gotPut = true
	}
	t.Cleanup(func() { getBuffer, putBuffer = previousGet, previousPut })
	if got := reuseBuffer("fresh"); got != "fresh" {
		t.Fatalf("reuseBuffer() = %q, want fresh", got)
	}
	if !gotPut {
		t.Fatal("reuseBuffer() did not return the buffer to the pool")
	}
}
