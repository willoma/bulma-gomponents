package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Container creates a container element. Its width depends on the applied
// modifiers among Widescreen, FullHD, MaxDesktop, MaxWidescreen or Fluid.
//
// http://willoma.github.io/bulma-gomponents/container.html
func Container(children ...any) *container {
	c := &container{Elem(html.Div, Class("container"))}
	c.With(children...)
	return c
}

type container struct {
	Element
}
