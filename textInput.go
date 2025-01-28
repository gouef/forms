package forms

import "github.com/gouef/html"

type TextInput struct {
	name         string
	label        *string
	htmlElement  *html.Html
	labelElement *html.Html
	required     bool
}

func NewTextInput(name string, label *string) *TextInput {
	htmlElement := html.El("input")
	htmlElement.AddAttribute("type", "")
	return &TextInput{
		name:         name,
		label:        label,
		htmlElement:  html.El("input"),
		labelElement: html.El("label").AddString(*label),
		required:     false,
	}
}

func (t *TextInput) GetName() string {
	return t.name
}

func (t *TextInput) GetLabel() *string {
	return t.label
}

func (t *TextInput) isRequired() {
	t.required = true
}

func (t *TextInput) GetHtmlElement() *html.Html {
	return t.htmlElement
}

func (t *TextInput) GetLabelElement() *html.Html {
	return t.labelElement
}
