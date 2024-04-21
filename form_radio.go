package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Radio creates a radio element, together with its label container.
//   - when a child is a string, it is added in the radio label
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the input element
//   - when a child is a gopmponents.Node with another type, it is added as a
//     child to the label element
//   - when a child is an Element, it is added in the radio label
//   - other childs are added as children to the input element
//
// The following modifiers change the radio behaviour:
//   - Checked: make it so the radio button is checked
//   - Disabled: disable the radio button
func Radio(children ...any) Element {
	return new(radio).With(children...)
}

// Checked, when provided as a child of Radio, makes it so the
// radio button is checked.
var Checked = gomponents.Attr("checked")

type radio struct {
	disabled      bool
	labelChildren []any
	inputChildren []any
}

func (r *radio) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			r.inputChildren = append(r.inputChildren, c...)
		case onLabel:
			r.labelChildren = append(r.labelChildren, c...)
		case string:
			r.labelChildren = append(r.labelChildren, c)
		case Class:
			switch c {
			case Disabled:
				r.inputChildren = append(r.inputChildren, html.Disabled())
				r.labelChildren = append(r.labelChildren, html.Disabled())
			default:
				r.inputChildren = append(r.inputChildren, c)
			}
		case gomponents.Node:
			if IsAttribute(c) {
				r.inputChildren = append(r.inputChildren, c)
			} else {
				r.labelChildren = append(r.labelChildren, c)
			}
		case Element:
			r.labelChildren = append(r.labelChildren, c)
		case []any:
			r.With(c...)
		default:
			r.inputChildren = append(r.inputChildren, c)
		}
	}

	return r
}

func (r *radio) Render(w io.Writer) error {
	return Elem(
		html.Label,
		Class("radio"),
		Elem(
			html.Input,
			html.Type("radio"),
			r.inputChildren,
		),
		" ",
		r.labelChildren,
	).Render(w)
}
