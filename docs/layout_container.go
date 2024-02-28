package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var container = c.NewPage(
	"Container", "Container", "/container",
	"https://bulma.io/documentation/layout/container/",
).Section(
	"Overview",
	"https://bulma.io/documentation/layout/container/#overview",
	b.Table(
		b.HeadRow(
			"Bulma classes",
			"Bulma-Gomponents component",
		),
		b.Row(
			el.Code(".container"),
			el.Code("b.Container"),
		),
		b.Row(
			el.Code(".container.is-widescreen"),
			el.Code("b.ContainerWidescreen"),
		),
		b.Row(
			el.Code(".container.is-fullhd"),
			el.Code("b.ContainerFullHD"),
		),
		b.Row(
			el.Code(".container.is-max-desktop"),
			el.Code("b.ContainerMaxDesktop"),
		),
		b.Row(
			el.Code(".container.is-max-widescreen"),
			el.Code("b.ContainerMaxWidescreen"),
		),
	),
).Section(
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
).Section(
	"Widescreen or FullHD only",
	"https://bulma.io/documentation/layout/container/#widescreen-or-fullhd-only",
	c.HorizontalExample(
		`b.ContainerWidescreen(
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$widescreen"), " breakpoint.",
	),
)`,
		b.ContainerWidescreen(
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$widescreen"), " breakpoint.",
			),
		),
	),
	c.HorizontalExample(
		`b.ContainerFullHD(
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$fullhd"), " breakpoint.",
	),
)`,
		b.ContainerFullHD(
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fullwidth"), " ", el.Em("until"), " the ", el.Code("$fullhd"), " breakpoint.",
			),
		),
	),
).Section(
	"Desktop and Widescreen maximum widths",
	"https://bulma.io/documentation/layout/container/#desktop-and-widescreen-maximum-widths",
	c.HorizontalExample(
		`b.ContainerMaxDesktop(
	b.Notification(
		b.Primary,
		"This container has a ", el.Code("max-width"), " of ", el.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
	),
)`,
		b.ContainerMaxDesktop(
			b.Notification(
				b.Primary,
				"This container has a ", el.Code("max-width"), " of ", el.Code("$desktop - $container-offset"), " on widescreen and fullhd.",
			),
		),
	),
	c.HorizontalExample(
		`b.ContainerMaxWidescreen(
	b.Notification(
		b.Primary,
		"This container has a ", el.Code("max-width"), " of ", el.Code("$widescreen - $container-offset"), " on fullhd.",
	),
)`,
		b.ContainerMaxWidescreen(
			b.Notification(
				b.Primary,
				"This container has a ", el.Code("max-width"), " of ", el.Code("$widescreen - $container-offset"), " on fullhd.",
			),
		),
	),
).Section(
	"Fluid container",
	"https://bulma.io/documentation/layout/container/#fluid-container",
	c.HorizontalExample(
		`b.ContainerFluid(
	b.Notification(
		b.Primary,
		"This container is ", el.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
	),
)`,
		b.ContainerFluid(
			b.Notification(
				b.Primary,
				"This container is ", el.Strong("fluid"), ": it will have a 32px gap on either side, on any viewport size.",
			),
		),
	),
)
