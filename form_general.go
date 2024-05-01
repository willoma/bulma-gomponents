package bulma

import (
	e "github.com/willoma/gomplements"
)

// Field creates a field element.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Field(children ...any) e.Element {
	return e.Div(e.Class("field"), children)
}

// Label creates a label element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Label(children ...any) e.Element {
	l := &label{e.Label(e.Class("label"))}
	l.With(children...)
	return l
}

type label struct {
	e.Element
}

func (l *label) Clone() e.Element {
	return &label{l.Element.Clone()}
}

// Control creates a control element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Control(children ...any) e.Element {
	return e.Div(e.Class("control"), children)
}

// Help creates a help element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Help(children ...any) e.Element {
	return e.P(e.Class("help"), children)
}

// FieldHorizontal creates a horizontal field, including an empty body if needed.
//
// https://willoma.github.io/bulma-gomponents/form.html
func FieldHorizontal(children ...any) e.Element {
	label := e.Div(e.Class("field-label"))
	body := e.Div(e.Class("field-body"))
	f := &fieldHorizontal{
		Element: e.Div(e.Class("field"), Horizontal, label, body),
		label:   label,
		body:    body,
	}
	f.With(children...)
	return f
}

type fieldHorizontal struct {
	e.Element
	label e.Element
	body  e.Element
}

func (f *fieldHorizontal) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onField:
			f.Element.With(c...)
		case onLabel:
			f.label.With(c...)
		case onBody:
			f.body.With(c...)
		case *label:
			f.label.With(c)
		case e.Element:
			f.body.With(c)
		case []any:
			f.With(c...)
		default:
			f.Element.With(c)
		}
	}

	return f
}

func (f *fieldHorizontal) Clone() e.Element {
	return &fieldHorizontal{
		Element: f.Element.Clone(),
		label:   f.label.Clone(),
		body:    f.body.Clone(),
	}
}
