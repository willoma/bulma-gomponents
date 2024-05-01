package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var skeletons = c.NewPage(
	"Skeletons", "Skeletons in bulma", "/skeletons",
	"",

	b.Content(
		e.P("The ", e.Code("b.SkeletonBlock"), " constructor creates a block skeleton."),
		e.P("The ", e.Code("b.SkeletonLines"), " constructor creates a block of skeleton lines. It takes the number of lines as its only argument"),
		e.P("Use the ", e.Code("b.Skeleton"), " modifier on any supported component to make it a skeleton. Use ", e.Code("b.HasSkeleton"), " for the ", e.Code("has-skeleton"), " class."),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/features/skeletons/",
).Subsection(
	"Skeleton block",
	"https://bulma.io/documentation/features/skeletons/#skeleton-block",

	c.Example(
		`b.SkeletonBlock()`,
		b.SkeletonBlock(),
	),

	c.Example(
		`b.SkeletonBlock(
	"Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Donec sed odio dui. Nullam quis risus eget urna mollis ornare vel eu leo. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
)`,
		b.SkeletonBlock(
			"Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Donec sed odio dui. Nullam quis risus eget urna mollis ornare vel eu leo. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		),
	),
).Subsection(
	"Skeleton lines",
	"https://bulma.io/documentation/features/skeletons/#skeleton-lines",

	c.Example(
		`b.SkeletonLines(5)`,
		b.SkeletonLines(5),
	),
).Subsection(
	"Bulma components with skeletons",
	"https://bulma.io/documentation/features/skeletons/#bulma-components-with-skeletons",

	b.Title4("Button"),
	c.Example(
		`b.Buttons(
	b.Button(b.Skeleton, "Button"),
	b.Button(b.Link, b.Skeleton, "Link"),
	b.Button(b.Primary, b.Skeleton, "Primary"),
	b.Button(b.Success, b.Skeleton, "Success"),
	b.Button(b.Info, b.Skeleton, "Info"),
	b.Button(b.Warning, b.Skeleton, "Warning"),
	b.Button(b.Danger, b.Skeleton, "Danger"),
)`,
		b.Buttons(
			b.Button(b.Skeleton, "Button"),
			b.Button(b.Link, b.Skeleton, "Link"),
			b.Button(b.Primary, b.Skeleton, "Primary"),
			b.Button(b.Success, b.Skeleton, "Success"),
			b.Button(b.Info, b.Skeleton, "Info"),
			b.Button(b.Warning, b.Skeleton, "Warning"),
			b.Button(b.Danger, b.Skeleton, "Danger"),
		),
	),

	b.Title4("Icon"),
	c.Example(
		`fa.Icon(fa.Solid, "reply", b.Skeleton)`,
		fa.Icon(fa.Solid, "reply", b.Skeleton),
	),

	b.Title4("Image"),
	c.Example(
		`b.ImageImg("https://placehold.co/16x16", b.ImgSq16, b.Skeleton),
b.ImageImg("https://placehold.co/32x32", b.ImgSq32, b.Skeleton),
b.ImageImg("https://placehold.co/64x64", b.ImgSq64, b.Skeleton),
b.ImageImg("https://placehold.co/128x128", b.ImgSq128, b.Skeleton)`,
		b.ImageImg("https://placehold.co/16x16", b.ImgSq16, b.Skeleton),
		b.ImageImg("https://placehold.co/32x32", b.ImgSq32, b.Skeleton),
		b.ImageImg("https://placehold.co/64x64", b.ImgSq64, b.Skeleton),
		b.ImageImg("https://placehold.co/128x128", b.ImgSq128, b.Skeleton),
	),

	b.Title4("Media object"),
	c.Example(
		`b.Media(
	b.MediaLeft(
		b.ImageImg("https://placehold.co/128x128", html.P, b.ImgSq64, b.Skeleton),
	),
	b.Content(
		b.Skeleton,
		e.P(
			e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
			e.Br(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
		),
	),
	b.Level(
		b.Mobile,
		b.LevelLeft(
			e.A(fa.Icon(fa.Solid, "reply", b.Small, b.Skeleton)),
			e.A(fa.Icon(fa.Solid, "retweet", b.Small, b.Skeleton)),
			e.A(fa.Icon(fa.Solid, "heart", b.Small, b.Skeleton)),
		),
	),
	b.MediaRight(
		b.Delete(e.AriaLabel("delete"), b.Skeleton),
	),
)`,
		b.Media(
			b.MediaLeft(
				b.ImageImg("https://placehold.co/128x128", html.P, b.ImgSq64, b.Skeleton),
			),
			b.Content(
				b.Skeleton,
				e.P(
					e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"),
					e.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.",
				),
			),
			b.Level(
				b.Mobile,
				b.LevelLeft(
					e.A(fa.Icon(fa.Solid, "reply", b.Small, b.Skeleton)),
					e.A(fa.Icon(fa.Solid, "retweet", b.Small, b.Skeleton)),
					e.A(fa.Icon(fa.Solid, "heart", b.Small, b.Skeleton)),
				),
			),
			b.MediaRight(
				b.Delete(e.AriaLabel("delete"), b.Skeleton),
			),
		),
	),
	c.Example(
		`b.Media(
	b.MediaLeft(b.ImageImg("https://placehold.co/128x128", html.P, b.ImgSq64, b.Skeleton)),
	b.Field(
		b.Control(
			html.P,
			b.Textarea(b.Skeleton, html.Placeholder("Add a comment...")),
		),
	),
	b.Level(
		b.LevelLeft(
			e.Div(
				b.ButtonA(b.Info, b.Skeleton, "Submit"),
			),
		),
		b.LevelRight(
			e.Div(
				b.Checkbox(
					b.OnLabel(b.Skeleton),
					"Press enter to submit",
				),
			),
		),
	),
)`,
		b.Media(
			b.MediaLeft(b.ImageImg("https://placehold.co/128x128", html.P, b.ImgSq64, b.Skeleton)),
			b.Field(
				b.Control(
					html.P,
					b.Textarea(b.Skeleton, html.Placeholder("Add a comment...")),
				),
			),
			b.Level(
				b.LevelLeft(
					e.Div(
						b.ButtonA(b.Info, b.Skeleton, "Submit"),
					),
				),
				b.LevelRight(
					e.Div(
						b.Checkbox(
							b.OnLabel(b.Skeleton),
							"Press enter to submit",
						),
					),
				),
			),
		),
	),

	b.Title4("Notification"),
	c.Example(
		`b.Notification(
	b.Skeleton,
	"Curabitur blandit tempus porttitor. Etiam porta sem malesuada magna mollis euismod. Cras justo odio, dapibus ac facilisis in, egestas eget quam.",
)`,
		b.Notification(
			b.Skeleton,
			"Curabitur blandit tempus porttitor. Etiam porta sem malesuada magna mollis euismod. Cras justo odio, dapibus ac facilisis in, egestas eget quam.",
		),
	),

	b.Title4("Tag"),
	c.Example(
		`b.Tag(b.Skeleton, "Tag"), " ",
b.Tag(b.Link, b.Skeleton, "Link"), " ",
b.Tag(b.Primary, b.Skeleton, "Primary"), " ",
b.Tag(b.Info, b.Skeleton, "Info"), " ",
b.Tag(b.Success, b.Skeleton, "Success"), " ",
b.Tag(b.Warning, b.Skeleton, "Warning"), " ",
b.Tag(b.Danger, b.Skeleton, "Danger")`,
		b.Tag(b.Skeleton, "Tag"), " ",
		b.Tag(b.Link, b.Skeleton, "Link"), " ",
		b.Tag(b.Primary, b.Skeleton, "Primary"), " ",
		b.Tag(b.Info, b.Skeleton, "Info"), " ",
		b.Tag(b.Success, b.Skeleton, "Success"), " ",
		b.Tag(b.Warning, b.Skeleton, "Warning"), " ",
		b.Tag(b.Danger, b.Skeleton, "Danger"),
	),

	b.Title4("Title and subtitle"),
	c.Example(
		`b.Title(b.Skeleton, "Title")`,
		b.Title(b.Skeleton, "Title"),
	),
	c.Example(
		`b.Title(b.HasSkeleton, "Title")`,
		b.Title(b.HasSkeleton, "Title"),
	),
	c.Example(
		`b.Subtitle(b.Skeleton, "Subtitle")`,
		b.Subtitle(b.Skeleton, "Subtitle"),
	),
	c.Example(
		`b.Subtitle(b.HasSkeleton, "Subtitle")`,
		b.Subtitle(b.HasSkeleton, "Subtitle"),
	),
	c.Example(
		`b.Title(b.Skeleton, b.Spaced, "Title"),
b.Subtitle(b.Skeleton, "Subtitle")`,
		b.Title(b.Skeleton, b.Spaced, "Title"),
		b.Subtitle(b.Skeleton, "Subtitle"),
	),
	c.Example(
		`b.Title(b.HasSkeleton, b.Spaced, "Title"),
b.Subtitle(b.HasSkeleton, "Subtitle")`,
		b.Title(b.HasSkeleton, b.Spaced, "Title"),
		b.Subtitle(b.HasSkeleton, "Subtitle"),
	),

	b.Title4("Form control"),
	c.Example(
		`b.InputText(b.Skeleton)`,
		b.InputText(b.Skeleton),
	),
	c.Example(
		`b.Textarea(b.Skeleton)`,
		b.Textarea(b.Skeleton),
	),
)
