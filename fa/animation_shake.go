package fa

import (
	e "github.com/willoma/gomplements"
)

type shake struct {
	animationBase
}

func Shake(options ...func(Animation)) Animation {
	a := &shake{}

	for _, o := range options {
		o(a)
	}

	return a
}

func (a *shake) ModifyParent(p e.Element) {
	p.With(Class("fa-shake"))

	a.animationBase.ModifyParent(p)
}
