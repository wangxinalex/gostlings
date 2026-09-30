// Concept: retrying work waits on a timer between attempts and stops at the first success.
// Task: call work up to attempts times, wait for the injected timer after each failure,
// and report how many attempts ran.
// Expected behavior: a first-attempt success creates no timer and reports one attempt; a
// later success stops immediately and reports the attempt that succeeded; running out of
// attempts returns the last error unchanged without waiting again; an attempt count below
// one is invalid, runs no work, and returns an error with zero attempts.
// Hint: loop over the attempts, return as soon as the work reports success, and otherwise
// create a timer through the retryTimer hook, wait for its channel, and stop it before the
// next attempt.
// Stuck?: time.NewTimer
package main

import "time"

var retryTimer = time.NewTimer

func retryWithTimer(attempts int, delay time.Duration, work func(attempt int) error) (int, error) {
	// TODO: retry after each failure with a timer wait, stop at the first success, and report the attempt count.
	return 0, nil
}

func main() {}
