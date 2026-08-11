// Concept: a service combines sync.Once initialization, protected state, and shutdown join.
// Task: allow active jobs until Shutdown, then wait for every started job to Finish.
// Hint: use Once for initialization, a mutex/Cond for active state, and Broadcast on shutdown.
package main

import "sync"

type serviceState struct {
	once        sync.Once
	mu          sync.Mutex
	cond        *sync.Cond
	initialized bool
	active      int
	closed      bool
}

var initializeService = func() {}

func newServiceState() *serviceState {
	state := &serviceState{}
	state.cond = sync.NewCond(&state.mu)
	return state
}

func (s *serviceState) Start() bool {
	// TODO: Initialize once, then reject new work after shutdown.
	return false
}

func (s *serviceState) Finish() {
	// TODO: Decrement active work and wake Wait callers.
}

func (s *serviceState) Shutdown() {
	// TODO: Prevent new work and wake all state waiters.
}

func (s *serviceState) Wait() {
	// TODO: Wait until all started work has called Finish.
}

func (s *serviceState) Initialized() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.initialized
}

func main() {}
