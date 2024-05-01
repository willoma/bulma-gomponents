package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Progress creates a progress bar.
//
// https://willoma.github.io/bulma-gomponents/progress.html
func Progress(value, max int, children ...any) e.Element {
	return e.Progress(
		e.Class("progress"),
		html.Value(strconv.Itoa(value)),
		html.Max(strconv.Itoa(max)),
		children,
	)
}

// ProgressIndeterminate creates an animated progress bar with indeterminate
// value.
//
// https://willoma.github.io/bulma-gomponents/progress.html
func ProgressIndeterminate(children ...any) e.Element {
	return e.Progress(
		e.Class("progress"),
		html.Max("100"),
		children,
	)
}
