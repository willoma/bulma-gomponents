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
		e.P("The ", e.Code("b.Grid"), " constructor creates a smart grid. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnFixedGrid(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="fixed-grid">`), " e.Element"},

			e.Code("b.OnGrid(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="grid">`), " e.Element"},

			[]any{e.Code("e.Element"), " not generated with ", e.Code("b.Cell(...)")},
			[]any{"Wrap e.Element in ", e.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{e.Code("string")},
			[]any{"Wrap string in ", e.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the grid",

			[]any{"Other ", e.Code("gomponents.Node")},
			[]any{"Wrap e.Element in ", e.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{e.Code("b.ColMin1"), ", ", e.Code("b.ColMin2"), ", ", e.Code("b.ColMin3"), ", ", e.Code("b.ColMin4"), ", ", e.Code("b.ColMin5"), ", ", e.Code("b.ColMin6"), ", ", e.Code("b.ColMin7"), ", ", e.Code("b.ColMin8"), ", ", e.Code("b.ColMin9"), ", ", e.Code("b.ColMin10"), ", ", e.Code("b.ColMin11"), " or ", e.Code("b.ColMin12")},
			"Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem",

			[]any{e.Code("b.Gap0"), ", ", e.Code("b.Gap05"), ", ", e.Code("b.Gap1"), ", ", e.Code("b.Gap15"), ", ", e.Code("b.Gap2"), ", ", e.Code("b.Gap25"), ", ", e.Code("b.Gap3"), ", ", e.Code("b.Gap35"), ", ", e.Code("b.Gap4"), ", ", e.Code("b.Gap45"), ", ", e.Code("b.Gap5"), ", ", e.Code("b.Gap55"), ", ", e.Code("b.Gap6"), ", ", e.Code("b.Gap65"), ", ", e.Code("b.Gap7"), ", ", e.Code("b.Gap75"), " or ", e.Code("b.Gap8")},
			"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",

			[]any{e.Code("b.ColGap0"), ", ", e.Code("b.ColGap05"), ", ", e.Code("b.ColGap1"), ", ", e.Code("b.ColGap15"), ", ", e.Code("b.ColGap2"), ", ", e.Code("b.ColGap25"), ", ", e.Code("b.ColGap3"), ", ", e.Code("b.ColGap35"), ", ", e.Code("b.ColGap4"), ", ", e.Code("b.ColGap45"), ", ", e.Code("b.ColGap5"), ", ", e.Code("b.ColGap55"), ", ", e.Code("b.ColGap6"), ", ", e.Code("b.ColGap65"), ", ", e.Code("b.ColGap7"), ", ", e.Code("b.ColGap75"), " or ", e.Code("b.ColGap8")},
			"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",

			[]any{e.Code("b.RowGap0"), ", ", e.Code("b.RowGap05"), ", ", e.Code("b.RowGap1"), ", ", e.Code("b.RowGap15"), ", ", e.Code("b.RowGap2"), ", ", e.Code("b.RowGap25"), ", ", e.Code("b.RowGap3"), ", ", e.Code("b.RowGap35"), ", ", e.Code("b.RowGap4"), ", ", e.Code("b.RowGap45"), ", ", e.Code("b.RowGap5"), ", ", e.Code("b.RowGap55"), ", ", e.Code("b.RowGap6"), ", ", e.Code("b.RowGap65"), ", ", e.Code("b.RowGap7"), ", ", e.Code("b.RowGap75"), " or ", e.Code("b.RowGap8")},
			"Set row gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
		),

		e.P("The ", e.Code("b.FixedGrid"), " constructor creates a fixed grid. It accepts the same children types as ", e.Code("b.Grid"), ", as well as the following children:"),
		b.DList(
			[]any{e.Code("b.Cols1"), ", ", e.Code("b.Cols2"), ", ", e.Code("b.Cols3"), ", ", e.Code("b.Cols4"), ", ", e.Code("b.Cols5"), ", ", e.Code("b.Cols6"), ", ", e.Code("b.Cols7"), ", ", e.Code("b.Cols8"), ", ", e.Code("b.Cols9"), ", ", e.Code("b.Cols10"), ", ", e.Code("b.Cols11"), " or ", e.Code("b.Cols12")},
			"Set columns count to 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 or 12",

			e.Code("b.AutoCount"),
			[]any{
				"Set the columns count to:",
				b.UList(
					b.MarginTop(b.Spacing0),
					"2 on mobile",
					"4 on tablet",
					"8 on desktop",
					"12 on widescreen",
					"16 on full HD",
				),
			},
		),
		e.P("The columns count may be breakpoint-based, with ", e.Code("b.Cols*.Mobile()"), " to ", e.Code("b.Cols*.FullHD()"), "."),

		e.P("The ", e.Code("b.Cell"), " constructor creates a single grid cell. The following children have a special meaning:"),
		b.DList(
			[]any{e.Code("b.ColStart1"), " to ", e.Code("b.ColStart12")},
			"Start the cell at the given column",

			e.Code("b.ColStartEnd"),
			"Start the cell at the last column",

			[]any{e.Code("b.ColFromEnd1"), " to ", e.Code("b.ColFromEnd12")},
			"End the cell at the given column",

			[]any{e.Code("b.ColSpan1"), " to ", e.Code("b.ColSpan12")},
			"Span the cell over this number of columns",

			[]any{e.Code("b.RowStart1"), " to ", e.Code("b.RowStart12")},
			"Start the cell at the given row",

			e.Code("b.RowStartEnd"),
			"Start the cell at the last row",

			[]any{e.Code("b.RowFromEnd1"), " to ", e.Code("b.RowFromEnd12")},
			"End the cell at the given row",

			[]any{e.Code("b.RowSpan1"), " to ", e.Code("b.RowSpan12")},
			"Span the cell over this number of rows",
		),
		e.P("The column and row start/end/span may be breakpoint-based, with ", e.Code(".Mobile()"), " to ", e.Code(".FullHD()"), "."),
	),
)
