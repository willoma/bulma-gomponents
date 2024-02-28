package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var typography = c.NewPage(
	"Typography", "Typography helpers", "/typography",
	"https://bulma.io/documentation/helpers/typography-helpers/",
).Section(
	"Sizes",
	"https://bulma.io/documentation/helpers/typography-helpers/#size",
	c.Example(
		`el.P(b.FontSize1, "b.FontSize1"),
el.P(b.FontSize2, "b.FontSize2"),
el.P(b.FontSize3, "b.FontSize3"),
el.P(b.FontSize4, "b.FontSize4"),
el.P(b.FontSize5, "b.FontSize5"),
el.P(b.FontSize6, "b.FontSize6"),
el.P(b.FontSize7, "b.FontSize7")`,
		el.P(b.FontSize1, "b.FontSize1"),
		el.P(b.FontSize2, "b.FontSize2"),
		el.P(b.FontSize3, "b.FontSize3"),
		el.P(b.FontSize4, "b.FontSize4"),
		el.P(b.FontSize5, "b.FontSize5"),
		el.P(b.FontSize6, "b.FontSize6 (default size)"),
		el.P(b.FontSize7, "b.FontSize7"),
	),
).Section(
	"Responsive size",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-size",
	b.Content(
		el.P("You can choose a ", el.Strong("specific"), " size for ", el.Em("each"), " viewport width, by calling one of the following functions on the size argument:"),
		b.UList(
			el.Code(".Mobile()"),
			el.Code(".Tablet()"),
			el.Code(".Desktop()"),
			el.Code(".Widescreen()"),
			el.Code(".FullHD()"),
		),
	),
	c.Example(
		`el.P(
	b.FontSize5.Mobile(),
	b.FontSize4.Tablet(),
	b.FontSize3.Desktop(),
	b.FontSize2.Widescreen(),
	b.FontSize1.FullHD(),
	"Example",
)`,
		el.P(
			b.FontSize5.Mobile(),
			b.FontSize4.Tablet(),
			b.FontSize3.Desktop(),
			b.FontSize2.Widescreen(),
			b.FontSize1.FullHD(),
			"Example",
		),
	),
).Section(
	"Alignment",
	"https://bulma.io/documentation/helpers/typography-helpers/#alignment",
	c.Example(
		`el.P(b.TextCentered, "Centered"),
el.P(b.TextJustified, "Justified"),
el.P(b.TextLeft, "Left"),
el.P(b.TextRight, "Right")`,
		el.P(b.TextCentered, "Centered"),
		el.P(b.TextJustified, "Justified"),
		el.P(b.TextLeft, "Left"),
		el.P(b.TextRight, "Right"),
	),
).Section(
	"Responsive Alignment",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-alignment",
	b.Content(
		el.P("You can ", el.Strong("align text"), " differently for each ", el.Strong("viewport width"), " viewport width, by calling one of the following functions on the size argument:"),
		b.UList(
			el.Code(".Mobile()"),
			el.Code(".Touch()"),
			el.Code(".TabletOnly()"),
			el.Code(".Tablet()"),
			el.Code(".DesktopOnly()"),
			el.Code(".Desktop()"),
			el.Code(".WidescreenOnly()"),
			el.Code(".Widescreen()"),
			el.Code(".FullHD()"),
		),
	),
).Section(
	"Text transformation",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-transformation",
	c.Example(
		`b.Table(
	b.HeadRow("Class", "Transformation"),
	b.Row(el.Code("b.Capitalized"), el.P(b.Capitalized, "First character of each word uppercased")),
	b.Row(el.Code("b.Lowercase"), el.P(b.Lowercase, "TRANSFORM aLL cHaRaCtErs to LOWERcase")),
	b.Row(el.Code("b.Uppercase"), el.P(b.Uppercase, "TRANSFORM aLL cHaRaCtErs to UPPERcase")),
	b.Row(el.Code("b.Italic"), el.P(b.Italic, "Transform all characters to italic")),
	b.Row(el.Code("b.Underlined"), el.P(b.Underlined, "Underline the text")),
)`,
		b.Table(
			b.HeadRow("Class", "Transformation"),
			b.Row(el.Code("b.Capitalized"), el.P(b.Capitalized, "First character of each word uppercased")),
			b.Row(el.Code("b.Lowercase"), el.P(b.Lowercase, "TRANSFORM aLL cHaRaCtErs to LOWERcase")),
			b.Row(el.Code("b.Uppercase"), el.P(b.Uppercase, "TRANSFORM aLL cHaRaCtErs to UPPERcase")),
			b.Row(el.Code("b.Italic"), el.P(b.Italic, "Transform all characters to italic")),
			b.Row(el.Code("b.Underlined"), el.P(b.Underlined, "Underline the text")),
		),
	),
).Section(
	"Text weight",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-weight",
	c.Example(
		`b.Table(
	b.HeadRow("Class", "Transformation"),
	b.Row(el.Code("b.TextLight"), el.P(b.TextLight, "Transform weight to light")),
	b.Row(el.Code("b.TextNormal"), el.P(b.TextNormal, "Transform weight to normal")),
	b.Row(el.Code("b.TextMedium"), el.P(b.TextMedium, "Transform weight to medium")),
	b.Row(el.Code("b.SemiBold"), el.P(b.SemiBold, "Transform weight to semi-bold")),
	b.Row(el.Code("b.Bold"), el.P(b.Bold, "Transform weight to bold")),
)`,
		b.Table(
			b.HeadRow("Class", "Transformation"),
			b.Row(el.Code("b.TextLight"), el.P(b.TextLight, "Transform weight to light")),
			b.Row(el.Code("b.TextNormal"), el.P(b.TextNormal, "Transform weight to normal")),
			b.Row(el.Code("b.TextMedium"), el.P(b.TextMedium, "Transform weight to medium")),
			b.Row(el.Code("b.SemiBold"), el.P(b.SemiBold, "Transform weight to semi-bold")),
			b.Row(el.Code("b.Bold"), el.P(b.Bold, "Transform weight to bold")),
		),
	),
).Section(
	"Font family",
	"https://bulma.io/documentation/helpers/typography-helpers/#font-family",
	c.Example(
		`b.Table(
	b.HeadRow("Class", "Transformation"),
	b.Row(el.Code("b.SansSerif"), el.P(b.SansSerif, "Set font family to sans serif")),
	b.Row(el.Code("b.Monospace"), el.P(b.Monospace, "Set font family to monospace")),
	b.Row(el.Code("b.TextPrimary"), el.P(b.TextPrimary, "Set font family to primary")),
	b.Row(el.Code("b.TextSecondary"), el.P(b.TextSecondary, "Set font family to secondary")),
	b.Row(el.Code("b.Code"), el.P(b.Code, "Set font family to code")),
)`,
		b.Table(
			b.HeadRow("Class", "Transformation"),
			b.Row(el.Code("b.SansSerif"), el.P(b.SansSerif, "Set font family to sans serif")),
			b.Row(el.Code("b.Monospace"), el.P(b.Monospace, "Set font family to monospace")),
			b.Row(el.Code("b.TextPrimary"), el.P(b.TextPrimary, "Set font family to primary")),
			b.Row(el.Code("b.TextSecondary"), el.P(b.TextSecondary, "Set font family to secondary")),
			b.Row(el.Code("b.Code"), el.P(b.Code, "Set font family to code")),
		),
	),
)
