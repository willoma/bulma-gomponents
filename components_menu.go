package bulma

import (
	"io"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
)

// Menu creates a menu.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func Menu(children ...any) e.Element {
	m := &menu{menu: e.Aside(e.Class("menu"))}
	m.With(children...)
	return m
}

type menu struct {
	menu            e.Element
	currentMenuList e.Element
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

func (m *menu) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onMenu:
			m.menu.With(c...)
		case string:
			m.flushCurrentMenuList()
			m.menu.With(MenuLabel(c))
		case *menuList:
			m.flushCurrentMenuList()
			m.currentMenuList = c
		case *menuLabel:
			m.flushCurrentMenuList()
			m.menu.With(c)
		case *menuEntry:
			m.addToMenuList(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
				m.menu.With(c)
			} else {
				m.addToMenuList(c)
			}
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

func (m *menu) Clone() e.Element {
	return &menu{
		menu:            m.menu.Clone(),
		currentMenuList: m.currentMenuList.Clone(),
	}
}

// MenuLabel creates a menu label.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuLabel(children ...any) e.Element {
	m := &menuLabel{e.P(e.Class("menu-label"))}
	m.With(children...)
	return m
}

type menuLabel struct {
	e.Element
}

func (m *menuLabel) Clone() e.Element {
	return &menuLabel{m.Element.Clone()}
}

// MenuList creates a menu list.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuList(children ...any) e.Element {
	m := &menuList{e.Ul(e.Class("menu-list"))}
	m.With(children...)
	return m
}

type menuList struct {
	e.Element
}

func (m *menuList) Clone() e.Element {
	return &menuList{m.Element.Clone()}
}

// MenuEntry creates an entry for a menu list.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuEntry(children ...any) e.Element {
	m := &menuEntry{li: e.Li()}
	m.With(children...)
	return m
}

type menuEntry struct {
	li      e.Element
	subList e.Element
}

func (m *menuEntry) addToSublist(child any) {
	if m.subList == nil {
		m.subList = MenuSublist()
	}
	m.subList.With(child)
}

func (m *menuEntry) With(children ...any) e.Element {
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

func (m *menuEntry) Clone() e.Element {
	return &menuEntry{
		li:      m.li.Clone(),
		subList: m.subList.Clone(),
	}
}

// MenuAHref creates an entry for a menu list, with a link wrapped inside it.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuAHref(href string, children ...any) e.Element {
	a := e.AHref(href)
	m := &menuAhref{Element: MenuEntry(a), a: a}
	m.With(children...)
	return m
}

type menuAhref struct {
	e.Element
	a e.Element
}

func (m *menuAhref) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onLi:
			m.Element.With(c...)
		case onA:
			m.a.With(c...)
		case e.Class:
			if c == Active {
				m.a.With(c)
			} else {
				m.Element.With(c)
			}
		case e.Classer, e.Classeser, e.Styles:
			m.Element.With(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (m *menuAhref) Clone() e.Element {
	return &menuAhref{
		Element: m.Element.Clone(),
		a:       m.a.Clone(),
	}
}

// MenuSublist creates a menu sublist.
//
// https://willoma.github.io/bulma-gomponents/menu.html
func MenuSublist(children ...any) e.Element {
	m := &menuSublist{e.Ul()}
	m.With(children...)
	return m
}

type menuSublist struct {
	e.Element
}

func (m *menuSublist) Clone() e.Element {
	return &menuSublist{m.Element.Clone()}
}
