// Concept: first-error cancellation in a worker pool
// Task: report one error, cancel remaining work, and wait for every worker to exit
// Expected behavior: a failing job returns an error without leaking or double-closing stop
// Hint: use a buffered error channel and sync.Once (or one coordinator) to close stop exactly once

package main

import "fmt"

type job struct {
	value int
	fail  bool
}

func run(workers int, jobs []job) error {
	// Thought: an error path still needs the full lifecycle: report the error,
	// broadcast cancellation, stop production, wait for workers, and return the
	// final error to the caller.
	return nil // TODO: add error reporting and cancellation to the worker pool
}

func main() {
	err := run(2, []job{{value: 1}, {fail: true}, {value: 2}})
	if err != nil {
		fmt.Println(err)
	}
}
