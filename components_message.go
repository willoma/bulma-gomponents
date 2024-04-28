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
		message: Elem(html.Article, Class("message")),
		body:    Elem(html.Div, Class("message-body")),
	}
	m.With(children...)
	return m
}

type message struct {
	message Element
	header  Element
	body    Element

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
			m.message.With(c...)
		case MessageTitle:
			m.addToHeader(Elem(html.P, string(c)))
		case *delete:
			m.addToHeader(c)
		case Class, Classer, Classeser, Styles:
			m.message.With(c)
		case gomponents.Node:
			if isAttribute(c) {
				m.message.With(c)
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
			m.message.With(m.header)
		}
		m.message.With(m.body)
	})
	return m.message.Render(w)
}

func (m *message) Clone() Element {
	return &message{
		message: m.message.Clone(),
		header:  m.header.Clone(),
		body:    m.body.Clone(),

		title:  m.title,
		delete: m.delete.Clone(),
	}
}
