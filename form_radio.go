package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
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
	b.Radio(e.Name("answer"), "Yes"),
	b.Radio(e.Name("answer"), "No"),
)`,
		b.Control(
			b.Radio(e.Name("answer"), "Yes"),
			b.Radio(e.Name("answer"), "No"),
		),
	),
	c.Example(
		`b.Control(
	b.Radio(e.Name("foobar"), "Foo"),
	b.Radio(e.Name("foobar"), "Bar", b.Checked),
)`,
		b.Control(
			b.Radio(e.Name("foobar"), "Foo"),
			b.Radio(e.Name("foobar"), "Bar", b.Checked),
		),
	),
	c.Example(
		`b.Control(
	b.Radio(e.Name("rsvp"), "Going"),
	b.Radio(e.Name("rsvp"), "Not going"),
	b.Radio(b.Disabled, e.Name("rsvp"), "Maybe"),
)`,
		b.Control(
			b.Radio(e.Name("rsvp"), "Going"),
			b.Radio(e.Name("rsvp"), "Not going"),
			b.Radio(b.Disabled, e.Name("rsvp"), "Maybe"),
		),
	),
)
