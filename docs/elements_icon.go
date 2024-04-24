package docs

import (
	"time"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var icon = c.NewPage(
	"Icon", "Icon", "/icon",
	"",
	b.Content(
		el.P(
			"The ", el.Code("b.Icon"), " constructor returns a container for an ", el.Strong("icon font"), ". The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.White"),
			"Set icon color to white",

			el.Code("b.Black"),
			"Set icon color to black",

			el.Code("b.Light"),
			"Set icon color to light",

			el.Code("b.Dark"),
			"Set icon color to dark",

			el.Code("b.Primary"),
			"Set icon color to primary",

			el.Code("b.Link"),
			"Set icon color to link",

			el.Code("b.Info"),
			"Set icon color to info",

			el.Code("b.Success"),
			"Set icon color to success",

			el.Code("b.Warning"),
			"Set icon color to warning",

			el.Code("b.Danger"),
			"Set icon color to danger",

			el.Code("b.BlackBis"),
			"Set icon color to black bis",

			el.Code("b.BlackTer"),
			"Set icon color to black ter",

			el.Code("b.GreyDarker"),
			"Set icon color to grey darker",

			el.Code("b.GreyDark"),
			"Set icon color to grey dark",

			el.Code("b.Grey"),
			"Set icon color to grey",

			el.Code("b.GreyLight"),
			"Set icon color to grey light",

			el.Code("b.GreyLigher"),
			"Set icon color to grey lighter",

			el.Code("b.WhiteTer"),
			"Set icon color to white ter",

			el.Code("b.WhiteBis"),
			"Set icon color to white bis",

			el.Code("b.PrimaryLight"),
			"Set icon color to primary light",

			el.Code("b.LinkLight"),
			"Set icon color to link light",

			el.Code("b.InfoLight"),
			"Set icon color to info light",

			el.Code("b.SuccessLight"),
			"Set icon color to success light",

			el.Code("b.WarningLight"),
			"Set icon color to warning light",

			el.Code("b.DangerLight"),
			"Set icon color to danger light",

			el.Code("b.PrimaryDark"),
			"Set icon color to primary dark",

			el.Code("b.LinkDark"),
			"Set icon color to link dark",

			el.Code("b.InfoDark"),
			"Set icon color to info dark",

			el.Code("b.SuccessDark"),
			"Set icon color to success dark",

			el.Code("b.WarningDark"),
			"Set icon color to warning dark",

			el.Code("b.DangerDark"),
			"Set icon color to danger dark",

			el.Code("b.Small"),
			"Set icon size to small",

			el.Code("b.Medium"),
			"Set icon size to medium",

			el.Code("b.Large"),
			"Set icon size to large",
		),

		el.P(
			"The ", el.Code("b.IconText"), " constructor creates an icon+text span container and embeds all its non-icon children into spans. The ", el.Code("b.FlexIconText"), " constructor creates a flex icon+text span container and embeds all its non-icon children into spans. The following children have a special meaning:",
		),
		b.DList(el.Code("b.White"),
			"Set icon color to white",

			el.Code("b.Black"),
			"Set icon color to black",

			el.Code("b.Light"),
			"Set icon color to light",

			el.Code("b.Dark"),
			"Set icon color to dark",

			el.Code("b.Primary"),
			"Set icon color to primary",

			el.Code("b.Link"),
			"Set icon color to link",

			el.Code("b.Info"),
			"Set icon color to info",

			el.Code("b.Success"),
			"Set icon color to success",

			el.Code("b.Warning"),
			"Set icon color to warning",

			el.Code("b.Danger"),
			"Set icon color to danger",

			el.Code("b.BlackBis"),
			"Set icon color to black bis",

			el.Code("b.BlackTer"),
			"Set icon color to black ter",

			el.Code("b.GreyDarker"),
			"Set icon color to grey darker",

			el.Code("b.GreyDark"),
			"Set icon color to grey dark",

			el.Code("b.Grey"),
			"Set icon color to grey",

			el.Code("b.GreyLight"),
			"Set icon color to grey light",

			el.Code("b.GreyLigher"),
			"Set icon color to grey lighter",

			el.Code("b.WhiteTer"),
			"Set icon color to white ter",

			el.Code("b.WhiteBis"),
			"Set icon color to white bis",

			el.Code("b.PrimaryLight"),
			"Set icon color to primary light",

			el.Code("b.LinkLight"),
			"Set icon color to link light",

			el.Code("b.InfoLight"),
			"Set icon color to info light",

			el.Code("b.SuccessLight"),
			"Set icon color to success light",

			el.Code("b.WarningLight"),
			"Set icon color to warning light",

			el.Code("b.DangerLight"),
			"Set icon color to danger light",

			el.Code("b.PrimaryDark"),
			"Set icon color to primary dark",

			el.Code("b.LinkDark"),
			"Set icon color to link dark",

			el.Code("b.InfoDark"),
			"Set icon color to info dark",

			el.Code("b.SuccessDark"),
			"Set icon color to success dark",

			el.Code("b.WarningDark"),
			"Set icon color to warning dark",

			el.Code("b.DangerDark"),
			"Set icon color to danger dark",
		),
	),
).Section(
	"Font Awesome", "",

	b.Content(
		el.P("The", el.Code("github.com/willoma/bulma-gomponents/fa"), " package provides helpers for ", el.Em("Font Awesome"), " icons."),
		el.P(
			"The ", el.Code("fa.Icon"), " constructor creates a ", el.Em("Font Awesome"), " icon, embedded in an icon container, dealing with appropriately applying children to the icon element or to the span container. It accepts the same values as ", el.Code("b.Icon"), ", as well as the ", el.Em("Font Awesome"), "-specific values described below. The following children have a special meaning:",
		),
		b.DList(
			el.Code("fa.OnFA(...)"),
			[]any{"Force childen to be applied to the Font Awesome ", el.Code("<i>"), " element"},

			el.Code("fa.OnIcon(...)"),
			[]any{"Force childen to be applied to the icon span"},
		),
		el.P("Other children are added to the icon span."),
		el.P(
			"The ", el.Code("fa.FA"), " constructor creates a ", el.Em("Font Awesome"), " icon. It accepts the ", el.Em("Font Awesome"), "-specific values described below.",
		),
	),
).Subsection(
	"First arguments", "Font Awesome::https://docs.fontawesome.com/web/add-icons/how-to#families--styles",

	b.Content(
		el.P("The ", el.Code("fa.Icon"), " function requires at least two arguments:"),
		el.Ol(
			el.Li(
				"The font style, one of ", el.Code("fa.Brand"), ", ", el.Code("fa.Duotone"), ", ", el.Code("fa.Light"), ", ", el.Code("fa.SharpLight"), ", ", el.Code("fa.Regular"), ", ", el.Code("fa.SharpRegular"), ", ", el.Code("fa.Solid"), ", ", el.Code("fa.SharpSolid"), ", ", el.Code("fa.Thin"), " or ", el.Code("fa.SharpThin"),
			),
			el.Li(
				"The icon name, ", el.Em("without the ", el.Code("fa-"), " prefix"), " (for example, ", el.Code(`"home"`), " or ", el.Code(`"spinner"`), ")",
			),
		),
	),
).Subsection(
	"Styling basics", "Font Awesome::https://docs.fontawesome.com/web/style/basics",

	b.Content(
		el.P("Arguments provided to ", el.Code("fa.Icon"), " are automatically applied to the ", el.Code("<span>"), " or ", el.Code("<i>"), " element, depending on its nature. If you need to explicitly apply a child to the ", el.Code("<i>"), " element, use ", el.Code("fa.FA"), " and include it manually as a child of ", el.Code("b.Icon"), "."),
	),
	c.Example(
		`el.Span(
	b.Style("font-size", "3em", "color", "Tomato"),
	fa.FA(
		fa.Solid, "camera",
	),
),
" ",
el.Span(
	b.Style("font-size", "48px", "color", "Dodgerblue"),
	fa.FA(
		fa.Solid, "camera",
	),
),
" ",
el.Span(
	b.Style("font-size", "3rem"),
	el.Span(
		b.Style("color", "Mediumslateblue"),
		fa.FA(
			fa.Solid, "camera",
		),
	),
)`,
		el.Span(
			b.Style("font-size", "3em", "color", "Tomato"),
			fa.FA(
				fa.Solid, "camera",
			),
		),
		" ",
		el.Span(
			b.Style("font-size", "48px", "color", "Dodgerblue"),
			fa.FA(
				fa.Solid, "camera",
			),
		),
		" ",
		el.Span(
			b.Style("font-size", "3rem"),
			el.Span(
				b.Style("color", "Mediumslateblue"),
				fa.FA(
					fa.Solid, "camera",
				),
			),
		),
	),
).Subsection(
	"Sizing icons", "Font Awesome::https://docs.fontawesome.com/web/style/size",

	b.Content(
		el.P(
			"The following size modifiers are available:",
		),
		b.UList(
			b.MarginVertical(b.Spacing0),
			el.Code("fa.Size2Xs"),
			el.Code("fa.SizeXs"),
			el.Code("fa.SizeSm"),
			el.Code("fa.SizeLg"),
			el.Code("fa.SizeXl"),
			el.Code("fa.Size2Xl"),
			el.Code("fa.Size1x"),
			el.Code("fa.Size2x"),
			el.Code("fa.Size3x"),
			el.Code("fa.Size4x"),
			el.Code("fa.Size5x"),
			el.Code("fa.Size6x"),
			el.Code("fa.Size7x"),
			el.Code("fa.Size8x"),
			el.Code("fa.Size9x"),
			el.Code("fa.Size10x"),
		),
	),
).Subsection(
	"Fixed width icons", "Font Awesome::https://docs.fontawesome.com/web/style/fixed-width",

	b.Content(
		el.P("In order to make all your icons the same width so they can easily vertically align, like in a list or navigation menu, use the", el.Code("fa.FixedWidth"), " modifier."),
		el.P("However, please note that ", el.Strong(el.Em("Bulma"), " icons are already square"), ", with a fixed width."),
	),
	c.Example(
		`el.Div(fa.FA(fa.Solid, "skating", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Skating"),
el.Div(fa.FA(fa.Solid, "skiing", fa.FixedWidth, b.Style("background", "SkyBlue")), " Skiing"),
el.Div(fa.FA(fa.Solid, "skiing-nordic", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Nordic Skiing"),
el.Div(fa.FA(fa.Solid, "snowboarding", fa.FixedWidth, b.Style("background", "SkyBlue")), " Snowboarding"),
el.Div(fa.FA(fa.Solid, "snowplow", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Snowplow"),`,
		b.Style("font-size", "1.5rem"),
		el.Div(fa.FA(fa.Solid, "skating", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Skating"),
		el.Div(fa.FA(fa.Solid, "skiing", fa.FixedWidth, b.Style("background", "SkyBlue")), " Skiing"),
		el.Div(fa.FA(fa.Solid, "skiing-nordic", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Nordic Skiing"),
		el.Div(fa.FA(fa.Solid, "snowboarding", fa.FixedWidth, b.Style("background", "SkyBlue")), " Snowboarding"),
		el.Div(fa.FA(fa.Solid, "snowplow", fa.FixedWidth, b.Style("background", "DodgerBlue")), " Snowplow"),
	),
).Subsection(
	"Icons in a list", "Font Awesome::https://docs.fontawesome.com/web/style/lists",

	b.Content(
		el.P("Use ", el.Em("fa.UList"), " and ", el.Em("fa.Li"), " to replace default bullets in unordered lists. You can also keep the semantics of an ordered list behind the scenes but use icon bullets visually with ", el.Em("fa.OList"), "."),
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
		el.P(
			"To arbitrarily rotate and flip icons, use one of the following modifiers on ", el.Code("fa.Icon"), ", ", el.Code("fa.FA"), " or ", el.Code("fa.Li"), ":",
		),
		b.UList(
			el.Code("fa.Rotate90"),
			el.Code("fa.Rotate180"),
			el.Code("fa.Rotate270"),
			el.Code("fa.FlipHorizontal"),
			el.Code("fa.FlipVertical"),
			el.Code("fa.FlipBoth"),
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
		el.P("In order to apply a custom rotation, use a ", el.Code("fa.Rotate"), " value."),
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
		el.P("When combining rotation and flipping, the ", el.Code("<i>"), " is automatically embedded in a ", el.Code("<span>"), " in order to apply all transformations."),
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
		el.P("The icons can be animated using one of the structures implementing ", el.Code("fa.Animation"), " (see below). All these structures accept these options"),
		b.DList(
			el.Code("fa.Delay(time.Duration)"),
			"Add an initial delay",

			el.Code("fa.DirectionNormal"),
			[]any{"Change the animation direction to normal"},

			el.Code("fa.DirectionReverse"),
			[]any{"Change the animation direction to reverse"},

			el.Code("fa.DirectionAlternate"),
			[]any{"Change the animation direction to alternate"},

			el.Code("fa.DirectionAlternateReverse"),
			[]any{"Change the animation direction to alternate reverse"},

			el.Code("fa.Duration(time.Duration)"),
			"Set duration",

			el.Code("fa.IterationCount(float64)"),
			"Set number of iterations",

			el.Code("fa.TimingEase"),
			[]any{"Set animation progress timing to ease"},

			el.Code("fa.TimingLinear"),
			[]any{"Set animation progress timing to linear"},

			el.Code("fa.TimingEaseIn"),
			[]any{"Set animation progress timing to ease-in"},

			el.Code("fa.TimingEaseOut"),
			[]any{"Set animation progress timing to ease-out"},

			el.Code("fa.TimingEaseInOut"),
			[]any{"Set animation progress timing to ease-in-out"},

			el.Code("fa.TimingStepStart"),
			[]any{"Set animation progress timing to step-start"},

			el.Code("fa.TimingStepEnd"),
			[]any{"Set animation progress timing to step-end"},
		),
		el.P("Values are rounded to the 2nd decimal. Please note you cannot use ", el.Code("0"), " as a value, because it is the default value in ", el.Em("Go"), " structures and, as such, is ignored by this library. If you need a zero value for one of the animation attributes, provide a very small float, for instance ", el.Code("0.001"), " which, after rounding, results in ", el.Em("0.00"), "."),
	),
	b.Content(
		el.P("The ", el.Code("fa.Beat"), " animation type scales an icon up or down. Additionally to the common attributes, it uses the ", el.Code("fa.MaxScale(float64)"), " attribute to set the maximum value the icon will scale."),
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
		el.P("The ", el.Code("fa.Fade"), " animation type fades an icon in and out. Additionally to the common attributes, it uses the ", el.Code("fa.MinOpacity(float64)"), " attribute to set the lowest opacity the icon will fade to and from."),
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
		el.P("The ", el.Code("fa.BeatFade"), " animation type scales and pulses an icon in and out. Additionally to the common attributes, it uses the ", el.Code("fa.MaxScale(float64)"), "and ", el.Code("fa.MinOpacity(float64)"), " attributes."),
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
		el.P("The ", el.Code("fa.Bounce"), " animation type bounces an icon up and down. Additionally to the common attributes, it uses the following attributes:"),
		b.DList(
			el.Code("fa.Rebound(float64)"),
			"Set the amount of rebound when landing after the jump",
			el.Code("fa.Height(float64)"),
			"Set the max height to jump when bouncing",
			el.Code("fa.StartScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) when starting to bounce",
			el.Code("fa.StartScaleY(float64)"),
			"Set the icon's vertical distortion (squish) when starting to bounce",
			el.Code("fa.JumpScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) at the top of the jump",
			el.Code("fa.JumpScaleY(float64)"),
			"Set the icon's vertical distortion (squish) at the top of the jump",
			el.Code("fa.LandScaleX(float64)"),
			"Set the icon's horizontal distortion (squish) when landing after the jump",
			el.Code("fa.LandScaleY(float64)"),
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
		el.P("The ", el.Code("fa.Flip"), " animation type rotates an icon in 3D space. By default, it rotates about the Y axis 180 degrees. Additionally to the common attributes, it uses the ", el.Code("fa.X(float64)"), ", ", el.Code("fa.Y(float64)"), ", ", el.Code("fa.Z(float64)"), " attributes to denote the axis of rotation (between ", el.Code("0"), " and ", el.Code("1"), ") and the ", el.Code("fa.Angle(float64)"), " attribute to set the rotation angle."),
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
		el.P("The ", el.Code("fa.Shake"), " animation type shakes back and forth"),
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
		el.P("The ", el.Code("fa.Spin"), " animation type make the icon spin 360° clockwise. Additionally to the common attributes, it uses the ", el.Code("fa.Pulse"), " attribute to make the rotation in 8 steps and the ", el.Code("fa.Reverse"), " attribute to make the icon spin counter-clockwise."),
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
		el.P("The ", el.Code("fa.Border"), "modifier creates a border with radius and padding around an icon. The ", el.Code("fa.PullLeft"), " modifier pulls the icon by floating it left and applying a right margin. The ", el.Code("fa.PullRight"), " modifier pulls the icon by floating it right and applying a left margin."),
	),
	c.Example(
		`el.P(
	fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		el.P(
			fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`el.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
),`,
		el.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	b.Content(
		el.P("The following modifiers allow customizing the border and pulling effects:"),
		b.DList(
			el.Code("fa.BorderColor"),
			"Set border color",
			el.Code("fa.BorderPadding"),
			"Set padding around icon",
			el.Code("fa.BorderRadius"),
			"Set border radius",
			el.Code("fa.BorderStyle"),
			"Set border style",
			el.Code("fa.BorderWidth"),
			"Set border width",
			el.Code("fa.PullMargin"),
			"Set margin around icon",
		),
	),
	c.Example(
		`el.P(
	fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft, fa.PullMargin("4em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		el.P(
			fa.FA(fa.Solid, "quote-left", fa.Size2x, fa.PullLeft, fa.PullMargin("4em")),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
	c.Example(
		`el.P(
	fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
	"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
)`,
		el.P(
			fa.FA(fa.Solid, "arrow-right", fa.Size2x, fa.PullRight, fa.Border, fa.BorderColor("inherit"), fa.BorderPadding("0.5em"), fa.BorderRadius("100%"), fa.BorderStyle("dotted"), fa.BorderWidth("0.5em")),
			"Gatsby believed in the green light, the orgastic future that year by year recedes before us. It eluded us then, but that’s no matter — tomorrow we will run faster, stretch our arms further... And one fine morning — So we beat on, boats against the current, borne back ceaselessly into the past.",
		),
	),
).Subsection(
	"Stacking icons", "Font Awesome::https://docs.fontawesome.com/web/style/stack",

	b.Content(
		el.P("To stack multiple icons, use the ", el.Code("fa.Stack"), " component builder. Then add the ", el.Code("fa.Stack1x"), " modifier to the regularly sized icon and add the ", el.Code("fa.Stack1x"), " modifier to the larger icon. ", el.Code("fa.Inverse"), " can be added to the icon with ", el.Code("fa.Stack1x"), " to help with a knock-out looking effect."),
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
	fa.FA(fa.Solid, "ban", fa.Stack2x, b.Style("color", "Tomato")),
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
			fa.FA(fa.Solid, "ban", fa.Stack2x, b.Style("color", "Tomato")),
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
	b.Style("vertical-align", "top"),
	fa.FA(fa.Solid, "circle", fa.Stack2x),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
), " ",
fa.FA(fa.Regular, "circle", fa.Size2x)`,
		fa.FA(fa.Regular, "circle", fa.Size2x), " ",
		fa.Stack(
			b.Style("vertical-align", "top"),
			fa.FA(fa.Regular, "circle", fa.Stack2x),
			fa.FA(fa.Solid, "flag", fa.Stack1x),
		), " ",
		fa.Stack(
			b.Style("vertical-align", "top"),
			fa.FA(fa.Solid, "circle", fa.Stack2x),
			fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse),
		), " ",
		fa.FA(fa.Regular, "circle", fa.Size2x),
	),
	c.Example(
		`fa.FA(fa.Regular, "circle", b.Style("vertical-align", "middle")), " ",
fa.Stack(
	b.Style("font-size", "0.5em"),
	fa.FA(fa.Regular, "circle", fa.Stack2x, b.Style("vertical-align", "middle")),
	fa.FA(fa.Solid, "flag", fa.Stack1x, b.Style("vertical-align", "middle")),
), " ",
fa.Stack(
	b.Style("font-size", "0.5em"),
	fa.FA(fa.Solid, "circle", fa.Stack2x, b.Style("vertical-align", "middle")),
	fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse, b.Style("vertical-align", "middle")),
), " ",
fa.FA(fa.Regular, "circle", b.Style("vertical-align", "middle"))`,
		fa.FA(fa.Regular, "circle", b.Style("vertical-align", "middle")), " ",
		fa.Stack(
			b.Style("font-size", "0.5em"),
			fa.FA(fa.Regular, "circle", fa.Stack2x, b.Style("vertical-align", "middle")),
			fa.FA(fa.Solid, "flag", fa.Stack1x, b.Style("vertical-align", "middle")),
		), " ",
		fa.Stack(
			b.Style("font-size", "0.5em"),
			fa.FA(fa.Solid, "circle", fa.Stack2x, b.Style("vertical-align", "middle")),
			fa.FA(fa.Solid, "flag", fa.Stack1x, fa.Inverse, b.Style("vertical-align", "middle")),
		), " ",
		fa.FA(fa.Regular, "circle", b.Style("vertical-align", "middle")),
	),
	b.Content(
		el.P("The following modifiers allow customizing stacked icons:"),
		b.DList(
			el.Code("fa.ZIndex"),
			"Set z-index of a stacked icon",
			el.Code("fa.InverseColor"),
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
		el.P("The following modifiers allow customizing duotone icons:"),
		b.DList(
			el.Code("fa.SwapOpacity"),
			"Swap the default opacity of each icon's layers",
			el.Code("fa.PrimaryOpacity"),
			"Set primary layer opacity",
			el.Code("fa.SecondaryOpacity"),
			"Set secondary layer opacity",
			el.Code("fa.PrimaryColor"),
			"Set primary layer color",
			el.Code("fa.SecondaryColor"),
			"Set secondary layer color",
		),
	),
	b.Message(
		b.Danger,
		"Please note these icons require pro plan, which ", el.Em("Bulma-Gomponents"), " do not have, therefore no example may be shown.",
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/icon/",

	c.Example(
		`b.Icon(el.I(b.Class("fas"), b.Class("fa-home"))),
b.Icon(fa.FA(fa.Solid, "home")),
fa.Icon(fa.Solid, "home")`,
		b.Icon(el.I(b.Class("fas"), b.Class("fa-home"))),
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
	el.P(
		"An invitation to ", b.IconText(fa.Icon(fa.Solid, "utensils"), "dinner"), " was soon afterwards dispatched; and already had Mrs. Bennet planned the courses that were to do credit to her housekeeping, when an answer arrived which deferred it all. Mr. Bingley was obliged to be in ", b.IconText(fa.Icon(fa.Solid, "city"), "town"), " the following day, and, consequently, unable to accept the honour of their ", b.IconText(fa.Icon(fa.Solid, "envelope-open-text"), "invitation"), ", etc.",
	),
	el.P(
		"Mrs. Bennet was quite disconcerted. She could not imagine what business he could have in town so soon after his ", b.IconText(fa.Icon(fa.Solid, "flag-checkered"), "arrival"), " in Hertfordshire; and she began to fear that he might be always ", b.IconText(fa.Icon(fa.Solid, "plane-departure"), "flying"), " about from one place to another, and never settled at Netherfield as he ought to be.",
	),
)`,
		b.Content(
			el.P(
				"An invitation to ", b.IconText(fa.Icon(fa.Solid, "utensils"), "dinner"), " was soon afterwards dispatched; and already had Mrs. Bennet planned the courses that were to do credit to her housekeeping, when an answer arrived which deferred it all. Mr. Bingley was obliged to be in ", b.IconText(fa.Icon(fa.Solid, "city"), "town"), " the following day, and, consequently, unable to accept the honour of their ", b.IconText(fa.Icon(fa.Solid, "envelope-open-text"), "invitation"), ", etc.",
			),
			el.P(
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
			"Your package will be delivered on ", el.Strong("Tuesday at 08:00"), ".",
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
			"Some information is missing from your ", b.AHref("#", "profile"), " details.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "ban", b.Danger),
			"Danger",
		),
		b.Block(
			html.P,
			"There was an error in your submission. ", b.AHref("#", "Please try again"), ".",
		)`,
		b.FlexIconText(
			fa.Icon(fa.Solid, "info-circle", b.Info),
			"Information",
		),
		b.Block(
			html.P,
			"Your package will be delivered on ", el.Strong("Tuesday at 08:00"), ".",
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
			"Some information is missing from your ", b.AHref("#", "profile"), " details.",
		),
		b.FlexIconText(
			fa.Icon(fa.Solid, "ban", b.Danger),
			"Danger",
		),
		b.Block(
			html.P,
			"There was an error in your submission. ", b.AHref("#", "Please try again"), ".",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/icon/#colors",
	c.Example(
		`b.Icon(b.Info, el.I(b.Class("fas"), b.Class("fa-info-circle"))),
b.Icon(b.Success, el.I(b.Class("fas"), b.Class("fa-check-square"))),
b.Icon(b.Warning, el.I(b.Class("fas"), b.Class("fa-exclamation-triangle"))),
b.Icon(b.Danger, el.I(b.Class("fas"), b.Class("fa-ban")))`,
		b.Icon(b.Info, el.I(b.Class("fas"), b.Class("fa-info-circle"))),
		b.Icon(b.Success, el.I(b.Class("fas"), b.Class("fa-check-square"))),
		b.Icon(b.Warning, el.I(b.Class("fas"), b.Class("fa-exclamation-triangle"))),
		b.Icon(b.Danger, el.I(b.Class("fas"), b.Class("fa-ban"))),
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
		el.P("Use a size modifier (", el.Code("b.Small"), ", etc) as an argument to ", el.Code("b.Icon"), " or ", el.Code("fa.Icon"), "."),
	),
).Subsection(
	"Font Awesome variations",
	"https://bulma.io/documentation/elements/icon/#font-awesome-variations",

	c.Example(
		`fa.Icon(fa.Solid, "home", fa.FixedWidth, b.BackgroundWarningLight),
el.Br(),
fa.Icon(fa.Solid, "home", fa.Border),
el.Br(),
fa.Icon(fa.Solid, "spinner", fa.Spin(fa.Pulse))`,
		fa.Icon(fa.Solid, "home", fa.FixedWidth, b.BackgroundWarningLight),
		el.Br(),
		fa.Icon(fa.Solid, "home", fa.Border),
		el.Br(),
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
