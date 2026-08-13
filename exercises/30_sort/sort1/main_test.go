package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[ada bob eva]\n[1 2 3 7]\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSortNamesAndNumbers(t *testing.T) {
	if got := sortNames([]string{"bob", "ada", "eva"}); !reflect.DeepEqual(got, []string{"ada", "bob", "eva"}) {
		t.Fatalf("sortNames() = %v, want [ada bob eva]", got)
	}
	if got := sortNumbers([]int{7, 1, 3, 2}); !reflect.DeepEqual(got, []int{1, 2, 3, 7}) {
		t.Fatalf("sortNumbers() = %v, want [1 2 3 7]", got)
	}
}
