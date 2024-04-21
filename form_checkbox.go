package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Checkbox creates a checkbox input element.
func Checkbox(children ...any) Element {
	return new(checkbox).With(children...)
}

type checkbox struct {
	labelChildren []any
	inputChildren []any
}

func (cb *checkbox) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			cb.inputChildren = append(cb.inputChildren, c...)
		case onLabel:
			cb.labelChildren = append(cb.labelChildren, c...)
		case string:
			cb.labelChildren = append(cb.labelChildren, c)
		case Class:
			switch c {
			case Disabled:
				cb.inputChildren = append(cb.inputChildren, html.Disabled())
				cb.labelChildren = append(cb.labelChildren, html.Disabled())
			default:
				cb.inputChildren = append(cb.inputChildren, c)
			}
		case gomponents.Node:
			if IsAttribute(c) {
				cb.inputChildren = append(cb.inputChildren, c)
			} else {
				cb.labelChildren = append(cb.labelChildren, c)
			}
		case Element:
			cb.labelChildren = append(cb.labelChildren, c)
		case []any:
			cb.With(c...)
		default:
			cb.inputChildren = append(cb.inputChildren, c)
		}
	}

	return cb
}

func (cb *checkbox) Render(w io.Writer) error {
	return Elem(
		html.Label,
		Class("checkbox"),
		Elem(
			html.Input,
			html.Type("checkbox"),
			cb.inputChildren,
		),
		" ",
		cb.labelChildren,
	).Render(w)
}
