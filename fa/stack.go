package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

const (
	Stack1x = Class("fa-stack-1x")
	Stack2x = Class("fa-stack-2x")
	Inverse = Class("fa-inverse")
)

// Stack stacks Font Awesome children icons.
func Stack(children ...any) e.Element {
	s := &stack{e.Span(e.Class("fa-stack"))}
	s.With(children...)
	return s
}

type stack struct {
	e.Element
}

// ZIndex sets the z-index of a stacked icon.
func ZIndex(value int) e.Styles {
	return e.Styles{"--fa-stack-z-index": strconv.Itoa(value)}
}

// InverseColor sets the color of an inversed stacked icon.
func InverseColor(value string) e.Styles {
	return e.Styles{"--fa-inverse": value}
}
