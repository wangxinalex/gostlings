package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestFanOutSquaresAndCloses(t *testing.T) {
	jobs := make(chan int, 3)
	for _, v := range []int{1, 2, 3} {
		jobs <- v
	}
	close(jobs)
	got := collect35(t, fanOut(make(chan struct{}), jobs, 2))
	sort.Ints(got)
	if want := []int{1, 4, 9}; !reflect.DeepEqual(got, want) {
		t.Fatalf("fanOut() = %v, want %v", got, want)
	}
}
func TestFanOutStopsBlockedResultSend(t *testing.T) {
	previous := onFanOutBeforeSend
	entered := make(chan struct{}, 1)
	onFanOutBeforeSend = func() { entered <- struct{}{} }
	t.Cleanup(func() { onFanOutBeforeSend = previous })
	stop, jobs := make(chan struct{}), make(chan int)
	out := fanOut(stop, jobs, 1)
	sent := make(chan struct{})
	go func() { jobs <- 4; close(sent) }()
	wait35(t, sent, "fanOut() did not receive its job")
	wait35(t, entered, "fanOut() did not begin its result send")
	close(stop)
	closed35(t, out)
}
func TestFanOutStopsBlockedJobReceive(t *testing.T) {
	stop, jobs := make(chan struct{}), make(chan int)
	out := fanOut(stop, jobs, 2)
	close(stop)
	closed35(t, out)
}
func collect35(t *testing.T, out <-chan int) []int {
	t.Helper()
	var got []int
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("fanOut() did not close")
		}
	}
}
func closed35(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("fanOut() sent after cancellation")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("fanOut() did not close after cancellation")
	}
}
func wait35(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}
