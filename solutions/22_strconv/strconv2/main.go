// Concept: parsing numbers from strings — strconv.ParseFloat
// Task: parse "3.14" as a float64, double it, then format and print to 2 decimal places
// Expected output: 6.28
// Hint: strconv.ParseFloat(s, 64) returns (float64, error); fmt.Sprintf("%.2f", f) formats to 2 decimal places (Go doc: strconv)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3.14"
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("%.2f\n", f*2)
}
