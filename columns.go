package docs

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
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
		e.P("The ", e.Code("b.Columns"), " constructor creates a columns container."),
		c.Modifiers(
			c.Row("b.Gapless", "No gap"),
			c.Row("b.ColumnGap(0)", "No gap"),
			c.Row("b.ColumnGap(1)", "0.25rem gap"),
			c.Row("b.ColumnGap(2)", "0.5rem gap"),
			c.Row("b.ColumnGap(3)", "0.75rem gap"),
			c.Row("b.ColumnGap(4)", "1rem gap"),
			c.Row("b.ColumnGap(5)", "1.25rem gap"),
			c.Row("b.ColumnGap(6)", "1.5rem gap"),
			c.Row("b.ColumnGap(7)", "1.75rem gap"),
			c.Row("b.ColumnGap(8)", "2rem gap"),
			c.Row("b.Centered", "Center columns"),
			c.Row("b.Desktop", "Allow columns only on desktops upward (not on tablets)"),
			c.Row("b.Mobile", "Allow columns on mobile phones too"),
			c.Row("b.Multiline", "Create a new line when columns do not fit in a single line"),
			c.Row("b.VCentered", "Center columns vertically"),
		),
		e.P("The gap may be breakpoint-based, with ", e.Code("b.ColumnGap*.Mobile()"), " to ", e.Code("b.ColumnGap*.FullHD()"), "."),
		c.Children(
			c.Row("b.Column(...any)", "Add column to the ", e.Code(`<div class="columns">`), " element"),
			c.Row("e.Element", "Wrap element in a column and add it to the ", e.Code(`<div class="columns">`), " element"),
			c.Row("string", "Wrap text in a column and add it to the ", e.Code(`<div class="columns">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<div class="columns">`), " element"),
			c.RowNodeElement("Wrap element in a column and add it to the ", e.Code(`<div class="columns">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="columns">`), " element"),
		),
		e.P("The ", e.Code("b.Column"), " constructor creates a single column."),
		c.Modifiers(
			c.Row2("b.Full", "b.Size(12)", "Set width to 100% (", frac(12, 12), ")"),
			c.Row("b.Size(11)", "Set width to 91.66% (", frac(11, 12), ")"),
			c.Row("b.Size(10)", "Set width to 83.33% (", frac(10, 12), ")"),
			c.Row("b.FourFifths", "Set width to 80% (", frac(4, 5), ")"),
			c.Row2("b.ThreeQuarters", "b.Size(9)", "Set width to 75% (", frac(3, 4), ", ", frac(9, 12), ")"),
			c.Row2("b.TwoThirds", "b.Size(8)", "Set width to 66.66% (", frac(2, 3), ", ", frac(8, 12), ")"),
			c.Row("b.ThreeFifths", "Set width to 60% (", frac(3, 5), ")"),
			c.Row("b.Size(7)", "Set width to 58.33% (", frac(7, 12), ")"),
			c.Row2("b.Half", "b.Size(6)", "Set width to 50% (", frac(1, 2), ", ", frac(6, 12), ")"),
			c.Row("b.Size(5)", "Set width to 41.66% (", frac(5, 12), ")"),
			c.Row("b.TwoFifths", "Set width to 40% (", frac(2, 5), ")"),
			c.Row2("b.OneThird", "b.Size(4)", "Set width to 33.33% (", frac(1, 3), ", ", frac(4, 12), ")"),
			c.Row2("b.OneQuarter", "b.Size(3)", "Set width to 25% (", frac(1, 4), ", ", frac(3, 12), ")"),
			c.Row("b.OneFifth", "Set width to 20% (", frac(1, 5), ")"),
			c.Row("b.Size(2)", "Set width to 16.66% (", frac(2, 12), ")"),
			c.Row("b.Size(1)", "Set width to 8.33% (", frac(1, 12), ")"),
			c.Row("b.Offset(11)", "Set offset to 91.66% (", frac(11, 12), ")"),
			c.Row("b.Offset(10)", "Set offset to 83.33% (", frac(10, 12), ")"),
			c.Row("b.OffsetFourFifths", "Set offset to 80% (", frac(4, 5), ")"),
			c.Row2("b.OffsetThreeQuarters", "b.Offset(9)", "Set offset to 75% (", frac(3, 4), ", ", frac(9, 12), ")"),
			c.Row2("b.OffsetTwoThirds", "b.Offset(8)", "Set offset to 66.66% (", frac(2, 3), ", ", frac(8, 12), ")"),
			c.Row("b.OffsetThreeFifths", "Set offset to 60% (", frac(3, 5), ")"),
			c.Row("b.Offset(7)", "Set offset to 58.33% (", frac(7, 12), ")"),
			c.Row2("b.OffsetHalf", "b.Offset(6)", "Set offset to 50% (", frac(1, 2), ", ", frac(6, 12), ")"),
			c.Row("b.Offset(5)", "Set offset to 41.66% (", frac(5, 12), ")"),
			c.Row("b.OffsetTwoFifths", "Set offset to 40% (", frac(2, 5), ")"),
			c.Row2("b.OffsetOneThird", "b.Offset(4)", "Set offset to 33.33% (", frac(1, 3), ", ", frac(4, 12), ")"),
			c.Row2("b.OffsetOneQuarter", "b.Offset(3)", "Set offset to 25% (", frac(1, 4), ", ", frac(3, 12), ")"),
			c.Row("b.OffsetOneFifth", "Set offset to 20% (", frac(1, 5), ")"),
			c.Row("b.Offset(2)", "Set offset to 16.66% (", frac(2, 12), ")"),
			c.Row("b.Offset(1)", "Set offset to 8.33% (", frac(1, 12), ")"),
			c.Row("b.Narrow", "The column takes only the width it needs"),
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
		b.Size(4),
		b.Offset(8),
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
		b.Size(11),
		b.Offset(1),
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
				b.Size(4),
				b.Offset(8),
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
				b.Size(11),
				b.Offset(1),
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
			b.Title(5, html.P, "Narrow column"),
			b.Subtitle(html.P, "This column is only 200px wide."),
		),
	),
	b.Column(
		b.Box(
			b.Title(5, html.P, "Flexible column"),
			b.Subtitle(html.P, "This column will take up the remaining space available."),
		),
	),
)`,
		b.Columns(
			b.Column(
				b.Narrow,
				b.Box(
					e.Styles{"width": "200px"},
					b.Title(5, html.P, "Narrow column"),
					b.Subtitle(html.P, "This column is only 200px wide."),
				),
			),
			b.Column(
				b.Box(
					b.Title(5, html.P, "Flexible column"),
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
	b.Content(e.P("Use ", e.Code("b.ColumnGap(0)"), " to ", e.Code("b.ColumnGap(8)"), ". The ", e.Code("is-variable"), " class is automatically set.")),
).Subsection(
	"Breakpoint based column gaps",
	"https://bulma.io/documentation/columns/gap/#breakpoint-based-column-gaps",
	c.Example(
		`b.Columns(
	b.ColumnGap(1).Mobile(),
	b.ColumnGap(0).Tablet(),
	b.ColumnGap(3).Desktop(),
	b.ColumnGap(8).Widescreen(),
	b.ColumnGap(2).FullHD(),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
	b.Column(e.P("Column")),
)`,
		b.Columns(
			b.ColumnGap(1).Mobile(),
			b.ColumnGap(0).Tablet(),
			b.ColumnGap(3).Desktop(),
			b.ColumnGap(8).Widescreen(),
			b.ColumnGap(2).FullHD(),
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
			b.Size(8),
			e.P("First column"),
		),
		b.Column(
			e.P("Second column with more content. This is so you can see the vertical alignment."),
		),
	)`,
		b.Columns(
			b.VCentered,
			b.Column(
				b.Size(8),
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
