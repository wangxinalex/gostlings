package main

import (
	"testing"
	"time"
)

func TestWatchBroadcastsDoneToEveryObserver(t *testing.T) {
	done := make(chan struct{})
	first, second := watch(done), watch(done)
	for _, out := range []<-chan string{first, second} {
		select {
		case got, ok := <-out:
			t.Fatalf("watch() received before done closed: (%q,%v)", got, ok)
		default:
		}
	}
	close(done)
	for _, out := range []<-chan string{first, second} {
		select {
		case got, ok := <-out:
			if !ok || got != "done" {
				t.Fatalf("watch() = (%q,%v), want (done,true)", got, ok)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatal("watch() did not report done")
		}
		select {
		case _, ok := <-out:
			if ok {
				t.Fatal("watch() did not close after its message")
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatal("watch() did not close")
		}
	}
}
