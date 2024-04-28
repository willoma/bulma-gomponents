package bulma

import "github.com/maragudk/gomponents/html"

// Section creates a section element.
//
// http://willoma.github.io/bulma-gomponents/section.html
func Section(children ...any) Element {
	return Elem(html.Section, Class("section"), children)
}
