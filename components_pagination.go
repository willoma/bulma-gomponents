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
	pagination := Elem(html.Nav).
		With(Class("pagination"))

	list := Elem(html.Ul).
		With(Class("pagination-list"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("pagination-link") || c.hasClass("pagination-ellipsis") {
				list.With(Elem(html.Li).With(c))
			} else {
				pagination.With(c)
			}
		default:
			pagination.With(c)
		}
	}

	return pagination.With(list)
}

// PaginationPrevious creates the "Previous" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationPrevious(children ...any) *Element {
	return Elem(html.A).
		With(Class("pagination-previous")).
		Withs(children)
}

// PaginationNext creates the "Next" link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Disabled: mark the link button as inactive
func PaginationNext(children ...any) *Element {
	return Elem(html.A).
		With(Class("pagination-next")).
		Withs(children)
}

// PaginationLink creates a single page link button for a pagination.
//
// The following modifiers change the link button behaviour:
//   - Current: mark this link button as being the current page
//   - Disabled: mark the link button as inactive
func PaginationLink(children ...any) *Element {
	return Elem(html.A).
		With(Class("pagination-link")).
		Withs(children)
}

// PaginationEllipsis creates an ellipsis element for a pagination.
func PaginationEllipsis() *Element {
	return Elem(html.Span).
		With(Class("pagination-ellipsis")).
		With(gomponents.Raw("&hellip;"))
}
