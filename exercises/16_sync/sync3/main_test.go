package main

import (
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	return string(b)
}

func TestOutput(t *testing.T) {
	got := captureStdout(main)
	gotLines := strings.Split(strings.TrimSpace(got), "\n")
	wantLines := strings.Split(strings.TrimSpace("config initialized\nrunning\nrunning\nrunning"), "\n")
	sort.Strings(gotLines)
	sort.Strings(wantLines)
	if !reflect.DeepEqual(gotLines, wantLines) {
		t.Errorf("got %q, want lines %q", got, wantLines)
	}
}
