package fa

import (
	b "github.com/willoma/bulma-gomponents"
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

func (a *spin) applyTo(e b.Element) {
	if a.pulse {
		e.With(Class("fa-spin-pulse"))
	} else {
		e.With(Class("fa-spin"))
	}

	if a.reverse {
		e.With(Class("fa-spin-reverse"))
	}

	a.animationBase.applyTo(e)
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
