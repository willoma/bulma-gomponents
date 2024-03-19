package easy

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

type (
	MessageTitle         string // set the header of a message
	MessageDeleteOnClick string // set a javascript action on click on the delete botton
)

// Message creates a message.
//   - when a child is a MessageTitle, the message header is set to that string
//   - when a child is a MessageDeleteOnClick, the "delete" button is added to
//     the header (only if a MessageTitle is provided) with that javascript
//     executed on click
//   - when a child is a gomponents.Node with type gomponents.ElementType, it
//     is added to the body
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the message element
//   - when a child is a b.Class, b.ColorClass or b.Styles, it is added to the
//     message element
//   - all other types are added to the body
//
// When there is no MessageTitle in the children, the message is displayed
// with the "message body only" style.
func Message(children ...any) *b.Element {
	m := &message{}
	m.addChildren(children)
	return m.elem()
}

type message struct {
	title           string
	delete          *b.Element
	messageChildren []any
	bodyChildren    []any
}

func (m *message) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case MessageTitle:
			m.title = string(c)
		case MessageDeleteOnClick:
			m.delete = b.Delete(
				b.OnClick(string(c)),
			)
		case b.Class, b.ColorClass, b.Styles:
			m.messageChildren = append(m.messageChildren, c)
		case gomponents.Node:
			if b.IsAttribute(c) {
				m.messageChildren = append(m.messageChildren, c)
			} else {
				m.bodyChildren = append(m.bodyChildren, c)
			}
		case []any:
			m.addChildren(c)
		default:
			m.bodyChildren = append(m.bodyChildren, c)
		}
	}
}

func (m *message) elem() *b.Element {
	message := b.Elem(html.Article).
		With(b.Class("message")).
		Withs(m.messageChildren)

	body := b.Elem(html.Div).
		With(b.Class("message-body")).
		Withs(m.bodyChildren)

	if m.title != "" {
		header := b.Elem(html.Div).
			With(b.Class("message-header")).
			With(b.Elem(html.P).
				With(m.title),
			)

		if m.delete != nil {
			header.With(m.delete)
		}
		message.With(header)
	}

	return message.With(body)
}
