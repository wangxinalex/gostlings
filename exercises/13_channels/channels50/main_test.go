package main

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestRunServiceRepliesAndReturnsOrderedResponses(t *testing.T) {
	first, second, third := make(chan response, 1), make(chan response, 1), make(chan response, 1)
	requests := []request{{value: 3, reply: first}, {value: 1, reply: second}, {value: 2, reply: third}}
	got, err := runService(make(chan struct{}), 2, requests)
	if err != nil {
		t.Fatalf("runService() error = %v, want nil", err)
	}
	if want := []response{{value: 6}, {value: 2}, {value: 4}}; !reflect.DeepEqual(got, want) {
		t.Fatalf("runService() = %#v, want ordered %#v", got, want)
	}
	for index, reply := range []<-chan response{first, second, third} {
		select {
		case got := <-reply:
			if got != want50(index) {
				t.Fatalf("reply %d = %#v, want %#v", index, got, want50(index))
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("runService() did not send reply %d", index)
		}
	}
}

func TestRunServiceHandlesEmptyRequestsAndPreStartCancellation(t *testing.T) {
	got, err := runService(make(chan struct{}), 2, nil)
	if err != nil || len(got) != 0 {
		t.Fatalf("runService(nil) = (%v, %v), want ([], nil)", got, err)
	}
	stop := make(chan struct{})
	close(stop)
	got, err = runService(stop, 2, []request{{value: 1}})
	if !errors.Is(err, errStopped) || len(got) != 0 {
		t.Fatalf("runService(stopped) = (%v, %v), want ([], errStopped)", got, err)
	}
}

func TestRunServiceBoundsWorkersUnderBackpressure(t *testing.T) {
	previous := processServiceRequest
	started := make(chan int, 3)
	release := make(chan struct{})
	processServiceRequest = func(current request) response {
		started <- current.value
		<-release
		return response{value: current.value * 2}
	}
	t.Cleanup(func() { processServiceRequest = previous })

	returned := make(chan struct {
		responses []response
		err       error
	}, 1)
	go func() {
		responses, err := runService(make(chan struct{}), 2, []request{{value: 1}, {value: 2}, {value: 3}})
		returned <- struct {
			responses []response
			err       error
		}{responses, err}
	}()
	for worker := 0; worker < 2; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("runService() did not start its bounded workers")
		}
	}
	select {
	case value := <-started:
		t.Fatalf("runService() started third request %d before capacity released", value)
	case <-time.After(100 * time.Millisecond):
	}
	close(release)
	select {
	case got := <-returned:
		if got.err != nil || !reflect.DeepEqual(got.responses, []response{{value: 2}, {value: 4}, {value: 6}}) {
			t.Fatalf("runService() = (%#v, %v), want ordered successful responses", got.responses, got.err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not finish after workers were released")
	}
}

func TestRunServicePropagatesRequestErrorAndRepliesOnce(t *testing.T) {
	bad := errors.New("bad request")
	reply := make(chan response, 1)
	got, err := runService(make(chan struct{}), 1, []request{{value: 1, reply: reply, err: bad}, {value: 2}})
	if !errors.Is(err, bad) {
		t.Fatalf("runService() error = %v, want %v", err, bad)
	}
	if len(got) != 0 {
		t.Fatalf("runService() results = %#v, want no successful responses after first error", got)
	}
	select {
	case response := <-reply:
		if !errors.Is(response.err, bad) {
			t.Fatalf("failure reply = %#v, want propagated error %v", response, bad)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not reply to the failing request")
	}
	select {
	case <-reply:
		t.Fatal("runService() sent more than one response on one reply channel")
	default:
	}
}

func TestRunServiceStopsBlockedPerRequestReply(t *testing.T) {
	previous := processServiceRequest
	started := make(chan struct{}, 1)
	processServiceRequest = func(current request) response {
		started <- struct{}{}
		return response{value: current.value * 2}
	}
	t.Cleanup(func() { processServiceRequest = previous })

	stop := make(chan struct{})
	returned := make(chan error, 1)
	go func() {
		_, err := runService(stop, 1, []request{{value: 3, reply: make(chan response)}})
		returned <- err
	}()
	select {
	case <-started:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not start the request with a blocked reply")
	}
	close(stop)
	select {
	case err := <-returned:
		if !errors.Is(err, errStopped) {
			t.Fatalf("runService() error = %v, want errStopped", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runService() did not cancel its blocked reply send")
	}
}

func want50(index int) response {
	return []response{{value: 6}, {value: 2}, {value: 4}}[index]
}
