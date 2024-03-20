package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type mediaPart struct {
	right    bool
	children []any
}

// Media creates a media element.
//   - when a child is a return value of MediaLeft, it is added in the left part
//     of the media element
//   - when a child is a return value of MediaRight, it is added in the right
//     part of the media element
//   - when a child is an Element, it is added in the content part of the media
//     element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as a direct child of to the media element
//   - when a child is a gomponents.Node with another type, it is added in the
//     content part of the media element
//   - other children types are added as direct children of to the media element
//
// Each of the left, content and right parts is only included if it has content.
func Media(children ...any) Element {
	m := &media{}
	m.addChildren(children)
	return m.elem()
}

type media struct {
	leftChildren    []any
	contentChildren []any
	rightChildren   []any
	elemChildren    []any
}

func (m *media) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case mediaPart:
			if c.right {
				m.rightChildren = append(m.rightChildren, c.children...)
			} else {
				m.leftChildren = append(m.leftChildren, c.children...)
			}
		case Element:
			m.contentChildren = append(m.contentChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				m.elemChildren = append(m.elemChildren, c)
			} else {
				m.contentChildren = append(m.contentChildren, c)
			}
		default:
			m.elemChildren = append(m.elemChildren, c)
		}
	}
}

func (m *media) elem() Element {
	e := Elem(html.Article, Class("media"), m.elemChildren)

	if len(m.leftChildren) > 0 {
		e.With(
			Elem(html.Div, Class("media-left"), m.leftChildren),
		)
	}

	if len(m.contentChildren) > 0 {
		e.With(
			Elem(html.Div, Class("media-content"), m.contentChildren),
		)
	}

	if len(m.rightChildren) > 0 {
		e.With(
			Elem(html.Div, Class("media-right"), m.rightChildren),
		)
	}

	return e
}

// MediaLeft marks children as belonging to the left part of a media element.
func MediaLeft(children ...any) mediaPart {
	return mediaPart{false, children}
}

// MediaRight marks children as belonging to the right part of a media element.
func MediaRight(children ...any) mediaPart {
	return mediaPart{true, children}
}
