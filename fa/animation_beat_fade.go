package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type beatFade struct {
	animationBase

	maxScale     float64
	zeroMaxScale bool

	minOpacity     float64
	zeroMinOpacity bool
}

func BeatFade(options ...func(any)) e.ParentModifier {
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
	} else if a.zeroMaxScale {
		p.With(e.Style("--fa-beat-fade-scale", "0"))
	}

	if a.minOpacity != 0 {
		p.With(e.Style(
			"--fa-beat-fade-opacity",
			strconv.FormatFloat(a.minOpacity, 'f', 2, 64),
		))
	} else if a.zeroMinOpacity {
		p.With(e.Style("--fa-beat-fade-opacity", "0"))
	}

	a.animationBase.modifyParent(p)
}

func (a *beatFade) If(cond bool) Animation {
	if cond {
		return a
	}

	return nil
}
