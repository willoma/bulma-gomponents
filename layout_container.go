package bulma

import (
	e "github.com/willoma/gomplements"
)

// Container creates a container element. Its width depends on the applied
// modifiers among Widescreen, FullHD, MaxDesktop, MaxWidescreen or Fluid.
//
// http://willoma.github.io/bulma-gomponents/container.html
func Container(children ...any) *container {
	c := &container{e.Div(e.Class("container"))}
	c.With(children...)
	return c
}

type container struct {
	e.Element
}

func (c *container) Clone() e.Element {
	return &container{c.Element.Clone()}
}
