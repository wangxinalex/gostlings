package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateForwardsValuesAndCloses(t *testing.T) {
	out := generateWithWatchdog(t, 1, 2, 3)
	var got []int
	deadline := time.After(500 * time.Millisecond)

	for {
		select {
		case value, ok := <-out:
			if !ok {
				want := []int{1, 2, 3}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("generate() = %v, want %v", got, want)
				}
				return
			}
			got = append(got, value)
		case <-deadline:
			t.Fatal("generate() did not close its output")
		}
	}
}

func TestGenerateWithNoValuesClosesPromptly(t *testing.T) {
	out := generateWithWatchdog(t)
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("generate() returned an unexpected value")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generate() did not close an empty output")
	}
}

func generateWithWatchdog(t *testing.T, values ...int) <-chan int {
	t.Helper()

	returned := make(chan (<-chan int), 1)
	go func() {
		returned <- generate(values...)
	}()

	select {
	case out := <-returned:
		return out
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generate() did not return its output channel")
		return nil
	}
}
