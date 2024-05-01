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
			"The ", e.Code("b.Card"), " constructor creates a card. Non-specific content is automatically grouped in a ", e.Code(`<div class="card-content">`), " e.Element. If you need to force starting a new card content, use the ", e.Code("b.SplitContent()"), " function. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnCard(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="card">`), " e.Element"},

			e.Code("b.OnContent(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="card-content">`), " e.Element being currently built"},

			e.Code("b.CardHeader(...)"),
			"Add items to the card header",

			e.Code("b.CardHeaderIcon(...)"),
			"Add an icon to the card header",

			e.Code("b.CardHeaderTitle(...)"),
			"Add a title to the card header",

			e.Code("b.CardFooter(...)"),
			"Add items to the card footer",

			e.Code("b.CardImage(...)"),
			"Add this image to the card",

			e.Code("b.CardImageImg(...)"),
			"Add this image to the card",

			[]any{"one of the class or style types defined in package ", e.Code("b")},
			[]any{"Apply the class or style to the ", e.Code(`<div class="card">`), " e.Element"},

			e.Code("e.Element"),
			"Add this e.Element to the card content",

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the card",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the card content",
		),
		e.P("Other children are added to the ", e.Code(`<div class="card">`), " e.Element."),
		e.P(
			"The ", e.Code("b.CardHeader"), " function marks children as being part of the card header. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.IconElem"),
			"Add this icon as the card header icon",

			e.Code("string"),
			"Add this text as the card header title",
		),
		e.P(
			"The ", e.Code("b.CardHeaderIcon"), " constructor creates a card header icon e.Element, as an alternative to providing an icon to ", e.Code("b.CardHeader"), " allowing to customize the icon. The ", e.Code("b.CardHeaderTitle"), " constructor creates a card header title e.Element, as an alternative to providing a string to ", e.Code("b.CardHeader"), " allowing to customize the title.",
		),
		e.P(
			"The ", e.Code("b.CardImage"), " constructor creates an image by calling ", e.Code("b.Image"), " and wrapping it into a card image e.Element. The ", e.Code("b.CardImageImg"), " constructor creates an image by calling ", e.Code("b.ImageImg"), " and wrapping it into a card image e.Element. The following child has a special meaning:",
		),
		b.DList(
			e.Code("b.OnCardImage(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="card-image">`), " e.Element"},
		),
		e.P("Other children are added to the ", e.Code("<figure>"), " e.Element."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/card/",
	c.Example(
		`b.Card(
	b.CardImage(
		b.Img4By3,
		e.ImgSrc(
			"https://bulma.io/assets/images/placeholders/1280x960.png",
			html.Alt("Placeholder image"),
		),
	),
	b.Media(
		b.MediaLeft(
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/96x96.png",
				html.Alt("Placeholder image"),
			),
		),
		b.Title4(html.P, "John Smith"),
		b.Subtitle6(html.P, "@johnsmith"),
	),
	b.Content(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", e.A("@bulmaio"), ". ", e.AHref("#", "#css"), " ", e.AHref("#", "#responsive"), e.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016"),
	),
)`,
		b.Card(
			b.CardImage(
				b.Img4By3,
				e.ImgSrc(
					"https://bulma.io/assets/images/placeholders/1280x960.png",
					html.Alt("Placeholder image"),
				),
			),
			b.Media(
				b.MediaLeft(
					b.ImageImg(
						"https://bulma.io/assets/images/placeholders/96x96.png",
						html.Alt("Placeholder image"),
					),
				),
				b.Title4(html.P, "John Smith"),
				b.Subtitle6(html.P, "@johnsmith"),
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
	b.MarginBottom(b.Spacing4),
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
			b.MarginBottom(b.Spacing4),
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
