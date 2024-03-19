package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Field creates a field element.
//
// The following modifiers change the field behaviour:
//   - Addons: attach the children controls (Input*, Button* or Select) together
//   - AddonsCentered: attach the children controls (Input*, Button* or Select)
//     together and center them (Addons is not required)
//   - AddonsCentered: attach the children controls (Input*, Button* or Select)
//     together and align them to the right (Addons is not required)
//   - Grouped: group the children controls together, without attaching them
//   - GroupedCentered: group the children controls together, without attaching
//     them and center them (Grouped is not required)
//   - GroupedRight: group the children controls together, without attaching
//     them and align them to the right (Grouped is not required)
//   - GroupedMultiline: group the children controls together, without attaching
//     them and allow them to fill up multiple lines (Grouped is not required)
//   - Horizontal: make the field horizontal ()
func Field(children ...any) *Element {
	return Elem(html.Div).
		With(Class("field")).
		Withs(children)
}

// Label creates a label element, to be used as a child of a Field.
func Label(children ...any) *Element {
	return Elem(html.Label).
		With(Class("label")).
		Withs(children)
}

// Control creates a control element, to be used as a child of a Field.
//
// The following modifiers change the control behaviour:
//   - IconsLeft: leave place on the left side of the Input* or Select child for
//     an icon - the Icon element must be a direct child of Control and have the
//     Left modifier
//   - IconsRight: leave place on the right side of the Input* or Select child
//     for an icon - the Icon element must be a direct child of Control and have
//     the Right modifier
//   - Expanded: make the contained element take the remaining space - to apply
//     this style to a Select, you also need to add the FullWidth modifier to
//     the Select element
func Control(children ...any) *Element {
	return Elem(html.Div).
		With(Class("control")).
		Withs(children)
}

// Help creates a help element, to be used as a child of a Field.
func Help(children ...any) *Element {
	return Elem(html.P).
		With(Class("help")).
		Withs(children)
}

// FieldBody creates a field-body element, to be used as a child of a
// FieldHorizontal.
func FieldBody(children ...any) *Element {
	return Elem(html.Div).
		With(Class("field-body")).
		Withs(children)
}

// FieldLabel creates a field-label element, to be used as a child of a
// FieldHorizontal.
func FieldLabel(children ...any) *Element {
	return Elem(html.Div).
		With(Class("field-label")).
		Withs(children)
}

// FieldHorizontal creates a horizontal field, including an empty body if
func FieldHorizontal(children ...any) *Element {
	e := Elem(html.Div).
		With(Class("field")).
		With(Horizontal)

	label := Elem(html.Div).
		With(Class("field-label"))

	body := Elem(html.Div).
		With(Class("field-body"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			switch {
			case c.hasClass("field-label"):
				label = c
			case c.hasClass("field-body"):
				body = c
			default:
				body.With(c)
			}
		default:
			e.With(c)
		}
	}

	return e.
		With(label).
		With(body)
}

type fieldHorizontal struct {
	label        *Element
	body         *Element
	bodyChildren []any
	elemChildren []any
}

func (f *fieldHorizontal) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			switch {
			case c.hasClass("field-label"):
				f.label = c
			case c.hasClass("field-body"):
				f.body = c
			default:
				f.bodyChildren = append(f.bodyChildren, c)
			}
		case []any:
			f.addChildren(c)
		default:
			f.elemChildren = append(f.elemChildren, c)
		}
	}
}

func (f *fieldHorizontal) elem() *Element {
	var label *Element
	if f.label != nil {
		label = f.label
	} else {
		label = Elem(html.Div).
			With(Class("field-label"))
	}

	var body *Element
	if f.body != nil {
		body = f.body.Withs(f.bodyChildren)
	} else {
		body = Elem(html.Div).
			With(Class("field-body")).
			Withs(f.bodyChildren)
	}

	return Elem(html.Div).
		With(Class("field")).
		With(Horizontal).
		Withs(f.elemChildren).
		With(label).
		With(body)
}
