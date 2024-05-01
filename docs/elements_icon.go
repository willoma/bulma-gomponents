package docs

import (
	"time"

	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var icon = c.NewPage(
	"Icon", "Icon", "/icon",
	"",
	b.Content(
		e.P(
			"The ", e.Code("b.Icon"), " constructor returns a container for an ", e.Strong("icon font"), ". The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.White"),
			"Set icon color to white",

			e.Code("b.Black"),
			"Set icon color to black",

			e.Code("b.Light"),
			"Set icon color to light",

			e.Code("b.Dark"),
			"Set icon color to dark",

			e.Code("b.Primary"),
			"Set icon color to primary",

			e.Code("b.Link"),
			"Set icon color to link",

			e.Code("b.Info"),
			"Set icon color to info",

			e.Code("b.Success"),
			"Set icon color to success",

			e.Code("b.Warning"),
			"Set icon color to warning",

			e.Code("b.Danger"),
			"Set icon color to danger",

			e.Code("b.BlackBis"),
			"Set icon color to black bis",

			e.Code("b.BlackTer"),
			"Set icon color to black ter",

			e.Code("b.GreyDarker"),
			"Set icon color to grey darker",

			e.Code("b.GreyDark"),
			"Set icon color to grey dark",

			e.Code("b.Grey"),
			"Set icon color to grey",

			e.Code("b.GreyLight"),
			"Set icon color to grey light",

			e.Code("b.GreyLigher"),
			"Set icon color to grey lighter",

			e.Code("b.WhiteTer"),
			"Set icon color to white ter",

			e.Code("b.WhiteBis"),
			"Set icon color to white bis",

			e.Code("b.PrimaryLight"),
			"Set icon color to primary light",

			e.Code("b.LinkLight"),
			"Set icon color to link light",

			e.Code("b.InfoLight"),
			"Set icon color to info light",

			e.Code("b.SuccessLight"),
			"Set icon color to success light",

			e.Code("b.WarningLight"),
			"Set icon color to warning light",

			e.Code("b.DangerLight"),
			"Set icon color to danger light",

			e.Code("b.PrimaryDark"),
			"Set icon color to primary dark",

			e.Code("b.LinkDark"),
			"Set icon color to link dark",

			e.Code("b.InfoDark"),
			"Set icon color to info dark",

			e.Code("b.SuccessDark"),
			"Set icon color to success dark",

			e.Code("b.WarningDark"),
			"Set icon color to warning dark",

			e.Code("b.DangerDark"),
			"Set icon color to danger dark",

			e.Code("b.Small"),
			"Set icon size to small",

			e.Code("b.Medium"),
			"Set icon size to medium",

			e.Code("b.Large"),
			"Set icon size to large",
		),

		e.P(
			"The ", e.Code("b.IconText"), " constructor creates an icon+text span container and embeds all its non-icon children into spans. The ", e.Code("b.FlexIconText"), " constructor creates a flex icon+text span container and embeds all its non-icon children into spans. The following children have a special meaning:",
		),
		b.DList(e.Code("b.White"),
			"Set icon color to white",

			e.Code("b.Black"),
			"Set icon color to black",

			e.Code("b.Light"),
			"Set icon color to light",

			e.Code("b.Dark"),
			"Set icon color to dark",

			e.Code("b.Primary"),
			"Set icon color to primary",

			e.Code("b.Link"),
			"Set icon color to link",

			e.Code("b.Info"),
			"Set icon color to info",

			e.Code("b.Success"),
			"Set icon color to success",

			e.Code("b.Warning"),
			"Set icon color to warning",

			e.Code("b.Danger"),
			"Set icon color to danger",

			e.Code("b.BlackBis"),
			"Set icon color to black bis",

			e.Code("b.BlackTer"),
			"Set icon color to black ter",

			e.Code("b.GreyDarker"),
			"Set icon color to grey darker",

			e.Code("b.GreyDark"),
			"Set icon color to grey dark",

			e.Code("b.Grey"),
			"Set icon color to grey",

			e.Code("b.GreyLight"),
			"Set icon color to grey light",

			e.Code("b.GreyLigher"),
			"Set icon color to grey lighter",

			e.Code("b.WhiteTer"),
			"Set icon color to white ter",

			e.Code("b.WhiteBis"),
			"Set icon color to white bis",

			e.Code("b.PrimaryLight"),
			"Set icon color to primary light",

			e.Code("b.LinkLight"),
			"Set icon color to link light",

			e.Code("b.InfoLight"),
			"Set icon color to info light",

			e.Code("b.SuccessLight"),
			"Set icon color to success light",

			e.Code("b.WarningLight"),
			"Set icon color to warning light",

			e.Code("b.DangerLight"),
			"Set icon color to danger light",

			e.Code("b.PrimaryDark"),
			"Set icon color to primary dark",

			e.Code("b.LinkDark"),
			"Set icon color to link dark",

			e.Code("b.InfoDark"),
			"Set icon color to info dark",

			e.Code("b.SuccessDark"),
			"Set icon color to success dark",

			e.Code("b.WarningDark"),
			"Set icon color to warning dark",

			e.Code("b.DangerDark"),
			"Set icon color to danger dark",
		),
	),
).Section(
	"Font Awesome", "",

	b.Content(
		e.P("The", e.Code("github.com/willoma/bulma-gomponents/fa"), " package provides helpers for ", e.Em("Font Awesome"), " icons."),
		e.P(
			"The ", e.Code("fa.Icon"), " constructor creates a ", e.Em("Font Awesome"), " icon, embedded in an icon container, dealing with appropriately applying children to the icon e.Element or to the span container. It accepts the same values as ", e.Code("b.Icon"), ", as well as the ", e.Em("Font Awesome"), "-specific values described below. The following children have a special meaning:",
		),
		b.DList(
			e.Code("fa.OnFA(...)"),
			[]any{"Force childen to be applied to the Font Awesome ", e.Code("<i>"), " e.Element"},

			e.Code("fa.OnIcon(...)"),
			[]any{"Force childen to be applied to the icon span"},
		),
		e.P("Other children are added to the icon span."),
		e.P(
			"The ", e.Code("fa.FA"), " constructor creates a ", e.Em("Font Awesome"), " icon. It accepts the ", e.Em("Font Awesome"), "-specific values described below.",
		),
	),
).Subsection(
	"First arguments", "Font Awesome::https://docs.fontawesome.com/web/add-icons/how-to#families--styles",

	b.Content(
		e.P("The ", e.Code("fa.Icon"), " function requires at least two arguments:"),
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
		e.P("Arguments provided to ", e.Code("fa.Icon"), " are automatically applied to the ", e.Code("<span>"), " or ", e.Code("<i>"), " e.Element, depending on its nature. If you need to explicitly apply a child to the ", e.Code("<i>"), " e.Element, use ", e.Code("fa.FA"), " and include it manually as a child of ", e.Code("b.Icon"), "."),
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

	b.Content(
		e.P(
			"The following size modifiers are available:",
		),
		b.UList(
			b.MarginVertical(b.Spacing0),
			e.Code("fa.Size2Xs"),
			e.Code("fa.SizeXs"),
			e.Code("fa.SizeSm"),
			e.Code("fa.SizeLg"),
			e.Code("fa.SizeXl"),
			e.Code("fa.Size2Xl"),
			e.Code("fa.Size1x"),
			e.Code("fa.Size2x"),
			e.Code("fa.Size3x"),
			e.Code("fa.Size4x"),
			e.Code("fa.Size5x"),
			e.Code("fa.Size6x"),
			e.Code("fa.Size7x"),
			e.Code("fa.Size8x"),
			e.Code("fa.Size9x"),
			e.Code("fa.Size10x"),
		),
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
		fa.Size2x,
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
		fa.Size2x,
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
		fa.Size2x,
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
		e.P("The icons can be animated using one of the structures implementing ", e.Code("fa.Animation"), " (see below). All these structures accept these options"),
		b.DList(
			e.Code("fa.Delay(time.Duration)"),
			"Add an initial delay",

			e.Code("fa.DirectionNormal"),
			[]any{"Change the animation direction to normal"},

			e.Code("fa.DirectionReverse"),
			[]any{"Change the animation direction to reverse"},

			e.Code("fa.DirectionAlternate"),
			[]any{"Change the animation direction to alternate"},

			e.Code("fa.DirectionAlternateReverse"),
			[]any{"Change the animation direction to alternate reverse"},

			e.Code("fa.Duration(time.Duration)"),
			"Set duration",

			e.Code("fa.IterationCount(float64)"),
			"Set number of iterations",

			e.Code("fa.TimingEase"),
			[]any{"Set animation progress timing to ease"},

			e.Code("fa.TimingLinear"),
			[]any{"Set animation progress timing to linear"},

			e.Code("fa.TimingEaseIn"),
			[]any{"Set animation progress timing to ease-in"},

			e.Code("fa.TimingEaseOut"),
			[]any{"Set animation progress timing to ease-out"},

			e.Code("fa.TimingEaseInOut"),
			[]any{"Set animation progress timing to ease-in-out"},

			e.Code("fa.TimingStepStart"),
			[]any{"Set animation progress timing to step-start"},

			e.Code("fa.TimingStepEnd"),
			[]any{"Set animation progress timing to step-end"},
		),
		e.P("Values are rounded to the 2nd decimal. Please note you cannot use ", e.Code("0"), " as a value, because it is the default value in ", e.Em("Go"), " structures and, as such, is ignored by this library. If you need a zero value for one of the animation attributes, provide a very small float, for instance ", e.Code("0.001"), " which, after rounding, results in ", e.Em("0.00"), "."),
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
		fa.Size2x,
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
		fa.Size2x,
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
		fa.Size2x,
		fa.FA(fa.Solid, "triangle-exclamation", fa.BeatFade()), " ",
		fa.FA(fa.Solid, "square-xmark", fa.BeatFade()), " ",
		fa.FA(fa.Solid, "poo-bolt", fa.BeatFade(fa.MinOpacity(0.1), fa.MaxScale(1.25))), " ",
		fa.FA(fa.Solid, "circle-info", fa.BeatFade(fa.MinOpacity(0.67), fa.MaxScale(1.075))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Bounce"), " animation type bounces an icon up and down. Additionally to the common attributes, it uses the following attributes:"),
		b.DList(
			e.Code("fa.Rebound(float64)"),
			"Set the amount of rebound when landing after the jump",
			e.Code("fa.Height(float64)"),
			"Set the max height to jump when bouncing",
			e.Code("fa.StartScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) when starting to bounce",
			e.Code("fa.StartScaleY(float64)"),
			"Set the icon's vertical distortion (squish) when starting to bounce",
			e.Code("fa.JumpScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) at the top of the jump",
			e.Code("fa.JumpScaleY(float64)"),
			"Set the icon's vertical distortion (squish) at the top of the jump",
			e.Code("fa.LandScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) when landing after the jump",
			e.Code("fa.LandScaleY(float64)"),
			"Set the icon's vertical distortion (squish) when landing after the jump",
		),
	),
	c.Example(
		`fa.FA(fa.Solid, "basketball", fa.Bounce()), " ",
fa.FA(fa.Solid, "volleyball", fa.Bounce()), " ",
fa.FA(fa.Solid, "frog", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1))), " ",
fa.FA(fa.Solid, "envelope", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1), fa.Rebound(0.001)))`,
		fa.Size2x,
		fa.FA(fa.Solid, "basketball", fa.Bounce()), " ",
		fa.FA(fa.Solid, "volleyball", fa.Bounce()), " ",
		fa.FA(fa.Solid, "frog", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1))), " ",
		fa.FA(fa.Solid, "envelope", fa.Bounce(fa.StartScaleX(1), fa.StartScaleY(1), fa.JumpScaleX(1), fa.JumpScaleY(1), fa.LandScaleX(1), fa.LandScaleY(1), fa.Rebound(0.001))),
	),
	b.Content(
		e.P("The ", e.Code("fa.Flip"), " animation type rotates an icon in 3D space. By default, it rotates about the Y axis 180 degrees. Additionally to the common attributes, it uses the ", e.Code("fa.X(float64)"), ", ", e.Code("fa.Y(float64)"), ", ", e.Code("fa.Z(float64)"), " attributes to denote the axis of rotation (between ", e.Code("0"), " and ", e.Code("1"), ") and the ", e.Code("fa.Angle(float64)"), " attribute to set the rotation angle."),
	),
	c.Example(
		`fa.FA(fa.Solid, "compact-disc", fa.Flip()), " ",
fa.FA(fa.Solid, "camera-rotate", fa.Flip()), " ",
fa.FA(fa.Solid, "floppy-disk", fa.Flip()), " ",
fa.FA(fa.Solid, "scroll", fa.Flip(fa.X(1), fa.Y(0.001))), " ",
fa.FA(fa.Solid, "money-check-dollar", fa.Flip(fa.Duration(3*time.Second)))`,
		fa.Size2x,
		fa.FA(fa.Solid, "compact-disc", fa.Flip()), " ",
		fa.FA(fa.Solid, "camera-rotate", fa.Flip()), " ",
		fa.FA(fa.Solid, "floppy-disk", fa.Flip()), " ",
		fa.FA(fa.Solid, "scroll", fa.Flip(fa.X(1), fa.Y(0.001))), " ",
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
		fa.Size2x,
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
		fa.Size2x,
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
	fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		e.P(
			fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		e.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	b.Content(
		e.P("The following modifiers allow customizing the border and pulling effects:"),
		b.DList(
			e.Code("fa.BorderColor"),
			"Set border color",
			e.Code("fa.BorderPadding"),
			"Set padding around icon",
			e.Code("fa.BorderRadius"),
			"Set border radius",
			e.Code("fa.BorderStyle"),
			"Set border style",
			e.Code("fa.BorderWidth"),
			"Set border width",
			e.Code("fa.PullMargin"),
			"Set margin around icon",
		),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft, fa.PullMargin("4em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		e.P(
			fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft, fa.PullMargin("4em")),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`e.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		e.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
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
	fa.Size2x,
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
			fa.Size2x,
			fa.FA(fa.Solid, "square", fa.Stack2x),
			fa.FA(fa.Solid, "terminal", fa.Stack1x, fa.Inverse),
		),
	),
	c.Example(
		`fa.FA(fa.Regular, "circle", fa.Size2x), " ",
fa.Stack(
	fa.FA(fa.Regular, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x),
), " ",
fa.Stack(
	e.Styles{"vertical-align": "top"},
	fa.FA(fa.Solid, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
), " ",
fa.FA(fa.Regular, "circle", fa.Size2x)`,
		fa.FA(fa.Regular, "circle", fa.Size2x), " ",
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
		fa.FA(fa.Regular, "circle", fa.Size2x),
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
			e.Code("fa.ZIndex"),
			"Set z-index of a stacked icon",
			e.Code("fa.InverseColor"),
			"Set color of an inversed stacked icon",
		),
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
		b.DList(
			e.Code("fa.SwapOpacity"),
			"Swap the default opacity of each icon's layers",
			e.Code("fa.PrimaryOpacity"),
			"Set primary layer opacity",
			e.Code("fa.SecondaryOpacity"),
			"Set secondary layer opacity",
			e.Code("fa.PrimaryColor"),
			"Set primary layer color",
			e.Code("fa.SecondaryColor"),
			"Set secondary layer color",
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
