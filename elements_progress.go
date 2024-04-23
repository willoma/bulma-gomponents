package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents/html"
)

// Progress creates a progress bar.
//
// https://willoma.github.io/bulma-gomponents/progress.html
func Progress(value, max int, children ...any) Element {
	return Elem(
		html.Progress,
		Class("progress"),
		html.Value(strconv.Itoa(value)),
		html.Max(strconv.Itoa(max)),
		children,
	)
}

// ProgressIndeterminate creates an animated progress bar with indeterminate
// value.
//
// https://willoma.github.io/bulma-gomponents/progress.html
func ProgressIndeterminate(children ...any) Element {
	return Elem(
		html.Progress,
		Class("progress"),
		html.Max("100"),
		children,
	)
}
