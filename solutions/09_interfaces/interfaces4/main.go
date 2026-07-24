// Concept: safe type assertions with the comma-ok form
// Task: this program panics because v is not an int; use the comma-ok form to print "not an int" instead
// Expected output: not an int
// Hint: n, ok := v.(int) sets ok to false on a failed assertion instead of panicking (Go Tour: Methods 15)

package main

import "fmt"

func main() {
	var v any = "hello"
	if n, ok := v.(int); ok {
		fmt.Println(n)
	} else {
		fmt.Println("not an int")
	}
}
