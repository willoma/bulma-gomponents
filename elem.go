package bulma

import (
	"fmt"
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type Element interface {
	gomponents.Node

	With(...any) Element
}

type ParentModifier interface {
	Element

	ModifyParent(parent Element)
}

type elemOption int

const (
	elemOptionSpanAroundNonIconsAlways elemOption = iota
	elemOptionSpanAroundNonIconsIfHasIcons
)

// Elem creates a base element, based on the provided gomponents.Node function.
func Elem(elemFn func(...gomponents.Node) gomponents.Node, children ...any) *element {
	e := &element{
		elemFn:           elemFn,
		classes:          map[string]bool{},
		stylesCollection: map[string]string{},
	}
	e.With(children...)
	return e
}

// element is a single element.
type element struct {
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

// With adds childs to the element. It accepts the following arguments types:
//   - func(...gomponents.Node) gomponents.Node: change the gomponents.Node
//     element for this element
//   - Class: add a class to the element
//   - Classer: add a class to the element
//   - Classeser: add multiple classes to the element
//   - ExternalClassesAndStyles: add one or multiple classes and CSS styles to
//     the element
//   - MultiClass: add multiple classes to the element
//   - Styles: add one or multiple CSS styles to the element
//   - ID: define the ID attribute for the element
//   - string: add a string to the element (using gomponents.Text)
//   - Element: add the provided element as a child
//   - container: add the provided element as a child
//   - gomponents.Node with type gomponents.AttributeType: add the provided
//     attribute to this element
//   - gomponents.Node with another time: add the provided element as a child
//   - []any: add the slice items by executing With recursively
//
// Any other type is ignored.
func (e *element) With(children ...any) Element {
	if e == nil {
		e = Elem(html.Div)
	}

	for _, c := range children {
		switch c := c.(type) {
		case elemOption:
			switch c {
			case elemOptionSpanAroundNonIconsIfHasIcons:
				e.spanAroundNonIconsIfHasIcons = true
			case elemOptionSpanAroundNonIconsAlways:
				e.spanAroundNonIconsAlways = true
			}
		case func(...gomponents.Node) gomponents.Node:
			e.elemFn = c
		case Class:
			e.classes[string(c)] = true
		case Classer:
			e.classes[string(c.Class())] = true
		case Classeser:
			for _, cl := range c.Classes() {
				e.classes[string(cl)] = true
			}
		case ExternalClassesAndStyles:
			cls, st := c.ClassesAndStyles()
			for _, cl := range cls {
				e.classes[string(cl)] = true
			}
			for prop, val := range st {
				e.stylesCollection[prop] = val
			}
		case MultiClass:
			for _, cl := range c.Responsive {
				e.classes[cl] = true
			}
			for _, cl := range c.Static {
				e.classes[cl] = true
			}
		case Styles:
			for prop, val := range c {
				e.stylesCollection[prop] = val
			}
		case ID:
			e.attributes = append(e.attributes, html.ID(string(c)))
		case string:
			e.elements = append(e.elements, gomponents.Text(c))
		case IconElem:
			e.hasIcons = true
			e.elements = append(e.elements, c)
		case ParentModifier:
			c.ModifyParent(e)
			e.elements = append(e.elements, c)
		case Element:
			e.elements = append(e.elements, c)
		case gomponents.Node:
			if IsAttribute(c) {
				e.attributes = append(e.attributes, c)
			} else {
				e.elements = append(e.elements, c)
			}
		case []any:
			e.With(c...)
		case fmt.Stringer:
			e.elements = append(e.elements, gomponents.Text(c.String()))
		}
	}

	return e
}

func (e *element) getChildren() []gomponents.Node {
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
			if ic, ok := c.(IconElem); ok {
				children = append(children, ic)
			} else {
				children = append(children, Elem(html.Span, c))
			}
		}
	} else {
		children = append(children, e.elements...)
	}

	return children
}

// Render renders the element. With this function, an Element is a
// gomponents.Node.
func (e *element) Render(w io.Writer) error {
	if e == nil {
		return nil
	}
	return e.elemFn(e.getChildren()...).Render(w)
}
