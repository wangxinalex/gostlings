package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateSendsValuesAndCloses(t *testing.T) {
	ch := generate(1, 2, 3)
	var got []int

	for {
		select {
		case value, ok := <-ch:
			if !ok {
				if !reflect.DeepEqual(got, []int{1, 2, 3}) {
					t.Fatalf("values = %v, want %v", got, []int{1, 2, 3})
				}
				return
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("generate did not close its output")
		}
	}
}

func TestSumDrainsUntilInputCloses(t *testing.T) {
	in := make(chan int)
	go func() {
		defer close(in)
		in <- 2
		in <- 4
	}()

	result := make(chan int, 1)
	go func() {
		result <- sum(in)
	}()

	select {
	case got := <-result:
		if got != 6 {
			t.Fatalf("sum() = %d, want 6", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("sum did not return after input closed")
	}
}
