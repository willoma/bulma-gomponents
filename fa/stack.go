package fa

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

const (
	Stack1x = Class("fa-stack-1x")
	Stack2x = Class("fa-stack-2x")
)

// Stack stacks Font-Awesome children icons.
//   - when a child is a Class, it is added as a b.Class to the stack element
//   - other children types are added as-is to the stack element
//
// Stacked icons must be created with the FA function.
func Stack(children ...any) *b.Element {
	e := b.Elem(html.Span).With(b.Class("fa-stack"))
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			e.With(b.Class(c))
		default:
			e.With(c)
		}
	}
	return e
}
