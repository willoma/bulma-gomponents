package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type MessageTitle string

// Message creates a message.
//
// https://willoma.github.io/bulma-gomponents/message.html
func Message(children ...any) Element {
	m := &message{
		Element: Elem(html.Article, Class("message")),
		body:    Elem(html.Div, Class("message-body")),
	}
	m.With(children...)
	return m
}

type message struct {
	Element
	header Element
	body   Element

	rendered sync.Once

	title  MessageTitle
	delete Element
}

func (m *message) addToHeader(children ...any) {
	if m.header == nil {
		m.header = Elem(html.Div, Class("message-header"))
	}
	m.header.With(children...)
}

func (m *message) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onHeader:
			m.addToHeader(c...)
		case onBody:
			m.body.With(c...)
		case onMessage:
			m.Element.With(c...)
		case MessageTitle:
			m.addToHeader(Elem(html.P, string(c)))
		case *delete:
			m.addToHeader(c)
		case Class, Classer, Classeser, Styles:
			m.Element.With(c)
		case gomponents.Node:
			if isAttribute(c) {
				m.Element.With(c)
			} else {
				m.body.With(c)
			}
		case []any:
			m.With(c...)
		default:
			m.body.With(c)
		}
	}

	return m
}

func (m *message) Render(w io.Writer) error {
	m.rendered.Do(func() {
		if m.header != nil {
			m.Element.With(m.header)
		}
		m.Element.With(m.body)
	})
	return m.Element.Render(w)
}
