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
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		return false
	}
	q.values = append(q.values, value)
	q.cond.Signal()
	return true
}

func (q *closableQueue) Pop() (int, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.values) == 0 && !q.closed {
		q.cond.Wait()
	}
	if len(q.values) == 0 {
		return 0, false
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}

func (q *closableQueue) Close() {
	q.mu.Lock()
	q.closed = true
	q.cond.Broadcast()
	q.mu.Unlock()
}

func main() {}
