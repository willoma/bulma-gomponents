package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Hero creates a hero element.
//
// http://willoma.github.io/bulma-gomponents/hero.html
func Hero(children ...any) Element {
	h := &hero{Element: Elem(html.Section, Class("hero"))}
	h.With(children...)
	return h
}

type hero struct {
	Element
	head Element
	body Element
	foot Element

	rendered sync.Once
}

func (h *hero) addToHead(children ...any) {
	if h.head == nil {
		h.head = Elem(html.Div, Class("hero-head"))
	}
	h.head.With(children...)
}

func (h *hero) addToBody(children ...any) {
	if h.body == nil {
		h.body = Elem(html.Div, Class("hero-body"))
	}
	h.body.With(children...)
}

func (h *hero) addToFoot(children ...any) {
	if h.foot == nil {
		h.foot = Elem(html.Div, Class("hero-foot"))
	}
	h.foot.With(children...)
}

func (h *hero) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onSection:
			h.Element.With(c...)
		case onBody:
			h.addToBody(c...)
		case heroHead:
			h.addToHead(c...)
		case heroFoot:
			h.addToFoot(c...)
		case Element:
			h.addToBody(c)
		case gomponents.Node:
			if isAttribute(c) {
				h.Element.With(c)
			} else {
				h.addToBody(c)
			}
		case []any:
			h.With(c...)
		default:
			h.Element.With(c)
		}
	}

	return h
}

func (h *hero) Render(w io.Writer) error {
	h.rendered.Do(func() {
		if h.head != nil {
			h.Element.With(h.head)
		}

		if h.body != nil {
			h.Element.With(h.body)
		}

		if h.foot != nil {
			h.Element.With(h.foot)
		}
	})

	return h.Element.Render(w)
}

// HeroHead marks children as belonging to the head part of a hero element.
//
// http://willoma.github.io/bulma-gomponents/hero.html
func HeroHead(children ...any) heroHead {
	return heroHead(children)
}

type heroHead []any

// HeroFoot marks children as belonging to the foot part of a hero element.
//
// http://willoma.github.io/bulma-gomponents/hero.html
func HeroFoot(children ...any) heroFoot {
	return heroFoot(children)
}

type heroFoot []any
