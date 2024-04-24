package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
)

type beat struct {
	animationBase

	maxScale float64
}

func Beat(options ...func(Animation)) Animation {
	a := &beat{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *beat) applyTo(e b.Element) {
	e.With(Class("fa-beat"))

	if a.maxScale != 0 {
		e.With(b.Style(
			"--fa-beat-scale",
			strconv.FormatFloat(a.maxScale, 'f', 2, 64),
		))
	}

	a.animationBase.applyTo(e)
}

func MaxScale(scale float64) func(Animation) {
	return func(a Animation) {
		switch a := a.(type) {
		case *beat:
			a.maxScale = scale
		case *beatFade:
			a.maxScale = scale
		}
	}
}
