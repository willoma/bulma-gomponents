package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func radioElem(disabled bool, children []any) *Element {
	input := Elem(html.Input).
		With(html.Type("radio"))

	label := Elem(html.Label).
		With(Class("radio")).
		With(input).
		With(" ")

	if disabled {
		label.With(html.Disabled())
		input.With(html.Disabled())
	}

	for _, c := range children {
		switch c := c.(type) {
		case string:
			label.With(c)
		case gomponents.Node:
			if IsAttribute(c) {
				input.With(c)
			} else {
				label.With(c)
			}
		case *Element, container:
			label.With(c)
		default:
			input.With(c)
		}
	}

	return label
}

// Radio creates a radio element, together with its label container.
//   - when a child is a string, it is added in the radio label
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the input element
//   - when a child is a gopmponents.Node with another type, it is added as a
//     child to the label element
//   - when a child is an *Element, it is added in the radio label
//   - when a child is a container, it is added in the radio label
//   - other childs are added as children to the input element
//
// The following modifiers change the radio behaviour:
//   - Checked: make it so the radio button is checked
func Radio(children ...any) *Element {
	return radioElem(false, children)
}

// RadioDisabled creates a disabled radio element, together with its label
// container.
//   - when a child is a string, it is added in the radio label
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the input element
//   - when a child is a gopmponents.Node with another type, it is added as a
//     child to the label element
//   - when a child is an *Element, it is added in the radio label
//   - when a child is a container, it is added in the radio label
//   - other childs are added as children to the input element
//
// The following modifiers change the radio behaviour:
//   - Checked: make it so the radio button is checked
func RadioDisabled(children ...any) *Element {
	return radioElem(true, children)
}

// Checked, when provided as a child of Radio or RadioDisabled, makes it so the
// radio button is checked.
var Checked = gomponents.Attr("checked")
