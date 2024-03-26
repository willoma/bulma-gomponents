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
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Breadcrumb"), " constructor creates a breadcrumb. It accepts the following values additionally to the standard set of children types:",
		),
		b.DList(
			el.Code("b.Inner(any)"),
			[]any{"forcibly apply the child to the ", el.Code("<ul>"), " element"},
			el.Code("b.Outer(any)"),
			[]any{"forcibly apply the child to the ", el.Code("<nav>"), " element"},
			el.Code("b.Centered"),
			"center the breadcrumb in its container",
			el.Code("b.Right"),
			"align the breadcrumb to the right",
			el.Code("b.ArrowSeparator"),
			[]any{"use an arrow (", el.Code("→"), ") as the separator"},
			el.Code("b.BulletSeparator"),
			[]any{"use a bullet (", el.Code("•"), ") as the separator"},
			el.Code("b.DotSeparator"),
			[]any{"use a dot (", el.Code("·"), ") as the separator"},
			el.Code("b.SucceedsSeparator"),
			[]any{`use a "succeeds" character (`, el.Code("≻"), `) as the separator`},
			el.Code("b.Small"),
			"set breadcrumb size to small",
			el.Code("b.Medium"),
			"set breadcrumb size to medium",
			el.Code("b.Large"),
			"set breadcrumb size to large",
		),
		el.P(
			"The ", el.Code("b.BreadcrumbEntry"), " constructor creates a breadcrumb entry (which is a regular <li> element). The ", el.Code("b.BreadcrumbAHref"), " constructor creates a breadcrumb link entry. The ", el.Code("b.BreadcrumbActiveAHref"), " constructor creates an active breadcrumb link entry.",
		),
	),
).Section(
	"Easy helper", "",

	b.Content(
		el.P("The ", el.Code("easy.Breadcrumb"), " constructor builds a breadcrumb from a single list of target/element pairs. Targets must be strings, which are used as the href attribute of the ", el.Code("<a>"), ` element. If the target starts with the "`, el.Code("!"), `" character, it makes it the active element. Content may be either a string, a single `, el.Code("b.Element"), ", a single ", el.Code("gomponents.Node"), ", or a ", el.Code("[]any"), ". When children are classes, styles or other attributes, they are applied to the breadcrumb component."),
	),
	c.Example(
		`easy.Breadcrumb(
	b.ArrowSeparator,

	"/",
	[]any{fa.Icon(fa.Solid, "home"), "Root"},

	"/profile",
	"Profile",

	"!/profile/email",
	"Email",
)`,
		easy.Breadcrumb(
			b.ArrowSeparator,

			"/",
			[]any{fa.Icon(fa.Solid, "home"), "Root"},

			"/profile",
			"Profile",

			"!/profile/email",
			"Email",
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
).Subsection(
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
).Subsection(
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
).Subsection(
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
