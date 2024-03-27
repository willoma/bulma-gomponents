package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type cell []any

// Cell creates a table cell.
func Cell(children ...any) cell {
	return cell(children)
}

type rowSection int

const (
	rowSectionBody rowSection = iota
	rowSectionHead
	rowSectionFoot
)

type row struct {
	section  rowSection
	el       func(...gomponents.Node) gomponents.Node
	children []any
}

func (r *row) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case cell:
			r.children = append(r.children, Elem(r.el, c...))
		case string:
			r.children = append(r.children, r.el(gomponents.Text(c)))
		case gomponents.Node:
			if IsAttribute(c) {
				r.children = append(r.children, c)
			} else {
				r.children = append(r.children, r.el(c))
			}
		case []any:
			r.With(c...)
		default:
			r.children = append(r.children, c)
		}
	}

	return r
}

func (r *row) Render(w io.Writer) error {
	return Elem(html.Tr, r.children...).Render(w)
}

// HeadRow creates a table header row (tr element).
//   - when a child is created by Cell, it is added as a th element
//   - when a child is a string, it is wrapped in a th element
//   - when a child is a gomponents.Node with type gomponents.ElementType, it
//     is wrapped in a th element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the tr element
//   - all other types are added as-is to the tr element
//
// The following modifiers change the row behaviour:
//   - Selected: mark the row as selected
func HeadRow(children ...any) Element {
	return (&row{section: rowSectionHead, el: html.Th}).With(children...)
}

// FootRow creates a table footer row (tr element).
//   - when a child is created by Cell, it is added as a th element
//   - when a child is a string, it is wrapped in a th element
//   - when a child is a gomponents.Node with type gomponents.ElementType, it
//     is wrapped in a th element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as-is to the tr element
//   - all other types are added as-is to the tr element
//
// The following modifiers change the row behaviour:
//   - Selected: mark the row as selected
func FootRow(children ...any) Element {
	return (&row{section: rowSectionFoot, el: html.Th}).With(children...)
}

// Row creates a table body row (tr element).
//   - when a child is created by Cell, it is added as a td element
//   - when a child is a string, it is wrapped in a td element
//   - when a child is a gomponents.Node with type gomponents.ElementType, it
//     is wrapped in a td element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as-is to the tr element
//   - all other types are added as-is to the tr element
//
// The following modifiers change the row behaviour:
//   - Selected: mark the row as selected
func Row(children ...any) Element {
	return (&row{section: rowSectionBody, el: html.Td}).With(children...)
}

// Table creates a table element.
//   - use HeadRow to add a row in the table header
//   - use FootRow to add a row in the table footer
//   - use Row to add a row in the table body
//
// The following modifiers change the table behaviour:
//   - Bordered: add borders
//   - Striped: add stripes
//   - Narrow: make the cells narrower
//   - Hoverable: add a hover effect on each body row
//   - FullWidth: take the whole width
func Table(children ...any) Element {
	return new(table).With(children...)
}

type table struct {
	children []any
	head     []gomponents.Node
	foot     []gomponents.Node
	body     []gomponents.Node
}

func (t *table) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *row:
			switch c.section {
			case rowSectionHead:
				t.head = append(t.head, c)
			case rowSectionFoot:
				t.foot = append(t.foot, c)
			case rowSectionBody:
				t.body = append(t.body, c)
			}
		case []any:
			t.With(c...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}

func (t *table) Render(w io.Writer) error {
	tb := Elem(html.Table, Class("table"), t.children)

	if len(t.head) > 0 {
		tb.With(html.THead(t.head...))
	}

	if len(t.foot) > 0 {
		tb.With(html.TFoot(t.foot...))
	}

	if len(t.body) > 0 {
		tb.With(html.TBody(t.body...))
	}

	return tb.Render(w)
}

// ScrollableTable creates a table in a table-container element, making the
// table scrollable.
//
// See the documentation on the Table function for modifiers details.
func ScrollableTable(children ...any) Element {
	return Elem(html.Div, Class("table-container"), Table(children...))
}
