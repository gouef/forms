package forms

import "github.com/gouef/html"

type Form struct {
	items    []interface{}
	renderer Renderer
}

func NewForm() *Form {
	renderer := NewSimpleRenderer()
	return &Form{
		items:    []interface{}{},
		renderer: renderer,
	}
}

func (f *Form) AddSelect() {

}

func (f *Form) SetRenderer(renderer Renderer) *Form {
	f.renderer = renderer
	return f
}

func (f *Form) GetRenderer() Renderer {
	res := html.El("form")
	res.Render()
	return f.renderer
}
