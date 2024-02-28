package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var columnsNesting = c.NewPage(
	"Nesting", "Nesting columns", "/columns/nesting",
	"https://bulma.io/documentation/columns/nesting/",

	c.Example(
		`b.Columns(
	b.Column(
		el.P(
			b.BackgroundColor(b.Info),
			"First column",
		),
		b.Columns(
			b.Mobile,
			b.Column(
				el.P(
					b.BackgroundColor(b.Info),
					"First nested column",
				),
			),
			b.Column(
				el.P(
					b.BackgroundColor(b.Info),
					"Second nested column",
				),
			),
		),
	),
	b.Column(
		el.P(
			b.BackgroundColor(b.Danger),
			"Second column",
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Half,
				el.P(
					b.BackgroundColor(b.Danger),
					"50%",
				),
			),
			b.Column(
				el.P(
					b.BackgroundColor(b.Danger),
					"Auto",
				),
			),
			b.Column(
				el.P(
					b.BackgroundColor(b.Danger),
					"Auto",
				),
			),
		),
	),
)`,
		b.Columns(
			b.Column(
				c.ColParagraph(
					b.Info.Background(),
					"First column",
				),
				b.Columns(
					b.Mobile,
					b.Column(
						c.ColParagraph(
							b.Info.Background(),
							"First nested column",
						),
					),
					b.Column(
						c.ColParagraph(
							b.Info.Background(),
							"Second nested column",
						),
					),
				),
			),
			b.Column(
				c.ColParagraph(
					b.Danger.Background(),
					"Second column",
				),
				b.Columns(
					b.Mobile,
					b.Column(
						b.Half,
						c.ColParagraph(
							b.Danger.Background(),
							"50%",
						),
					),
					b.Column(
						c.ColParagraph(
							b.Danger.Background(),
							"Auto",
						),
					),
					b.Column(
						c.ColParagraph(
							b.Danger.Background(),
							"Auto",
						),
					),
				),
			),
		),
	),
)
