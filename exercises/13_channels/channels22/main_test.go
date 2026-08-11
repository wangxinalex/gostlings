package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeForwardsEveryInputAndCloses(t *testing.T) {
	first := buffered(1, 4)
	second := buffered(2, 3)
	got := collectMerged(t, merge(first, second))
	sort.Ints(got)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}

func TestMergeWithZeroInputsCloses(t *testing.T) {
	got := collectMerged(t, merge())
	if len(got) != 0 {
		t.Fatalf("merge() = %v, want no values", got)
	}
}

func TestMergeDoesNotIgnoreABlockedInput(t *testing.T) {
	blocked := make(chan int)
	out := merge(buffered(2), blocked)

	sent := make(chan struct{})
	go func() {
		blocked <- 3
		close(blocked)
		close(sent)
	}()

	got := collectMerged(t, out)
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not receive from its blocked input")
	}
	sort.Ints(got)
	if want := []int{2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}

func TestMergeClosesOnlyAfterEveryForwarderExits(t *testing.T) {
	previous := onForwarderExit
	exited := make(chan struct{}, 2)
	release := make(chan struct{})
	onForwarderExit = func() {
		exited <- struct{}{}
		<-release
	}
	t.Cleanup(func() { onForwarderExit = previous })

	out := merge(buffered(), buffered())
	for forwarder := 0; forwarder < 2; forwarder++ {
		select {
		case <-exited:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("forwarder %d did not reach its exit hook", forwarder+1)
		}
	}
	select {
	case _, ok := <-out:
		if !ok {
			t.Fatal("merge() closed before every forwarder exited")
		}
		t.Fatal("merge() produced an unexpected value")
	default:
	}

	close(release)
	if got := collectMerged(t, out); len(got) != 0 {
		t.Fatalf("merge() = %v, want no values", got)
	}
}

func buffered(values ...int) <-chan int {
	in := make(chan int, len(values))
	for _, value := range values {
		in <- value
	}
	close(in)
	return in
}

func collectMerged(t *testing.T, out <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case value, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, value)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("merge() did not close")
		}
	}
}
