package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var columnsOptions = c.NewPage(
	"Options", "Column options", "/columns/options",
	"https://bulma.io/documentation/columns/options/",
).Section(
	"Vertical alignment",
	"https://bulma.io/documentation/columns/options/#vertical-alignment",
	c.Example(
		`b.Columns(
	b.VCentered,
	b.Column(
		b.Size8,
		el.P("First column"),
	),
	b.Column(
		el.P("Second column with more content. This is so you can see the vertical alignment."),
	),
)`,
		b.Columns(
			b.VCentered,
			b.Column(
				b.Size8,
				c.ColParagraph("First column"),
			),
			b.Column(
				c.ColParagraph("Second column with more content. This is so you can see the vertical alignment."),
			),
		),
	),
).Section(
	"Multiline",
	"https://bulma.io/documentation/columns/options/#multiline",
	c.Example(
		`b.Columns(
	b.Multiline,
	b.Mobile,
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.Half,
		el.P("is-half"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(
		el.P("Auto"),
	),
)`,
		b.Columns(
			b.Multiline,
			b.Mobile,
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.Half,
				c.ColParagraph("is-half"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(
				c.ColParagraph("Auto"),
			),
		),
	),
).Section(
	"Centering columns",
	"https://bulma.io/documentation/columns/options/#centering-columns",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Centered,
	b.Column(
		b.Half,
		el.P(el.Code("is-half")),
	),
)`,
		b.Columns(
			b.Mobile,
			b.Centered,
			b.Column(
				b.Half,
				c.ColParagraph(el.Code("is-half")),
			),
		),
	),
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Multiline,
	b.Centered,
	b.Column(
		b.Narrow,
		el.P(
			el.Code("is-narrow"),
			el.Br(),
			"First column",
		),
	),
	b.Column(
		b.Narrow,
		el.P(
			el.Code("is-narrow"),
			el.Br(),
			"Our second column",
		),
	),
	b.Column(
		b.Narrow,
		el.P(
			el.Code("is-narrow"),
			el.Br(),
			"Third column",
		),
	),
	b.Column(
		b.Narrow,
		el.P(
			el.Code("is-narrow"),
			el.Br(),
			"The fourth column",
		),
	),
	b.Column(
		b.Narrow,
		el.P(
			el.Code("is-narrow"),
			el.Br(),
			"Fifth column",
		),
	),
)`,
		b.Columns(
			b.Mobile,
			b.Multiline,
			b.Centered,
			b.Column(
				b.Narrow,
				c.ColParagraph(
					el.Code("is-narrow"),
					el.Br(),
					"First column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					el.Code("is-narrow"),
					el.Br(),
					"Our second column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					el.Code("is-narrow"),
					el.Br(),
					"Third column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					el.Code("is-narrow"),
					el.Br(),
					"The fourth column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					el.Code("is-narrow"),
					el.Br(),
					"Fifth column",
				),
			),
		),
	),
)
