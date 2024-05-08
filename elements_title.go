package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var title = c.NewPage(
	"Title", "Title and subtitle", "/title",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Title"), " constructor creates a title. The ", e.Code("b.Subtitle"), " constructor creates a subtitle.",
		),
		c.Modifiers(
			c.Row("b.Spaced", "Maintain normal spacing between title and the following subtitle"),
			c.Row("1", "Set size to 1"),
			c.Row("2", "Set size to 2"),
			c.Row("3", "Set size to 3"),
			c.Row("4", "Set size to 4"),
			c.Row("5", "Set size to 5"),
			c.Row("6", "Set size to 6"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/title/",
	c.Example(
		`b.Title("Title"),
b.Subtitle("Subtitle")`,
		b.Title("Title"),
		b.Subtitle("Subtitle"),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/title/#sizes",
	c.Example(
		`b.Title(1, "Title 1"),
b.Title(2, "Title 2"),
b.Title(3, "Title 3"),
b.Title(4, "Title 4"),
b.Title(5, "Title 5"),
b.Title(6, "Title 6")`,
		b.Title(1, "Title 1"),
		b.Title(2, "Title 2"),
		b.Title(3, "Title 3"),
		b.Title(4, "Title 4"),
		b.Title(5, "Title 5"),
		b.Title(6, "Title 6"),
	),
	c.Example(
		`b.Subtitle(1, "Subtitle 1"),
b.Subtitle(2, "Subtitle 2"),
b.Subtitle(3, "Subtitle 3"),
b.Subtitle(4, "Subtitle 4"),
b.Subtitle(5, "Subtitle 5"),
b.Subtitle(6, "Subtitle 6")`,
		b.Subtitle(1, "Subtitle 1"),
		b.Subtitle(2, "Subtitle 2"),
		b.Subtitle(3, "Subtitle 3"),
		b.Subtitle(4, "Subtitle 4"),
		b.Subtitle(5, "Subtitle 5"),
		b.Subtitle(6, "Subtitle 6"),
	),
	c.Example(
		`b.Block(
	b.Title(1, html.P, "Title 1"),
	b.Subtitle(3, html.P, "Subtitle 3"),
),
b.Block(
	b.Title(2, html.P, "Title 2"),
	b.Subtitle(4, html.P, "Subtitle 4"),
),
b.Block(
	b.Title(3, html.P, "Title 3"),
	b.Subtitle(5, html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title(1, html.P, "Title 1"),
			b.Subtitle(3, html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title(2, html.P, "Title 2"),
			b.Subtitle(4, html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title(3, html.P, "Title 3"),
			b.Subtitle(5, html.P, "Subtitle 5"),
		),
	),
	c.Example(
		`b.Block(
	b.Title(1, html.P, b.Spaced, "Title 1"),
	b.Subtitle(3, html.P, "Subtitle 3"),
),
b.Block(
	b.Title(2, html.P, b.Spaced, "Title 2"),
	b.Subtitle(4, html.P, "Subtitle 4"),
),
b.Block(
	b.Title(3, html.P, b.Spaced, "Title 3"),
	b.Subtitle(5, html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title(1, html.P, b.Spaced, "Title 1"),
			b.Subtitle(3, html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title(2, html.P, b.Spaced, "Title 2"),
			b.Subtitle(4, html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title(3, html.P, b.Spaced, "Title 3"),
			b.Subtitle(5, html.P, "Subtitle 5"),
		),
	),
)
