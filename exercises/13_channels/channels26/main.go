// Concept: a basic worker pool has a jobs producer, workers, and one results closer.
// Task: square every item in jobs with workers workers and return all results in completion order.
// Expected behavior: every job is processed once; an empty jobs slice returns an empty result slice.
// Hint: start a producer goroutine that sends jobs then closes jobsCh. Each worker ranges jobsCh and sends to results.
//       Workers acknowledge exit to a coordinator; the coordinator closes results after every acknowledgement.

package main

import "fmt"

var processJob = func(value int) int { return value * value }

func run(workers int, jobs []int) []int {
	return nil // TODO: produce jobs, fan out workers, and collect results after coordinator close
}

func main() {
	fmt.Println(run(2, []int{1, 2, 3, 4}))
}
