package main

import (
	"errors"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeResultsPreservesValuesAndErrors(t *testing.T) {
	first, second := make(chan result, 1), make(chan result, 1)
	bad := errors.New("bad job")
	first <- result{value: 2}
	second <- result{value: 5, err: bad}
	close(first)
	close(second)
	got := collect36(t, mergeResults(first, second))
	sort.Slice(got, func(i, j int) bool { return got[i].value < got[j].value })
	if len(got) != 2 || got[0].value != 2 || got[0].err != nil || got[1].value != 5 || !errors.Is(got[1].err, bad) {
		t.Fatalf("mergeResults() = %#v, want preserved value and error envelopes", got)
	}
}
func TestMergeResultsClosesForNoInputs(t *testing.T) {
	if got := collect36(t, mergeResults()); !reflect.DeepEqual(got, []result(nil)) {
		t.Fatalf("mergeResults() = %#v, want no results", got)
	}
}
func collect36(t *testing.T, out <-chan result) []result {
	t.Helper()
	var got []result
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("mergeResults() did not close")
		}
	}
}
