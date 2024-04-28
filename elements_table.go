package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// TCell creates a table cell.
func TCell(children ...any) Element {
	c := &tcell{Element: Elem(html.Td)}
	c.With(children...)
	return c
}

type tcell struct {
	Element

	hasChangedElem bool
}

func (c *tcell) With(children ...any) Element {
	for _, ch := range children {
		switch ch := ch.(type) {
		case func(...gomponents.Node) gomponents.Node:
			if !c.hasChangedElem {
				c.hasChangedElem = true
				c.Element.With(ch)
			}
		case []any:
			c.With(ch...)
		default:
			c.Element.With(ch)
		}
	}

	return c
}

func (c *tcell) Clone() Element {
	return &tcell{
		Element:        c.Element.Clone(),
		hasChangedElem: c.hasChangedElem,
	}
}

// HeadRow creates a table header row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func HeadRow(children ...any) Element {
	return newRow(rowSectionHead, html.Th, children...)
}

// FootRow creates a table footer row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func FootRow(children ...any) Element {
	return newRow(rowSectionFoot, html.Th, children...)
}

// Row creates a table body row (tr element).
//
// https://willoma.github.io/bulma-gomponents/table.html
func Row(children ...any) Element {
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
		Element: Elem(html.Tr),
		section: section,
		elemFn:  elemFn,
	}
	r.With(children...)
	return r
}

type row struct {
	Element

	section rowSection
	elemFn  func(...gomponents.Node) gomponents.Node
}

func (r *row) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *tcell:
			c.With(r.elemFn)
			r.Element.With(c)
		case string:
			r.Element.With(TCell(r.elemFn, gomponents.Text(c)))
		case gomponents.Node:
			if isAttribute(c) {
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

func (r *row) Clone() Element {
	return &row{
		Element: r.Element.Clone(),
		section: r.section,
		elemFn:  r.elemFn,
	}
}

// Table creates a table element.
//
// https://willoma.github.io/bulma-gomponents/table.html
func Table(children ...any) Element {
	t := &table{table: Elem(html.Table, Class("table"))}
	t.With(children...)
	return t
}

type table struct {
	table Element
	head  Element
	body  Element
	foot  Element

	rendered sync.Once
}

func (t *table) addToHead(children ...any) {
	if t.head == nil {
		t.head = Elem(html.THead)
	}
	t.head.With(children...)
}

func (t *table) addToBody(children ...any) {
	if t.body == nil {
		t.body = Elem(html.TBody)
	}
	t.body.With(children...)
}

func (t *table) addToFoot(children ...any) {
	if t.foot == nil {
		t.foot = Elem(html.TFoot)
	}
	t.foot.With(children...)
}

func (t *table) With(children ...any) Element {
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

func (t *table) Clone() Element {
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
func ScrollableTable(children ...any) Element {
	return Elem(html.Div, Class("table-container"), Table(children...))
}
