package bulma

import (
	"io"
	"sync"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
)

type MessageTitle string

// Message creates a message.
//
// https://willoma.github.io/bulma-gomponents/message.html
func Message(children ...any) e.Element {
	m := &message{
		message: e.Article(e.Class("message")),
		body:    e.Div(e.Class("message-body")),
	}
	m.With(children...)
	return m
}

type message struct {
	message e.Element
	header  e.Element
	body    e.Element

	rendered sync.Once

	title  MessageTitle
	delete e.Element
}

func (m *message) addToHeader(children ...any) {
	if m.header == nil {
		m.header = e.Div(e.Class("message-header"))
	}
	m.header.With(children...)
}

func (m *message) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onHeader:
			m.addToHeader(c...)
		case onBody:
			m.body.With(c...)
		case onMessage:
			m.message.With(c...)
		case MessageTitle:
			m.addToHeader(e.P(string(c)))
		case *delete:
			m.addToHeader(c)
		case e.Class, e.Classer, e.Classeser, e.Styles:
			m.message.With(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (m *message) Clone() e.Element {
	return &message{
		message: m.message.Clone(),
		header:  m.header.Clone(),
		body:    m.body.Clone(),

		title:  m.title,
		delete: m.delete.Clone(),
	}
}
