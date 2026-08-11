// Concept: closing one done channel broadcasts completion to every observer.
// Task: return a stream that reports "done" once the shared done channel closes.
// Expected behavior: every independent watch call receives "done" then its stream closes.
// Hint: wait for <-done in a goroutine, send one string on a buffered output, then close that output.
package main

import "fmt"

func watch(done <-chan struct{}) <-chan string {
	return nil // TODO: wait for done, publish one message, and close out
}
func main() { fmt.Println(watch(make(chan struct{}))) }
