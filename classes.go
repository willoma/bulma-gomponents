package bulma

import e "github.com/willoma/gomplements"

// "is-" classes
const (
	Active                     = e.Class("is-active")                        // Button*, Dropdown, Dropup, NavbarItem, PanelTabs/a
	AlignContentBaseline       = e.Class("is-align-content-baseline")        // any child of an element with Flex
	AlignContentCenter         = e.Class("is-align-content-center")          // any child of an element with Flex
	AlignContentEnd            = e.Class("is-align-content-end")             // any child of an element with Flex
	AlignContentFlexEnd        = e.Class("is-align-content-flex-end")        // any child of an element with Flex
	AlignContentFlexStart      = e.Class("is-align-content-flex-start")      // any child of an element with Flex
	AlignContentSpaceAround    = e.Class("is-align-content-space-around")    // any child of an element with Flex
	AlignContentSpaceBetween   = e.Class("is-align-content-space-between")   // any child of an element with Flex
	AlignContentSpaceEvenly    = e.Class("is-align-content-space-evenly")    // any child of an element with Flex
	AlignContentStart          = e.Class("is-align-content-start")           // any child of an element with Flex
	AlignContentStretch        = e.Class("is-align-content-stretch")         // any child of an element with Flex
	AlignItemsBaseline         = e.Class("is-align-items-baseline")          // any child of an element with Flex
	AlignItemsCenter           = e.Class("is-align-items-center")            // any child of an element with Flex
	AlignItemsEnd              = e.Class("is-align-items-end")               // any child of an element with Flex
	AlignItemsFlexEnd          = e.Class("is-align-items-flex-end")          // any child of an element with Flex
	AlignItemsFlexStart        = e.Class("is-align-items-flex-start")        // any child of an element with Flex
	AlignItemsSelfEnd          = e.Class("is-align-items-self-end")          // any child of an element with Flex
	AlignItemsSelfStart        = e.Class("is-align-items-self-start")        // any child of an element with Flex
	AlignItemsStart            = e.Class("is-align-items-start")             // any child of an element with Flex
	AlignItemsStretch          = e.Class("is-align-items-stretch")           // any child of an element with Flex
	AlignSelfAuto              = e.Class("is-align-self-auto")               // any child of an element with Flex
	AlignSelfBaseline          = e.Class("is-align-self-baseline")           // any child of an element with Flex
	AlignSelfCenter            = e.Class("is-align-self-center")             // any child of an element with Flex
	AlignSelfFlexEnd           = e.Class("is-align-self-flex-end")           // any child of an element with Flex
	AlignSelfFlexStart         = e.Class("is-align-self-flex-start")         // any child of an element with Flex
	AlignSelfStretch           = e.Class("is-align-self-stretch")            // any child of an element with Flex
	Arrowless                  = e.Class("is-arrowless")                     // NavbarLink
	Bordered                   = e.Class("is-bordered")                      // Table
	Boxed                      = e.Class("is-boxed")                         // NavbarDropdown, Tabs, File
	Capitalized                = e.Class("is-capitalized")                   // any text element
	Centered                   = e.Class("is-centered")                      // Columns, Buttons, Breadcrumb, Pagination, Tabs, File
	Clearfix                   = e.Class("is-clearfix")                      // any element
	Clickable                  = e.Class("is-clickable")                     // any element
	Clipped                    = e.Class("is-clipped")                       // any element
	Code                       = e.Class("is-family-code")                   // any text element
	Current                    = e.Class("is-current")                       // PaginationLink
	Desktop                    = e.Class("is-desktop")                       // Columns
	Disabled                   = e.Class("is-disabled")                      // PaginationLink, PaginationNext, PaginationPrevious
	Expanded                   = e.Class("is-expanded")                      // NavbarItem, NavbarAHref, Control
	FixedBottom                = e.Class("is-fixed-bottom")                  // Navbar
	FixedTop                   = e.Class("is-fixed-top")                     // Navbar
	Flex                       = ResponsiveClass("is-flex")                  // any element
	FlexColumn                 = e.Class("is-flex-direction-column")         // any child of an element with Flex
	FlexColumnReverse          = e.Class("is-flex-direction-column-reverse") // any child of an element with Flex
	FlexNowrap                 = e.Class("is-flex-wrap-nowrap")              // any child of an element with Flex
	FlexRow                    = e.Class("is-flex-direction-row")            // any child of an element with Flex
	FlexRowReverse             = e.Class("is-flex-direction-row-reverse")    // any child of an element with Flex
	FlexWrap                   = e.Class("is-flex-wrap-wrap")                // any child of an element with Flex
	FlexWrapReverse            = e.Class("is-flex-wrap-wrap-reverse")        // any child of an element with Flex
	Fluid                      = e.Class("is-fluid")                         // Container
	Focused                    = e.Class("is-focused")                       // Button*, Input*, Select, SelectMultiple
	FourFifths                 = ResponsiveClass("is-four-fifths")           // Column
	Full                       = ResponsiveClass("is-full")                  // Column
	FullHD                     = e.Class("is-fullhd")                        // Container
	FullHeight                 = e.Class("is-fullheight")                    // Hero
	FullHeightWithNavbar       = e.Class("is-fullheight-with-navbar")        // Hero
	FullWidth                  = e.Class("is-fullwidth")                     // Button*, Image, ImageImg, Table, Tabs, File, Select
	Gapless                    = e.Class("is-gapless")                       // Columns
	Ghost                      = e.Class("is-ghost")                         // Button*
	Grouped                    = e.Class("is-grouped")                       // Field
	Half                       = ResponsiveClass("is-half")                  // Column
	HalfHeight                 = e.Class("is-halfheight")                    // Hero
	Hidden                     = ResponsiveClass("is-hidden")                // any element
	Horizontal                 = e.Class("is-horizontal")                    // Field
	Hoverable                  = e.Class("is-hoverable")                     // Table, Dropdown, Dropup, NavbarItem
	Hovered                    = e.Class("is-hovered")                       // Button*, Input*, Select, SelectMultiple
	Img1By1                    = e.Class("is-1by1")                          // Image, ImageImg
	Img5By4                    = e.Class("is-5by4")                          // Image, ImageImg
	Img4By3                    = e.Class("is-4by3")                          // Image, ImageImg
	Img3By2                    = e.Class("is-3by2")                          // Image, ImageImg
	Img5By3                    = e.Class("is-5by3")                          // Image, ImageImg
	Img16By9                   = e.Class("is-16by9")                         // Image, ImageImg
	Img2By1                    = e.Class("is-2by1")                          // Image, ImageImg
	Img3By1                    = e.Class("is-3by1")                          // Image, ImageImg
	Img4By5                    = e.Class("is-4by5")                          // Image, ImageImg
	Img3By4                    = e.Class("is-3by4")                          // Image, ImageImg
	Img2By3                    = e.Class("is-2by3")                          // Image, ImageImg
	Img3By5                    = e.Class("is-3by5")                          // Image, ImageImg
	Img9By16                   = e.Class("is-9by16")                         // Image, ImageImg
	Img1By2                    = e.Class("is-1by2")                          // Image, ImageImg
	Img1By3                    = e.Class("is-1by3")                          // Image, ImageImg
	ImgSq16                    = e.Class("is-16x16")                         // Image, ImageImg
	ImgSq24                    = e.Class("is-24x24")                         // Image, ImageImg
	ImgSq32                    = e.Class("is-24x24")                         // Image, ImageImg
	ImgSq48                    = e.Class("is-48x48")                         // Image, ImageImg
	ImgSq64                    = e.Class("is-64x64")                         // Image, ImageImg
	ImgSq96                    = e.Class("is-96x96")                         // Image, ImageImg
	ImgSq128                   = e.Class("is-128x128")                       // Image, ImageImg
	ImgSquare                  = e.Class("is-square")                        // Image, ImageImg
	Inline                     = ResponsiveClass("is-inline")                // any element
	InlineBlock                = ResponsiveClass("is-inline-block")          // any element
	InlineFlex                 = ResponsiveClass("is-inline-flex")           // any element
	Inverted                   = e.Class("is-inverted")                      // Button*
	Invisible                  = ResponsiveClass("is-invisible")             // any element
	Italic                     = e.Class("is-italic")                        // any text element
	JustifyContentCenter       = e.Class("is-justify-content-center")        // any child of an element with Flex
	JustifyContentEnd          = e.Class("is-justify-content-end")           // any child of an element with Flex
	JustifyContentFlexEnd      = e.Class("is-justify-content-flex-end")      // any child of an element with Flex
	JustifyContentFlexStart    = e.Class("is-justify-content-flex-start")    // any child of an element with Flex
	JustifyContentLeft         = e.Class("is-justify-content-left")          // any child of an element with Flex
	JustifyContentRight        = e.Class("is-justify-content-right")         // any child of an element with Flex
	JustifyContentSpaceAround  = e.Class("is-justify-content-space-around")  // any child of an element with Flex
	JustifyContentSpaceBetween = e.Class("is-justify-content-space-between") // any child of an element with Flex
	JustifyContentSpaceEvenly  = e.Class("is-justify-content-space-evenly")  // any child of an element with Flex
	JustifyContentStart        = e.Class("is-justify-content-start")         // any child of an element with Flex
	Large                      = e.Class("is-large")                         // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Left                       = e.Class("is-left")                          // Label/input, Label/select
	Loading                    = e.Class("is-loading")                       // Button*, Input*, Select, SelectMultiple
	Lowercase                  = e.Class("is-lowercase")                     // any text element
	MaxDesktop                 = e.Class("is-max-desktop")                   // Container
	MaxTablet                  = e.Class("is-max-tablet")                    // Container
	MaxWidescreen              = e.Class("is-max-widescreen")                // Container
	Medium                     = e.Class("is-medium")                        // Button*, Buttons, Content, Delete, Icon, Tag, Breadcrumb, Pagination, File, Hero, Section
	Mobile                     = e.Class("is-mobile")                        // Columns, Level
	Monospace                  = e.Class("is-family-monospace")              // any text element
	Multiline                  = e.Class("is-multiline")                     // Columns
	Multiple                   = e.Class("is-multiple")                      // Select
	Narrow                     = ResponsiveClass("is-narrow")                // Column, Table
	NavbarFixedBottom          = e.Class("has-navbar-fixed-bottom")          // HTML/body
	NavbarFixedTop             = e.Class("has-navbar-fixed-top")             // HTML/body
	Normal                     = e.Class("is-normal")                        // Button*, Content, Tag, File
	OffsetFourFifths           = ResponsiveClass("is-offset-four-fifths")    // Column
	OffsetHalf                 = ResponsiveClass("is-offset-half")           // Column
	OffsetOneFifth             = ResponsiveClass("is-offset-one-fifth")      // Column
	OffsetOneQuarter           = ResponsiveClass("is-offset-one-quarter")    // Column
	OffsetOneThird             = ResponsiveClass("is-offset-one-third")      // Column
	OffsetThreeFifths          = ResponsiveClass("is-offset-three-fifths")   // Column
	OffsetThreeQuarters        = ResponsiveClass("is-offset-three-quarters") // Column
	OffsetTwoFifths            = ResponsiveClass("is-offset-two-fifths")     // Column
	OffsetTwoThirds            = ResponsiveClass("is-offset-two-thirds")     // Column
	OlLowerAlpha               = e.Class("is-lower-alpha")                   // Content/ol
	OlLowerRoman               = e.Class("is-lower-roman")                   // Content/ol
	OlUpperAlpha               = e.Class("is-upper-alpha")                   // Content/ol
	OlUpperRoman               = e.Class("is-upper-roman")                   // Content/ol
	OneFifth                   = ResponsiveClass("is-one-fifth")             // Column
	OneQuarter                 = ResponsiveClass("is-one-quarter")           // Column
	OneThird                   = ResponsiveClass("is-one-third")             // Column
	Outlined                   = e.Class("is-outlined")                      // Button*
	Overlay                    = e.Class("is-overlay")                       // any element
	PositionAbsolute           = e.Class("is-position-absolute")             // any element
	PositionFixed              = e.Class("is-position-fixed")                // any element
	PositionRelative           = e.Class("is-position-relative")             // any element
	PositionStatic             = e.Class("is-position-static")               // any element
	PositionSticky             = e.Class("is-position-sticky")               // any element
	PulledLeft                 = e.Class("is-pulled-left")                   // any element
	PulledRight                = e.Class("is-pulled-right")                  // any element
	Radiusless                 = e.Class("is-radiusless")                    // any element
	Relative                   = e.Class("is-relative")                      // any element
	Responsive                 = e.Class("is-responsive")                    // Button
	Right                      = e.Class("is-right")                         // Buttons, Breadcrumb, NavbarDropdown, Tabs, File, Label/input, Label/select
	Rounded                    = e.Class("is-rounded")                       // Button*, Image/img, ImageImg, Rounded, Pagination, Input*, Select, SelectMultiple
	SansSerif                  = e.Class("is-family-sans-serif")             // any text element
	Selected                   = e.Class("is-selected")                      // Button*, Table
	Shadowless                 = e.Class("is-shadowless")                    // any element
	Skeleton                   = e.Class("is-skeleton")                      // any components
	Small                      = e.Class("is-small")                         // Button*, Buttons, Content, Delete, Icon, Breadcrumb, Pagination, File, Hero
	Spaced                     = e.Class("is-spaced")                        // Title*, Subtitle*, Navbar
	SrOnly                     = e.Class("is-sr-only")                       // any element
	Static                     = e.Class("is-static")                        // Button*, Input*
	Striped                    = e.Class("is-striped")                       // Table
	Tab                        = e.Class("is-tab")                           // NavbarItem, NavbarAHref
	FamilyPrimary              = e.Class("is-family-primary")                // any text element
	FamilySecondary            = e.Class("is-family-secondary")              // any text element
	ThreeFifths                = ResponsiveClass("is-three-fifths")          // Column
	ThreeQuarters              = ResponsiveClass("is-three-quarters")        // Column
	Toggle                     = e.Class("is-toggle")                        // Tabs
	Transparent                = e.Class("is-transparent")                   // Navbar
	TwoFifths                  = ResponsiveClass("is-two-fifths")            // Column
	TwoThirds                  = ResponsiveClass("is-two-thirds")            // Column
	Underlined                 = e.Class("is-underlined")                    // any text element
	Unselectable               = e.Class("is-unselectable")                  // any element
	Up                         = e.Class("is-up")                            // Dropdown
	Uppercase                  = e.Class("is-uppercase")                     // any text element
	VCentered                  = e.Class("is-vcentered")                     // Columns
	VisibilityBlock            = ResponsiveClass("is-block")                 // any element
	Widescreen                 = e.Class("is-widescreen")                    // Container
)

// "has-" classes
const (
	Addons            = e.Class("has-addons")                 // Buttons, Tags, Field
	ArrowSeparator    = e.Class("has-arrow-separator")        // Breadcrumb
	AutoCount         = e.Class("has-auto-count")             // FixedGrid
	BulletSeparator   = e.Class("has-bullet-separator")       // Breadcrumb
	DotSeparator      = e.Class("has-dot-separator")          // Breadcrumb
	FixedSize         = e.Class("has-fixed-size")             // Textarea
	IconsLeft         = e.Class("has-icons-left")             // Control
	IconsRight        = e.Class("has-icons-right")            // Control
	RadiusSmall       = e.Class("has-radius-small")           // any component
	RadiusNormal      = e.Class("has-radius-normal")          // any component
	RadiusLarge       = e.Class("has-radius-large")           // any component
	RadiusRounded     = e.Class("has-radius-rounded")         // any component
	Ratio             = e.Class("has-ratio")                  // Image/*
	Shadow            = e.Class("has-shadow")                 // Navbar
	HasSkeleton       = e.Class("has-skeleton")               // any component
	SucceedsSeparator = e.Class("has-succeeds-separator")     // Breadcrumb
	TextCentered      = ResponsiveClass("has-text-centered")  // any text element
	TextJustified     = ResponsiveClass("has-text-justified") // any text element
	TextLeft          = ResponsiveClass("has-text-left")      // any text element
	TextRight         = ResponsiveClass("has-text-right")     // any text element
	WeightBold        = e.Class("has-text-weight-bold")       // any text element
	WeightLight       = e.Class("has-text-weight-light")      // any text element
	WeightMedium      = e.Class("has-text-weight-medium")     // any text element
	WeightNormal      = e.Class("has-text-weight-normal")     // any text element
	WeightSemiBold    = e.Class("has-text-weight-semibold")   // any text element
)

// "is-" classes with automatic adding of other class
var (
	GroupedCentered  = e.Classes{"is-grouped", "is-grouped-centered"}  // Field
	GroupedMultiline = e.Classes{"is-grouped", "is-grouped-multiline"} // Field
	GroupedRight     = e.Classes{"is-grouped", "is-grouped-right"}     // Field
	ToggleRounded    = e.Classes{"is-toggle", "is-toggle-rounded"}     // Tabs
)

// "has-" classes with automatic adding of other class
var (
	AddonsCentered = e.Classes{"has-addons", "has-addons-centered"} // Field
	AddonsRight    = e.Classes{"has-addons", "has-addons-right"}    // Field
)

func changeSizePrefix(prefix string, class e.Class) e.Class {
	if class == Small || class == Normal || class == Medium || class == Large {
		return e.Class(prefix) + class[3:]
	}
	return class
}
