// Concept: sync.Cond waits for a predicate while holding a mutex.
// Task: implement a queue whose Pop waits until Push makes a value available.
// Hint: wait in a for loop, then remove one value while the lock is held.
package main

import "sync"

type intQueue struct {
	mu     sync.Mutex
	cond   *sync.Cond
	values []int
}

func newIntQueue() *intQueue {
	q := &intQueue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *intQueue) Push(value int) {
	// TODO: Append under the lock and signal a waiting Pop.
}

func (q *intQueue) Pop() int {
	// TODO: Wait in a predicate loop until a value is available, then remove it.
	return 0
}

func main() {}
