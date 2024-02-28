package fa

// Rotation (see https://fontawesome.com/docs/web/style/rotate)
const (
	Rotate90       = FaClass("fa-rotate-90")
	Rotate180      = FaClass("fa-rotate-180")
	Rotate270      = FaClass("fa-rotate-270")
	FlipHorizontal = FaClass("fa-flip-horizontal")
	FlipVertical   = FaClass("fa-flip-vertical")
	FlipBoth       = FaClass("fa-flip-both")
)

type Rotate float64
