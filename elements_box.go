package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var box = c.NewPage(
	"Box", "Box", "/box",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Box"), " constructor returns a white box with some padding and a shadow.",
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
			b.InputEmail(e.Placeholder("e.g. alex@example.com")),
		),
	),
	b.Field(
		b.Label("Password"),
		b.Control(
			b.InputPassword(e.Placeholder("********")),
		),
	),
	b.Button(b.Primary, "Sign in"),
)`,
		b.Box(
			html.FormEl,
			b.Field(
				b.Label("Email"),
				b.Control(
					b.InputEmail(e.Placeholder("e.g. alex@example.com")),
				),
			),
			b.Field(
				b.Label("Password"),
				b.Control(
					b.InputPassword(e.Placeholder("********")),
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
			e.P(
				e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
				e.Br(),
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
			),
		),
		b.Level(
			b.Mobile,
			b.LevelLeft(
				e.A(e.AriaLabel("reply"), fa.Icon(fa.Solid, "reply", b.Small)),
				e.A(e.AriaLabel("retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
				e.A(e.AriaLabel("like"), fa.Icon(fa.Solid, "heart", b.Small)),
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
					e.P(
						e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
						e.Br(),
						"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
					),
				),
				b.Level(
					b.Mobile,
					b.LevelLeft(
						e.A(e.AriaLabel("reply"), fa.Icon(fa.Solid, "reply", b.Small)),
						e.A(e.AriaLabel("retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
						e.A(e.AriaLabel("like"), fa.Icon(fa.Solid, "heart", b.Small)),
					),
				),
			),
		),
	),
)
