// Concept: unbuffered channels — send and receive must happen in different goroutines
// Task: this program deadlocks because it sends and receives in the same goroutine; move the send into a goroutine
// Expected output: hi
// Hint: an unbuffered channel blocks send until a receive is ready — do the send in a goroutine (Go Tour: Concurrency 2)

package main

import "fmt"

func main() {
	ch := make(chan string)

	// 思路：无缓冲 channel 的发送必须等到接收者同时出现。
	// TODO: Move the send into a goroutine so this doesn't deadlock.
	ch <- "hi"

	fmt.Println(<-ch)
}
