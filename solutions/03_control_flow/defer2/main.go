// Concept: deferred calls run after the return value is set — and can change named results
// Task: add ONE deferred statement so countdown returns 3 even though it returns 0 first
// Expected output: 3
// Hint: with a named result (n int), a deferred closure can assign to n after the
//       return statement has evaluated; defer func() { n = 3 }() (Go Tour: Flowcontrol 12-13)

package main

import "fmt"

func countdown() (n int) {
	defer func() { n = 3 }()
	return 0
}

func main() {
	fmt.Println(countdown())
}
