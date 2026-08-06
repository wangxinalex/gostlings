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
	if err := os.WriteFile("demo.txt", []byte("hello, disk!"), 0o644); err != nil {
		fmt.Println(err)
		return
	}
	data, err := os.ReadFile("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	if err := os.Remove("demo.txt"); err != nil {
		fmt.Println(err)
	}
}
