package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Delete creates a delete button.
//
// https://willoma.github.io/bulma-gomponents/delete.html
func Delete(children ...any) Element {
	return &delete{Elem(html.Button, Class("delete"), children)}
}

type delete struct {
	Element
}

func (d *delete) Clone() Element {
	return &delete{d.Element.Clone()}
}
