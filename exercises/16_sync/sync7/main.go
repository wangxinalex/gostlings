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
	// TODO: Snapshot under the lock, transform outside it, and commit under the lock.
	return 0
}

func (s *protectedState) valueSnapshot() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

func main() {}
