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
			c.Row("b.TitleLevel1", "Set size to 1"),
			c.Row("b.TitleLevel2", "Set size to 2"),
			c.Row("b.TitleLevel3", "Set size to 3"),
			c.Row("b.TitleLevel4", "Set size to 4"),
			c.Row("b.TitleLevel5", "Set size to 5"),
			c.Row("b.TitleLevel6", "Set size to 6"),
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
		`b.Title(b.TitleLevel1, "Title 1"),
b.Title(b.TitleLevel2, "Title 2"),
b.Title(b.TitleLevel3, "Title 3"),
b.Title(b.TitleLevel4, "Title 4"),
b.Title(b.TitleLevel5, "Title 5"),
b.Title(b.TitleLevel6, "Title 6")`,
		b.Title(b.TitleLevel1, "Title 1"),
		b.Title(b.TitleLevel2, "Title 2"),
		b.Title(b.TitleLevel3, "Title 3"),
		b.Title(b.TitleLevel4, "Title 4"),
		b.Title(b.TitleLevel5, "Title 5"),
		b.Title(b.TitleLevel6, "Title 6"),
	),
	c.Example(
		`b.Subtitle(b.TitleLevel1, "Subtitle 1"),
b.Subtitle(b.TitleLevel2, "Subtitle 2"),
b.Subtitle(b.TitleLevel3, "Subtitle 3"),
b.Subtitle(b.TitleLevel4, "Subtitle 4"),
b.Subtitle(b.TitleLevel5, "Subtitle 5"),
b.Subtitle(b.TitleLevel6, "Subtitle 6")`,
		b.Subtitle(b.TitleLevel1, "Subtitle 1"),
		b.Subtitle(b.TitleLevel2, "Subtitle 2"),
		b.Subtitle(b.TitleLevel3, "Subtitle 3"),
		b.Subtitle(b.TitleLevel4, "Subtitle 4"),
		b.Subtitle(b.TitleLevel5, "Subtitle 5"),
		b.Subtitle(b.TitleLevel6, "Subtitle 6"),
	),
	c.Example(
		`b.Block(
	b.Title(b.TitleLevel1, html.P, "Title 1"),
	b.Subtitle(b.TitleLevel3, html.P, "Subtitle 3"),
),
b.Block(
	b.Title(b.TitleLevel2, html.P, "Title 2"),
	b.Subtitle(b.TitleLevel4, html.P, "Subtitle 4"),
),
b.Block(
	b.Title(b.TitleLevel3, html.P, "Title 3"),
	b.Subtitle(b.TitleLevel5, html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title(b.TitleLevel1, html.P, "Title 1"),
			b.Subtitle(b.TitleLevel3, html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title(b.TitleLevel2, html.P, "Title 2"),
			b.Subtitle(b.TitleLevel4, html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title(b.TitleLevel3, html.P, "Title 3"),
			b.Subtitle(b.TitleLevel5, html.P, "Subtitle 5"),
		),
	),
	c.Example(
		`b.Block(
	b.Title(b.TitleLevel1, html.P, b.Spaced, "Title 1"),
	b.Subtitle(b.TitleLevel3, html.P, "Subtitle 3"),
),
b.Block(
	b.Title(b.TitleLevel2, html.P, b.Spaced, "Title 2"),
	b.Subtitle(b.TitleLevel4, html.P, "Subtitle 4"),
),
b.Block(
	b.Title(b.TitleLevel3, html.P, b.Spaced, "Title 3"),
	b.Subtitle(b.TitleLevel5, html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title(b.TitleLevel1, html.P, b.Spaced, "Title 1"),
			b.Subtitle(b.TitleLevel3, html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title(b.TitleLevel2, html.P, b.Spaced, "Title 2"),
			b.Subtitle(b.TitleLevel4, html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title(b.TitleLevel3, html.P, b.Spaced, "Title 3"),
			b.Subtitle(b.TitleLevel5, html.P, "Subtitle 5"),
		),
	),
)
