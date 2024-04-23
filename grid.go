package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Grid creates a smart grid.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func Grid(children ...any) Element {
	g := &grid{Elem(html.Div, Class("grid"))}
	g.With(children...)
	return g
}

type grid struct {
	Element
}

func (g *grid) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *cell:
			g.Element.With(c)
		case Element:
			g.Element.With(Cell(c))
		case string:
			g.Element.With(Cell(c))
		case gomponents.Node:
			if isAttribute(c) {
				g.Element.With(c)
			} else {
				g.Element.With(Cell(c))
			}
		case []any:
			g.Element.With(c...)
		default:
			g.Element.With(c)
		}
	}

	return g
}

// FixedGrid creates a fixed grid.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func FixedGrid(children ...any) Element {
	grid := Grid()
	f := &fixedGrid{
		Element: Elem(html.Div, Class("fixed-grid"), grid),
		grid:    grid,
	}
	f.With(children...)
	return f
}

type fixedGrid struct {
	Element
	grid Element
}

func (f *fixedGrid) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onFixedGrid:
			f.Element.With(c...)
		case onGrid:
			f.grid.With(c...)
		case Class:
			switch c {
			case AutoCount,
				Cols1, Cols2, Cols3, Cols4, Cols5, Cols6,
				Cols7, Cols8, Cols9, Cols10, Cols11, Cols12:
				f.Element.With(c)
			default:
				f.grid.With(c)
			}
		case []any:
			f.With(c...)
		default:
			f.grid.With(c)
		}
	}

	return f
}

// Cell creates a single grid cell.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func Cell(children ...any) Element {
	c := &cell{Elem(html.Div, Class("cell"))}
	c.With(children...)
	return c
}

type cell struct {
	Element
}
