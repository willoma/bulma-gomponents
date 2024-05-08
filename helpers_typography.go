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
		"Change the size, weight, and other font properties of the text.",
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/helpers/typography-helpers/#size",
	c.Table(
		b.HeadRow("Modifier", "Font size"),
		c.Row("b.FontSize(1)", e.Span(b.FontSize(1), "Set font size to 1"), " (3 rem)"),
		c.Row("b.FontSize(2)", e.Span(b.FontSize(2), "Set font size to 2"), " (2.5 rem)"),
		c.Row("b.FontSize(3)", e.Span(b.FontSize(3), "Set font size to 3"), " (2 rem)"),
		c.Row("b.FontSize(4)", e.Span(b.FontSize(4), "Set font size to 4"), " (1.5 rem)"),
		c.Row("b.FontSize(5)", e.Span(b.FontSize(5), "Set font size to 5"), " (1.25 rem)"),
		c.Row("b.FontSize(6)", e.Span(b.FontSize(6), "Set font size to 6 (default size)"), " (1 rem)"),
		c.Row("b.FontSize(7)", e.Span(b.FontSize(7), "Set font size to 7"), " (0.75 rem)"),
	),
).Section(
	"Responsive size",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-size",
	b.Content(
		e.P("You can choose a ", e.Strong("specific"), " size for ", e.Em("each"), " viewport width, by calling one of the following methods on the size modifier:"),
		c.Table(
			b.HeadRow("Method", "Action"),
			c.Row(".Mobile()", "Apply on mobile"),
			c.Row(".Touch()", "Apply on mobile and tablet"),
			c.Row(".Tablet()", "Apply on tablet and larger"),
			c.Row(".Desktop()", "Apply on desktop and larger"),
			c.Row(".Widescreen()", "Apply on widescreen and full HD"),
			c.Row(".FullHD()", "Apply on full HD"),
		),
	),
).Section(
	"Alignment",
	"https://bulma.io/documentation/helpers/typography-helpers/#alignment",
	c.Table(
		b.HeadRow("Modifier", "Alignment"),
		c.Row("b.TextCentered", "Centered"),
		c.Row("b.TextJustified", "Justified"),
		c.Row("b.TextLeft", "Left"),
		c.Row("b.TextRight", "Right"),
	),
).Section(
	"Responsive Alignment",
	"https://bulma.io/documentation/helpers/typography-helpers/#responsive-alignment",
	b.Content(
		e.P("You can ", e.Strong("align text"), " differently for each ", e.Strong("viewport width"), " viewport width, by calling one of the following methods on the size modifier:"),
		c.Table(
			b.HeadRow("Method", "Action"),
			c.Row(".Mobile()", "Apply on mobile"),
			c.Row(".Touch()", "Apply on mobile and tablet"),
			c.Row(".TabletOnly()", "Apply on tablet"),
			c.Row(".Tablet()", "Apply on tablet and larger"),
			c.Row(".DesktopOnly()", "Apply on desktop"),
			c.Row(".Desktop()", "Apply on desktop and larger"),
			c.Row(".WidescreenOnly()", "Apply on widescreen"),
			c.Row(".Widescreen()", "Apply on widescreen and full HD"),
			c.Row(".FullHD()", "Apply on full HD"),
		),
	),
).Section(
	"Text transformation",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-transformation",
	c.Table(
		b.HeadRow("Modifier", "Transformation"),
		c.Row("b.Capitalized", b.Capitalized, "First character of each word uppercased"),
		c.Row("b.Lowercase", b.Lowercase, "TRANSFORM aLL cHaRaCtErs to LOWERcase"),
		c.Row("b.Uppercase", b.Uppercase, "TRANSFORM aLL cHaRaCtErs to UPPERcase"),
		c.Row("b.Italic", b.Italic, "Transform all characters to italic"),
		c.Row("b.Underlined", b.Underlined, "Underline the text"),
	),
).Section(
	"Text weight",
	"https://bulma.io/documentation/helpers/typography-helpers/#text-weight",
	c.Table(
		b.HeadRow("Modifier", "Weight"),
		c.Row("b.WeightLight", b.WeightLight, "Transform weight to light"),
		c.Row("b.WeightNormal", b.WeightNormal, "Transform weight to normal"),
		c.Row("b.WeightMedium", b.WeightMedium, "Transform weight to medium"),
		c.Row("b.WeightSemiBold", b.WeightSemiBold, "Transform weight to semi-bold"),
		c.Row("b.WeightBold", b.WeightBold, "Transform weight to bold"),
	),
).Section(
	"Font family",
	"https://bulma.io/documentation/helpers/typography-helpers/#font-family",
	c.Table(
		b.HeadRow("Modifier", "Family"),
		c.Row("b.SansSerif", b.SansSerif, "Set font family to sans serif"),
		c.Row("b.Monospace", b.Monospace, "Set font family to monospace"),
		c.Row("b.FamilyPrimary", b.FamilyPrimary, "Set font family to primary"),
		c.Row("b.FamilySecondary", b.FamilySecondary, "Set font family to secondary"),
		c.Row("b.Code", b.Code, "Set font family to code"),
	),
)
