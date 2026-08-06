// Concept: launching goroutines with the go keyword
// Task: run say as a goroutine so the message actually prints before main exits
// Expected output: hello
// Hint: use go before the function call, then sleep briefly in main so the
//       goroutine has time to finish. The sleep is only a placeholder — the
//       proper way to wait (sync.WaitGroup) is the next exercise (Go Tour: Concurrency 1)

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
}

func main() {
	// TODO: Launch say("hello") as a goroutine, then sleep briefly so it can finish.
	say("hello")
}
