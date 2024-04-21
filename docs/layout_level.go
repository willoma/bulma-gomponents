package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var level = c.NewPage(
	"Level", "Level", "/level",
	"",

	b.Content(
		el.P("The ", el.Code("b.Level"), " constructor creates a level. The following children have a special meaning:"),
		b.DList(
			el.Code("b.LevelLeft(...)"),
			[]any{"Add children to the left part of the level"},

			el.Code("b.LevelRight(...)"),
			[]any{"Add children to the right part of the level"},

			el.Code("b.Element"),
			[]any{"Add the ", el.Code("level-item"), " class to the element and apply the element to the level"},

			el.Code("string"),
			[]any{"Wrap the string into a ", el.Code("b.Element"), " div with class ", el.Code("level-item"), " and apply it to the level"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the level",

			[]any{"Other ", el.Code("gomponents.Node")},
			[]any{"Wrap the element into a ", el.Code("b.Element"), " div with class ", el.Code("level-item"), " and apply it to the level"},
		),
		el.P("Children in ", el.Code("b.LevelLeft"), " and ", el.Code("b.LevelRight"), " have the following rules:"),
		b.DList(

			el.Code("b.Element"),
			[]any{"Add the ", el.Code("level-item"), " class to the element and apply the element to the level section"},

			el.Code("string"),
			[]any{"Wrap the string into a ", el.Code("b.Element"), " div with class ", el.Code("level-item"), " and apply it to the level section"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the level section",

			[]any{"Other ", el.Code("gomponents.Node")},
			[]any{"Wrap the element into a ", el.Code("b.Element"), " div with class ", el.Code("level-item"), " and apply it to the level section"},
		),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/level/",
	c.HorizontalExample(
		`b.Level(
	b.LevelLeft(
		el.Div(
			b.Subtitle5(el.Strong("123"), " posts"),
		),
		el.Div(
			b.Field(
				b.Addons,
				b.Control(b.InputText(html.Placeholder("Find a post"))),
				b.Control(b.Button("Search")),
			),
		),
	),
	b.LevelRight(
		el.P(el.Strong("All")),
		el.P(el.A("Published")),
		el.P(el.A("Drafts")),
		el.P(el.A("Deleted")),
		el.P(b.ButtonA(b.Success, "New")),
	),
)`,
		b.Level(
			b.LevelLeft(
				el.Div(
					b.Subtitle5(el.Strong("123"), " posts"),
				),
				el.Div(
					b.Field(
						b.Addons,
						b.Control(b.InputText(html.Placeholder("Find a post"))),
						b.Control(b.Button("Search")),
					),
				),
			),
			b.LevelRight(
				el.P(el.Strong("All")),
				el.P(el.A("Published")),
				el.P(el.A("Drafts")),
				el.P(el.A("Deleted")),
				el.P(b.ButtonA(b.Success, "New")),
			),
		),
	),
).Subsection(
	"Centered level",
	"https://bulma.io/documentation/layout/level/#centered-level",
	c.HorizontalExample(
		`b.Level(
	el.Div(
		b.TextCentered,
		el.Div(el.P("Tweets"), b.Title(html.P, "3,456")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Following"), b.Title(html.P, "123")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Followers"), b.Title(html.P, "456K")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Likes"), b.Title(html.P, "789")),
	),
)`,
		b.Level(
			el.Div(
				b.TextCentered,
				el.Div(el.P("Tweets"), b.Title(html.P, "3,456")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Following"), b.Title(html.P, "123")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Followers"), b.Title(html.P, "456K")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Likes"), b.Title(html.P, "789")),
			),
		),
	),
	c.HorizontalExample(
		`b.Level(
	el.P(b.TextCentered, el.A("Home")),
	el.P(b.TextCentered, el.A("Menu")),
	el.P(b.ImgSrc("https://bulma.io/assets/images/bulma-type.png", html.Alt(""), b.Style("height", "30px"))),
	el.P(b.TextCentered, el.A("Reservations")),
	el.P(b.TextCentered, el.A("Contact")),
)`,
		b.Level(
			el.P(b.TextCentered, el.A("Home")),
			el.P(b.TextCentered, el.A("Menu")),
			el.P(b.ImgSrc("https://bulma.io/assets/images/bulma-type.png", html.Alt(""), b.Style("height", "30px"))),
			el.P(b.TextCentered, el.A("Reservations")),
			el.P(b.TextCentered, el.A("Contact")),
		),
	),
).Subsection(
	"Mobile level",
	"https://bulma.io/documentation/layout/level/#mobile-level",
	c.HorizontalExample(
		`b.Level(
	b.Mobile,
	el.Div(
		b.TextCentered,
		el.Div(el.P("Tweets"), b.Title(html.P, "3,456")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Following"), b.Title(html.P, "123")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Followers"), b.Title(html.P, "456K")),
	),
	el.Div(
		b.TextCentered,
		el.Div(el.P("Likes"), b.Title(html.P, "789")),
	),
)`,
		b.Level(
			b.Mobile,
			el.Div(
				b.TextCentered,
				el.Div(el.P("Tweets"), b.Title(html.P, "3,456")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Following"), b.Title(html.P, "123")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Followers"), b.Title(html.P, "456K")),
			),
			el.Div(
				b.TextCentered,
				el.Div(el.P("Likes"), b.Title(html.P, "789")),
			),
		),
	),
)
