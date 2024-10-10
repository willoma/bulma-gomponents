package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Tabs create a tabs container.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func Tabs(children ...any) e.Element {
	list := e.Ul()
	t := &tabs{
		tabs: e.Div(e.Class("tabs")),
		list: list,
	}
	t.With(children...)
	return t
}

type tabs struct {
	tabs                  e.Element
	list                  e.Element
	intermediateContainer *container
}

func (t *tabs) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onTabs:
			t.tabs.With(c...)
		case onUl:
			t.list.With(c...)
		case e.Class, e.Classer, e.Classeser, e.Styles:
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
	tabs := t.tabs.Clone()

	if t.intermediateContainer != nil {
		tabs.With(t.intermediateContainer.Clone().With(t.list))
	} else {
		tabs.With(t.list)
	}

	return tabs.Render(w)
}

func (t *tabs) Clone() e.Element {
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
func TabLink(children ...any) e.Element {
	a := &spanAroundNonIconsIfHasIcons{elemFn: html.A}
	t := &tabLink{
		Element: e.Li(a),
		a:       a,
	}
	t.With(children...)
	return t
}

type tabLink struct {
	e.Element
	a e.Element
}

func (t *tabLink) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class:
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

func (t *tabLink) Clone() e.Element {
	return &tabLink{
		Element: t.Element.Clone(),
		a:       t.a.Clone(),
	}
}

// TabAHref creates a tab link with an a element.
//
// https://willoma.github.io/bulma-gomponents/tabs.html
func TabAHref(href string, children ...any) e.Element {
	return TabLink(html.Href(href), children)
}
