package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type fade struct {
	animationBase

	minOpacity     float64
	zeroMinOpacity bool
}

func Fade(options ...func(any)) e.ParentModifier {
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
	} else if a.zeroMinOpacity {
		p.With(e.Style("--fa-fade-opacity", "0"))
	}

	a.animationBase.modifyParent(p)
}

func (a *fade) If(cond bool) e.ParentModifier {
	return &conditionalAnimation{animation: a, cond: cond}
}

func MinOpacity(opacity float64) func(any) {
	return func(a any) {
		switch a := a.(type) {
		case *fade:
			a.zeroMinOpacity = opacity == 0
			a.minOpacity = opacity
		case *beatFade:
			a.zeroMinOpacity = opacity == 0
			a.minOpacity = opacity
		}
	}
}
