// Concept: io.Copy streams data from a Reader to a Writer
// Task: use io.Copy to copy from the reader to the builder, then print the builder's content
// Expected output: streamed: hello world
// Hint: io.Copy(dst, src) copies until EOF; strings.Builder implements io.Writer (Go doc: io)

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	var sb strings.Builder

	io.Copy(&sb, r)
	fmt.Println("streamed:", sb.String())
}
