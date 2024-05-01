package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
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

func (a *fade) ModifyParent(p e.Element) {
	p.With(Class("fa-fade"))

	if a.minOpacity != 0 {
		p.With(e.Style(
			"--fa-fade-opacity",
			strconv.FormatFloat(a.minOpacity, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(p)
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
