package bulma

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// Pagination creates a pagination element.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func Pagination(children ...any) e.Element {
	list := e.Ul(e.Class("pagination-list"))
	p := &pagination{
		Element: e.Nav(e.Class("pagination"), list),
		list:    list,
	}
	p.With(children...)
	return p
}

type pagination struct {
	e.Element
	list e.Element
}

func (p *pagination) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onList:
			p.list.With(c...)
		case *paginationEllipsis, *paginationLink:
			p.list.With(c)
		case []any:
			p.With(c...)
		default:
			p.Element.With(c)
		}
	}

	return p
}

func (p *pagination) Clone() e.Element {
	return &pagination{
		Element: p.Element.Clone(),
		list:    p.list.Clone(),
	}
}

// PaginationPrevious creates the "Previous" link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationPrevious(children ...any) e.Element {
	return e.A(e.Class("pagination-previous"), children)
}

// PaginationNext creates the "Next" link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationNext(children ...any) e.Element {
	return e.A(e.Class("pagination-next"), children)
}

// PaginationLink creates a single page link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationLink(children ...any) e.Element {
	p := &paginationLink{e.A(e.Class("pagination-link"))}
	p.With(children...)
	return p
}

type paginationLink struct {
	e.Element
}

func (p *paginationLink) Clone() e.Element {
	return &paginationLink{p.Element.Clone()}
}

// PaginationAHref creates a single page link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationAHref(href string, children ...any) e.Element {
	return PaginationLink(html.Href(href), children)
}

// PaginationEllipsis creates an ellipsis element for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationEllipsis(children ...any) e.Element {
	p := &paginationEllipsis{
		e.Span(e.Class("pagination-ellipsis"), gomponents.Raw("&hellip;")),
	}
	p.With(children...)
	return p
}

type paginationEllipsis struct {
	e.Element
}

func (p *paginationEllipsis) Clone() e.Element {
	return &paginationEllipsis{p.Element.Clone()}
}
