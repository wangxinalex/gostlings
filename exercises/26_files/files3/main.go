// Concept: bufio.Scanner reads input line by line
// Task: open numbers.txt, scan it line by line, and print the sum of the numbers
// Expected output: sum: 6
// Hint: sc := bufio.NewScanner(f); for sc.Scan() { ... strconv.Atoi(sc.Text()) ... };
//       check sc.Err() after the loop, and defer f.Close() right after opening.
//       Run from this directory (`cd exercises/26_files/files3 && go run .`) or
//       verify with `go test ./exercises/26_files/files3` (Go doc: bufio)

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
