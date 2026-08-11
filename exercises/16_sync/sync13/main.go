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
	// TODO: Wait while at the bound; reject acquisition after Close.
	return false
}

func (s *boundedState) Release() {
	// TODO: Decrement active safely and wake a waiting producer.
}

func (s *boundedState) Close() {
	// TODO: Mark closed and wake every waiter.
}

func (s *boundedState) Active() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active
}

func main() {}
