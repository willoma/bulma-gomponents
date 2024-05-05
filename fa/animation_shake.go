package fa

import (
	e "github.com/willoma/gomplements"
)

type shake struct {
	animation
}

func Shake(options ...func(any)) e.ParentModifier {
	a := &shake{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *shake) ModifyParent(p e.Element) {
	p.With(Class("fa-shake"))

	a.animation.ModifyParent(p)
}
