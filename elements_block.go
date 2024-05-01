package bulma

import (
	e "github.com/willoma/gomplements"
)

// Block creates a block element.
//
// https://willoma.github.io/bulma-gomponents/block.html
func Block(children ...any) e.Element {
	return e.Div(e.Class("block"), children)
}
