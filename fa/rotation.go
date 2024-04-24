package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
)

type rotateOrFlip string

func (r rotateOrFlip) Class() b.Class {
	return b.Class(r)
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

func (r Rotate) ModifyParent(parent b.Element) {
	parent.With(b.Class("fa-rotate-by"))
	parent.With(b.Style(
		"--fa-rotate-angle", strconv.FormatFloat(float64(r), 'f', -1, 64)+"deg",
	))
}
