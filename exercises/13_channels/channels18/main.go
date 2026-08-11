// Concept: worker start and join with raw channels
// Task: start count workers, broadcast stop to them, and close done after every worker exits
// Expected behavior: one close(stop) reaches every worker; zero workers complete immediately
// Hint: each worker waits on <-stop then sends one signal to a buffered exited channel.
//       A coordinator receives count exit signals and is the only goroutine that closes done.

package main

import "fmt"

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	return nil // TODO: start workers, collect their exit signals, and close done once
}

func main() {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)
	<-done
	fmt.Println("workers stopped")
}
