package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Radio creates a radio element, together with its label container.
//   - when a child is marked with b.Inner, it is forcibly applied to the <input> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <label> element
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
func Radio(children ...any) Element {
	return new(radio).With(children...)
}

// RadioDisabled creates a disabled radio element, together with its label
// container.
//   - when a child is marked with b.Inner, it is forcibly applied to the <input> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <label> element
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
func RadioDisabled(children ...any) Element {
	return (&radio{disabled: true}).With(children...)
}

// Checked, when provided as a child of Radio or RadioDisabled, makes it so the
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
		case *ApplyToInner:
			r.inputChildren = append(r.inputChildren, c.Child)
		case *ApplyToOuter:
			r.labelChildren = append(r.labelChildren, c.Child)
		case string:
			r.labelChildren = append(r.labelChildren, c)
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
	input := Elem(html.Input, html.Type("radio"), r.inputChildren)
	label := Elem(html.Label, Class("radio"), input, " ", r.labelChildren)

	if r.disabled {
		label.With(html.Disabled())
		input.With(html.Disabled())
	}

	return label.Render(w)
}
