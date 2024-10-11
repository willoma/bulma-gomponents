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
		e.P("The ", e.Code("b.Radio"), " constructor creates a radio input, together with its label container."),
		c.Modifiers(
			c.Row("b.Disabled", "Disable the radio button"),
		),
		c.Children(
			c.Row("b.OnInput(...any)", "Apply children to the ", e.Code("<input>"), " element"),
			c.Row("b.OnLabel(...any)", "Apply children to the ", e.Code("<label>"), " element"),
			c.Row("string", "Add to the ", e.Code("<label>"), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code("<input>"), " element"),
			c.RowNodeElement("Add element to the ", e.Code("<label>"), " element"),
			c.Row("e.Element", "Add element to the ", e.Code("<label>"), " element"),
			c.RowDefault("Apply child to the ", e.Code("<input>"), " element"),
		),
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
).Subsection(
	"List of Radio Buttons", "https://bulma.io/documentation/form/radio/#list-of-radio-buttons",
	c.Example(
		`b.Radios(
	b.Radio(e.Name("rsvp"), "Going"),
	b.Radio(e.Name("rsvp"), "Not going"),
	b.Radio(b.Disabled, e.Name("rsvp"), "Maybe"),
)`,
		b.Radios(
			b.Radio(e.Name("rsvp"), "Going"),
			b.Radio(e.Name("rsvp"), "Not going"),
			b.Radio(b.Disabled, e.Name("rsvp"), "Maybe"),
		),
	),
)
