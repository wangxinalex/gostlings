// Concept: reading files with os.ReadFile
// Task: read the file data.txt and print its contents; print the error on failure
// Expected output: hello from file
// Hint: os.ReadFile("data.txt") returns ([]byte, error); print string(data) on
//       success. data.txt already ends with a newline, so use fmt.Print — not
//       Println — or the output gets a blank line. NOTE: this exercise reads a
//       file next to main.go, so run it from its own directory
//       (`cd exercises/26_files/files1 && go run .`) or verify with
//       `go test ./exercises/26_files/files1` (Go doc: os)

package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(data))
}
