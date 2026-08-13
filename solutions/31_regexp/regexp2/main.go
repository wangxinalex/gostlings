// Concept: compiling a regexp and finding all matches
// Task: complete findNumbers so it returns every digit sequence in s
// Expected output: [123 45]
// Hint: re, err := regexp.Compile(`\d+`); re.FindAllString(s, -1) (Go doc: regexp)

package main

import (
	"fmt"
	"regexp"
)

func findNumbers(s string) []string {
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		return nil
	}
	return re.FindAllString(s, -1)
}

func main() {
	fmt.Println(findNumbers("a=123 b=45"))
}
