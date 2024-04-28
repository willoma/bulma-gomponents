package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Menu creates a menu.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func Menu(children ...any) Element {
	m := &menu{menu: Elem(html.Aside, Class("menu"))}
	m.With(children...)
	return m
}

type menu struct {
	menu            Element
	currentMenuList Element
}

func (m *menu) flushCurrentMenuList() {
	if m.currentMenuList != nil {
		m.menu.With(m.currentMenuList)
		m.currentMenuList = nil
	}
}

func (m *menu) addToMenuList(child any) {
	if m.currentMenuList == nil {
		m.currentMenuList = MenuList()
	}
	m.currentMenuList.With(child)
}

func (m *menu) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onMenu:
			m.menu.With(c...)
		case string:
			m.flushCurrentMenuList()
			m.menu.With(MenuLabel(c))
		case *menuLabel:
			m.flushCurrentMenuList()
			m.menu.With(c)
		case *menuEntry:
			m.addToMenuList(c)
		case gomponents.Node:
			if isAttribute(c) {
				m.menu.With(c)
			} else {
				m.addToMenuList(c)
			}
		case Element:
			m.addToMenuList(c)
		case []any:
			m.With(c...)
		default:
			m.menu.With(c)
		}
	}

	return m
}

func (m *menu) Render(w io.Writer) error {
	m.flushCurrentMenuList()
	return m.menu.Render(w)
}

func (m *menu) Clone() Element {
	return &menu{
		menu:            m.menu.Clone(),
		currentMenuList: m.currentMenuList.Clone(),
	}
}

// MenuLabel creates a menu label.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuLabel(children ...any) Element {
	m := &menuLabel{Elem(html.P, Class("menu-label"))}
	m.With(children...)
	return m
}

type menuLabel struct {
	Element
}

func (m *menuLabel) Clone() Element {
	return &menuLabel{m.Element.Clone()}
}

// MenuList creates a menu list.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuList(children ...any) Element {
	m := &menuList{Elem(html.Ul, Class("menu-list"))}
	m.With(children...)
	return m
}

type menuList struct {
	Element
}

func (m *menuList) Clone() Element {
	return &menuList{m.Element.Clone()}
}

// MenuEntry creates an entry for a menu list.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuEntry(children ...any) Element {
	m := &menuEntry{li: Elem(html.Li, Class("menu-list"))}
	m.With(children...)
	return m
}

type menuEntry struct {
	li      Element
	subList Element
}

func (m *menuEntry) addToSublist(child any) {
	if m.subList == nil {
		m.subList = MenuSublist()
	}
	m.subList.With(child)
}

func (m *menuEntry) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onLi:
			m.li.With(c)
		case *menuEntry:
			m.addToSublist(c)
		case *menuSublist:
			m.subList = c
		case []any:
			m.li.With(c...)
		default:
			m.li.With(c)
		}
	}
	return m
}

func (m *menuEntry) Render(w io.Writer) error {
	if m.subList != nil {
		m.li.With(m.subList)
		m.subList = nil
	}
	return m.li.Render(w)
}

func (m *menuEntry) Clone() Element {
	return &menuEntry{
		li:      m.li.Clone(),
		subList: m.subList.Clone(),
	}
}

// MenuAHref creates an entry for a menu list, with a link wrapped inside it.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuAHref(href string, children ...any) Element {
	a := AHref(href)
	m := &menuAhref{Element: MenuEntry(a), a: a}
	m.With(children...)
	return m
}

type menuAhref struct {
	Element
	a Element
}

func (m *menuAhref) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onLi:
			m.Element.With(c...)
		case onA:
			m.a.With(c...)
		case Class:
			if c == Active {
				m.a.With(c)
			} else {
				m.Element.With(c)
			}
		case Classer, Classeser, Styles:
			m.Element.With(c)
		case gomponents.Node:
			if isAttribute(c) {
				m.Element.With(c)
			} else {
				m.a.With(c)
			}
		case []any:
			m.With(c...)
		default:
			m.a.With(c)
		}
	}

	return m
}

func (m *menuAhref) Clone() Element {
	return &menuAhref{
		Element: m.Element.Clone(),
		a:       m.a.Clone(),
	}
}

// MenuSublist creates a menu sublist.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuSublist(children ...any) Element {
	m := &menuSublist{Elem(html.Ul)}
	m.With(children...)
	return m
}

type menuSublist struct {
	Element
}

func (m *menuSublist) Clone() Element {
	return &menuSublist{m.Element.Clone()}
}
