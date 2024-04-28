package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
)

type bounce struct {
	animationBase

	rebound     float64
	height      float64
	startScaleX float64
	startScaleY float64
	jumpScaleX  float64
	jumpScaleY  float64
	landScaleX  float64
	landScaleY  float64
}

func Bounce(options ...func(Animation)) Animation {
	a := &bounce{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *bounce) ModifyParent(e b.Element) {
	e.With(Class("fa-bounce"))

	if a.rebound != 0 {
		e.With(b.Style(
			"--fa-bounce-rebound",
			strconv.FormatFloat(a.rebound, 'f', 2, 64),
		))
	}

	if a.height != 0 {
		e.With(b.Style(
			"--fa-bounce-height",
			strconv.FormatFloat(a.height, 'f', 2, 64),
		))
	}

	if a.startScaleX != 0 {
		e.With(b.Style(
			"--fa-bounce-start-scale-x",
			strconv.FormatFloat(a.startScaleX, 'f', 2, 64),
		))
	}

	if a.startScaleY != 0 {
		e.With(b.Style(
			"--fa-bounce-start-scale-y",
			strconv.FormatFloat(a.startScaleY, 'f', 2, 64),
		))
	}

	if a.jumpScaleX != 0 {
		e.With(b.Style(
			"--fa-bounce-jump-scale-x",
			strconv.FormatFloat(a.jumpScaleX, 'f', 2, 64),
		))
	}

	if a.jumpScaleY != 0 {
		e.With(b.Style(
			"--fa-bounce-jump-scale-y",
			strconv.FormatFloat(a.jumpScaleY, 'f', 2, 64),
		))
	}

	if a.landScaleX != 0 {
		e.With(b.Style(
			"--fa-bounce-land-scale-x",
			strconv.FormatFloat(a.landScaleX, 'f', 2, 64),
		))
	}

	if a.landScaleY != 0 {
		e.With(b.Style(
			"--fa-bounce-land-scale-y",
			strconv.FormatFloat(a.landScaleY, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(e)
}

func Rebound(rebound float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.rebound = rebound
		}
	}
}

func Height(height float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.height = height
		}
	}
}

func StartScaleX(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.startScaleX = scale
		}
	}
}

func StartScaleY(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.startScaleY = scale
		}
	}
}

func JumpScaleX(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.jumpScaleX = scale
		}
	}
}

func JumpScaleY(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.jumpScaleY = scale
		}
	}
}

func LandScaleX(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.landScaleX = scale
		}
	}
}

func LandScaleY(scale float64) func(Animation) {
	return func(a Animation) {
		if b, ok := a.(*bounce); ok {
			b.landScaleY = scale
		}
	}
}
