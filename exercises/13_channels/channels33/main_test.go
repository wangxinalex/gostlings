package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCollectForwardsSourcesAndCloses(t *testing.T) {
	first, second := make(chan int, 2), make(chan int, 1)
	first <- 1
	first <- 3
	second <- 2
	close(first)
	close(second)
	got := collect33(t, collect([]<-chan int{first, second}))
	sort.Ints(got)
	if want := []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("collect() = %v, want %v", got, want)
	}
}
func TestCollectClosesOnlyAfterEveryForwarderExits(t *testing.T) {
	previous := onCollectorExit
	exited, release := make(chan struct{}, 2), make(chan struct{})
	onCollectorExit = func() { exited <- struct{}{}; <-release }
	t.Cleanup(func() { onCollectorExit = previous })
	first, second := make(chan int), make(chan int)
	close(first)
	close(second)
	out := collect([]<-chan int{first, second})
	for range []int{0, 1} {
		wait33(t, exited, "collect() did not wait for a forwarder")
	}
	select {
	case _, ok := <-out:
		if !ok {
			t.Fatal("collect() closed before all forwarders exited")
		}
		t.Fatal("collect() produced an unexpected value")
	default:
	}
	close(release)
	collect33(t, out)
}
func collect33(t *testing.T, out <-chan int) []int {
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
			t.Fatal("collect() did not close")
		}
	}
}
func wait33(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}
