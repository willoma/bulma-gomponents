package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var formRadio = c.NewPage(
	"Radio", "Radio button", "/form/radio",
	"",

	b.Content(
		el.P("The ", el.Code("b.Radio"), " constructor creates a radio element, together with its label container. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<input>"), " element"},

			el.Code("b.OnLabel(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<label>"), " element"},

			el.Code("string"),
			"Add as a child to the label",

			el.Code("b.Disabled"),
			"Disable the radio button",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the input",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the label",

			el.Code("b.Element"),
			"Add this element to the label",
		),
		el.P("Other children are added to the ", el.Code("<input>"), " element."),
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
