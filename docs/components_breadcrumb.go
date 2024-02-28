package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var breadcrumb = c.NewPage(
	"Breadcrumb", "Breadcrumb", "/breadcrumb",
	"https://bulma.io/documentation/components/breadcrumb/",
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
	b.Content(el.P("Additionally, ", el.Em("Bulma-Gomponents"), " provides a helper for simple href/text breadcrumbs, ", el.Code("easy.Breadcrumb"), ":")),
	c.Example(
		`easy.Breadcrumb(
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb", // The exclamation mark states this entry is active
)`,
		easy.Breadcrumb(
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
).Section(
	"Alignment",
	"https://bulma.io/documentation/components/breadcrumb/#alignment",
	c.Example(
		`b.Breadcrumb(
	b.Centered,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
),
easy.Breadcrumb(
	b.Centered,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		b.Breadcrumb(
			b.Centered,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
		easy.Breadcrumb(
			b.Centered,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Right,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbActiveAHref("#", "Breadcrumb"),
),
easy.Breadcrumb(
	b.Right,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		b.Breadcrumb(
			b.Right,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbActiveAHref("#", "Breadcrumb"),
		),
		easy.Breadcrumb(
			b.Right,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
).Section(
	"Icons",
	"https://bulma.io/documentation/components/breadcrumb/#icons",
	c.Example(
		`b.Breadcrumb(
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "home"), "Bulma"),
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "book"), "Documentation"),
	b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "puzzle-piece"), "Components"),
	b.BreadcrumbActiveAHref("#", fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
),
easy.Breadcrumb(
	"#", []any{fa.Icon(fa.Solid, "home"), "Bulma"},
	"#", []any{fa.Icon(fa.Solid, "book"), "Documentation"},
	"#", []any{fa.Icon(fa.Solid, "puzzle-piece"), "Components"},
	"!#", []any{fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"},
)`,
		b.Breadcrumb(
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "home"), "Bulma"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "book"), "Documentation"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "puzzle-piece"), "Components"),
			b.BreadcrumbActiveAHref("#", fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
		),
		easy.Breadcrumb(
			"#", []any{fa.Icon(fa.Solid, "home"), "Bulma"},
			"#", []any{fa.Icon(fa.Solid, "book"), "Documentation"},
			"#", []any{fa.Icon(fa.Solid, "puzzle-piece"), "Components"},
			"!#", []any{fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"},
		),
	),
).Section(
	"Alternative separators",
	"https://bulma.io/documentation/components/breadcrumb/#alternative-separators",
	c.Example(
		`easy.Breadcrumb(
	b.ArrowSeparator,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.ArrowSeparator,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`easy.Breadcrumb(
	b.BulletSeparator,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
`,
		easy.Breadcrumb(
			b.BulletSeparator,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`easy.Breadcrumb(
	b.DotSeparator,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.DotSeparator,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`easy.Breadcrumb(
	b.SucceedsSeparator,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.SucceedsSeparator,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/components/breadcrumb/#sizes",
	c.Example(
		`easy.Breadcrumb(
	b.Small,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.Small,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`easy.Breadcrumb(
	b.Medium,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.Medium,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
	c.Example(
		`easy.Breadcrumb(
	b.Large,
	"#", "Bulma",
	"#", "Documentation",
	"#", "Components",
	"!#", "Breadcrumb",
)`,
		easy.Breadcrumb(
			b.Large,
			"#", "Bulma",
			"#", "Documentation",
			"#", "Components",
			"!#", "Breadcrumb",
		),
	),
)
