package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var visibility = c.NewPage(
	"Visibility", "Responsive helpers", "/visibility",
	"https://bulma.io/documentation/helpers/visibility-helpers/",
).Section(
	"Show",
	"https://bulma.io/documentation/helpers/visibility-helpers/#show",
	b.Content(
		e.P("You can use one of the following display classes:"),
		b.UList(
			e.Code("b.VisibilityBlock"),
			e.Code("b.Flex"),
			e.Code("b.Inline"),
			e.Code("b.InlineBlock"),
			e.Code("b.InlineFlex"),
		),
		e.P("Modifiers to display on specific breakpoints:"),
		b.UList(
			e.Code(".Mobile()"),
			e.Code(".TabletOnly()"),
			e.Code(".DesktopOnly()"),
			e.Code(".WidescreenOnly()"),
		),
		e.P("Modifiers to display up to or from a specific breakpoint:"),
		b.UList(
			e.Code(".Touch()"),
			e.Code(".Tablet()"),
			e.Code(".Desktop()"),
			e.Code(".Widescreen()"),
			e.Code(".FullHD()"),
		),
	),
).Section(
	"Hide",
	"https://bulma.io/documentation/helpers/visibility-helpers/#hide",
	b.Content(
		e.P("Modifiers to hide on specific breakpoints:"),
		b.UList(
			e.Code("b.Hidden.Mobile()"),
			e.Code("b.Hidden.TabletOnly()"),
			e.Code("b.Hidden.DesktopOnly()"),
			e.Code("b.Hidden.WidescreenOnly()"),
		),
		e.P("Modifiers to hide up to or from a specific breakpoint:"),
		b.UList(
			e.Code("b.Hidden.Touch()"),
			e.Code("b.Hidden.Tablet()"),
			e.Code("b.Hidden.Desktop()"),
			e.Code("b.Hidden.Widescreen()"),
			e.Code("b.Hidden.FullHD()"),
		),
	),
).Section(
	"Other visibility helpers",
	"https://bulma.io/documentation/helpers/visibility-helpers/#other-visibility-helpers",
	b.Table(
		b.HeadRow("Modifier", "Result"),
		b.Row(e.Code("b.Invisible"), b.Cell("Adds visibility ", e.Strong("hidden"))),
		b.Row(e.Code("b.SrOnly"), b.Cell("Hide elements ", e.Strong("visually"), " but keep the element available to be announced by a ", e.Strong("screen reader"))),
	),
)
