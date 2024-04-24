package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
)

type fade struct {
	animationBase

	minOpacity float64
}

func Fade(options ...func(Animation)) Animation {
	a := &fade{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *fade) applyTo(e b.Element) {
	e.With(Class("fa-fade"))

	if a.minOpacity != 0 {
		e.With(b.Style(
			"--fa-fade-opacity",
			strconv.FormatFloat(a.minOpacity, 'f', 2, 64),
		))
	}

	a.animationBase.applyTo(e)
}

func MinOpacity(opacity float64) func(Animation) {
	return func(a Animation) {
		switch a := a.(type) {
		case *fade:
			a.minOpacity = opacity
		case *beatFade:
			a.minOpacity = opacity
		}
	}
}
