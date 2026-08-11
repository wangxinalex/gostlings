// Concept: several forwarders need one coordinator to close their shared output.
// Task: collect all source values into one output stream.
// Expected behavior: all sources drain; output closes only after every forwarder exits.
// Hint: give each forwarder an exit acknowledgement. A coordinator receives every acknowledgement,
//       then closes out; no forwarder owns the shared output close.
package main

import "fmt"

var onCollectorExit = func() {}

func collect(sources []<-chan int) <-chan int {
	return nil // TODO: forward every source and coordinate close(out)
}
func main() { fmt.Println(collect(nil)) }
