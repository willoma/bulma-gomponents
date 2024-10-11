package docs

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var mediaObject = c.NewPage(
	"Media Object", "Media Object", "/media-object",
	"",

	b.Content(
		e.P("The ", e.Code("b.Media"), " constructor creates a media."),
		c.Children(
			c.Row("b.OnMedia(...any)", "Apply children to the ", e.Code(`<article class="media">`), " element"),
			c.Row("b.OnContent(...any)", "Apply children to the ", e.Code(`<div class="media-content">`), " element"),
			c.Row("b.MediaLeft(...any)", "Apply children to the ", e.Code(`<div class="media-left">`), " element"),
			c.Row("b.MediaRight(...any)", "Apply children to the ", e.Code(`<div class="media-right">`), " element"),
			c.Row("e.Element", "Add element the ", e.Code(`<div class="media-content">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<article class="media">`), " element"),
			c.RowNodeElement("Apply child to the ", e.Code(`<div class="media-content">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<article class="media">`), " element"),
		),
		e.P("Each of the left, content and right parts is only included if it has content."),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/media-object/",
	c.Example(
		`b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/assets/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Content(
		e.P(
			e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
			e.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
		),
		b.Level(
			b.Mobile,
			b.LevelLeft(
				e.A(fa.Icon(fa.Solid, "reply", b.Small)),
				e.A(fa.Icon(fa.Solid, "retweet", b.Small)),
				e.A(fa.Icon(fa.Solid, "heart", b.Small)),
			),
		),
	),
	b.MediaRight(b.Delete()),
)`,
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/assets/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Content(
				e.P(
					e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
					e.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
				),
				b.Level(
					b.Mobile,
					b.LevelLeft(
						e.A(fa.Icon(fa.Solid, "reply", b.Small)),
						e.A(fa.Icon(fa.Solid, "retweet", b.Small)),
						e.A(fa.Icon(fa.Solid, "heart", b.Small)),
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
			"https://bulma.io/assets/images/placeholders/128x128.png",
			e.P, b.ImgSq64,
		),
	),
	b.Field(
		b.Control(
			b.Textarea(
				e.Placeholder("Add a comment..."),
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
					"https://bulma.io/assets/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Field(
				b.Control(
					b.Textarea(
						e.Placeholder("Add a comment..."),
					),
				),
			),
			b.Level(
				b.LevelLeft(
					e.Div(
						b.ButtonA(b.Info, "Submit"),
					),
				),
				b.LevelRight(
					b.Checkbox("Press enter to submit"),
				),
			),
		),
	),
).Subsection(
	"Nesting",
	"https://bulma.io/documentation/layout/media-object/#nesting",
	c.Example(
		`b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/assets/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Content(
		e.P(
			e.Strong("Barbara Middleton"),
			e.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porta eros lacus, nec ultricies elit blandit non. Suspendisse pellentesque mauris sit amet dolor blandit rutrum. Nunc in tempus turpis.",
			e.Br(),
			e.Small(e.A("Like"), " · ", e.A("Reply"), " · 3 hrs"),
		),
	),
	b.Media(
		b.MediaLeft(
			html.Figure,
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/96x96.png",
				html.P, b.ImgSq48,
			),
		),
		b.Content(
			e.P(
				e.Strong("Sean Brown"),
				e.Br(),
				"Donec sollicitudin urna eget eros malesuada sagittis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam blandit nisl a nulla sagittis, a lobortis leo feugiat.",
				e.Br(),
				e.Small(e.A("Like"), " · ", e.A("Reply"), " · 2 hrs"),
			),
		),
		b.Media("Vivamus quis semper metus, non tincidunt dolor. Vivamus in mi eu lorem cursus ullamcorper sit amet nec massa."),
		b.Media("Morbi vitae diam et purus tincidunt porttitor vel vitae augue. Praesent malesuada metus sed pharetra euismod. Cras tellus odio, tincidunt iaculis diam non, porta aliquet tortor."),
	),
	b.Media(
		b.MediaLeft(
			html.Figure,
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/96x96.png",
				html.P, b.ImgSq48,
			),
		),
		b.Content(
			e.P(
				e.Strong("Kayli Eunice"),
				e.Br(),
				"Sed convallis scelerisque mauris, non pulvinar nunc mattis vel. Maecenas varius felis sit amet magna vestibulum euismod malesuada cursus libero. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Phasellus lacinia non nisl id feugiat.",
				e.Br(),
				e.Small(e.A("Like"), " · ", e.A("Reply"), " · 2 hrs"),
			),
		),
	),
),
b.Media(
	b.MediaLeft(
		html.Figure,
		b.ImageImg(
			"https://bulma.io/assets/images/placeholders/128x128.png",
			html.P, b.ImgSq64,
		),
	),
	b.Field(b.Control(
		b.Textarea(
			e.Placeholder("Add a comment..."),
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
					"https://bulma.io/assets/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Content(
				e.P(
					e.Strong("Barbara Middleton"),
					e.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porta eros lacus, nec ultricies elit blandit non. Suspendisse pellentesque mauris sit amet dolor blandit rutrum. Nunc in tempus turpis.",
					e.Br(),
					e.Small(e.A("Like"), " · ", e.A("Reply"), " · 3 hrs"),
				),
			),
			b.Media(
				b.MediaLeft(
					html.Figure,
					b.ImageImg(
						"https://bulma.io/assets/images/placeholders/96x96.png",
						html.P, b.ImgSq48,
					),
				),
				b.Content(
					e.P(
						e.Strong("Sean Brown"),
						e.Br(),
						"Donec sollicitudin urna eget eros malesuada sagittis. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam blandit nisl a nulla sagittis, a lobortis leo feugiat.",
						e.Br(),
						e.Small(e.A("Like"), " · ", e.A("Reply"), " · 2 hrs"),
					),
				),
				b.Media("Vivamus quis semper metus, non tincidunt dolor. Vivamus in mi eu lorem cursus ullamcorper sit amet nec massa."),
				b.Media("Morbi vitae diam et purus tincidunt porttitor vel vitae augue. Praesent malesuada metus sed pharetra euismod. Cras tellus odio, tincidunt iaculis diam non, porta aliquet tortor."),
			),
			b.Media(
				b.MediaLeft(
					html.Figure,
					b.ImageImg(
						"https://bulma.io/assets/images/placeholders/96x96.png",
						html.P, b.ImgSq48,
					),
				),
				b.Content(
					e.P(
						e.Strong("Kayli Eunice"),
						e.Br(),
						"Sed convallis scelerisque mauris, non pulvinar nunc mattis vel. Maecenas varius felis sit amet magna vestibulum euismod malesuada cursus libero. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Phasellus lacinia non nisl id feugiat.",
						e.Br(),
						e.Small(e.A("Like"), " · ", e.A("Reply"), " · 2 hrs"),
					),
				),
			),
		),
		b.Media(
			b.MediaLeft(
				html.Figure,
				b.ImageImg(
					"https://bulma.io/assets/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
				),
			),
			b.Field(b.Control(
				b.Textarea(
					e.Placeholder("Add a comment..."),
				),
			)),
			b.Field(b.Control(
				b.Button("Post comment"),
			)),
		),
	),
)
