// Concept: select with default — a non-blocking channel operation
// Task: try to receive once and continue immediately when no value is ready
// Expected behavior: an empty input returns "no value"
// Hint: default runs when no channel case is ready; it must not be used in a hot loop without work or backoff

package main

import "fmt"

func tryReceive(ch <-chan int) string {
	// 思路：default 让 select 立即返回；它表达“现在试一次”，
	// 而不是“保证最终一定收到”。
	return "" // TODO: add receive and default cases
}

func main() {
	fmt.Println(tryReceive(make(chan int)))
}
