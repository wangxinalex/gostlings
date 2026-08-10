// Concept: first-error cancellation in a worker pool
// Task: report one error, cancel remaining work, and wait for every worker to exit
// Expected behavior: a failing job returns an error without leaking or double-closing stop
// Hint: use a buffered error channel and sync.Once (or one coordinator) to close stop exactly once

package main

import "fmt"

type job struct {
	value int
	fail  bool
}

func run(workers int, jobs []job) error {
	// 思路：错误路径也必须走完整生命周期：报告错误、广播取消、停止生产、
	// 等待 worker，再由调用者拿到最终 error。
	return nil // TODO: add error reporting and cancellation to the worker pool
}

func main() {
	err := run(2, []job{{value: 1}, {fail: true}, {value: 2}})
	if err != nil {
		fmt.Println(err)
	}
}
