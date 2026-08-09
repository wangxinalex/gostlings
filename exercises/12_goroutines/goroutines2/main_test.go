package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	gotLines := strings.Split(strings.TrimSpace(got), "\n")
	wantLines := strings.Split(strings.TrimSpace("worker 0 done\nworker 1 done\nworker 2 done"), "\n")
	sort.Strings(gotLines)
	sort.Strings(wantLines)
	if !reflect.DeepEqual(gotLines, wantLines) {
		t.Errorf("got %q, want lines %q", got, wantLines)
	}
}
