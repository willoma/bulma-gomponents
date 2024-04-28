package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Radio creates a radio element, together with its label container.
//
// https://willoma.github.io/bulma-gomponents/form/radio.html
func Radio(children ...any) Element {
	input := Elem(html.Input, html.Type("radio"))
	r := &radio{
		Element: Elem(html.Label, Class("radio"), input, " "),
		input:   input,
	}
	r.With(children...)
	return r
}

// Checked, when provided as a child of Radio, makes it so the
// radio button is checked.
//
// https://willoma.github.io/bulma-gomponents/form/radio.html
var Checked = gomponents.Attr("checked")

type radio struct {
	Element
	input Element
}

func (r *radio) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			r.input.With(c...)
		case onLabel:
			r.Element.With(c...)
		case string:
			r.Element.With(c)
		case Class:
			switch c {
			case Disabled:
				r.input.With(html.Disabled())
				r.Element.With(html.Disabled())
			default:
				r.input.With(c)
			}
		case gomponents.Node:
			if isAttribute(c) {
				r.input.With(c)
			} else {
				r.Element.With(c)
			}
		case Element:
			r.Element.With(c)
		case []any:
			r.With(c...)
		default:
			r.input.With(c)
		}
	}

	return r
}

func (r *radio) Clone() Element {
	return &radio{
		Element: r.Element.Clone(),
		input:   r.input.Clone(),
	}
}
