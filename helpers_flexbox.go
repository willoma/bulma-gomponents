package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var flexbox = c.NewPage(
	"Flexbox", "Flexbox helpers", "/flexbox",
	"https://bulma.io/documentation/helpers/flexbox-helpers/",
).Section(
	"Flex direction",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-direction",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.FlexRow", e.Code("flex-direction: row")),
		c.Row("b.FlexRowReverse", e.Code("flex-direction: row-reverse")),
		c.Row("b.FlexColumn", e.Code("flex-direction: column")),
		c.Row("b.FlexColumnReverse", e.Code("flex-direction: column-reverse")),
	),
).Section(
	"Flex wrap",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-wrap",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.FlexNowrap", e.Code("flex-wrap: nowrap")),
		c.Row("b.FlexWrap", e.Code("flex-wrap: wrap")),
		c.Row("b.FlexWrapReverse", e.Code("flex-wrap: wrap-reverse")),
	),
).Section(
	"Justify content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#justify-content",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.JustifyContentFlexStart", e.Code("justify-content: flex-start")),
		c.Row("b.JustifyContentFlexEnd", e.Code("justify-content: flex-end")),
		c.Row("b.JustifyContentCenter", e.Code("justify-content: center")),
		c.Row("b.JustifyContentSpaceBetween", e.Code("justify-content: space-between")),
		c.Row("b.JustifyContentSpaceAround", e.Code("justify-content: space-around")),
		c.Row("b.JustifyContentSpaceEvenly", e.Code("justify-content: space-evenly")),
		c.Row("b.JustifyContentStart", e.Code("justify-content: start")),
		c.Row("b.JustifyContentEnd", e.Code("justify-content: end")),
		c.Row("b.JustifyContentLeft", e.Code("justify-content: left")),
		c.Row("b.JustifyContentRight", e.Code("justify-content: right")),
	),
).Section(
	"Align content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-content",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.AlignContentFlexStart", e.Code("align-content: flex-start")),
		c.Row("b.AlignContentFlexEnd", e.Code("align-content: flex-end")),
		c.Row("b.AlignContentCenter", e.Code("align-content: center")),
		c.Row("b.AlignContentSpaceBetween", e.Code("align-content: space-between")),
		c.Row("b.AlignContentSpaceAround", e.Code("align-content: space-around")),
		c.Row("b.AlignContentSpaceEvenly", e.Code("align-content: space-evenly")),
		c.Row("b.AlignContentStretch", e.Code("align-content: stretch")),
		c.Row("b.AlignContentStart", e.Code("align-content: start")),
		c.Row("b.AlignContentEnd", e.Code("align-content: end")),
		c.Row("b.AlignContentBaseline", e.Code("align-content: baseline")),
	),
).Section(
	"Align items",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-items",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.AlignItemsStretch", e.Code("align-items: stretch")),
		c.Row("b.AlignItemsFlexStart", e.Code("align-items: flex-start")),
		c.Row("b.AlignItemsFlexEnd", e.Code("align-items: flex-end")),
		c.Row("b.AlignItemsCenter", e.Code("align-items: center")),
		c.Row("b.AlignItemsBaseline", e.Code("align-items: baseline")),
		c.Row("b.AlignItemsStart", e.Code("align-items: start")),
		c.Row("b.AlignItemsEnd", e.Code("align-items: end")),
		c.Row("b.AlignItemsSelfStart", e.Code("align-items: self-start")),
		c.Row("b.AlignItemsSelfEnd", e.Code("align-items: self-end")),
	),
).Section(
	"Align self",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-self",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.AlignSelfAuto", e.Code("align-self: audo")),
		c.Row("b.AlignSelfFlexStart", e.Code("align-self: flex-start")),
		c.Row("b.AlignSelfFlexEnd", e.Code("align-self: flex-end")),
		c.Row("b.AlignSelfCenter", e.Code("align-self: center")),
		c.Row("b.AlignSelfBaseline", e.Code("align-self: baseline")),
		c.Row("b.AlignSelfStretch", e.Code("align-self: stretch")),
	),
).Section(
	"Flex grow and flex shrink",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-grow-and-flex-shrink",
	c.Table(
		b.HeadRow("Modifier", "Property: Value"),
		c.Row("b.FlexGrow(0)", e.Code("flex-grow: 0")),
		c.Row("b.FlexGrow(1)", e.Code("flex-grow: 1")),
		c.Row("b.FlexGrow(2)", e.Code("flex-grow: 2")),
		c.Row("b.FlexGrow(3)", e.Code("flex-grow: 3")),
		c.Row("b.FlexGrow(4)", e.Code("flex-grow: 4")),
		c.Row("b.FlexGrow(5())", e.Code("flex-grow: 5")),
		c.Row("b.FlexShrink(0)", e.Code("flex-shrink: 0")),
		c.Row("b.FlexShrink(1)", e.Code("flex-shrink: 1")),
		c.Row("b.FlexShrink(2)", e.Code("flex-shrink: 2")),
		c.Row("b.FlexShrink(3)", e.Code("flex-shrink: 3")),
		c.Row("b.FlexShrink(4)", e.Code("flex-shrink: 4")),
		c.Row("b.FlexShrink(5)", e.Code("flex-shrink: 5")),
	),
)
