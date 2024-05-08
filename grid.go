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
			c.RowTo("b.ColMin1", "b.ColMin12", "Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem"),
			c.RowMultiple(
				[]string{"b.Gap0", "b.Gap05", "b.Gap1", "b.Gap15", "b.Gap2", "b.Gap25", "b.Gap3", "b.Gap35", "b.Gap4", "b.Gap45", "b.Gap5", "b.Gap55", "b.Gap6", "b.Gap65", "b.Gap7", "b.Gap75", "b.Gap8"},
				"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowMultiple(
				[]string{"b.ColGap0", "b.ColGap05", "b.ColGap1", "b.ColGap15", "b.ColGap2", "b.ColGap25", "b.ColGap3", "b.ColGap35", "b.ColGap4", "b.ColGap45", "b.ColGap5", "b.ColGap55", "b.ColGap6", "b.ColGap65", "b.ColGap7", "b.ColGap75", "b.ColGap8"},
				"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowMultiple(
				[]string{"b.RowGap0", "b.RowGap05", "b.RowGap1", "b.RowGap15", "b.RowGap2", "b.RowGap25", "b.RowGap3", "b.RowGap35", "b.RowGap4", "b.RowGap45", "b.RowGap5", "b.RowGap55", "b.RowGap6", "b.RowGap65", "b.RowGap7", "b.RowGap75", "b.RowGap8"},
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
			c.RowTo("b.ColMin1", "b.ColMin12", "Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem"),
			c.RowMultiple(
				[]string{"b.Gap0", "b.Gap05", "b.Gap1", "b.Gap15", "b.Gap2", "b.Gap25", "b.Gap3", "b.Gap35", "b.Gap4", "b.Gap45", "b.Gap5", "b.Gap55", "b.Gap6", "b.Gap65", "b.Gap7", "b.Gap75", "b.Gap8"},
				"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowMultiple(
				[]string{"b.ColGap0", "b.ColGap05", "b.ColGap1", "b.ColGap15", "b.ColGap2", "b.ColGap25", "b.ColGap3", "b.ColGap35", "b.ColGap4", "b.ColGap45", "b.ColGap5", "b.ColGap55", "b.ColGap6", "b.ColGap65", "b.ColGap7", "b.ColGap75", "b.ColGap8"},
				"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowMultiple(
				[]string{"b.RowGap0", "b.RowGap05", "b.RowGap1", "b.RowGap15", "b.RowGap2", "b.RowGap25", "b.RowGap3", "b.RowGap35", "b.RowGap4", "b.RowGap45", "b.RowGap5", "b.RowGap55", "b.RowGap6", "b.RowGap65", "b.RowGap7", "b.RowGap75", "b.RowGap8"},
				"Set row gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
			),
			c.RowTo("b.Cols1", "b.Cols12", "Set columns count to 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 or 12"),
			c.Row(
				"b.AutoCount",
				"Set columns count to:",
				b.UList(
					b.MarginTop(b.Spacing0),
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
		e.P("The columns count may be breakpoint-based, with ", "b.Cols*.Mobile()"), " to ", e.Code("b.Cols*.FullHD()"), ".",

		e.P("The ", e.Code("b.Cell"), " constructor creates a single grid cell."),
		c.Modifiers(
			c.RowTo("b.ColStart1", "b.ColStart12", "Start the cell at the given column"),
			c.Row("b.ColStartEnd", "Start the cell at the last column"),
			c.RowTo("b.ColFromEnd1", "b.ColFromEnd12", "End the cell at the given column"),
			c.RowTo("b.ColSpan1", "b.ColSpan12", "Span the cell over this number of columns"),
			c.RowTo("b.RowStart1", "b.RowStart12", "Start the cell at the given row"),
			c.Row("b.RowStartEnd", "Start the cell at the last row"),
			c.RowTo("b.RowFromEnd1", "b.RowFromEnd12", "End the cell at the given row"),
			c.RowTo("b.RowSpan1", "b.RowSpan12", "Span the cell over this number of rows"),
		),
		e.P("The column and row start/end/span may be breakpoint-based, with ", ".Mobile()"), " to ", e.Code(".FullHD()"), ".",
	),
)
