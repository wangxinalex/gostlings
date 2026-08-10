// Concept: fan-out — several workers consume one jobs channel
// Task: square every job and close results after every worker exits
// Expected behavior: one squared result per job, followed by a closed output
// Hint: workers range over jobs; a coordinator waits for them before closing results

package main

import "fmt"

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	// Thought: closing jobs tells every worker that no new work remains. results
	// must be closed after all workers exit, or a worker may send to a closed channel.
	return nil // TODO: start workers and coordinate result closure
}

func main() {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(2, jobs) {
		fmt.Println(result)
	}
}
