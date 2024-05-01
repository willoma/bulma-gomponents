package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var breadcrumb = c.NewPage(
	"Breadcrumb", "Breadcrumb", "/breadcrumb",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Breadcrumb"), " constructor creates a breadcrumb. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnNav(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<nav>"), " element"},

			e.Code("b.OnUl(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<ul>"), " element"},

			[]any{"one of the class or style types defined in package ", e.Code("b")},
			[]any{"Apply the class or style to the ", e.Code(`<nav class="breadcrumb">`), " element"},

			e.Code("b.Centered"),
			"Center the breadcrumb in its container",

			e.Code("b.Right"),
			"Align the breadcrumb to the right",

			e.Code("b.ArrowSeparator"),
			[]any{"Use an arrow (", e.Code("→"), ") as the separator"},

			e.Code("b.BulletSeparator"),
			[]any{"Use a bullet (", e.Code("•"), ") as the separator"},

			e.Code("b.DotSeparator"),
			[]any{"Use a dot (", e.Code("·"), ") as the separator"},

			e.Code("b.SucceedsSeparator"),
			[]any{`Use a "succeeds" character (`, e.Code("≻"), `) as the separator`},

			e.Code("b.Small"),
			"Set breadcrumb size to small",

			e.Code("b.Medium"),
			"Set breadcrumb size to medium",

			e.Code("b.Large"),
			"Set breadcrumb size to large",
		),
		e.P("Other children are added to the ", e.Code("<ul>"), " element."),
		e.P(
			"The ", e.Code("b.BreadcrumbEntry"), " constructor creates a breadcrumb entry (which is a regular <li> e.Element). The ", e.Code("b.BreadcrumbAHref"), " constructor creates a breadcrumb link entry. The ", e.Code("b.BreadcrumbActiveAHref"), " constructor creates an active breadcrumb link entry.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/breadcrumb/",
	c.Example(
		`b.Breadcrumb(
	b.BreadcrumbEntry(
		e.AHref("#", "Bulma"),
	),
	b.BreadcrumbEntry(
		e.AHref("#", "Documentation"),
	),
	b.BreadcrumbEntry(
		e.AHref("#", "Components"),
	),
	b.BreadcrumbEntry(
		b.Active,
		e.AHref("#", e.AriaCurrentPage, "Breadcrumb"),
	),
)`,
		b.Breadcrumb(
			b.BreadcrumbEntry(
				e.AHref("#", "Bulma"),
			),
			b.BreadcrumbEntry(
				e.AHref("#", "Documentation"),
			),
			b.BreadcrumbEntry(
				e.AHref("#", "Components"),
			),
			b.BreadcrumbEntry(
				b.Active,
				e.AHref("#", e.AriaCurrentPage, "Breadcrumb"),
			),
		),
	),
	b.Content(e.P(e.Em("Bulma-Gomponents"), " provides the ", e.Code("b.BreadcrumbHref"), " and ", e.Code("b.BreadcrumbActiveAHref"), "helpers for regular breadcrumb links:")),
	c.Example(
		`b.Breadcrumb(
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
).Subsection(
	"Alignment",
	"https://bulma.io/documentation/components/breadcrumb/#alignment",
	c.Example(
		`b.Breadcrumb(
	b.Centered,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Centered,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Right,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Right,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
).Subsection(
	"Icons",
	"https://bulma.io/documentation/components/breadcrumb/#icons",
	c.Example(
		`b.Breadcrumb(
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "home"), "Bulma"),
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "book"), "Documentation"),
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "puzzle-piece"), "Components"),
	b.BreadcrumbActiveAHref("#", fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "home"), "Bulma"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "book"), "Documentation"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "puzzle-piece"), "Components"),
			b.BreadcrumbActiveAHref("#", fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
		),
	),
).Subsection(
	"Alternative separators",
	"https://bulma.io/documentation/components/breadcrumb/#alternative-separators",
	c.Example(
		`easy.Breadcrumb(
	b.ArrowSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.ArrowSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.BulletSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
`,
		b.Breadcrumb(
			b.BulletSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.DotSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.DotSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.SucceedsSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.SucceedsSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/components/breadcrumb/#sizes",
	c.Example(
		`b.Breadcrumb(
	b.Small,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Small,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Medium,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Medium,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Large,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Large,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
	),
)
