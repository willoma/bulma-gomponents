package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var section = c.NewPage(
	"Section", "Section", "/section",
	"",

	b.Content(
		e.P("The ", e.Code("b.Section"), " constructor creates a section. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Medium"),
			"Set spacing to medium",

			e.Code("b.Large"),
			"Set spacing to large",
		),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/section/",
	c.Example(
		`b.Section(
	b.Title("Section"),
	b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Title("Section"),
			b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
		),
	),
	c.Example(
		`b.Section(
	b.Medium,
	b.Title("Section"),
	b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Medium,
			b.Title("Section"),
			b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
		),
	),
	c.Example(
		`b.Section(
	b.Large,
	b.Title("Medium section"),
	b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Large,
			b.Title("Large section"),
			b.Subtitle("A simple container to divide your page into ", e.Strong("sections"), ", like the one you're currently reading."),
		),
	),
)
