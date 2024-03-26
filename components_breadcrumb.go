package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Breadcrumb creates a breadcrumb.
//   - when a child is marked with b.Inner, it is forcibly applied to the <ul> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <nav> element
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
func Breadcrumb(children ...any) Element {
	return new(breadcrumb).With(children...)
}

type breadcrumb struct {
	ulChildren  []any
	navChildren []any
}

func (b *breadcrumb) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *ApplyToInner:
			b.ulChildren = append(b.ulChildren, c.Child)
		case *ApplyToOuter:
			b.navChildren = append(b.navChildren, c.Child)
		case Class:
			b.navChildren = append(b.navChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				b.navChildren = append(b.navChildren, c)
			} else {
				b.ulChildren = append(b.ulChildren, c)
			}
		case []any:
			b.With(c...)
		default:
			b.ulChildren = append(b.ulChildren, c)
		}
	}
	return b
}

func (b *breadcrumb) Render(w io.Writer) error {
	return Elem(
		html.Nav,
		Class("breadcrumb"),
		html.Aria("label", "breadcrumbs"),
		b.navChildren,
		elemOptionSpanAroundNonIconsIfHasIcons,
		Elem(html.Ul, b.ulChildren...),
	).Render(w)
}

// BreadcrumbEntry creates a generic breadcrumb entry.
func BreadcrumbEntry(children ...any) Element {
	return Elem(html.Li, children...)
}

// BreadcrumbAHref creates a breadcrumb entry which contains
// a link to the provided URL.
//
// It is better than BreadcrumbEntry(AHref(href, children...)),
// because it ensures text is enclosed in span if a child is an icon.
func BreadcrumbAHref(href string, children ...any) Element {
	return Elem(
		html.Li,
		AHref(
			href,
			elemOptionSpanAroundNonIconsIfHasIcons,
			children,
		),
	)
}

// BreadcrumbActiveAHref creates an active breadcrumb entry which contains
// a link to the provided URL.
//
// It is better than BreadcrumbEntry(Active, AHref(href, children)),
// because it ensures text is enclosed in span if a child is an icon.
func BreadcrumbActiveAHref(href string, children ...any) Element {
	return Elem(
		html.Li, Active,
		AHref(
			href,
			elemOptionSpanAroundNonIconsIfHasIcons,
			children,
			html.Aria("current", "page"),
		),
	)
}
