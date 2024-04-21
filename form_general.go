package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Field creates a field element.
//
// The following modifiers change the field behaviour:
//   - Addons: attach the children controls (Input*, Button* or Select) together
//   - AddonsCentered: attach the children controls (Input*, Button* or Select)
//     together and center them (Addons is not required)
//   - AddonsRight: attach the children controls (Input*, Button* or Select)
//     together and alight them to the right (Addons is not required)
//   - Grouped: group the children controls together, without attaching them
//   - GroupedCentered: group the children controls together, without attaching
//     them and center them (Grouped is not required)
//   - GroupedRight: group the children controls together, without attaching
//     them and align them to the right (Grouped is not required)
//   - GroupedMultiline: group the children controls together, without attaching
//     them and allow them to fill up multiple lines (Grouped is not required)
//   - Horizontal: make the field horizontal
func Field(children ...any) Element {
	return Elem(html.Div, Class("field"), children)
}

// Label creates a label element, to be used as a child of a Field.
func Label(children ...any) Element {
	return new(label).With(children...)
}

type label struct {
	children []any
}

func (l *label) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			l.With(c...)
		default:
			l.children = append(l.children, c)
		}
	}

	return l
}

func (l *label) Render(w io.Writer) error {
	return Elem(html.Label, Class("label"), l.children).Render(w)
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

// FieldHorizontal creates a horizontal field, including an empty body if needed.
func FieldHorizontal(children ...any) Element {
	return new(fieldHorizontal).With(children...)
}

type fieldHorizontal struct {
	labelChildren []any
	bodyChildren  []any
	fieldChildren []any
}

func (f *fieldHorizontal) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onField:
			f.fieldChildren = append(f.fieldChildren, c...)
		case onLabel:
			f.labelChildren = append(f.labelChildren, c...)
		case onBody:
			f.bodyChildren = append(f.bodyChildren, c...)
		case *label:
			f.labelChildren = append(f.labelChildren, c)
		case Element:
			f.bodyChildren = append(f.bodyChildren, c)
		case []any:
			f.With(c...)
		default:
			f.fieldChildren = append(f.fieldChildren, c)
		}
	}

	return f
}

func (f *fieldHorizontal) Render(w io.Writer) error {
	return Elem(
		html.Div,
		Class("field"),
		Horizontal,
		f.fieldChildren,
		Elem(html.Div, Class("field-label"), f.labelChildren),
		Elem(html.Div, Class("field-body"), f.bodyChildren),
	).Render(w)
}
