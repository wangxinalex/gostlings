package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRunLabelsKeepsEveryLoopValue(t *testing.T) {
	got := runLabels([]string{"red", "green", "blue"})
	want := []string{"blue", "green", "red"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runLabels() = %v, want every label once", got)
	}
}
