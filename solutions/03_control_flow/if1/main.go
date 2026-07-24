// Concept: if/else conditions
// Task: fix the condition so the program reports the number correctly
// Expected output: odd
// Hint: an if statement runs its branch only when the condition is true (Go Tour: Flowcontrol 5)

package main

import "fmt"

func main() {
	n := 9
	if n%2 != 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}
