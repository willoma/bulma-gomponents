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
		e.P("The ", e.Code("b.Checkbox"), " constructor creates a checkbox input. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<input>"), " e.Element"},

			e.Code("b.OnLabel(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<label>"), " e.Element"},

			e.Code("string"),
			"Add as a child to the label",

			e.Code("b.Disabled"),
			"Disable the checkbox input",

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
