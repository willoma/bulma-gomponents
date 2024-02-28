package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var columnsResponsiveness = c.NewPage(
	"Responsiveness", "Columns responsiveness", "/columns/responsiveness",
	"https://bulma.io/documentation/columns/responsiveness/",
).Section(
	"Mobile columns",
	"https://bulma.io/documentation/columns/responsiveness/#mobile-columns",
	c.Example(
		`b.Columns(
	b.Mobile,
	b.Column(el.P("1")),
	b.Column(el.P("2")),
	b.Column(el.P("3")),
	b.Column(el.P("4")),
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
	b.Column(el.P("1")),
	b.Column(el.P("2")),
	b.Column(el.P("3")),
	b.Column(el.P("4")),
)`,
		b.Columns(
			b.Desktop,
			b.Column(c.ColParagraph("1")),
			b.Column(c.ColParagraph("2")),
			b.Column(c.ColParagraph("3")),
			b.Column(c.ColParagraph("4")),
		),
	),
).Section(
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
		el.P(
			el.Code("is-three-quarters-mobile"),
			el.Br(),
			el.Code("is-two-thirds-tablet"),
			el.Br(),
			el.Code("is-half-desktop"),
			el.Br(),
			el.Code("is-one-third-widescreen"),
			el.Br(),
			el.Code("is-one-quarter-fullhd"),
		),
	),
	b.Column(el.P("2")),
	b.Column(el.P("3")),
	b.Column(el.P("4")),
	b.Column(el.P("5")),
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
					el.Code("is-three-quarters-mobile"),
					el.Br(),
					el.Code("is-two-thirds-tablet"),
					el.Br(),
					el.Code("is-half-desktop"),
					el.Br(),
					el.Code("is-one-third-widescreen"),
					el.Br(),
					el.Code("is-one-quarter-fullhd"),
				),
			),
			b.Column(c.ColParagraph("2")),
			b.Column(c.ColParagraph("3")),
			b.Column(c.ColParagraph("4")),
			b.Column(c.ColParagraph("5")),
		),
	),
)
