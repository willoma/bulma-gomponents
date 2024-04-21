package bulma

import (
	"io"
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Select creates a dropdown select element. It should contain one or multiple
// Option elements.
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
	return new(selectEl).With(children...)
}

type selectEl struct {
	multiple       bool
	selectChildren []any
	divChildren    []any
}

func (s *selectEl) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onSelect:
			s.selectChildren = append(s.selectChildren, c...)
		case onDiv:
			s.divChildren = append(s.divChildren, c...)
		case Class:
			switch c {
			case Hovered, Focused:
				s.selectChildren = append(s.selectChildren, c)
			case Multiple:
				s.divChildren = append(s.divChildren, Multiple)
				s.selectChildren = append(s.selectChildren, html.Multiple())
			case Disabled:
				s.selectChildren = append(s.selectChildren, html.Disabled())
			default:
				s.divChildren = append(s.divChildren, c)
			}
		case []any:
			s.With(c...)
		default:
			s.selectChildren = append(s.selectChildren, c)
		}
	}

	return s
}

func (s *selectEl) Render(w io.Writer) error {
	return Elem(
		html.Div,
		Class("select"),
		s.divChildren,
		Elem(html.Select, s.selectChildren...),
	).Render(w)
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

func Size(size int) gomponents.Node {
	return gomponents.Attr("size", strconv.Itoa(size))
}
