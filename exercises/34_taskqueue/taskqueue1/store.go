// Concept: the store owns ordering, id assignment, and the copy boundary.
// Task: implement Add, List, and Complete on the in-memory queue.
// Expected behavior: focused tests pass (run `go test ./exercises/34_taskqueue/taskqueue1`)
// Hint: the store decides ids and keeps insertion order, so a listing reads back in
//       the order tasks were added. Handing out the store's own slice lets a caller
//       rewrite stored state, and a missing id is a distinct failure that the
//       command layer reports rather than a panic.
// Stuck?: validate before appending; a sentinel error plus fmt's wrapping verb
//       keeps the cause available to errors.Is.

package main

import (
	"errors"
	"fmt"
)

// ErrNotFound reports a lookup by an id the queue does not hold.
var ErrNotFound = errors.New("task not found")

// Store is the queue as the command layer sees it.
type Store interface {
	Add(title string, priority int) (Task, error)
	List() []Task
	Complete(id int) error
}

// Memory keeps the queue in a slice and hands out ids.
type Memory struct {
	tasks []Task
	next  int
}

// NewMemory returns an empty queue whose first id is 1.
func NewMemory() *Memory { return &Memory{next: 1} }

// Add validates a new task and appends it to the queue.
func (m *Memory) Add(title string, priority int) (Task, error) {
	// TODO: build the candidate with the next id and open status, reject it when
	//       validation fails, and append it so listing keeps insertion order.
	return Task{}, nil
}

// List returns the queue in insertion order.
func (m *Memory) List() []Task {
	// TODO: return a copy of the queue, so a caller cannot change stored state.
	return nil
}

// Complete marks the task with this id as done.
func (m *Memory) Complete(id int) error {
	// TODO: mark the matching task done, or report the id as not found.
	return nil
}

// NotFound wraps the sentinel with the id that was searched for.
func NotFound(id int) error {
	return fmt.Errorf("task %d: %w", id, ErrNotFound)
}
