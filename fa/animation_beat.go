package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type beat struct {
	animationBase

	maxScale     float64
	zeroMaxScale bool
}

func Beat(options ...func(any)) e.ParentModifier {
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
	} else if a.zeroMaxScale {
		p.With(e.Style("--fa-beat-scale", "0"))
	}

	a.animationBase.modifyParent(p)
}

func (a *beat) If(cond bool) Animation {
	if cond {
		return a
	}

	return nil
}

func MaxScale(scale float64) func(any) {
	return func(a any) {
		switch a := a.(type) {
		case *beat:
			a.zeroDelay = scale == 0
			a.maxScale = scale
		case *beatFade:
			a.zeroDelay = scale == 0
			a.maxScale = scale
		}
	}
}
