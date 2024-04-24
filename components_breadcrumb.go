package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Breadcrumb creates a breadcrumb.
//
// https://willoma.github.io/bulma-gomponents/breadcrumb.html
func Breadcrumb(children ...any) Element {
	ul := Elem(html.Ul)
	b := &breadcrumb{
		Element: Elem(
			html.Nav,
			Class("breadcrumb"),
			html.Aria("label", "breadcrumbs"),
			elemOptionSpanAroundNonIconsIfHasIcons,
			ul,
		),
		ul: ul,
	}
	b.With(children...)
	return b
}

type breadcrumb struct {
	Element
	ul Element
}

func (b *breadcrumb) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onUl:
			b.ul.With(c...)
		case onNav:
			b.Element.With(c...)
		case Class, Classer, Classeser, ExternalClassesAndStyles, Styles:
			b.Element.With(c)
		case gomponents.Node:
			if isAttribute(c) {
				b.Element.With(c)
			} else {
				b.ul.With(c)
			}
		case []any:
			b.With(c...)
		default:
			b.ul.With(c)
		}
	}
	return b
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
