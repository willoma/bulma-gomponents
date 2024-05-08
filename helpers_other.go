package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
)

var other = c.NewPage(
	"Other", "Other helpers", "/other",
	"https://bulma.io/documentation/helpers/other-helpers/",
	c.Modifiers(
		c.Row("b.Clearfix", "Fix the element's floating children"),
		c.Row("b.PulledLeft", "Move the element to the ", e.Strong("left")),
		c.Row("b.PulledRight", "Move the element to the ", e.Strong("right")),
		c.Row("b.Overlay", "Completely ", e.Strong("cover"), " the first positioned parent"),
		c.Row("b.Clipped", "Add overflow ", e.Strong("hidden")),
		c.Row("b.Radiusless", "Remove any ", e.Strong("radius")),
		c.Row("b.Shadowless", "Remove any ", e.Strong("shadow")),
		c.Row("b.Unselectable", "Prevent the text from being ", e.Strong("selectable")),
		c.Row("b.Clickable", "Apply ", e.Code("cursor: pointer !important"), " to the element"),
		c.Row("b.Relative", "Apply ", e.Code("position: relative"), " to the element"),
	),
).Section(
	"Positioning helpers",
	"",

	c.Modifiers(
		c.Row("b.PositionAbsolute", "Remove the element from the document flow and position it relative to its closed positioned ancestor (", e.Strong("absolute positioning"), ")"),
		c.Row("b.PositionFixed", "Remove the element from the document flow and position it relative to the viewport (", e.Strong("fixed positioning"), ")"),
		c.Row("b.PositionRelative", "Position the element according to the document flow and offset relative to itself (", e.Strong("relative positioning"), ")"),
		c.Row("b.PositionStatic", "Position the element according to the document flow with no offset (", e.Strong("static positioning"), ")"),
		c.Row("b.PositionSticky", "Position the element according to the document flow and offset relative to its nearest scrolling ancestor (", e.Strong("sticky positioning"), ")"),
	),
)
