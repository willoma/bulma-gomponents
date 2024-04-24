package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents/html"
)

// Tabs create a tabs container.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func Tabs(children ...any) Element {
	list := Elem(html.Ul)
	t := &tabs{
		Element: Elem(html.Div, Class("tabs")),
		list:    list,
	}
	t.With(children...)
	return t
}

type tabs struct {
	Element
	list                  Element
	intermediateContainer *container

	rendered sync.Once
}

func (t *tabs) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onTabs:
			t.Element.With(c...)
		case onUl:
			t.list.With(c...)
		case Class, Classer, Classeser, Styles:
			t.Element.With(c)
		case *container:
			t.intermediateContainer = c
		case []any:
			t.With(c...)
		default:
			t.list.With(c)
		}
	}

	return t
}

func (t *tabs) Render(w io.Writer) error {
	t.rendered.Do(func() {
		var target Element
		if t.intermediateContainer != nil {
			target = t.intermediateContainer
			t.Element.With(target)
		} else {
			target = t.Element
		}

		target.With(t.list)
	})

	return t.Element.Render(w)
}

// TabLink creates a tab entry which is a link. Use html.Href as an argument
// to define a link target if needed.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func TabLink(children ...any) Element {
	a := Elem(html.A, elemOptionSpanAroundNonIconsIfHasIcons)
	t := &tabLink{
		Element: Elem(html.Li, a),
		a:       a,
	}
	t.With(children...)
	return t
}

type tabLink struct {
	Element
	a Element
	// active   bool
	// children []any
}

func (t *tabLink) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			if c == Active {
				t.Element.With(c)
			} else {
				t.a.With(c)
			}
		case []any:
			t.With(c...)
		default:
			t.a.With(c)
		}
	}

	return t
}

// TabAHref creates a tab link with an a element.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func TabAHref(href string, children ...any) Element {
	return new(tabLink).With(html.Href(href), children)
}
