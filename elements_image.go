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
			"The ", e.Code("b.Image"), " constructor creates a Bulma image e.Element. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.ImgSq16"),
			"Fixed square image, of 16x16px",

			e.Code("b.ImgSq24"),
			"Fixed square image, of 24x24px",

			e.Code("b.ImgSq32"),
			"Fixed square image, of 32x32px",

			e.Code("b.ImgSq48"),
			"Fixed square image, of 48x48px",

			e.Code("b.ImgSq64"),
			"Fixed square image, of 64x64px",

			e.Code("b.ImgSq96"),
			"Fixed square image, of 96x96px",

			e.Code("b.ImgSq128"),
			"Fixed square image, of 128x128px",

			e.Code("b.ImgSquare"),
			"Force image ratio to square",

			e.Code("b.Img1By1"),
			"Force image ratio to 1 by 1",

			e.Code("b.Img5By4"),
			"Force image ratio to 5 by 4",

			e.Code("b.Img4By3"),
			"Force image ratio to 4 by 3",

			e.Code("b.Img3By2"),
			"Force image ratio to 3 by 2",

			e.Code("b.Img5By3"),
			"Force image ratio to 5 by 3",

			e.Code("b.Img16By9"),
			"Force image ratio to 16 by 9",

			e.Code("b.Img2By1"),
			"Force image ratio to 2 by 1",

			e.Code("b.Img3By1"),
			"Force image ratio to 3 by 1",

			e.Code("b.Img4By5"),
			"Force image ratio to 4 by 5",

			e.Code("b.Img3By4"),
			"Force image ratio to 3 by 4",

			e.Code("b.Img2By3"),
			"Force image ratio to 2 by 3",

			e.Code("b.Img3By5"),
			"Force image ratio to 3 by 5",

			e.Code("b.Img9By16"),
			"Force image ratio to 9 by 16",

			e.Code("b.Img1By2"),
			"Force image ratio to 1 by 2",

			e.Code("b.Img1By3"),
			"Force image ratio to 1 by 3",

			e.Code("b.FullWidth"),
			"Make the image take the whole width of its container",
		),
		e.P(
			"Use ", e.Code("e.ImgSrc"), " to create an <img> e.Element with the provided URL as its src attribute. Apply ", e.Code("b.Rounded"), " to the inner image to make it rounded, associated with an ", e.Code("b.Img*By*"), " modifier on the ", e.Code("b.Image"), " e.Element. Apply ", e.Code("b.Ratio"), " to an inner e.Element to apply the parent ratio to that e.Element.",
		),
		e.P(
			"The ", e.Code("b.ImageImg"), " constructor creates a Bulma image e.Element with an inner img. It accepts the same values as ", e.Code("b.Image"), ". The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnImg(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<img>"), " e.Element"},

			e.Code("b.OnFigure(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<figure>"), " e.Element"},

			e.Code("b.Rounded"),
			[]any{"Make the image rounded (class applied to the ", e.Code("<img>"), " e.Element)"},

			e.Code("b.ImgAlt"),
			"Define the image alt text",
		),
		e.P(
			"Other children are added to the ", e.Code("<figure>"), " e.Element.",
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
	"Arbitrary ratios with any e.Element",
	"https://bulma.io/documentation/elements/image/#arbitrary-ratios-with-any-e.Element",
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
