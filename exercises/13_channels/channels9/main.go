// Concept: a done channel reports completion without carrying data
// Task: close the done channel when the background operation finishes
// Expected behavior: complete() returns a channel that closes promptly
// Hint: the goroutine should defer close(done); a close is a broadcast notification

package main

import "fmt"

func complete() <-chan struct{} {
	done := make(chan struct{})
	// 思路：done 只表示“结束了”，不承载结果；关闭它可以同时唤醒所有等待者。
	return done // TODO: start work and close done when it finishes
}

func main() {
	<-complete()
	fmt.Println("completed")
}
