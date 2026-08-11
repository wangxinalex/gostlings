// Concept: sync.Pool reuses temporary objects without making reuse a correctness requirement.
// Task: borrow a bytes.Buffer, reset it, write value, and return the buffer to the pool.
// Hint: Get, Reset, use, Put; callers should receive a copied string before Put.
package main

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}
var getBuffer = func() *bytes.Buffer { return bufferPool.Get().(*bytes.Buffer) }
var putBuffer = func(buffer *bytes.Buffer) { bufferPool.Put(buffer) }

func reuseBuffer(value string) string {
	buffer := getBuffer()
	buffer.Reset()
	buffer.WriteString(value)
	result := buffer.String()
	putBuffer(buffer)
	return result
}

func main() {}
