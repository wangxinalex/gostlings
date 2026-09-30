// Concept: an internally built semaphore bounds how many jobs run at once.
// Task: run every job with at most limit jobs active at a time, and return the
// results in input order.
// Expected behavior: no more than limit calls to work run concurrently; the
// returned slice matches job order; an empty job list returns an empty slice; a
// limit below one still processes every job.
// Hint: build the limiter yourself as a buffered channel holding one permit per
//       slot. A worker takes a permit before calling work and puts it back
//       afterward, so the buffer itself enforces the bound. Carry each job's
//       index with its result, because completion order is not input order.
// Stuck?: sending into the prefilled permit channel acquires a permit; receiving
//       from it releases one.

package main

import "fmt"

func parallel(limit int, jobs []int, work func(int) int) []int {
	// TODO: bound the active calls to work with a permit channel, then return the
	//       results in the order of jobs.
	return nil
}

func main() { fmt.Println(parallel(1, nil, func(value int) int { return value })) }
