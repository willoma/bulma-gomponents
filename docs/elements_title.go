package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var title = c.NewPage(
	"Title", "Title and subtitle", "/title",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Title"), " constructor creates a title. The ", e.Code("b.Subtitle"), " constructor creates a subtitle. These constructors accept the ", e.Code("b.Spaced"), " modifier additionally to the standard set of children types, to maintain normal spacing between title and the following subtitle - otherwise, a title followed by a subtitle are grouped together.",
		),
		e.P(
			"The following constructors create titles and subtitles of sizes 1 to 6:",
			b.UList(
				e.Code("b.Title1"),
				e.Code("b.Title2"),
				e.Code("b.Title3"),
				e.Code("b.Title4"),
				e.Code("b.Title5"),
				e.Code("b.Title6"),
				e.Code("b.Subtitle1"),
				e.Code("b.Subtitle2"),
				e.Code("b.Subtitle3"),
				e.Code("b.Subtitle4"),
				e.Code("b.Subtitle5"),
				e.Code("b.Subtitle6"),
			),
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
		`b.Title1("Title 1"),
b.Title2("Title 2"),
b.Title3("Title 3"),
b.Title4("Title 4"),
b.Title5("Title 5"),
b.Title6("Title 6")`,
		b.Title1("Title 1"),
		b.Title2("Title 2"),
		b.Title3("Title 3"),
		b.Title4("Title 4"),
		b.Title5("Title 5"),
		b.Title6("Title 6"),
	),
	c.Example(
		`b.Subtitle1("Subtitle 1"),
b.Subtitle2("Subtitle 2"),
b.Subtitle3("Subtitle 3"),
b.Subtitle4("Subtitle 4"),
b.Subtitle5("Subtitle 5"),
b.Subtitle6("Subtitle 6")`,
		b.Subtitle1("Subtitle 1"),
		b.Subtitle2("Subtitle 2"),
		b.Subtitle3("Subtitle 3"),
		b.Subtitle4("Subtitle 4"),
		b.Subtitle5("Subtitle 5"),
		b.Subtitle6("Subtitle 6"),
	),
	c.Example(
		`b.Block(
	b.Title1(html.P, "Title 1"),
	b.Subtitle3(html.P, "Subtitle 3"),
),
b.Block(
	b.Title2(html.P, "Title 2"),
	b.Subtitle4(html.P, "Subtitle 4"),
),
b.Block(
	b.Title3(html.P, "Title 3"),
	b.Subtitle5(html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title1(html.P, "Title 1"),
			b.Subtitle3(html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title2(html.P, "Title 2"),
			b.Subtitle4(html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title3(html.P, "Title 3"),
			b.Subtitle5(html.P, "Subtitle 5"),
		),
	),
	c.Example(
		`b.Block(
	b.Title1(html.P, b.Spaced, "Title 1"),
	b.Subtitle3(html.P, "Subtitle 3"),
),
b.Block(
	b.Title2(html.P, b.Spaced, "Title 2"),
	b.Subtitle4(html.P, "Subtitle 4"),
),
b.Block(
	b.Title3(html.P, b.Spaced, "Title 3"),
	b.Subtitle5(html.P, "Subtitle 5"),
)`,
		b.Block(
			b.Title1(html.P, b.Spaced, "Title 1"),
			b.Subtitle3(html.P, "Subtitle 3"),
		),
		b.Block(
			b.Title2(html.P, b.Spaced, "Title 2"),
			b.Subtitle4(html.P, "Subtitle 4"),
		),
		b.Block(
			b.Title3(html.P, b.Spaced, "Title 3"),
			b.Subtitle5(html.P, "Subtitle 5"),
		),
	),
)
