// Concept: the io.Reader interface — reading data from any source
// Task: use io.ReadAll to read all data from the reader and print it as a string
// Expected output: Hello, io.Reader!
// Hint: io.ReadAll(r) reads everything from an io.Reader into a []byte (Go doc: io)

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, io.Reader!")
	// TODO: Use io.ReadAll to read all data from r, then print it as a string.
	fmt.Println("not implemented")
}
