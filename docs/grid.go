package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var grid = c.NewPage(
	"Grid", "Grid", "/grid",
	"",

	b.Content(
		el.P("The ", el.Code("b.Grid"), " constructor creates a smart grid. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnFixedGrid(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="fixed-grid">`), " element"},

			el.Code("b.OnGrid(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="grid">`), " element"},

			[]any{el.Code("b.Element"), " not generated with ", el.Code("b.Cell(...)")},
			[]any{"Wrap element in ", el.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{el.Code("string")},
			[]any{"Wrap string in ", el.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the grid",

			[]any{"Other ", el.Code("gomponents.Node")},
			[]any{"Wrap element in ", el.Code("b.Cell(...)"), " and add it to the grid"},

			[]any{el.Code("b.ColMin1"), ", ", el.Code("b.ColMin2"), ", ", el.Code("b.ColMin3"), ", ", el.Code("b.ColMin4"), ", ", el.Code("b.ColMin5"), ", ", el.Code("b.ColMin6"), ", ", el.Code("b.ColMin7"), ", ", el.Code("b.ColMin8"), ", ", el.Code("b.ColMin9"), ", ", el.Code("b.ColMin10"), ", ", el.Code("b.ColMin11"), " or ", el.Code("b.ColMin12")},
			"Set minimum column width to 1.5rem, 3rem, 4.5rem, 6rem, 7.5rem, 9rem, 10.5rem, 12rem, 13.5rem, 15rem, 16.5rem or 18rem",

			[]any{el.Code("b.Gap0"), ", ", el.Code("b.Gap05"), ", ", el.Code("b.Gap1"), ", ", el.Code("b.Gap15"), ", ", el.Code("b.Gap2"), ", ", el.Code("b.Gap25"), ", ", el.Code("b.Gap3"), ", ", el.Code("b.Gap35"), ", ", el.Code("b.Gap4"), ", ", el.Code("b.Gap45"), ", ", el.Code("b.Gap5"), ", ", el.Code("b.Gap55"), ", ", el.Code("b.Gap6"), ", ", el.Code("b.Gap65"), ", ", el.Code("b.Gap7"), ", ", el.Code("b.Gap75"), " or ", el.Code("b.Gap8")},
			"Set gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",

			[]any{el.Code("b.ColGap0"), ", ", el.Code("b.ColGap05"), ", ", el.Code("b.ColGap1"), ", ", el.Code("b.ColGap15"), ", ", el.Code("b.ColGap2"), ", ", el.Code("b.ColGap25"), ", ", el.Code("b.ColGap3"), ", ", el.Code("b.ColGap35"), ", ", el.Code("b.ColGap4"), ", ", el.Code("b.ColGap45"), ", ", el.Code("b.ColGap5"), ", ", el.Code("b.ColGap55"), ", ", el.Code("b.ColGap6"), ", ", el.Code("b.ColGap65"), ", ", el.Code("b.ColGap7"), ", ", el.Code("b.ColGap75"), " or ", el.Code("b.ColGap8")},
			"Set column gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",

			[]any{el.Code("b.RowGap0"), ", ", el.Code("b.RowGap05"), ", ", el.Code("b.RowGap1"), ", ", el.Code("b.RowGap15"), ", ", el.Code("b.RowGap2"), ", ", el.Code("b.RowGap25"), ", ", el.Code("b.RowGap3"), ", ", el.Code("b.RowGap35"), ", ", el.Code("b.RowGap4"), ", ", el.Code("b.RowGap45"), ", ", el.Code("b.RowGap5"), ", ", el.Code("b.RowGap55"), ", ", el.Code("b.RowGap6"), ", ", el.Code("b.RowGap65"), ", ", el.Code("b.RowGap7"), ", ", el.Code("b.RowGap75"), " or ", el.Code("b.RowGap8")},
			"Set row gap to 0, 0.25rem, 0.5rem, 0.75rem, 1rem, 1.25rem, 1.5rem, 1.75rem, 2rem, 2.25rem, 2.5rem, 2.75rem, 3rem, 3.25rem, 3.5rem, 3.75rem or 4rem",
		),

		el.P("The ", el.Code("b.FixedGrid"), " constructor creates a fixed grid. It accepts the same children types as ", el.Code("b.Grid"), ", as well as the following children:"),
		b.DList(
			[]any{el.Code("b.Cols1"), ", ", el.Code("b.Cols2"), ", ", el.Code("b.Cols3"), ", ", el.Code("b.Cols4"), ", ", el.Code("b.Cols5"), ", ", el.Code("b.Cols6"), ", ", el.Code("b.Cols7"), ", ", el.Code("b.Cols8"), ", ", el.Code("b.Cols9"), ", ", el.Code("b.Cols10"), ", ", el.Code("b.Cols11"), " or ", el.Code("b.Cols12")},
			"Set columns count to 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 or 12",

			el.Code("b.AutoCount"),
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
		el.P("The columns count may be breakpoint-based, with ", el.Code("b.Cols*.Mobile()"), " to ", el.Code("b.Cols*.FullHD()"), "."),

		el.P("The ", el.Code("b.Cell"), " constructor creates a single grid cell. The following children have a special meaning:"),
		b.DList(
			[]any{el.Code("b.ColStart1"), " to ", el.Code("b.ColStart12")},
			"Start the cell at the given column",

			el.Code("b.ColStartEnd"),
			"Start the cell at the last column",

			[]any{el.Code("b.ColFromEnd1"), " to ", el.Code("b.ColFromEnd12")},
			"End the cell at the given column",

			[]any{el.Code("b.ColSpan1"), " to ", el.Code("b.ColSpan12")},
			"Span the cell over this number of columns",

			[]any{el.Code("b.RowStart1"), " to ", el.Code("b.RowStart12")},
			"Start the cell at the given row",

			el.Code("b.RowStartEnd"),
			"Start the cell at the last row",

			[]any{el.Code("b.RowFromEnd1"), " to ", el.Code("b.RowFromEnd12")},
			"End the cell at the given row",

			[]any{el.Code("b.RowSpan1"), " to ", el.Code("b.RowSpan12")},
			"Span the cell over this number of rows",
		),
		el.P("The column and row start/end/span may be breakpoint-based, with ", el.Code(".Mobile()"), " to ", el.Code(".FullHD()"), "."),
	),
)
