package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Content creates a content element.
//
// https://willoma.github.io/bulma-gomponents/content.html
func Content(children ...any) Element {
	c := &content{Elem(html.Div, Class("content"))}
	c.With(children...)
	return c
}

type content struct {
	Element
}

func (ct *content) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case noP:
			ct.withNoP(c)
		case string:
			ct.Element.With(Elem(html.P, c))
		case []any:
			ct.With(c...)
		default:
			ct.Element.With(c)
		}
	}

	return ct
}

func (ct *content) withNoP(children []any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			ct.With(c...)
		default:
			ct.Element.With(c)
		}
	}

	return ct
}

func NoP(children ...any) noP {
	return noP(children)
}

type noP []any
