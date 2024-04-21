package bulma

import "github.com/maragudk/gomponents/html"

// SkeletonBlock creates a block skeleton.
func SkeletonBlock(children ...any) Element {
	return Elem(html.Div, Class("skeleton-block"), children)
}

// SkeletonLines create a block of skeleton lines.
func SkeletonLines(count int) Element {
	elem := Elem(html.Div, Class("skeleton-lines"))

	for i := 0; i < count; i++ {
		elem.With(Elem(html.Div))
	}

	return elem
}
