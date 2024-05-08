package fa

import "strconv"

// Font Awesome icons sizes
var (
	Size2Xs = Class("fa-2xs")
	SizeXs  = Class("fa-xs")
	SizeSm  = Class("fa-sm")
	SizeLg  = Class("fa-lg")
	SizeXl  = Class("fa-xl")
	Size2Xl = Class("fa-2xl")
)

// Font Awesome icons sizes as multipliers (works with 1 to 10).
func Size(size int) Class {
	return Class("fa-" + strconv.Itoa(size) + "x")
}
