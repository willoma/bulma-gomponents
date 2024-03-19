package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Breadcrumb creates a breadcrumb.
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
func Breadcrumb(children ...any) *Element {
	b := &breadcrumb{}
	b.addChildren(children)
	return b.elem()
}

type breadcrumb struct {
	ulChildren  []any
	navChildren []any
}

func (b *breadcrumb) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			b.navChildren = append(b.navChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				b.navChildren = append(b.navChildren, c)
			} else {
				b.ulChildren = append(b.ulChildren, c)
			}
		case []any:
			b.addChildren(c)
		default:
			b.ulChildren = append(b.ulChildren, c)
		}
	}
}

func (b *breadcrumb) elem() *Element {
	nav := Elem(html.Nav).
		With(Class("breadcrumb")).
		With(html.Aria("label", "breadcrumbs")).
		Withs(b.navChildren)

	nav.spanAroundNonIconsIfHasIcons = true

	return nav.With(Elem(html.Ul).Withs(b.ulChildren))
}

// BreadcrumbEntry creates a generic breadcrumb entry.
func BreadcrumbEntry(children ...any) *Element {
	return Elem(html.Li).Withs(children)
}

// BreadcrumbAHref creates a breadcrumb entry which contains
// a link to the provided URL.
//
// It is better than BreadcrumbEntry(AHref(href, children...)),
// because it ensures text is enclosed in span if a child is an icon.
func BreadcrumbAHref(href string, children ...any) *Element {
	ahref := AHref(href, children...)
	ahref.spanAroundNonIconsIfHasIcons = true

	return Elem(html.Li).With(ahref)
}

// BreadcrumbActiveAHref creates an active breadcrumb entry which contains
// a link to the provided URL.
//
// It is better than BreadcrumbEntry(Active, AHref(href, children)),
// because it ensures text is enclosed in span if a child is an icon.
func BreadcrumbActiveAHref(href string, children ...any) *Element {
	ahref := AHref(href, children...).
		With(html.Aria("current", "page"))

	ahref.spanAroundNonIconsIfHasIcons = true

	return Elem(html.Li).With(Active).With(ahref)
}
