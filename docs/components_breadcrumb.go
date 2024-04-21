package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var breadcrumb = c.NewPage(
	"Breadcrumb", "Breadcrumb", "/breadcrumb",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Breadcrumb"), " constructor creates a breadcrumb. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnNav(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<nav>"), " element"},

			el.Code("b.OnUl(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<ul>"), " element"},

			[]any{"one of the class or style types defined in package ", el.Code("b")},
			[]any{"Apply the class or style to the ", el.Code(`<nav class="breadcrumb">`), " element"},

			el.Code("b.Centered"),
			"Center the breadcrumb in its container",

			el.Code("b.Right"),
			"Align the breadcrumb to the right",

			el.Code("b.ArrowSeparator"),
			[]any{"Use an arrow (", el.Code("→"), ") as the separator"},

			el.Code("b.BulletSeparator"),
			[]any{"Use a bullet (", el.Code("•"), ") as the separator"},

			el.Code("b.DotSeparator"),
			[]any{"Use a dot (", el.Code("·"), ") as the separator"},

			el.Code("b.SucceedsSeparator"),
			[]any{`Use a "succeeds" character (`, el.Code("≻"), `) as the separator`},

			el.Code("b.Small"),
			"Set breadcrumb size to small",

			el.Code("b.Medium"),
			"Set breadcrumb size to medium",

			el.Code("b.Large"),
			"Set breadcrumb size to large",
		),
		el.P("Other children are added to the ", el.Code("<ul>"), " element."),
		el.P(
			"The ", el.Code("b.BreadcrumbEntry"), " constructor creates a breadcrumb entry (which is a regular <li> element). The ", el.Code("b.BreadcrumbAHref"), " constructor creates a breadcrumb link entry. The ", el.Code("b.BreadcrumbActiveAHref"), " constructor creates an active breadcrumb link entry.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/breadcrumb/",
	c.Example(
		`b.Breadcrumb(
	b.BreadcrumbEntry(
		b.AHref("#", "Bulma"),
	),
	b.BreadcrumbEntry(
		b.AHref("#", "Documentation"),
	),
	b.BreadcrumbEntry(
		b.AHref("#", "Components"),
	),
	b.BreadcrumbEntry(
		b.Active,
		b.AHref("#", html.Aria("current", "page"), "Breadcrumb"),
	),
)`,
		b.Breadcrumb(
			b.BreadcrumbEntry(
				b.AHref("#", "Bulma"),
			),
			b.BreadcrumbEntry(
				b.AHref("#", "Documentation"),
			),
			b.BreadcrumbEntry(
				b.AHref("#", "Components"),
			),
			b.BreadcrumbEntry(
				b.Active,
				b.AHref("#", html.Aria("current", "page"), "Breadcrumb"),
			),
		),
	),
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the ", el.Code("b.BreadcrumbHref"), " and ", el.Code("b.BreadcrumbActiveAHref"), "helpers for regular breadcrumb links:")),
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
