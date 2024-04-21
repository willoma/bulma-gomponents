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
// The following modifiers change the message color:
//   - Dark
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// When there is no MessageHeader in the children, the message is displayed
// with the "message body only" style.
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

func (m *message) ensureHeader(children ...any) {
	if m.header == nil {
		m.header = Elem(html.Div, Class("message-header"), children)
	} else {
		m.header.With(children...)
	}
}

func (m *message) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onHeader:
			m.ensureHeader(c...)
		case onBody:
			m.body.With(c...)
		case onMessage:
			m.Element.With(c...)
		case MessageTitle:
			m.ensureHeader(Elem(html.P, string(c)))
		case *delete:
			m.ensureHeader(c)
		case Class, Classer, Classeser, ExternalClassesAndStyles, MultiClass, Styles:
			m.Element.With(c)
		case gomponents.Node:
			if IsAttribute(c) {
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
