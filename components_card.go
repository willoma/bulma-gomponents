package docs

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var card = c.NewPage(
	"Card", "Card", "/card",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Card"), " constructor creates a card.",
		),
		c.Children(
			c.Row("b.OnCard(...any)", "Apply children to the ", e.Code(`<div class="card">`), " element"),
			c.Row("b.OnContent(...any)", "Apply children to the current ", e.Code(`<div class="card-content">`), " element"),
			c.Row("b.SplitContent()", "Force starting a new ", e.Code(`<div class="card-content">`), " element"),
			c.Row("b.CardHeader(...any)", "Apply children to the ", e.Code(`<header class="card-header">`), " element"),
			c.Row("b.CardHeaderIcon(...any", "Apply an icon to the ", e.Code(`<header class="card-header">`), " element"),
			c.Row("b.CardHeaderTitle(...any)", "Add a title to the ", e.Code(`<header class="card-header">`), " element"),
			c.Row("b.CardFooter(...any)", "Apply children to the ", e.Code(`<header class="card-footer">`), " element"),
			c.Row("b.CardImage(...any)", "Add image to the ", e.Code(`<div class="card">`), " element"),
			c.Row("b.CardImageImg(...any)", "Add image to the ", e.Code(`<div class="card">`), " element"),
			c.Row("e.Element", "Add element to the ", e.Code(`<div class="card-content">`), " element"),
			c.RowNodeAttribute("Apply child to the ", e.Code(`<div class="card">`), " element"),
			c.RowNodeElement("Apply child to the ", e.Code(`<div class="card-content">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="card-content">`), " element"),
		),
		e.P(
			"The ", e.Code("b.CardHeader"), " function marks children as being part of the card header.",
		),
		c.Children(
			c.Row("b.IconElem", "Add this icon as the card header icon"),
			c.Row("string", "Add this text as the card header title"),
			c.RowDefault("Apply child to the card header"),
		),
		e.P(
			"The ", e.Code("b.CardHeaderIcon"), " constructor creates a card header icon element, as an alternative to providing an icon to ", e.Code("b.CardHeader"), ", allowing to customize the icon. The ", e.Code("b.CardHeaderTitle"), " constructor creates a card header title element, as an alternative to providing a string to ", e.Code("b.CardHeader"), ", allowing to customize the title.",
		),
		e.P(
			"The ", e.Code("b.CardImage"), " constructor creates an image by calling ", e.Code("b.Image"), " and wrapping it into a card image element. The ", e.Code("b.CardImageImg"), " constructor creates an image by calling ", e.Code("b.ImageImg"), " and wrapping it into a card image element.",
		),
		c.Children(
			c.Row("b.OnCardImage(...any)", "Apply children to the ", e.Code(`<div class="card-image">`), " element"),
			c.RowDefault("Apply children to the ", e.Code(`<figure class="image">`), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/card/",
	c.Example(
		`b.Card(
	b.CardImage(
		b.Img4By3,
		e.ImgSrc(
			"https://bulma.io/assets/images/placeholders/1280x960.png",
			e.Alt("Placeholder image"),
		),
	),
	b.Media(
		b.MediaLeft(
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/96x96.png",
				e.Alt("Placeholder image"),
			),
		),
		b.Title(4, html.P, "John Smith"),
		b.Subtitle(6, html.P, "@johnsmith"),
	),
	b.Content(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", e.A("@bulmaio"), ". ", e.AHref("#", "#css"), " ", e.AHref("#", "#responsive"), e.Br(), e.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016"),
	),
)`,
		b.Card(
			b.CardImage(
				b.Img4By3,
				e.ImgSrc(
					"https://bulma.io/assets/images/placeholders/1280x960.png",
					e.Alt("Placeholder image"),
				),
			),
			b.Media(
				b.MediaLeft(
					b.ImageImg(
						"https://bulma.io/assets/images/placeholders/96x96.png",
						e.Alt("Placeholder image"),
					),
				),
				b.Title(4, html.P, "John Smith"),
				b.Subtitle(6, html.P, "@johnsmith"),
			),
			b.Content(
				b.NoP("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", e.A("@bulmaio"), ". ", e.AHref("#", "#css"), " ", e.AHref("#", "#responsive"), e.Br(), e.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016")),
			),
		),
	),
).Subsection(
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
				e.AriaLabel("more options"),
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
b.Block(
	b.Card(
		b.CardHeaderTitle("Card header"),
		b.CardHeaderIcon(
			e.AriaLabel("more options"),
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
						e.AriaLabel("more options"),
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
		b.Block(
			b.Card(
				b.CardHeaderTitle("Card header"),
				b.CardHeaderIcon(
					e.AriaLabel("more options"),
					fa.Icon(fa.Solid, "angle-down"),
				),
			),
		),
	),
	c.Example(
		`b.Card(
	b.MarginBottom(4),
	b.CardImage(
		b.Img4By3,
		e.ImgSrc(
			"https://bulma.io/assets/images/placeholders/1280x960.png",
			html.Alt("Placeholder image"),
		),
	),
),
b.Card(
	b.CardImageImg(
		"https://bulma.io/assets/images/placeholders/1280x960.png",
		b.Img4By3,
	),
)`,
		b.Card(
			b.MarginBottom(4),
			b.CardImage(
				b.Img4By3,
				e.ImgSrc(
					"https://bulma.io/assets/images/placeholders/1280x960.png",
					html.Alt("Placeholder image"),
				),
			),
		),
		b.Card(
			b.CardImageImg(
				"https://bulma.io/assets/images/placeholders/1280x960.png",
				b.Img4By3,
			),
		),
	),
	c.Example(
		`b.Card(
	b.Content(
		"Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum.",
	),
)`,
		b.Card(
			b.Content(
				"Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum.",
			),
		),
	),
	c.Example(
		`b.Card(
	b.CardFooter(
		e.AHref("#", "Save"),
		e.AHref("#", "Edit"),
		e.AHref("#", "Delete"),
	),
)`,
		b.Card(
			b.CardFooter(
				e.AHref("#", "Save"),
				e.AHref("#", "Edit"),
				e.AHref("#", "Delete"),
			),
		),
	),
).Subsection(
	"Examples",
	"https://bulma.io/documentation/components/card/#examples",
	c.Example(
		`b.Card(
	b.CardHeader(
		"Component",
		b.CardHeaderIcon(
			e.AriaLabel("more options"),
			fa.Icon(fa.Solid, "angle-down"),
		),
	),
	b.Content(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", e.A("@bulmaio"), ". ", e.AHref("#", "#css"), " ", e.AHref("#", "#responsive"), e.Br(), e.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016"),
	),
	b.CardFooter(
		e.AHref("#", "Save"),
		e.AHref("#", "Edit"),
		e.AHref("#", "Delete"),
	),
)`,
		b.Card(
			b.CardHeader(
				"Component",
				b.CardHeaderIcon(
					e.AriaLabel("more options"),
					fa.Icon(fa.Solid, "angle-down"),
				),
			),
			b.Content(
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", e.A("@bulmaio"), ". ", e.AHref("#", "#css"), " ", e.AHref("#", "#responsive"), e.Br(), e.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016"),
			),
			b.CardFooter(
				e.AHref("#", "Save"),
				e.AHref("#", "Edit"),
				e.AHref("#", "Delete"),
			),
		),
	),
	c.Example(
		`b.Card(
	b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
	b.Subtitle(html.P, "Jeff Atwood"),
	b.CardFooter(
		e.P(
			e.Span(
				"View on ", e.AHref("https://twitter.com/codinghorror/status/506010907021828096", "Twitter"),
			),
		),
		e.P(
			e.Span(
				"Share on ", e.AHref("#", "Facebook"),
			),
		),
	),
)`,
		b.Card(
			b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
			b.Subtitle(html.P, "Jeff Atwood"),
			b.CardFooter(
				e.P(
					e.Span(
						"View on ", e.AHref("https://twitter.com/codinghorror/status/506010907021828096", "Twitter"),
					),
				),
				e.P(
					e.Span(
						"Share on ", e.AHref("#", "Facebook"),
					),
				),
			),
		),
	),
)
