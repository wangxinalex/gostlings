// Concept: writing formatted output to any io.Writer with fmt.Fprintf
// Task: complete greet so it writes "Hello, Ada!" to w
// Expected output: Hello, Ada!
// Hint: fmt.Fprintf(w, "Hello, %s!", name) (Go doc: fmt)

package main

import (
	"fmt"
	"io"
	"os"
)

func greet(w io.Writer, name string) error {
	// TODO: Write the greeting to w and return any error.
	return nil
}

func main() {
	if err := greet(os.Stdout, "Ada"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println()
}
