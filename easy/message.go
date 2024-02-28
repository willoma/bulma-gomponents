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
	message := b.Elem(html.Article).
		With(b.Class("message"))

	body := b.Elem(html.Div).
		With(b.Class("message-body"))

	var (
		title  string
		delete *b.Element
	)
	for _, c := range children {
		switch c := c.(type) {
		case MessageTitle:
			title = string(c)
		case MessageDeleteOnClick:
			delete = b.Delete(
				b.OnClick(string(c)),
			)
		case b.Class, b.ColorClass, b.Styles:
			message.With(c)
		case gomponents.Node:
			if b.IsAttribute(c) {
				message.With(c)
			} else {
				body.With(c)
			}
		default:
			body.With(c)
		}
	}

	if title != "" {
		header := b.Elem(html.Div).
			With(b.Class("message-header")).
			With(b.Elem(html.P).
				With(title),
			)

		if delete != nil {
			header.With(delete)
		}
		message.With(header)
	}

	return message.With(body)
}
