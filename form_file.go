package bulma

import (
	"github.com/maragudk/gomponents/html"
)

const (
	fileNameAutoUpdateScript = `if (this.files.length>0) { this.parentNode.getElementsByClassName("file-name")[0].textContent = this.files[0].name }`
)

// FileName is a modifier that, when applied to File, defines the default value
// for the file-name element.
//
// https://willoma.github.io/bulma-gomponents/form/file.html
type FileName string

// FileNameAutoUpdate is a modifier that, when applied to File, defines the
// default value for the file-name element, and adds the needed javascript
// event to update the file name.
//
// https://willoma.github.io/bulma-gomponents/form/file.html
type FileNameAutoUpdate string

// File creates a file input element.
//
// https://willoma.github.io/bulma-gomponents/form/file.html
func File(children ...any) Element {
	input := Elem(html.Input, Class("file-input"), html.Type("file"))
	cta := Elem(html.Span, Class("file-cta"))
	label := Elem(
		html.Label,
		Class("file-label"),
		input,
		cta,
	)
	f := &file{
		Element: Elem(html.Div, Class("file"), label),
		label:   label,
		input:   input,
		cta:     cta,
	}
	f.With(children...)
	return f
}

type file struct {
	Element
	label Element
	input Element
	cta   Element
}

func (f *file) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCTA:
			f.cta.With(c...)
		case onDiv:
			f.Element.With(c...)
		case onInput:
			f.input.With(c...)
		case Class, Classer, Classeser, ExternalClassesAndStyles, MultiClass, Styles:
			f.Element.With(c)
		case string:
			f.cta.With(Elem(html.Span, Class("file-label"), c))
		case FileName:
			f.Element.With(Class("has-name"))
			f.label.With(Elem(html.Span, Class("file-name"), string(c)))
		case FileNameAutoUpdate:
			f.Element.With(Class("has-name"))
			f.label.With(Elem(html.Span, Class("file-name"), string(c)))
			f.input.With(On("change", fileNameAutoUpdateScript))
		case IconElem:
			c.SetIconClass(Class("file-icon"))
			f.cta.With(c)
		case Element:
			f.cta.With(c)
		case []any:
			f.With(c...)
		default:
			f.input.With(c)
		}
	}

	return f
}
