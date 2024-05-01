package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type flip struct {
	animationBase

	x     float64
	y     float64
	z     float64
	angle float64
}

func Flip(options ...func(Animation)) Animation {
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
	}

	if a.y != 0 {
		p.With(e.Style(
			"--fa-flip-y",
			strconv.FormatFloat(a.y, 'f', 2, 64),
		))
	}

	if a.z != 0 {
		p.With(e.Style(
			"--fa-flip-z",
			strconv.FormatFloat(a.z, 'f', 2, 64),
		))
	}

	if a.angle != 0 {
		p.With(e.Style(
			"--fa-flip-angle",
			strconv.FormatFloat(a.angle, 'f', 2, 64),
		))
	}

	a.animationBase.ModifyParent(p)
}

func X(x float64) func(Animation) {
	return func(a Animation) {
		if f, ok := a.(*flip); ok {
			f.x = x
		}
	}
}

func Y(y float64) func(Animation) {
	return func(a Animation) {
		if f, ok := a.(*flip); ok {
			f.y = y
		}
	}
}

func Z(z float64) func(Animation) {
	return func(a Animation) {
		if f, ok := a.(*flip); ok {
			f.z = z
		}
	}
}

func Angle(angle float64) func(Animation) {
	return func(a Animation) {
		if f, ok := a.(*flip); ok {
			f.angle = angle
		}
	}
}
