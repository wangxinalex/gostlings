package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRunBatchesJoinsBeforeStartingNextBatch(t *testing.T) {
	firstStarted := make(chan struct{})
	allowFirst := make(chan struct{})
	secondStarted := make(chan struct{}, 1)
	original := runBatchJob
	runBatchJob = func(batch, job int) string {
		if batch == 0 {
			close(firstStarted)
			<-allowFirst
		} else {
			secondStarted <- struct{}{}
		}
		return fmt.Sprintf("job %d done", job)
	}
	defer func() { runBatchJob = original }()

	completed := make(chan [][]string, 1)
	go func() { completed <- runBatches([][]int{{1}, {2}}) }()

	select {
	case <-firstStarted:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("first batch did not start")
	}
	select {
	case <-secondStarted:
		t.Fatal("second batch started before the first batch was joined")
	default:
	}

	close(allowFirst)
	select {
	case got := <-completed:
		want := [][]string{{"job 1 done"}, {"job 2 done"}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("runBatches() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBatches did not complete")
	}
}
