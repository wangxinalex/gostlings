// Concept: ranging over a slice in a template
// Task: complete render so it joins words with commas using a range action
// Expected output: go,rust,python
// Hint: parse `{{range $i, $v := .}}{{if $i}},{{end}}{{$v}}{{end}}` (Go doc: text/template)

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
