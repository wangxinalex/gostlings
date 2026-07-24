// Concept: deferred cleanup runs even when a panic unwinds the stack
// Task: the file is never closed because the function panics; add ONE line so the file is always closed
// Expected output: file opened: data.txt
// file closed
// Hint: defer schedules a call to run when the surrounding function returns — even on panic (Go Tour: Flowcontrol 12)

package main

import "fmt"

type File struct{ Name string }

func (f *File) Close() { fmt.Println("file closed") }

func processFile() {
	f := &File{Name: "data.txt"}
	// TODO: Make sure the file gets closed even though the code below panics.
	fmt.Println("file opened:", f.Name)
	panic("processing error")
}

func main() {
	defer func() {
		recover() // Swallow the panic so the program exits normally.
	}()
	processFile()
}
