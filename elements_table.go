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

func rowChildren(elem func(...gomponents.Node) gomponents.Node, children []any) *Element {
	e := Elem(html.Tr)

	for _, c := range children {
		switch c := c.(type) {
		case cell:
			e.With(Elem(elem).Withs(c))
		case string:
			e.With(elem(gomponents.Text(c)))
		case gomponents.Node:
			if IsAttribute(c) {
				e.With(c)
			} else {
				e.With(elem(c))
			}
		default:
			e.With(c)
		}
	}

	return e
}

type rowSection int

const (
	rowSectionBody rowSection = iota
	rowSectionHead
	rowSectionFoot
)

type row struct {
	section rowSection
	row     *Element
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
func HeadRow(children ...any) row {
	return row{rowSectionHead, rowChildren(html.Th, children)}
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
func FootRow(children ...any) row {
	return row{rowSectionFoot, rowChildren(html.Th, children)}
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
func Row(children ...any) row {
	return row{rowSectionBody, rowChildren(html.Td, children)}
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
	table := Elem(html.Table).
		With(Class("table"))

	var head, foot, body []gomponents.Node

	for _, c := range children {
		switch c := c.(type) {
		case row:
			switch c.section {
			case rowSectionHead:
				head = append(head, c.row)
			case rowSectionFoot:
				foot = append(foot, c.row)
			case rowSectionBody:
				body = append(body, c.row)
			}
		default:
			table.With(c)
		}
	}

	if len(head) > 0 {
		table.With(html.THead(head...))
	}
	if len(foot) > 0 {
		table.With(html.TFoot(foot...))
	}
	if len(body) > 0 {
		table.With(html.TBody(body...))
	}

	return table
}

// ScrollableTable creates a table in a table-container element, making the
// table scrollable.
//
// See the documentation on the Table function for modifiers details.
func ScrollableTable(children ...any) *Element {
	return Elem(html.Div).
		With(Class("table-container")).
		With(Table(children...))
}
