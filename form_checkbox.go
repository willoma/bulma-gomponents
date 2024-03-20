package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Checkbox creates a checkbox input element.
func Checkbox(children ...any) *Element {
	cb := &checkbox{}
	cb.addChildren(children)
	return cb.elem()
}

// CheckboxDisabled creates a disabled checkbox input element.
func CheckboxDisabled(children ...any) *Element {
	cb := &checkbox{disabled: true}
	cb.addChildren(children)
	return cb.elem()
}

type checkbox struct {
	disabled      bool
	labelChildren []any
	inputChildren []any
}

func (cb *checkbox) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case string:
			cb.labelChildren = append(cb.labelChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				cb.inputChildren = append(cb.inputChildren, c)
			} else {
				cb.labelChildren = append(cb.labelChildren, c)
			}
		case *Element, container:
			cb.labelChildren = append(cb.labelChildren, c)
		case []any:
			cb.addChildren(c)
		default:
			cb.inputChildren = append(cb.inputChildren, c)
		}
	}
}

func (cb *checkbox) elem() *Element {
	input := Elem(html.Input, html.Type("checkbox"), cb.inputChildren)

	label := Elem(
		html.Label,
		Class("checkbox"),
		input,
		" ",
		cb.labelChildren,
	)

	if cb.disabled {
		label.With(html.Disabled())
		input.With(html.Disabled())
	}

	return label
}
