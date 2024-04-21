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
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Card"), " constructor creates a card. Non-specific content is automatically grouped in a ", el.Code(`<div class="card-content">`), " element. If you need to force starting a new card content, use the ", el.Code("b.SplitContent()"), " function. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnCard(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="card">`), " element"},

			el.Code("b.OnContent(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="card-content">`), " element being currently built"},

			el.Code("b.CardHeader(...)"),
			"Add items to the card header",

			el.Code("b.CardHeaderIcon(...)"),
			"Add an icon to the card header",

			el.Code("b.CardHeaderTitle(...)"),
			"Add a title to the card header",

			el.Code("b.CardFooter(...)"),
			"Add items to the card footer",

			el.Code("b.CardImage(...)"),
			"Add this image to the card",

			el.Code("b.CardImageImg(...)"),
			"Add this image to the card",

			[]any{"one of the class or style types defined in package ", el.Code("b")},
			[]any{"Apply the class or style to the ", el.Code(`<div class="card">`), " element"},

			el.Code("b.Element"),
			"Add this element to the card content",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the card",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the card content",
		),
		el.P("Other children are added to the ", el.Code(`<div class="card">`), " element."),
		el.P(
			"The ", el.Code("b.CardHeader"), " function marks children as being part of the card header. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.IconElem"),
			"Add this icon as the card header icon",

			el.Code("string"),
			"Add this text as the card header title",
		),
		el.P(
			"The ", el.Code("b.CardHeaderIcon"), " constructor creates a card header icon element, as an alternative to providing an icon to ", el.Code("b.CardHeader"), " allowing to customize the icon. The ", el.Code("b.CardHeaderTitle"), " constructor creates a card header title element, as an alternative to providing a string to ", el.Code("b.CardHeader"), " allowing to customize the title.",
		),
		el.P(
			"The ", el.Code("b.CardImage"), " constructor creates an image by calling ", el.Code("b.Image"), " and wrapping it into a card image element. The ", el.Code("b.CardImageImg"), " constructor creates an image by calling ", el.Code("b.ImageImg"), " and wrapping it into a card image element. The following child has a special meaning:",
		),
		b.DList(
			el.Code("b.OnCardImage(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="card-image">`), " element"},
		),
		el.P("Other children are added to the ", el.Code("<figure>"), " element."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/card/",
	c.Example(
		`b.Card(
	b.CardImage(
		b.Img4By3,
		b.ImgSrc(
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
		gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
	),
)`,
		b.Card(
			b.CardImage(
				b.Img4By3,
				b.ImgSrc(
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
				b.NoP("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. ", el.A("@bulmaio"), ". ", b.AHref("#", "#css"), " ", b.AHref("#", "#responsive"), el.Br(), el.Time(gomponents.Attr("datetime", "2016-1-1"), "11:09 PM - 1 Jan 2016")),
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
b.Block(
	b.Card(
		b.CardHeaderTitle("Card header"),
		b.CardHeaderIcon(
			html.Aria("label", "more options"),
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
		b.Block(
			b.Card(
				b.CardHeaderTitle("Card header"),
				b.CardHeaderIcon(
					html.Aria("label", "more options"),
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
				b.ImgSrc(
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
		gomponents.Text("Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum."),
	),
)`,
		b.Card(
			b.Content(
				gomponents.Text("Lorem ipsum leo risus, porta ac consectetur ac, vestibulum at eros. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum."),
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
).Subsection(
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
	b.Content(
		gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
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
			b.Content(
				gomponents.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris. "), el.A("@bulmaio"), gomponents.Text(". "), b.AHref("#", "#css"), gomponents.Text(" "), b.AHref("#", "#responsive"), el.Br(), html.Time(gomponents.Attr("datetime", "2016-1-1"), gomponents.Text("11:09 PM - 1 Jan 2016")),
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
	b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
	b.Subtitle(html.P, "Jeff Atwood"),
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
			b.Title(html.P, "“There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.”"),
			b.Subtitle(html.P, "Jeff Atwood"),
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
