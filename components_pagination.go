package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Pagination creates a pagination element.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func Pagination(children ...any) Element {
	list := Elem(html.Ul, Class("pagination-list"))
	p := &pagination{
		Element: Elem(html.Nav, Class("pagination"), list),
		list:    list,
	}
	p.With(children...)
	return p
}

type pagination struct {
	Element
	list Element
}

func (p *pagination) With(children ...any) Element {
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

func (p *pagination) Clone() Element {
	return &pagination{
		Element: p.Element.Clone(),
		list:    p.list.Clone(),
	}
}

// PaginationPrevious creates the "Previous" link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationPrevious(children ...any) Element {
	return Elem(html.A, Class("pagination-previous"), children)
}

// PaginationNext creates the "Next" link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationNext(children ...any) Element {
	return Elem(html.A, Class("pagination-next"), children)
}

// PaginationLink creates a single page link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationLink(children ...any) Element {
	p := &paginationLink{Elem(html.A, Class("pagination-link"))}
	p.With(children...)
	return p
}

type paginationLink struct {
	Element
}

func (p *paginationLink) Clone() Element {
	return &paginationLink{p.Element.Clone()}
}

// PaginationAHref creates a single page link button for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationAHref(href string, children ...any) Element {
	return PaginationLink(html.Href(href), children)
}

// PaginationEllipsis creates an ellipsis element for a pagination.
//
// https://willoma.github.io/bulma-gomponents/pagination.html
func PaginationEllipsis(children ...any) Element {
	p := &paginationEllipsis{
		Elem(html.Span, Class("pagination-ellipsis"), gomponents.Raw("&hellip;")),
	}
	p.With(children...)
	return p
}

type paginationEllipsis struct {
	Element
}

func (p *paginationEllipsis) Clone() Element {
	return &paginationEllipsis{p.Element.Clone()}
}
