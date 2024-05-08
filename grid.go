package bulma

import (
	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

// Grid creates a smart grid.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func Grid(children ...any) e.Element {
	g := &grid{e.Div(e.Class("grid"))}
	g.With(children...)
	return g
}

type grid struct {
	e.Element
}

func (g *grid) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case *cell:
			g.Element.With(c)
		case e.Element:
			g.Element.With(Cell(c))
		case string:
			g.Element.With(Cell(c))
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (g *grid) Clone() e.Element {
	return &grid{g.Element.Clone()}
}

// FixedGrid creates a fixed grid.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func FixedGrid(children ...any) e.Element {
	grid := Grid()
	f := &fixedGrid{
		Element: e.Div(e.Class("fixed-grid"), grid),
		grid:    grid,
	}
	f.With(children...)
	return f
}

type fixedGrid struct {
	e.Element
	grid e.Element
}

func (f *fixedGrid) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onFixedGrid:
			f.Element.With(c...)
		case onGrid:
			f.grid.With(c...)
		case cols:
			f.Element.With(c)
		case e.Class:
			switch c {
			case AutoCount:
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

func (f *fixedGrid) Clone() e.Element {
	return &fixedGrid{
		Element: f.Element.Clone(),
		grid:    f.grid.Clone(),
	}
}

// Cell creates a single grid cell.
//
// https://willoma.github.io/bulma-gomponents/grid.html
func Cell(children ...any) e.Element {
	c := &cell{e.Div(e.Class("cell"))}
	c.With(children...)
	return c
}

type cell struct {
	e.Element
}

func (c *cell) Clone() e.Element {
	return &cell{c.Element.Clone()}
}
