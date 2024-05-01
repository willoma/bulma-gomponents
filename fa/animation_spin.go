package fa

import (
	e "github.com/willoma/gomplements"
)

type spin struct {
	animationBase

	pulse   bool
	reverse bool
}

func Spin(options ...func(Animation)) Animation {
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

	a.animationBase.ModifyParent(p)
}

func Pulse(a Animation) {
	if s, ok := a.(*spin); ok {
		s.pulse = true
	}
}

func Reverse(a Animation) {
	if s, ok := a.(*spin); ok {
		s.reverse = true
	}
}
