package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var columnsGap = c.NewPage(
	"Gap", "Columns gap", "/columns/gap",
	"https://bulma.io/documentation/columns/gap/",
).Section(
	"Gapless",
	"https://bulma.io/documentation/columns/gap/#gapless",
	c.Example(
		`b.Columns(
	b.Gapless,
	b.Column(el.P("No gap")),
	b.Column(el.P("No gap")),
	b.Column(el.P("No gap")),
	b.Column(el.P("No gap")),
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
).Section(
	"Variable gap",
	"https://bulma.io/documentation/columns/gap/#variable-gap",
	b.Content(el.P("Use ", el.Code("b.Gap0"), " to ", el.Code("b.Gap8"), ". The ", el.Code("is-variable"), " class is automatically set.")),
).Section(
	"Breakpoint based column gaps",
	"https://bulma.io/documentation/columns/gap/#breakpoint-based-column-gaps",
	c.Example(
		`b.Columns(
	b.Gap1.Mobile(),
	b.Gap0.Tablet(),
	b.Gap3.Desktop(),
	b.Gap8.Widescreen(),
	b.Gap2.FullHD(),
	b.Column(el.P("Column")),
	b.Column(el.P("Column")),
	b.Column(el.P("Column")),
	b.Column(el.P("Column")),
	b.Column(el.P("Column")),
	b.Column(el.P("Column")),
)`,
		b.Columns(
			b.Gap1.Mobile(),
			b.Gap0.Tablet(),
			b.Gap3.Desktop(),
			b.Gap8.Widescreen(),
			b.Gap2.FullHD(),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
			b.Column(c.ColParagraph("Column")),
		),
	),
)
