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
//   - when a child is an *Element, it is added in the content part of the media
//     element
//   - when a child is a return value of Container*, it is added in the content
//     part of the media element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as a direct child of to the media element
//   - when a child is a gomponents.Node with another type, it is added in the
//     content part of the media element
//   - other children types are added as direct children of to the media element
//
// Each of the left, content and right parts is only included if it has content.
func Media(children ...any) *Element {
	e := Elem(html.Article).
		With(Class("media"))

	var left, content, right []any
	for _, c := range children {
		switch c := c.(type) {
		case mediaPart:
			if c.right {
				right = append(right, c.children...)
			} else {
				left = append(left, c.children...)
			}
		case *Element, container:
			content = append(content, c)
		case gomponents.Node:
			if IsAttribute(c) {
				e.With(c)
			} else {
				content = append(content, c)
			}
		default:
			e.With(c)
		}
	}

	if len(left) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("media-left")).
				Withs(left),
		)
	}

	if len(content) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("media-content")).
				Withs(content),
		)
	}

	if len(right) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("media-right")).
				Withs(right),
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
