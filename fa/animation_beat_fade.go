package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type beatFade struct {
	animationBase

	maxScale   float64
	minOpacity float64
}

func BeatFade(options ...func(Animation)) Animation {
	a := &beatFade{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *beatFade) ModifyParent(p e.Element) {
	p.With(Class("fa-beat-fade"))

	if a.maxScale != 0 {
		p.With(e.Style(
			"--fa-beat-fade-scale",
			strconv.FormatFloat(a.maxScale, 'f', 2, 64),
		))
	}

	if a.minOpacity != 0 {
		p.With(e.Style(
			"--fa-beat-fade-opacity",
			strconv.FormatFloat(a.minOpacity, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(p)
}
