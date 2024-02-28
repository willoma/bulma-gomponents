package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var columnsBasics = c.NewPage(
	"Basics", "Columns powered by Flexbox", "/columns/index",
	"https://bulma.io/documentation/columns/basics/",

	c.Example(
		`b.Columns(
	b.Column(el.P("First column")),
	b.Column(el.P("Second column")),
	b.Column(el.P("Third column")),
	b.Column(el.P("Fourth column")),
)`,
		b.Columns(
			b.Column(c.ColParagraph("First column")),
			b.Column(c.ColParagraph("Second column")),
			b.Column(c.ColParagraph("Third column")),
			b.Column(c.ColParagraph("Fourth column")),
		),
	),
)
