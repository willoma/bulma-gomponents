package docs

import (
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var container = c.NewPage(
	"Container", "Container", "/container",
	"",

	b.Content(
		el.P("The ", el.Code("b.Container"), " constructor creates a container element. Its ", el.Strong("maximum width"), " depends on the applied modifier:"),
	),
	b.Table(
		b.HeadRow(
			"Modifier",
			"Up to 1023px",
			"Desktop (1024px-1215px)",
			"Widescreen (1216px-1407px)",
			"FullHD (1408px and above)",
		),
		b.Row(
			"None",
			b.TCell(b.Light, "Full width"),
			"960px",
			"1152px",
			"1344px",
		),
		b.Row(
			el.Code("b.Widescreen"),
			b.TCell(html.ColSpan("2"), b.Light, "Full width"),
			"1152px",
			"1344px",
		),
		b.Row(
			el.Code("b.FullHD"),
			b.TCell(html.ColSpan("3"), b.Light, "Full width"),
			"1344px",
		),
		b.Row(
			el.Code("b.MaxDesktop"),
			b.TCell(b.Light, "Full width"),
			b.TCell(html.ColSpan("3"), "960px"),
		),
		b.Row(
			el.Code("b.MaxWidescreen"),
			b.TCell(b.Light, "Full width"),
			"960px",
			b.TCell(html.ColSpan("2"), "1152px"),
		),
		b.Row(
			el.Code("b.Fluid"),
			b.TCell(html.ColSpan("4"), "100% with 32px gap on both sides"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/layout/container/",
).Subsection(
	"Default behavior",
	"https://bulma.io/documentation/layout/container/#default-behavior",
	c.HorizontalExample(
		`b.Container(
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("centered"), " on desktop and larger viewports.",
	),
)`,
		b.Container(
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("centered"), " on desktop and larger viewports.",
			),
		),
	),
).Subsection(
	"Widescreen or FullHD only",
	"https://bulma.io/documentation/layout/container/#widescreen-or-fullhd-only",
	c.HorizontalExample(
		`b.Container(
	b.Widescreen,
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$widescreen"), " breakpoint.",
	),
)`,
		b.Container(
			b.Widescreen,
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$widescreen"), " breakpoint.",
			),
		),
	),
	c.HorizontalExample(
		`b.Container(
	b.FullHD,
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$fullhd"), " breakpoint.",
	),
)`,
		b.Container(
			b.FullHD,
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$fullhd"), " breakpoint.",
			),
		),
	),
).Subsection(
	"Desktop and Widescreen maximum widths",
	"https://bulma.io/documentation/layout/container/#desktop-and-widescreen-maximum-widths",
	c.HorizontalExample(
		`b.Container(
	b.MaxDesktop,
	b.Notification(
		b.Primary,
		"This container has a ", el.Code("max-width"), " of ", el.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
	),
)`,
		b.Container(
			b.MaxDesktop,
			b.Notification(
				b.Primary,
				"This container has a ", el.Code("max-width"), " of ", el.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
			),
		),
	),
	c.HorizontalExample(
		`b.Container(
	b.MaxWidescreen,
	b.Notification(
		b.Primary,
		"This container has a ", el.Code("max-width"), " of ", el.Code("$widescreen - $container-offset"), " on fullhd.",
	),
)`,
		b.Container(
			b.MaxWidescreen,
			b.Notification(
				b.Primary,
				"This container has a ", el.Code("max-width"), " of ", el.Code("$widescreen - $container-offset"), " on fullhd.",
			),
		),
	),
).Subsection(
	"Fluid container",
	"https://bulma.io/documentation/layout/container/#fluid-container",
	c.HorizontalExample(
		`b.Container(
	b.Fluid,
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
	),
)`,
		b.Container(
			b.Fluid,
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
			),
		),
	),
)
