package forms

import "github.com/gouef/html"

type FormElement interface {
	GetHtmlElement() *html.Html
}

type Container interface {
	Render() string
	GetForm() html.Html
	AddText(name string, label *string)
	AddPassword(name string, label *string)
	AddTextArea(name string, label *string)
	AddEmail(name string, label *string)
	AddInteger(name string, label *string)
	AddFloat(name string, label *string)
	AddDate(name string, label *string)
	AddTime(name string, label *string, inSeconds *bool)
	AddDateTime(name string, label *string, inSeconds *bool)
	AddUpload(name string, label *string)
	AddMultiUpload(name string, label *string)
	AddHidden(name string, defaultValue *string)
	AddCheckbox(name string, caption *string)
	AddRadioList(name string, label *string, items []string)
	AddCheckboxList(name string, label *string, items []string)
	AddSelect(name string, label *string, items []string)
	AddMultiSelect(name string, label *string, items []string)
	AddColor(name string, label *string)
	AddSubmit(name string, caption *string)
	AddButton(name string, caption *string)
	AddImageButton(name string, src *string, alt *string)
}
