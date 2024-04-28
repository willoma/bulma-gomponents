package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var box = c.NewPage(
	"Box", "Box", "/box",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Box"), " constructor returns a white box with some padding and a shadow.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/box/",
	c.Example(
		`b.Box("I'm in a box.")`,
		b.Box("I'm in a box."),
	),

	c.Example(
		`b.Box(
	html.FormEl,
	b.Field(
		b.Label("Email"),
		b.Control(
			b.InputEmail(html.Placeholder("e.g. alex@example.com")),
		),
	),
	b.Field(
		b.Label("Password"),
		b.Control(
			b.InputPassword(html.Placeholder("********")),
		),
	),
	b.Button(b.Primary, "Sign in"),
)`,
		b.Box(
			html.FormEl,
			b.Field(
				b.Label("Email"),
				b.Control(
					b.InputEmail(html.Placeholder("e.g. alex@example.com")),
				),
			),
			b.Field(
				b.Label("Password"),
				b.Control(
					b.InputPassword(html.Placeholder("********")),
				),
			),
			b.Button(b.Primary, "Sign in"),
		),
	),

	c.Example(
		`b.Box(
	b.Media(
		b.MediaLeft(
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/128x128.png",
				b.ImgSq64,
				b.ImgAlt("Image"),
			),
		),
		b.Content(
			el.P(
				el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
				el.Br(),
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
			),
		),
		b.Level(
			b.Mobile,
			b.LevelLeft(
				el.A(html.Aria("level", "reply"), fa.Icon(fa.Solid, "reply", b.Small)),
				el.A(html.Aria("level", "retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
				el.A(html.Aria("level", "like"), fa.Icon(fa.Solid, "heart", b.Small)),
			),
		),
	),
)`,
		b.Box(
			b.Media(
				b.MediaLeft(
					b.ImageImg(
						"https://bulma.io/assets/images/placeholders/128x128.png",
						b.ImgSq64,
						b.ImgAlt("Image"),
					),
				),
				b.Content(
					el.P(
						el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
						el.Br(),
						"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
					),
				),
				b.Level(
					b.Mobile,
					b.LevelLeft(
						el.A(html.Aria("level", "reply"), fa.Icon(fa.Solid, "reply", b.Small)),
						el.A(html.Aria("level", "retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
						el.A(html.Aria("level", "like"), fa.Icon(fa.Solid, "heart", b.Small)),
					),
				),
			),
		),
	),
)
