package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type rotateOrFlip string

func (r rotateOrFlip) Class() e.Class {
	return e.Class(r)
}

func (r rotateOrFlip) If(cond bool) e.ParentModifier {
	return &conditionalRotateOrFlip{rotateOrFlip: r, cond: cond}
}

type conditionalRotateOrFlip struct {
	rotateOrFlip rotateOrFlip
	cond         bool
}

func (r *conditionalRotateOrFlip) ModifyParent(p e.Element) {
	if r.cond {
		p.With(r.rotateOrFlip)
	}
}

// Rotation (see https://fontawesome.com/docs/web/style/rotate)
const (
	Rotate90       = rotateOrFlip("fa-rotate-90")
	Rotate180      = rotateOrFlip("fa-rotate-180")
	Rotate270      = rotateOrFlip("fa-rotate-270")
	FlipHorizontal = rotateOrFlip("fa-flip-horizontal")
	FlipVertical   = rotateOrFlip("fa-flip-vertical")
	FlipBoth       = rotateOrFlip("fa-flip-both")
)

type Rotate float64

func (r Rotate) ModifyParent(parent e.Element) {
	parent.With(e.Class("fa-rotate-by"))
	parent.With(e.Style(
		"--fa-rotate-angle", strconv.FormatFloat(float64(r), 'f', -1, 64)+"deg",
	))
}

func (r Rotate) If(cond bool) e.ParentModifier {
	return &conditionalRotate{rotate: r, cond: cond}
}

type conditionalRotate struct {
	rotate Rotate
	cond   bool
}

func (r *conditionalRotate) ModifyParent(p e.Element) {
	if r.cond {
		p.With(r.rotate)
	}
}
