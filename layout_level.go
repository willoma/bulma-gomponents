package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var level = c.NewPage(
	"Level", "Level", "/level",
	"",

	b.Content(
		e.P("The ", e.Code("b.Level"), " constructor creates a level. The following children have a special meaning:"),
		b.DList(
			e.Code("b.LevelLeft(...)"),
			[]any{"Add children to the left part of the level"},

			e.Code("b.LevelRight(...)"),
			[]any{"Add children to the right part of the level"},

			e.Code("e.Element"),
			[]any{"Add the ", e.Code("level-item"), " class to the element and apply the element to the level"},

			e.Code("string"),
			[]any{"Wrap the string into a ", e.Code("e.Element"), " div with class ", e.Code("level-item"), " and apply it to the level"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the level",

			[]any{"Other ", e.Code("gomponents.Node")},
			[]any{"Wrap the element into a ", e.Code("e.Element"), " div with class ", e.Code("level-item"), " and apply it to the level"},
		),
		e.P("Children in ", e.Code("b.LevelLeft"), " and ", e.Code("b.LevelRight"), " have the following rules:"),
		b.DList(

			e.Code("e.Element"),
			[]any{"Add the ", e.Code("level-item"), " class to the element and apply the element to the level section"},

			e.Code("string"),
			[]any{"Wrap the string into a ", e.Code("e.Element"), " div with class ", e.Code("level-item"), " and apply it to the level section"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the level section",

			[]any{"Other ", e.Code("gomponents.Node")},
			[]any{"Wrap the element into a ", e.Code("e.Element"), " div with class ", e.Code("level-item"), " and apply it to the level section"},
		),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/level/",
	c.HorizontalExample(
		`b.Level(
	b.LevelLeft(
		e.Div(
			b.Subtitle5(e.Strong("123"), " posts"),
		),
		e.Div(
			b.Field(
				b.Addons,
				b.Control(b.InputText(e.Placeholder("Find a post"))),
				b.Control(b.Button("Search")),
			),
		),
	),
	b.LevelRight(
		e.P(e.Strong("All")),
		e.P(e.A("Published")),
		e.P(e.A("Drafts")),
		e.P(e.A("Deleted")),
		e.P(b.ButtonA(b.Success, "New")),
	),
)`,
		b.Level(
			b.LevelLeft(
				e.Div(
					b.Subtitle5(e.Strong("123"), " posts"),
				),
				e.Div(
					b.Field(
						b.Addons,
						b.Control(b.InputText(e.Placeholder("Find a post"))),
						b.Control(b.Button("Search")),
					),
				),
			),
			b.LevelRight(
				e.P(e.Strong("All")),
				e.P(e.A("Published")),
				e.P(e.A("Drafts")),
				e.P(e.A("Deleted")),
				e.P(b.ButtonA(b.Success, "New")),
			),
		),
	),
).Subsection(
	"Centered level",
	"https://bulma.io/documentation/layout/level/#centered-level",
	c.HorizontalExample(
		`b.Level(
	e.Div(
		b.TextCentered,
		e.Div(e.P("Tweets"), b.Title(html.P, "3,456")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Following"), b.Title(html.P, "123")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Followers"), b.Title(html.P, "456K")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Likes"), b.Title(html.P, "789")),
	),
)`,
		b.Level(
			e.Div(
				b.TextCentered,
				e.Div(e.P("Tweets"), b.Title(html.P, "3,456")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Following"), b.Title(html.P, "123")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Followers"), b.Title(html.P, "456K")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Likes"), b.Title(html.P, "789")),
			),
		),
	),
	c.HorizontalExample(
		`b.Level(
	e.P(b.TextCentered, e.A("Home")),
	e.P(b.TextCentered, e.A("Menu")),
	e.P(e.ImgSrc("https://bulma.io/assets/images/bulma-type.png", e.Alt(""), e.Styles{"height": "30px"})),
	e.P(b.TextCentered, e.A("Reservations")),
	e.P(b.TextCentered, e.A("Contact")),
)`,
		b.Level(
			e.P(b.TextCentered, e.A("Home")),
			e.P(b.TextCentered, e.A("Menu")),
			e.P(e.ImgSrc("https://bulma.io/assets/images/bulma-type.png", e.Alt(""), e.Styles{"height": "30px"})),
			e.P(b.TextCentered, e.A("Reservations")),
			e.P(b.TextCentered, e.A("Contact")),
		),
	),
).Subsection(
	"Mobile level",
	"https://bulma.io/documentation/layout/level/#mobile-level",
	c.HorizontalExample(
		`b.Level(
	b.Mobile,
	e.Div(
		b.TextCentered,
		e.Div(e.P("Tweets"), b.Title(html.P, "3,456")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Following"), b.Title(html.P, "123")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Followers"), b.Title(html.P, "456K")),
	),
	e.Div(
		b.TextCentered,
		e.Div(e.P("Likes"), b.Title(html.P, "789")),
	),
)`,
		b.Level(
			b.Mobile,
			e.Div(
				b.TextCentered,
				e.Div(e.P("Tweets"), b.Title(html.P, "3,456")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Following"), b.Title(html.P, "123")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Followers"), b.Title(html.P, "456K")),
			),
			e.Div(
				b.TextCentered,
				e.Div(e.P("Likes"), b.Title(html.P, "789")),
			),
		),
	),
)
