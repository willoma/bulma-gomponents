package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Breadcrumb creates a breadcrumb.
//
// https://willoma.github.io/bulma-gomponents/breadcrumb.html
func Breadcrumb(children ...any) e.Element {
	ul := e.Ul()
	b := &breadcrumb{
		Element: e.Nav(
			e.Class("breadcrumb"),
			e.AriaLabel("breadcrumbs"),
			ul,
		),
		ul: ul,
	}
	b.With(children...)
	return b
}

type breadcrumb struct {
	e.Element
	ul e.Element
}

func (b *breadcrumb) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onUl:
			b.ul.With(c...)
		case onNav:
			b.Element.With(c...)
		case e.Class, e.Classer, e.Classeser, e.Styles:
			b.Element.With(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (b *breadcrumb) Clone() e.Element {
	return &breadcrumb{
		Element: b.Element.Clone(),
		ul:      b.ul.Clone(),
	}
}

// BreadcrumbEntry creates a generic breadcrumb entry.
func BreadcrumbEntry(children ...any) e.Element {
	return e.Li(children...)
}

// BreadcrumbAHref creates a breadcrumb entry which contains
// a link to the provided URL.
//
// It is better than BreadcrumbEntry(AHref(href, children...)),
// because it ensures text is enclosed in span if a child is an icon.
func BreadcrumbAHref(href string, children ...any) e.Element {
	ahref := &spanAroundNonIconsIfHasIcons{elemFn: html.A}
	ahref.With(html.Href(href)).With(children...)

	return e.Li(ahref)
}

type breadcrumbAHref struct {
	e.Element
}

func (b *breadcrumbAHref) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class:
			if c == Active {
				b.Element.With(e.AriaCurrentPage)
			}
			b.With(c)
		case []any:
			b.With(c...)
		default:
			b.Element.With(c)
		}
	}

	return b
}
