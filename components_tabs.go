package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Tabs create a tabs container.
//   - when a child is marked with b.Inner, it is forcibly applied to the <ul> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <div class="tabs"> element
//
// Its arguments must include TabLink. The ul element is automatically created.
//
// If a child is the Container function or one of its derivative (Container*),
// this function is executed and its result is used as an intermediate container.
//
// The following modifiers change the tabs behaviour:
//   - Centered: center the tabs
//   - Right: align the tabs to the right
//   - Boxed: draw boxed tabs
//   - Toggle: button-looking tabs
//   - ToggleRounded: rounded button-looking tabs
//   - FullWidth: take the whole width
//
// The following modifiers change the tabs size:
//   - Small
//   - Medium
//   - Large
func Tabs(children ...any) Element {
	return new(tabs).With(children...)
}

type tabs struct {
	intermediateContainer *container
	tabsChildren          []any
	ulChildren            []any
}

func (t *tabs) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onTabs:
			t.tabsChildren = append(t.tabsChildren, c...)
		case onUl:
			t.ulChildren = append(t.ulChildren, c...)
		case Class, Classer, Classeser, ExternalClassesAndStyles, MultiClass, Styles:
			t.tabsChildren = append(t.tabsChildren, c)
		case *container:
			t.intermediateContainer = c
		case []any:
			t.With(c...)
		default:
			t.ulChildren = append(t.ulChildren, c)
		}
	}

	return t
}

func (t *tabs) Render(w io.Writer) error {
	tabsEl := Elem(html.Div, Class("tabs"), t.tabsChildren)

	content := Elem(html.Ul, t.ulChildren...)

	if t.intermediateContainer != nil {
		return tabsEl.With(t.intermediateContainer.With(content)).Render(w)
	}

	return tabsEl.With(content).Render(w)
}

// TabLink creates a tab entry which is a link. Use html.Href as an argument
// to define a link target if needed.
func TabLink(children ...any) Element {
	return new(tabLink).With(children...)
}

type tabLink struct {
	active   bool
	children []any
}

func (t *tabLink) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			if c == Active {
				t.active = true
			} else {
				t.children = append(t.children, c)
			}
		case []any:
			t.With(c...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}

func (t *tabLink) Render(w io.Writer) error {
	li := Elem(html.Li)
	if t.active {
		li.With(Active)
	}

	return li.With(
		Elem(
			html.A,
			elemOptionSpanAroundNonIconsIfHasIcons,
			t.children,
		),
	).Render(w)
}

func TabAHref(href string, children ...any) Element {
	return new(tabLink).With(html.Href(href), children)
}
