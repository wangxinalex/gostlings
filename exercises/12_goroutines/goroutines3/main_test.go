package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRunWithArgsPreservesEveryLabel(t *testing.T) {
	got := runWithArgs([]string{"north", "south", "east"})
	want := []string{"east", "north", "south"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runWithArgs() = %v, want %v", got, want)
	}
}

func TestRunWithArgsHandlesEmptyInput(t *testing.T) {
	if got := runWithArgs(nil); len(got) != 0 {
		t.Fatalf("runWithArgs(nil) returned %d labels, want 0", len(got))
	}
}
