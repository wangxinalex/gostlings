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
	s.once.Do(func() {
		initializeService()
		s.mu.Lock()
		s.initialized = true
		s.mu.Unlock()
	})
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return false
	}
	s.active++
	return true
}

func (s *serviceState) Finish() {
	s.mu.Lock()
	if s.active > 0 {
		s.active--
	}
	s.cond.Broadcast()
	s.mu.Unlock()
}

func (s *serviceState) Shutdown() {
	s.mu.Lock()
	s.closed = true
	s.cond.Broadcast()
	s.mu.Unlock()
}

func (s *serviceState) Wait() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for s.active > 0 {
		s.cond.Wait()
	}
}

func (s *serviceState) Initialized() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.initialized
}

func main() {}
