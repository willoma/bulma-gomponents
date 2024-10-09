package bulma

import (
	"io"
	"sync"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// TCell creates a table cell.
func TCell(children ...any) e.Element {
	c := &tcell{Element: e.Td()}
	c.With(children...)
	return c
}

type tcell struct {
	e.Element
}

func (c *tcell) Clone() e.Element {
	return &tcell{
		Element: c.Element.Clone(),
	}
}

// Th creates a table cell with header type.
func Th(children ...any) e.Element {
	c := &tcellPredefined{Element: e.Th()}
	c.With(children...)
	return c
}

// Td creates a table cell with data type.
func Td(children ...any) e.Element {
	c := &tcellPredefined{Element: e.Td()}
	c.With(children...)
	return c
}

type tcellPredefined struct {
	e.Element
}

func (c *tcellPredefined) Clone() e.Element {
	return &tcellPredefined{
		Element: c.Element.Clone(),
	}
}

// HeadRow creates a table header row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func HeadRow(children ...any) e.Element {
	return newRow(rowSectionHead, html.Th, children...)
}

// FootRow creates a table footer row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func FootRow(children ...any) e.Element {
	return newRow(rowSectionFoot, html.Th, children...)
}

// Row creates a table body row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func Row(children ...any) e.Element {
	return newRow(rowSectionBody, html.Td, children...)
}

type rowSection int

const (
	rowSectionBody rowSection = iota
	rowSectionHead
	rowSectionFoot
)

func newRow(section rowSection, elemFn func(...gomponents.Node) gomponents.Node, children ...any) *row {
	r := &row{
		Element: e.Tr(),
		section: section,
		elemFn:  elemFn,
	}
	r.With(children...)
	return r
}

type row struct {
	e.Element

	section rowSection
	elemFn  func(...gomponents.Node) gomponents.Node
}

func (r *row) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case *tcell:
			c.With(r.elemFn)
			r.Element.With(c)
		case *tcellPredefined:
			r.Element.With(c)
		case string:
			r.Element.With(TCell(r.elemFn, c))
		case gomponents.Node:
			if e.IsAttribute(c) {
				r.Element.With(c)
			} else {
				r.Element.With(TCell(r.elemFn, c))
			}
		case []any:
			r.With(c...)
		default:
			r.Element.With(c)
		}
	}

	return r
}

func (r *row) Clone() e.Element {
	return &row{
		Element: r.Element.Clone(),
		section: r.section,
		elemFn:  r.elemFn,
	}
}

// Table creates a table element.
//
// https://willoma.github.io/bulma-gomponents/table.html
func Table(children ...any) e.Element {
	t := &table{table: e.Table(e.Class("table"))}
	t.With(children...)
	return t
}

type table struct {
	table e.Element
	head  e.Element
	body  e.Element
	foot  e.Element

	rendered sync.Once
}

func (t *table) addToHead(children ...any) {
	if t.head == nil {
		t.head = e.THead()
	}
	t.head.With(children...)
}

func (t *table) addToBody(children ...any) {
	if t.body == nil {
		t.body = e.TBody()
	}
	t.body.With(children...)
}

func (t *table) addToFoot(children ...any) {
	if t.foot == nil {
		t.foot = e.TFoot()
	}
	t.foot.With(children...)
}

func (t *table) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onHead:
			t.addToHead(c...)
		case onBody:
			t.addToBody(c...)
		case onFoot:
			t.addToFoot(c...)
		case *row:
			switch c.section {
			case rowSectionHead:
				t.addToHead(c)
			case rowSectionBody:
				t.addToBody(c)
			case rowSectionFoot:
				t.addToFoot(c)
			}
		case []any:
			t.With(c...)
		default:
			t.table.With(c)
		}
	}

	return t
}

func (t *table) Render(w io.Writer) error {
	t.rendered.Do(func() {
		if t.head != nil {
			t.table.With(t.head)
		}
		if t.body != nil {
			t.table.With(t.body)
		}
		if t.foot != nil {
			t.table.With(t.foot)
		}
	})

	return t.table.Render(w)
}

func (t *table) Clone() e.Element {
	return &table{
		table: t.table.Clone(),
		head:  t.head.Clone(),
		body:  t.body.Clone(),
		foot:  t.foot.Clone(),
	}
}

// ScrollableTable creates a table in a table-container element, making the
// table scrollable.
//
// https://willoma.github.io/bulma-gomponents/table.html
func ScrollableTable(children ...any) e.Element {
	return e.Div(e.Class("table-container"), Table(children...))
}
