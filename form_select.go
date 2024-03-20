package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type Size int

type Name string

// Select creates a dropdown select element. It should contain one or multiple
// Option elements.
//   - when a child is a Class, it is applied to the select element if it is
//     Hovered or Focused, otherwise it is applied to the div element
//   - when a child is an Element, it is added as a child to the select element
//   - when a child is a Name, it is used as the name attribute of the select
//     element
//   - other children are added as children to the div element
//
// The following modifiers change the select behaviour:
//   - Rounded: rounded select
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the select (in place of
//     the arrow)
//
// The following modifiers change the select color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the select size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func Select(children ...any) Element {
	s := &selectEl{}
	s.addChildren(children)
	return s.elem()
}

// SelectMultiple creates a multiple select element. It should contain one or
// multiple Option elements.
//   - when a child is a Size, it defines the select size
//   - when a child is a Class, it is applied to the select element if it is
//     Hovered or Focused, otherwise it is applied to the div element
//   - when a child is an Element, it is added as a child to the select element
//   - when a child is a Name, it is used as the name attribute of the select
//     element
//   - other children types are added as children to the div element
//
// The following modifiers change the select behaviour:
//   - Rounded: rounded select
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the select (in place of
//     the arrow)
//
// The following modifiers change the select color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the select size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func SelectMultiple(children ...any) Element {
	s := &selectEl{multiple: true}
	s.addChildren(children)
	return s.elem()
}

type selectEl struct {
	multiple       bool
	selectChildren []any
	divChildren    []any
}

func (s *selectEl) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Size:
			s.selectChildren = append(
				s.selectChildren,
				gomponents.Attr(
					"size",
					strconv.Itoa(int(c)),
				),
			)
		case Class:
			switch c {
			case Hovered, Focused:
				s.selectChildren = append(s.selectChildren, c)
			default:
				s.divChildren = append(s.divChildren, c)
			}
		case Element:
			s.selectChildren = append(s.selectChildren, c)
		case Name:
			s.selectChildren = append(s.selectChildren, html.Name(string(c)))
		case []any:
			s.addChildren(c)
		default:
			s.divChildren = append(s.divChildren, c)
		}
	}
}

func (s *selectEl) elem() Element {
	div := Elem(html.Div, Class("select"), s.divChildren)
	sel := Elem(html.Select, s.selectChildren...)

	if s.multiple {
		div.With(Class("is-multiple"))
		sel.With(html.Multiple())
	}

	return div.With(sel)
}

// Option creates an option element, to be used as a child of a Select or
// SelectMultiple. The value argument is used as the option value attribute.
func Option(value string, children ...any) Element {
	return Elem(html.Option, html.Value(value), children)
}

// OptionSelected creates a selected option element, to be used as a child of a
// Select or SelectMultiple. The value argument is used as the option value
// attribute.
func OptionSelected(value string, children ...any) Element {
	return Elem(html.Option, html.Value(value), html.Selected(), children)
}
