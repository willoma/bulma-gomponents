package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Checkbox creates a checkbox input element.
//   - when a child is marked with b.Inner, it is forcibly applied to the <input> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <label> element
func Checkbox(children ...any) Element {
	return new(checkbox).With(children...)
}

// CheckboxDisabled creates a disabled checkbox input element.
//   - when a child is marked with b.Inner, it is forcibly applied to the <input> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <label> element
func CheckboxDisabled(children ...any) Element {
	return (&checkbox{disabled: true}).With(children...)
}

type checkbox struct {
	disabled      bool
	labelChildren []any
	inputChildren []any
}

func (cb *checkbox) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *ApplyToInner:
			cb.inputChildren = append(cb.inputChildren, c.Child)
		case *ApplyToOuter:
			cb.labelChildren = append(cb.labelChildren, c.Child)
		case string:
			cb.labelChildren = append(cb.labelChildren, c)
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

	return label.Render(w)
}
