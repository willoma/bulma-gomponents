package bulma

import (
	"io"
	"sync"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
)

// Hero creates a hero element.
//
// http://willoma.github.io/bulma-gomponents/hero.html
func Hero(children ...any) e.Element {
	h := &hero{hero: e.Section(e.Class("hero"))}
	h.With(children...)
	return h
}

type hero struct {
	hero e.Element
	head e.Element
	body e.Element
	foot e.Element

	rendered sync.Once
}

func (h *hero) addToHead(children ...any) {
	if h.head == nil {
		h.head = e.Div(e.Class("hero-head"))
	}
	h.head.With(children...)
}

func (h *hero) addToBody(children ...any) {
	if h.body == nil {
		h.body = e.Div(e.Class("hero-body"))
	}
	h.body.With(children...)
}

func (h *hero) addToFoot(children ...any) {
	if h.foot == nil {
		h.foot = e.Div(e.Class("hero-foot"))
	}
	h.foot.With(children...)
}

func (h *hero) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onSection:
			h.hero.With(c...)
		case onBody:
			h.addToBody(c...)
		case heroHead:
			h.addToHead(c...)
		case heroFoot:
			h.addToFoot(c...)
		case e.Element:
			h.addToBody(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
				h.hero.With(c)
			} else {
				h.addToBody(c)
			}
		case []any:
			h.With(c...)
		default:
			h.hero.With(c)
		}
	}

	return h
}

func (h *hero) Render(w io.Writer) error {
	h.rendered.Do(func() {
		if h.head != nil {
			h.hero.With(h.head)
		}

		if h.body != nil {
			h.hero.With(h.body)
		}

		if h.foot != nil {
			h.hero.With(h.foot)
		}
	})

	return h.hero.Render(w)
}

func (h *hero) Clone() e.Element {
	return &hero{
		hero: h.hero.Clone(),
		head: h.head.Clone(),
		body: h.body.Clone(),
		foot: h.foot.Clone(),
	}
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
