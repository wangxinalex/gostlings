// Concept: writing files with os.WriteFile
// Task: write "hello, disk!" to demo.txt, read it back, print the contents,
//       then remove the file
// Expected output: hello, disk!
// Hint: os.WriteFile("demo.txt", []byte("hello, disk!"), 0o644) creates or
//       overwrites the file; read it back with os.ReadFile, then clean up with
//       os.Remove. Run from this directory (`cd exercises/26_files/files2 &&
//       go run .`) or verify with `go test ./exercises/26_files/files2` (Go doc: os)

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
