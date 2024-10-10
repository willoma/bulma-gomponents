package bulma

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

// ColNum is used for classes based on column numbers (ColSpan, Offset, etc -
// things related to Grid or Columns). Values may be 1 to 12. Any other value
// will be treated as "end", which may be used for ColStart and RowStart.
type ColNum int

func (c ColNum) String() string {
	if c < 1 || c > 12 {
		return "end"
	}

	return strconv.FormatInt(int64(c), 10)
}

// ColFromEnd defines the column number from the end of the grid for a Cell.
func ColFromEnd(end ColNum) ResponsiveClass {
	return ResponsiveClass("is-col-from-end-" + end.String())
}

// ColMin defines the minimum column width for a Grid (valid values are 1 to 32).
func ColMin(min int) e.Class {
	if min < 1 {
		min = 1
	} else if min > 32 {
		min = 32
	}

	return e.Class("is-col-min-" + strconv.Itoa(min))
}

// ColSpan defines the number of columns to span over for a Cell.
func ColSpan(num ColNum) e.Class {
	return e.Class("is-col-span-" + num.String())
}

// ColStart defines the column number for a Cell.
func ColStart(start ColNum) ResponsiveClass {
	return ResponsiveClass("is-col-start-" + start.String())
}

// RowFromEnd defines the row number from the end of the grid for a Cell.
func RowFromEnd(end ColNum) ResponsiveClass {
	return ResponsiveClass("is-row-from-end-" + end.String())
}

// RowSpan defines the number of rows to span over for a Cell.
func RowSpan(num ColNum) e.Class {
	return e.Class("is-row-span-" + num.String())
}

// RowStart defines the row number for a Cell.
func RowStart(start ColNum) ResponsiveClass {
	return ResponsiveClass("is-row-start-" + start.String())
}

// GapSize is a gap size. Values may be 0 to 8, in 0.5 steps. Other values are
// ignored
type GapSize float64

func (g GapSize) String() string {
	if g == 0 || g == 1 || g == 2 || g == 3 || g == 4 || g == 5 || g == 6 || g == 7 || g == 8 {
		return strconv.FormatInt(int64(g), 10)
	}

	if g == 0.5 || g == 1.5 || g == 2.5 || g == 3.5 || g == 4.5 || g == 5.5 || g == 6.5 || g == 7.5 {
		return strconv.FormatFloat(float64(g), 'f', 1, 64)
	}

	return ""
}

// ColGap defines the column gap for a Grid.
func ColGap(gap GapSize) e.Class {
	g := gap.String()
	if g == "" {
		return ""
	}
	return e.Class("is-column-gap-" + g)
}

// Gap defines the gap for a Grid.
func Gap(gap GapSize) e.Class {
	g := gap.String()
	if g == "" {
		return ""
	}
	return e.Class("is-gap-" + g)
}

// RowGap defines the row gap for a Grid.
func RowGap(gap GapSize) e.Class {
	g := gap.String()
	if g == "" {
		return ""
	}
	return e.Class("is-row-gap-" + g)
}

type ColsDef struct {
	ResponsiveClass
}

// Cols defines the columns count for a fixed grid
func Cols(num ColNum) ColsDef {
	return ColsDef{ResponsiveClass("has-" + num.String() + "-cols")}
}

// Offset defines the numeric offset for a Column. Accepted values are 1 to 11.
func Offset(offset ColNum) ResponsiveClass {
	return ResponsiveClass("is-offset-" + offset.String())
}

// Size defines the numeric offset for a Column. Accepted values are 1 to 12.
func Size(size ColNum) ResponsiveClass {
	return ResponsiveClass("is-" + size.String())
}

// ColumnGap defines the gap size for Columns. Accepted values are 0 to 8.
func ColumnGap(gap int) *ResponsiveClasses {
	if gap < 0 || gap > 8 {
		return &ResponsiveClasses{}
	}

	return &ResponsiveClasses{[]string{"is-" + strconv.FormatInt(int64(gap), 10)}, []string{"is-variable"}}
}

// FontSize defines the font size. Accepted values are 1 to 7.
func FontSize(size int) ResponsiveClass {
	if size < 1 || size > 7 {
		return ""
	}

	return ResponsiveClass("is-size-" + strconv.FormatInt(int64(size), 10))
}

// FlexGrow defines the flex-grow value of any child of an element with Flex.
// Accepted values are 0 to 5.
func FlexGrow(size int) e.Class {
	if size < 0 || size > 5 {
		return ""
	}

	return e.Class("is-flex-grow-" + strconv.FormatInt(int64(size), 10))
}

// FlexShrink defines the flex-shrink value of any child of an element with Flex.
// Accepted values are 0 to 5.
func FlexShrink(size int) e.Class {
	if size < 0 || size > 5 {
		return ""
	}

	return e.Class("is-flex-shrink-" + strconv.FormatInt(int64(size), 10))
}
