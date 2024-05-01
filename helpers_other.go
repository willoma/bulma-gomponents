package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var other = c.NewPage(
	"Other", "Other helpers", "/other",
	"https://bulma.io/documentation/helpers/other-helpers/",
	b.Table(
		b.HeadRow("Modifier", "Action"),
		b.Row(e.Code("b.Clearfix"), b.Cell("Fixes an e.Element's floating children")),
		b.Row(e.Code("b.PulledLeft"), b.Cell("Moves an e.Element to the ", e.Strong("left"))),
		b.Row(e.Code("b.PulledRight"), b.Cell("Moves an e.Element to the ", e.Strong("right"))),
		b.Row(e.Code("b.Overlay"), b.Cell("Completely ", e.Strong("covers"), " the first positioned parent")),
		b.Row(e.Code("b.Clipped"), b.Cell("Adds overflow ", e.Strong("hidden"))),
		b.Row(e.Code("b.Radiusless"), b.Cell("Removes any ", e.Strong("radius"))),
		b.Row(e.Code("b.Shadowless"), b.Cell("Removes any ", e.Strong("shadow"))),
		b.Row(e.Code("b.Unselectable"), b.Cell("Prevents the text from being ", e.Strong("selectable"))),
		b.Row(e.Code("b.Clickable"), b.Cell("Applies ", e.Code("cursor: pointer !important"), " to the element")),
		b.Row(e.Code("b.Relative"), b.Cell("Applies ", e.Code("position: relative"), " to the element")),
	),
)
