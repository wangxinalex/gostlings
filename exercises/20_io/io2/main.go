// Concept: io.Copy streams data from a Reader to a Writer
// Task: use io.Copy to copy from the reader to the builder, then print the builder's content
// Expected output: streamed: hello world
// Hint: let the io package move the bytes until the source ends; the destination
//       only has to satisfy the writer interface (Go doc: io)
// Stuck?: io.Copy takes the destination first, then the source.

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	var sb strings.Builder

	// TODO: Use io.Copy to copy from r into &sb.
	//       Then print with the prefix "streamed: ".

	fmt.Println("not implemented")
}
