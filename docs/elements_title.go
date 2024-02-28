package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var title = c.NewPage(
	"Title", "Title and subtitle", "/title",
	"https://bulma.io/documentation/elements/title/",
	c.Example(
		`b.Title("Title"),
b.Subtitle("Subtitle")`,
		b.Title("Title"),
		b.Subtitle("Subtitle"),
	),
).Section(
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
		`b.Title1(html.P, "Title 1"),
b.Subtitle3(html.P, "Subtitle 3"),
b.Title2(html.P, "Title 2"),
b.Subtitle4(html.P, "Subtitle 4"),
b.Title3(html.P, "Title 3"),
b.Subtitle5(html.P, "Subtitle 5")`,
		b.Title1(html.P, "Title 1"),
		b.Subtitle3(html.P, "Subtitle 3"),
		b.Title2(html.P, "Title 2"),
		b.Subtitle4(html.P, "Subtitle 4"),
		b.Title3(html.P, "Title 3"),
		b.Subtitle5(html.P, "Subtitle 5"),
	),
	c.Example(
		`b.Title1(html.P, b.Spaced, "Title 1"),
b.Subtitle3(html.P, "Subtitle 3"),
b.Title2(html.P, b.Spaced, "Title 2"),
b.Subtitle4(html.P, "Subtitle 4"),
b.Title3(html.P, b.Spaced, "Title 3"),
b.Subtitle5(html.P, "Subtitle 5")`,
		b.Title1(html.P, b.Spaced, "Title 1"),
		b.Subtitle3(html.P, "Subtitle 3"),
		b.Title2(html.P, b.Spaced, "Title 2"),
		b.Subtitle4(html.P, "Subtitle 4"),
		b.Title3(html.P, b.Spaced, "Title 3"),
		b.Subtitle5(html.P, "Subtitle 5"),
	),
)
