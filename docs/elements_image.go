package docs

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var image = c.NewPage(
	"Image", "Image", "/image",
	"https://bulma.io/documentation/elements/image/",

	c.Example(
		`b.Image(
	b.ImgSq128,
	el.Img(html.Src("https://bulma.io/images/placeholders/128x128.png")),
)`,
		b.Image(
			b.ImgSq128,
			el.Img(html.Src("https://bulma.io/images/placeholders/128x128.png")),
		),
	),
	b.Content(el.P("Bulma-Gomponents provides the ", el.Code("b.ImgSrc"), " helper:")),
	c.Example(
		`b.Image(
	b.ImgSq128,
	b.ImgSrc("https://bulma.io/images/placeholders/128x128.png"),
)`,
		b.Image(
			b.ImgSq128,
			b.ImgSrc("https://bulma.io/images/placeholders/128x128.png"),
		),
	),
	b.Content(el.P("Bulma-Gomponents also provides the ", el.Code("b.ImageImg"), " helper, for the simple ", el.Em("figure/img"), " case:")),
	c.Example(
		`b.ImageImg("https://bulma.io/images/placeholders/128x128.png", b.ImgSq128)`,
		b.ImageImg("https://bulma.io/images/placeholders/128x128.png", b.ImgSq128),
	),
).Section(
	"Fixed square images",
	"https://bulma.io/documentation/elements/image/#fixed-square-images",
	b.Content(el.P("Use ", el.Code("b.ImgSq16"), " to ", el.Code("b.ImgSq128"), ".")),
).Section(
	"Rounded images",
	"https://bulma.io/documentation/elements/image/#rounded-images",
	c.Example(
		`b.Image(
			b.ImgSq128,
			b.ImgSrc(
				"https://bulma.io/images/placeholders/128x128.png",
				b.Rounded,
			),
		)`,
		b.Image(
			b.ImgSq128,
			b.ImgSrc(
				"https://bulma.io/images/placeholders/128x128.png",
				b.Rounded,
			),
		),
	),
	b.Content(el.P("Bulma-Gomponents assigns the ", el.Code("b.Rounded"), " attribute to the ", el.Code("<img>"), " tag when it is provided to the ", el.Code("b.ImageImg"), " helper:")),
	c.Example(
		`b.ImageImg(
	"https://bulma.io/images/placeholders/128x128.png",
	b.ImgSq128, b.Rounded,
)`,
		b.ImageImg(
			"https://bulma.io/images/placeholders/128x128.png",
			b.ImgSq128, b.Rounded,
		),
	),
).Section(
	"Retina images",
	"https://bulma.io/documentation/elements/image/#retina-images",
	c.Example(
		`b.ImageImg(
	"https://bulma.io/images/placeholders/256x256.png",
	b.ImgSq128,
)`,
		b.ImageImg(
			"https://bulma.io/images/placeholders/256x256.png",
			b.ImgSq128,
		),
	),
).Section(
	"Responsive images with ratios",
	"https://bulma.io/documentation/elements/image/#responsive-images-with-ratios",
	b.Content(el.P("Use ", el.Code("b.Img*By*"), ".")),
).Section(
	"Arbitrary ratios with any element",
	"https://bulma.io/documentation/elements/image/#arbitrary-ratios-with-any-element",
	c.Example(
		`b.Image(
	b.Img16By9,
	el.IFrame(
		b.Ratio,
		html.Width("640"),
		html.Height("360"),
		html.Src("https://www.youtube.com/embed/YE7VzlLtp-4"),
		gomponents.Attr("frameborder", "0"),
		gomponents.Attr("allowfullscreen"),
	)
)`,
		b.Image(
			b.Img16By9,
			el.IFrame(
				b.Ratio,
				html.Width("640"),
				html.Height("360"),
				html.Src("https://www.youtube.com/embed/YE7VzlLtp-4"),
				gomponents.Attr("frameborder", "0"),
				gomponents.Attr("allowfullscreen"),
			),
		),
	),
)
