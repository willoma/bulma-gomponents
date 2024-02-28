package docs

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var card = c.NewPage(
	"Card", "Card", "/card",
	"https://bulma.io/documentation/components/card/",
	c.Example(
		`b.Card(
	b.CardImage(
		b.Img4By3,
		b.ImgSrc(
			"https://bulma.io/images/placeholders/1280x960.png",
			html.Alt("Placeholder image"),
		),
	),
	b.CardContent(
		b.Media(
			b.MediaLeft(
				b.ImageImg(
					"https://bulma.io/images/placeholders/96x96.png",
					html.Alt("Placeholder image"),
				),
			),
			b.Title4(html.P, "John Smith"),
			b.Subtitle6(html.P, "@johnsmith"),
		),
		b.Content(
			gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
		),
	),
)`,
		b.Card(
			b.CardImage(
				b.Img4By3,
				b.ImgSrc(
					"https://bulma.io/images/placeholders/1280x960.png",
					html.Alt("Placeholder image"),
				),
			),
			b.CardContent(
				b.Media(
					b.MediaLeft(
						b.ImageImg(
							"https://bulma.io/images/placeholders/96x96.png",
							html.Alt("Placeholder image"),
						),
					),
					b.Title4(html.P, "John Smith"),
					b.Subtitle6(html.P, "@johnsmith"),
				),
				b.Content(
					gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
				),
			),
		),
	),
).Section(
	"Card parts",
	"https://bulma.io/documentation/components/card/#card-parts",
	c.Example(
		`b.Block(
	b.Card(
		b.CardHeader(
			b.CardHeaderTitle(
				"Card header",
			),
			b.CardHeaderIcon(
				html.Aria("label", "more options"),
				fa.Icon(fa.Solid, "angle-down"),
			),
		),
	),
),
b.Block(
	b.Card(
		b.CardHeader(
			"Card header",
			fa.Icon(fa.Solid, "angle-down"),
		),
	),
)`,
		b.Block(
			b.Card(
				b.CardHeader(
					b.CardHeaderTitle(
						"Card header",
					),
					b.CardHeaderIcon(
						html.Aria("label", "more options"),
						fa.Icon(fa.Solid, "angle-down"),
					),
				),
			),
		),
		b.Block(
			b.Card(
				b.CardHeader(
					"Card header",
					fa.Icon(fa.Solid, "angle-down"),
				),
			),
		),
	),
	c.Example(
		`b.Card(
	b.MarginBottom(b.Spacing4),
	b.CardImage(
		b.Img4By3,
		b.ImgSrc(
			"https://bulma.io/images/placeholders/1280x960.png",
			html.Alt("Placeholder image"),
		),
	),
),
b.Card(
	b.CardImageImg(
		"https://bulma.io/images/placeholders/1280x960.png",
		b.Img4By3,
	),
)`,
		b.Card(
			b.MarginBottom(b.Spacing4),
			b.CardImage(
				b.Img4By3,
				b.ImgSrc(
					"https://bulma.io/images/placeholders/1280x960.png",
					html.Alt("Placeholder image"),
				),
			),
		),
		b.Card(
			b.CardImageImg(
				"https://bulma.io/images/placeholders/1280x960.png",
				b.Img4By3,
			),
		),
	),
	c.Example(
		`b.Card(
	b.CardContent(
		b.Content(
			gomponents.Text("Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum."),
		),
	),
)`,
		b.Card(
			b.CardContent(
				b.Content(
					gomponents.Text("Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum."),
				),
			),
		),
	),
	c.Example(
		`b.Card(
	b.CardFooter(
		b.AHref("#", "Save"),
		b.AHref("#", "Edit"),
		b.AHref("#", "Delete"),
	),
)`,
		b.Card(
			b.CardFooter(
				b.AHref("#", "Save"),
				b.AHref("#", "Edit"),
				b.AHref("#", "Delete"),
			),
		),
	),
).Section(
	"Examples",
	"https://bulma.io/documentation/components/card/#examples",
	c.Example(
		`b.Card(
	b.CardHeader(
		"Component",
		b.CardHeaderIcon(
			html.Aria("label", "more options"),
			fa.Icon(fa.Solid, "angle-down"),
		),
	),
	b.CardContent(
		b.Content(
			gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
		),
	),
	b.CardFooter(
		b.AHref("#", "Save"),
		b.AHref("#", "Edit"),
		b.AHref("#", "Delete"),
	),
)`,
		b.Card(
			b.CardHeader(
				"Component",
				b.CardHeaderIcon(
					html.Aria("label", "more options"),
					fa.Icon(fa.Solid, "angle-down"),
				),
			),
			b.CardContent(
				b.Content(
					gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
				),
			),
			b.CardFooter(
				b.AHref("#", "Save"),
				b.AHref("#", "Edit"),
				b.AHref("#", "Delete"),
			),
		),
	),
	c.Example(
		`b.Card(
	b.CardContent(
		b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
		b.Subtitle(html.P, "Jeff Atwood"),
	),
	b.CardFooter(
		el.P(
			el.Span(
				"View on ", b.AHref("https://twitter.com/codinghorror/status/506010907021828096", "Twitter"),
			),
		),
		el.P(
			el.Span(
				"Share on ", b.AHref("#", "Facebook"),
			),
		),
	),
)`,
		b.Card(
			b.CardContent(
				b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
				b.Subtitle(html.P, "Jeff Atwood"),
			),
			b.CardFooter(
				el.P(
					el.Span(
						"View on ", b.AHref("https://twitter.com/codinghorror/status/506010907021828096", "Twitter"),
					),
				),
				el.P(
					el.Span(
						"Share on ", b.AHref("#", "Facebook"),
					),
				),
			),
		),
	),
)
