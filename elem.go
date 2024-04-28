package bulma

import (
	"bytes"
	"fmt"
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type Element interface {
	gomponents.Node

	With(...any) Element

	Clone() Element
}

type ParentModifier interface {
	ModifyParent(parent Element)
}

type ParentModifierAndElement interface {
	Element
	ParentModifier
}

type ParentModifierAndNode interface {
	gomponents.Node
	ParentModifier
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

// isAttribute returns true if the provided argument is a gomponents.Node with
// type gomponents.AttributeType.
func isAttribute(node any) bool {
	desc, ok := node.(nodeTypeDescriber)
	return ok && desc.Type() == gomponents.AttributeType
}

// With adds childs to the element. It accepts the following arguments types:
//   - func(...gomponents.Node) gomponents.Node: change the gomponents.Node
//     element for this element
//   - Class: add a class to the element
//   - Classer: add a class to the element
//   - Classeser: add multiple classes to the element
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
		case ParentModifierAndElement:
			c.ModifyParent(e)
			e.elements = append(e.elements, c)
		case ParentModifierAndNode:
			c.ModifyParent(e)
			e.elements = append(e.elements, c)
		case ParentModifier:
			c.ModifyParent(e)
		case Element:
			e.elements = append(e.elements, c)
		case gomponents.Node:
			if isAttribute(c) {
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

func (e *element) Clone() Element {
	classes := make(map[string]bool, len(e.classes))
	for class, ok := range e.classes {
		classes[class] = ok
	}

	stylesCollection := make(map[string]string, len(e.stylesCollection))
	for style, value := range e.stylesCollection {
		stylesCollection[style] = value
	}

	attributes := make([]gomponents.Node, len(e.attributes))
	for i, attr := range e.attributes {
		attributes[i] = attr
	}

	elements := make([]gomponents.Node, len(e.elements))
	for i, elem := range e.elements {
		switch elem := elem.(type) {
		case Element:
			elements[i] = elem.Clone()
		default:
			elements[i] = elem
		}
	}

	return &element{
		elemFn: e.elemFn,

		hasIcons:                     e.hasIcons,
		spanAroundNonIconsIfHasIcons: e.spanAroundNonIconsIfHasIcons,
		spanAroundNonIconsAlways:     e.spanAroundNonIconsAlways,

		classes:          classes,
		stylesCollection: stylesCollection,

		attributes: attributes,
		elements:   elements,
	}
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

// Prepare pre-renders a node in memory for future uses.
func Prepare(e Element) gomponents.Node {
	var buf bytes.Buffer
	e.Render(&buf)

	if modifier, ok := e.(ParentModifier); ok {
		return &preparedElementWithParentModifier{
			preparedElement: preparedElement{buf.Bytes()},
			modifyFn:        modifier.ModifyParent,
		}
	}

	return &preparedElement{content: buf.Bytes()}
}

type preparedElement struct {
	content []byte
}

type preparedElementWithParentModifier struct {
	preparedElement
	modifyFn func(Element)
}

func (e *preparedElement) Render(w io.Writer) error {
	_, err := w.Write(e.content)
	return err
}

func (e *preparedElementWithParentModifier) ModifyParent(parent Element) {
	e.modifyFn(parent)
}
