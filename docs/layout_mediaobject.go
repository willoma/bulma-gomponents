package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var mediaObject = c.NewPage(
	"Media Object", "Media Object", "/media-object",
	"https://bulma.io/documentation/layout/media-object/",
	c.Example(
		`b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Content(
		el.P(
			el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
			html.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
		),
		b.Level(
			b.Mobile,
			b.LevelLeft(
				el.A(
					b.LevelItem,
					fa.Icon(fa.Solid, "reply", b.Small),
				),
				b.LevelItem(
					html.A,
					fa.Icon(fa.Solid, "retweet", b.Small),
				),
				el.A(
					b.Class("level-item"),
					fa.Icon(fa.Solid, "heart", b.Small),
				),
			),
		),
	),
	b.MediaRight(b.Delete()),
)`,
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Content(
				el.P(
					el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
					html.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
				),
				b.Level(
					b.Mobile,
					b.LevelLeft(
						el.A(
							b.LevelItem,
							fa.Icon(fa.Solid, "reply", b.Small),
						),
						b.LevelItem(
							html.A,
							fa.Icon(fa.Solid, "retweet", b.Small),
						),
						el.A(
							b.Class("level-item"),
							fa.Icon(fa.Solid, "heart", b.Small),
						),
					),
				),
			),
			b.MediaRight(b.Delete()),
		),
	),
	c.Example(
		`b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/images/placeholders/128x128.png",
			el.P, b.ImgSq64,
		),
	),
	b.Field(
		b.Control(
			b.Textarea(
				html.Placeholder("Add a comment..."),
			),
		),
	),
	b.Level(
		b.LevelLeft(
			b.LevelItem(
				b.ButtonA(b.Info, "Submit"),
			),
		),
		b.LevelRight(
			b.Checkbox("Press enter to submit"),
		),
	),
)`,
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/images/placeholders/128x128.png",
					el.P, b.ImgSq64,
				),
			),
			b.Field(
				b.Control(
					b.Textarea(
						html.Placeholder("Add a comment..."),
					),
				),
			),
			b.Level(
				b.LevelLeft(
					b.LevelItem(
						b.ButtonA(b.Info, "Submit"),
					),
				),
				b.LevelRight(
					b.Checkbox("Press enter to submit"),
				),
			),
		),
	),
).Section(
	"Nesting",
	"https://bulma.io/documentation/layout/media-object/#nesting",
	c.Example(
		`b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Content(
		el.P(
			el.Strong("Barbara Middleton"),
			el.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porta eros lacus, nec ultricies elit blandit non. Suspendisse pellentesque mauris sit amet dolor blandit rutrum. Nunc in tempus turpis.",
			el.Br(),
			el.Small(el.A("Like"), " · ", el.A("Reply"), " · 3 hrs"),
		),
	),
	b.Media(
		b.MediaLeft(
			html.Figure,
			b.ImageImg(
				"https://bulma.io/images/placeholders/96x96.png",
				html.P, b.ImgSq48,
			),
		),
		b.Content(
			el.P(
				el.Strong("Sean Brown"),
				el.Br(),
				"Donec sollicitudin urna eget eros malesuada sagittis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam blandit nisl a nulla sagittis, a lobortis leo feugiat.",
				el.Br(),
				el.Small(el.A("Like"), " · ", el.A("Reply"), " · 2 hrs"),
			),
		),
		b.Media("Vivamus quis semper metus, non tincidunt dolor. Vivamus in mi eu lorem cursus ullamcorper sit amet nec massa."),
		b.Media("Morbi vitae diam et purus tincidunt porttitor vel vitae augue. Praesent malesuada metus sed pharetra euismod. Cras tellus odio, tincidunt iaculis diam non, porta aliquet tortor."),
	),
	b.Media(
		b.MediaLeft(
			html.Figure,
			b.ImageImg(
				"https://bulma.io/images/placeholders/96x96.png",
				html.P, b.ImgSq48,
			),
		),
		b.Content(
			el.P(
				el.Strong("Kayli Eunice"),
				el.Br(),
				"Sed convallis scelerisque mauris, non pulvinar nunc mattis vel. Maecenas varius felis sit amet magna vestibulum euismod malesuada cursus libero. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Phasellus lacinia non nisl id feugiat.",
				el.Br(),
				el.Small(el.A("Like"), " · ", el.A("Reply"), " · 2 hrs"),
			),
		),
	),
),
b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Field(b.Control(
		b.Textarea(
			html.Placeholder("Add a comment..."),
		),
	)),
	b.Field(b.Control(
		b.Button("Post comment"),
	)),
)`,
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Content(
				el.P(
					el.Strong("Barbara Middleton"),
					el.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porta eros lacus, nec ultricies elit blandit non. Suspendisse pellentesque mauris sit amet dolor blandit rutrum. Nunc in tempus turpis.",
					el.Br(),
					el.Small(el.A("Like"), " · ", el.A("Reply"), " · 3 hrs"),
				),
			),
			b.Media(
				b.MediaLeft(
					html.Figure,
					b.ImageImg(
						"https://bulma.io/images/placeholders/96x96.png",
						html.P, b.ImgSq48,
					),
				),
				b.Content(
					el.P(
						el.Strong("Sean Brown"),
						el.Br(),
						"Donec sollicitudin urna eget eros malesuada sagittis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam blandit nisl a nulla sagittis, a lobortis leo feugiat.",
						el.Br(),
						el.Small(el.A("Like"), " · ", el.A("Reply"), " · 2 hrs"),
					),
				),
				b.Media("Vivamus quis semper metus, non tincidunt dolor. Vivamus in mi eu lorem cursus ullamcorper sit amet nec massa."),
				b.Media("Morbi vitae diam et purus tincidunt porttitor vel vitae augue. Praesent malesuada metus sed pharetra euismod. Cras tellus odio, tincidunt iaculis diam non, porta aliquet tortor."),
			),
			b.Media(
				b.MediaLeft(
					html.Figure,
					b.ImageImg(
						"https://bulma.io/images/placeholders/96x96.png",
						html.P, b.ImgSq48,
					),
				),
				b.Content(
					el.P(
						el.Strong("Kayli Eunice"),
						el.Br(),
						"Sed convallis scelerisque mauris, non pulvinar nunc mattis vel. Maecenas varius felis sit amet magna vestibulum euismod malesuada cursus libero. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Phasellus lacinia non nisl id feugiat.",
						el.Br(),
						el.Small(el.A("Like"), " · ", el.A("Reply"), " · 2 hrs"),
					),
				),
			),
		),
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Field(b.Control(
				b.Textarea(
					html.Placeholder("Add a comment..."),
				),
			)),
			b.Field(b.Control(
				b.Button("Post comment"),
			)),
		),
	),
)
