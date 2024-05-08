package components

import (
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
)

func Table(children ...any) e.Element {
	return b.Table(
		b.Bordered,
		b.Narrow,
	).With(children...)
}

func Modifiers(children ...any) e.Element {
	return Table(
		b.FontSize7,
		b.HeadRow("Modifier", "Action"),
	).With(children...)
}

func Children(children ...any) e.Element {
	return Table(
		b.FontSize7,
		b.HeadRow("Child", "Action"),
	).With(children...)
}

func Row(el string, description ...any) e.Element {
	return b.Row(
		e.Code(el),
		b.TCell(description...),
	)
}

func RowNoCode(el any, description ...any) e.Element {
	return b.Row(
		el,
		b.TCell(description...),
	)
}

func Row2(el1, el2 string, description ...any) e.Element {
	return b.Row(
		b.TCell(e.Code(el1), ", ", e.Code(el2)),
		b.TCell(description...),
	)
}

func RowTo(el1, el2 string, description ...any) e.Element {
	return b.Row(
		b.TCell(e.Code(el1), " to ", e.Code(el2)),
		b.TCell(description...),
	)
}

func RowMultiple(els []string, description ...any) e.Element {
	cell := b.TCell()
	for i, el := range els {
		cell.With(e.Code(el))
		switch {
		case i < len(els)-2:
			cell.With(", ")
		case i == len(els)-2:
			cell.With(" or ")
		}
	}
	return b.Row(
		cell,
		b.TCell(description...),
	)
}

func RowClassesStyles(description ...any) e.Element {
	return b.Row(
		b.TCell(e.Code("e.Class"), ", ", e.Code("e.Classer"), ", ", e.Code("e.Classeser"), " or ", e.Code("e.Styles")),
		b.TCell(description...),
	)
}

func RowNodeAttribute(description ...any) e.Element {
	return b.Row(
		b.TCell(e.Code("gomponents.Node"), "of type ", e.Code("gomponents.AttributeType")),
		b.TCell(description...),
	)
}

func RowNodeElement(description ...any) e.Element {
	return b.Row(
		b.TCell(e.Code("gomponents.Node"), "of type ", e.Code("gomponents.ElementType")),
		b.TCell(description...),
	)
}

func RowDefault(description ...any) e.Element {
	return b.Row(
		"Anything else",
		b.TCell(description...),
	)
}
