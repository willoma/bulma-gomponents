package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Pagination creates a pagination element.
//   - when a child is marked with b.OnList, it is forcibly applied to the <ul class="pagination-list"> element
//
// The results of PaginationLink and PaginationEllipsis are automatically added
// in the auto-created "pagination-list" element.
//
// The following modifiers change the pagination behaviour:
//   - Centered: center the pagination list ("Previous" on left, "Next" on right)
//   - Right: align the pagination list to the right ("Previous" and "Next" on left)
//   - Rounded: have rounded pagination buttons
//
// The following modifiers change the pagination size:
//   - Small
//   - Medium
//   - Large
func Pagination(children ...any) Element {
	return new(pagination).With(children...)
}

type pagination struct {
	listChildren       []any
	paginationChildren []any
}

func (p *pagination) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onList:
			p.listChildren = append(p.listChildren, c)
		case *paginationEllipsis, *paginationLink:
			p.listChildren = append(p.listChildren, Elem(html.Li, c))
		case []any:
			p.With(c...)
		default:
			p.paginationChildren = append(p.paginationChildren, c)
		}
	}

	return p
}

func (p *pagination) Render(w io.Writer) error {
	return Elem(
		html.Nav,
		Class("pagination"),
		p.paginationChildren,
		Elem(html.Ul, Class("pagination-list"), p.listChildren),
	).Render(w)
}

// PaginationPrevious creates the "Previous" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationPrevious(children ...any) Element {
	return Elem(html.A, Class("pagination-previous"), children)
}

// PaginationNext creates the "Next" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationNext(children ...any) Element {
	return Elem(html.A, Class("pagination-next"), children)
}

// PaginationLink creates a single page link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Current: mark this link button as being the current page
//   - Disabled: mark the link button as inactive
func PaginationLink(children ...any) Element {
	return new(paginationLink).With(children...)
}

type paginationLink struct {
	children []any
}

func (p *paginationLink) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			p.With(c...)
		default:
			p.children = append(p.children, c)
		}
	}

	return p
}

func (p *paginationLink) Render(w io.Writer) error {
	return Elem(html.A, Class("pagination-link"), p.children).Render(w)
}

// PaginationAHref creates a single page link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Current: mark this link button as being the current page
//   - Disabled: mark the link button as inactive
func PaginationAHref(href string, children ...any) Element {
	return new(paginationLink).With(html.Href(href), children)
}

// PaginationEllipsis creates an ellipsis element for a pagination.
func PaginationEllipsis(children ...any) Element {
	return new(paginationEllipsis).With(children...)
}

type paginationEllipsis struct {
	children []any
}

func (p *paginationEllipsis) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			p.With(c...)
		default:
			p.children = append(p.children, c)
		}
	}

	return p
}

func (p *paginationEllipsis) Render(w io.Writer) error {
	return Elem(html.Span, Class("pagination-ellipsis"), gomponents.Raw("&hellip;"), p.children).Render(w)
}
