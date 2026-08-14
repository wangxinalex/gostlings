// Concept: launching and joining one goroutine
// Task: start the greeting task in a goroutine and wait for it before main returns
// Expected output: hello
// Hint: call wg.Add(1) before go, defer wg.Done() inside the goroutine, and call wg.Wait().
// Version note: this repository targets Go 1.26, so sync.WaitGroup.Go is available.
// On older Go versions, use the Add/go/defer Done form shown in the hint.

package main

import "fmt"

var greeting = func() {
	fmt.Println("hello")
}

func main() {
	// TODO: Start a greeting goroutine that prints "hello" and wait for it to finish.
}
