// Concept: a worker pool needs a defined behavior when there are no workers.
// Task: square every job with the requested workers.
// Expected behavior: empty jobs returns an empty slice; zero workers returns an empty slice without consuming jobs.
// Hint: handle workers < 1 before starting the jobs producer. Otherwise use worker exit acknowledgements
//
//	and let one coordinator close results after every worker exits.
package main

import "fmt"

func run(workers int, jobs []int) []int {
	return nil // TODO: define zero-worker behavior and coordinate results
}
func main() { fmt.Println(run(2, []int{1, 2, 3})) }
