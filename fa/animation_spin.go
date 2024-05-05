package fa

import (
	e "github.com/willoma/gomplements"
)

type spin struct {
	animation

	pulse   bool
	reverse bool
}

func Spin(options ...func(any)) e.ParentModifier {
	a := &spin{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *spin) ModifyParent(p e.Element) {
	if a.pulse {
		p.With(Class("fa-spin-pulse"))
	} else {
		p.With(Class("fa-spin"))
	}

	if a.reverse {
		p.With(Class("fa-spin-reverse"))
	}

	a.animation.ModifyParent(p)
}

func Pulse(a any) {
	if s, ok := a.(*spin); ok {
		s.pulse = true
	}
}

func Reverse(a any) {
	if s, ok := a.(*spin); ok {
		s.reverse = true
	}
}
