package main

import (
	"errors"
	"sort"
	"testing"
	"time"
)

func TestServeRepliesReportsResultsAndClosesDoneAfterResults(t *testing.T) {
	jobs := make(chan request, 3)
	first, second, failed := make(chan response, 1), make(chan response, 1), make(chan response, 1)
	bad := errors.New("bad request")
	jobs <- request{value: 2, reply: first}
	jobs <- request{value: 4, reply: second}
	jobs <- request{value: 6, reply: failed, err: bad}
	close(jobs)

	results, done := serve(make(chan struct{}), jobs)
	got := collect47(t, results)
	sort.Slice(got, func(i, j int) bool { return got[i].value < got[j].value })
	if len(got) != 3 || got[0].value != 4 || got[0].err != nil || got[1].value != 8 || got[1].err != nil || got[2].value != 12 || !errors.Is(got[2].err, bad) {
		t.Fatalf("serve() results = %#v, want doubled responses including the request error", got)
	}
	for _, want := range []struct {
		reply <-chan response
		value int
		err   error
	}{
		{reply: first, value: 4},
		{reply: second, value: 8},
		{reply: failed, value: 12, err: bad},
	} {
		select {
		case got := <-want.reply:
			if got.value != want.value || !errors.Is(got.err, want.err) {
				t.Fatalf("serve() reply = %#v, want value %d and error %v", got, want.value, want.err)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatal("serve() did not send a per-request reply")
		}
	}
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not close done after results")
	}
}

func TestServeClosesResultsAndDoneForEmptyJobs(t *testing.T) {
	jobs := make(chan request)
	close(jobs)
	results, done := serve(make(chan struct{}), jobs)
	collect47(t, results)
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not finish empty jobs")
	}
}

func TestServeStopsBlockedResultPublicationAndJoinsBeforeDone(t *testing.T) {
	previous := onServeBeforeResult
	previousWorkers := serveWorkerCount
	beforeResult := make(chan struct{}, 1)
	release := make(chan struct{})
	onServeBeforeResult = func() {
		beforeResult <- struct{}{}
		<-release
	}
	serveWorkerCount = 1
	t.Cleanup(func() {
		onServeBeforeResult = previous
		serveWorkerCount = previousWorkers
	})

	stop, jobs := make(chan struct{}), make(chan request, 2)
	pendingReply := make(chan response, 1)
	jobs <- request{value: 3}
	jobs <- request{value: 5, reply: pendingReply}
	results, done := serve(stop, jobs)
	select {
	case <-beforeResult:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not begin publishing its result")
	}
	close(stop)
	close(release)
	got := collect47(t, results)
	for _, result := range got {
		if result.value == 10 {
			t.Fatalf("serve() published a result for a pending request after stop: %#v", result)
		}
	}
	select {
	case reply := <-pendingReply:
		t.Fatalf("serve() replied to a pending request after stop: %#v", reply)
	default:
	}
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("serve() did not join workers before done")
	}
}

func collect47(t *testing.T, results <-chan response) []response {
	t.Helper()
	var got []response
	for {
		select {
		case result, ok := <-results:
			if !ok {
				return got
			}
			got = append(got, result)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("serve() results did not close")
		}
	}
}
