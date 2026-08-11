// Concept: sync.Cond.Broadcast wakes every waiter during shutdown.
// Task: make a queue that returns closed when it is empty and closed.
// Hint: Close sets the predicate and Broadcasts; Pop checks both values and closed in a loop.
package main

import "sync"

type closableQueue struct {
	mu     sync.Mutex
	cond   *sync.Cond
	values []int
	closed bool
}

func newClosableQueue() *closableQueue {
	q := &closableQueue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *closableQueue) Push(value int) bool {
	// TODO: Reject pushes after close; signal a waiter after a successful push.
	return false
}

func (q *closableQueue) Pop() (int, bool) {
	// TODO: Wait while open and empty; return false after close and drain.
	return 0, false
}

func (q *closableQueue) Close() {
	// TODO: Mark closed and broadcast so every waiter can exit.
}

func main() {}
