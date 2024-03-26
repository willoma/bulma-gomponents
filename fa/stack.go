package fa

import (
	"io"
	"strconv"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

const (
	Stack1x = Class("fa-stack-1x")
	Stack2x = Class("fa-stack-2x")
	Inverse = Class("fa-inverse")
)

// Stack stacks Font Awesome children icons.
func Stack(children ...any) b.Element {
	return new(stack).With(children...)
}

type stack struct {
	children []any
}

func (s *stack) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
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

// ZIndex sets the z-index of a stacked icon
func ZIndex(value int) b.Styles {
	return b.Style("--fa-stack-z-index", strconv.Itoa(value))
}

// InverseColor sets the color of an inversed stacked icon
func InverseColor(value string) b.Styles {
	return b.Style("--fa-inverse", value)
}
