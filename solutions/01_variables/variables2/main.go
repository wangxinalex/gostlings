// Concept: assignment versus redeclaration
// Task: fix this program so it compiles and runs
// Expected output: second
// Hint: := declares a new variable and may only be used once per variable; use = to assign a new value (Go Tour: Basics 10)

package main

import "fmt"

func main() {
	message := "first"
	message = "second"
	fmt.Println(message)
}
