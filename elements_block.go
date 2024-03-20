package bulma

import "github.com/maragudk/gomponents/html"

// Block creates a block element.
func Block(children ...any) *Element {
	return Elem(html.Div, Class("block"), children)
}
