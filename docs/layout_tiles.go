package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var tiles = c.NewPage(
	"Tiles", "Tiles powered by Flexbox", "/tiles",
	"https://bulma.io/documentation/layout/tiles/",
	b.Content(el.P(
		"To build intricate 2-dimensional layouts, you only need a ", el.Strong("single element"), ": the ", el.Code("b.Tile"), ".",
	)),
).Section(
	"Example",
	"https://bulma.io/documentation/layout/tiles/#example",
	c.HorizontalExample(
		`b.Tile(
	b.VTile(
		b.Size8,
		b.Tile(
			b.VTile(
				b.Notification(
					html.Article, b.Warning,
					b.Title(html.P, "Vertical..."),
					b.Subtitle(html.P, "Top tile"),
				),
				b.Notification(
					html.Article, b.Warning,
					b.Title(html.P, "...tiles"),
					b.Subtitle(html.P, "Bottom tile"),
				),
			),
			b.Tile(
				b.Notification(
					html.Article, b.Info,
					b.Title(html.P, "Middle tile"),
					b.Subtitle(html.P, "With an image"),
					b.ImageImg(
						"https://bulma.io/images/placeholders/640x480.png",
						b.Img4By3,
					),
				),
			),
		),

		b.Tile(
			b.Notification(
				html.Article, b.Danger,
				b.Title(html.P, "Wide tile"),
				b.Subtitle(html.P, "Aligned with the right tile"),
				b.Content(
				// Content
				),
			),
		),
	),
	b.Tile(
		b.Notification(
			html.Article, b.Success,
			b.Content(
				b.Title(html.P, "Tall tile"),
				b.Subtitle(html.P, "With even more content"),
				b.Content(
				// Content
				),
			),
		),
	),
)`,
		b.Tile(
			b.VTile(
				b.Size8,
				b.Tile(
					b.VTile(
						b.Notification(
							html.Article, b.Warning,
							b.Title(html.P, "Vertical..."),
							b.Subtitle(html.P, "Top tile"),
						),
						b.Notification(
							html.Article, b.Warning,
							b.Title(html.P, "...tiles"),
							b.Subtitle(html.P, "Bottom tile"),
						),
					),
					b.Tile(
						b.Notification(
							html.Article, b.Info,
							b.Title(html.P, "Middle tile"),
							b.Subtitle(html.P, "With an image"),
							b.ImageImg(
								"https://bulma.io/images/placeholders/640x480.png",
								b.Img4By3,
							),
						),
					),
				),

				b.Tile(
					b.Notification(
						html.Article, b.Danger,
						b.Title(html.P, "Wide tile"),
						b.Subtitle(html.P, "Aligned with the right tile"),
						b.Content(
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
						),
					),
				),
			),
			b.Tile(
				b.Notification(
					html.Article, b.Success,
					b.Content(
						b.Title(html.P, "Tall tile"),
						b.Subtitle(html.P, "With even more content"),
						b.Content(
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
							"Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor.",
							"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
						),
					),
				),
			),
		),
	),
).Section(
	"Modifiers",
	"https://bulma.io/documentation/layout/tiles/#modifiers",
	b.Content(
		el.Ul(
			el.Li(
				el.Strong("3 contextual"), " elements",
				el.Ul(
					el.Li("the ", el.Em("ancestor"), " modifier is automatically added when a tile's parent is not a tile"),
					el.Li("the ", el.Em("parent"), " modifier is automatically added when a tile has non-tile child elements"),
					el.Li("the ", el.Em("child"), " modifier is automatically added to contained elements: do ", el.Strong("not"), " define child tiles manually"),
				),
			),
			el.Li(
				el.Strong("1 directional"), " modifier",
				el.Ul(
					el.Li("the ", el.Em("vertical"), " modifier is added by using", el.Code("b.VTile"), " instead of ", el.Code("b.Tile")),
				),
			),
			el.Li(
				el.Strong("12 horizontal size"), " modifiers",
				el.Ul(
					el.Li("from ", el.Code("b.Size1")),
					el.Li("to ", el.Code("b.Size12")),
				),
			),
		),
		el.P("Please note the ", el.Code("b.Tile"), " magic for contextual elements only works with ", el.Em("Bulma-Gomponents"), " generated elements as children, not with regular ", el.Em("Gomponents"), " nodes."),
	),
).Section(
	"How it works: Nesting",
	"https://bulma.io/documentation/layout/tiles/#how-it-works-nesting",
	c.HorizontalExample(
		`b.Tile(
	b.VTile(
		b.Size4,
		b.Box(
			b.Title(html.P, "One"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
		b.Box(
			b.Title(html.P, "Two"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
	),
	b.Tile(
		b.Box(
			b.Title(html.P, "Three"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
			el.P("Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor."),
			el.P("Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo."),
		),
	),
)`,
		b.Tile(
			b.VTile(
				b.Size4,
				b.Box(
					b.Title(html.P, "One"),
					el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
				b.Box(
					b.Title(html.P, "Two"),
					el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
			),
			b.Tile(
				b.Box(
					b.Title(html.P, "Three"),
					el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
					el.P("Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor."),
					el.P("Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo."),
				),
			),
		),
	),
).Section(
	"Nesting requirements",
	"https://bulma.io/documentation/layout/tiles/#nesting-requirements",
	b.Content(
		el.Pre(`b.Tile
|
└───b.Tile
	|
	└───*b.Element`),
		el.Pre(`b.Tile
|
├───b.VTile b.Size8
|   |
|   ├───b.Tile
|   |   |
|   |   ├───b.VTile
|   |   |   ├───*b.Element
|   |   |   └───*b.Element
|   |   |
|   |   └───b.Tile
|   |       └───*b.Element
|   |
|   └───b.Tile
|       └───*b.Element
|
└───b.Tile
	└───*b.Element`),
	),
	c.HorizontalExample(
		`b.Tile(
	b.VTile(b.Size8,
		b.Tile(
			b.VTile(
				b.Box(html.Article,
					// Put any content you want
				),
				b.Box(html.Article,
					// Put any content you want
				),
			),
			b.Tile(
				b.Box(html.Article,
					// Put any content you want
				),
			),
		),
		b.Tile(
			b.Box(html.Article,
				// Put any content you want
			),
		),
	),
	b.Tile(
		b.Box(html.Article,
			// Put any content you want
		),
	),
)`,
		b.Tile(
			b.VTile(b.Size8,
				b.Tile(
					b.VTile(
						b.Box(html.Article,
							b.Title(html.P, "Vertical tiles"),
							b.Subtitle(html.P, "Top box"),
						),
						b.Box(html.Article,
							b.Title(html.P, "Vertical tiles"),
							b.Subtitle(html.P, "Bottom box"),
						),
					),
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Middle box"),
							b.Subtitle(html.P, "With an image"),
							b.ImageImg("https://bulma.io/images/placeholders/640x480.png", b.Img4By3),
						),
					),
				),
				b.Tile(
					b.Box(html.Article,
						b.Title(html.P, "Wide column"),
						b.Subtitle(html.P, "Aligned with the right column"),
						b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
					),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Tall column"),
					b.Subtitle(html.P, "With even more content"),
					b.Content(
						"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
						"Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor.",
						"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
					),
				),
			),
		),
	),
).Section(
	"3 columns",
	"https://bulma.io/documentation/layout/tiles/#3-columns",
	c.HorizontalExample(
		`b.Tile(
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Hello World"),
			b.Subtitle(html.P, "What is up?"),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Foo"),
			b.Subtitle(html.P, "Bar"),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Third column"),
			b.Subtitle(html.P, "With some content"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
	),
),
b.Tile(
	b.VTile(b.Size8,
		b.Tile(
			b.VTile(
				b.Box(html.Article,
					b.Title(html.P, "Vertical tiles"),
					b.Subtitle(html.P, "Top box"),
				),
				b.Box(html.Article,
					b.Title(html.P, "Vertical tiles"),
					b.Subtitle(html.P, "Bottom box"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Middle box"),
					b.Subtitle(html.P, "With an image"),
					b.ImageImg("https://bulma.io/images/placeholders/640x480.png", b.Img4By3),
				),
			),
		),
		b.Tile(
			b.Box(html.Article,
				b.Title(html.P, "Wide column"),
				b.Subtitle(html.P, "Aligned with the right column"),
				b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
			),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Tall column"),
			b.Subtitle(html.P, "With even more content"),
			b.Content(
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
				"Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor.",
				"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
			),
		),
	),
),
b.Tile(
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Side column"),
			b.Subtitle(html.P, "With some content"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
	),
	b.Tile(b.Size8,
		b.Box(html.Article,
			b.Title(html.P, "Main column"),
			b.Subtitle(html.P, "With some content"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
	),
)`,
		b.Tile(
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Hello World"),
					b.Subtitle(html.P, "What is up?"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Foo"),
					b.Subtitle(html.P, "Bar"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Third column"),
					b.Subtitle(html.P, "With some content"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
			),
		),
		b.Tile(
			b.VTile(b.Size8,
				b.Tile(
					b.VTile(
						b.Box(html.Article,
							b.Title(html.P, "Vertical tiles"),
							b.Subtitle(html.P, "Top box"),
						),
						b.Box(html.Article,
							b.Title(html.P, "Vertical tiles"),
							b.Subtitle(html.P, "Bottom box"),
						),
					),
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Middle box"),
							b.Subtitle(html.P, "With an image"),
							b.ImageImg("https://bulma.io/images/placeholders/640x480.png", b.Img4By3),
						),
					),
				),
				b.Tile(
					b.Box(html.Article,
						b.Title(html.P, "Wide column"),
						b.Subtitle(html.P, "Aligned with the right column"),
						b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
					),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Tall column"),
					b.Subtitle(html.P, "With even more content"),
					b.Content(
						"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
						"Suspendisse varius ligula in molestie lacinia. Maecenas varius eget ligula a sagittis. Pellentesque interdum, nisl nec interdum maximus, augue diam porttitor lorem, et sollicitudin felis neque sit amet erat. Maecenas imperdiet felis nisi, fringilla luctus felis hendrerit sit amet. Aenean vitae gravida diam, finibus dignissim turpis. Sed eget varius ligula, at volutpat tortor.",
						"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
					),
				),
			),
		),
		b.Tile(
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Side column"),
					b.Subtitle(html.P, "With some content"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
			),
			b.Tile(b.Size8,
				b.Box(html.Article,
					b.Title(html.P, "Main column"),
					b.Subtitle(html.P, "With some content"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
			),
		),
	),
).Section(
	"4 columns",
	"https://bulma.io/documentation/layout/tiles/#4-columns",
	c.HorizontalExample(
		`b.Tile(
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "One"),
			b.Subtitle(html.P, "Subtitle"),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Two"),
			b.Subtitle(html.P, "Subtitle"),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Three"),
			b.Subtitle(html.P, "Subtitle"),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Four"),
			b.Subtitle(html.P, "Subtitle"),
		),
	),
),
b.Tile(
	b.VTile(b.Size9,
		b.Tile(
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Five"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
				),
			),
			b.VTile(b.Size8,
				b.Tile(
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Six"),
							b.Subtitle(html.P, "Subtitle"),
						),
					),
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Seven"),
							b.Subtitle(html.P, "Subtitle"),
						),
					),
				),
				b.Tile(
					b.Box(html.Article,
						b.Title(html.P, "Eight"),
						b.Subtitle(html.P, "Subtitle"),
					),
				),
			),
		),
		b.Tile(
			b.Tile(b.Size8,
				b.Box(html.Article,
					b.Title(html.P, "Nine"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Ten"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
				),
			),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Content(
				b.Title(html.P, "Eleven"),
				b.Subtitle(html.P, "Subtitle"),
				b.Content(
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
					"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
				),
			),
		),
	),
),
b.Tile(
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Twelve"),
			b.Subtitle(html.P, "Subtitle"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut."),
		),
	),
	b.Tile(b.Size6,
		b.Box(html.Article,
			b.Title(html.P, "Thirteen"),
			b.Subtitle(html.P, "Subtitle"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
		),
	),
	b.Tile(
		b.Box(html.Article,
			b.Title(html.P, "Fourteen"),
			b.Subtitle(html.P, "Subtitle"),
			b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut."),
		),
	),
)`,
		b.Tile(
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "One"),
					b.Subtitle(html.P, "Subtitle"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Two"),
					b.Subtitle(html.P, "Subtitle"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Three"),
					b.Subtitle(html.P, "Subtitle"),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Four"),
					b.Subtitle(html.P, "Subtitle"),
				),
			),
		),
		b.Tile(
			b.VTile(b.Size9,
				b.Tile(
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Five"),
							b.Subtitle(html.P, "Subtitle"),
							b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
						),
					),
					b.VTile(b.Size8,
						b.Tile(
							b.Tile(
								b.Box(html.Article,
									b.Title(html.P, "Six"),
									b.Subtitle(html.P, "Subtitle"),
								),
							),
							b.Tile(
								b.Box(html.Article,
									b.Title(html.P, "Seven"),
									b.Subtitle(html.P, "Subtitle"),
								),
							),
						),
						b.Tile(
							b.Box(html.Article,
								b.Title(html.P, "Eight"),
								b.Subtitle(html.P, "Subtitle"),
							),
						),
					),
				),
				b.Tile(
					b.Tile(b.Size8,
						b.Box(html.Article,
							b.Title(html.P, "Nine"),
							b.Subtitle(html.P, "Subtitle"),
							b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
						),
					),
					b.Tile(
						b.Box(html.Article,
							b.Title(html.P, "Ten"),
							b.Subtitle(html.P, "Subtitle"),
							b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam."),
						),
					),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Content(
						b.Title(html.P, "Eleven"),
						b.Subtitle(html.P, "Subtitle"),
						b.Content(
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam semper diam at erat pulvinar, at pulvinar felis blandit. Vestibulum volutpat tellus diam, consequat gravida libero rhoncus ut. Morbi maximus, leo sit amet vehicula eleifend, nunc dui porta orci, quis semper odio felis ut quam.",
							"Integer sollicitudin, tortor a mattis commodo, velit urna rhoncus erat, vitae congue lectus dolor consequat libero. Donec leo ligula, maximus et pellentesque sed, gravida a metus. Cras ullamcorper a nunc ac porta. Aliquam ut aliquet lacus, quis faucibus libero. Quisque non semper leo.",
						),
					),
				),
			),
		),
		b.Tile(
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Twelve"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut."),
				),
			),
			b.Tile(b.Size6,
				b.Box(html.Article,
					b.Title(html.P, "Thirteen"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis."),
				),
			),
			b.Tile(
				b.Box(html.Article,
					b.Title(html.P, "Fourteen"),
					b.Subtitle(html.P, "Subtitle"),
					b.Content("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut."),
				),
			),
		),
	),
)
