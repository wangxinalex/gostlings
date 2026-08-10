// Concept: semaphore — a buffered channel can limit active work
// Task: run work concurrently but never exceed limit active jobs
// Expected behavior: results stay in input order and active work is bounded by limit
// Hint: acquire a token with tokens <- struct{}{} and release it with <-tokens

package main

import "fmt"

func parallel(limit int, jobs []int, work func(int) int) []int {
	// 思路：channel 容量表达“同时允许多少个工作”；它不是任务结果队列，
	// 每个 goroutine 必须在 work 前取得 token，并在结束后归还。
	return nil // TODO: add a buffered token channel and wait for all results
}

func main() {
	results := parallel(2, []int{1, 2, 3, 4}, func(value int) int { return value * value })
	fmt.Println(results)
}
