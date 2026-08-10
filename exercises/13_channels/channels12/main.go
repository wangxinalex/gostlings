// Concept: broadcast cancellation stops a group of goroutines
// Task: start count workers and close the returned done channel after all stop
// Expected behavior: closing stop lets every worker exit and done closes
// Hint: each worker waits on stop; a WaitGroup coordinator closes done once

package main

import (
	"fmt"
	"time"
)

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	// 思路：stop 是广播信号，所有 worker 都接收同一个关闭事件；
	// done 则表示整个 worker group 已经收尾。
	return done // TODO: start workers and close done after all of them exit
}

func main() {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)
	select {
	case <-done:
		fmt.Println("workers stopped")
	case <-time.After(time.Second):
		fmt.Println("workers still running")
	}
}
