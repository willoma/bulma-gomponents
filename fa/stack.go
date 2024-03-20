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
	s := &stack{}
	s.addChildren(children)
	return s.elem()
}

type stack struct {
	children []any
}

func (s *stack) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			s.children = append(s.children, b.Class(c))
		case []any:
			s.addChildren(c)
		default:
			s.children = append(s.children, c)
		}
	}
}

func (s *stack) elem() *b.Element {
	return b.Elem(html.Span, b.Class("fa-stack"), s.children)
}
