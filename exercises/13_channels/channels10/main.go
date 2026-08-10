// Concept: cancellation-aware sends prevent a producer from being stranded
// Task: stop producing when the caller closes stop, even if nobody receives output
// Expected behavior: cancellation closes the output instead of leaving a blocked goroutine
// Hint: select between out <- value and <-stop for every potentially blocking send

package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := 1; value <= 3; value++ {
			// TODO: Make this send respond to stop.
			out <- value
		}
	}()
	return out
}

func main() {
	for value := range produce(make(chan struct{})) {
		fmt.Println(value)
	}
}
