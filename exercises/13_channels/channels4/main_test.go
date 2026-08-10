package main

import "testing"

func TestReadReportsClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	got, ok := read(ch)
	if got != 0 || ok {
		t.Fatalf("read(closed channel) = (%d, %v), want (0, false)", got, ok)
	}
}

func TestReadDistinguishesBufferedZeroValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 0
	close(ch)

	got, ok := read(ch)
	if got != 0 || !ok {
		t.Fatalf("read(buffered zero) = (%d, %v), want (0, true)", got, ok)
	}
}
