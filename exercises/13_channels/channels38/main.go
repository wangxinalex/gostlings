// Problem: several observers need to learn that one shared operation finished.
// Without this pattern: sending one completion value wakes only one observer or
// requires knowing the observer count in advance.
// Channels: done carries no data and is closed by the operation owner; each
// returned output carries one observer-specific message and is closed by watch.
// Timeline: close(done) -> every watcher receives -> each watcher sends "done" -> closes its output
// Hint: wait for `<-done` in a goroutine, send one string on a capacity-one output,
// then close that output. A closed done channel can be received by every watcher.
package main

import "fmt"

func watch(done <-chan struct{}) <-chan string {
	return nil // TODO: wait for done, publish one message, and close out
}
func main() { fmt.Println(watch(make(chan struct{}))) }
