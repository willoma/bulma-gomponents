package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents/html"
)

// Progress creates a progress bar.
//
// The following modifiers change the progress bar color:
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
//
// The following modifiers change the progress bar size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func Progress(value, max int, children ...any) *Element {
	return Elem(html.Progress).
		With(Class("progress")).
		With(html.Value(strconv.Itoa(value))).
		With(html.Max(strconv.Itoa(max))).
		Withs(children)
}

// ProgressIndeterminate creates an animated progress bar with indeterminate
// value.
//
// The following modifiers change the progress bar color:
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
//
// The following modifiers change the progress bar size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func ProgressIndeterminate(children ...any) *Element {
	return Elem(html.Progress).
		With(Class("progress")).
		With(html.Max("100")).
		Withs(children)
}
