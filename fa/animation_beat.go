package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
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

func (a *beat) ModifyParent(p e.Element) {
	p.With(Class("fa-beat"))

	if a.maxScale != 0 {
		p.With(e.Style(
			"--fa-beat-scale",
			strconv.FormatFloat(a.maxScale, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(p)
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
