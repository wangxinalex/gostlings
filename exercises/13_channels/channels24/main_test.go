package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestDrainDisablesClosedInputs(t *testing.T) {
	first := make(chan int, 2)
	second := make(chan int, 2)
	first <- 1
	first <- 3
	second <- 2
	second <- 4
	close(first)
	close(second)

	got := drain(first, second)
	sort.Ints(got)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("drain() = %v, want %v", got, want)
	}
}
