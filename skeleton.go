package bulma

import (
	e "github.com/willoma/gomplements"
)

// SkeletonBlock creates a block skeleton.
func SkeletonBlock(children ...any) e.Element {
	return e.Div(e.Class("skeleton-block"), children)
}

// SkeletonLines create a block of skeleton lines.
func SkeletonLines(count int) e.Element {
	elem := e.Div(e.Class("skeleton-lines"))

	for i := 0; i < count; i++ {
		elem.With(e.Div())
	}

	return elem
}
