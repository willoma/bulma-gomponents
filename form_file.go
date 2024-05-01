package bulma

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
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
func File(children ...any) e.Element {
	input := e.Input(e.Class("file-input"), html.Type("file"))
	cta := e.Span(e.Class("file-cta"))
	label := e.Label(
		e.Class("file-label"),
		input,
		cta,
	)
	f := &file{
		Element: e.Div(e.Class("file"), label),
		label:   label,
		input:   input,
		cta:     cta,
	}
	f.With(children...)
	return f
}

type file struct {
	e.Element
	label e.Element
	input e.Element
	cta   e.Element
}

func (f *file) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCTA:
			f.cta.With(c...)
		case onDiv:
			f.Element.With(c...)
		case onInput:
			f.input.With(c...)
		case e.Class, e.Classer, e.Classeser, e.Styles:
			f.Element.With(c)
		case string:
			f.cta.With(e.Span(e.Class("file-label"), c))
		case FileName:
			f.Element.With(e.Class("has-name"))
			f.label.With(e.Span(e.Class("file-name"), string(c)))
		case FileNameAutoUpdate:
			f.Element.With(e.Class("has-name"))
			f.label.With(e.Span(e.Class("file-name"), string(c)))
			f.input.With(e.On("change", fileNameAutoUpdateScript))
		case IconElem:
			c.SetIconClass(e.Class("file-icon"))
			f.cta.With(c)
		case e.Element:
			f.cta.With(c)
		case []any:
			f.With(c...)
		default:
			f.input.With(c)
		}
	}

	return f
}

func (f *file) Clone() e.Element {
	return &file{
		Element: f.Element.Clone(),
		label:   f.label.Clone(),
		input:   f.input.Clone(),
		cta:     f.cta.Clone(),
	}
}
