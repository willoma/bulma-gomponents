package bulma

// Class is an Element modifier to apply a HTML class.
type Class string

// MultiClass is an Element modifier to apply multiple HTML classes.
type MultiClass struct {
	Responsive []string // Classes that may be modified with one of the responsive functions
	Static     []string // Classes that will not be modified with responsive functions
}

// "is-" classes
const (
	Active                     = Class("is-active")                        // Button*, Dropdown, Dropup, NavbarItem, PanelTabs/a
	AlignContentBaseline       = Class("is-align-content-baseline")        // any child of an element with Flex
	AlignContentCenter         = Class("is-align-content-center")          // any child of an element with Flex
	AlignContentEnd            = Class("is-align-content-end")             // any child of an element with Flex
	AlignContentFlexEnd        = Class("is-align-content-flex-end")        // any child of an element with Flex
	AlignContentFlexStart      = Class("is-align-content-flex-start")      // any child of an element with Flex
	AlignContentSpaceAround    = Class("is-align-content-space-around")    // any child of an element with Flex
	AlignContentSpaceBetween   = Class("is-align-content-space-between")   // any child of an element with Flex
	AlignContentSpaceEvenly    = Class("is-align-content-space-evenly")    // any child of an element with Flex
	AlignContentStart          = Class("is-align-content-start")           // any child of an element with Flex
	AlignContentStretch        = Class("is-align-content-stretch")         // any child of an element with Flex
	AlignItemsBaseline         = Class("is-align-items-baseline")          // any child of an element with Flex
	AlignItemsCenter           = Class("is-align-items-center")            // any child of an element with Flex
	AlignItemsEnd              = Class("is-align-items-end")               // any child of an element with Flex
	AlignItemsFlexEnd          = Class("is-align-items-flex-end")          // any child of an element with Flex
	AlignItemsFlexStart        = Class("is-align-items-flex-start")        // any child of an element with Flex
	AlignItemsSelfEnd          = Class("is-align-items-self-end")          // any child of an element with Flex
	AlignItemsSelfStart        = Class("is-align-items-self-start")        // any child of an element with Flex
	AlignItemsStart            = Class("is-align-items-start")             // any child of an element with Flex
	AlignItemsStretch          = Class("is-align-items-stretch")           // any child of an element with Flex
	AlignSelfAuto              = Class("is-align-self-auto")               // any child of an element with Flex
	AlignSelfBaseline          = Class("is-align-self-baseline")           // any child of an element with Flex
	AlignSelfCenter            = Class("is-align-self-center")             // any child of an element with Flex
	AlignSelfFlexEnd           = Class("is-align-self-flex-end")           // any child of an element with Flex
	AlignSelfFlexStart         = Class("is-align-self-flex-start")         // any child of an element with Flex
	AlignSelfStretch           = Class("is-align-self-stretch")            // any child of an element with Flex
	Arrowless                  = Class("is-arrowless")                     // NavbarLink
	Bordered                   = Class("is-bordered")                      // Table
	Boxed                      = Class("is-boxed")                         // NavbarDropdown, Tabs, File
	Capitalized                = Class("is-capitalized")                   // any text element
	Centered                   = Class("is-centered")                      // Columns, Buttons, Breadcrumb, Pagination, Tabs, File
	Clearfix                   = Class("is-clearfix")                      // any element
	Clickable                  = Class("is-clickable")                     // any element
	Clipped                    = Class("is-clipped")                       // any element
	Code                       = Class("is-family-code")                   // any text element
	Current                    = Class("is-current")                       // PaginationLink
	Desktop                    = Class("is-desktop")                       // Columns
	Disabled                   = Class("is-disabled")                      // PaginationLink, PaginationNext, PaginationPrevious
	Expanded                   = Class("is-expanded")                      // NavbarItem, NavbarAHref, Control
	FixedBottom                = Class("is-fixed-bottom")                  // Navbar
	FixedTop                   = Class("is-fixed-top")                     // Navbar
	Flex                       = Class("is-flex")                          // any element (responsive variants are supported)
	FlexColumn                 = Class("is-flex-direction-column")         // any child of an element with Flex
	FlexColumnReverse          = Class("is-flex-direction-column-reverse") // any child of an element with Flex
	FlexGrow0                  = Class("is-flex-grow-0")                   // any child of an element with Flex
	FlexGrow1                  = Class("is-flex-grow-1")                   // any child of an element with Flex
	FlexGrow2                  = Class("is-flex-grow-2")                   // any child of an element with Flex
	FlexGrow3                  = Class("is-flex-grow-3")                   // any child of an element with Flex
	FlexGrow4                  = Class("is-flex-grow-4")                   // any child of an element with Flex
	FlexGrow5                  = Class("is-flex-grow-5")                   // any child of an element with Flex
	FlexNowrap                 = Class("is-flex-wrap-nowrap")              // any child of an element with Flex
	FlexRow                    = Class("is-flex-direction-row")            // any child of an element with Flex
	FlexRowReverse             = Class("is-flex-direction-row-reverse")    // any child of an element with Flex
	FlexShrink0                = Class("is-flex-shrink-0")                 // any child of an element with Flex
	FlexShrink1                = Class("is-flex-shrink-1")                 // any child of an element with Flex
	FlexShrink2                = Class("is-flex-shrink-2")                 // any child of an element with Flex
	FlexShrink3                = Class("is-flex-shrink-3")                 // any child of an element with Flex
	FlexShrink4                = Class("is-flex-shrink-4")                 // any child of an element with Flex
	FlexShrink5                = Class("is-flex-shrink-5")                 // any child of an element with Flex
	FlexWrap                   = Class("is-flex-wrap-wrap")                // any child of an element with Flex
	FlexWrapReverse            = Class("is-flex-wrap-wrap-reverse")        // any child of an element with Flex
	Focused                    = Class("is-focused")                       // Button*, Input*, Select, SelectMultiple
	FontSize1                  = Class("is-size-1")                        // any text element (responsive variants are supported)
	FontSize2                  = Class("is-size-2")                        // any text element (responsive variants are supported)
	FontSize3                  = Class("is-size-3")                        // any text element (responsive variants are supported)
	FontSize4                  = Class("is-size-4")                        // any text element (responsive variants are supported)
	FontSize5                  = Class("is-size-5")                        // any text element (responsive variants are supported)
	FontSize6                  = Class("is-size-6")                        // any text element (responsive variants are supported)
	FontSize7                  = Class("is-size-7")                        // any text element (responsive variants are supported)
	FourFifths                 = Class("is-four-fifths")                   // Column
	Full                       = Class("is-full")                          // Column
	FullHeight                 = Class("fullheight")                       // Hero
	FullHeightWithNavbar       = Class("fullheight-with-navbar")           // Hero
	FullWidth                  = Class("is-fullwidth")                     // Button*, Image, ImageImg, Table, Tabs, File, Select
	Gapless                    = Class("is-gapless")                       // Columns
	Ghost                      = Class("is-ghost")                         // Button*
	Grouped                    = Class("is-grouped")                       // Field
	Half                       = Class("is-half")                          // Column
	HalfHeight                 = Class("halfheight")                       // Hero
	Hidden                     = Class("is-hidden")                        // any element (responsive variants are supported)
	Horizontal                 = Class("is-horizontal")                    // Field
	Hoverable                  = Class("is-hoverable")                     // Table, Dropdown, Dropup, NavbarItem
	Hovered                    = Class("is-hovered")                       // Button*, Input*, Select, SelectMultiple
	Img1By1                    = Class("is-1by1")                          // Image, ImageImg
	Img5By4                    = Class("is-5by4")                          // Image, ImageImg
	Img4By3                    = Class("is-4by3")                          // Image, ImageImg
	Img3By2                    = Class("is-3by2")                          // Image, ImageImg
	Img5By3                    = Class("is-5by3")                          // Image, ImageImg
	Img16By9                   = Class("is-16by9")                         // Image, ImageImg
	Img2By1                    = Class("is-2by1")                          // Image, ImageImg
	Img3By1                    = Class("is-3by1")                          // Image, ImageImg
	Img4By5                    = Class("is-4by5")                          // Image, ImageImg
	Img3By4                    = Class("is-3by4")                          // Image, ImageImg
	Img2By3                    = Class("is-2by3")                          // Image, ImageImg
	Img3By5                    = Class("is-3by5")                          // Image, ImageImg
	Img9By16                   = Class("is-9by16")                         // Image, ImageImg
	Img1By2                    = Class("is-1by2")                          // Image, ImageImg
	Img1By3                    = Class("is-1by3")                          // Image, ImageImg
	ImgSq16                    = Class("is-16x16")                         // Image, ImageImg
	ImgSq24                    = Class("is-24x24")                         // Image, ImageImg
	ImgSq32                    = Class("is-24x24")                         // Image, ImageImg
	ImgSq48                    = Class("is-48x48")                         // Image, ImageImg
	ImgSq64                    = Class("is-64x64")                         // Image, ImageImg
	ImgSq96                    = Class("is-96x96")                         // Image, ImageImg
	ImgSq128                   = Class("is-128x128")                       // Image, ImageImg
	ImgSquare                  = Class("is-square")                        // Image, ImageImg
	Inline                     = Class("is-inline")                        // any element (responsive variants are supported)
	InlineBlock                = Class("is-inline-block")                  // any element (responsive variants are supported)
	InlineFlex                 = Class("is-inline-flex")                   // any element (responsive variants are supported)
	Inverted                   = Class("is-inverted")                      // Button*
	Invisible                  = Class("is-invisible")                     // any element
	Italic                     = Class("is-italic")                        // any text element
	JustifyContentCenter       = Class("is-justify-content-center")        // any child of an element with Flex
	JustifyContentEnd          = Class("is-justify-content-end")           // any child of an element with Flex
	JustifyContentFlexEnd      = Class("is-justify-content-flex-end")      // any child of an element with Flex
	JustifyContentFlexStart    = Class("is-justify-content-flex-start")    // any child of an element with Flex
	JustifyContentLeft         = Class("is-justify-content-left")          // any child of an element with Flex
	JustifyContentRight        = Class("is-justify-content-right")         // any child of an element with Flex
	JustifyContentSpaceAround  = Class("is-justify-content-space-around")  // any child of an element with Flex
	JustifyContentSpaceBetween = Class("is-justify-content-space-between") // any child of an element with Flex
	JustifyContentSpaceEvenly  = Class("is-justify-content-space-evenly")  // any child of an element with Flex
	JustifyContentStart        = Class("is-justify-content-start")         // any child of an element with Flex
	Large                      = Class("is-large")                         // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Left                       = Class("is-left")                          // Label/input, Label/select
	Loading                    = Class("is-loading")                       // Button*, Input*, Select, SelectMultiple
	Lowercase                  = Class("is-lowercase")                     // any text element
	Medium                     = Class("is-medium")                        // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Mobile                     = Class("is-mobile")                        // Columns, Level
	Monospace                  = Class("is-family-monospace")              // any text element
	Multiline                  = Class("is-multiline")                     // Columns
	Narrow                     = Class("is-narrow")                        // Column, Table
	NavbarFixedBottom          = Class("has-navbar-fixed-bottom")          // HTML/body
	NavbarFixedTop             = Class("has-navbar-fixed-top")             // HTML/body
	Normal                     = Class("is-normal")                        // Button*, Content, Tag, File
	Offset1                    = Class("is-offset-1")                      // Column
	Offset2                    = Class("is-offset-2")                      // Column
	Offset3                    = Class("is-offset-3")                      // Column
	Offset4                    = Class("is-offset-4")                      // Column
	Offset5                    = Class("is-offset-5")                      // Column
	Offset6                    = Class("is-offset-6")                      // Column
	Offset7                    = Class("is-offset-7")                      // Column
	Offset8                    = Class("is-offset-8")                      // Column
	Offset9                    = Class("is-offset-9")                      // Column
	Offset10                   = Class("is-offset-10")                     // Column
	Offset11                   = Class("is-offset-11")                     // Column
	OffsetFourFifths           = Class("is-offset-four-fifths")            // Column
	OffsetHalf                 = Class("is-offset-half")                   // Column
	OffsetOneFifth             = Class("is-offset-one-fifth")              // Column
	OffsetOneQuarter           = Class("is-offset-one-quarter")            // Column
	OffsetOneThird             = Class("is-offset-one-third")              // Column
	OffsetThreeFifths          = Class("is-offset-three-fifths")           // Column
	OffsetThreeQuarters        = Class("is-offset-three-quarters")         // Column
	OffsetTwoFifths            = Class("is-offset-two-fifths")             // Column
	OffsetTwoThirds            = Class("is-offset-two-thirds")             // Column
	OlLowerAlpha               = Class("is-lower-alpha")                   // Content/ol
	OlLowerRoman               = Class("is-lower-roman")                   // Content/ol
	OlUpperAlpha               = Class("is-upper-alpha")                   // Content/ol
	OlUpperRoman               = Class("is-upper-roman")                   // Content/ol
	OneFifth                   = Class("is-one-fifth")                     // Column
	OneQuarter                 = Class("is-one-quarter")                   // Column
	OneThird                   = Class("is-one-third")                     // Column
	Outlined                   = Class("is-outlined")                      // Button*
	Overlay                    = Class("is-overlay")                       // any element
	PulledLeft                 = Class("is-pulled-left")                   // any element
	PulledRight                = Class("is-pulled-right")                  // any element
	Radiusless                 = Class("is-radiusless")                    // any element
	Relative                   = Class("is-relative")                      // any element
	Responsive                 = Class("is-responsive")                    // Button
	Right                      = Class("is-right")                         // Buttons, Breadcrumb, NavbarDropdown, Tabs, File, Label/input, Label/select
	Rounded                    = Class("is-rounded")                       // Button*, Image/img, ImageImg, Rounded, Pagination, Input*, Select, SelectMultiple
	SansSerif                  = Class("is-family-sans-serif")             // any text element
	Selected                   = Class("is-selected")                      // Button*, Table
	Shadowless                 = Class("is-shadowless")                    // any element
	Size1                      = Class("is-1")                             // Column, Tile
	Size2                      = Class("is-2")                             // Column, Tile
	Size3                      = Class("is-3")                             // Column, Tile
	Size4                      = Class("is-4")                             // Column, Tile
	Size5                      = Class("is-5")                             // Column, Tile
	Size6                      = Class("is-6")                             // Column, Tile
	Size7                      = Class("is-7")                             // Column, Tile
	Size8                      = Class("is-8")                             // Column, Tile
	Size9                      = Class("is-9")                             // Column, Tile
	Size10                     = Class("is-10")                            // Column, Tile
	Size11                     = Class("is-11")                            // Column, Tile
	Size12                     = Class("is-12")                            // Column, Tile
	Small                      = Class("is-small")                         // Button*, Buttons, Content, Delete, Icon, Breadcrumb, Pagination, File, Hero
	Spaced                     = Class("is-spaced")                        // Title*, Subtitle*, Navbar
	SrOnly                     = Class("is-sr-only")                       // any element
	Static                     = Class("is-static")                        // Button*, Input*
	Striped                    = Class("is-striped")                       // Table
	Tab                        = Class("is-tab")                           // NavbarItem, NavbarAHref
	Text                       = Class("is-text")                          // Button*
	TextPrimary                = Class("is-family-primary")                // any text element
	TextSecondary              = Class("is-family-secondary")              // any text element
	ThreeFifths                = Class("is-three-fifths")                  // Column
	ThreeQuarters              = Class("is-three-quarters")                // Column
	Toggle                     = Class("is-toggle")                        // Tabs
	Transparent                = Class("is-transparent")                   // Navbar
	TwoFifths                  = Class("is-two-fifths")                    // Column
	TwoThirds                  = Class("is-two-thirds")                    // Column
	Underlined                 = Class("is-underlined")                    // any text element
	Unselectable               = Class("is-unselectable")                  // any element
	Uppercase                  = Class("is-uppercase")                     // any text element
	VCentered                  = Class("is-vcentered")                     // Columns
	VisibilityBlock            = Class("is-block")                         // any element (responsive variants are supported)
)

// "has-" classes
const (
	Addons            = Class("has-addons")           // Buttons, Tags, Field
	ArrowSeparator    = Class("has-arrow-separator")  // Breadcrumb
	Bold              = Class("has-text-weight-bold") // any text element
	BulletSeparator   = Class("has-bullet-separator") // Breadcrumb
	DotSeparator      = Class("has-dot-separator")    // Breadcrumb
	FixedSize         = Class("has-fixed-size")       // Textarea
	HasDropdown       = Class("has-dropdown")         // NavbarItem
	Heading           = Class("heading")
	IconsLeft         = Class("has-icons-left")           // Control
	IconsRight        = Class("has-icons-right")          // Control
	SemiBold          = Class("has-text-weight-semibold") // any text element
	Ratio             = Class("has-ratio")                // Image/*
	Shadow            = Class("has-shadow")               // Navbar
	SucceedsSeparator = Class("has-succeeds-separator")   // Breadcrumb
	TextCentered      = Class("has-text-centered")        // any text element (responsive variants are supported)
	TextJustified     = Class("has-text-justified")       // any text element (responsive variants are supported)
	TextLeft          = Class("has-text-left")            // any text element (responsive variants are supported)
	TextLight         = Class("has-text-weight-light")    // any text element
	TextMedium        = Class("has-text-weight-medium")   // any text element
	TextNormal        = Class("has-text-weight-normal")   // any text element
	TextRight         = Class("has-text-right")           // any text element (responsive variants are supported)
)

// "is-" classes with automatic adding of other class
var (
	Gap0             = MultiClass{[]string{"is-0"}, []string{"is-variable"}}              // Columns
	Gap1             = MultiClass{[]string{"is-1"}, []string{"is-variable"}}              // Columns
	Gap2             = MultiClass{[]string{"is-2"}, []string{"is-variable"}}              // Columns
	Gap3             = MultiClass{[]string{"is-3"}, []string{"is-variable"}}              // Columns
	Gap4             = MultiClass{[]string{"is-4"}, []string{"is-variable"}}              // Columns
	Gap5             = MultiClass{[]string{"is-5"}, []string{"is-variable"}}              // Columns
	Gap6             = MultiClass{[]string{"is-6"}, []string{"is-variable"}}              // Columns
	Gap7             = MultiClass{[]string{"is-7"}, []string{"is-variable"}}              // Columns
	Gap8             = MultiClass{[]string{"is-8"}, []string{"is-variable"}}              // Columns
	GroupedCentered  = MultiClass{Static: []string{"is-grouped", "is-grouped-centered"}}  // Field
	GroupedMultiline = MultiClass{Static: []string{"is-grouped", "is-grouped-multiline"}} // Field
	GroupedRight     = MultiClass{Static: []string{"is-grouped", "is-grouped-right"}}     // Field
	ToggleRounded    = MultiClass{nil, []string{"is-toggle", "is-toggle-rounded"}}        // Tabs
)

// "has-" classes with automatic adding of other class
var (
	AddonsCentered = MultiClass{Static: []string{"has-addons", "has-addons-centered"}} // Field
	AddonsRight    = MultiClass{Static: []string{"has-addons", "has-addons-right"}}    // Field
	HasDropup      = MultiClass{Static: []string{"has-dropdown", "has-dropdown-up"}}   // NavbarItem
)

func changeSizePrefix(prefix string, class Class) Class {
	if class == Small || class == Normal || class == Medium || class == Large {
		return Class(prefix) + class[3:]
	}
	return class
}
