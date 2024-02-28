package bulma

import "github.com/maragudk/gomponents/html"

// Box creates a box element.
func Box(children ...any) *Element {
	return Elem(html.Div).
		With(Class("box")).
		Withs(children)
}
