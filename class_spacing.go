package bulma

import e "github.com/willoma/gomplements"

// Spacing is used for spacing sizes, used in Margin* and Padding*.
type Spacing string

const (
	Spacing0    Spacing = "0"
	Spacing1    Spacing = "1"
	Spacing2    Spacing = "2"
	Spacing3    Spacing = "3"
	Spacing4    Spacing = "4"
	Spacing5    Spacing = "5"
	Spacing6    Spacing = "6"
	SpacingAuto Spacing = "auto"
)

// Margin sets all margins on the element.
func Margin(s Spacing) e.Class {
	return e.Class("m-" + string(s))
}

// MarginTop sets the top margin on the element.
func MarginTop(s Spacing) e.Class {
	return e.Class("mt-" + string(s))
}

// MarginRight sets the right margin on the element.
func MarginRight(s Spacing) e.Class {
	return e.Class("mr-" + string(s))
}

// MarginBottom sets the bottom margin on the element.
func MarginBottom(s Spacing) e.Class {
	return e.Class("mb-" + string(s))
}

// MarginLeft sets the left margin on the element.
func MarginLeft(s Spacing) e.Class {
	return e.Class("ml-" + string(s))
}

// MarginHorizontal sets the left and right margins on the element.
func MarginHorizontal(s Spacing) e.Class {
	return e.Class("mx-" + string(s))
}

// MarginVertical sets the top and bottom margins on the element.
func MarginVertical(s Spacing) e.Class {
	return e.Class("my-" + string(s))
}

// Padding sets all paddings on the element.
func Padding(s Spacing) e.Class {
	return e.Class("p-" + string(s))
}

// PaddingTop sets the top padding on the element.
func PaddingTop(s Spacing) e.Class {
	return e.Class("pt-" + string(s))
}

// PaddingRight sets the right padding on the element.
func PaddingRight(s Spacing) e.Class {
	return e.Class("pr-" + string(s))
}

// PaddingBottom sets the bottom padding on the element.
func PaddingBottom(s Spacing) e.Class {
	return e.Class("pb-" + string(s))
}

// PaddingLeft sets the left padding on the element.
func PaddingLeft(s Spacing) e.Class {
	return e.Class("pl-" + string(s))
}

// PaddingHorizontal sets the left and right paddings on the element.
func PaddingHorizontal(s Spacing) e.Class {
	return e.Class("px-" + string(s))
}

// PaddingVertical sets the top and bottom paddings on the element.
func PaddingVertical(s Spacing) e.Class {
	return e.Class("py-" + string(s))
}
