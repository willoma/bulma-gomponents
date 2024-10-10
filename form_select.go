package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Select creates a dropdown select element. It should contain one or multiple
// Option elements.
//
// https://willoma.github.io/bulma-gomponents/form/select.html
func Select(children ...any) e.Element {
	selectHTMLElement := e.Select()
	s := &selectEl{
		Element:           e.Div(e.Class("select"), selectHTMLElement),
		selectHTMLElement: selectHTMLElement,
	}
	s.With(children...)
	return s
}

type selectEl struct {
	e.Element
	selectHTMLElement e.Element
}

func (s *selectEl) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onSelect:
			s.selectHTMLElement.With(c...)
		case onDiv:
			s.Element.With(c...)
		case *option:
			s.selectHTMLElement.With(c)
		case e.Class:
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

func (s *selectEl) Clone() e.Element {
	return &selectEl{
		Element:           s.Element.Clone(),
		selectHTMLElement: s.selectHTMLElement.Clone(),
	}
}

// Option creates an option element, to be used as a child of a Select or
// SelectMultiple. The value argument is used as the option value attribute.
//
// https://willoma.github.io/bulma-gomponents/form/select.html
func Option(value string, children ...any) e.Element {
	o := &option{e.Option(html.Value(value))}
	o.With(children...)
	return o
}

type option struct {
	e.Element
}

func (o *option) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class:
			if c == Selected {
				o.Element.With(html.Selected())
			} else {
				o.Element.With(c)
			}
		case []any:
			o.With(c...)
		default:
			o.Element.With(c)
		}
	}

	return o
}

func (o *option) Clone() e.Element {
	return &option{o.Element.Clone()}
}

func SelectSize(size int) gomponents.Node {
	return gomponents.Attr("size", strconv.Itoa(size))
}
