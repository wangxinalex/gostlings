// Concept: a request can carry a private reply channel.
// Task: double every request value and reply on that request's reply channel.
// Expected behavior: each caller receives its own doubled value; done closes after requests closes.
// Hint: range over requests in one server goroutine, send value*2 on request.reply, and defer close(done).
package main

import "fmt"

type request struct {
	value int
	reply chan int
}

func serve(requests <-chan request) <-chan struct{} {
	return nil // TODO: reply to every request and close done after input closure
}

func main() { fmt.Println(serve(nil)) }
