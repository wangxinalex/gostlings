// Concept: launching goroutines with the go keyword
// Task: run say as a goroutine so the message actually prints before main exits
// Expected output: hello
// Hint: use go before the function call; you may need a short time.Sleep in main to let it finish (Go Tour: Concurrency 1)

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
}

func main() {
	go say("hello")
	time.Sleep(100 * time.Millisecond)
}
