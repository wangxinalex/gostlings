// Concept: composing interfaces from smaller interfaces
// Task: declare Reader, Closer, and ReadCloser, then write drainAndClose
// Expected output: go is fun
// closed: true
// 3 2 1
// closed: true
// Hint: an interface can embed another interface, so ReadCloser needs no methods of its own;
//       a value satisfies ReadCloser only when it implements Read and Close. Read until the
//       reader reports that it is exhausted, then close it (Go Tour: Methods 9-11)

package main

import "fmt"

type wordReader struct {
	words  []string
	index  int
	closed bool
}

func (w *wordReader) Read() (string, bool) {
	if w.index >= len(w.words) {
		return "", false
	}
	word := w.words[w.index]
	w.index++
	return word, true
}

func (w *wordReader) Close() error {
	w.closed = true
	return nil
}

type countdown struct {
	next   int
	closed bool
}

func (c *countdown) Read() (string, bool) {
	if c.next <= 0 {
		return "", false
	}
	value := fmt.Sprintf("%d", c.next)
	c.next--
	return value, true
}

func (c *countdown) Close() error {
	c.closed = true
	return nil
}

// TODO: Declare Reader with one Read method that returns the next value and a bool
//       reporting whether a value was produced.

// TODO: Declare Closer with one Close method returning an error.

// TODO: Declare ReadCloser by embedding Reader and Closer. A value satisfies ReadCloser
//       only when it implements both of those methods.

// TODO: Write drainAndClose(rc ReadCloser) string: read values until the reader is
//       exhausted, close it, and return the collected values joined by single spaces.

func main() {
	words := &wordReader{words: []string{"go", "is", "fun"}}
	fmt.Println(drainAndClose(words))
	fmt.Println("closed:", words.closed)

	numbers := &countdown{next: 3}
	fmt.Println(drainAndClose(numbers))
	fmt.Println("closed:", numbers.closed)
}
