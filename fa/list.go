package fa

import (
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

// Li is a list element in a UList or a OList.
type Li struct {
	Icon  b.Element
	Child any
}

// UList replaces bullets with icons in ul lists. When the Icon attribute is
// nil, the fa-minus icon is used. When not nil, the Icon attribute should be
// a return value of FA.
//
// If the Child element is of type []any, all its elements are added as children
// to the li element.
func UList(lines ...Li) b.Element {
	var children []any

	for _, li := range lines {
		var icon b.Element

		if li.Icon == nil {
			icon = FA(Solid, "minus")
		} else {
			icon = li.Icon
		}

		children = append(
			children,
			el.Li(el.Span(b.Class("fa-li"), icon), li.Child),
		)
	}

	return el.Ul(b.Class("fa-ul"), children)
}

// OList replaces numbers with icons in ol lists. When the Icon attribute is
// nil, the fa-minus icon is used. When not nil, the Icon attribute should be
// a return value of FA.
func OList(lines ...Li) b.Element {
	var children []any

	for _, li := range lines {
		var icon b.Element

		if li.Icon == nil {
			icon = FA(Solid, "minus")
		} else {
			icon = li.Icon
		}

		children = append(
			children,
			el.Li(el.Span(b.Class("fa-li"), icon), li.Child),
		)
	}

	return el.Ol(b.Class("fa-ul"), children)
}
