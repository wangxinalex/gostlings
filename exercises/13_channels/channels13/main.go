// Concept: done signaling — a closed channel broadcasts completion without data
// Task: return a done channel that closes when the asynchronous work finishes
// Expected behavior: receiving from done reports a closed channel and no value
// Hint: make done, start a goroutine, and write defer close(done) at the top of it.
//       Do not send a value: closing is the completion signal.

package main

import "fmt"

func complete() <-chan struct{} {
	return nil // TODO: close a done channel from the completing goroutine
}

func main() {
	<-complete()
	fmt.Println("complete")
}
