// Problem: a pool API must define what zero workers means instead of leaving a
// producer blocked forever.
// Without this pattern: starting a jobs producer with no consumer deadlocks, and
// the caller cannot distinguish “no capacity” from slow work.
// Channels: jobs and results are internal; workers own job receives, and one
// coordinator owns result closure. The zero-worker policy is an API decision.
// Timeline: workers < 1 -> return empty result; otherwise produce -> workers -> close results
// Hint: handle workers < 1 before starting the jobs producer. For positive
// workers, reuse the pool protocol: close jobs after production, acknowledge
// every worker exit, close results once, and collect results to completion.
package main

import "fmt"

func run(workers int, jobs []int) []int {
	return nil // TODO: define zero-worker behavior and coordinate results
}
func main() { fmt.Println(run(2, []int{1, 2, 3})) }
