package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var container = c.NewPage(
	"Container", "Container", "/container",
	"",

	b.Content(
		e.P("The ", e.Code("b.Container"), " constructor creates a container element. Its ", e.Strong("maximum width"), " depends on the applied modifier:"),
	),
	c.Table(
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
			e.Code("b.Widescreen"),
			b.TCell(e.ColSpan("2"), b.Light, "Full width"),
			"1152px",
			"1344px",
		),
		b.Row(
			e.Code("b.FullHD"),
			b.TCell(e.ColSpan("3"), b.Light, "Full width"),
			"1344px",
		),
		b.Row(
			e.Code("b.MaxDesktop"),
			b.TCell(b.Light, "Full width"),
			b.TCell(e.ColSpan("3"), "960px"),
		),
		b.Row(
			e.Code("b.MaxWidescreen"),
			b.TCell(b.Light, "Full width"),
			"960px",
			b.TCell(e.ColSpan("2"), "1152px"),
		),
		b.Row(
			e.Code("b.Fluid"),
			b.TCell(e.ColSpan("4"), "100% with 32px gap on both sides"),
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
		"This container is ", e.Strong("centered"), " on desktop and larger viewports.",
	),
)`,
		b.Container(
			b.Notification(
				b.Primary,
				"This container is ", e.Strong("centered"), " on desktop and larger viewports.",
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
		"This container is ", e.Strong("fullwidth"), " ", e.Em("until"), " the ", e.Code("$widescreen"), " breakpoint.",
	),
)`,
		b.Container(
			b.Widescreen,
			b.Notification(
				b.Primary,
				"This container is ", e.Strong("fullwidth"), " ", e.Em("until"), " the ", e.Code("$widescreen"), " breakpoint.",
			),
		),
	),
	c.HorizontalExample(
		`b.Container(
	b.FullHD,
	b.Notification(
		b.Primary,
		"This container is ", e.Strong("fullwidth"), " ", e.Em("until"), " the ", e.Code("$fullhd"), " breakpoint.",
	),
)`,
		b.Container(
			b.FullHD,
			b.Notification(
				b.Primary,
				"This container is ", e.Strong("fullwidth"), " ", e.Em("until"), " the ", e.Code("$fullhd"), " breakpoint.",
			),
		),
	),
).Subsection(
	"Tablet, Desktop and Widescreen maximum widths",
	"https://bulma.io/documentation/layout/container/#desktop-and-widescreen-maximum-widths",
	c.HorizontalExample(
		`b.Container(
	b.MaxTablet,
	b.Notification(
		b.Primary,
		"This container has a ", e.Code("max-width"), " of ", e.Code("$tablet - $container-offset"), ".",
	),
)`,
		b.Container(
			b.MaxTablet,
			b.Notification(
				b.Primary,
				"This container has a ", e.Code("max-width"), " of ", e.Code("$tablet - $container-offset"), ".",
			),
		),
	),
	c.HorizontalExample(
		`b.Container(
	b.MaxDesktop,
	b.Notification(
		b.Primary,
		"This container has a ", e.Code("max-width"), " of ", e.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
	),
)`,
		b.Container(
			b.MaxDesktop,
			b.Notification(
				b.Primary,
				"This container has a ", e.Code("max-width"), " of ", e.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
			),
		),
	),
	c.HorizontalExample(
		`b.Container(
	b.MaxWidescreen,
	b.Notification(
		b.Primary,
		"This container has a ", e.Code("max-width"), " of ", e.Code("$widescreen - $container-offset"), " on fullhd.",
	),
)`,
		b.Container(
			b.MaxWidescreen,
			b.Notification(
				b.Primary,
				"This container has a ", e.Code("max-width"), " of ", e.Code("$widescreen - $container-offset"), " on fullhd.",
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
		"This container is ", e.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
	),
)`,
		b.Container(
			b.Fluid,
			b.Notification(
				b.Primary,
				"This container is ", e.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
			),
		),
	),
)
