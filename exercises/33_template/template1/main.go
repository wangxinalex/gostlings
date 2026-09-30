// Concept: rendering text templates with text/template
// Task: complete render so it fills the {{.Name}} placeholder
// Expected output: Hello, Ada!
// Hint: parse the template text, execute it with the data into a buffer, and
//       return the buffer's contents; both steps can fail (Go doc: text/template)
// Stuck?: template.New's Parse and Execute methods; bytes.Buffer satisfies io.Writer.

package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type person struct {
	Name string
}

func render(tmplText string, data any) (string, error) {
	// TODO: Parse tmplText, execute it with data, and return the rendered string.
	return "", nil
}

func main() {
	out, err := render("Hello, {{.Name}}!", person{Name: "Ada"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
