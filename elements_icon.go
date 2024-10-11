package docs

import (
	"time"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var icon = c.NewPage(
	"Icon", "Icon", "/icon",
	"",
	b.Content(
		e.P(
			"The ", e.Code("b.Icon"), " constructor creates a container for an ", e.Strong("icon font"), ".",
		),
		c.Modifiers(
			c.Row("b.White", "Set color to white"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.BlackBis", "Set color to black bis"),
			c.Row("b.BlackTer", "Set color to black ter"),
			c.Row("b.GreyDarker", "Set color to grey darker"),
			c.Row("b.GreyDark", "Set color to grey dark"),
			c.Row("b.Grey", "Set color to grey"),
			c.Row("b.GreyLight", "Set color to grey light"),
			c.Row("b.GreyLigher", "Set color to grey lighter"),
			c.Row("b.WhiteTer", "Set color to white ter"),
			c.Row("b.WhiteBis", "Set color to white bis"),
			c.Row("b.PrimaryLight", "Set color to primary light"),
			c.Row("b.LinkLight", "Set color to link light"),
			c.Row("b.InfoLight", "Set color to info light"),
			c.Row("b.SuccessLight", "Set color to success light"),
			c.Row("b.WarningLight", "Set color to warning light"),
			c.Row("b.DangerLight", "Set color to danger light"),
			c.Row("b.PrimaryDark", "Set color to primary dark"),
			c.Row("b.LinkDark", "Set color to link dark"),
			c.Row("b.InfoDark", "Set color to info dark"),
			c.Row("b.SuccessDark", "Set color to success dark"),
			c.Row("b.WarningDark", "Set color to warning dark"),
			c.Row("b.DangerDark", "Set color to danger dark"),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),

		e.P(
			"The ", e.Code("b.IconText"), " constructor creates an icon+text span container and embeds all its non-icon children into spans. The ", e.Code("b.FlexIconText"), " constructor creates a flex icon+text span container and embeds all its non-icon children into spans.",
		),
		c.Modifiers(
			c.Row("b.White", "Set color to white"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.BlackBis", "Set color to black bis"),
			c.Row("b.BlackTer", "Set color to black ter"),
			c.Row("b.GreyDarker", "Set color to grey darker"),
			c.Row("b.GreyDark", "Set color to grey dark"),
			c.Row("b.Grey", "Set color to grey"),
			c.Row("b.GreyLight", "Set color to grey light"),
			c.Row("b.GreyLigher", "Set color to grey lighter"),
			c.Row("b.WhiteTer", "Set color to white ter"),
			c.Row("b.WhiteBis", "Set color to white bis"),
			c.Row("b.PrimaryLight", "Set color to primary light"),
			c.Row("b.LinkLight", "Set color to link light"),
			c.Row("b.InfoLight", "Set color to info light"),
			c.Row("b.SuccessLight", "Set color to success light"),
			c.Row("b.WarningLight", "Set color to warning light"),
			c.Row("b.DangerLight", "Set color to danger light"),
			c.Row("b.PrimaryDark", "Set color to primary dark"),
			c.Row("b.LinkDark", "Set color to link dark"),
			c.Row("b.InfoDark", "Set color to info dark"),
			c.Row("b.SuccessDark", "Set color to success dark"),
			c.Row("b.WarningDark", "Set color to warning dark"),
			c.Row("b.DangerDark", "Set color to danger dark"),
		),
	),
).Section(
	"Font Awesome", "",

	b.Content(
		e.P("The", e.Code("github.com/willoma/bulma-gomponents/fa"), " package provides helpers for ", e.Em("Font Awesome"), " icons."),
		e.P(
			"The ", e.Code("fa.Icon"), " constructor creates a ", e.Em("Font Awesome"), " icon, embedded in an icon container, dealing with appropriately applying children to the icon element or to the span container. The ", e.Code("fa.FA"), " constructor creates a ", e.Em("Font Awesome"), " icon only, without the icon container. These constructors accept the same modifiers as ", e.Code("b.Icon"), ", as well as the ", e.Em("Font Awesome"), "-specific values described in following sections.",
		),
		c.Children(
			c.Row("fa.OnFA(...any)", "Apply children to the Font Awesome ", e.Code("<i>"), " element"),
			c.Row("fa.OnIcon(...any)", "Apply children to the ", e.Code("<span>"), " element"),
			c.RowDefault("Apply children to the ", e.Code("<span>"), " element"),
		),
	),
).Subsection(
	"First arguments", "Font Awesome::https://docs.fontawesome.com/web/add-icons/how-to#families--styles",

	b.Content(
		e.P("The ", e.Code("fa.Icon"), " and ", e.Code("fa.FA"), " functions require at least two arguments:"),
		e.Ol(
			e.Li(
				"The font style, one of ", e.Code("fa.Brand"), ", ", e.Code("fa.Duotone"), ", ", e.Code("fa.Light"), ", ", e.Code("fa.SharpLight"), ", ", e.Code("fa.Regular"), ", ", e.Code("fa.SharpRegular"), ", ", e.Code("fa.Solid"), ", ", e.Code("fa.SharpSolid"), ", ", e.Code("fa.Thin"), " or ", e.Code("fa.SharpThin"),
			),
			e.Li(
				"The icon name, ", e.Em("without the ", e.Code("fa-"), " prefix"), " (for example, ", e.Code(`"home"`), " or ", e.Code(`"spinner"`), ")",
			),
		),
	),
).Subsection(
	"Styling basics", "Font Awesome::https://docs.fontawesome.com/web/style/basics",

	b.Content(
		e.P("Arguments provided to ", e.Code("fa.Icon"), " are automatically applied to the ", e.Code("<span>"), " or ", e.Code("<i>"), " element, depending on its nature. If you need to explicitly apply a child to the ", e.Code("<i>"), " element, use ", e.Code("fa.FA"), " and include it manually as a child of ", e.Code("b.Icon"), "."),
	),
	c.Example(
		`e.Span(
	e.Styles{"font-size", "3em", "color": "Tomato"},
	fa.FA(
		fa.Solid, "camera",
	),
),
" ",
e.Span(
	e.Styles{"font-size", "48px", "color": "Dodgerblue"},
	fa.FA(
		fa.Solid, "camera",
	),
),
" ",
e.Span(
	e.Styles{"font-size": "3rem"},
	e.Span(
		e.Styles{"color": "Mediumslateblue"},
		fa.FA(
			fa.Solid, "camera",
		),
	),
)`,
		e.Span(
			e.Styles{"font-size": "3em", "color": "Tomato"},
			fa.FA(
				fa.Solid, "camera",
			),
		),
		" ",
		e.Span(
			e.Styles{"font-size": "48px", "color": "Dodgerblue"},
			fa.FA(
				fa.Solid, "camera",
			),
		),
		" ",
		e.Span(
			e.Styles{"font-size": "3rem"},
			e.Span(
				e.Styles{"color": "Mediumslateblue"},
				fa.FA(
					fa.Solid, "camera",
				),
			),
		),
	),
).Subsection(
	"Sizing icons", "Font Awesome::https://docs.fontawesome.com/web/style/size",

	c.Modifiers(
		c.Row("fa.Size(2)s", "Set size to 2xs"),
		c.Row("fa.Size()s", "Set size to xs"),
		c.Row("fa.SizeSm", "Set size to sm"),
		c.Row("fa.SizeLg", "Set size to lg"),
		c.Row("fa.Size()l", "Set size to xl"),
		c.Row("fa.Size(2)l", "Set size to 2xl"),
		c.RowTo("fa.Size(0)", "fa.Size(10)", "Set size to the provided integer value"),
	),
).Subsection(
	"Fixed width icons", "Font Awesome::https://docs.fontawesome.com/web/style/fixed-width",

	b.Content(
		e.P("In order to make all your icons the same width so they can easily vertically align, like in a list or navigation menu, use the", e.Code("fa.FixedWidth"), " modifier."),
		e.P("However, please note that ", e.Strong(e.Em("Bulma"), " icons are already square"), ", with a fixed width."),
	),
	c.Example(
		`e.Div(fa.FA(fa.Solid, "skating", fa.FixedWidth, e.Styles{"background": "DodgerBlue")), " Skating"},
e.Div(fa.FA(fa.Solid, "skiing", fa.FixedWidth, e.Styles{"background": "SkyBlue")), " Skiing"},
e.Div(fa.FA(fa.Solid, "skiing-nordic", fa.FixedWidth, e.Styles{"background": "DodgerBlue")), " Nordic Skiing"},
e.Div(fa.FA(fa.Solid, "snowboarding", fa.FixedWidth, e.Styles{"background": "SkyBlue")), " Snowboarding"},
e.Div(fa.FA(fa.Solid, "snowplow", fa.FixedWidth, e.Styles{"background": "DodgerBlue")), " Snowplow"},`,
		e.Styles{"font-size": "1.5rem"},
		e.Div(fa.FA(fa.Solid, "skating", fa.FixedWidth, e.Styles{"background": "DodgerBlue"}), " Skating"),
		e.Div(fa.FA(fa.Solid, "skiing", fa.FixedWidth, e.Styles{"background": "SkyBlue"}), " Skiing"),
		e.Div(fa.FA(fa.Solid, "skiing-nordic", fa.FixedWidth, e.Styles{"background": "DodgerBlue"}), " Nordic Skiing"),
		e.Div(fa.FA(fa.Solid, "snowboarding", fa.FixedWidth, e.Styles{"background": "SkyBlue"}), " Snowboarding"),
		e.Div(fa.FA(fa.Solid, "snowplow", fa.FixedWidth, e.Styles{"background": "DodgerBlue"}), " Snowplow"),
	),
).Subsection(
	"Icons in a list", "Font Awesome::https://docs.fontawesome.com/web/style/lists",

	b.Content(
		e.P("Use ", e.Em("fa.UList"), " and ", e.Em("fa.Li"), " to replace default bullets in unordered lists. You can also keep the semantics of an ordered list behind the scenes but use icon bullets visually with ", e.Em("fa.OList"), "."),
	),
	c.Example(
		`fa.UList(
	fa.Li(fa.Solid, "check-square", "List icons can"),
	fa.Li(fa.Solid, "check-square", "be used to"),
	fa.Li(fa.Solid, "spinner", fa.Spin(fa.Pulse), "replace bullets"),
	fa.Li(fa.Regular, "square", "in lists"),
)`,
		fa.UList(
			fa.Li(fa.Solid, "check-square", "List icons can"),
			fa.Li(fa.Solid, "check-square", "be used to"),
			fa.Li(fa.Solid, "spinner", fa.Spin(fa.Pulse), "replace bullets"),
			fa.Li(fa.Regular, "square", "in lists"),
		),
	),
	c.Example(
		`fa.OList(
	fa.Li(fa.Solid, "check-square", "List icons can"),
	fa.Li(fa.Solid, "check-square", "be used to"),
	fa.Li(fa.Solid, "spinner", fa.Spin(fa.Pulse), "replace bullets"),
	fa.Li(fa.Regular, "square", "in lists"),
)`,
		fa.OList(
			fa.Li(fa.Solid, "check-square", "List icons can"),
			fa.Li(fa.Solid, "check-square", "be used to"),
			fa.Li(fa.Solid, "spinner", fa.Spin(fa.Pulse), "replace bullets"),
			fa.Li(fa.Regular, "square", "in lists"),
		),
	),
).Subsection(
	"Rotate icons", "Font Awesome::https://docs.fontawesome.com/web/style/rotate",

	b.Content(
		e.P(
			"To arbitrarily rotate and flip icons, use one of the following modifiers on ", e.Code("fa.Icon"), ", ", e.Code("fa.FA"), " or ", e.Code("fa.Li"), ":",
		),
		b.UList(
			e.Code("fa.Rotate90"),
			e.Code("fa.Rotate180"),
			e.Code("fa.Rotate270"),
			e.Code("fa.FlipHorizontal"),
			e.Code("fa.FlipVertical"),
			e.Code("fa.FlipBoth"),
		),
	),
	c.Example(
		`fa.FA(fa.Solid, "snowboarding"), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate90), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate180), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate270), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipHorizontal), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipVertical), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipBoth)`,
		fa.Size(2),
		fa.FA(fa.Solid, "snowboarding"), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate90), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate180), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate270), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipHorizontal), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipVertical), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipBoth),
	),
	b.Content(
		e.P("In order to apply a custom rotation, use a ", e.Code("fa.Rotate"), " value."),
	),
	c.Example(
		`fa.FA(fa.Solid, "snowboarding"), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(45)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(-45)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(19)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(80)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(0.25)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(-0.25))`,
		fa.Size(2),
		fa.FA(fa.Solid, "snowboarding"), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(45)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(-45)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(19)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(80)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(0.25)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(-0.25)),
	),
	b.Content(
		e.P("When combining rotation and flipping, the ", e.Code("<i>"), " is automatically embedded in a ", e.Code("<span>"), " in order to apply all transformations."),
	),
	c.Example(
		`fa.FA(fa.Solid, "snowboarding", fa.Rotate90, fa.FlipHorizontal), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipHorizontal, fa.Rotate90), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipVertical, fa.Rotate270), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate270, fa.FlipVertical), " ",
fa.FA(fa.Solid, "snowboarding", fa.FlipBoth, fa.Rotate(120)), " ",
fa.FA(fa.Solid, "snowboarding", fa.Rotate(20), fa.FlipBoth), " "`,
		fa.Size(2),
		fa.FA(fa.Solid, "snowboarding", fa.Rotate90, fa.FlipHorizontal), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipHorizontal, fa.Rotate90), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipVertical, fa.Rotate270), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate270, fa.FlipVertical), " ",
		fa.FA(fa.Solid, "snowboarding", fa.FlipBoth, fa.Rotate(120)), " ",
		fa.FA(fa.Solid, "snowboarding", fa.Rotate(20), fa.FlipBoth), " ",
	),
).Subsection(
	"Animating icons", "Font Awesome::https://docs.fontawesome.com/web/style/animate",

	b.Content(
		e.P("The icons can be animated using one of the structures implementing ", e.Code("fa.Animation"), ". Values are rounded to the 2nd decimal."),
		c.Modifiers(
			c.Row("fa.Delay(time.Duration)", "Add an initial delay"),
			c.Row("fa.DirectionNormal", "Change the animation direction to normal"),
			c.Row("fa.DirectionReverse", "Change the animation direction to reverse"),
			c.Row("fa.DirectionAlternate", "Change the animation direction to alternate"),
			c.Row("fa.DirectionAlternateReverse", "Change the animation direction to alternate reverse"),
			c.Row("fa.Duration(time.Duration)", "Set duration"),
			c.Row("fa.IterationCount(float64)", "Set number of iterations"),
			c.Row("fa.TimingEase", "Set animation progress timing to ease"),
			c.Row("fa.TimingLinear", "Set animation progress timing to linear"),
			c.Row("fa.TimingEaseIn", "Set animation progress timing to ease-in"),
			c.Row("fa.TimingEaseOut", "Set animation progress timing to ease-out"),
			c.Row("fa.TimingEaseInOut", "Set animation progress timing to ease-in-out"),
			c.Row("fa.TimingStepStart", "Set animation progress timing to step-start"),
			c.Row("fa.TimingStepEnd", "Set animation progress timing to step-end"),
		),
	),
	b.Content(
		e.P("The ", e.Code("fa.Beat"), " animation type scales an icon up or down. Additionally to the common attributes, it uses the ", e.Code("fa.MaxScale(float64)"), " attribute to set the maximum value the icon will scale."),
	),
	c.Example(
		`fa.FA(fa.Solid, "circle-plus", fa.Beat()), " ",
fa.FA(fa.Solid, "heart", fa.Beat()), " ",
fa.FA(fa.Solid, "heart", fa.Beat(fa.Duration(500*time.Millisecond))), " ",
fa.FA(fa.Solid, "heart", fa.Beat(fa.Duration(2*time.Second))), " ",
fa.FA(fa.Solid, "heart", fa.Beat(fa.MaxScale(2))),`,
		fa.Size(2),
		fa.FA(fa.Solid, "circle-plus", fa.Beat()), " ",
		fa.FA(fa.Solid, "heart", fa.Beat()), " ",
		fa.FA(fa.Solid, "heart", fa.Beat(fa.Duration(500*time.Millisecond))), " ",
		fa.FA(fa.Solid, "heart", fa.Beat(fa.Duration(2*time.Second))), " ",
		fa.FA(fa.Solid, "heart", fa.Beat(fa.MaxScale(2))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Fade"), " animation type fades an icon in and out. Additionally to the common attributes, it uses the ", e.Code("fa.MinOpacity(float64)"), " attribute to set the lowest opacity the icon will fade to and from."),
	),
	c.Example(
		`fa.FA(fa.Solid, "triangle-exclamation", fa.Fade(), " ",
fa.FA(fa.Solid, "skull-crossbones", fa.Fade()), " ",
fa.FA(fa.Solid, "desktop", fa.Fade()), " ",
fa.FA(fa.Solid, "i-cursor", fa.Fade(fa.Duration(2 * time.Second), fa.MinOpacity(0.6)))`,
		fa.Size(2),
		fa.FA(fa.Solid, "triangle-exclamation", fa.Fade()), " ",
		fa.FA(fa.Solid, "skull-crossbones", fa.Fade()), " ",
		fa.FA(fa.Solid, "desktop", fa.Fade()), " ",
		fa.FA(fa.Solid, "i-cursor", fa.Fade(fa.Duration(2*time.Second), fa.MinOpacity(0.6))),
	),
	b.Content(
		e.P("The ", e.Code("fa.BeatFade"), " animation type scales and pulses an icon in and out. Additionally to the common attributes, it uses the ", e.Code("fa.MaxScale(float64)"), "and ", e.Code("fa.MinOpacity(float64)"), " attributes."),
	),
	c.Example(
		`fa.FA(fa.Solid, "triangle-exclamation", fa.BeatFade()), " ",
fa.FA(fa.Solid, "square-xmark", fa.BeatFade()), " ",
fa.FA(fa.Solid, "poo-bolt", fa.BeatFade(fa.MinOpacity(0.1), fa.MaxScale(1.25))), " ",
fa.FA(fa.Solid, "circle-info", fa.BeatFade(fa.MinOpacity(0.67), fa.MaxScale(1.075)))`,
		fa.Size(2),
		fa.FA(fa.Solid, "triangle-exclamation", fa.BeatFade()), " ",
		fa.FA(fa.Solid, "square-xmark", fa.BeatFade()), " ",
		fa.FA(fa.Solid, "poo-bolt", fa.BeatFade(fa.MinOpacity(0.1), fa.MaxScale(1.25))), " ",
		fa.FA(fa.Solid, "circle-info", fa.BeatFade(fa.MinOpacity(0.67), fa.MaxScale(1.075))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Bounce"), " animation type bounces an icon up and down. Additionally to the common attributes, it uses the following attributes:"),
		b.DList(
			c.Row("fa.Rebound(float64)", "Set the amount of rebound when landing after the jump"), c.Row("fa.Height(float64)", "Set the max height to jump when bouncing"), c.Row("fa.StartScaleX(float64)", "Set the icon's horizontal distortion (squish) when starting to bounce"), c.Row("fa.StartScaleY(float64)", "Set the icon's vertical distortion (squish) when starting to bounce"), c.Row("fa.JumpScaleX(float64)", "Set the icon's horizontal distortion (squish) at the top of the jump"), c.Row("fa.JumpScaleY(float64)", "Set the icon's vertical distortion (squish) at the top of the jump"), c.Row("fa.LandScaleX(float64)", "Set the icon's horizontal distortion (squish) when landing after the jump"), c.Row("fa.LandScaleY(float64)", "Set the icon's vertical distortion (squish) when landing after the jump")),
	),
	c.Example(
		`fa.FA(fa.Solid, "basketball", fa.Bounce()), " ",
fa.FA(fa.Solid, "volleyball", fa.Bounce()), " ",
fa.FA(fa.Solid, "frog", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1))), " ",
fa.FA(fa.Solid, "envelope", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1), fa.Rebound(0)))`,
		fa.Size(2),
		fa.FA(fa.Solid, "basketball", fa.Bounce()), " ",
		fa.FA(fa.Solid, "volleyball", fa.Bounce()), " ",
		fa.FA(fa.Solid, "frog", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1))), " ",
		fa.FA(fa.Solid, "envelope", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1), fa.Rebound(0))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Flip"), " animation type rotates an icon in 3D space. By default, it rotates about the Y axis 180 degrees. Additionally to the common attributes, it uses the ", e.Code("fa.X(float64)"), ", ", e.Code("fa.Y(float64)"), ", ", e.Code("fa.Z(float64)"), " attributes to denote the axis of rotation (between ", e.Code("0"), " and ", e.Code("1"), ") and the ", e.Code("fa.Angle(float64)"), " attribute to set the rotation angle."),
	),
	c.Example(
		`fa.FA(fa.Solid, "compact-disc", fa.Flip()), " ",
fa.FA(fa.Solid, "camera-rotate", fa.Flip()), " ",
fa.FA(fa.Solid, "floppy-disk", fa.Flip()), " ",
fa.FA(fa.Solid, "scroll", fa.Flip(fa.X(1), fa.Y(0))), " ",
fa.FA(fa.Solid, "money-check-dollar", fa.Flip(fa.Duration(3*time.Second)))`,
		fa.Size(2),
		fa.FA(fa.Solid, "compact-disc", fa.Flip()), " ",
		fa.FA(fa.Solid, "camera-rotate", fa.Flip()), " ",
		fa.FA(fa.Solid, "floppy-disk", fa.Flip()), " ",
		fa.FA(fa.Solid, "scroll", fa.Flip(fa.X(1), fa.Y(0))), " ",
		fa.FA(fa.Solid, "money-check-dollar", fa.Flip(fa.Duration(3*time.Second))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Shake"), " animation type shakes back and forth"),
	),
	c.Example(
		`fa.FA(fa.Solid, "bell", fa.Shake()), " ",
fa.FA(fa.Solid, "lock", fa.Shake()), " ",
fa.FA(fa.Solid, "stopwatch", fa.Shake()), " ",
fa.FA(fa.Solid, "bomb", fa.Shake())`,
		fa.Size(2),
		fa.FA(fa.Solid, "bell", fa.Shake()), " ",
		fa.FA(fa.Solid, "lock", fa.Shake()), " ",
		fa.FA(fa.Solid, "stopwatch", fa.Shake()), " ",
		fa.FA(fa.Solid, "bomb", fa.Shake()),
	),
	b.Content(
		e.P("The ", e.Code("fa.Spin"), " animation type make the icon spin 360° clockwise. Additionally to the common attributes, it uses the ", e.Code("fa.Pulse"), " attribute to make the rotation in 8 steps and the ", e.Code("fa.Reverse"), " attribute to make the icon spin counter-clockwise."),
	),
	c.Example(
		`fa.FA(fa.Solid, "sync", fa.Spin()), " ",
fa.FA(fa.Solid, "circle-notch", fa.Spin()), " ",
fa.FA(fa.Solid, "cog", fa.Spin()), " ",
fa.FA(fa.Solid, "cog", fa.Spin(fa.Reverse)), " ",
fa.FA(fa.Solid, "spinner", fa.Spin(fa.Pulse)), " ",
fa.FA(fa.Solid, "spinner", fa.Spin(fa.Pulse, fa.Reverse))`,
		fa.Size(2),
		fa.FA(fa.Solid, "sync", fa.Spin()), " ",
		fa.FA(fa.Solid, "circle-notch", fa.Spin()), " ",
		fa.FA(fa.Solid, "cog", fa.Spin()), " ",
		fa.FA(fa.Solid, "cog", fa.Spin(fa.Reverse)), " ",
		fa.FA(fa.Solid, "spinner", fa.Spin(fa.Pulse)), " ",
		fa.FA(fa.Solid, "spinner", fa.Spin(fa.Pulse, fa.Reverse)),
	),
).Subsection(
	"Bordered and pulled icons", "Font Awesome::https://docs.fontawesome.com/web/style/pull",

	b.Content(
		e.P("The ", e.Code("fa.Border"), "modifier creates a border with radius and padding around an icon. The ", e.Code("fa.PullLeft"), " modifier pulls the icon by floating it left and applying a right margin. The ", e.Code("fa.PullRight"), " modifier pulls the icon by floating it right and applying a left margin."),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "quote-left", fa.Size(2), fa.PullLeft),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		e.P(
			fa.FA(fa.Solid, "quote-left", fa.Size(2), fa.PullLeft),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size(2), fa.PullRight, fa.Border),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		e.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size(2), fa.PullRight, fa.Border),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	b.Content(
		e.P("The following modifiers allow customizing the border and pulling effects:"),
		b.DList(
			c.Row("fa.BorderColor", "Set border color"), c.Row("fa.BorderPadding", "Set padding around icon"), c.Row("fa.BorderRadius", "Set border radius"), c.Row("fa.BorderStyle", "Set border style"), c.Row("fa.BorderWidth", "Set border width"), c.Row("fa.PullMargin", "Set margin around icon")),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "quote-left", fa.Size(2), fa.PullLeft, fa.PullMargin("4em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		e.P(
			fa.FA(fa.Solid, "quote-left", fa.Size(2), fa.PullLeft, fa.PullMargin("4em")),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size(2), fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		e.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size(2), fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
).Subsection(
	"Stacking icons", "Font Awesome::https://docs.fontawesome.com/web/style/stack",

	b.Content(
		e.P("To stack multiple icons, use the ", e.Code("fa.Stack"), " component builder. Then add the ", e.Code("fa.Stack1x"), " modifier to the regularly sized icon and add the ", e.Code("fa.Stack1x"), " modifier to the larger icon. ", e.Code("fa.Inverse"), " can be added to the icon with ", e.Code("fa.Stack1x"), " to help with a knock-out looking effect."),
	),

	c.Example(
		`fa.Stack(
	fa.FA(fa.Solid, "square", fa.Stack2x),
	fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse),
), " ",
fa.Stack(
	fa.FA(fa.Solid, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
), " ",
fa.Stack(
	fa.FA(fa.Solid, "camera", fa.Stack1x),
	fa.FA(fa.Solid, "ban", fa.Stack2x, e.Styles{"color": "Tomato"}),
), " ",
fa.Stack(
	fa.FA(fa.Solid, "square", fa.Stack2x),
	fa.FA(fa.Solid, "terminal", fa.Stack1x, fa.Inverse),
), " ",
fa.Stack(
	fa.Size(2),
	fa.FA(fa.Solid, "square", fa.Stack2x),
	fa.FA(fa.Solid, "terminal", fa.Stack1x, fa.Inverse),
)`,
		fa.Stack(
			fa.FA(fa.Solid, "square", fa.Stack2x),
			fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse),
		), " ",
		fa.Stack(
			fa.FA(fa.Solid, "circle", fa.Stack2x),
			fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
		), " ",
		fa.Stack(
			fa.FA(fa.Solid, "camera", fa.Stack1x),
			fa.FA(fa.Solid, "ban", fa.Stack2x, e.Styles{"color": "Tomato"}),
		), " ",
		fa.Stack(
			fa.FA(fa.Solid, "square", fa.Stack2x),
			fa.FA(fa.Solid, "terminal", fa.Stack1x, fa.Inverse),
		), " ",
		fa.Stack(
			fa.Size(2),
			fa.FA(fa.Solid, "square", fa.Stack2x),
			fa.FA(fa.Solid, "terminal", fa.Stack1x, fa.Inverse),
		),
	),
	c.Example(
		`fa.FA(fa.Regular, "circle", fa.Size(2)), " ",
fa.Stack(
	fa.FA(fa.Regular, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x),
), " ",
fa.Stack(
	e.Styles{"vertical-align": "top"},
	fa.FA(fa.Solid, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
), " ",
fa.FA(fa.Regular, "circle", fa.Size(2))`,
		fa.FA(fa.Regular, "circle", fa.Size(2)), " ",
		fa.Stack(
			e.Styles{"vertical-align": "top"},
			fa.FA(fa.Regular, "circle", fa.Stack2x),
			fa.FA(fa.Solid, "flag", fa.Stack1x),
		), " ",
		fa.Stack(
			e.Styles{"vertical-align": "top"},
			fa.FA(fa.Solid, "circle", fa.Stack2x),
			fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
		), " ",
		fa.FA(fa.Regular, "circle", fa.Size(2)),
	),
	c.Example(
		`fa.FA(fa.Regular, "circle", e.Styles{"vertical-align": "middle"}), " ",
fa.Stack(
	e.Styles{"font-size": "0.5em"},
	fa.FA(fa.Regular, "circle", fa.Stack2x, e.Styles{"vertical-align": "middle"}),
	fa.FA(fa.Solid, "flag", fa.Stack1x, e.Styles{"vertical-align": "middle"}),
), " ",
fa.Stack(
	e.Styles{"font-size": "0.5em"},
	fa.FA(fa.Solid, "circle", fa.Stack2x, e.Styles{"vertical-align": "middle"}),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse, e.Styles{"vertical-align": "middle"}),
), " ",
fa.FA(fa.Regular, "circle", e.Styles{"vertical-align": "middle"})`,
		fa.FA(fa.Regular, "circle", e.Styles{"vertical-align": "middle"}), " ",
		fa.Stack(
			e.Styles{"font-size": "0.5em"},
			fa.FA(fa.Regular, "circle", fa.Stack2x, e.Styles{"vertical-align": "middle"}),
			fa.FA(fa.Solid, "flag", fa.Stack1x, e.Styles{"vertical-align": "middle"}),
		), " ",
		fa.Stack(
			e.Styles{"font-size": "0.5em"},
			fa.FA(fa.Solid, "circle", fa.Stack2x, e.Styles{"vertical-align": "middle"}),
			fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse, e.Styles{"vertical-align": "middle"}),
		), " ",
		fa.FA(fa.Regular, "circle", e.Styles{"vertical-align": "middle"}),
	),
	b.Content(
		e.P("The following modifiers allow customizing stacked icons:"),
		b.DList(
			c.Row("fa.ZIndex", "Set z-index of a stacked icon"), c.Row("fa.InverseColor", "Set color of an inversed stacked icon")),
	),
	c.Example(
		`fa.Stack(
	fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse, fa.ZIndex(2)),
	fa.FA(fa.Solid, "square", fa.Stack2x, fa.ZIndex(1)),
), " ",
fa.Stack(
	fa.InverseColor("#1da1f2"),
	fa.FA(fa.Solid, "square", fa.Stack2x,
	fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse),
)`,
		fa.Stack(
			fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse, fa.ZIndex(2)),
			fa.FA(fa.Solid, "square", fa.Stack2x, fa.ZIndex(1)),
		), " ",
		fa.Stack(
			fa.InverseColor("#1da1f2"),
			fa.FA(fa.Solid, "square", fa.Stack2x),
			fa.FA(fa.Brand, "twitter", fa.Stack1x, fa.Inverse),
		),
	),
).Subsection(
	"Duotone icons", "Font Awesome::https://docs.fontawesome.com/web/style/duotone",

	b.Content(
		e.P("The following modifiers allow customizing duotone icons:"),
		c.Modifiers(
			c.Row("fa.SwapOpacity", "Swap the default opacity of each icon's layers"),
			c.Row("fa.PrimaryOpacity(float64)", "Set primary layer opacity"),
			c.Row("fa.SecondaryOpacity(float64)", "Set secondary layer opacity"),
			c.Row("fa.DuoOpacities(primary, secondary float64)", "Set primary and secondary layers opacities"),
			c.Row("fa.PrimaryColor(string)", "Set primary layer color"),
			c.Row("fa.SecondaryColor(string)", "Set secondary layer color"),
			c.Row("fa.DuoColors(primary, secondary string)", "Set primary and secondary colors"),
		),
	),
	b.Message(
		b.Danger,
		"Please note these icons require pro plan, which ", e.Em("Bulma-Gomponents"), " do not have, therefore no example may be shown.",
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/icon/",

	c.Example(
		`b.Icon(e.I(e.Class("fas"), e.Class("fa-home"))),
b.Icon(fa.FA(fa.Solid, "home")),
fa.Icon(fa.Solid, "home")`,
		b.Icon(e.I(e.Class("fas"), e.Class("fa-home"))),
		b.Icon(fa.FA(fa.Solid, "home")),
		fa.Icon(fa.Solid, "home"),
	),
).Subsection(
	"Icon text",
	"https://bulma.io/documentation/elements/icon/#icon-text",
	c.Example(
		`b.IconText(
	fa.Icon(fa.Solid, "home"),
	"Home",
)`,
		b.IconText(
			fa.Icon(fa.Solid, "home"),
			"Home",
		),
	),
	c.Example(
		`b.IconText(
	fa.Icon(fa.Solid, "train"),
	"Paris",
	fa.Icon(fa.Solid, "arrow-right"),
	"Budapest",
	fa.Icon(fa.Solid, "arrow-right"),
	"Bucharest",
	fa.Icon(fa.Solid, "arrow-right"),
	"Istanbul",
	fa.Icon(fa.Solid, "flag-checkered"),
)`,
		b.IconText(
			fa.Icon(fa.Solid, "train"),
			"Paris",
			fa.Icon(fa.Solid, "arrow-right"),
			"Budapest",
			fa.Icon(fa.Solid, "arrow-right"),
			"Bucharest",
			fa.Icon(fa.Solid, "arrow-right"),
			"Istanbul",
			fa.Icon(fa.Solid, "flag-checkered"),
		),
	),
	c.Example(
		`b.Content(
	e.P(
		"An invitation to ", b.IconText(fa.Icon(fa.Solid, "utensils"), "dinner"), " was soon afterwards dispatched; and already had Mrs. Bennet planned the courses that were to do credit to her housekeeping, when an answer arrived which deferred it all. Mr. Bingley was obliged to be in ", b.IconText(fa.Icon(fa.Solid, "city"), "town"), " the following day, and, consequently, unable to accept the honour of their ", b.IconText(fa.Icon(fa.Solid, "envelope-open-text"), "invitation"), ", etc.",
	),
	e.P(
		"Mrs. Bennet was quite disconcerted. She could not imagine what business he could have in town so soon after his ", b.IconText(fa.Icon(fa.Solid, "flag-checkered"), "arrival"), " in Hertfordshire; and she began to fear that he might be always ", b.IconText(fa.Icon(fa.Solid, "plane-departure"), "flying"), " about from one place to another, and never settled at Netherfield as he ought to be.",
	),
)`,
		b.Content(
			e.P(
				"An invitation to ", b.IconText(fa.Icon(fa.Solid, "utensils"), "dinner"), " was soon afterwards dispatched; and already had Mrs. Bennet planned the courses that were to do credit to her housekeeping, when an answer arrived which deferred it all. Mr. Bingley was obliged to be in ", b.IconText(fa.Icon(fa.Solid, "city"), "town"), " the following day, and, consequently, unable to accept the honour of their ", b.IconText(fa.Icon(fa.Solid, "envelope-open-text"), "invitation"), ", etc.",
			),
			e.P(
				"Mrs. Bennet was quite disconcerted. She could not imagine what business he could have in town so soon after his ", b.IconText(fa.Icon(fa.Solid, "flag-checkered"), "arrival"), " in Hertfordshire; and she began to fear that he might be always ", b.IconText(fa.Icon(fa.Solid, "plane-departure"), "flying"), " about from one place to another, and never settled at Netherfield as he ought to be.",
			),
		),
	),
	c.Example(
		`b.FlexIconText(
			fa.Icon(fa.Solid, "info-circle", b.Info),
			"Information",
		),
		b.Block(
			html.P,
			"Your package will be delivered on ", e.Strong("Tuesday at 08:00"), ".",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "check-square", b.Success),
			"Success",
		),
		b.Block(
			html.P,
			"Your image has been successfully uploaded.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "exclamation-triangle", b.Warning),
			"Warning",
		),
		b.Block(
			html.P,
			"Some information is missing from your ", e.AHref("#", "profile"), " details.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "ban", b.Danger),
			"Danger",
		),
		b.Block(
			html.P,
			"There was an error in your submission. ", e.AHref("#", "Please try again"), ".",
		)`,
		b.FlexIconText(
			fa.Icon(fa.Solid, "info-circle", b.Info),
			"Information",
		),
		b.Block(
			html.P,
			"Your package will be delivered on ", e.Strong("Tuesday at 08:00"), ".",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "check-square", b.Success),
			"Success",
		),
		b.Block(
			html.P,
			"Your image has been successfully uploaded.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "exclamation-triangle", b.Warning),
			"Warning",
		),
		b.Block(
			html.P,
			"Some information is missing from your ", e.AHref("#", "profile"), " details.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "ban", b.Danger),
			"Danger",
		),
		b.Block(
			html.P,
			"There was an error in your submission. ", e.AHref("#", "Please try again"), ".",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/icon/#colors",
	c.Example(
		`b.Icon(b.Info, e.I(e.Class("fas"), e.Class("fa-info-circle"))),
b.Icon(b.Success, e.I(e.Class("fas"), e.Class("fa-check-square"))),
b.Icon(b.Warning, e.I(e.Class("fas"), e.Class("fa-exclamation-triangle"))),
b.Icon(b.Danger, e.I(e.Class("fas"), e.Class("fa-ban")))`,
		b.Icon(b.Info, e.I(e.Class("fas"), e.Class("fa-info-circle"))),
		b.Icon(b.Success, e.I(e.Class("fas"), e.Class("fa-check-square"))),
		b.Icon(b.Warning, e.I(e.Class("fas"), e.Class("fa-exclamation-triangle"))),
		b.Icon(b.Danger, e.I(e.Class("fas"), e.Class("fa-ban"))),
	),
	c.Example(
		`fa.Icon(fa.Solid, "info-circle", b.Info),
fa.Icon(fa.Solid, "check-square", b.Success),
fa.Icon(fa.Solid, "exclamation-triangle", b.Warning),
fa.Icon(fa.Solid, "ban", b.Danger)`,
		fa.Icon(fa.Solid, "info-circle", b.Info),
		fa.Icon(fa.Solid, "check-square", b.Success),
		fa.Icon(fa.Solid, "exclamation-triangle", b.Warning),
		fa.Icon(fa.Solid, "ban", b.Danger),
	),
	c.Example(
		`b.IconText(
	b.Info,
	fa.Icon(fa.Solid, "info-circle"),
	"Info",
),
" ",
b.IconText(
	b.Success,
	fa.Icon(fa.Solid, "check-square"),
	"Success",
),
" ",
b.IconText(
	b.Warning,
	fa.Icon(fa.Solid, "exclamation-triangle"),
	"Warning",
),
" ",
b.IconText(
	b.Danger,
	fa.Icon(fa.Solid, "ban"),
	"Danger",
)`,
		b.IconText(
			b.Info,
			fa.Icon(fa.Solid, "info-circle"),
			"Info",
		),
		" ",
		b.IconText(
			b.Success,
			fa.Icon(fa.Solid, "check-square"),
			"Success",
		),
		" ",
		b.IconText(
			b.Warning,
			fa.Icon(fa.Solid, "exclamation-triangle"),
			"Warning",
		),
		" ",
		b.IconText(
			b.Danger,
			fa.Icon(fa.Solid, "ban"),
			"Danger",
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/icon/#sizes",
	b.Content(
		e.P("Use a size modifier (", e.Code("b.Small"), ", etc) as an argument to ", e.Code("b.Icon"), " or ", e.Code("fa.Icon"), "."),
	),
).Subsection(
	"Font Awesome variations",
	"https://bulma.io/documentation/elements/icon/#font-awesome-variations",

	c.Example(
		`fa.Icon(fa.Solid, "home", fa.FixedWidth, b.BackgroundWarningLight),
e.Br(),
fa.Icon(fa.Solid, "home", fa.Border),
e.Br(),
fa.Icon(fa.Solid, "spinner", fa.Spin(fa.Pulse))`,
		fa.Icon(fa.Solid, "home", fa.FixedWidth, b.BackgroundWarningLight),
		e.Br(),
		fa.Icon(fa.Solid, "home", fa.Border),
		e.Br(),
		fa.Icon(fa.Solid, "spinner", fa.Spin(fa.Pulse)),
	),
	c.Example(
		`b.Icon(
	b.Medium,
	fa.Stack(
		fa.SizeSm,
		fa.FA(fa.Solid, "circle", fa.Stack2x),
		fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
	),
)`,
		b.Icon(
			b.Medium,
			fa.Stack(
				fa.SizeSm,
				fa.FA(fa.Solid, "circle", fa.Stack2x),
				fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
			),
		),
	),
	c.Example(
		`b.Icon(
	b.Large,
	fa.Stack(
		fa.FA(fa.Solid, "camera", fa.Stack1x),
		fa.FA(fa.Solid, "ban", fa.Stack2x, b.Danger),
		fa.SizeLg,
	),
)`,
		b.Icon(
			b.Large,
			fa.Stack(
				fa.FA(fa.Solid, "camera", fa.Stack1x),
				fa.FA(fa.Solid, "ban", fa.Stack2x, b.Danger),
				fa.SizeLg,
			),
		),
	),
).Subsection(
	"Material Design Icons",
	"https://bulma.io/documentation/elements/icon/#material-design-icons",
	b.Content("Nothing specific here."),
).Subsection(
	"Ionicons",
	"https://bulma.io/documentation/elements/icon/#ionicons",
	b.Content("Nothing specific here."),
)
