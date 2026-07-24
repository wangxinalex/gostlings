// Concept: strconv.Atoi and strconv.Itoa — string↔int conversion
// Task: convert the string "42" to an int, add 1 to it, and convert back to a string to print
// Expected output: 43
// Hint: strconv.Atoi returns (int, error); strconv.Itoa returns string (Go doc: strconv)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "42"
	n, _ := strconv.Atoi(s)
	n++
	fmt.Println(strconv.Itoa(n))
}
