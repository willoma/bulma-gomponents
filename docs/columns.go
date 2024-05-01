package docs

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

func frac(num, denom int) gomponents.Node {
	return gomponents.El(
		"math",
		gomponents.El(
			"mfrac",
			gomponents.El(
				"mi",
				gomponents.Raw(strconv.Itoa(num)),
			),
			gomponents.El(
				"mi",
				gomponents.Raw(strconv.Itoa(denom)),
			),
		),
	)
}

var columns = c.NewPage(
	"Columns", "Columns powered by Flexbox", "/columns",
	"",

	b.Content(
		e.P("The ", e.Code("b.Columns"), " constructor creates a columns container. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Column(...)"),
			"Add the column",

			e.Code("e.Element"),
			"Wrap the element in a column and add it",

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the columns container",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Wrap the element in a column and add it",

			e.Code("b.Gapless"),
			"No gap",

			e.Code("b.ColumnGap0"),
			"No gap",

			e.Code("b.ColumnGap1"),
			"0.25rem gap",

			e.Code("b.ColumnGap2"),
			"0.5rem gap",

			e.Code("b.ColumnGap3"),
			"0.75rem gap",

			e.Code("b.ColumnGap4"),
			"1rem gap",

			e.Code("b.ColumnGap5"),
			"1.25rem gap",

			e.Code("b.ColumnGap6"),
			"1.5rem gap",

			e.Code("b.ColumnGap7"),
			"1.75rem gap",

			e.Code("b.ColumnGap8"),
			"2rem gap",

			e.Code("b.Centered"),
			"Center columns",

			e.Code("b.Desktop"),
			"Allow columns only on desktops upward (not on tablets)",

			e.Code("b.Mobile"),
			"Allow columns on mobile phones too",

			e.Code("b.Multiline"),
			"Create a new line when columns do not fit in a single line",

			e.Code("b.VCentered"),
			"Center columns vertically",
		),
		e.P("The gap may be breakpoint-based, with ", e.Code("b.ColumnGap*.Mobile()"), " to ", e.Code("b.ColumnGap*.FullHD()"), "."),

		e.P("The ", e.Code("b.Column"), " constructor creates a single column. The following children have a special meaning:"),
		b.DList(
			[]any{e.Code("b.Full"), ", ", e.Code("b.Size12")},
			[]any{"Set column width to 100% (", frac(12, 12), ")"},

			e.Code("b.Size11"),
			[]any{"Set column width to 91.66% (", frac(11, 12), ")"},

			e.Code("b.Size10"),
			[]any{"Set column width to 83.33% (", frac(10, 12), ")"},

			e.Code("b.FourFifths"),
			[]any{"Set column width to 80% (", frac(4, 5), ")"},

			[]any{e.Code("b.ThreeQuarters"), ", ", e.Code("b.Size9")},
			[]any{"Set column width to 75% (", frac(3, 4), ", ", frac(9, 12), ")"},

			[]any{e.Code("b.TwoThirds"), ", ", e.Code("b.Size8")},
			[]any{"Set column width to 66.66% (", frac(2, 3), ", ", frac(8, 12), ")"},

			e.Code("b.ThreeFifths"),
			[]any{"Set column width to 60% (", frac(3, 5), ")"},

			e.Code("b.Size7"),
			[]any{"Set column width to 58.33% (", frac(7, 12), ")"},

			[]any{e.Code("b.Half"), ", ", e.Code("b.Size6")},
			[]any{"Set column width to 50% (", frac(1, 2), ", ", frac(6, 12), ")"},

			e.Code("b.Size5"),
			[]any{"Set column width to 41.66% (", frac(5, 12), ")"},

			e.Code("b.TwoFifths"),
			[]any{"Set column width to 40% (", frac(2, 5), ")"},

			[]any{e.Code("b.OneThird"), ", ", e.Code("b.Size4")},
			[]any{"Set column width to 33.33% (", frac(1, 3), ", ", frac(4, 12), ")"},

			[]any{e.Code("b.OneQuarter"), ", ", e.Code("b.Size3")},
			[]any{"Set column width to 25% (", frac(1, 4), ", ", frac(3, 12), ")"},

			e.Code("b.OneFifth"),
			[]any{"Set column width to 20% (", frac(1, 5), ")"},

			e.Code("b.Size2"),
			[]any{"Set column width to 16.66% (", frac(2, 12), ")"},

			e.Code("b.Size1"),
			[]any{"Set column width to 8.33% (", frac(1, 12), ")"},

			e.Code("b.Offset11"),
			[]any{"Set column offset to 91.66% (", frac(11, 12), ")"},

			e.Code("b.Offset10"),
			[]any{"Set column offset to 83.33% (", frac(10, 12), ")"},

			e.Code("b.OffsetFourFifths"),
			[]any{"Set column offset to 80% (", frac(4, 5), ")"},

			[]any{e.Code("b.OffsetThreeQuarters"), ", ", e.Code("b.Offset9")},
			[]any{"Set column offset to 75% (", frac(3, 4), ", ", frac(9, 12), ")"},

			[]any{e.Code("b.OffsetTwoThirds"), ", ", e.Code("b.Offset8")},
			[]any{"Set column offset to 66.66% (", frac(2, 3), ", ", frac(8, 12), ")"},

			e.Code("b.OffsetThreeFifths"),
			[]any{"Set column offset to 60% (", frac(3, 5), ")"},

			e.Code("b.Offset7"),
			[]any{"Set column offset to 58.33% (", frac(7, 12), ")"},

			[]any{e.Code("b.OffsetHalf"), ", ", e.Code("b.Offset6")},
			[]any{"Set column offset to 50% (", frac(1, 2), ", ", frac(6, 12), ")"},

			e.Code("b.Offset5"),
			[]any{"Set column offset to 41.66% (", frac(5, 12), ")"},

			e.Code("b.OffsetTwoFifths"),
			[]any{"Set column offset to 40% (", frac(2, 5), ")"},

			[]any{e.Code("b.OffsetOneThird"), ", ", e.Code("b.Offset4")},
			[]any{"Set column offset to 33.33% (", frac(1, 3), ", ", frac(4, 12), ")"},

			[]any{e.Code("b.OffsetOneQuarter"), ", ", e.Code("b.Offset3")},
			[]any{"Set column offset to 25% (", frac(1, 4), ", ", frac(3, 12), ")"},

			e.Code("b.OffsetOneFifth"),
			[]any{"Set column offset to 20% (", frac(1, 5), ")"},

			e.Code("b.Offset2"),
			[]any{"Set column offset to 16.66% (", frac(2, 12), ")"},

			e.Code("b.Offset1"),
			[]any{"Set column offset to 8.33% (", frac(1, 12), ")"},

			e.Code("b.Narrow"),
			"The column takes only the width it needs",
		),
		e.P("The width, offset and narrow behaviour may be breakpoint-based, by calling ", e.Code(".Mobile()"), " to ", e.Code(".FullHD()"), " on the modifier(s)."),
	),
).Section(
	"Bulma example: Basics", "https://bulma.io/documentation/columns/basics/",

	c.Example(
		`b.Columns(
	b.Column(e.P("First column")),
	e.P("Second column"),
	b.Column(e.P("Third column")),
	e.P("Fourth column"),
)`,
		b.Columns(
			b.Column(c.ColParagraph("First column")),
			c.ColParagraph("Second column"),
			b.Column(c.ColParagraph("Third column")),
			c.ColParagraph("Fourth column"),
		),
	),
).Section(
	"Bulma examples: Sizes", "https://bulma.io/documentation/columns/sizes/",

	c.Example(
		`b.Columns(
	b.Column(
		b.SizeFourFifths,
		e.P("is-four-fifths"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.ThreeQuarters,
		e.P("is-three-quarters"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.TwoThirds,
		e.P("is-two-thirds"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.ThreeFifths,
		e.P("is-three-fifths"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.Half,
		e.P("is-half"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.TwoFifths,
		e.P("is-two-fifths"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneThird,
		e.P("is-one-third"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
),
b.Columns(
	b.Column(
		b.OneFifth,
		e.P("is-one-fifth"),
	),
	b.Column(e.P("Auto")),
	b.Column(e.P("Auto")),
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
).Subsection(
	"12 columns system",
	"https://bulma.io/documentation/columns/sizes/",

	b.Content(e.P("Use ", e.Code("bulma.ColSize1"), " to ", e.Code("bulma.ColSize12"), ".")),
).Subsection(
	"Offset",
	"https://bulma.io/documentation/columns/sizes/#offset",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Column(
		b.Half,
		b.OffsetOneQuarter,
		e.P(
			e.Code("is-half"),
			e.Br(),
			e.Code("is-offset-one-quarter"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.ThreeFifths,
		b.OffsetOneFifth,
		e.P(
			e.Code("is-three-fifths"),
			e.Br(),
			e.Code("is-offset-one-fifth"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.Size4,
		b.Offset8,
		e.P(
			e.Code("is-4"),
			e.Br(),
			e.Code("is-offset-8"),
		),
	),
),
b.Columns(
	b.Mobile,
	b.Column(
		b.Size11,
		b.Offset1,
		e.P(
			e.Code("is-11"),
			e.Br(),
			e.Code("is-offset-1"),
		),
	),
)`,
		b.Columns(
			b.Mobile,
			b.Column(
				b.Half,
				b.OffsetOneQuarter,
				c.ColParagraph(
					e.Code("is-half"),
					e.Br(),
					e.Code("is-offset-one-quarter"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.ThreeFifths,
				b.OffsetOneFifth,
				c.ColParagraph(
					e.Code("is-three-fifths"),
					e.Br(),
					e.Code("is-offset-one-fifth"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Size4,
				b.Offset8,
				c.ColParagraph(
					e.Code("is-4"),
					e.Br(),
					e.Code("is-offset-8"),
				),
			),
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Size11,
				b.Offset1,
				c.ColParagraph(
					e.Code("is-11"),
					e.Br(),
					e.Code("is-offset-1"),
				),
			),
		),
	),
).Subsection(
	"Narrow column",
	"https://bulma.io/documentation/columns/sizes/#narrow-column",
	c.Example(
		`b.Columns(
	b.Column(
		b.Narrow,
		b.Box(
			e.Styles{"width": "200px"},
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
					e.Styles{"width": "200px"},
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
).Section(
	"Bulma examples: Responsiveness", "https://bulma.io/documentation/columns/responsiveness/",
).Subsection(
	"Mobile columns",
	"https://bulma.io/documentation/columns/responsiveness/#mobile-columns",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Column(e.P("1")),
	b.Column(e.P("2")),
	b.Column(e.P("3")),
	b.Column(e.P("4")),
)`,
		b.Columns(
			b.Mobile,
			b.Column(c.ColParagraph("1")),
			b.Column(c.ColParagraph("2")),
			b.Column(c.ColParagraph("3")),
			b.Column(c.ColParagraph("4")),
		),
	),
	c.Example(
		`b.Columns(
	b.ColDesktop,
	b.Column(e.P("1")),
	b.Column(e.P("2")),
	b.Column(e.P("3")),
	b.Column(e.P("4")),
)`,
		b.Columns(
			b.Desktop,
			b.Column(c.ColParagraph("1")),
			b.Column(c.ColParagraph("2")),
			b.Column(c.ColParagraph("3")),
			b.Column(c.ColParagraph("4")),
		),
	),
).Subsection(
	"Different column sizes per breakpoint",
	"https://bulma.io/documentation/columns/responsiveness/#different-column-sizes-per-breakpoint",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Column(
		b.ThreeQuarters.Mobile(),
		b.TwoThirds.Tablet(),
		b.Half.Desktop(),
		b.OneThird.Widescreen(),
		b.OneQuarter.FullHD(),
		e.P(
			e.Code("is-three-quarters-mobile"),
			e.Br(),
			e.Code("is-two-thirds-tablet"),
			e.Br(),
			e.Code("is-half-desktop"),
			e.Br(),
			e.Code("is-one-third-widescreen"),
			e.Br(),
			e.Code("is-one-quarter-fullhd"),
		),
	),
	b.Column(e.P("2")),
	b.Column(e.P("3")),
	b.Column(e.P("4")),
	b.Column(e.P("5")),
)`,
		b.Columns(
			b.Mobile,
			b.Column(
				b.ThreeQuarters.Mobile(),
				b.TwoThirds.Tablet(),
				b.Half.Desktop(),
				b.OneThird.Widescreen(),
				b.OneQuarter.FullHD(),
				c.ColParagraph(
					e.Code("is-three-quarters-mobile"),
					e.Br(),
					e.Code("is-two-thirds-tablet"),
					e.Br(),
					e.Code("is-half-desktop"),
					e.Br(),
					e.Code("is-one-third-widescreen"),
					e.Br(),
					e.Code("is-one-quarter-fullhd"),
				),
			),
			b.Column(c.ColParagraph("2")),
			b.Column(c.ColParagraph("3")),
			b.Column(c.ColParagraph("4")),
			b.Column(c.ColParagraph("5")),
		),
	),
).Section(
	"Bulma examples: Nesting", "https://bulma.io/documentation/columns/nesting/",

	c.Example(
		`b.Columns(
	b.Column(
		e.P(
			b.BackgroundInfo,
			"First column",
		),
		b.Columns(
			b.Mobile,
			b.Column(
				e.P(
					b.BackgroundInfo,
					"First nested column",
				),
			),
			b.Column(
				e.P(
					b.BackgroundInfo,
					"Second nested column",
				),
			),
		),
	),
	b.Column(
		e.P(
			b.BackgroundDanger,
			"Second column",
		),
		b.Columns(
			b.Mobile,
			b.Column(
				b.Half,
				e.P(
					b.BackgroundDanger,
					"50%",
				),
			),
			b.Column(
				e.P(
					b.BackgroundDanger,
					"Auto",
				),
			),
			b.Column(
				e.P(
					b.BackgroundDanger,
					"Auto",
				),
			),
		),
	),
)`,
		b.Columns(
			b.Column(
				c.ColParagraph(
					b.BackgroundInfo,
					"First column",
				),
				b.Columns(
					b.Mobile,
					b.Column(
						c.ColParagraph(
							b.BackgroundInfo,
							"First nested column",
						),
					),
					b.Column(
						c.ColParagraph(
							b.BackgroundInfo,
							"Second nested column",
						),
					),
				),
			),
			b.Column(
				c.ColParagraph(
					b.BackgroundDanger,
					"Second column",
				),
				b.Columns(
					b.Mobile,
					b.Column(
						b.Half,
						c.ColParagraph(
							b.BackgroundDanger,
							"50%",
						),
					),
					b.Column(
						c.ColParagraph(
							b.BackgroundDanger,
							"Auto",
						),
					),
					b.Column(
						c.ColParagraph(
							b.BackgroundDanger,
							"Auto",
						),
					),
				),
			),
		),
	),
).Section(
	"Bulma examples: Gap", "https://bulma.io/documentation/columns/gap/",
).Subsection(
	"Gapless",
	"https://bulma.io/documentation/columns/gap/#gapless",
	c.Example(
		`b.Columns(
	b.Gapless,
	b.Column(e.P("No gap")),
	b.Column(e.P("No gap")),
	b.Column(e.P("No gap")),
	b.Column(e.P("No gap")),
)`,
		b.Columns(
			b.Gapless,
			b.Column(c.ColParagraph("No gap")),
			b.Column(c.ColParagraph("No gap")),
			b.Column(c.ColParagraph("No gap")),
			b.Column(c.ColParagraph("No gap")),
		),
	),
	c.Example(
		`b.Columns(
	b.Gapless,
	b.Multiline,
	b.Mobile,
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.Half,
		e.P("is-half"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		b.OneQuarter,
		e.P("is-one-quarter"),
	),
	b.Column(
		e.P("Auto"),
	),
)`,
		b.Columns(
			b.Gapless,
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
).Subsection(
	"Variable gap",
	"https://bulma.io/documentation/columns/gap/#variable-gap",
	b.Content(e.P("Use ", e.Code("b.ColumnGap0"), " to ", e.Code("b.ColumnGap8"), ". The ", e.Code("is-variable"), " class is automatically set.")),
).Subsection(
	"Breakpoint based column gaps",
	"https://bulma.io/documentation/columns/gap/#breakpoint-based-column-gaps",
	c.Example(
		`b.Columns(
	b.ColumnGap1.Mobile(),
	b.ColumnGap0.Tablet(),
	b.ColumnGap3.Desktop(),
	b.ColumnGap8.Widescreen(),
	b.ColumnGap2.FullHD(),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
)`,
		b.Columns(
			b.ColumnGap1.Mobile(),
			b.ColumnGap0.Tablet(),
			b.ColumnGap3.Desktop(),
			b.ColumnGap8.Widescreen(),
			b.ColumnGap2.FullHD(),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
		),
	),
).Section(
	"Bulma examples: Options", "https://bulma.io/documentation/columns/options/",
).Subsection(
	"Vertical alignment",
	"https://bulma.io/documentation/columns/options/#vertical-alignment",
	c.Example(
		`b.Columns(
		b.VCentered,
		b.Column(
			b.Size8,
			e.P("First column"),
		),
		b.Column(
			e.P("Second column with more content. This is so you can see the vertical alignment."),
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
).Subsection(
	"Multiline",
	"https://bulma.io/documentation/columns/options/#multiline",
	c.Example(
		`b.Columns(
		b.Multiline,
		b.Mobile,
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.Half,
			e.P("is-half"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			b.OneQuarter,
			e.P("is-one-quarter"),
		),
		b.Column(
			e.P("Auto"),
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
).Subsection(
	"Centering columns",
	"https://bulma.io/documentation/columns/options/#centering-columns",
	c.Example(
		`b.Columns(
		b.Mobile,
		b.Centered,
		b.Column(
			b.Half,
			e.P(e.Code("is-half")),
		),
	)`,
		b.Columns(
			b.Mobile,
			b.Centered,
			b.Column(
				b.Half,
				c.ColParagraph(e.Code("is-half")),
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
			e.P(
				e.Code("is-narrow"),
				e.Br(),
				"First column",
			),
		),
		b.Column(
			b.Narrow,
			e.P(
				e.Code("is-narrow"),
				e.Br(),
				"Our second column",
			),
		),
		b.Column(
			b.Narrow,
			e.P(
				e.Code("is-narrow"),
				e.Br(),
				"Third column",
			),
		),
		b.Column(
			b.Narrow,
			e.P(
				e.Code("is-narrow"),
				e.Br(),
				"The fourth column",
			),
		),
		b.Column(
			b.Narrow,
			e.P(
				e.Code("is-narrow"),
				e.Br(),
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
					e.Code("is-narrow"),
					e.Br(),
					"First column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					e.Code("is-narrow"),
					e.Br(),
					"Our second column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					e.Code("is-narrow"),
					e.Br(),
					"Third column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					e.Code("is-narrow"),
					e.Br(),
					"The fourth column",
				),
			),
			b.Column(
				b.Narrow,
				c.ColParagraph(
					e.Code("is-narrow"),
					e.Br(),
					"Fifth column",
				),
			),
		),
	),
)
