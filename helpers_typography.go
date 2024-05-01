package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var typography = c.NewPage(
	"Typography", "Typography helpers", "/typography",
	"https://bulma.io/documentation/helpers/typography-helpers/",

	b.Content(
		"Change the size, weight, and other font properties of the text",
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/helpers/typography-helpers/#size",
	c.Example(
		`e.P(b.FontSize1, "b.FontSize1"),
e.P(b.FontSize2, "b.FontSize2"),
e.P(b.FontSize3, "b.FontSize3"),
e.P(b.FontSize4, "b.FontSize4"),
e.P(b.FontSize5, "b.FontSize5"),
e.P(b.FontSize6, "b.FontSize6"),
e.P(b.FontSize7, "b.FontSize7")`,
		e.P(b.FontSize1, "b.FontSize1"),
		e.P(b.FontSize2, "b.FontSize2"),
		e.P(b.FontSize3, "b.FontSize3"),
		e.P(b.FontSize4, "b.FontSize4"),
		e.P(b.FontSize5, "b.FontSize5"),
		e.P(b.FontSize6, "b.FontSize6 (default size)"),
		e.P(b.FontSize7, "b.FontSize7"),
	),
).Section(
	"Responsive size",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-size",
	b.Content(
		e.P("You can choose a ", e.Strong("specific"), " size for ", e.Em("each"), " viewport width, by calling one of the following functions on the size argument:"),
		b.UList(
			e.Code(".Mobile()"),
			e.Code(".Tablet()"),
			e.Code(".Desktop()"),
			e.Code(".Widescreen()"),
			e.Code(".FullHD()"),
		),
	),
	c.Example(
		`e.P(
	b.FontSize5.Mobile(),
	b.FontSize4.Tablet(),
	b.FontSize3.Desktop(),
	b.FontSize2.Widescreen(),
	b.FontSize1.FullHD(),
	"Example",
)`,
		e.P(
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
		`e.P(b.TextCentered, "Centered"),
e.P(b.TextJustified, "Justified"),
e.P(b.TextLeft, "Left"),
e.P(b.TextRight, "Right")`,
		e.P(b.TextCentered, "Centered"),
		e.P(b.TextJustified, "Justified"),
		e.P(b.TextLeft, "Left"),
		e.P(b.TextRight, "Right"),
	),
).Section(
	"Responsive Alignment",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-alignment",
	b.Content(
		e.P("You can ", e.Strong("align text"), " differently for each ", e.Strong("viewport width"), " viewport width, by calling one of the following functions on the size argument:"),
		b.UList(
			e.Code(".Mobile()"),
			e.Code(".Touch()"),
			e.Code(".TabletOnly()"),
			e.Code(".Tablet()"),
			e.Code(".DesktopOnly()"),
			e.Code(".Desktop()"),
			e.Code(".WidescreenOnly()"),
			e.Code(".Widescreen()"),
			e.Code(".FullHD()"),
		),
	),
).Section(
	"Text transformation",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-transformation",
	c.Example(
		`b.Table(
	b.HeadRow("Modifier", "Transformation"),
	b.Row(e.Code("b.Capitalized"), e.P(b.Capitalized, "First character of each word uppercased")),
	b.Row(e.Code("b.Lowercase"), e.P(b.Lowercase, "TRANSFORM aLL cHaRaCtErs to LOWERcase")),
	b.Row(e.Code("b.Uppercase"), e.P(b.Uppercase, "TRANSFORM aLL cHaRaCtErs to UPPERcase")),
	b.Row(e.Code("b.Italic"), e.P(b.Italic, "Transform all characters to italic")),
	b.Row(e.Code("b.Underlined"), e.P(b.Underlined, "Underline the text")),
)`,
		b.Table(
			b.HeadRow("Modifier", "Transformation"),
			b.Row(e.Code("b.Capitalized"), e.P(b.Capitalized, "First character of each word uppercased")),
			b.Row(e.Code("b.Lowercase"), e.P(b.Lowercase, "TRANSFORM aLL cHaRaCtErs to LOWERcase")),
			b.Row(e.Code("b.Uppercase"), e.P(b.Uppercase, "TRANSFORM aLL cHaRaCtErs to UPPERcase")),
			b.Row(e.Code("b.Italic"), e.P(b.Italic, "Transform all characters to italic")),
			b.Row(e.Code("b.Underlined"), e.P(b.Underlined, "Underline the text")),
		),
	),
).Section(
	"Text weight",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-weight",
	c.Example(
		`b.Table(
	b.HeadRow("Modifier", "Transformation"),
	b.Row(e.Code("b.WeightLight"), e.P(b.WeightLight, "Transform weight to light")),
	b.Row(e.Code("b.WeightNormal"), e.P(b.WeightNormal, "Transform weight to normal")),
	b.Row(e.Code("b.WeightMedium"), e.P(b.WeightMedium, "Transform weight to medium")),
	b.Row(e.Code("b.SemiBold"), e.P(b.SemiBold, "Transform weight to semi-bold")),
	b.Row(e.Code("b.Bold"), e.P(b.Bold, "Transform weight to bold")),
)`,
		b.Table(
			b.HeadRow("Modifier", "Transformation"),
			b.Row(e.Code("b.WeightLight"), e.P(b.WeightLight, "Transform weight to light")),
			b.Row(e.Code("b.WeightNormal"), e.P(b.WeightNormal, "Transform weight to normal")),
			b.Row(e.Code("b.WeightMedium"), e.P(b.WeightMedium, "Transform weight to medium")),
			b.Row(e.Code("b.WeightSemiBold"), e.P(b.WeightSemiBold, "Transform weight to semi-bold")),
			b.Row(e.Code("b.WeightBold"), e.P(b.WeightBold, "Transform weight to bold")),
		),
	),
).Section(
	"Font family",
	"https://bulma.io/documentation/helpers/typography-helpers/#font-family",
	c.Example(
		`b.Table(
	b.HeadRow("Modifier", "Transformation"),
	b.Row(e.Code("b.SansSerif"), e.P(b.SansSerif, "Set font family to sans serif")),
	b.Row(e.Code("b.Monospace"), e.P(b.Monospace, "Set font family to monospace")),
	b.Row(e.Code("b.FamilyPrimary"), e.P(b.FamilyPrimary, "Set font family to primary")),
	b.Row(e.Code("b.FamilySecondary"), e.P(b.FamilySecondary, "Set font family to secondary")),
	b.Row(e.Code("b.Code"), e.P(b.Code, "Set font family to code")),
)`,
		b.Table(
			b.HeadRow("Modifier", "Transformation"),
			b.Row(e.Code("b.SansSerif"), e.P(b.SansSerif, "Set font family to sans serif")),
			b.Row(e.Code("b.Monospace"), e.P(b.Monospace, "Set font family to monospace")),
			b.Row(e.Code("b.FamilyPrimary"), e.P(b.FamilyPrimary, "Set font family to primary")),
			b.Row(e.Code("b.FamilySecondary"), e.P(b.FamilySecondary, "Set font family to secondary")),
			b.Row(e.Code("b.Code"), e.P(b.Code, "Set font family to code")),
		),
	),
)
