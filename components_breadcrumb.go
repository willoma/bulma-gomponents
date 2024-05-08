package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var breadcrumb = c.NewPage(
	"Breadcrumb", "Breadcrumb", "/breadcrumb",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Breadcrumb"), " constructor creates a breadcrumb.",
		),
		c.Modifiers(
			c.Row("b.Centered", "Center the breadcrumb in its container"),
			c.Row("b.Right", "Align the breadcrumb to the right"),
			c.Row("b.ArrowSeparator", "Use an arrow (", e.Code("→"), ") as the separator"),
			c.Row("b.BulletSeparator", "Use a bullet (", e.Code("•"), ") as the separator"),
			c.Row("b.DotSeparator", "Use a dot (", e.Code("·"), ") as the separator"),
			c.Row("b.SucceedsSeparator", `Use a "succeeds" character (`, e.Code("≻"), `) as the separator`),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),
		c.Children(
			c.Row("b.OnUl(...any)", "Apply children to the ", e.Code("<ul>"), " element"),
			c.Row("b.OnNav(...any)", "Apply children to the ", e.Code("<nav>"), " element"),
			c.RowClassesStyles("Apply child to the ", e.Code(`<nav>`), " element"),
			c.RowNodeAttribute("Apply child to the ", e.Code(`<nav>`), " element"),
			c.RowNodeElement("Apply child to the ", e.Code(`<ul>`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<ul>`), " element"),
		),
		e.P(
			"The ", e.Code("b.BreadcrumbEntry"), " constructor creates a breadcrumb entry (which is a regular ", e.Code("<li>"), " element).",
		),
		e.P(
			"The ", e.Code("b.BreadcrumbAHref"), " constructor creates a breadcrumb link entry.",
		),
		c.Modifiers(
			c.Row("b.Active", "Apply the active style and add the ", e.Code(`aria-current="page"`), " attribute"),
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
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
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
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Centered,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Right,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Right,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
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
	b.BreadcrumbAHref("#", b.Active, fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "home"), "Bulma"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "book"), "Documentation"),
			b.BreadcrumbAHref("#", fa.Icon(fa.Solid, "puzzle-piece"), "Components"),
			b.BreadcrumbAHref("#", b.Active, fa.Icon(fa.Solid, "thumbs-up"), "Breadcrumb"),
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
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.ArrowSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.BulletSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
`,
		b.Breadcrumb(
			b.BulletSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.DotSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.DotSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.SucceedsSeparator,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.SucceedsSeparator,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
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
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Small,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Medium,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Medium,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
	c.Example(
		`b.Breadcrumb(
	b.Large,
	b.BreadcrumbAHref("#", "Bulma"),
	b.BreadcrumbAHref("#", "Documentation"),
	b.BreadcrumbAHref("#", "Components"),
	b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
)`,
		b.Breadcrumb(
			b.Large,
			b.BreadcrumbAHref("#", "Bulma"),
			b.BreadcrumbAHref("#", "Documentation"),
			b.BreadcrumbAHref("#", "Components"),
			b.BreadcrumbAHref("#", b.Active, "Breadcrumb"),
		),
	),
)
