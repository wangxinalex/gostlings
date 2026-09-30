// Concept: parsing numbers from strings — strconv.ParseFloat
// Task: parse "3.14" as a float64, double it, then format and print to 2 decimal places
// Expected output: 6.28
// Hint: the strconv package parses text into a float and reports failure
//       separately; format the result with two decimals through fmt (Go doc: strconv)
// Stuck?: strconv.ParseFloat takes a bit size (64 here); fmt's precision verb sets decimals.

package main

import (
	"fmt"
)

func main() {
	s := "3.14"
	// TODO: Parse s as float64, double it, and print with 2 decimal places.
	fmt.Println(s)
}
