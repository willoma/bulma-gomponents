package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Menu creates a menu.
func Menu(children ...any) Element {
	return Elem(html.Aside, Class("menu"), children)
}

// MenuLabel creates a menu label.
func MenuLabel(children ...any) Element {
	return Elem(html.P, Class("menu-label"), children)
}

// MenuList creates a menu list.
func MenuList(children ...any) Element {
	return Elem(html.Ul, Class("menu-list"), children)
}

// MenuEntry creates an entry for a menu list.
func MenuEntry(children ...any) Element {
	return Elem(html.Li, children...)
}

// MenuAHref creates an entry for a menu list, with a link wrapped inside it.
func MenuAHref(href string, children ...any) Element {
	return (&menuAhref{href: href}).With(children...)
}

type menuAhref struct {
	href          string
	entryChildren []any
	aChildren     []any
}

func (m *menuAhref) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *ApplyToInner:
			m.aChildren = append(m.aChildren, c)
		case *ApplyToOuter:
			m.entryChildren = append(m.entryChildren, c)
		case Class:
			if c == Active {
				m.aChildren = append(m.aChildren, c)
			} else {
				m.entryChildren = append(m.entryChildren, c)
			}
		case ColorClass, ExternalClass, ExternalClassesAndStyles, MultiClass, Styles:
			m.entryChildren = append(m.entryChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				m.entryChildren = append(m.entryChildren, c)
			} else {
				m.aChildren = append(m.aChildren, c)
			}
		case []any:
			m.With(c...)
		default:
			m.aChildren = append(m.aChildren, c)
		}
	}
	return m
}

func (m *menuAhref) Render(w io.Writer) error {
	return MenuEntry(AHref(m.href, m.aChildren...), m.entryChildren).Render(w)
}

// MenuSublist creates a menu sublist.
func MenuSublist(children ...any) Element {
	return Elem(html.Ul, children...)
}
