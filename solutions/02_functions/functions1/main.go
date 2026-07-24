// Concept: defining and calling functions
// Task: fix this program so it compiles and runs
// Expected output: Hello from a function!
// Hint: a function is defined with the func keyword; printing text requires importing the fmt package (Go Tour: Basics 4)

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from a function!")
}

func main() {
	sayHello()
}
