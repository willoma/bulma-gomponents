package bulma

import (
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
func Radio(children ...any) Element {
	r := &radio{}
	r.addChildren(children)
	return r.elem()
}

// RadioDisabled creates a disabled radio element, together with its label
// container.
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
	r := &radio{disabled: true}
	r.addChildren(children)
	return r.elem()
}

// Checked, when provided as a child of Radio or RadioDisabled, makes it so the
// radio button is checked.
var Checked = gomponents.Attr("checked")

type radio struct {
	disabled      bool
	labelChildren []any
	inputChildren []any
}

func (r *radio) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
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
			r.addChildren(c)
		default:
			r.inputChildren = append(r.inputChildren, c)
		}
	}
}

func (r *radio) elem() Element {
	input := Elem(html.Input, html.Type("radio"), r.inputChildren)
	label := Elem(html.Label, Class("radio"), input, " ", r.labelChildren)

	if r.disabled {
		label.With(html.Disabled())
		input.With(html.Disabled())
	}

	return label
}
