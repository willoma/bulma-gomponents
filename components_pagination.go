package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var pagination = c.NewPage(
	"Pagination", "Pagination", "/pagination",
	"",
	b.Content(
		e.P("The ", e.Code("b.Pagination"), " constructor creates a pagination. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnList(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<ul class="pagination-list">`), " e.Element"},

			e.Code("b.PaginationLink(...)"),
			"Add a button-looking link to the pagination list",

			e.Code("b.PaginationEllipsis()"),
			"Add an ellipsis to the pagination list",

			e.Code("b.Centered"),
			`Center the pagination list ("Previous" button on the left side, "Next" on the right side)`,

			e.Code("b.Right"),
			`Align the pagination list to the right ("Previous" and "Next" buttons on the left side)`,

			e.Code("b.Rounded"),
			"Have rounded pagination buttons",

			e.Code("b.Small"),
			"Set pagination size to small",

			e.Code("b.Medium"),
			"Set pagination size to medium",

			e.Code("b.Large"),
			"Set pagination size to large",
		),
		e.P("Other children are added to the ", e.Code(`<div class="pagination">`), " e.Element."),
		e.P("The ", e.Code("b.PaginationLink"), " and ", e.Code("b.PaginationAHref"), " constructors create button-looking links for the pagination list section. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Current"),
			"Mark this link as being the current page",

			e.Code("b.Disabled"),
			"Mark this link as inactive",
		),
		e.P("The ", e.Code("b.PaginationEllipsis"), " constructor creates an ellipsis for the pagination list section."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/pagination/",
	c.HorizontalExample(
		`b.Pagination(
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.PaginationPrevious(b.Disabled, "Previous", html.TitleAttr("This is the first page")),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Page 1"), b.Current, "1"),
	b.PaginationLink(e.AriaLabel("Goto page 2"), "2"),
	b.PaginationLink(e.AriaLabel("Goto page 3"), "3"),
)`,
		b.Pagination(
			b.PaginationPrevious(b.Disabled, "Previous", html.TitleAttr("This is the first page")),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Page 1"), b.Current, "1"),
			b.PaginationLink(e.AriaLabel("Goto page 2"), "2"),
			b.PaginationLink(e.AriaLabel("Goto page 3"), "3"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Centered,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Centered,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Right,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Right,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
).Subsection(
	"Styles",
	"https://bulma.io/documentation/components/pagination/#styles",
	c.HorizontalExample(
		`b.Pagination(
	b.Rounded,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Rounded,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/components/pagination/#sizes",
	c.HorizontalExample(
		`b.Pagination(
	b.Small,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Small,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Medium,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Medium,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Large,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
	b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
	b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Large,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(e.AriaLabel("Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 45"), "45"),
			b.PaginationLink(e.AriaLabel("Goto page 46"), b.Current, "46"),
			b.PaginationLink(e.AriaLabel("Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(e.AriaLabel("Goto page 86"), "86"),
		),
	),
)
