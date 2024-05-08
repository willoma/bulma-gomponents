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
		e.P("The ", e.Code("b.Level"), " constructor creates a level."),
		c.Children(
			c.Row("b.LevelLeft(...any)", "Apply children to the ", e.Code(`<div class="level-left">`), " element"),
			c.Row("b.LevelRight(...any)", "Apply children to the ", e.Code(`<div class="level-right">`), " element"),
			c.Row("e.Element", "Add the ", e.Code("level-item"), " class to the element and apply the element to the ", e.Code(`<nav class="level">`), " element"),
			c.Row("string", "Wrap text into an ", e.Code("<p>"), " element with class ", e.Code("level-item"), " and apply it to the ", e.Code(`<nav class="level">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<nav class="level">`), " element"),
			c.RowNodeElement("Wrap element into a ", e.Code("<div>"), " element with class ", e.Code("level-item"), " and apply it to the ", e.Code(`<nav class="level">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<nav class="level">`), " element"),
		),
		e.P("The ", e.Code("b.LevelLeft"), " and ", e.Code("b.LevelRight"), " functions create content for the level left or right sections."),
		c.Children(
			c.Row("e.Element", "Add the ", e.Code("level-item"), " class to the element"),
			c.Row("string", "Wrap text into an ", e.Code("<p>"), " element with class ", e.Code("level-item")),
			c.RowNodeAttribute("Apply attribute to the level section"),
			c.RowNodeElement("Wrap element into a ", e.Code("<div>"), " element with class ", e.Code("level-item")),
		),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/level/",
	c.HorizontalExample(
		`b.Level(
	b.LevelLeft(
		e.Div(
			b.Subtitle(5, e.Strong("123"), " posts"),
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
					b.Subtitle(5, e.Strong("123"), " posts"),
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
