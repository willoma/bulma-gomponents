package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var columnsSizes = c.NewPage(
	"Sizes", "Column sizes", "/columns/sizes",
	"https://bulma.io/documentation/columns/sizes/",

	c.Example(
		`b.Columns(
	b.Column(
		b.SizeFourFifths,
		el.P("is-four-fifths"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.ThreeQuarters,
		el.P("is-three-quarters"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.TwoThirds,
		el.P("is-two-thirds"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.ThreeFifths,
		el.P("is-three-fifths"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.Half,
		el.P("is-half"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.TwoFifths,
		el.P("is-two-fifths"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneThird,
		el.P("is-one-third"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneQuarter,
		el.P("is-one-quarter"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneFifth,
		el.P("is-one-fifth"),
	),
	b.Column(el.P("Auto")),
	b.Column(el.P("Auto")),
)`,
		b.Columns(
			b.Column(
				b.FourFifths,
				c.ColParagraph("is-four-fifths"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.ThreeQuarters,
				c.ColParagraph("is-three-quarters"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.TwoThirds,
				c.ColParagraph("is-two-thirds"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.ThreeFifths,
				c.ColParagraph("is-three-fifths"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.Half,
				c.ColParagraph("is-half"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.TwoFifths,
				c.ColParagraph("is-two-fifths"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.OneThird,
				c.ColParagraph("is-one-third"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.OneQuarter,
				c.ColParagraph("is-one-quarter"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
		b.Columns(
			b.Column(
				b.OneFifth,
				c.ColParagraph("is-one-fifth"),
			),
			b.Column(c.ColParagraph("Auto")),
			b.Column(c.ColParagraph("Auto")),
		),
	),
).Section(
	"12 columns system",
	"https://bulma.io/documentation/columns/sizes/",

	b.Content(el.P("Use ", el.Code("bulma.ColSize1"), " to ", el.Code("bulma.ColSize12"), ".")),
).Section(
	"Offset",
	"https://bulma.io/documentation/columns/sizes/#offset",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Column(
		b.Half,
		b.OffsetOneQuarter,
		el.P(
			el.Code("is-half"),
			el.Br(),
			el.Code("is-offset-one-quarter"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.ThreeFifths,
		b.OffsetOneFifth,
		el.P(
			el.Code("is-three-fifths"),
			el.Br(),
			el.Code("is-offset-one-fifth"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.Size4,
		b.Offset8,
		el.P(
			el.Code("is-4"),
			el.Br(),
			el.Code("is-offset-8"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.Size11,
		b.Offset1,
		el.P(
			el.Code("is-11"),
			el.Br(),
			el.Code("is-offset-1"),
		),
	),
)`,
		b.Columns(
			b.Mobile,
			b.Column(
				b.Half,
				b.OffsetOneQuarter,
				c.ColParagraph(
					el.Code("is-half"),
					el.Br(),
					el.Code("is-offset-one-quarter"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.ThreeFifths,
				b.OffsetOneFifth,
				c.ColParagraph(
					el.Code("is-three-fifths"),
					el.Br(),
					el.Code("is-offset-one-fifth"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Size4,
				b.Offset8,
				c.ColParagraph(
					el.Code("is-4"),
					el.Br(),
					el.Code("is-offset-8"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Size11,
				b.Offset1,
				c.ColParagraph(
					el.Code("is-11"),
					el.Br(),
					el.Code("is-offset-1"),
				),
			),
		),
	),
).Section(
	"Narrow column",
	"https://bulma.io/documentation/columns/sizes/#narrow-column",
	c.Example(
		`b.Columns(
	b.Column(
		b.Narrow,
		b.Box(
			b.Style("width", "200px"),
			b.Title5(html.P, "Narrow column"),
			b.Subtitle(html.P, "This column is only 200px wide."),
		),
	),
	b.Column(
		b.Box(
			b.Title5(html.P, "Flexible column"),
			b.Subtitle(html.P, "This column will take up the remaining space available."),
		),
	),
)`,
		b.Columns(
			b.Column(
				b.Narrow,
				b.Box(
					b.Style("width", "200px"),
					b.Title5(html.P, "Narrow column"),
					b.Subtitle(html.P, "This column is only 200px wide."),
				),
			),
			b.Column(
				b.Box(
					b.Title5(html.P, "Flexible column"),
					b.Subtitle(html.P, "This column will take up the remaining space available."),
				),
			),
		),
	),
)
