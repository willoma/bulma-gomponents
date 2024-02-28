package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Elem creates a base element, based on the provided gomponents.Node function.
func Elem(elemFn func(...gomponents.Node) gomponents.Node) *Element {
	e := &Element{
		elemFn:           elemFn,
		classes:          map[string]bool{},
		stylesCollection: map[string]string{},
	}
	return e
}

// Element is a single element.
type Element struct {
	elemFn func(...gomponents.Node) gomponents.Node

	hasIcons                     bool
	spanAroundNonIconsIfHasIcons bool
	spanAroundNonIconsAlways     bool

	classes          map[string]bool
	stylesCollection map[string]string

	attributes []gomponents.Node
	elements   []gomponents.Node
}

type nodeTypeDescriber interface {
	Type() gomponents.NodeType
}

// IsAttribute returns true if the provided argument is a gomponents.Node with
// type gomponents.AttributeType.
func IsAttribute(node any) bool {
	desc, ok := node.(nodeTypeDescriber)
	return ok && desc.Type() == gomponents.AttributeType
}

// With adds a child to the element. It accepts the following arguments types:
//   - func(...gomponents.Node) gomponents.Node: change the gomponents.Node
//     element for this element
//   - Class: add a class to the element
//   - MultiClass: add multiple classes to the element
//   - ColorClass: add a color class to the element
//   - Styles: add one or multiple CSS styles to the element
//   - string: add a string to the element (using gomponents.Text)
//   - *Element: add the provided element as a child
//   - container: add the provided element as a child
//   - func(...any) *Element: apply the following parts of that element to
//     the current element: classes, styles, attributes, children
//   - gomponents.Node with type gomponents.AttributeType: add the provided
//     attribute to this element
//   - gomponents.Node with another time: add the provided element as a child
//
// Any other type is ignored.
func (e *Element) With(c any) *Element {
	if e == nil {
		e = Elem(html.Div)
	}

	switch c := c.(type) {
	case func(...gomponents.Node) gomponents.Node:
		e.elemFn = c
	case Class:
		e.classes[string(c)] = true
	case MultiClass:
		for _, cl := range c.Responsive {
			e.classes[cl] = true
		}
		for _, cl := range c.Static {
			e.classes[cl] = true
		}
	case ColorClass:
		e.classes["is-"+c.class] = true
		if c.light {
			e.classes["is-light"] = true
		}
	case Styles:
		for prop, val := range c {
			e.stylesCollection[prop] = val
		}
	case string:
		e.elements = append(e.elements, gomponents.Text(c))
	case *Element:
		if c.hasClass("icon") {
			e.hasIcons = true
		}

		if c.hasClass("tile") && !e.hasClass("tile") {
			c.classes["is-ancestor"] = true
		}

		if c.hasClass("navbar") {
			if c.hasClass(string(FixedBottom)) {
				e.With(NavbarFixedBottom)
			} else if c.hasClass(string(FixedTop)) {
				e.With(NavbarFixedTop)
			}
		}

		e.elements = append(e.elements, c)
	case container:
		elem := Element(*c)
		e.elements = append(e.elements, &elem)
	case func(...any) *Element:
		other := c()
		for c := range other.classes {
			e.classes[c] = true
		}
		for p, v := range other.stylesCollection {
			e.stylesCollection[p] = v
		}
		e.attributes = append(e.attributes, other.attributes...)
		e.elements = append(e.elements, other.elements...)
	case gomponents.Node:
		if IsAttribute(c) {
			e.attributes = append(e.attributes, c)
		} else {
			e.elements = append(e.elements, c)
		}
	}
	return e
}

// Withs adds multiple childs to the element. It simply calls With, for all
// elements in the slice.
func (e *Element) Withs(children []any) *Element {
	for _, c := range children {
		e.With(c)
	}
	return e
}

func (e *Element) hasClass(cl string) bool {
	return e.classes[cl]
}

func (e *Element) getChildren() []gomponents.Node {
	if e == nil {
		return nil
	}
	children := []gomponents.Node{}

	classes := ""
	for cl := range e.classes {
		classes += " " + cl
	}
	if classes != "" {
		children = append(children, html.Class(string(classes[1:])))
	}

	styles := ""
	for prop, val := range e.stylesCollection {
		styles += prop + ":" + val + ";"
	}
	if styles != "" {
		children = append(children, html.StyleAttr(styles))
	}

	for _, attr := range e.attributes {
		children = append(children, attr)
	}

	if e.spanAroundNonIconsAlways || (e.spanAroundNonIconsIfHasIcons && e.hasIcons) {
		for _, c := range e.elements {
			switch c := c.(type) {
			case *Element:
				if c.hasClass("icon") {
					children = append(children, c)
				} else {
					children = append(children, Elem(html.Span).With(c))
				}
			default:
				children = append(children, Elem(html.Span).With(c))
			}
		}
	} else {
		children = append(children, e.elements...)
	}

	return children
}

// Render renders the element. With this function, an *Element is a
// gomponents.Node.
func (e *Element) Render(w io.Writer) error {
	if e == nil {
		return nil
	}
	return e.elemFn(e.getChildren()...).Render(w)
}
