// Concept: worker pool — combine fan-out, result collection, and close ordering
// Task: process every job with a fixed number of workers and return all squares
// Expected behavior: one result per job; empty jobs return promptly
// Hint: close jobs after sending, wait for workers, then close results while the caller ranges results

package main

import "fmt"

func run(workers int, jobs []int) []int {
	// 思路：生产者关闭 jobs，worker 因 range 结束；只有确认所有 worker 退出后，
	// 协调者才能关闭 results，调用者的 range 才能返回完整结果。
	return nil // TODO: build the jobs/results lifecycle
}

func main() {
	fmt.Println(run(2, []int{1, 2, 3, 4}))
}
