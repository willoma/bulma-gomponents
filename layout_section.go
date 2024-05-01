package bulma

import (
	e "github.com/willoma/gomplements"
)

// Section creates a section element.
//
// http://willoma.github.io/bulma-gomponents/section.html
func Section(children ...any) e.Element {
	return e.Section(e.Class("section"), children)
}
