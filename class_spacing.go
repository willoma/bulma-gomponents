package bulma

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
func Margin(s Spacing) Class {
	return Class("m-" + string(s))
}

// MarginTop sets the top margin on the element.
func MarginTop(s Spacing) Class {
	return Class("mt-" + string(s))
}

// MarginRight sets the right margin on the element.
func MarginRight(s Spacing) Class {
	return Class("mr-" + string(s))
}

// MarginBottom sets the bottom margin on the element.
func MarginBottom(s Spacing) Class {
	return Class("mb-" + string(s))
}

// MarginLeft sets the left margin on the element.
func MarginLeft(s Spacing) Class {
	return Class("ml-" + string(s))
}

// MarginHorizontal sets the left and right margins on the element.
func MarginHorizontal(s Spacing) Class {
	return Class("mx-" + string(s))
}

// MarginVertical sets the top and bottom margins on the element.
func MarginVertical(s Spacing) Class {
	return Class("my-" + string(s))
}

// Padding sets all paddings on the element.
func Padding(s Spacing) Class {
	return Class("p-" + string(s))
}

// PaddingTop sets the top padding on the element.
func PaddingTop(s Spacing) Class {
	return Class("pt-" + string(s))
}

// PaddingRight sets the right padding on the element.
func PaddingRight(s Spacing) Class {
	return Class("pr-" + string(s))
}

// PaddingBottom sets the bottom padding on the element.
func PaddingBottom(s Spacing) Class {
	return Class("pb-" + string(s))
}

// PaddingLeft sets the left padding on the element.
func PaddingLeft(s Spacing) Class {
	return Class("pl-" + string(s))
}

// PaddingHorizontal sets the left and right paddings on the element.
func PaddingHorizontal(s Spacing) Class {
	return Class("px-" + string(s))
}

// PaddingVertical sets the top and bottom paddings on the element.
func PaddingVertical(s Spacing) Class {
	return Class("py-" + string(s))
}
