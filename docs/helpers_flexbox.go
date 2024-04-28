package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var flexbox = c.NewPage(
	"Flexbox", "Flexbox helpers", "/flexbox",
	"https://bulma.io/documentation/helpers/flexbox-helpers/",
).Section(
	"Flex direction",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-direction",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.FlexRow"), el.Code("flex-direction: row")),
		b.Row(el.Code("b.FlexRowReverse"), el.Code("flex-direction: row-reverse")),
		b.Row(el.Code("b.FlexColumn"), el.Code("flex-direction: column")),
		b.Row(el.Code("b.FlexColumnReverse"), el.Code("flex-direction: column-reverse")),
	),
).Section(
	"Flex wrap",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-wrap",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.FlexNowrap"), el.Code("flex-wrap: nowrap")),
		b.Row(el.Code("b.FlexWrap"), el.Code("flex-wrap: wrap")),
		b.Row(el.Code("b.FlexWrapReverse"), el.Code("flex-wrap: wrap-reverse")),
	),
).Section(
	"Justify content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#justify-content",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.JustifyContentFlexStart"), el.Code("justify-content: flex-start")),
		b.Row(el.Code("b.JustifyContentFlexEnd"), el.Code("justify-content: flex-end")),
		b.Row(el.Code("b.JustifyContentCenter"), el.Code("justify-content: center")),
		b.Row(el.Code("b.JustifyContentSpaceBetween"), el.Code("justify-content: space-between")),
		b.Row(el.Code("b.JustifyContentSpaceAround"), el.Code("justify-content: space-around")),
		b.Row(el.Code("b.JustifyContentSpaceEvenly"), el.Code("justify-content: space-evenly")),
		b.Row(el.Code("b.JustifyContentStart"), el.Code("justify-content: start")),
		b.Row(el.Code("b.JustifyContentEnd"), el.Code("justify-content: end")),
		b.Row(el.Code("b.JustifyContentLeft"), el.Code("justify-content: left")),
		b.Row(el.Code("b.JustifyContentRight"), el.Code("justify-content: right")),
	),
).Section(
	"Align content",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-content",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.AlignContentFlexStart"), el.Code("align-content: flex-start")),
		b.Row(el.Code("b.AlignContentFlexEnd"), el.Code("align-content: flex-end")),
		b.Row(el.Code("b.AlignContentCenter"), el.Code("align-content: center")),
		b.Row(el.Code("b.AlignContentSpaceBetween"), el.Code("align-content: space-between")),
		b.Row(el.Code("b.AlignContentSpaceAround"), el.Code("align-content: space-around")),
		b.Row(el.Code("b.AlignContentSpaceEvenly"), el.Code("align-content: space-evenly")),
		b.Row(el.Code("b.AlignContentStretch"), el.Code("align-content: stretch")),
		b.Row(el.Code("b.AlignContentStart"), el.Code("align-content: start")),
		b.Row(el.Code("b.AlignContentEnd"), el.Code("align-content: end")),
		b.Row(el.Code("b.AlignContentBaseline"), el.Code("align-content: baseline")),
	),
).Section(
	"Align items",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-items",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.AlignItemsStretch"), el.Code("align-items: stretch")),
		b.Row(el.Code("b.AlignItemsFlexStart"), el.Code("align-items: flex-start")),
		b.Row(el.Code("b.AlignItemsFlexEnd"), el.Code("align-items: flex-end")),
		b.Row(el.Code("b.AlignItemsCenter"), el.Code("align-items: center")),
		b.Row(el.Code("b.AlignItemsBaseline"), el.Code("align-items: baseline")),
		b.Row(el.Code("b.AlignItemsStart"), el.Code("align-items: start")),
		b.Row(el.Code("b.AlignItemsEnd"), el.Code("align-items: end")),
		b.Row(el.Code("b.AlignItemsSelfStart"), el.Code("align-items: self-start")),
		b.Row(el.Code("b.AlignItemsSelfEnd"), el.Code("align-items: self-end")),
	),
).Section(
	"Align self",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#align-self",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.AlignSelfAuto"), el.Code("align-self: audo")),
		b.Row(el.Code("b.AlignSelfFlexStart"), el.Code("align-self: flex-start")),
		b.Row(el.Code("b.AlignSelfFlexEnd"), el.Code("align-self: flex-end")),
		b.Row(el.Code("b.AlignSelfCenter"), el.Code("align-self: center")),
		b.Row(el.Code("b.AlignSelfBaseline"), el.Code("align-self: baseline")),
		b.Row(el.Code("b.AlignSelfStretch"), el.Code("align-self: stretch")),
	),
).Section(
	"Flex grow and flex shrink",
	"https://bulma.io/documentation/helpers/flexbox-helpers/#flex-grow-and-flex-shrink",
	b.Table(
		b.HeadRow("Modifier", "Property: Value"),
		b.Row(el.Code("b.FlexGrow0"), el.Code("flex-grow: 0")),
		b.Row(el.Code("b.FlexGrow1"), el.Code("flex-grow: 1")),
		b.Row(el.Code("b.FlexGrow2"), el.Code("flex-grow: 2")),
		b.Row(el.Code("b.FlexGrow3"), el.Code("flex-grow: 3")),
		b.Row(el.Code("b.FlexGrow4"), el.Code("flex-grow: 4")),
		b.Row(el.Code("b.FlexGrow5"), el.Code("flex-grow: 5")),
		b.Row(el.Code("b.FlexShrink0"), el.Code("flex-shrink: 0")),
		b.Row(el.Code("b.FlexShrink1"), el.Code("flex-shrink: 1")),
		b.Row(el.Code("b.FlexShrink2"), el.Code("flex-shrink: 2")),
		b.Row(el.Code("b.FlexShrink3"), el.Code("flex-shrink: 3")),
		b.Row(el.Code("b.FlexShrink4"), el.Code("flex-shrink: 4")),
		b.Row(el.Code("b.FlexShrink5"), el.Code("flex-shrink: 5")),
	),
)
