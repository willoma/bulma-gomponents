package easy

import (
	"strings"

	"github.com/maragudk/gomponents"

	b "github.com/willoma/bulma-gomponents"
)

// Breadcrumb creates a breadcrumb with multiple links. It requires its children
// to be pairs of link targets / element content.
//
// The following modifiers change the breadcrumb behaviour:
//   - Centered: center the breadcrumb in its container
//   - Right: align the breadcrumb on the right
//
// The following modifiers change the separator:
//   - ArrowSeparator: use an arrow as the separator
//   - BulletSeparator: use a bullet as the separator
//   - DotSeparator: use a small dot as the separator
//   - SucceedsSeparator: use a "succeed" character as the separator
//
// The following modifiers change the breadcrumb size:
//   - Small
//   - Medium
//   - Large
//
// Targets must be strings, which are used as the href attribute of the a
// element. If the target starts with the "!" character, it makes it the active
// element.
//
// Content may be either a string, a single Element, a single gomponents.Node, or a []any.
//
// For instance:
//
//	easy.Breadcrumb(
//		b.ArrowSeparator,
//		"/", []any{fa.Icon(fa.Solid, "home"), "Root"},
//		"/profile", "Profile",
//	)
func Breadcrumb(children ...any) b.Element {
	newChildren := []any{}

	currentHref := ""

	addLink := func(children ...any) {
		if strings.HasPrefix(currentHref, "!") {
			newChildren = append(
				newChildren,
				b.BreadcrumbActiveAHref(currentHref[1:], children...),
			)
		} else {
			newChildren = append(
				newChildren,
				b.BreadcrumbAHref(currentHref, children...),
			)
		}
		currentHref = ""
	}

	for _, c := range children {
		switch c := c.(type) {
		case b.Class, b.ColorClass, b.ExternalClass, b.ExternalClassesAndStyles, b.MultiClass, b.Styles:
			newChildren = append(newChildren, c)

		case string:
			if currentHref == "" {
				currentHref = c
			} else {
				addLink(c)
			}

		case gomponents.Node:
			if b.IsAttribute(c) {
				newChildren = append(newChildren, c)
			} else {
				addLink(c)
			}

		case []any:
			addLink(c...)

		default:
			addLink(c)
		}
	}

	return b.Breadcrumb(newChildren...)
}
