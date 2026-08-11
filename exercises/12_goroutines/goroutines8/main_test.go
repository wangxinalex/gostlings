package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRunWorkersWithInputReturnsOneResultPerJob(t *testing.T) {
	got := runWorkersWithInput([]int{7, 1, 4})
	want := []string{"job 1 received", "job 4 received", "job 7 received"}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runWorkersWithInput() = %v, want %v", got, want)
	}
}
