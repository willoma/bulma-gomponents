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
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Image"), " constructor creates a Bulma image element. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.ImgSq16"),
			"Fixed square image, of 16x16px",

			el.Code("b.ImgSq24"),
			"Fixed square image, of 24x24px",

			el.Code("b.ImgSq32"),
			"Fixed square image, of 32x32px",

			el.Code("b.ImgSq48"),
			"Fixed square image, of 48x48px",

			el.Code("b.ImgSq64"),
			"Fixed square image, of 64x64px",

			el.Code("b.ImgSq96"),
			"Fixed square image, of 96x96px",

			el.Code("b.ImgSq128"),
			"Fixed square image, of 128x128px",

			el.Code("b.ImgSquare"),
			"Force image ratio to square",

			el.Code("b.Img1By1"),
			"Force image ratio to 1 by 1",

			el.Code("b.Img5By4"),
			"Force image ratio to 5 by 4",

			el.Code("b.Img4By3"),
			"Force image ratio to 4 by 3",

			el.Code("b.Img3By2"),
			"Force image ratio to 3 by 2",

			el.Code("b.Img5By3"),
			"Force image ratio to 5 by 3",

			el.Code("b.Img16By9"),
			"Force image ratio to 16 by 9",

			el.Code("b.Img2By1"),
			"Force image ratio to 2 by 1",

			el.Code("b.Img3By1"),
			"Force image ratio to 3 by 1",

			el.Code("b.Img4By5"),
			"Force image ratio to 4 by 5",

			el.Code("b.Img3By4"),
			"Force image ratio to 3 by 4",

			el.Code("b.Img2By3"),
			"Force image ratio to 2 by 3",

			el.Code("b.Img3By5"),
			"Force image ratio to 3 by 5",

			el.Code("b.Img9By16"),
			"Force image ratio to 9 by 16",

			el.Code("b.Img1By2"),
			"Force image ratio to 1 by 2",

			el.Code("b.Img1By3"),
			"Force image ratio to 1 by 3",

			el.Code("b.FullWidth"),
			"Make the image take the whole width of its container",
		),
		el.P(
			"Use ", el.Code("b.ImgSrc"), " to create an <img> element with the provided URL as its src attribute. Apply ", el.Code("b.Rounded"), " to the inner image to make it rounded, associated with an ", el.Code("b.Img*By*"), " modifier on the ", el.Code("b.Image"), " element. Apply ", el.Code("b.Ratio"), " to an inner element to apply the parent ratio to that element.",
		),
		el.P(
			"The ", el.Code("b.ImageImg"), " constructor creates a Bulma image element with an inner img. It accepts the same values as ", el.Code("b.Image"), ". The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnImg(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<img>"), " element"},

			el.Code("b.OnFigure(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<figure>"), " element"},

			el.Code("b.Rounded"),
			[]any{"Make the image rounded (class applied to the ", el.Code("<img>"), " element)"},

			el.Code("b.ImgAlt"),
			"Define the image alt text",
		),
		el.P(
			"Other children are added to the ", el.Code("<figure>"), " element.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/image/",

	c.Example(
		`b.Image(
	b.ImgSq128,
	el.Img(html.Src("https://bulma.io/assets/images/placeholders/128x128.png")),
)`,
		b.Image(
			b.ImgSq128,
			el.Img(html.Src("https://bulma.io/assets/images/placeholders/128x128.png")),
		),
	),
	c.Example(
		`b.Image(
	b.ImgSq128,
	b.ImgSrc("https://bulma.io/assets/images/placeholders/128x128.png"),
)`,
		b.Image(
			b.ImgSq128,
			b.ImgSrc("https://bulma.io/assets/images/placeholders/128x128.png"),
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
		`el.P("16x16px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/16x16.png", b.ImgSq16),
el.P("24x24px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/24x24.png", b.ImgSq24),
el.P("32x32px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/32x32.png", b.ImgSq32),
el.P("48x48px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/48x48.png", b.ImgSq48),
el.P("64x64px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/64x64.png", b.ImgSq64),
el.P("96x96px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/96x96.png", b.ImgSq96),
el.P("128x128px:"),
b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128)`,
		el.P("16x16px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/16x16.png", b.ImgSq16),
		el.P("24x24px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/24x24.png", b.ImgSq24),
		el.P("32x32px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/32x32.png", b.ImgSq32),
		el.P("48x48px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/48x48.png", b.ImgSq48),
		el.P("64x64px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/64x64.png", b.ImgSq64),
		el.P("96x96px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/96x96.png", b.ImgSq96),
		el.P("128x128px:"),
		b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq128),
	),
).Subsection(
	"Rounded images",
	"https://bulma.io/documentation/elements/image/#rounded-images",
	c.Example(
		`b.Image(
			b.ImgSq128,
			b.ImgSrc(
				"https://bulma.io/assets/images/placeholders/128x128.png",
				b.Rounded,
			),
		)`,
		b.Image(
			b.ImgSq128,
			b.ImgSrc(
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
		`el.Div(
	b.Style("width", "10rem"),
	el.P("Square (or 1 by 1):"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x480.png",
		b.ImgSquare,
	),
	el.P("1 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x480.png",
		b.Img1By1,
	),
	el.P("5 by 4:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/600x480.png",
		b.Img5By4,
	),
	el.P("4 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x480.png",
		b.Img4By3,
	),
	el.P("3 by 2:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x320.png",
		b.Img3By2,
	),
	el.P("5 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/800x480.png",
		b.Img5By3,
	),
	el.P("16 by 9:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x360.png",
		b.Img16By9,
	),
	el.P("2 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/640x320.png",
		b.Img2By1,
	),
	el.P("3 by 1:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/720x240.png",
		b.Img3By1,
	),
	el.P("4 by 5:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x600.png",
		b.Img4By5,
	),
	el.P("3 by 4:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x640.png",
		b.Img3By4,
	),
	el.P("2 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/320x480.png",
		b.Img2By3,
	),
	el.P("3 by 5:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/480x800.png",
		b.Img3By5,
	),
	el.P("9 by 16:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/360x640.png",
		b.Img9By16,
	),
	el.P("1 by 2:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/320x640.png",
		b.Img1By2,
	),
	el.P("1 by 3:"),
	b.ImageImg(
		"https://bulma.io/assets/images/placeholders/240x720.png",
		b.Img1By3,
	),
)`,
		el.Div(
			b.Style("width", "10rem"),
			el.P("Square (or 1 by 1):"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x480.png",
				b.ImgSquare,
			),
			el.P("1 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x480.png",
				b.Img1By1,
			),
			el.P("5 by 4:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/600x480.png",
				b.Img5By4,
			),
			el.P("4 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x480.png",
				b.Img4By3,
			),
			el.P("3 by 2:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x320.png",
				b.Img3By2,
			),
			el.P("5 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/800x480.png",
				b.Img5By3,
			),
			el.P("16 by 9:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x360.png",
				b.Img16By9,
			),
			el.P("2 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/640x320.png",
				b.Img2By1,
			),
			el.P("3 by 1:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/720x240.png",
				b.Img3By1,
			),
			el.P("4 by 5:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x600.png",
				b.Img4By5,
			),
			el.P("3 by 4:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x640.png",
				b.Img3By4,
			),
			el.P("2 by 3:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/320x480.png",
				b.Img2By3,
			),
			el.P("3 by 5:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/480x800.png",
				b.Img3By5,
			),
			el.P("9 by 16:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/360x640.png",
				b.Img9By16,
			),
			el.P("1 by 2:"),
			b.ImageImg(
				"https://bulma.io/assets/images/placeholders/320x640.png",
				b.Img1By2,
			),
			el.P("1 by 3:"),
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
