// Concept: function values and closures
// Task: complete makeGreeter so it returns a function that prefixes each name
//       with the greeting it received
// Expected output: Hello, Alice!
// Hello, Bob!
// Hint: functions can be values; return `func(name string) string { ... }` and
//       use the captured prefix inside it (function values and closures are not
//       covered in the Go Tour)

package main

import "fmt"

func makeGreeter(prefix string) func(string) string {
	// TODO: Return a function that adds prefix before each name.
}

func main() {
	greet := makeGreeter("Hello, ")
	fmt.Println(greet("Alice") + "!")
	fmt.Println(greet("Bob") + "!")
}
