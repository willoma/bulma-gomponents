package bulma

import (
	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

// Columns creates a columns container.
//
// https://willoma.github.io/bulma-gomponents/columns.html
func Columns(children ...any) e.Element {
	c := &columns{e.Div(e.Class("columns"))}
	c.With(children...)
	return c
}

type columns struct {
	e.Element
}

func (cols *columns) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case *column:
			cols.Element.With(c)
		case e.Element:
			cols.Element.With(Column(c))
		case string:
			cols.Element.With(Column(c))
		case gomponents.Node:
			if e.IsAttribute(c) {
				cols.Element.With(c)
			} else {
				cols.Element.With(Column(c))
			}
		case []any:
			cols.With(c...)
		default:
			cols.Element.With(c)
		}
	}

	return cols
}

func (cols *columns) Clone() e.Element {
	return &columns{cols.Element.Clone()}
}

// Column creates a single column.
//
// https://willoma.github.io/bulma-gomponents/columns.html
func Column(children ...any) e.Element {
	c := &column{e.Div(e.Class("column"))}
	c.With(children...)
	return c
}

type column struct {
	e.Element
}

func (col *column) Clone() e.Element {
	return &column{col.Element.Clone()}
}
