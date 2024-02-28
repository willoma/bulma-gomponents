package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var formRadio = c.NewPage(
	"Radio", "Radio button", "/form/radio",
	"https://bulma.io/documentation/form/radio/",
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
	b.RadioDisabled(html.Name("rsvp"), "Maybe"),
)`,
		b.Control(
			b.Radio(html.Name("rsvp"), "Going"),
			b.Radio(html.Name("rsvp"), "Not going"),
			b.RadioDisabled(html.Name("rsvp"), "Maybe"),
		),
	),
)
