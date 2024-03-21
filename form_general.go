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
func Field(children ...any) Element {
	return Elem(html.Div, Class("field"), children)
}

// Label creates a label element, to be used as a child of a Field.
func Label(children ...any) Element {
	return Elem(html.Label, Class("label"), children)
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
func Control(children ...any) Element {
	return Elem(html.Div, Class("control"), children)
}

// Help creates a help element, to be used as a child of a Field.
func Help(children ...any) Element {
	return Elem(html.P, Class("help"), children)
}

// FieldBody creates a field-body element, to be used as a child of a
// FieldHorizontal.
func FieldBody(children ...any) Element {
	return &fieldBody{Elem(html.Div, Class("field-body"), children)}
}

type fieldBody struct {
	Element
}

// FieldLabel creates a field-label element, to be used as a child of a
// FieldHorizontal.
func FieldLabel(children ...any) Element {
	return &fieldLabel{Elem(html.Div, Class("field-label"), children)}
}

type fieldLabel struct {
	Element
}

// FieldHorizontal creates a horizontal field, including an empty body if
func FieldHorizontal(children ...any) Element {
	f := &fieldHorizontal{}
	f.addChildren(children)
	return f.elem()
}

type fieldHorizontal struct {
	label        Element
	body         Element
	bodyChildren []any
	elemChildren []any
}

func (f *fieldHorizontal) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *fieldBody:
			f.body = c
		case *fieldLabel:
			f.label = c
		case Element:
			f.bodyChildren = append(f.bodyChildren, c)
		case []any:
			f.addChildren(c)
		default:
			f.elemChildren = append(f.elemChildren, c)
		}
	}
}

func (f *fieldHorizontal) elem() Element {
	var label Element
	if f.label != nil {
		label = f.label
	} else {
		label = Elem(html.Div, Class("field-label"))
	}

	var body Element
	if f.body != nil {
		body = f.body.With(f.bodyChildren...)
	} else {
		body = Elem(html.Div, Class("field-body"), f.bodyChildren)
	}

	return Elem(
		html.Div,
		Class("field"),
		Horizontal,
		f.elemChildren,
		label,
		body,
	)
}
