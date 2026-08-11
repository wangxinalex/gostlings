package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestRunBatchesJoinsBeforeStartingNextBatch(t *testing.T) {
	firstStarted := make(chan struct{})
	allowFirst := make(chan struct{})
	var mu sync.Mutex
	var events []string
	record := func(event string) {
		mu.Lock()
		events = append(events, event)
		mu.Unlock()
	}
	originalJob := runBatchJob
	originalBatchStart := onBatchStart
	runBatchJob = func(batch, job int) string {
		if batch == 0 {
			record("batch 0 job start")
			close(firstStarted)
			<-allowFirst
			record("batch 0 job finish")
		} else {
			record("batch 1 job finish")
		}
		return fmt.Sprintf("job %d done", job)
	}
	onBatchStart = func(batch int) {
		record(fmt.Sprintf("batch %d start", batch))
	}
	defer func() {
		runBatchJob = originalJob
		onBatchStart = originalBatchStart
	}()

	completed := make(chan [][]string, 1)
	go func() { completed <- runBatches([][]int{{1}, {2}}) }()

	select {
	case <-firstStarted:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("first batch did not start")
	}

	close(allowFirst)
	var got [][]string
	select {
	case got = <-completed:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runBatches did not complete")
	}

	mu.Lock()
	gotEvents := append([]string(nil), events...)
	mu.Unlock()
	finishFirst := -1
	startSecond := -1
	for index, event := range gotEvents {
		switch event {
		case "batch 0 job finish":
			finishFirst = index
		case "batch 1 start":
			startSecond = index
		}
	}
	if finishFirst == -1 || startSecond == -1 {
		t.Fatalf("event trace = %v, want first-batch finish and second-batch start", gotEvents)
	}
	if finishFirst > startSecond {
		t.Fatalf("event trace = %v, want batch 0 job finish before batch 1 start", gotEvents)
	}
	want := [][]string{{"job 1 done"}, {"job 2 done"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runBatches() = %v, want %v", got, want)
	}
}
