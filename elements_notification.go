package bulma

import (
	e "github.com/willoma/gomplements"
)

// Notification creates a notification element.
//
// https://willoma.github.io/bulma-gomponents/notification.html
func Notification(children ...any) e.Element {
	return e.Div(e.Class("notification"), children)
}
