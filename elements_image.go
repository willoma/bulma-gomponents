package docs

import (
	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var image = c.NewPage(
	"Image", "Image", "/image",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Image"), " constructor creates a Bulma image element.",
		),
		c.Modifiers(
			c.Row("b.ImgSq16", "Fixed square image, of 16x16px"),
			c.Row("b.ImgSq24", "Fixed square image, of 24x24px"),
			c.Row("b.ImgSq32", "Fixed square image, of 32x32px"),
			c.Row("b.ImgSq48", "Fixed square image, of 48x48px"),
			c.Row("b.ImgSq64", "Fixed square image, of 64x64px"),
			c.Row("b.ImgSq96", "Fixed square image, of 96x96px"),
			c.Row("b.ImgSq128", "Fixed square image, of 128x128px"),
			c.Row("b.ImgSquare", "Force image ratio to square"),
			c.Row("b.Img1By1", "Force image ratio to 1 by 1"),
			c.Row("b.Img5By4", "Force image ratio to 5 by 4"),
			c.Row("b.Img4By3", "Force image ratio to 4 by 3"),
			c.Row("b.Img3By2", "Force image ratio to 3 by 2"),
			c.Row("b.Img5By3", "Force image ratio to 5 by 3"),
			c.Row("b.Img16By9", "Force image ratio to 16 by 9"),
			c.Row("b.Img2By1", "Force image ratio to 2 by 1"),
			c.Row("b.Img3By1", "Force image ratio to 3 by 1"),
			c.Row("b.Img4By5", "Force image ratio to 4 by 5"),
			c.Row("b.Img3By4", "Force image ratio to 3 by 4"),
			c.Row("b.Img2By3", "Force image ratio to 2 by 3"),
			c.Row("b.Img3By5", "Force image ratio to 3 by 5"),
			c.Row("b.Img9By16", "Force image ratio to 9 by 16"),
			c.Row("b.Img1By2", "Force image ratio to 1 by 2"),
			c.Row("b.Img1By3", "Force image ratio to 1 by 3"),
			c.Row("b.FullWidth", "Make the image take the whole width of its container"),
		),
		e.P(
			"Use ", e.Code("e.ImgSrc"), " to create an ", e.Code("<img>"), " element with the provided URL as its src attribute. Apply ", e.Code("b.Rounded"), " to the inner image to make it rounded, associated with an ", e.Code("b.Img*By*"), " modifier on the ", e.Code("b.Image"), " element. Apply ", e.Code("b.Ratio"), " to an inner element to apply the parent ratio to that element.",
		),
		e.P(
			"The ", e.Code("b.ImageImg"), " constructor creates a Bulma image element with an inner ", e.Code("<img>"), ". It accepts the same modifiers as ", e.Code("b.Image"), ".",
		),
		c.Children(
			c.Row("b.OnImg(...any)", "Apply children to the ", e.Code("<img>"), " element"),
			c.Row("b.OnFigure(...any)", "Apply children to the ", e.Code("<figure>"), " element"),
			c.Row("b.Rounded", "Make the image rounded (class applied to the ", e.Code("<img>"), " element)"),
			c.Row("b.ImgAlt", "Define the image alt text"),
			c.RowDefault("Apply children to the ", e.Code("<figure>"), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/image/",

	c.Example(
		`b.Image(
	b.ImgSq128,
	e.Img(e.Src("https://bulma.io/assets/images/placeholders/128x128.png")),
)`,
		b.Image(
			b.ImgSq128,
			e.Img(e.Src("https://bulma.io/assets/images/placeholders/128x128.png")),
		),
	),
	c.Example(
		`b.Image(
	b.ImgSq128,
	e.ImgSrc("https://bulma.io/assets/images/placeholders/128x128.png"),
)`,
		b.Image(
			b.ImgSq128,
			e.ImgSrc("https://bulma.io/assets/images/placeholders/128x128.png"),
		),
	),
	c.Example(
		`b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128)`,
		b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128),
	),
).Subsection(
	"Fixed square images",
	"https://bulma.io/documentation/elements/image/#fixed-square-images",

	c.Example(
		`e.P("16x16px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/16x16.png", b.ImgSq16),
e.P("24x24px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/24x24.png", b.ImgSq24),
e.P("32x32px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/32x32.png", b.ImgSq32),
e.P("48x48px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/48x48.png", b.ImgSq48),
e.P("64x64px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/64x64.png", b.ImgSq64),
e.P("96x96px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/96x96.png", b.ImgSq96),
e.P("128x128px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128)`,
		e.P("16x16px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/16x16.png", b.ImgSq16),
		e.P("24x24px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/24x24.png", b.ImgSq24),
		e.P("32x32px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/32x32.png", b.ImgSq32),
		e.P("48x48px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/48x48.png", b.ImgSq48),
		e.P("64x64px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/64x64.png", b.ImgSq64),
		e.P("96x96px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/96x96.png", b.ImgSq96),
		e.P("128x128px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128),
	),
).Subsection(
	"Rounded images",
	"https://bulma.io/documentation/elements/image/#rounded-images",
	c.Example(
		`b.Image(
			b.ImgSq128,
			e.ImgSrc(
				"https://bulma.io/assets/images/placeholders/128x128.png",
				b.Rounded,
			),
		)`,
		b.Image(
			b.ImgSq128,
			e.ImgSrc(
				"https://bulma.io/assets/images/placeholders/128x128.png",
				b.Rounded,
			),
		),
	),
	c.Example(
		`b.ImageImg(
	"https://bulma.io/assets/images/placeholders/128x128.png",
	b.ImgSq128, b.Rounded,
)`,
		b.ImageImg(
			"https://bulma.io/assets/images/placeholders/128x128.png",
			b.ImgSq128, b.Rounded,
		),
	),
).Subsection(
	"Retina images",
	"https://bulma.io/documentation/elements/image/#retina-images",
	c.Example(
		`b.ImageImg(
	"https://bulma.io/assets/images/placeholders/256x256.png",
	b.ImgSq128,
)`,
		b.ImageImg(
			"https://bulma.io/assets/images/placeholders/256x256.png",
			b.ImgSq128,
		),
	),
).Subsection(
	"Responsive images with ratios",
	"https://bulma.io/documentation/elements/image/#responsive-images-with-ratios",

	c.Example(
		`e.Div(
	e.Styles{"width": "10rem"},
	e.P("Square (or 1 by 1):"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x480.png",
		b.ImgSquare,
	),
	e.P("1 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x480.png",
		b.Img1By1,
	),
	e.P("5 by 4:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/600x480.png",
		b.Img5By4,
	),
	e.P("4 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x480.png",
		b.Img4By3,
	),
	e.P("3 by 2:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x320.png",
		b.Img3By2,
	),
	e.P("5 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/800x480.png",
		b.Img5By3,
	),
	e.P("16 by 9:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x360.png",
		b.Img16By9,
	),
	e.P("2 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x320.png",
		b.Img2By1,
	),
	e.P("3 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/720x240.png",
		b.Img3By1,
	),
	e.P("4 by 5:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x600.png",
		b.Img4By5,
	),
	e.P("3 by 4:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x640.png",
		b.Img3By4,
	),
	e.P("2 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/320x480.png",
		b.Img2By3,
	),
	e.P("3 by 5:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x800.png",
		b.Img3By5,
	),
	e.P("9 by 16:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/360x640.png",
		b.Img9By16,
	),
	e.P("1 by 2:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/320x640.png",
		b.Img1By2,
	),
	e.P("1 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/240x720.png",
		b.Img1By3,
	),
)`,
		e.Div(
			e.Styles{"width": "10rem"},
			e.P("Square (or 1 by 1):"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x480.png",
				b.ImgSquare,
			),
			e.P("1 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x480.png",
				b.Img1By1,
			),
			e.P("5 by 4:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/600x480.png",
				b.Img5By4,
			),
			e.P("4 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x480.png",
				b.Img4By3,
			),
			e.P("3 by 2:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x320.png",
				b.Img3By2,
			),
			e.P("5 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/800x480.png",
				b.Img5By3,
			),
			e.P("16 by 9:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x360.png",
				b.Img16By9,
			),
			e.P("2 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x320.png",
				b.Img2By1,
			),
			e.P("3 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/720x240.png",
				b.Img3By1,
			),
			e.P("4 by 5:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x600.png",
				b.Img4By5,
			),
			e.P("3 by 4:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x640.png",
				b.Img3By4,
			),
			e.P("2 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/320x480.png",
				b.Img2By3,
			),
			e.P("3 by 5:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x800.png",
				b.Img3By5,
			),
			e.P("9 by 16:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/360x640.png",
				b.Img9By16,
			),
			e.P("1 by 2:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/320x640.png",
				b.Img1By2,
			),
			e.P("1 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/240x720.png",
				b.Img1By3,
			),
		),
	),
).Subsection(
	"Arbitrary ratios with any element",
	"https://bulma.io/documentation/elements/image/#arbitrary-ratios-with-any-element",
	c.Example(
		`b.Image(
	b.Img16By9,
	e.IFrame(
		b.Ratio,
		e.Width("640"),
		e.Height("360"),
		e.Src("https://www.youtube.com/embed/YE7VzlLtp-4"),
		gomponents.Attr("frameborder", "0"),
		gomponents.Attr("allowfullscreen"),
	)
)`,
		b.Image(
			b.Img16By9,
			e.IFrame(
				b.Ratio,
				e.Width("640"),
				e.Height("360"),
				e.Src("https://www.youtube.com/embed/YE7VzlLtp-4"),
				gomponents.Attr("frameborder", "0"),
				gomponents.Attr("allowfullscreen"),
			),
		),
	),
)
