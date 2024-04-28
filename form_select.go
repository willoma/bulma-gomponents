package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Select creates a dropdown select element. It should contain one or multiple
// Option elements.
//
// https://willoma.github.io/bulma-gomponents/form/select.html
func Select(children ...any) Element {
	selectHTMLElement := Elem(html.Select)
	s := &selectEl{
		Element:           Elem(html.Div, Class("select"), selectHTMLElement),
		selectHTMLElement: selectHTMLElement,
	}
	s.With(children...)
	return s
}

type selectEl struct {
	Element
	selectHTMLElement Element
}

func (s *selectEl) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onSelect:
			s.selectHTMLElement.With(c...)
		case onDiv:
			s.Element.With(c...)
		case *option:
			s.selectHTMLElement.With(c)
		case Class:
			switch c {
			case Hovered, Focused:
				s.selectHTMLElement.With(c)
			case Multiple:
				s.Element.With(c)
				s.selectHTMLElement.With(html.Multiple())
			case Disabled:
				s.selectHTMLElement.With(html.Disabled())
			default:
				s.Element.With(c)
			}
		case []any:
			s.With(c...)
		default:
			s.Element.With(c)
		}
	}

	return s
}

func (s *selectEl) Clone() Element {
	return &selectEl{
		Element:           s.Element.Clone(),
		selectHTMLElement: s.selectHTMLElement.Clone(),
	}
}

// Option creates an option element, to be used as a child of a Select or
// SelectMultiple. The value argument is used as the option value attribute.
//
// https://willoma.github.io/bulma-gomponents/form/select.html
func Option(value string, children ...any) Element {
	o := &option{Elem(html.Option, html.Value(value))}
	o.With(children...)
	return o
}

// OptionSelected creates a selected option element, to be used as a child of a
// Select or SelectMultiple. The value argument is used as the option value
// attribute.
//
// https://willoma.github.io/bulma-gomponents/form/select.html
func OptionSelected(value string, children ...any) Element {
	o := &option{Elem(html.Option, html.Value(value), html.Selected())}
	o.With(children...)
	return o
}

type option struct {
	Element
}

func (o *option) Clone() Element {
	return &option{o.Element.Clone()}
}

func Size(size int) gomponents.Node {
	return gomponents.Attr("size", strconv.Itoa(size))
}
