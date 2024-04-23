package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Checkbox creates a checkbox input element.
//
// https://willoma.github.io/bulma-gomponents/form/checkbox.html
func Checkbox(children ...any) Element {
	input := Elem(html.Input, html.Type("checkbox"))
	cb := &checkbox{
		Element: Elem(
			html.Label,
			Class("checkbox"),
			input,
			" ",
		),
		input: input,
	}
	cb.With(children...)
	return cb
}

type checkbox struct {
	Element
	input Element
}

func (cb *checkbox) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			cb.input.With(c...)
		case onLabel:
			cb.Element.With(c...)
		case string:
			cb.Element.With(c)
		case Class:
			switch c {
			case Disabled:
				cb.input.With(html.Disabled())
				cb.Element.With(html.Disabled())
			default:
				cb.input.With(c)
			}
		case gomponents.Node:
			if isAttribute(c) {
				cb.input.With(c)
			} else {
				cb.Element.With(c)
			}
		case Element:
			cb.Element.With(c)
		case []any:
			cb.With(c...)
		default:
			cb.input.With(c)
		}
	}

	return cb
}
