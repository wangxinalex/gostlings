// Concept: defer and LIFO order
// Task: add three defer statements so the program prints 3, 2, 1
// Expected output: 3, 2, 1 (each on its own line)
// Hint: deferred calls run in last-in-first-out order when the function returns (Go Tour: Flowcontrol 12-13)

package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
