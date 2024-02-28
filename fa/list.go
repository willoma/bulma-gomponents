package fa

import (
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

// Li is a list element in a UList or a OList.
type Li struct {
	Icon  *b.Element
	Child any
}

// UList replaces bullets with icons in ul lists. When the Icon attribute is
// nil, the fa-minus icon is used. When not nil, the Icon attribute should be
// a return value of FA.
//
// If the Child element is of type []any, all its elements are added as children
// to the li element.
func UList(lines ...Li) *b.Element {
	elem := el.Ul(b.Class("fa-ul"))

	for _, li := range lines {
		liElem := el.Li()

		if li.Icon == nil {
			liElem.With(el.Span(b.Class("fa-li"), FA(Solid, "minus")))
		} else {
			liElem.With(el.Span(b.Class("fa-li"), li.Icon))
		}

		switch c := li.Child.(type) {
		case []any:
			liElem.Withs(c)
		default:
			liElem.With(c)
		}

		elem.With(liElem)
	}

	return elem
}

// OList replaces numbers with icons in ol lists. When the Icon attribute is
// nil, the fa-minus icon is used. When not nil, the Icon attribute should be
// a return value of FA.
//
// If the Child element is of type []any, all its elements are added as children
// to the li element.
func OList(lines ...Li) *b.Element {
	elem := el.Ol(b.Class("fa-ul"))

	for _, li := range lines {
		liElem := el.Li()

		if li.Icon == nil {
			liElem.With(el.Span(b.Class("fa-li"), FA(Solid, "minus")))
		} else {
			liElem.With(el.Span(b.Class("fa-li"), li.Icon))
		}

		switch c := li.Child.(type) {
		case []any:
			liElem.Withs(c)
		default:
			liElem.With(c)
		}

		elem.With(liElem)
	}

	return elem
}
