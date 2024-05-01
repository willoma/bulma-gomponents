package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Checkbox creates a checkbox input element.
//
// https://willoma.github.io/bulma-gomponents/form/checkbox.html
func Checkbox(children ...any) e.Element {
	input := e.Input(html.Type("checkbox"))
	cb := &checkbox{
		Element: e.Label(
			e.Class("checkbox"),
			input,
			" ",
		),
		input: input,
	}
	cb.With(children...)
	return cb
}

type checkbox struct {
	e.Element
	input e.Element
}

func (cb *checkbox) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			cb.input.With(c...)
		case onLabel:
			cb.Element.With(c...)
		case string:
			cb.Element.With(c)
		case e.Class:
			switch c {
			case Disabled:
				cb.input.With(html.Disabled())
				cb.Element.With(html.Disabled())
			default:
				cb.input.With(c)
			}
		case gomponents.Node:
			if e.IsAttribute(c) {
				cb.input.With(c)
			} else {
				cb.Element.With(c)
			}
		case e.Element:
			cb.Element.With(c)
		case []any:
			cb.With(c...)
		default:
			cb.input.With(c)
		}
	}

	return cb
}

func (cb *checkbox) Clone() e.Element {
	return &checkbox{
		Element: cb.Element.Clone(),
		input:   cb.input.Clone(),
	}
}
