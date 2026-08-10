// Concept: multi-stage pipelines compose independent channel lifecycles
// Task: build a pipeline that doubles each value and then adds one
// Expected behavior: [1,2,3] becomes [3,5,7], then output closes
// Hint: implement transform once, then compose two transform stages

package main

import "fmt"

func transform(in <-chan int, fn func(int) int) <-chan int {
	return nil // TODO: range in, apply fn, send out, and close out
}

func pipeline(in <-chan int) <-chan int {
	// Thought: downstream range can finish only after upstream closes its output.
	// Every stage owns the protocol of reading to completion and closing its output.
	return nil // TODO: compose double and add-one stages
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range pipeline(in) {
		fmt.Println(value)
	}
}
