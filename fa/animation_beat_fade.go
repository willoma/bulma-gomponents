package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
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

func (a *beatFade) ModifyParent(e b.Element) {
	e.With(Class("fa-beat-fade"))

	if a.maxScale != 0 {
		e.With(b.Style(
			"--fa-beat-fade-scale",
			strconv.FormatFloat(a.maxScale, 'f', 2, 64),
		))
	}

	if a.minOpacity != 0 {
		e.With(b.Style(
			"--fa-beat-fade-opacity",
			strconv.FormatFloat(a.minOpacity, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(e)
}
