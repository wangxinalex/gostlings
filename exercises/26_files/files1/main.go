// Concept: reading files with os.ReadFile
// Task: read the file data.txt and print its contents; print the error on failure
// Expected output: hello from file
// Hint: the os package reads a whole file into bytes and returns an error
//       separately; print the bytes as a string, and use fmt.Print rather than
//       Println because data.txt already ends with a newline. NOTE: this exercise
//       reads a file next to main.go, so run it from its own directory
//       (`cd exercises/26_files/files1 && go run .`) or verify with
//       `go test ./exercises/26_files/files1` (Go doc: os)
// Stuck?: os.ReadFile returns the bytes and an error.

package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Read data.txt with os.ReadFile. Print the file contents with
	//       fmt.Print on success, or print the error and return early on failure.
}
