package fa

import (
	b "github.com/willoma/bulma-gomponents"
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

func (a *shake) ModifyParent(e b.Element) {
	e.With(Class("fa-shake"))

	a.animationBase.ModifyParent(e)
}
