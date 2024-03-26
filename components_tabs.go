package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Tabs create a tabs container.
//
// Its arguments must include TabsLink. The ul element is automatically created.
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
	intermediateContainer Element
	tabsChildren          []any
	contentChildren       []any
}

func (t *tabs) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class, ColorClass, ExternalClass, ExternalClassesAndStyles, MultiClass, Styles:
			t.tabsChildren = append(t.tabsChildren, c)
		case func(children ...any) container:
			t.intermediateContainer = c()
		case []any:
			t.With(c...)
		default:
			t.contentChildren = append(t.contentChildren, c)
		}
	}

	return t
}

func (t *tabs) Render(w io.Writer) error {
	tabsEl := Elem(html.Div, Class("tabs"), t.tabsChildren)

	content := Elem(html.Ul, t.contentChildren...)

	if t.intermediateContainer != nil {
		return tabsEl.With(t.intermediateContainer.With(content)).Render(w)
	}

	return tabsEl.With(content).Render(w)
}

// TabsLink creates a tab entry which is a link. Use html.Href as an argument
// to define a link target if needed.
func TabsLink(children ...any) Element {
	return new(tabsLink).With(children...)
}

type tabsLink struct {
	active   bool
	children []any
}

func (t *tabsLink) With(children ...any) Element {
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

func (t *tabsLink) Render(w io.Writer) error {
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
