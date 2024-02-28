package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var section = c.NewPage(
	"Section", "Section", "/section",
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
