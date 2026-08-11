// Concept: a condition variable coordinates a bounded shared resource.
// Task: Acquire waits at the limit, Release frees capacity, and Close wakes all waiters.
// Hint: protect active/closed with the mutex and wait in a predicate loop.
package main

import "sync"

type boundedState struct {
	mu     sync.Mutex
	cond   *sync.Cond
	limit  int
	active int
	closed bool
}

func newBoundedState(limit int) *boundedState {
	state := &boundedState{limit: limit}
	state.cond = sync.NewCond(&state.mu)
	return state
}

func (s *boundedState) Acquire() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for s.active >= s.limit && !s.closed {
		s.cond.Wait()
	}
	if s.closed {
		return false
	}
	s.active++
	return true
}

func (s *boundedState) Release() {
	s.mu.Lock()
	if s.active > 0 {
		s.active--
	}
	s.cond.Signal()
	s.mu.Unlock()
}

func (s *boundedState) Close() {
	s.mu.Lock()
	s.closed = true
	s.cond.Broadcast()
	s.mu.Unlock()
}

func (s *boundedState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active
}

func main() {}
