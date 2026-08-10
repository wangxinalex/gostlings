// Concept: select multiplexes channel operations
// Task: receive from the input that becomes ready first
// Expected behavior: the ready fast input wins while the silent input does not block the call
// Hint: put one receive from each input in a select

package main

import "fmt"

func receiveFast(fast, slow <-chan string) string {
	// 思路：select 会同时等待多个 channel 操作；不要先接收 slow，
	// 因为 slow 可能永远没有发送者。
	return <-fast // TODO: wait for whichever input is ready first
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string)
	fast <- "fast lane"
	fmt.Println(receiveFast(fast, slow))
}
