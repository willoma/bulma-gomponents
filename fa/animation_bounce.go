package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type bounce struct {
	animationBase

	rebound     float64
	zeroRebound bool

	height     float64
	zeroHeight bool

	startScaleX     float64
	zeroStartScaleX bool

	startScaleY     float64
	zeroStartScaleY bool

	jumpScaleX     float64
	zeroJumpScaleX bool

	jumpScaleY     float64
	zeroJumpScaleY bool

	landScaleX     float64
	zeroLandScaleX bool

	landScaleY     float64
	zeroLandScaleY bool
}

func Bounce(options ...func(any)) e.ParentModifier {
	a := &bounce{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *bounce) ModifyParent(p e.Element) {
	p.With(Class("fa-bounce"))

	if a.rebound != 0 {
		p.With(e.Style(
			"--fa-bounce-rebound",
			strconv.FormatFloat(a.rebound, 'f', 2, 64),
		))
	} else if a.zeroRebound {
		p.With(e.Style("--fa-bounce-rebound", "0"))
	}

	if a.height != 0 {
		p.With(e.Style(
			"--fa-bounce-height",
			strconv.FormatFloat(a.height, 'f', 2, 64),
		))
	} else if a.zeroHeight {
		p.With(e.Style("--fa-bounce-height", "0"))
	}

	if a.startScaleX != 0 {
		p.With(e.Style(
			"--fa-bounce-start-scale-x",
			strconv.FormatFloat(a.startScaleX, 'f', 2, 64),
		))
	} else if a.zeroStartScaleX {
		p.With(e.Style("--fa-bounce-start-scale-x", "0"))
	}

	if a.startScaleY != 0 {
		p.With(e.Style(
			"--fa-bounce-start-scale-y",
			strconv.FormatFloat(a.startScaleY, 'f', 2, 64),
		))
	} else if a.zeroStartScaleY {
		p.With(e.Style("--fa-bounce-start-scale-y", "0"))
	}

	if a.jumpScaleX != 0 {
		p.With(e.Style(
			"--fa-bounce-jump-scale-x",
			strconv.FormatFloat(a.jumpScaleX, 'f', 2, 64),
		))
	} else if a.zeroJumpScaleX {
		p.With(e.Style("--fa-bounce-jump-scale-x", "0"))
	}

	if a.jumpScaleY != 0 {
		p.With(e.Style(
			"--fa-bounce-jump-scale-y",
			strconv.FormatFloat(a.jumpScaleY, 'f', 2, 64),
		))
	} else if a.zeroJumpScaleY {
		p.With(e.Style("--fa-bounce-jump-scale-y", "0"))
	}

	if a.landScaleX != 0 {
		p.With(e.Style(
			"--fa-bounce-land-scale-x",
			strconv.FormatFloat(a.landScaleX, 'f', 2, 64),
		))
	} else if a.zeroLandScaleX {
		p.With(e.Style("--fa-bounce-land-scale-x", "0"))
	}

	if a.landScaleY != 0 {
		p.With(e.Style(
			"--fa-bounce-land-scale-y",
			strconv.FormatFloat(a.landScaleY, 'f', 2, 64),
		))
	} else if a.zeroLandScaleY {
		p.With(e.Style("--fa-bounce-land-scale-y", "0"))
	}

	a.animationBase.modifyParent(p)
}

func (a *bounce) If(cond bool) e.ParentModifier {
	return &conditionalAnimation{animation: a, cond: cond}
}

func Rebound(rebound float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroRebound = rebound == 0
			b.rebound = rebound
		}
	}
}

func Height(height float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroHeight = height == 0
			b.height = height
		}
	}
}

func StartScaleX(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroStartScaleX = scale == 0
			b.startScaleX = scale
		}
	}
}

func StartScaleY(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroStartScaleY = scale == 0
			b.startScaleY = scale
		}
	}
}

func JumpScaleX(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroJumpScaleX = scale == 0
			b.jumpScaleX = scale
		}
	}
}

func JumpScaleY(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroJumpScaleY = scale == 0
			b.jumpScaleY = scale
		}
	}
}

func LandScaleX(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroLandScaleX = scale == 0
			b.landScaleX = scale
		}
	}
}

func LandScaleY(scale float64) func(any) {
	return func(a any) {
		if b, ok := a.(*bounce); ok {
			b.zeroLandScaleY = scale == 0
			b.landScaleY = scale
		}
	}
}
