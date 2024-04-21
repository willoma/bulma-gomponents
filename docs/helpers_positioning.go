package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var positioning = c.NewPage(
	"Positioning", "Positioning helpers", "/positioning",
	"",

	b.Content("The following positioning modifiers are available:"),
	b.Table(
		b.HeadRow("Modifier", "Action"),
		b.Row(
			el.Code("b.PositionAbsolute"),
			b.Cell("Remove the element from the document flow and position it relative to its closed positioned ancestor (", el.Em("absolute positioning"), ")"),
		),
		b.Row(
			el.Code("b.PositionFixed"),
			b.Cell("Remove the element from the document flow and position it relative to the viewport (", el.Em("fixed positioning"), ")"),
		),
		b.Row(
			el.Code("b.PositionRelative"),
			b.Cell("Position the element according to the document flow and offset relative to itself (", el.Em("relative positioning"), ")"),
		),
		b.Row(
			el.Code("b.PositionStatic"),
			b.Cell("Position the element according to the document flow with no offset (", el.Em("static positioning"), ")"),
		),
		b.Row(
			el.Code("b.PositionSticky"),
			b.Cell("Position the element according to the document flow and offset relative to its nearest scrolling ancestor (", el.Em("sticky positioning"), ")"),
		),
	),
	b.Notification(
		b.Warning,
		"As of 2024-04-20, the official Bulma documentation has a bug and does not provide a page for positioning helpers.",
	),
)
