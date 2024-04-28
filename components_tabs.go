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
		tabs: Elem(html.Div, Class("tabs")),
		list: list,
	}
	t.With(children...)
	return t
}

type tabs struct {
	tabs                  Element
	list                  Element
	intermediateContainer *container

	rendered sync.Once
}

func (t *tabs) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onTabs:
			t.tabs.With(c...)
		case onUl:
			t.list.With(c...)
		case Class, Classer, Classeser, Styles:
			t.tabs.With(c)
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
			t.tabs.With(target)
		} else {
			target = t.tabs
		}

		target.With(t.list)
	})

	return t.tabs.Render(w)
}

func (t *tabs) Clone() Element {
	return &tabs{
		tabs:                  t.tabs.Clone(),
		list:                  t.list.Clone(),
		intermediateContainer: t.intermediateContainer,
	}
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

func (t *tabLink) Clone() Element {
	return &tabLink{
		Element: t.Element.Clone(),
		a:       t.a.Clone(),
	}
}

// TabAHref creates a tab link with an a element.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func TabAHref(href string, children ...any) Element {
	return TabLink(html.Href(href), children)
}
