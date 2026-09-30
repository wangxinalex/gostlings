// Concept: ranging over a slice in a template
// Task: complete render so it joins words with commas using a range action
// Expected output: go,rust,python
// Hint: iterate the slice with the template's range action and print a comma
//       before every element except the first; compare the index against zero in
//       an if action to decide (Go doc: text/template)
// Stuck?: the range action can bind both an index and a value variable.

package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func render(words []string) (string, error) {
	const tmplText = `{{range $i, $v := .}}{{if $i}},{{end}}{{$v}}{{end}}`
	// TODO: Parse tmplText, execute it with words, and return the rendered string.
	return "", nil
}

func main() {
	out, err := render([]string{"go", "rust", "python"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
