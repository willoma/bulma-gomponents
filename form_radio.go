package bulma

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// Radios creates a list of radio elements.
//
// https://willoma.github.io/bulma-gomponents/form/radio.html
func Radios(children ...any) e.Element {
	return e.Div(e.Class("radios"), children)
}

// Radio creates a radio element, together with its label container.
//
// https://willoma.github.io/bulma-gomponents/form/radio.html
func Radio(children ...any) e.Element {
	input := e.Input(html.Type("radio"))
	r := &radio{
		Element: e.Label(e.Class("radio"), input, " "),
		input:   input,
	}
	r.With(children...)
	return r
}

// Checked, when provided as a child of Radio or Checkbox, makes it so the
// radio button is checked.
//
// https://willoma.github.io/bulma-gomponents/form/radio.html
var Checked = gomponents.Attr("checked")

type radio struct {
	e.Element
	input e.Element
}

func (r *radio) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			r.input.With(c...)
		case onLabel:
			r.Element.With(c...)
		case string:
			r.Element.With(c)
		case e.Class:
			switch c {
			case Disabled:
				r.input.With(html.Disabled())
				r.Element.With(html.Disabled())
			default:
				r.input.With(c)
			}
		case gomponents.Node:
			if e.IsAttribute(c) {
				r.input.With(c)
			} else {
				r.Element.With(c)
			}
		case []any:
			r.With(c...)
		default:
			r.input.With(c)
		}
	}

	return r
}

func (r *radio) Clone() e.Element {
	return &radio{
		Element: r.Element.Clone(),
		input:   r.input.Clone(),
	}
}
