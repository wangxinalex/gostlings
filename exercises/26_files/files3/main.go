// Concept: bufio.Scanner reads input line by line
// Task: open numbers.txt, scan it line by line, and print the sum of the numbers
// Expected output: sum: 6
// Hint: bufio scans a file one line at a time; convert each line with strconv,
//       check the scanner's error after the loop, and close the file as soon as
//       you open it. Run from this directory (`cd exercises/26_files/files3 &&
//       go run .`) or verify with `go test ./exercises/26_files/files3` (Go doc: bufio)
// Stuck?: bufio.NewScanner's Scan and Text methods; strconv.Atoi per line.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// TODO: Open numbers.txt, defer f.Close(), scan it with a bufio.Scanner,
	//       accumulate the parsed numbers, and print "sum: <total>".
	//       Print the error and return early if opening, parsing, or scanning fails.
}
