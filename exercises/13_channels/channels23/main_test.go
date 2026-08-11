package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeForwardsInputsBeforeTheyClose(t *testing.T) {
	got := collectCancellableMerge(t, merge(make(chan struct{}), cancellableBuffered(1, 4), cancellableBuffered(2, 3)))
	sort.Ints(got)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("merge() = %v, want %v", got, want)
	}
}

func TestMergeStopsWhileWaitingForInput(t *testing.T) {
	stop := make(chan struct{})
	out := merge(stop, make(chan int))
	close(stop)
	waitForCancellableMergeClose(t, out)
}

func TestMergeStopsWhileItsOutputSendIsBlocked(t *testing.T) {
	previous := onMergeBeforeSend
	beforeSend := make(chan struct{}, 1)
	onMergeBeforeSend = func() { beforeSend <- struct{}{} }
	t.Cleanup(func() { onMergeBeforeSend = previous })

	stop := make(chan struct{})
	in := make(chan int)
	out := merge(stop, in)
	sent := make(chan struct{})
	go func() {
		in <- 6
		close(sent)
	}()
	select {
	case <-sent:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not receive its input")
	}
	select {
	case <-beforeSend:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not begin its output send")
	}

	close(stop)
	waitForCancellableMergeClose(t, out)
}

func cancellableBuffered(values ...int) <-chan int {
	in := make(chan int, len(values))
	for _, value := range values {
		in <- value
	}
	close(in)
	return in
}

func collectCancellableMerge(t *testing.T, out <-chan int) []int {
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

func waitForCancellableMergeClose(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("merge() sent a value after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not close after cancellation")
	}
}
