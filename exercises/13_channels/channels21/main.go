// Concept: fan-out workers share one jobs channel, then fan in to one result stream.
// Task: start workers workers that square every job, and close out only after all workers exit.
// Expected behavior: every job produces one square; one coordinator owns close(out).
// Hint: give each worker one buffered exited <- struct{}{} acknowledgement after its jobs range ends.
//       A separate goroutine receives workers acknowledgements, then closes out exactly once.

package main

import "fmt"

var onSquareWorkerStart = func() {}
var onSquareWorkerExit = func() {}

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	return nil // TODO: fan out jobs, fan in squares, and let one coordinator close out
}

func main() {
	jobs := make(chan int, 2)
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(2, jobs) {
		fmt.Println(result)
	}
}
