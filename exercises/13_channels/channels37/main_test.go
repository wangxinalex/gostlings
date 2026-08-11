package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCollectOrStopForwardsThenCloses(t *testing.T) {
	work := make(chan int, 2)
	work <- 4
	work <- 9
	close(work)
	if got := collect37(t, collectOrStop(make(chan struct{}), work)); !reflect.DeepEqual(got, []int{4, 9}) {
		t.Fatalf("collectOrStop() = %v, want [4 9]", got)
	}
}
func TestCollectOrStopLetsStopAwareProducerFinishWhenConsumerAbandons(t *testing.T) {
	stop, work := make(chan struct{}), make(chan int)
	producerDone := make(chan struct{})
	go func() {
		defer close(producerDone)
		for _, v := range []int{1, 2} {
			select {
			case work <- v:
			case <-stop:
				return
			}
		}
	}()
	out := collectOrStop(stop, work)
	firstSent := make(chan struct{})
	go func() { <-out; close(firstSent) }()
	wait37(t, firstSent, "collectOrStop() did not forward the first value")
	close(stop)
	closed37(t, out)
	wait37(t, producerDone, "stop-aware producer did not finish after stop")
}
func collect37(t *testing.T, out <-chan int) []int {
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
			t.Fatal("collectOrStop() did not close")
		}
	}
}
func closed37(t *testing.T, out <-chan int) {
	t.Helper()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("collectOrStop() sent after stop")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("collectOrStop() did not close after stop")
	}
}
func wait37(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal(message)
	}
}
