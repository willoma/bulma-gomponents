package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var pagination = c.NewPage(
	"Pagination", "Pagination", "/pagination",
	"",
	b.Content(
		el.P("The ", el.Code("b.Pagination"), " constructor creates a pagination. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnList(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<ul class="pagination-list">`), " element"},

			el.Code("b.PaginationLink(...)"),
			"Add a button-looking link to the pagination list",

			el.Code("b.PaginationEllipsis()"),
			"Add an ellipsis to the pagination list",

			el.Code("b.Centered"),
			`Center the pagination list ("Previous" button on the left side, "Next" on the right side)`,

			el.Code("b.Right"),
			`Align the pagination list to the right ("Previous" and "Next" buttons on the left side)`,

			el.Code("b.Rounded"),
			"Have rounded pagination buttons",

			el.Code("b.Small"),
			"Set pagination size to small",

			el.Code("b.Medium"),
			"Set pagination size to medium",

			el.Code("b.Large"),
			"Set pagination size to large",
		),
		el.P("Other children are added to the ", el.Code(`<div class="pagination">`), " element."),
		el.P("The ", el.Code("b.PaginationLink"), " and ", el.Code("b.PaginationAHref"), " constructors create button-looking links for the pagination list section. The following children have a special meaning:"),
		b.DList(
			el.Code("b.Current"),
			"Mark this link as being the current page",

			el.Code("b.Disabled"),
			"Mark this link as inactive",
		),
		el.P("The ", el.Code("b.PaginationEllipsis"), " constructor creates an ellipsis for the pagination list section."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/pagination/",
	c.HorizontalExample(
		`b.Pagination(
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.PaginationPrevious(b.Disabled, "Previous", html.TitleAttr("This is the first page")),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Page 1"), b.Current, "1"),
	b.PaginationLink(html.Aria("label", "Goto page 2"), "2"),
	b.PaginationLink(html.Aria("label", "Goto page 3"), "3"),
)`,
		b.Pagination(
			b.PaginationPrevious(b.Disabled, "Previous", html.TitleAttr("This is the first page")),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Page 1"), b.Current, "1"),
			b.PaginationLink(html.Aria("label", "Goto page 2"), "2"),
			b.PaginationLink(html.Aria("label", "Goto page 3"), "3"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Centered,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Centered,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Right,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Right,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
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
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Rounded,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
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
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Small,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Medium,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Medium,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
		),
	),
	c.HorizontalExample(
		`b.Pagination(
	b.Large,
	b.PaginationPrevious("Previous"),
	b.PaginationNext("Next page"),
	b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
	b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
	b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
	b.PaginationEllipsis(),
	b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
)`,
		b.Pagination(
			b.Large,
			b.PaginationPrevious("Previous"),
			b.PaginationNext("Next page"),
			b.PaginationLink(html.Aria("label", "Goto page 1"), "1"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 45"), "45"),
			b.PaginationLink(html.Aria("label", "Goto page 46"), b.Current, "46"),
			b.PaginationLink(html.Aria("label", "Goto page 47"), "47"),
			b.PaginationEllipsis(),
			b.PaginationLink(html.Aria("label", "Goto page 86"), "86"),
		),
	),
)
