package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestDrainDisablesClosedInputsAfterTheirBufferedValues(t *testing.T) {
	first, second := make(chan int, 2), make(chan int, 2)
	first <- 4
	first <- 1
	second <- 3
	second <- 2
	close(first)
	close(second)

	got := drain(first, second)
	sort.Ints(got)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("drain() = %v, want exactly %v; closed inputs must not keep yielding zero values", got, want)
	}
}

func TestDrainEndsWhenEveryInputIsAlreadyClosed(t *testing.T) {
	first, second := make(chan int), make(chan int)
	close(first)
	close(second)

	returned := make(chan []int, 1)
	go func() { returned <- drain(first, second) }()

	select {
	case got := <-returned:
		if len(got) != 0 {
			t.Fatalf("drain() = %v, want no values from closed empty inputs", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("drain() kept selecting closed inputs instead of retiring them")
	}
}

func TestDrainKeepsReadingTheOpenInput(t *testing.T) {
	closed, open := make(chan int), make(chan int, 1)
	close(closed)
	open <- 7
	close(open)

	returned := make(chan []int, 1)
	go func() { returned <- drain(closed, open) }()

	select {
	case got := <-returned:
		if !reflect.DeepEqual(got, []int{7}) {
			t.Fatalf("drain() = %v, want [7] from the input that was still open", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("drain() stopped before the open input was drained")
	}
}
