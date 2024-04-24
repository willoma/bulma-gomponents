package bulma

type Class string

type Classes []Class

func (c Classes) Classes() []Class {
	return c
}

type Classer interface {
	Class() Class
}

type Classeser interface {
	Classes() []Class
}

type ExternalClassesAndStyles interface {
	ClassesAndStyles() ([]Class, Styles)
}

type Styles map[string]string

// Style applies styles to the element.
func Style(propertiesValues ...string) Styles {
	stylesObj := Styles{}
	var limit int
	if len(propertiesValues)%2 == 0 {
		limit = len(propertiesValues)
	} else {
		limit = len(propertiesValues) - 1
	}
	for i := 0; i < limit; i += 2 {
		stylesObj[propertiesValues[i]] = propertiesValues[i+1]
	}
	return stylesObj
}

// "is-" classes
const (
	Active                     = Class("is-active")                          // Button*, Dropdown, Dropup, NavbarItem, PanelTabs/a
	AlignContentBaseline       = Class("is-align-content-baseline")          // any child of an element with Flex
	AlignContentCenter         = Class("is-align-content-center")            // any child of an element with Flex
	AlignContentEnd            = Class("is-align-content-end")               // any child of an element with Flex
	AlignContentFlexEnd        = Class("is-align-content-flex-end")          // any child of an element with Flex
	AlignContentFlexStart      = Class("is-align-content-flex-start")        // any child of an element with Flex
	AlignContentSpaceAround    = Class("is-align-content-space-around")      // any child of an element with Flex
	AlignContentSpaceBetween   = Class("is-align-content-space-between")     // any child of an element with Flex
	AlignContentSpaceEvenly    = Class("is-align-content-space-evenly")      // any child of an element with Flex
	AlignContentStart          = Class("is-align-content-start")             // any child of an element with Flex
	AlignContentStretch        = Class("is-align-content-stretch")           // any child of an element with Flex
	AlignItemsBaseline         = Class("is-align-items-baseline")            // any child of an element with Flex
	AlignItemsCenter           = Class("is-align-items-center")              // any child of an element with Flex
	AlignItemsEnd              = Class("is-align-items-end")                 // any child of an element with Flex
	AlignItemsFlexEnd          = Class("is-align-items-flex-end")            // any child of an element with Flex
	AlignItemsFlexStart        = Class("is-align-items-flex-start")          // any child of an element with Flex
	AlignItemsSelfEnd          = Class("is-align-items-self-end")            // any child of an element with Flex
	AlignItemsSelfStart        = Class("is-align-items-self-start")          // any child of an element with Flex
	AlignItemsStart            = Class("is-align-items-start")               // any child of an element with Flex
	AlignItemsStretch          = Class("is-align-items-stretch")             // any child of an element with Flex
	AlignSelfAuto              = Class("is-align-self-auto")                 // any child of an element with Flex
	AlignSelfBaseline          = Class("is-align-self-baseline")             // any child of an element with Flex
	AlignSelfCenter            = Class("is-align-self-center")               // any child of an element with Flex
	AlignSelfFlexEnd           = Class("is-align-self-flex-end")             // any child of an element with Flex
	AlignSelfFlexStart         = Class("is-align-self-flex-start")           // any child of an element with Flex
	AlignSelfStretch           = Class("is-align-self-stretch")              // any child of an element with Flex
	Arrowless                  = Class("is-arrowless")                       // NavbarLink
	Bordered                   = Class("is-bordered")                        // Table
	Boxed                      = Class("is-boxed")                           // NavbarDropdown, Tabs, File
	Capitalized                = Class("is-capitalized")                     // any text element
	Centered                   = Class("is-centered")                        // Columns, Buttons, Breadcrumb, Pagination, Tabs, File
	Clearfix                   = Class("is-clearfix")                        // any element
	Clickable                  = Class("is-clickable")                       // any element
	Clipped                    = Class("is-clipped")                         // any element
	Code                       = Class("is-family-code")                     // any text element
	ColFromEnd1                = ResponsiveClass("is-col-from-end-1")        // Cell
	ColFromEnd2                = ResponsiveClass("is-col-from-end-2")        // Cell
	ColFromEnd3                = ResponsiveClass("is-col-from-end-3")        // Cell
	ColFromEnd4                = ResponsiveClass("is-col-from-end-4")        // Cell
	ColFromEnd5                = ResponsiveClass("is-col-from-end-5")        // Cell
	ColFromEnd6                = ResponsiveClass("is-col-from-end-6")        // Cell
	ColFromEnd7                = ResponsiveClass("is-col-from-end-7")        // Cell
	ColFromEnd8                = ResponsiveClass("is-col-from-end-8")        // Cell
	ColFromEnd9                = ResponsiveClass("is-col-from-end-9")        // Cell
	ColFromEnd10               = ResponsiveClass("is-col-from-end-10")       // Cell
	ColFromEnd11               = ResponsiveClass("is-col-from-end-11")       // Cell
	ColFromEnd12               = ResponsiveClass("is-col-from-end-12")       // Cell
	ColGap0                    = Class("is-column-gap-0")                    // Grid
	ColGap05                   = Class("is-column-gap-0.5")                  // Grid
	ColGap1                    = Class("is-column-gap-1")                    // Grid
	ColGap15                   = Class("is-column-gap-1.5")                  // Grid
	ColGap2                    = Class("is-column-gap-2")                    // Grid
	ColGap25                   = Class("is-column-gap-2.5")                  // Grid
	ColGap3                    = Class("is-column-gap-3")                    // Grid
	ColGap35                   = Class("is-column-gap-3.5")                  // Grid
	ColGap4                    = Class("is-column-gap-4")                    // Grid
	ColGap45                   = Class("is-column-gap-4.5")                  // Grid
	ColGap5                    = Class("is-column-gap-5")                    // Grid
	ColGap55                   = Class("is-column-gap-5.5")                  // Grid
	ColGap6                    = Class("is-column-gap-6")                    // Grid
	ColGap65                   = Class("is-column-gap-6.5")                  // Grid
	ColGap7                    = Class("is-column-gap-7")                    // Grid
	ColGap75                   = Class("is-column-gap-7.5")                  // Grid
	ColGap8                    = Class("is-column-gap-8")                    // Grid
	ColMin1                    = Class("is-col-min-1")                       // Grid
	ColMin2                    = Class("is-col-min-2")                       // Grid
	ColMin3                    = Class("is-col-min-3")                       // Grid
	ColMin4                    = Class("is-col-min-4")                       // Grid
	ColMin5                    = Class("is-col-min-5")                       // Grid
	ColMin6                    = Class("is-col-min-6")                       // Grid
	ColMin7                    = Class("is-col-min-7")                       // Grid
	ColMin8                    = Class("is-col-min-8")                       // Grid
	ColMin9                    = Class("is-col-min-9")                       // Grid
	ColMin10                   = Class("is-col-min-10")                      // Grid
	ColMin11                   = Class("is-col-min-11")                      // Grid
	ColMin12                   = Class("is-col-min-12")                      // Grid
	ColSpan1                   = ResponsiveClass("is-col-span-1")            // Cell
	ColSpan2                   = ResponsiveClass("is-col-span-2")            // Cell
	ColSpan3                   = ResponsiveClass("is-col-span-3")            // Cell
	ColSpan4                   = ResponsiveClass("is-col-span-4")            // Cell
	ColSpan5                   = ResponsiveClass("is-col-span-5")            // Cell
	ColSpan6                   = ResponsiveClass("is-col-span-6")            // Cell
	ColSpan7                   = ResponsiveClass("is-col-span-7")            // Cell
	ColSpan8                   = ResponsiveClass("is-col-span-8")            // Cell
	ColSpan9                   = ResponsiveClass("is-col-span-9")            // Cell
	ColSpan10                  = ResponsiveClass("is-col-span-10")           // Cell
	ColSpan11                  = ResponsiveClass("is-col-span-11")           // Cell
	ColSpan12                  = ResponsiveClass("is-col-span-12")           // Cell
	ColStart1                  = ResponsiveClass("is-col-start-1")           // Cell
	ColStart2                  = ResponsiveClass("is-col-start-2")           // Cell
	ColStart3                  = ResponsiveClass("is-col-start-3")           // Cell
	ColStart4                  = ResponsiveClass("is-col-start-4")           // Cell
	ColStart5                  = ResponsiveClass("is-col-start-5")           // Cell
	ColStart6                  = ResponsiveClass("is-col-start-6")           // Cell
	ColStart7                  = ResponsiveClass("is-col-start-7")           // Cell
	ColStart8                  = ResponsiveClass("is-col-start-8")           // Cell
	ColStart9                  = ResponsiveClass("is-col-start-9")           // Cell
	ColStart10                 = ResponsiveClass("is-col-start-10")          // Cell
	ColStart11                 = ResponsiveClass("is-col-start-11")          // Cell
	ColStart12                 = ResponsiveClass("is-col-start-12")          // Cell
	ColStartEnd                = Class("is-col-start-end")                   // Cell
	Current                    = Class("is-current")                         // PaginationLink
	Desktop                    = Class("is-desktop")                         // Columns
	Disabled                   = Class("is-disabled")                        // PaginationLink, PaginationNext, PaginationPrevious
	Expanded                   = Class("is-expanded")                        // NavbarItem, NavbarAHref, Control
	FixedBottom                = Class("is-fixed-bottom")                    // Navbar
	FixedTop                   = Class("is-fixed-top")                       // Navbar
	Flex                       = ResponsiveClass("is-flex")                  // any element
	FlexColumn                 = Class("is-flex-direction-column")           // any child of an element with Flex
	FlexColumnReverse          = Class("is-flex-direction-column-reverse")   // any child of an element with Flex
	FlexGrow0                  = Class("is-flex-grow-0")                     // any child of an element with Flex
	FlexGrow1                  = Class("is-flex-grow-1")                     // any child of an element with Flex
	FlexGrow2                  = Class("is-flex-grow-2")                     // any child of an element with Flex
	FlexGrow3                  = Class("is-flex-grow-3")                     // any child of an element with Flex
	FlexGrow4                  = Class("is-flex-grow-4")                     // any child of an element with Flex
	FlexGrow5                  = Class("is-flex-grow-5")                     // any child of an element with Flex
	FlexNowrap                 = Class("is-flex-wrap-nowrap")                // any child of an element with Flex
	FlexRow                    = Class("is-flex-direction-row")              // any child of an element with Flex
	FlexRowReverse             = Class("is-flex-direction-row-reverse")      // any child of an element with Flex
	FlexShrink0                = Class("is-flex-shrink-0")                   // any child of an element with Flex
	FlexShrink1                = Class("is-flex-shrink-1")                   // any child of an element with Flex
	FlexShrink2                = Class("is-flex-shrink-2")                   // any child of an element with Flex
	FlexShrink3                = Class("is-flex-shrink-3")                   // any child of an element with Flex
	FlexShrink4                = Class("is-flex-shrink-4")                   // any child of an element with Flex
	FlexShrink5                = Class("is-flex-shrink-5")                   // any child of an element with Flex
	FlexWrap                   = Class("is-flex-wrap-wrap")                  // any child of an element with Flex
	FlexWrapReverse            = Class("is-flex-wrap-wrap-reverse")          // any child of an element with Flex
	Fluid                      = Class("is-fluid")                           // Container
	Focused                    = Class("is-focused")                         // Button*, Input*, Select, SelectMultiple
	FontSize1                  = ResponsiveClass("is-size-1")                // any text element
	FontSize2                  = ResponsiveClass("is-size-2")                // any text element
	FontSize3                  = ResponsiveClass("is-size-3")                // any text element
	FontSize4                  = ResponsiveClass("is-size-4")                // any text element
	FontSize5                  = ResponsiveClass("is-size-5")                // any text element
	FontSize6                  = ResponsiveClass("is-size-6")                // any text element
	FontSize7                  = ResponsiveClass("is-size-7")                // any text element
	FourFifths                 = ResponsiveClass("is-four-fifths")           // Column
	Full                       = ResponsiveClass("is-full")                  // Column
	FullHD                     = Class("is-fullhd")                          // Container
	FullHeight                 = Class("is-fullheight")                      // Hero
	FullHeightWithNavbar       = Class("is-fullheight-with-navbar")          // Hero
	FullWidth                  = Class("is-fullwidth")                       // Button*, Image, ImageImg, Table, Tabs, File, Select
	Gapless                    = Class("is-gapless")                         // Columns
	Gap0                       = Class("is-gap-0")                           // Grid
	Gap05                      = Class("is-gap-0.5")                         // Grid
	Gap1                       = Class("is-gap-1")                           // Grid
	Gap15                      = Class("is-gap-1.5")                         // Grid
	Gap2                       = Class("is-gap-2")                           // Grid
	Gap25                      = Class("is-gap-2.5")                         // Grid
	Gap3                       = Class("is-gap-3")                           // Grid
	Gap35                      = Class("is-gap-3.5")                         // Grid
	Gap4                       = Class("is-gap-4")                           // Grid
	Gap45                      = Class("is-gap-4.5")                         // Grid
	Gap5                       = Class("is-gap-5")                           // Grid
	Gap55                      = Class("is-gap-5.5")                         // Grid
	Gap6                       = Class("is-gap-6")                           // Grid
	Gap65                      = Class("is-gap-6.5")                         // Grid
	Gap7                       = Class("is-gap-7")                           // Grid
	Gap75                      = Class("is-gap-7.5")                         // Grid
	Gap8                       = Class("is-gap-8")                           // Grid
	Ghost                      = Class("is-ghost")                           // Button*
	Grouped                    = Class("is-grouped")                         // Field
	Half                       = ResponsiveClass("is-half")                  // Column
	HalfHeight                 = Class("is-halfheight")                      // Hero
	Hidden                     = ResponsiveClass("is-hidden")                // any element
	Horizontal                 = Class("is-horizontal")                      // Field
	Hoverable                  = Class("is-hoverable")                       // Table, Dropdown, Dropup, NavbarItem
	Hovered                    = Class("is-hovered")                         // Button*, Input*, Select, SelectMultiple
	Img1By1                    = Class("is-1by1")                            // Image, ImageImg
	Img5By4                    = Class("is-5by4")                            // Image, ImageImg
	Img4By3                    = Class("is-4by3")                            // Image, ImageImg
	Img3By2                    = Class("is-3by2")                            // Image, ImageImg
	Img5By3                    = Class("is-5by3")                            // Image, ImageImg
	Img16By9                   = Class("is-16by9")                           // Image, ImageImg
	Img2By1                    = Class("is-2by1")                            // Image, ImageImg
	Img3By1                    = Class("is-3by1")                            // Image, ImageImg
	Img4By5                    = Class("is-4by5")                            // Image, ImageImg
	Img3By4                    = Class("is-3by4")                            // Image, ImageImg
	Img2By3                    = Class("is-2by3")                            // Image, ImageImg
	Img3By5                    = Class("is-3by5")                            // Image, ImageImg
	Img9By16                   = Class("is-9by16")                           // Image, ImageImg
	Img1By2                    = Class("is-1by2")                            // Image, ImageImg
	Img1By3                    = Class("is-1by3")                            // Image, ImageImg
	ImgSq16                    = Class("is-16x16")                           // Image, ImageImg
	ImgSq24                    = Class("is-24x24")                           // Image, ImageImg
	ImgSq32                    = Class("is-24x24")                           // Image, ImageImg
	ImgSq48                    = Class("is-48x48")                           // Image, ImageImg
	ImgSq64                    = Class("is-64x64")                           // Image, ImageImg
	ImgSq96                    = Class("is-96x96")                           // Image, ImageImg
	ImgSq128                   = Class("is-128x128")                         // Image, ImageImg
	ImgSquare                  = Class("is-square")                          // Image, ImageImg
	Inline                     = ResponsiveClass("is-inline")                // any element
	InlineBlock                = ResponsiveClass("is-inline-block")          // any element
	InlineFlex                 = ResponsiveClass("is-inline-flex")           // any element
	Inverted                   = Class("is-inverted")                        // Button*
	Invisible                  = ResponsiveClass("is-invisible")             // any element
	Italic                     = Class("is-italic")                          // any text element
	JustifyContentCenter       = Class("is-justify-content-center")          // any child of an element with Flex
	JustifyContentEnd          = Class("is-justify-content-end")             // any child of an element with Flex
	JustifyContentFlexEnd      = Class("is-justify-content-flex-end")        // any child of an element with Flex
	JustifyContentFlexStart    = Class("is-justify-content-flex-start")      // any child of an element with Flex
	JustifyContentLeft         = Class("is-justify-content-left")            // any child of an element with Flex
	JustifyContentRight        = Class("is-justify-content-right")           // any child of an element with Flex
	JustifyContentSpaceAround  = Class("is-justify-content-space-around")    // any child of an element with Flex
	JustifyContentSpaceBetween = Class("is-justify-content-space-between")   // any child of an element with Flex
	JustifyContentSpaceEvenly  = Class("is-justify-content-space-evenly")    // any child of an element with Flex
	JustifyContentStart        = Class("is-justify-content-start")           // any child of an element with Flex
	Large                      = Class("is-large")                           // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Left                       = Class("is-left")                            // Label/input, Label/select
	Loading                    = Class("is-loading")                         // Button*, Input*, Select, SelectMultiple
	Lowercase                  = Class("is-lowercase")                       // any text element
	MaxDesktop                 = Class("is-max-desktop")                     // Container
	MaxWidescreen              = Class("is-max-widescreen")                  // Container
	Medium                     = Class("is-medium")                          // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Mobile                     = Class("is-mobile")                          // Columns, Level
	Monospace                  = Class("is-family-monospace")                // any text element
	Multiline                  = Class("is-multiline")                       // Columns
	Multiple                   = Class("is-multiple")                        // Select
	Narrow                     = ResponsiveClass("is-narrow")                // Column, Table
	NavbarFixedBottom          = Class("has-navbar-fixed-bottom")            // HTML/body
	NavbarFixedTop             = Class("has-navbar-fixed-top")               // HTML/body
	Normal                     = Class("is-normal")                          // Button*, Content, Tag, File
	Offset1                    = ResponsiveClass("is-offset-1")              // Column
	Offset2                    = ResponsiveClass("is-offset-2")              // Column
	Offset3                    = ResponsiveClass("is-offset-3")              // Column
	Offset4                    = ResponsiveClass("is-offset-4")              // Column
	Offset5                    = ResponsiveClass("is-offset-5")              // Column
	Offset6                    = ResponsiveClass("is-offset-6")              // Column
	Offset7                    = ResponsiveClass("is-offset-7")              // Column
	Offset8                    = ResponsiveClass("is-offset-8")              // Column
	Offset9                    = ResponsiveClass("is-offset-9")              // Column
	Offset10                   = ResponsiveClass("is-offset-10")             // Column
	Offset11                   = ResponsiveClass("is-offset-11")             // Column
	OffsetFourFifths           = ResponsiveClass("is-offset-four-fifths")    // Column
	OffsetHalf                 = ResponsiveClass("is-offset-half")           // Column
	OffsetOneFifth             = ResponsiveClass("is-offset-one-fifth")      // Column
	OffsetOneQuarter           = ResponsiveClass("is-offset-one-quarter")    // Column
	OffsetOneThird             = ResponsiveClass("is-offset-one-third")      // Column
	OffsetThreeFifths          = ResponsiveClass("is-offset-three-fifths")   // Column
	OffsetThreeQuarters        = ResponsiveClass("is-offset-three-quarters") // Column
	OffsetTwoFifths            = ResponsiveClass("is-offset-two-fifths")     // Column
	OffsetTwoThirds            = ResponsiveClass("is-offset-two-thirds")     // Column
	OlLowerAlpha               = Class("is-lower-alpha")                     // Content/ol
	OlLowerRoman               = Class("is-lower-roman")                     // Content/ol
	OlUpperAlpha               = Class("is-upper-alpha")                     // Content/ol
	OlUpperRoman               = Class("is-upper-roman")                     // Content/ol
	OneFifth                   = ResponsiveClass("is-one-fifth")             // Column
	OneQuarter                 = ResponsiveClass("is-one-quarter")           // Column
	OneThird                   = ResponsiveClass("is-one-third")             // Column
	Outlined                   = Class("is-outlined")                        // Button*
	Overlay                    = Class("is-overlay")                         // any element
	PositionAbsolute           = Class("is-position-absolute")               // any element
	PositionFixed              = Class("is-position-fixed")                  // any element
	PositionRelative           = Class("is-position-relative")               // any element
	PositionStatic             = Class("is-position-static")                 // any element
	PositionSticky             = Class("is-position-sticky")                 // any element
	PulledLeft                 = Class("is-pulled-left")                     // any element
	PulledRight                = Class("is-pulled-right")                    // any element
	Radiusless                 = Class("is-radiusless")                      // any element
	Relative                   = Class("is-relative")                        // any element
	Responsive                 = Class("is-responsive")                      // Button
	Right                      = Class("is-right")                           // Buttons, Breadcrumb, NavbarDropdown, Tabs, File, Label/input, Label/select
	Rounded                    = Class("is-rounded")                         // Button*, Image/img, ImageImg, Rounded, Pagination, Input*, Select, SelectMultiple
	RowFromEnd1                = ResponsiveClass("is-row-from-end-1")        // Cell
	RowFromEnd2                = ResponsiveClass("is-row-from-end-2")        // Cell
	RowFromEnd3                = ResponsiveClass("is-row-from-end-3")        // Cell
	RowFromEnd4                = ResponsiveClass("is-row-from-end-4")        // Cell
	RowFromEnd5                = ResponsiveClass("is-row-from-end-5")        // Cell
	RowFromEnd6                = ResponsiveClass("is-row-from-end-6")        // Cell
	RowFromEnd7                = ResponsiveClass("is-row-from-end-7")        // Cell
	RowFromEnd8                = ResponsiveClass("is-row-from-end-8")        // Cell
	RowFromEnd9                = ResponsiveClass("is-row-from-end-9")        // Cell
	RowFromEnd10               = ResponsiveClass("is-row-from-end-10")       // Cell
	RowFromEnd11               = ResponsiveClass("is-row-from-end-11")       // Cell
	RowFromEnd12               = ResponsiveClass("is-row-from-end-12")       // Cell
	RowGap0                    = Class("is-row-gap-0")                       // Grid
	RowGap05                   = Class("is-row-gap-0.5")                     // Grid
	RowGap1                    = Class("is-row-gap-1")                       // Grid
	RowGap15                   = Class("is-row-gap-1.5")                     // Grid
	RowGap2                    = Class("is-row-gap-2")                       // Grid
	RowGap25                   = Class("is-row-gap-2.5")                     // Grid
	RowGap3                    = Class("is-row-gap-3")                       // Grid
	RowGap35                   = Class("is-row-gap-3.5")                     // Grid
	RowGap4                    = Class("is-row-gap-4")                       // Grid
	RowGap45                   = Class("is-row-gap-4.5")                     // Grid
	RowGap5                    = Class("is-row-gap-5")                       // Grid
	RowGap55                   = Class("is-row-gap-5.5")                     // Grid
	RowGap6                    = Class("is-row-gap-6")                       // Grid
	RowGap65                   = Class("is-row-gap-6.5")                     // Grid
	RowGap7                    = Class("is-row-gap-7")                       // Grid
	RowGap75                   = Class("is-row-gap-7.5")                     // Grid
	RowGap8                    = Class("is-row-gap-8")                       // Grid
	RowSpan1                   = ResponsiveClass("is-row-span-1")            // Cell
	RowSpan2                   = ResponsiveClass("is-row-span-2")            // Cell
	RowSpan3                   = ResponsiveClass("is-row-span-3")            // Cell
	RowSpan4                   = ResponsiveClass("is-row-span-4")            // Cell
	RowSpan5                   = ResponsiveClass("is-row-span-5")            // Cell
	RowSpan6                   = ResponsiveClass("is-row-span-6")            // Cell
	RowSpan7                   = ResponsiveClass("is-row-span-7")            // Cell
	RowSpan8                   = ResponsiveClass("is-row-span-8")            // Cell
	RowSpan9                   = ResponsiveClass("is-row-span-9")            // Cell
	RowSpan10                  = ResponsiveClass("is-row-span-10")           // Cell
	RowSpan11                  = ResponsiveClass("is-row-span-11")           // Cell
	RowSpan12                  = ResponsiveClass("is-row-span-12")           // Cell
	RowStart1                  = ResponsiveClass("is-row-start-1")           // Cell
	RowStart2                  = ResponsiveClass("is-row-start-2")           // Cell
	RowStart3                  = ResponsiveClass("is-row-start-3")           // Cell
	RowStart4                  = ResponsiveClass("is-row-start-4")           // Cell
	RowStart5                  = ResponsiveClass("is-row-start-5")           // Cell
	RowStart6                  = ResponsiveClass("is-row-start-6")           // Cell
	RowStart7                  = ResponsiveClass("is-row-start-7")           // Cell
	RowStart8                  = ResponsiveClass("is-row-start-8")           // Cell
	RowStart9                  = ResponsiveClass("is-row-start-9")           // Cell
	RowStart10                 = ResponsiveClass("is-row-start-10")          // Cell
	RowStart11                 = ResponsiveClass("is-row-start-11")          // Cell
	RowStart12                 = ResponsiveClass("is-row-start-12")          // Cell
	RowStartEnd                = Class("is-row-start-end")                   // Cell
	SansSerif                  = Class("is-family-sans-serif")               // any text element
	Selected                   = Class("is-selected")                        // Button*, Table
	Shadowless                 = Class("is-shadowless")                      // any element
	Size1                      = ResponsiveClass("is-1")                     // Column
	Size2                      = ResponsiveClass("is-2")                     // Column
	Size3                      = ResponsiveClass("is-3")                     // Column
	Size4                      = ResponsiveClass("is-4")                     // Column
	Size5                      = ResponsiveClass("is-5")                     // Column
	Size6                      = ResponsiveClass("is-6")                     // Column
	Size7                      = ResponsiveClass("is-7")                     // Column
	Size8                      = ResponsiveClass("is-8")                     // Column
	Size9                      = ResponsiveClass("is-9")                     // Column
	Size10                     = ResponsiveClass("is-10")                    // Column
	Size11                     = ResponsiveClass("is-11")                    // Column
	Size12                     = ResponsiveClass("is-12")                    // Column
	Skeleton                   = Class("is-skeleton")                        // any components
	Small                      = Class("is-small")                           // Button*, Buttons, Content, Delete, Icon, Breadcrumb, Pagination, File, Hero
	Spaced                     = Class("is-spaced")                          // Title*, Subtitle*, Navbar
	SrOnly                     = Class("is-sr-only")                         // any element
	Static                     = Class("is-static")                          // Button*, Input*
	Striped                    = Class("is-striped")                         // Table
	Tab                        = Class("is-tab")                             // NavbarItem, NavbarAHref
	FamilyPrimary              = Class("is-family-primary")                  // any text element
	FamilySecondary            = Class("is-family-secondary")                // any text element
	ThreeFifths                = ResponsiveClass("is-three-fifths")          // Column
	ThreeQuarters              = ResponsiveClass("is-three-quarters")        // Column
	Toggle                     = Class("is-toggle")                          // Tabs
	Transparent                = Class("is-transparent")                     // Navbar
	TwoFifths                  = ResponsiveClass("is-two-fifths")            // Column
	TwoThirds                  = ResponsiveClass("is-two-thirds")            // Column
	Underlined                 = Class("is-underlined")                      // any text element
	Unselectable               = Class("is-unselectable")                    // any element
	Up                         = Class("is-up")                              // Dropdown
	Uppercase                  = Class("is-uppercase")                       // any text element
	VCentered                  = Class("is-vcentered")                       // Columns
	VisibilityBlock            = ResponsiveClass("is-block")                 // any element
	Widescreen                 = Class("is-widescreen")                      // Container
)

// "has-" classes
const (
	Addons            = Class("has-addons")                   // Buttons, Tags, Field
	ArrowSeparator    = Class("has-arrow-separator")          // Breadcrumb
	AutoCount         = Class("has-auto-count")               // FixedGrid
	BulletSeparator   = Class("has-bullet-separator")         // Breadcrumb
	Cols1             = ResponsiveClass("has-1-cols")         // FixedGrid
	Cols2             = ResponsiveClass("has-2-cols")         // FixedGrid
	Cols3             = ResponsiveClass("has-3-cols")         // FixedGrid
	Cols4             = ResponsiveClass("has-4-cols")         // FixedGrid
	Cols5             = ResponsiveClass("has-5-cols")         // FixedGrid
	Cols6             = ResponsiveClass("has-6-cols")         // FixedGrid
	Cols7             = ResponsiveClass("has-7-cols")         // FixedGrid
	Cols8             = ResponsiveClass("has-8-cols")         // FixedGrid
	Cols9             = ResponsiveClass("has-9-cols")         // FixedGrid
	Cols10            = ResponsiveClass("has-10-cols")        // FixedGrid
	Cols11            = ResponsiveClass("has-11-cols")        // FixedGrid
	Cols12            = ResponsiveClass("has-12-cols")        // FixedGrid
	DotSeparator      = Class("has-dot-separator")            // Breadcrumb
	FixedSize         = Class("has-fixed-size")               // Textarea
	IconsLeft         = Class("has-icons-left")               // Control
	IconsRight        = Class("has-icons-right")              // Control
	Ratio             = Class("has-ratio")                    // Image/*
	Shadow            = Class("has-shadow")                   // Navbar
	HasSkeleton       = Class("has-skeleton")                 // any component
	SucceedsSeparator = Class("has-succeeds-separator")       // Breadcrumb
	TextCentered      = ResponsiveClass("has-text-centered")  // any text element
	TextJustified     = ResponsiveClass("has-text-justified") // any text element
	TextLeft          = ResponsiveClass("has-text-left")      // any text element
	TextRight         = ResponsiveClass("has-text-right")     // any text element
	WeightBold        = Class("has-text-weight-bold")         // any text element
	WeightLight       = Class("has-text-weight-light")        // any text element
	WeightMedium      = Class("has-text-weight-medium")       // any text element
	WeightNormal      = Class("has-text-weight-normal")       // any text element
	WeightSemiBold    = Class("has-text-weight-semibold")     // any text element
)

// "is-" classes with automatic adding of other class
var (
	ColumnGap0       = ResponsiveClasses{[]string{"is-0"}, []string{"is-variable"}} // Columns
	ColumnGap1       = ResponsiveClasses{[]string{"is-1"}, []string{"is-variable"}} // Columns
	ColumnGap2       = ResponsiveClasses{[]string{"is-2"}, []string{"is-variable"}} // Columns
	ColumnGap3       = ResponsiveClasses{[]string{"is-3"}, []string{"is-variable"}} // Columns
	ColumnGap4       = ResponsiveClasses{[]string{"is-4"}, []string{"is-variable"}} // Columns
	ColumnGap5       = ResponsiveClasses{[]string{"is-5"}, []string{"is-variable"}} // Columns
	ColumnGap6       = ResponsiveClasses{[]string{"is-6"}, []string{"is-variable"}} // Columns
	ColumnGap7       = ResponsiveClasses{[]string{"is-7"}, []string{"is-variable"}} // Columns
	ColumnGap8       = ResponsiveClasses{[]string{"is-8"}, []string{"is-variable"}} // Columns
	GroupedCentered  = Classes{"is-grouped", "is-grouped-centered"}                 // Field
	GroupedMultiline = Classes{"is-grouped", "is-grouped-multiline"}                // Field
	GroupedRight     = Classes{"is-grouped", "is-grouped-right"}                    // Field
	ToggleRounded    = Classes{"is-toggle", "is-toggle-rounded"}                    // Tabs
)

// "has-" classes with automatic adding of other class
var (
	AddonsCentered = Classes{"has-addons", "has-addons-centered"} // Field
	AddonsRight    = Classes{"has-addons", "has-addons-right"}    // Field
)

func changeSizePrefix(prefix string, class Class) Class {
	if class == Small || class == Normal || class == Medium || class == Large {
		return Class(prefix) + class[3:]
	}
	return class
}
