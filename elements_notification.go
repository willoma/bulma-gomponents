package bulma

import "github.com/maragudk/gomponents/html"

// Notification creates a notification element.
//
// When you provide a Delete button as a child, you may add the following to
// the Delete children in order to close the notification:
//
//	OnClick(JSRemoveThisNotification)
//
// The following modifiers change the notification color:
//   - White
//   - Black
//   - Light
//   - Dark
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
func Notification(children ...any) *Element {
	return Elem(html.Div, Class("notification"), children)
}
