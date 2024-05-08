package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var formCheckbox = c.NewPage(
	"Checkbox", "Checkbox", "/form/checkbox",
	"",

	b.Content(
		e.P("The ", e.Code("b.Checkbox"), " constructor creates a checkbox input, together with its label container."),
		c.Modifiers(
			c.Row("b.Disabled", "Disable the checkbox input"),
		),
		c.Children(
			c.Row("b.OnInput(...any)", "Apply children to the ", e.Code("<input>"), " element"),
			c.Row("b.OnLabel(...aby)", "Apply children to the ", e.Code("<label>"), " element"),
			c.Row("string", "Add to the ", e.Code("<label>"), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code("<input>"), " element"),
			c.RowNodeElement("Add element to the ", e.Code("<label>"), " element"),
			c.Row("e.Element", "Add element to the ", e.Code("<label>"), " element"),
			c.RowDefault("Apply child to the ", e.Code("<input>"), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/checkbox/",
	c.Example(
		`b.Checkbox("Remember me")`,
		b.Checkbox("Remember me"),
	),
	c.Example(
		`b.Checkbox("I agree to the ", e.AHref("#", "terms and conditions"))`,
		b.Checkbox("I agree to the ", e.AHref("#", "terms and conditions")),
	),
	c.Example(
		`b.Checkbox(b.Disabled, "Save my preferences")`,
		b.Checkbox(b.Disabled, "Save my preferences"),
	),
)
