package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

const (
	Stack1x = Class("fa-stack-1x")
	Stack2x = Class("fa-stack-2x")
	Inverse = Class("fa-inverse")
)

// Stack stacks Font Awesome children icons.
func Stack(children ...any) b.Element {
	s := &stack{el.Span(b.Class("fa-stack"))}
	s.With(children...)
	return s
}

type stack struct {
	b.Element
}

// ZIndex sets the z-index of a stacked icon.
func ZIndex(value int) b.Styles {
	return b.Style("--fa-stack-z-index", strconv.Itoa(value))
}

// InverseColor sets the color of an inversed stacked icon.
func InverseColor(value string) b.Styles {
	return b.Style("--fa-inverse", value)
}
