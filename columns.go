package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Columns creates a columns container.
//
// Each child is automatically wrapped in a new Column element if:
//   - it is an Element but not generated with Column
//   - it is a gomponents.Node with type gomponents.ElementType
//
// The following modifiers change the gap:
//   - Gapless: no gap
//   - Gap0: no gap ("is-variable" is automatically added)
//   - Gap1: 0.25rem gap ("is-variable" is automatically added)
//   - Gap2: 0.5rem gap ("is-variable" is automatically added)
//   - Gap3: 0.75rem gap ("is-variable" is automatically added)
//   - Gap4: 1rem gap ("is-variable" is automatically added)
//   - Gap5: 1.25rem gap ("is-variable" is automatically added)
//   - Gap6: 1.5rem gap ("is-variable" is automatically added)
//   - Gap7: 1.75rem gap ("is-variable" is automatically added)
//   - Gap8: 2rem gap ("is-variable" is automatically added)
//
// The gap may be breakpoint-based, with Gap*.Mobile() to Gap*.FullHD().
//
// The following modifiers change the columns behaviour:
//
//   - Centered: center columns
//   - Desktop: allow columns only on desktops upward (not on tablets)
//   - Mobile: allow columns on mobile phones too
//   - Multiline: create a new line when columns do not fit in a single line
//   - VCentered: align columns vertically
func Columns(children ...any) Element {
	return (&columns{}).With(children...)
}

type columns struct {
	children []any
}

func (cols *columns) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *column:
			cols.children = append(cols.children, c)
		case Element:
			cols.children = append(cols.children, Column(c))
		case gomponents.Node:
			if IsAttribute(c) {
				cols.children = append(cols.children, c)
			} else {
				cols.children = append(cols.children, Column(c))
			}
		case []any:
			cols.With(c...)
		default:
			cols.children = append(cols.children, c)
		}
	}

	return cols
}

func (cols *columns) Render(w io.Writer) error {
	return Elem(html.Div, Class("columns"), cols.children).Render(w)
}

// Column creates a single column.
//
// The following modifiers change its width:
//   - Full, Size12: 100%
//   - Size11: 91.66%
//   - Size10: 83.33%
//   - FourFifths: 80%
//   - ThreeQuarters, Size9: 75%
//   - TwoThirds, Size8: 66.66%
//   - ThreeFifths: 60%
//   - Size7: 58.33%
//   - Half, Size6: 50%
//   - Size5: 41.66%
//   - TwoFifths: 40%
//   - OneThird, Size4: 33.33%
//   - OneQuarter, Size3: 25%
//   - OneFifth: 20%
//   - Size2: 16.66%
//   - Size1: 8.33%
//
// The following modifiers add an offset before the column:
//   - Offset11: 91.66%
//   - Offset10: 83.33%
//   - OffsetFourFifths: 80%
//   - OffsetThreeQuarters, Offset9: 75%
//   - OffsetTwoThirds, Offset8: 66%
//   - OffsetThreeFifths: 60%
//   - Offset7: 58.33%
//   - OffsetHalf, Offset6: 50%
//   - Offset5: 41.66%
//   - OffsetTwoFifths: 40%
//   - OffsetOneThird, Offset4: 33%
//   - OffsetOneQuarter, Offset3: 25%
//   - OffsetOneFifth: 20%
//   - Offset2: 16.66%
//   - Offset1: 8.33%
//
// The width may be breakpoint-based, by calling .Mobile() to .FullHD() on the size modifier(s).
//
// The following modifiers change the columns behaviour:
//   - Narrow: the column takes only the width it needs
//
// the narrow behaviour may be breakpoint-based, by using Narrow.Mobile() to Narrow.FullHD().
func Column(children ...any) Element {
	return (&column{}).With(children...)
}

type column struct {
	children []any
}

func (col *column) With(children ...any) Element {
	for _, c := range children {
		if slice, ok := c.([]any); ok {
			col.With(slice...)
		} else {
			col.children = append(col.children, c)
		}
	}

	return col
}

func (col *column) Render(w io.Writer) error {
	return Elem(html.Div, Class("column"), col.children).Render(w)
}
