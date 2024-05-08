package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var spacing = c.NewPage(
	"Spacing", "Spacing helpers", "/spacing",
	"https://bulma.io/documentation/helpers/spacing-helpers/",
	b.Content(e.P("Spacing is applied with the ", e.Code("b.Margin*"), " and ", e.Code("b.Padding*"), " functions, associated with ", e.Code("b.Spacing"), " constants:")),
	b.Columns(
		c.Table(
			b.HeadRow("Property", "Function"),
			b.Row("margin", e.Code("b.Margin")),
			b.Row("margin-top", e.Code("b.MarginTop")),
			b.Row("margin-right", e.Code("b.MarginRight")),
			b.Row("margin-bottom", e.Code("b.MarginBottom")),
			b.Row("margin-left", e.Code("b.MarginLeft")),
			b.Row("margin-left and margin-right", e.Code("b.MarginHorizontal")),
			b.Row("margin-top and margin-bottom", e.Code("b.MarginVertical")),
			b.Row("padding", e.Code("b.Padding")),
			b.Row("padding-top", e.Code("b.PaddingTop")),
			b.Row("padding-right", e.Code("b.PaddingRight")),
			b.Row("padding-bottom", e.Code("b.PaddingBottom")),
			b.Row("padding-left", e.Code("b.PaddingLeft")),
			b.Row("padding-left and padding-right", e.Code("b.PaddingHorizontal")),
			b.Row("padding-top and padding-bottom", e.Code("b.PaddingVertical")),
		),
		c.Table(
			b.HeadRow("Spacing constant", "Value"),
			b.Row(e.Code("b.Spacing0"), "0"),
			b.Row(e.Code("b.Spacing1"), "0.25rem"),
			b.Row(e.Code("b.Spacing2"), "0.5rem"),
			b.Row(e.Code("b.Spacing3"), "0.75rem"),
			b.Row(e.Code("b.Spacing4"), "1rem"),
			b.Row(e.Code("b.Spacing5"), "1.5rem"),
			b.Row(e.Code("b.Spacing6"), "3rem"),
		),
	),
).Section(
	"Configuration",
	"https://bulma.io/documentation/helpers/spacing-helpers/#configuration",
	b.Content(e.P("Classes names and spacing values are ", e.Strong("not"), " configurable with ", e.Em("Bulma-Gomponents"), ".")),
)
