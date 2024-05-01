package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var flexbox = c.NewPage(
	"Flexbox", "Flexbox helpers", "/flexbox",
	"https://bulma.io/documentation/helpers/flexbox-helpers/",
).Section(
	"Flex direction",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-direction",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.FlexRow"), e.Code("flex-direction: row")),
		b.Row(e.Code("b.FlexRowReverse"), e.Code("flex-direction: row-reverse")),
		b.Row(e.Code("b.FlexColumn"), e.Code("flex-direction: column")),
		b.Row(e.Code("b.FlexColumnReverse"), e.Code("flex-direction: column-reverse")),
	),
).Section(
	"Flex wrap",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-wrap",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.FlexNowrap"), e.Code("flex-wrap: nowrap")),
		b.Row(e.Code("b.FlexWrap"), e.Code("flex-wrap: wrap")),
		b.Row(e.Code("b.FlexWrapReverse"), e.Code("flex-wrap: wrap-reverse")),
	),
).Section(
	"Justify content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#justify-content",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.JustifyContentFlexStart"), e.Code("justify-content: flex-start")),
		b.Row(e.Code("b.JustifyContentFlexEnd"), e.Code("justify-content: flex-end")),
		b.Row(e.Code("b.JustifyContentCenter"), e.Code("justify-content: center")),
		b.Row(e.Code("b.JustifyContentSpaceBetween"), e.Code("justify-content: space-between")),
		b.Row(e.Code("b.JustifyContentSpaceAround"), e.Code("justify-content: space-around")),
		b.Row(e.Code("b.JustifyContentSpaceEvenly"), e.Code("justify-content: space-evenly")),
		b.Row(e.Code("b.JustifyContentStart"), e.Code("justify-content: start")),
		b.Row(e.Code("b.JustifyContentEnd"), e.Code("justify-content: end")),
		b.Row(e.Code("b.JustifyContentLeft"), e.Code("justify-content: left")),
		b.Row(e.Code("b.JustifyContentRight"), e.Code("justify-content: right")),
	),
).Section(
	"Align content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-content",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.AlignContentFlexStart"), e.Code("align-content: flex-start")),
		b.Row(e.Code("b.AlignContentFlexEnd"), e.Code("align-content: flex-end")),
		b.Row(e.Code("b.AlignContentCenter"), e.Code("align-content: center")),
		b.Row(e.Code("b.AlignContentSpaceBetween"), e.Code("align-content: space-between")),
		b.Row(e.Code("b.AlignContentSpaceAround"), e.Code("align-content: space-around")),
		b.Row(e.Code("b.AlignContentSpaceEvenly"), e.Code("align-content: space-evenly")),
		b.Row(e.Code("b.AlignContentStretch"), e.Code("align-content: stretch")),
		b.Row(e.Code("b.AlignContentStart"), e.Code("align-content: start")),
		b.Row(e.Code("b.AlignContentEnd"), e.Code("align-content: end")),
		b.Row(e.Code("b.AlignContentBaseline"), e.Code("align-content: baseline")),
	),
).Section(
	"Align items",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-items",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.AlignItemsStretch"), e.Code("align-items: stretch")),
		b.Row(e.Code("b.AlignItemsFlexStart"), e.Code("align-items: flex-start")),
		b.Row(e.Code("b.AlignItemsFlexEnd"), e.Code("align-items: flex-end")),
		b.Row(e.Code("b.AlignItemsCenter"), e.Code("align-items: center")),
		b.Row(e.Code("b.AlignItemsBaseline"), e.Code("align-items: baseline")),
		b.Row(e.Code("b.AlignItemsStart"), e.Code("align-items: start")),
		b.Row(e.Code("b.AlignItemsEnd"), e.Code("align-items: end")),
		b.Row(e.Code("b.AlignItemsSelfStart"), e.Code("align-items: self-start")),
		b.Row(e.Code("b.AlignItemsSelfEnd"), e.Code("align-items: self-end")),
	),
).Section(
	"Align self",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-self",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.AlignSelfAuto"), e.Code("align-self: audo")),
		b.Row(e.Code("b.AlignSelfFlexStart"), e.Code("align-self: flex-start")),
		b.Row(e.Code("b.AlignSelfFlexEnd"), e.Code("align-self: flex-end")),
		b.Row(e.Code("b.AlignSelfCenter"), e.Code("align-self: center")),
		b.Row(e.Code("b.AlignSelfBaseline"), e.Code("align-self: baseline")),
		b.Row(e.Code("b.AlignSelfStretch"), e.Code("align-self: stretch")),
	),
).Section(
	"Flex grow and flex shrink",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-grow-and-flex-shrink",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(e.Code("b.FlexGrow0"), e.Code("flex-grow: 0")),
		b.Row(e.Code("b.FlexGrow1"), e.Code("flex-grow: 1")),
		b.Row(e.Code("b.FlexGrow2"), e.Code("flex-grow: 2")),
		b.Row(e.Code("b.FlexGrow3"), e.Code("flex-grow: 3")),
		b.Row(e.Code("b.FlexGrow4"), e.Code("flex-grow: 4")),
		b.Row(e.Code("b.FlexGrow5"), e.Code("flex-grow: 5")),
		b.Row(e.Code("b.FlexShrink0"), e.Code("flex-shrink: 0")),
		b.Row(e.Code("b.FlexShrink1"), e.Code("flex-shrink: 1")),
		b.Row(e.Code("b.FlexShrink2"), e.Code("flex-shrink: 2")),
		b.Row(e.Code("b.FlexShrink3"), e.Code("flex-shrink: 3")),
		b.Row(e.Code("b.FlexShrink4"), e.Code("flex-shrink: 4")),
		b.Row(e.Code("b.FlexShrink5"), e.Code("flex-shrink: 5")),
	),
)
