// Concept: type assertions on the empty interface
// Task: assert v as a string so this program compiles and runs
// Expected output: length: 5
// Hint: s := v.(string) extracts the string stored in v; the empty interface any can hold a value of any type (Go Tour: Methods 14-15)

package main

import "fmt"

func main() {
	var v any = "hello"
	// TODO: Assert v as a string and assign it to s.
	fmt.Println("length:", len(s))
}
