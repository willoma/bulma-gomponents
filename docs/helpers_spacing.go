package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var spacing = c.NewPage(
	"Spacing", "Spacing helpers", "/spacing",
	"https://bulma.io/documentation/helpers/spacing-helpers/",
	b.Content(el.P("Spacing is applied with the ", el.Code("b.Margin*"), " and ", el.Code("b.Padding*"), " functions, associated with ", el.Code("b.Spacing"), " constants:")),
	b.Columns(
		b.Table(
			b.HeadRow("Property", "Function"),
			b.Row("margin", el.Code("b.Margin")),
			b.Row("margin-top", el.Code("b.MarginTop")),
			b.Row("margin-right", el.Code("b.MarginRight")),
			b.Row("margin-bottom", el.Code("b.MarginBottom")),
			b.Row("margin-left", el.Code("b.MarginLeft")),
			b.Row("margin-left and margin-right", el.Code("b.MarginHorizontal")),
			b.Row("margin-top and margin-bottom", el.Code("b.MarginVertical")),
			b.Row("padding", el.Code("b.Padding")),
			b.Row("padding-top", el.Code("b.PaddingTop")),
			b.Row("padding-right", el.Code("b.PaddingRight")),
			b.Row("padding-bottom", el.Code("b.PaddingBottom")),
			b.Row("padding-left", el.Code("b.PaddingLeft")),
			b.Row("padding-left and padding-right", el.Code("b.PaddingHorizontal")),
			b.Row("padding-top and padding-bottom", el.Code("b.PaddingVertical")),
		),
		b.Table(
			b.HeadRow("Spacing constant", "Value"),
			b.Row(el.Code("b.Spacing0"), "0"),
			b.Row(el.Code("b.Spacing1"), "0.25rem"),
			b.Row(el.Code("b.Spacing2"), "0.5rem"),
			b.Row(el.Code("b.Spacing3"), "0.75rem"),
			b.Row(el.Code("b.Spacing4"), "1rem"),
			b.Row(el.Code("b.Spacing5"), "1.5rem"),
			b.Row(el.Code("b.Spacing6"), "3rem"),
		),
	),
).Section(
	"Configuration",
	"https://bulma.io/documentation/helpers/spacing-helpers/#configuration",
	b.Content(el.P("Classes names and spacing values are ", el.Strong("not"), " configurable with ", el.Em("Bulma-Gomponents"), ".")),
)
