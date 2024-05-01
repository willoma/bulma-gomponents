package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var formRadio = c.NewPage(
	"Radio", "Radio button", "/form/radio",
	"",

	b.Content(
		e.P("The ", e.Code("b.Radio"), " constructor creates a radio e.Element, together with its label container. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<input>"), " e.Element"},

			e.Code("b.OnLabel(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<label>"), " e.Element"},

			e.Code("string"),
			"Add as a child to the label",

			e.Code("b.Disabled"),
			"Disable the radio button",

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the input",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the label",

			e.Code("e.Element"),
			"Add this e.Element to the label",
		),
		e.P("Other children are added to the ", e.Code("<input>"), " e.Element."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/radio/",
	c.Example(
		`b.Control(
	b.Radio(html.Name("answer"), "Yes"),
	b.Radio(html.Name("answer"), "No"),
)`,
		b.Control(
			b.Radio(html.Name("answer"), "Yes"),
			b.Radio(html.Name("answer"), "No"),
		),
	),
	c.Example(
		`b.Control(
	b.Radio(html.Name("foobar"), "Foo"),
	b.Radio(html.Name("foobar"), "Bar", b.Checked),
)`,
		b.Control(
			b.Radio(html.Name("foobar"), "Foo"),
			b.Radio(html.Name("foobar"), "Bar", b.Checked),
		),
	),
	c.Example(
		`b.Control(
	b.Radio(html.Name("rsvp"), "Going"),
	b.Radio(html.Name("rsvp"), "Not going"),
	b.Radio(b.Disabled, html.Name("rsvp"), "Maybe"),
)`,
		b.Control(
			b.Radio(html.Name("rsvp"), "Going"),
			b.Radio(html.Name("rsvp"), "Not going"),
			b.Radio(b.Disabled, html.Name("rsvp"), "Maybe"),
		),
	),
)
