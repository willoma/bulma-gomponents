package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Grid creates a smart grid.
//
// Each child is automatically wrapped in a new Cell element if:
//   - it is an Element but not generated with Cell
//   - it is a string
//   - it is a gomponents.Node with type gomponents.ElementType
//
// The following modifiers change the grid behaviour:
//   - ColMin1: set minimum column width to 1.5rem
//   - ColMin2: set minimum column width to 3rem
//   - ColMin3: set minimum column width to 4.5rem
//   - ColMin4: set minimum column width to 6rem
//   - ColMin5: set minimum column width to 7.5rem
//   - ColMin6: set minimum column width to 9rem
//   - ColMin7: set minimum column width to 10.5rem
//   - ColMin8: set minimum column width to 12rem
//   - ColMin9: set minimum column width to 13.5rem
//   - ColMin10: set minimum column width to 15rem
//   - ColMin11: set minimum column width to 16.5rem
//   - ColMin12: set minimum column width to 18rem
//   - Gap0: no gap
//   - Gap05: 0.25rem gap
//   - Gap1: 0.5rem gap
//   - Gap15: 0.75rem gap
//   - Gap2: 1rem gap
//   - Gap25: 1.25rem gap
//   - Gap3: 1.5rem gap
//   - Gap35: 1.75rem gap
//   - Gap4: 2rem gap
//   - Gap45: 2.25rem gap
//   - Gap5: 2.5rem gap
//   - Gap55: 2.75rem gap
//   - Gap6: 3rem gap
//   - Gap65: 3.25rem gap
//   - Gap7: 3.5rem gap
//   - Gap75: 3.75rem gap
//   - Gap8: 4rem gap
//   - ColGap0: no column gap
//   - ColGap05: 0.25rem column gap
//   - ColGap1: 0.5rem column gap
//   - ColGap15: 0.75rem column gap
//   - ColGap2: 1rem column gap
//   - ColGap25: 1.25rem column gap
//   - ColGap3: 1.5rem column gap
//   - ColGap35: 1.75rem column gap
//   - ColGap4: 2rem column gap
//   - ColGap45: 2.25rem column gap
//   - ColGap5: 2.5rem column gap
//   - ColGap55: 2.75rem column gap
//   - ColGap6: 3rem column gap
//   - ColGap65: 3.25rem column gap
//   - ColGap7: 3.5rem column gap
//   - ColGap75: 3.75rem column gap
//   - ColGap8: 4rem column gap
//   - RowGap0: no row gap
//   - RowGap05: 0.25rem row gap
//   - RowGap1: 0.5rem row gap
//   - RowGap15: 0.75rem row gap
//   - RowGap2: 1rem row gap
//   - RowGap25: 1.25rem row gap
//   - RowGap3: 1.5rem row gap
//   - RowGap35: 1.75rem row gap
//   - RowGap4: 2rem row gap
//   - RowGap45: 2.25rem row gap
//   - RowGap5: 2.5rem row gap
//   - RowGap55: 2.75rem row gap
//   - RowGap6: 3rem row gap
//   - RowGap65: 3.25rem row gap
//   - RowGap7: 3.5rem row gap
//   - RowGap75: 3.75rem row gap
//   - RowGap8: 4rem row gap
func Grid(children ...any) Element {
	return new(grid).With(children...)
}

type grid struct {
	children []any
}

func (g *grid) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *cell:
			g.children = append(g.children, c)
		case Element:
			g.children = append(g.children, Cell(c))
		case string:
			g.children = append(g.children, Cell(c))
		case gomponents.Node:
			if IsAttribute(c) {
				g.children = append(g.children, c)
			} else {
				g.children = append(g.children, Cell(c))
			}
		case []any:
			g.With(c...)
		default:
			g.children = append(g.children, c)
		}
	}

	return g
}

func (g *grid) Render(w io.Writer) error {
	return Elem(html.Div, Class("grid"), g.children).Render(w)
}

// FixedGrid creates a fixed grid.
//
// It accepts the same children types as Grid, as well as the following modifiers:
//   - Cols1: Set columns count to 1
//   - Cols2: Set columns count to 2
//   - Cols3: Set columns count to 3
//   - Cols4: Set columns count to 4
//   - Cols5: Set columns count to 5
//   - Cols6: Set columns count to 6
//   - Cols7: Set columns count to 7
//   - Cols8: Set columns count to 8
//   - Cols9: Set columns count to 9
//   - Cols10: Set columns count to 10
//   - Cols11: Set columns count to 11
//   - Cols12: Set columns count to 12
//   - AutoCount: Set the following automatic columns count: Mobile: 2 columns,
//     Tablet: 4 columns, Desktop: 8 columns, Widescreen: 12 columns,
//     Full HD: 16 columns
//
// The columns count may be breakpoint-based, with Cols*.Mobile() to Cols*.FullHD().
func FixedGrid(children ...any) Element {
	return (&fixedGrid{grid: new(grid)}).With(children...)
}

type fixedGrid struct {
	children []any
	grid     *grid
}

func (f *fixedGrid) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onFixedGrid:
			f.children = append(f.children, c...)
		case onGrid:
			f.grid.With(c...)
		case Class:
			switch c {
			case AutoCount,
				Cols1, Cols2, Cols3, Cols4, Cols5, Cols6,
				Cols7, Cols8, Cols9, Cols10, Cols11, Cols12:
				f.children = append(f.children, c)
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

func (f *fixedGrid) Render(w io.Writer) error {
	return Elem(html.Div, Class("fixed-grid"), f.children, f.grid).Render(w)
}

// Cell creates a single grid cell.
//
// The following modifiers change its behaviour:
//   - ColStart1 to ColStart12: start the cell at the given column
//   - ColStartEnd: start the cell at the last column
//   - ColFromEnd1 to ColFromEnd12: end the cell at the given column
//   - ColSpan1 to ColSpan12: span the cell over this number of columns
//   - RowStart1 to RowStart12: start the cell at the given row
//   - RowStartEnd: start the cell at the last row
//   - RowFromEnd1 to RowFromEnd12: end the cell at the given row
//   - RowSpan1 to RowSpan12: span the cell over this number of rows
//
// These modifiers may be breakpoint-based, with .Mobile() to .FullHD().
func Cell(children ...any) Element {
	return new(cell).With(children...)
}

type cell struct {
	children []any
}

func (cl *cell) With(children ...any) Element {
	for _, c := range children {
		if slice, ok := c.([]any); ok {
			cl.With(slice...)
		} else {
			cl.children = append(cl.children, c)
		}
	}

	return cl
}

func (cl *cell) Render(w io.Writer) error {
	return Elem(html.Div, Class("cell"), cl.children).Render(w)
}
