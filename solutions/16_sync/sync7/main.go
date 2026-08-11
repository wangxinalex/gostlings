// Concept: take a protected snapshot before doing slow work.
// Task: compute outside the mutex, then commit the result safely.
// Hint: lock only to copy/read and commit; the transform callback must run without the lock.
package main

import "sync"

type protectedState struct {
	mu    sync.Mutex
	value int
}

func (s *protectedState) updateState(transform func(int) int) int {
	s.mu.Lock()
	snapshot := s.value
	s.mu.Unlock()

	next := transform(snapshot)
	s.mu.Lock()
	s.value = next
	s.mu.Unlock()
	return next
}

func (s *protectedState) valueSnapshot() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

func main() {}
