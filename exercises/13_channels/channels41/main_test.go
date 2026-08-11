package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestDrainDisablesAlreadyClosedInputsAfterTheirBufferedValues(t *testing.T) {
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
