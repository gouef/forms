package forms

import (
	_ "embed"
	"html/template"
)

type Renderer interface {
	Render() template.HTML
}

type SimpleRenderer struct {
}

func NewSimpleRenderer() Renderer {
	return &SimpleRenderer{}
}

func (r *SimpleRenderer) Render() template.HTML {
	return template.HTML("")
}

//go:embed templates/form.gohtml
var formTemplate string
