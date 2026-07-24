// Concept: deferred cleanup always runs — even after a panic
// Task: the file should be closed whether or not a panic occurs; add deferred cleanup
// Expected output: file opened
// file closed
// Hint: defer runs even when a panic happens; use a named return to set the error in the deferred func if needed (Go doc: builtin)

package main

import "fmt"

type File struct{ Name string }

func (f *File) Close() { fmt.Println("file closed") }

func processFile() {
	f := &File{Name: "data.txt"}
	defer f.Close()
	fmt.Println("file opened")
	panic("processing error")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			// Swallow the panic so we can see the output.
		}
	}()
	processFile()
}
