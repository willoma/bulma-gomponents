package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type flip struct {
	animationBase

	x     float64
	zeroX bool

	y     float64
	zeroY bool

	z     float64
	zeroZ bool

	angle     float64
	zeroAngle bool
}

func Flip(options ...func(any)) e.ParentModifier {
	a := &flip{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *flip) ModifyParent(p e.Element) {
	p.With(Class("fa-flip"))

	if a.x != 0 {
		p.With(e.Style(
			"--fa-flip-x",
			strconv.FormatFloat(a.x, 'f', 2, 64),
		))
	} else if a.zeroX {
		p.With(e.Style("--fa-flip-x", "0"))
	}

	if a.y != 0 {
		p.With(e.Style(
			"--fa-flip-y",
			strconv.FormatFloat(a.y, 'f', 2, 64),
		))
	} else if a.zeroY {
		p.With(e.Style("--fa-flip-y", "0"))
	}

	if a.z != 0 {
		p.With(e.Style(
			"--fa-flip-z",
			strconv.FormatFloat(a.z, 'f', 2, 64),
		))
	} else if a.zeroZ {
		p.With(e.Style("--fa-flip-z", "0"))
	}

	if a.angle != 0 {
		p.With(e.Style(
			"--fa-flip-angle",
			strconv.FormatFloat(a.angle, 'f', 2, 64),
		))
	} else if a.zeroAngle {
		p.With(e.Style("--fa-flip-angle", "0"))
	}

	a.animationBase.modifyParent(p)
}

func (a *flip) If(cond bool) e.ParentModifier {
	return &conditionalAnimation{animation: a, cond: cond}
}

func X(x float64) func(any) {
	return func(a any) {
		if f, ok := a.(*flip); ok {
			f.zeroX = x == 0
			f.x = x
		}
	}
}

func Y(y float64) func(any) {
	return func(a any) {
		if f, ok := a.(*flip); ok {
			f.zeroY = y == 0
			f.y = y
		}
	}
}

func Z(z float64) func(any) {
	return func(a any) {
		if f, ok := a.(*flip); ok {
			f.zeroZ = z == 0
			f.z = z
		}
	}
}

func Angle(angle float64) func(any) {
	return func(a any) {
		if f, ok := a.(*flip); ok {
			f.zeroAngle = angle == 0
			f.angle = angle
		}
	}
}
