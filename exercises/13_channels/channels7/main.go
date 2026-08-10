// Concept: select timeout — a caller should not wait forever for a result
// Task: return the value if it arrives quickly, otherwise return a timeout message
// Expected behavior: a silent input returns "timed out" promptly
// Hint: add a case for <-time.After(50 * time.Millisecond)

package main

import (
	"fmt"
	"time"
)

func await(ch <-chan string) string {
	// 思路：超时是 select 的一个竞争分支；它只结束当前等待，
	// 不会自动停止仍在运行的生产者，后面还要学习取消。
	timeout := time.After(50 * time.Millisecond)
	_ = timeout
	return <-ch // TODO: add the timeout branch
}

func main() {
	fmt.Println(await(make(chan string)))
}
