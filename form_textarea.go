package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Textarea creates a textarea element.
//
// The following modifiers change the textarea behaviour:
//   - Rows(int): set the textarea heights, in number of rows
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - html.Disabled(): disable the input
//   - html.ReadOnly(): forbid modifications
//   - Static: remove specific styling but maintain vertical spacing
//   - FixedSize: disable the textarea resizing possibility
//
// The following modifiers change the textarea color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the textarea size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func Textarea(children ...any) *Element {
	return Elem(html.Textarea).
		With(Class("textarea")).
		Withs(children)
}

// Rows changes a textarea height.
func Rows(rows int) gomponents.Node {
	return html.Rows(strconv.Itoa(rows))
}
