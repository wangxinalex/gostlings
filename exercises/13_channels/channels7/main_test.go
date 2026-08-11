package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateReturnsAllValuesThroughReceiveOnlyChannel(t *testing.T) {
	values := make(chan []int, 1)
	go func() {
		values <- receiveAll(generate(2, 4, 6))
	}()

	select {
	case got := <-values:
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("receiveAll(generate()) = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generate() did not close after sending its values")
	}
}

func TestGenerateClosesEmptyReceiveOnlyChannel(t *testing.T) {
	values := make(chan []int, 1)
	go func() {
		values <- receiveAll(generate())
	}()

	select {
	case got := <-values:
		if len(got) != 0 {
			t.Fatalf("receiveAll(generate()) = %v, want no values", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generate() did not close an empty output")
	}
}
