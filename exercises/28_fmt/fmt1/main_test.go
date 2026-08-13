package main

import (
	"gostlings/internal/testutil"
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	got := testutil.CaptureStdout(t, main)
	const want = "{Ada 36}\n{Name:Ada Age:36}\nmain.Person{Name:\"Ada\", Age:36}\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormats(t *testing.T) {
	got := formats(Person{Name: "Ada", Age: 36})
	want := []string{"{Ada 36}", "{Name:Ada Age:36}", "main.Person{Name:\"Ada\", Age:36}"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("formats() = %v, want %v", got, want)
	}
}
