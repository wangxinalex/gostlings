package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "[{bob 40} {ada 36} {eve 29}]\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestByAgeDesc(t *testing.T) {
	got := byAgeDesc([]Person{{Name: "ada", Age: 36}, {Name: "bob", Age: 40}, {Name: "eve", Age: 29}})
	want := []Person{{Name: "bob", Age: 40}, {Name: "ada", Age: 36}, {Name: "eve", Age: 29}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("byAgeDesc() = %v, want %v", got, want)
	}
}
