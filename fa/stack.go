package fa

import (
	"io"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

const (
	Stack1x = Class("fa-stack-1x")
	Stack2x = Class("fa-stack-2x")
)

// Stack stacks Font Awesome children icons.
//   - when a child is a Class, it is added as a b.Class to the stack element
//   - other children types are added as-is to the stack element
//
// Stacked icons must be created with the FA function.
func Stack(children ...any) b.Element {
	return new(stack).With(children...)
}

type stack struct {
	children []any
}

func (s *stack) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			s.children = append(s.children, b.Class(c))
		case []any:
			s.With(c...)
		default:
			s.children = append(s.children, c)
		}
	}

	return s
}

func (s *stack) Render(w io.Writer) error {
	return b.Elem(html.Span, b.Class("fa-stack"), s.children).Render(w)
}
