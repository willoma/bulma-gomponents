package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var icon = c.NewPage(
	"Icon", "Icon", "/icon",
	"https://bulma.io/documentation/elements/icon/",

	c.Example(
		`b.Icon(el.I(b.Class("fas"), b.Class("fa-home")))`,
		b.Icon(el.I(b.Class("fas"), b.Class("fa-home"))),
	),
	b.Content(el.P("The ", el.Code("bulma/fa"), " package includes helper functions and values for using ", el.Em("Font Awesome"), " icons with ", el.Em("Bulma-Gomponents"), ":")),
	c.Example(
		`b.Icon(fa.FA(fa.Solid, "home")),
fa.Icon(fa.Solid, "home")`,
		b.Icon(fa.FA(fa.Solid, "home")),
		fa.Icon(fa.Solid, "home"),
	),
).Section(
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
			"Some information is missing from your ", b.AHref("#", "profile"), "details.",
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
			"Some information is missing from your ", b.AHref("#", "profile"), "details.",
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
).Section(
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
b.IconText(
	b.Success,
	fa.Icon(fa.Solid, "check-square"),
	"Success",
),
b.IconText(
	b.Warning,
	fa.Icon(fa.Solid, "exclamation-triangle"),
	"Warning",
),
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
		b.IconText(
			b.Success,
			fa.Icon(fa.Solid, "check-square"),
			"Success",
		),
		b.IconText(
			b.Warning,
			fa.Icon(fa.Solid, "exclamation-triangle"),
			"Warning",
		),
		b.IconText(
			b.Danger,
			fa.Icon(fa.Solid, "ban"),
			"Danger",
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/elements/icon/#sizes",
	b.Content(
		el.P("Use a size modifier (", el.Code("b.Small"), ", etc) as an argument to ", el.Code("b.Icon"), " or ", el.Code("fa.Icon"), "."),
	),
).Section(
	"Font Awesome variations",
	"https://bulma.io/documentation/elements/icon/#font-awesome-variations",
	b.Content(
		b.DList(
			"Fixed width", el.Code("fa.FixedWidth"),
			"Bordered", el.Code("fa.Border"),
			"Animated", el.Code("fa.Pulse"),
		),
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
).Section(
	"Material Design Icons",
	"https://bulma.io/documentation/elements/icon/#material-design-icons",
	b.Content("Nothing specific here."),
).Section(
	"Ionicons",
	"https://bulma.io/documentation/elements/icon/#ionicons",
	b.Content("Nothing specific here."),
)
