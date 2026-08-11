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
	q.mu.Lock()
	q.values = append(q.values, value)
	q.mu.Unlock()
	q.cond.Signal()
}

func (q *intQueue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.values) == 0 {
		q.cond.Wait()
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

func main() {}
