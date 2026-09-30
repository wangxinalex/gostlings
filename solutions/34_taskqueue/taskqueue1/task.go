// Concept: a domain type owns its validation and its display form.
// Task: implement Validate and Status.String.
// Expected behavior: focused tests pass (run `go test ./exercises/34_taskqueue/taskqueue1`)
// Hint: a task can be wrong in more than one way, so collect every problem instead
//       of returning at the first one; a caller then asks whether one specific
//       failure happened. Render an unrecognised status distinctly rather than
//       falling back to a valid one.
// Stuck?: errors.Join combines the independent failures; fmt's string verb uses
//       the String method.

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Status is where a task sits in its lifecycle.
type Status string

// The two statuses a queue can hold.
const (
	Open Status = "open"
	Done Status = "done"
)

// The causes a task can be rejected for.
var (
	ErrTitleRequired    = errors.New("title is required")
	ErrPriorityNegative = errors.New("priority must not be negative")
)

// Task is one unit of work in the queue.
type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority int    `json:"priority"`
	Status   Status `json:"status"`
}

// Validate reports every problem with the task.
func (t Task) Validate() error {
	var problems []error
	if strings.TrimSpace(t.Title) == "" {
		problems = append(problems, ErrTitleRequired)
	}
	if t.Priority < 0 {
		problems = append(problems, ErrPriorityNegative)
	}
	return errors.Join(problems...)
}

// String renders the status for output.
func (s Status) String() string {
	switch s {
	case Open:
		return string(Open)
	case Done:
		return string(Done)
	default:
		return "unknown status " + string(s)
	}
}

// Label renders a task for humans.
func (t Task) Label() string {
	return fmt.Sprintf("%d [%s] %s", t.ID, t.Status, t.Title)
}
