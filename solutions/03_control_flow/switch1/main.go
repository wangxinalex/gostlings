// Concept: switch without a condition
// Task: add the default branch so a score below 60 prints "fail"
// Expected output: fail
// Hint: a switch without a condition matches the first true case; default runs when no case matches (Go Tour: Flowcontrol 9)

package main

import "fmt"

func main() {
	switch score := 50; {
	case score >= 90:
		fmt.Println("excellent")
	case score >= 60:
		fmt.Println("pass")
	default:
		fmt.Println("fail")
	}
}
