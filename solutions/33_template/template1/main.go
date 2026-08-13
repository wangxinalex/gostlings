// Concept: rendering text templates with text/template
// Task: complete render so it fills the {{.Name}} placeholder
// Expected output: Hello, Ada!
// Hint: template.New("greeting").Parse(tmplText); tmpl.Execute(&buf, data) (Go doc: text/template)

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
	tmpl, err := template.New("render").Parse(tmplText)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	out, err := render("Hello, {{.Name}}!", person{Name: "Ada"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
