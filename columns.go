package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Columns creates a columns container.
//
// https://willoma.github.io/bulma-gomponents/columns.html
func Columns(children ...any) Element {
	c := &columns{Elem(html.Div, Class("columns"))}
	c.With(children...)
	return c
}

type columns struct {
	Element
}

func (cols *columns) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *column:
			cols.Element.With(c)
		case Element:
			cols.Element.With(Column(c))
		case string:
			cols.Element.With(Column(c))
		case gomponents.Node:
			if isAttribute(c) {
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

func (cols *columns) Clone() Element {
	return &columns{cols.Element.Clone()}
}

// Column creates a single column.
//
// https://willoma.github.io/bulma-gomponents/columns.html
func Column(children ...any) Element {
	c := &column{Elem(html.Div, Class("column"))}
	c.With(children...)
	return c
}

type column struct {
	Element
}

func (col *column) Clone() Element {
	return &column{col.Element.Clone()}
}
