package bulma

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

// Spacing is used for spacing sizes, used in Margin* and Padding*. Values may be 0 to 6. Any other value will be treated as "auto".
type Spacing int

func (s Spacing) String() string {
	if s < 0 || s > 6 {
		return "auto"
	}
	return strconv.FormatInt(int64(s), 10)
}

// Margin sets all margins on the element.
func Margin(s Spacing) e.Class {
	return e.Class("m-" + s.String())
}

// MarginTop sets the top margin on the element.
func MarginTop(s Spacing) e.Class {
	return e.Class("mt-" + s.String())
}

// MarginRight sets the right margin on the element.
func MarginRight(s Spacing) e.Class {
	return e.Class("mr-" + s.String())
}

// MarginBottom sets the bottom margin on the element.
func MarginBottom(s Spacing) e.Class {
	return e.Class("mb-" + s.String())
}

// MarginLeft sets the left margin on the element.
func MarginLeft(s Spacing) e.Class {
	return e.Class("ml-" + s.String())
}

// MarginHorizontal sets the left and right margins on the element.
func MarginHorizontal(s Spacing) e.Class {
	return e.Class("mx-" + s.String())
}

// MarginVertical sets the top and bottom margins on the element.
func MarginVertical(s Spacing) e.Class {
	return e.Class("my-" + s.String())
}

// Padding sets all paddings on the element.
func Padding(s Spacing) e.Class {
	return e.Class("p-" + s.String())
}

// PaddingTop sets the top padding on the element.
func PaddingTop(s Spacing) e.Class {
	return e.Class("pt-" + s.String())
}

// PaddingRight sets the right padding on the element.
func PaddingRight(s Spacing) e.Class {
	return e.Class("pr-" + s.String())
}

// PaddingBottom sets the bottom padding on the element.
func PaddingBottom(s Spacing) e.Class {
	return e.Class("pb-" + s.String())
}

// PaddingLeft sets the left padding on the element.
func PaddingLeft(s Spacing) e.Class {
	return e.Class("pl-" + s.String())
}

// PaddingHorizontal sets the left and right paddings on the element.
func PaddingHorizontal(s Spacing) e.Class {
	return e.Class("px-" + s.String())
}

// PaddingVertical sets the top and bottom paddings on the element.
func PaddingVertical(s Spacing) e.Class {
	return e.Class("py-" + s.String())
}
