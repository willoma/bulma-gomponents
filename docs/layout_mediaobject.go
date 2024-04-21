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
	"",

	b.Content(
		el.P("The ", el.Code("b.Media"), " constructor creates a media. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnContent(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<>"), " element"},

			el.Code("b.OnMedia(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<>"), " element"},

			el.Code("b.MediaLeft(...)"),
			[]any{"Apply children to the left section"},

			el.Code("b.MediaRight(...)"),
			[]any{"Apply children to the right section"},

			el.Code("b.Element"),
			"Apply to the content section",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the media",

			[]any{"Other ", el.Code("gomponents.Node")},
			[]any{"Apply to the content section"},
		),
		el.P("Each of the left, content and right parts is only included if it has content. Other children are added to the ", el.Code(`<article class="media">`), " element."),
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
		el.P(
			el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
			html.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
		),
		b.Level(
			b.Mobile,
			b.LevelLeft(
				el.A(fa.Icon(fa.Solid, "reply", b.Small)),
				el.A(fa.Icon(fa.Solid, "retweet", b.Small)),
				el.A(fa.Icon(fa.Solid, "heart", b.Small)),
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
				el.P(
					el.Strong("John Smith"), " ", el.Small("@johnsmith"), " ", el.Small("31m"),
					html.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
				),
				b.Level(
					b.Mobile,
					b.LevelLeft(
						el.A(fa.Icon(fa.Solid, "reply", b.Small)),
						el.A(fa.Icon(fa.Solid, "retweet", b.Small)),
						el.A(fa.Icon(fa.Solid, "heart", b.Small)),
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
					"https://bulma.io/assets/images/placeholders/128x128.png",
					html.P, b.ImgSq64,
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
					el.Div(
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
				"https://bulma.io/assets/images/placeholders/96x96.png",
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
				"https://bulma.io/assets/images/placeholders/96x96.png",
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
			"https://bulma.io/assets/images/placeholders/128x128.png",
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
					"https://bulma.io/assets/images/placeholders/128x128.png",
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
						"https://bulma.io/assets/images/placeholders/96x96.png",
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
						"https://bulma.io/assets/images/placeholders/96x96.png",
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
					"https://bulma.io/assets/images/placeholders/128x128.png",
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
