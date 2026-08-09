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
	wantLines := strings.Split(strings.TrimSpace("0\n1\n2"), "\n")
	sort.Strings(gotLines)
	sort.Strings(wantLines)
	if !reflect.DeepEqual(gotLines, wantLines) {
		t.Errorf("got %q, want lines %q", got, wantLines)
	}
}
