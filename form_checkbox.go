package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func checkboxElem(disabled bool, children []any) *Element {
	input := Elem(html.Input).
		With(html.Type("checkbox"))

	label := Elem(html.Label).
		With(Class("checkbox")).
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

// Checkbox creates a checkbox input element.
func Checkbox(children ...any) *Element {
	return checkboxElem(false, children)
}

// CheckboxDisabled creates a disabled checkbox input element.
func CheckboxDisabled(children ...any) *Element {
	return checkboxElem(true, children)
}
