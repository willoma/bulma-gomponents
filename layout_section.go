package bulma

import "github.com/maragudk/gomponents/html"

// Section creates a section element.
//
// The following modifiers change the section behaviour:
//   - Medium: medium spacing
//   - Large: large spacing
func Section(children ...any) *Element {
	return Elem(html.Section, Class("section"), children)
}
