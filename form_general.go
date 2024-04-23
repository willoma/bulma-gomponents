package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Field creates a field element.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Field(children ...any) Element {
	return Elem(html.Div, Class("field"), children)
}

// Label creates a label element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Label(children ...any) Element {
	l := &label{Elem(html.Label, Class("label"))}
	l.With(children...)
	return l
}

type label struct {
	Element
}

// Control creates a control element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Control(children ...any) Element {
	return Elem(html.Div, Class("control"), children)
}

// Help creates a help element, to be used as a child of a Field.
//
// https://willoma.github.io/bulma-gomponents/form.html
func Help(children ...any) Element {
	return Elem(html.P, Class("help"), children)
}

// FieldHorizontal creates a horizontal field, including an empty body if needed.
//
// https://willoma.github.io/bulma-gomponents/form.html
func FieldHorizontal(children ...any) Element {
	label := Elem(html.Div, Class("field-label"))
	body := Elem(html.Div, Class("field-body"))
	f := &fieldHorizontal{
		Element: Elem(html.Div, Class("field"), Horizontal, label, body),
		label:   label,
		body:    body,
	}
	f.With(children...)
	return f
}

type fieldHorizontal struct {
	Element
	label Element
	body  Element
}

func (f *fieldHorizontal) With(children ...any) Element {
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
		case Element:
			f.body.With(c)
		case []any:
			f.With(c...)
		default:
			f.Element.With(c)
		}
	}

	return f
}
