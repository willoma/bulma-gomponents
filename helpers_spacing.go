package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var spacing = c.NewPage(
	"Spacing", "Spacing helpers", "/spacing",
	"https://bulma.io/documentation/helpers/spacing-helpers/",
	b.Content(e.P("Spacing is applied with the ", e.Code("b.Margin*"), " and ", e.Code("b.Padding*"), " functions, associated with a spacing value:")),
	b.Columns(
		c.Table(
			b.HeadRow("Property", "Function"),
			c.RowNoCode("margin", e.Code("b.Margin")),
			c.RowNoCode("margin-top", e.Code("b.MarginTop")),
			c.RowNoCode("margin-right", e.Code("b.MarginRight")),
			c.RowNoCode("margin-bottom", e.Code("b.MarginBottom")),
			c.RowNoCode("margin-left", e.Code("b.MarginLeft")),
			c.RowNoCode("margin-left and margin-right", e.Code("b.MarginHorizontal")),
			c.RowNoCode("margin-top and margin-bottom", e.Code("b.MarginVertical")),
			c.RowNoCode("padding", e.Code("b.Padding")),
			c.RowNoCode("padding-top", e.Code("b.PaddingTop")),
			c.RowNoCode("padding-right", e.Code("b.PaddingRight")),
			c.RowNoCode("padding-bottom", e.Code("b.PaddingBottom")),
			c.RowNoCode("padding-left", e.Code("b.PaddingLeft")),
			c.RowNoCode("padding-left and padding-right", e.Code("b.PaddingHorizontal")),
			c.RowNoCode("padding-top and padding-bottom", e.Code("b.PaddingVertical")),
		),
		c.Table(
			b.HeadRow("Spacing value", "Value"),
			c.Row("b.Spacing0", "0"),
			c.Row("b.Spacing1", "0.25rem"),
			c.Row("b.Spacing2", "0.5rem"),
			c.Row("b.Spacing3", "0.75rem"),
			c.Row("b.Spacing4", "1rem"),
			c.Row("b.Spacing5", "1.5rem"),
			c.Row("b.Spacing6", "3rem"),
			c.RowDefault("automatic spacing"),
		),
	),
).Section(
	"Configuration",
	"https://bulma.io/documentation/helpers/spacing-helpers/#configuration",
	b.Content(e.P("Classes names and spacing values are ", e.Strong("not"), " configurable with ", e.Em("Bulma-Gomponents"), ".")),
)
