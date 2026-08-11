// Concept: launching and joining one goroutine
// Task: start the greeting task in a goroutine and wait for it before main returns

package main

import (
	"fmt"
	"sync"
)

var greeting = func() {
	fmt.Println("hello")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		greeting()
	}()
	wg.Wait()
}
