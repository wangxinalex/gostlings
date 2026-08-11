package main

import (
	"context"
	"testing"
	"time"
)

func TestMergeForwardsValuesAndErrorsThenCloses(t *testing.T) {
	first := make(chan mergeResult, 2)
	second := make(chan mergeResult, 1)
	first <- mergeResult{value: 1}
	first <- mergeResult{err: errMergeTest}
	second <- mergeResult{value: 2}
	close(first)
	close(second)
	out := merge(context.Background(), first, second)
	count := 0
	for {
		select {
		case _, ok := <-out:
			if !ok {
				goto closed
			}
			count++
		case <-time.After(500 * time.Millisecond):
			t.Fatal("merge output did not close")
		}
	}
closed:
	if count != 3 {
		t.Fatalf("merged result count = %d, want 3", count)
	}
}

func TestMergeClosesWhenContextIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	source := make(chan mergeResult)
	out := merge(ctx, source)
	cancel()
	select {
	case _, ok := <-out:
		if ok {
			for range out {
			}
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("merge() did not close after cancellation")
	}
}

var errMergeTest = context.Canceled
