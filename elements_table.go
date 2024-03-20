package bulma

import (
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

func (r *row) addChildren(children []any) {
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
			r.addChildren(c)
		default:
			r.children = append(r.children, c)
		}
	}
}

func (r *row) elem() *Element {
	return Elem(html.Tr, r.children...)
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
func HeadRow(children ...any) *row {
	r := &row{section: rowSectionHead, el: html.Th}
	r.addChildren(children)
	return r
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
func FootRow(children ...any) *row {
	r := &row{section: rowSectionFoot, el: html.Th}
	r.addChildren(children)
	return r
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
func Row(children ...any) *row {
	r := &row{section: rowSectionBody, el: html.Td}
	r.addChildren(children)
	return r
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
//   - Table: add a hover effect on each body row
//   - FullWidth: take the whole width
func Table(children ...any) *Element {
	t := &table{}
	t.addChildren(children)
	return t.elem()
}

type table struct {
	children []any
	head     []gomponents.Node
	foot     []gomponents.Node
	body     []gomponents.Node
}

func (t *table) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *row:
			switch c.section {
			case rowSectionHead:
				t.head = append(t.head, c.elem())
			case rowSectionFoot:
				t.foot = append(t.foot, c.elem())
			case rowSectionBody:
				t.body = append(t.body, c.elem())
			}
		case []any:
			t.addChildren(c)
		default:
			t.children = append(t.children, c)
		}
	}
}

func (t *table) elem() *Element {
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

	return tb
}

// ScrollableTable creates a table in a table-container element, making the
// table scrollable.
//
// See the documentation on the Table function for modifiers details.
func ScrollableTable(children ...any) *Element {
	return Elem(html.Div, Class("table-container"), Table(children...))
}
