package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var positioning = c.NewPage(
	"Positioning", "Positioning helpers", "/positioning",
	"",

	b.Content("The following positioning modifiers are available:"),
	b.Table(
		b.HeadRow("Modifier", "Action"),
		b.Row(
			e.Code("b.PositionAbsolute"),
			b.Cell("Remove the element from the document flow and position it relative to its closed positioned ancestor (", e.Em("absolute positioning"), ")"),
		),
		b.Row(
			e.Code("b.PositionFixed"),
			b.Cell("Remove the element from the document flow and position it relative to the viewport (", e.Em("fixed positioning"), ")"),
		),
		b.Row(
			e.Code("b.PositionRelative"),
			b.Cell("Position the element according to the document flow and offset relative to itself (", e.Em("relative positioning"), ")"),
		),
		b.Row(
			e.Code("b.PositionStatic"),
			b.Cell("Position the element according to the document flow with no offset (", e.Em("static positioning"), ")"),
		),
		b.Row(
			e.Code("b.PositionSticky"),
			b.Cell("Position the element according to the document flow and offset relative to its nearest scrolling ancestor (", e.Em("sticky positioning"), ")"),
		),
	),
	b.Notification(
		b.Warning,
		"As of 2024-04-20, the official Bulma documentation has a bug and does not provide a page for positioning helpers.",
	),
)
