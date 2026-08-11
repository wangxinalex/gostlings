// Concept: cancellable producer — a blocked send needs an exit path
// Task: produce increasing values until stop closes, then close the output
// Expected behavior: the first value is 1; closing stop also releases a blocked send
// Hint: start a goroutine with defer close(out). For every out <- value, use select with
//       a second case <-stop: return, so an abandoned receiver cannot leave it blocked.

package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	return nil // TODO: send values with a stop case and close out when stopping
}

func main() {
	stop := make(chan struct{})
	out := produce(stop)
	fmt.Println(<-out)
	close(stop)
}
