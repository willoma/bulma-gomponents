package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var section = c.NewPage(
	"Section", "Section", "/section",
	"",

	b.Content(
		el.P("The ", el.Code("b.Section"), " constructor creates a section. The following children have a special meaning:"),
		b.DList(
			el.Code("b.Medium"),
			"Set spacing to medium",

			el.Code("b.Large"),
			"Set spacing to large",
		),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/section/",
	c.Example(
		`b.Section(
	b.Title("Section"),
	b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Title("Section"),
			b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
		),
	),
	c.Example(
		`b.Section(
	b.Medium,
	b.Title("Section"),
	b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Medium,
			b.Title("Section"),
			b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
		),
	),
	c.Example(
		`b.Section(
	b.Large,
	b.Title("Medium section"),
	b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
)`,
		b.Section(
			b.Large,
			b.Title("Large section"),
			b.Subtitle("A simple container to divide your page into ", el.Strong("sections"), ", like the one you're currently reading."),
		),
	),
)
