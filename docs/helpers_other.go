package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var other = c.NewPage(
	"Other", "Other helpers", "/other",
	"https://bulma.io/documentation/helpers/other-helpers/",
	b.Table(
		b.HeadRow("Modifier", "Action"),
		b.Row(el.Code("b.Clearfix"), b.Cell("Fixes an element's floating children")),
		b.Row(el.Code("b.PulledLeft"), b.Cell("Moves an element to the ", el.Strong("left"))),
		b.Row(el.Code("b.PulledRight"), b.Cell("Moves an element to the ", el.Strong("right"))),
		b.Row(el.Code("b.Overlay"), b.Cell("Completely ", el.Strong("covers"), " the first positioned parent")),
		b.Row(el.Code("b.Clipped"), b.Cell("Adds overflow ", el.Strong("hidden"))),
		b.Row(el.Code("b.Radiusless"), b.Cell("Removes any ", el.Strong("radius"))),
		b.Row(el.Code("b.Shadowless"), b.Cell("Removes any ", el.Strong("shadow"))),
		b.Row(el.Code("b.Unselectable"), b.Cell("Prevents the text from being ", el.Strong("selectable"))),
		b.Row(el.Code("b.Clickable"), b.Cell("Applies ", el.Code("cursor: pointer !important"), " to the element")),
		b.Row(el.Code("b.Relative"), b.Cell("Applies ", el.Code("position: relative"), " to the element")),
	),
)
