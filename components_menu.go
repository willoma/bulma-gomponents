package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Menu creates a menu.
func Menu(children ...any) *Element {
	return Elem(html.Aside).
		With(Class("menu")).
		Withs(children)
}

// MenuLabel creates a menu label.
func MenuLabel(children ...any) *Element {
	return Elem(html.P).
		With(Class("menu-label")).
		Withs(children)
}

// MenuList creates a menu list.
func MenuList(children ...any) *Element {
	return Elem(html.Ul).
		With(Class("menu-list")).
		Withs(children)
}

// MenuEntry creates an entry for a menu list.
func MenuEntry(children ...any) *Element {
	return Elem(html.Li).Withs(children)
}

// MenuSublist creates a menu sublist.
func MenuSublist(children ...any) *Element {
	return Elem(html.Ul).Withs(children)
}
