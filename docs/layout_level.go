package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var level = c.NewPage(
	"Level", "Level", "/level",
	"https://bulma.io/documentation/layout/level/",
	c.HorizontalExample(
		`b.Level(
	b.LevelLeft(
		b.Subtitle5(el.Strong("123"), " posts"),
		b.Field(
			b.Addons,
			b.Control(
				b.InputText(html.Placeholder("Find a post")),
			),
			b.Control(
				b.Button("Search"),
			),
		),
	),
	b.LevelRight(
		b.LevelItem(html.P, el.Strong("All")),
		b.LevelItem(html.P, el.A("Published")),
		b.LevelItem(html.P, el.A("Drafts")),
		b.LevelItem(html.P, el.A("Deleted")),
		b.LevelItem(html.P, b.ButtonA(b.Success, "New")),
	),
)`,
		b.Level(
			b.LevelLeft(
				b.Subtitle5(el.Strong("123"), " posts"),
				b.Field(
					b.Addons,
					b.Control(
						b.InputText(html.Placeholder("Find a post")),
					),
					b.Control(
						b.Button("Search"),
					),
				),
			),
			b.LevelRight(
				b.LevelItem(html.P, el.Strong("All")),
				b.LevelItem(html.P, el.A("Published")),
				b.LevelItem(html.P, el.A("Drafts")),
				b.LevelItem(html.P, el.A("Deleted")),
				b.LevelItem(html.P, b.ButtonA(b.Success, "New")),
			),
		),
	),
).Section(
	"Centered level",
	"https://bulma.io/documentation/layout/level/#centered-level",
	c.HorizontalExample(
		`b.Level(
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Tweets"),
			b.Title(html.P, "3,456"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Following"),
			b.Title(html.P, "123"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Followers"),
			b.Title(html.P, "456K"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Likes"),
			b.Title(html.P, "789"),
		),
	),
)`,
		b.Level(
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Tweets"),
					b.Title(html.P, "3,456"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Following"),
					b.Title(html.P, "123"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Followers"),
					b.Title(html.P, "456K"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Likes"),
					b.Title(html.P, "789"),
				),
			),
		),
	),
	c.HorizontalExample(
		`b.Level(
	b.LevelItem(
		html.P,
		b.TextCentered,
		el.A("Home"),
	),
	b.LevelItem(
		html.P,
		b.TextCentered,
		el.A("Menu"),
	),
	b.LevelItem(
		html.P,
		b.ImgSrc("https://bulma.io/images/bulma-type.png", html.Alt(""), b.Style("height", "30px")),
	),
	b.LevelItem(
		html.P,
		b.TextCentered,
		el.A("Reservations"),
	),
	b.LevelItem(
		html.P,
		b.TextCentered,
		el.A("Contact"),
	),
)`,
		b.Level(
			b.LevelItem(
				html.P,
				b.TextCentered,
				el.A("Home"),
			),
			b.LevelItem(
				html.P,
				b.TextCentered,
				el.A("Menu"),
			),
			b.LevelItem(
				html.P,
				b.ImgSrc("https://bulma.io/images/bulma-type.png", html.Alt(""), b.Style("height", "30px")),
			),
			b.LevelItem(
				html.P,
				b.TextCentered,
				el.A("Reservations"),
			),
			b.LevelItem(
				html.P,
				b.TextCentered,
				el.A("Contact"),
			),
		),
	),
).Section(
	"Mobile level",
	"https://bulma.io/documentation/layout/level/#mobile-level",
	c.HorizontalExample(
		`b.Level(
	b.Mobile,
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Tweets"),
			b.Title(html.P, "3,456"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Following"),
			b.Title(html.P, "123"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Followers"),
			b.Title(html.P, "456K"),
		),
	),
	b.LevelItem(
		b.TextCentered,
		el.Div(
			el.P(b.Heading, "Likes"),
			b.Title(html.P, "789"),
		),
	),
)`,
		b.Level(
			b.Mobile,
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Tweets"),
					b.Title(html.P, "3,456"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Following"),
					b.Title(html.P, "123"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Followers"),
					b.Title(html.P, "456K"),
				),
			),
			b.LevelItem(
				b.TextCentered,
				el.Div(
					el.P(b.Heading, "Likes"),
					b.Title(html.P, "789"),
				),
			),
		),
	),
)
