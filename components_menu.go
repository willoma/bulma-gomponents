package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Menu creates a menu.
func Menu(children ...any) *Element {
	return Elem(html.Aside, Class("menu"), children)
}

// MenuLabel creates a menu label.
func MenuLabel(children ...any) *Element {
	return Elem(html.P, Class("menu-label"), children)
}

// MenuList creates a menu list.
func MenuList(children ...any) *Element {
	return Elem(html.Ul, Class("menu-list"), children)
}

// MenuEntry creates an entry for a menu list.
func MenuEntry(children ...any) *Element {
	return Elem(html.Li, children...)
}

// MenuSublist creates a menu sublist.
func MenuSublist(children ...any) *Element {
	return Elem(html.Ul, children...)
}
