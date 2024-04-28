package bulma

import "github.com/maragudk/gomponents/html"

// Notification creates a notification element.
//
// https://willoma.github.io/bulma-gomponents/notification.html
func Notification(children ...any) Element {
	return Elem(html.Div, Class("notification"), children)
}
