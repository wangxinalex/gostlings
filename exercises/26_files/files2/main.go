// Concept: writing files with os.WriteFile
// Task: write "hello, disk!" to demo.txt, read it back, print the contents,
//       then remove the file
// Expected output: hello, disk!
// Hint: the os package writes bytes with a permission mode, reads them back, and
//       removes the file; cleanup is part of the task. Run from this directory
//       (`cd exercises/26_files/files2 && go run .`) or verify with
//       `go test ./exercises/26_files/files2` (Go doc: os)
// Stuck?: os.WriteFile takes the bytes and a mode literal such as 0o644.

package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Write "hello, disk!" to demo.txt with os.WriteFile (mode 0o644),
	//       read it back with os.ReadFile, print the contents, then remove
	//       demo.txt with os.Remove. Print the error and return early on failure.
}
