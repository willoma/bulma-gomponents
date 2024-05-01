package bulma

import (
	e "github.com/willoma/gomplements"
)

// Delete creates a delete button.
//
// https://willoma.github.io/bulma-gomponents/delete.html
func Delete(children ...any) e.Element {
	return &delete{e.Button(e.Class("delete"), children)}
}

type delete struct {
	e.Element
}

func (d *delete) Clone() e.Element {
	return &delete{d.Element.Clone()}
}
