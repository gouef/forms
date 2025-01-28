package forms

import "github.com/gouef/html"

type BaseInput interface {
	GetName() string
	GetLabel() *string
	isRequired()
	GetHtmlElement() *html.Html
	GetLabelElement() *html.Html
}
