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

// NotificationCloseButton adds a "delete" button to a notification,
// with the adequate javascript in order to close it.
func NotificationCloseButton() e.Element {
	return Delete(e.OnClick(JSRemoveThisNotification))
}
