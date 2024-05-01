package bulma

import (
	e "github.com/willoma/gomplements"
)

// Box creates a box element.
//
// https://willoma.github.io/bulma-gomponents/box.html
func Box(children ...any) e.Element {
	return e.Div(e.Class("box"), children)
}
