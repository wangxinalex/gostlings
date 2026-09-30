// Concept: collecting one error per job and reporting it after a full join
// Task: run every job concurrently, wait for all of them, and return the first non-nil error
// Expected behavior: every job runs and returns nil when they all succeed; otherwise the first
//                  non-nil error in job order is returned, and only after every job has
//                  finished, so later errors are dropped. An empty job list returns nil.
// Hint: launch each job in its own goroutine with its own slot in an errors slice, join the whole
//       group with sync.WaitGroup, and scan the slots in job order only after the join returns.
//       Later errors are dropped, and an empty job list needs no special wait.
// Version note: pass the job and its error slot into each worker; explicit parameters avoid
// depending on Go 1.22's changed loop-variable capture semantics.

package main

func runAll(jobs []func() error) error {
	// TODO: Run every job in its own goroutine, join all of them, then return the
	// first non-nil error in job order, or nil when every job succeeded.
	return nil
}

func main() {}
