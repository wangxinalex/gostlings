// Concept: extracting captured groups with FindStringSubmatch
// Task: complete parseDate so it returns year, month, and day from "YYYY-MM-DD"
// Expected output: 2026 08 13
// Hint: capture the three parts in the pattern and read them back as a slice of
//       submatches, where element zero is the whole match (Go doc: regexp)
// Stuck?: regexp.MustCompile for a constant pattern; the submatch method returns
//       the full match followed by the three groups.

package main

import (
	"fmt"
	"regexp"
)

func parseDate(s string) (year, month, day string) {
	// TODO: Return the three captured groups.
	return "", "", ""
}

func main() {
	y, m, d := parseDate("2026-08-13")
	fmt.Println(y, m, d)
}
