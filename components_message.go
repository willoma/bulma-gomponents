package bulma

import (
	"github.com/maragudk/gomponents/html"
)

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
	return Elem(html.Article, Class("message"), children)
}

// MessageHeader creates a message header.
func MessageHeader(children ...any) Element {
	return Elem(html.Div, Class("message-header"), children)
}

// MessageBody creates a message body.
func MessageBody(children ...any) Element {
	return Elem(html.Div, Class("message-body"), children)
}
