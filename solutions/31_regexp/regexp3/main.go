// Concept: extracting captured groups with FindStringSubmatch
// Task: complete parseDate so it returns year, month, and day from "YYYY-MM-DD"
// Expected output: 2026 08 13
// Hint: re := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`); re.FindStringSubmatch(s) (Go doc: regexp)

package main

import (
	"fmt"
	"regexp"
)

func parseDate(s string) (year, month, day string) {
	re := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	m := re.FindStringSubmatch(s)
	if len(m) != 4 {
		return "", "", ""
	}
	return m[1], m[2], m[3]
}

func main() {
	y, m, d := parseDate("2026-08-13")
	fmt.Println(y, m, d)
}
