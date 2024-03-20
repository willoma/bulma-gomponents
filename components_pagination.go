package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Pagination creates a pagination element.
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
func Pagination(children ...any) *Element {
	p := pagination{}
	p.addChildren(children)
	return p.elem()
}

type pagination struct {
	listChildren       []any
	paginationChildren []any
}

func (p *pagination) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("pagination-link") || c.hasClass("pagination-ellipsis") {
				p.listChildren = append(p.listChildren, Elem(html.Li, c))
			} else {
				p.paginationChildren = append(p.paginationChildren, c)
			}
		case []any:
			p.addChildren(c)
		default:
			p.paginationChildren = append(p.paginationChildren, c)
		}
	}
}

func (p *pagination) elem() *Element {
	return Elem(
		html.Nav,
		Class("pagination"),
		p.paginationChildren,
		Elem(html.Ul, Class("pagination-list"), p.listChildren),
	)
}

// PaginationPrevious creates the "Previous" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationPrevious(children ...any) *Element {
	return Elem(html.A, Class("pagination-previous"), children)
}

// PaginationNext creates the "Next" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationNext(children ...any) *Element {
	return Elem(html.A, Class("pagination-next"), children)
}

// PaginationLink creates a single page link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Current: mark this link button as being the current page
//   - Disabled: mark the link button as inactive
func PaginationLink(children ...any) *Element {
	return Elem(html.A, Class("pagination-link"), children)
}

// PaginationEllipsis creates an ellipsis element for a pagination.
func PaginationEllipsis() *Element {
	return Elem(html.Span, Class("pagination-ellipsis"), gomponents.Raw("&hellip;"))
}
