package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

// Media creates a media element.
//
// http://willoma.github.io/bulma-gomponents/media-object.html
func Media(children ...any) e.Element {
	m := &media{media: e.Article(e.Class("media"))}
	m.With(children...)
	return m
}

type media struct {
	media   e.Element
	left    e.Element
	content e.Element
	right   e.Element

	rendered sync.Once
}

func (m *media) addToLeft(children ...any) {
	if m.left == nil {
		m.left = e.Div(e.Class("media-left"))
	}
	m.left.With(children...)
}

func (m *media) addToContent(children ...any) {
	if m.content == nil {
		m.content = e.Div(e.Class("media-content"))
	}
	m.content.With(children...)
}

func (m *media) addToRight(children ...any) {
	if m.right == nil {
		m.right = e.Div(e.Class("media-right"))
	}
	m.right.With(children...)
}

func (m *media) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onMedia:
			m.media.With(c...)
		case onContent:
			m.addToContent(c...)
		case mediaLeft:
			m.addToLeft(c...)
		case mediaRight:
			m.addToRight(c...)
		case e.Element:
			m.addToContent(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
				m.media.With(c)
			} else {
				m.addToContent(c)
			}
		case []any:
			m.With(c...)
		default:
			m.media.With(c)
		}
	}

	return m
}

func (m *media) Render(w io.Writer) error {
	m.rendered.Do(func() {
		if m.left != nil {
			m.media.With(m.left)
		}

		if m.content != nil {
			m.media.With(m.content)
		}

		if m.right != nil {
			m.media.With(m.right)
		}
	})

	return m.media.Render(w)
}

func (m *media) Clone() e.Element {
	return &media{
		media:   m.media.Clone(),
		left:    m.left.Clone(),
		content: m.content.Clone(),
		right:   m.right.Clone(),
	}
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
