package bulma

import (
	e "github.com/willoma/gomplements"
)

// Content creates a content element.
//
// https://willoma.github.io/bulma-gomponents/content.html
func Content(children ...any) e.Element {
	c := &content{e.Div(e.Class("content"))}
	c.With(children...)
	return c
}

type content struct {
	e.Element
}

func (ct *content) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case noP:
			ct.withNoP(c)
		case string:
			ct.Element.With(e.P(c))
		case []any:
			ct.With(c...)
		default:
			ct.Element.With(c)
		}
	}

	return ct
}

func (ct *content) withNoP(children []any) e.Element {
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

func (ct *content) Clone() e.Element {
	return &content{ct.Element.Clone()}
}

func NoP(children ...any) noP {
	return noP(children)
}

type noP []any
