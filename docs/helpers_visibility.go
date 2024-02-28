package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var visibility = c.NewPage(
	"Visibility", "Responsive helpers", "/visibility",
	"https://bulma.io/documentation/helpers/visibility-helpers/",
).Section(
	"Show",
	"https://bulma.io/documentation/helpers/visibility-helpers/#show",
	b.Content(
		el.P("You can use one of the following display classes:"),
		b.UList(
			el.Code("b.VisibilityBlock"),
			el.Code("b.Flex"),
			el.Code("b.Inline"),
			el.Code("b.InlineBlock"),
			el.Code("b.InlineFlex"),
		),
		el.P("Modifiers to display on specific breakpoints:"),
		b.UList(
			el.Code(".Mobile()"),
			el.Code(".TabletOnly()"),
			el.Code(".DesktopOnly()"),
			el.Code(".WidescreenOnly()"),
		),
		el.P("Modifiers to display up to or from a specific breakpoint:"),
		b.UList(
			el.Code(".Touch()"),
			el.Code(".Tablet()"),
			el.Code(".Desktop()"),
			el.Code(".Widescreen()"),
			el.Code(".FullHD()"),
		),
	),
).Section(
	"Hide",
	"https://bulma.io/documentation/helpers/visibility-helpers/#hide",
	b.Content(
		el.P("Modifiers to hide on specific breakpoints:"),
		b.UList(
			el.Code("b.Hidden.Mobile()"),
			el.Code("b.Hidden.TabletOnly()"),
			el.Code("b.Hidden.DesktopOnly()"),
			el.Code("b.Hidden.WidescreenOnly()"),
		),
		el.P("Modifiers to hide up to or from a specific breakpoint:"),
		b.UList(
			el.Code("b.Hidden.Touch()"),
			el.Code("b.Hidden.Tablet()"),
			el.Code("b.Hidden.Desktop()"),
			el.Code("b.Hidden.Widescreen()"),
			el.Code("b.Hidden.FullHD()"),
		),
	),
).Section(
	"Other visibility helpers",
	"https://bulma.io/documentation/helpers/visibility-helpers/#other-visibility-helpers",
	b.Table(
		b.HeadRow("Class", "Result"),
		b.Row(el.Code("b.Invisible"), b.Cell("Adds visibility ", el.Strong("hidden"))),
		b.Row(el.Code("b.SrOnly"), b.Cell("Hide elements ", el.Strong("visually"), " but keep the element available to be announced by a ", el.Strong("screen reader"))),
	),
)
