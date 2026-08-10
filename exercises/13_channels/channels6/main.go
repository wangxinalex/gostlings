// Concept: select multiplexes channel operations
// Task: receive from whichever input is ready without blocking on a silent input
// Expected behavior: a ready input is returned; if multiple inputs are ready, either may be selected
// Hint: put one receive from each input in a select. select does not give earlier
//       cases priority when multiple cases are ready.

package main

import "fmt"

func receiveFast(fast, slow <-chan string) string {
	// 思路：select 会同时等待多个 channel 操作；不要先接收 slow，
	// 也不要假设 fast 一定先到，因为 slow 可能先 ready，或者两个都 ready。
	// 当多个 case 同时 ready 时，select 的选择是不确定的。
	return <-fast // TODO: wait for whichever input is ready first
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string)
	fast <- "fast lane"
	fmt.Println(receiveFast(fast, slow))
}
