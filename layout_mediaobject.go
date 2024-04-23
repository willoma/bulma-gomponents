package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Media creates a media element.
//
// http://willoma.github.io/bulma-gomponents/media-object.html
func Media(children ...any) Element {
	m := &media{Element: Elem(html.Article, Class("media"))}
	m.With(children...)
	return m
}

type media struct {
	Element
	left    Element
	content Element
	right   Element

	rendered sync.Once
}

func (m *media) addToLeft(children ...any) {
	if m.left == nil {
		m.left = Elem(html.Div, Class("media-left"))
	}
	m.left.With(children...)
}

func (m *media) addToContent(children ...any) {
	if m.content == nil {
		m.content = Elem(html.Div, Class("media-content"))
	}
	m.content.With(children...)
}

func (m *media) addToRight(children ...any) {
	if m.right == nil {
		m.right = Elem(html.Div, Class("media-right"))
	}
	m.right.With(children...)
}

func (m *media) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onMedia:
			m.Element.With(c...)
		case onContent:
			m.addToContent(c...)
		case mediaLeft:
			m.addToLeft(c...)
		case mediaRight:
			m.addToRight(c...)
		case Element:
			m.addToContent(c)
		case gomponents.Node:
			if isAttribute(c) {
				m.Element.With(c)
			} else {
				m.addToContent(c)
			}
		case []any:
			m.With(c...)
		default:
			m.Element.With(c)
		}
	}

	return m
}

func (m *media) Render(w io.Writer) error {
	m.rendered.Do(func() {
		if m.left != nil {
			m.Element.With(m.left)
		}

		if m.content != nil {
			m.Element.With(m.content)
		}

		if m.right != nil {
			m.Element.With(m.right)
		}
	})

	return m.Element.Render(w)
}

// MediaLeft marks children as belonging to the left part of a media element.
//
// http://willoma.github.io/bulma-gomponents/media-object.html
func MediaLeft(children ...any) mediaLeft {
	return mediaLeft(children)
}

type mediaLeft []any

// MediaRight marks children as belonging to the right part of a media element.
//
// http://willoma.github.io/bulma-gomponents/media-object.html
func MediaRight(children ...any) mediaRight {
	return mediaRight(children)
}

type mediaRight []any
