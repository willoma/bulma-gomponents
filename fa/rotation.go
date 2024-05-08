package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

type rotateOrFlip string

func (r rotateOrFlip) Class() e.Class {
	return e.Class(r)
}

func (r rotateOrFlip) If(cond bool) rotateOrFlip {
	if cond {
		return r
	}

	return ""
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
	if r == 0 {
		return
	}

	parent.With(e.Class("fa-rotate-by"))
	parent.With(e.Style(
		"--fa-rotate-angle", strconv.FormatFloat(float64(r), 'f', -1, 64)+"deg",
	))
}

func (r Rotate) If(cond bool) Rotate {
	if cond {
		return r
	}

	return 0
}
