package bulma

import "github.com/maragudk/gomponents/html"

// Footer creates a footer element.
func Footer(children ...any) *Element {
	return Elem(html.Footer).
		With(Class("footer")).
		Withs(children)
}
