package main

import (
	"sync"
	"testing"
	"time"
)

func TestServiceInitializesOnceAndJoinsShutdown(t *testing.T) {
	state := newServiceState()
	previous := initializeService
	var mu sync.Mutex
	initializations := 0
	initializeService = func() { mu.Lock(); initializations++; mu.Unlock() }
	t.Cleanup(func() { initializeService = previous })
	const callers = 12
	started := make(chan bool, callers)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); started <- state.Start() }()
	}
	wg.Wait()
	for i := 0; i < callers; i++ {
		if !<-started {
			t.Fatal("Start() rejected work before shutdown")
		}
	}
	if !state.Initialized() || initializations != 1 {
		t.Fatalf("initialized=%v calls=%d, want true and one call", state.Initialized(), initializations)
	}
	state.Shutdown()
	waitDone := make(chan struct{})
	go func() { state.Wait(); close(waitDone) }()
	select {
	case <-waitDone:
		t.Fatal("Wait() returned before active work finished")
	case <-time.After(20 * time.Millisecond):
	}
	for i := 0; i < callers; i++ {
		state.Finish()
	}
	select {
	case <-waitDone:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Wait() did not return after all Finish calls")
	}
}

func TestServiceRejectsNewWorkAfterShutdown(t *testing.T) {
	state := newServiceState()
	state.Shutdown()
	if state.Start() {
		t.Fatal("Start() succeeded after Shutdown")
	}
	state.Wait()
}
