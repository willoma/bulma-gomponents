package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var grid = c.NewPage(
	"Grid", "Grid", "/grid",
	"",

	b.Content(
		e.P("The ", "b.Grid"), " constructor creates a smart grid.",
		c.Modifiers(
			c.RowTo("b.ColMin(1)", "b.ColMin(12)", "Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem"),
			c.RowToDetail(
				"b.Gap(0)", "b.Gap(8)", "with 0.5 steps",
				"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowToDetail(
				"b.ColGap(0)", "b.ColGap(8)", "with 0.5 steps",
				"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowToDetail(
				"b.RowGap(0)", "b.RowGap(8)", "with 0.5 steps",
				"Set row gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
		),
		c.Children(
			c.Row("b.Cell(...any)", "Add cell to the ", e.Code(`<div class="grid">`), " element"),
			c.Row("e.Element", "Wrap element in a cell and add it to the ", e.Code(`<div class="grid">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<div class="grid">`), " element"),
			c.RowNodeElement("Wrap element in a cell and add it to the ", e.Code(`<div class="grid">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="grid">`), " element"),
		),

		e.P("The ", "b.FixedGrid"), " constructor creates a fixed grid.",
		c.Modifiers(
			c.RowTo("b.ColMin(1)", "b.ColMin(12)", "Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem"),
			c.RowToDetail(
				"b.Gap(0)", "b.Gap(8)", "with 0.5 steps",
				"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowToDetail(
				"b.ColGap(0)", "b.ColGap(8)", "with 0.5 steps",
				"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowToDetail(
				"b.RowGap(0)", "b.RowGap(8)", "with 0.5 steps",
				"Set row gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowTo("b.Cols(1)", "b.Cols(12)", "Set columns count to 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 or 12"),
			c.Row(
				"b.AutoCount",
				"Set columns count to:",
				b.UList(
					b.MarginTop(0),
					"2 on mobile",
					"4 on tablet",
					"8 on desktop",
					"12 on widescreen",
					"16 on full HD",
				),
			),
		),
		c.Children(
			c.Row("b.OnFixedGrid(...any)", "Apply children to the ", e.Code(`<div class="fixed-grid">`), " element"),
			c.Row("b.OnGrid(...any)", "Apply children to the ", e.Code(`<div class="grid">`), " element"),
			c.Row("b.Cell(...any)", "Add cell to the ", e.Code(`<div class="grid">`), " element"),
			c.Row("e.Element", "Wrap element in a cell and add it to the ", e.Code(`<div class="grid">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<div class="grid">`), " element"),
			c.RowNodeElement("Wrap element in a cell and add it to the ", e.Code(`<div class="grid">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="grid">`), " element"),
		),
		e.P("The columns count may be breakpoint-based, with ", "b.Cols(...).Mobile()"), " to ", e.Code("b.Cols(...).FullHD()"), ".",

		e.P("The ", e.Code("b.Cell"), " constructor creates a single grid cell."),
		c.Modifiers(
			c.RowTo("b.ColStart(1)", "b.ColStart(12)", "Start the cell at the given column (other values are treated as ", e.Code("end"), ", ie. the last column)"),
			c.RowTo("b.ColFromEnd(1)", "b.ColFromEnd(12)", "End the cell at the given column"),
			c.RowTo("b.ColSpan(1)", "b.ColSpan(12)", "Span the cell over this number of columns"),
			c.RowTo("b.RowStart(1)", "b.RowStart(12)", "Start the cell at the given row (other values are treated as ", e.Code("end"), ", ie. the last row)"),
			c.RowTo("b.RowFromEnd(1)", "b.RowFromEnd(12)", "End the cell at the given row"),
			c.RowTo("b.RowSpan(1)", "b.RowSpan(12)", "Span the cell over this number of rows"),
		),
		e.P("The column and row start/end/span may be breakpoint-based, with ", ".Mobile()"), " to ", e.Code(".FullHD()"), ".",
	),
)
