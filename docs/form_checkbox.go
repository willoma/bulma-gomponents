package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var formCheckbox = c.NewPage(
	"Checkbox", "Checkbox", "/form/checkbox",
	"",

	b.Content(
		el.P("The ", el.Code("b.Checkbox"), " constructor creates a checkbox input. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<input>"), " element"},

			el.Code("b.OnLabel(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<label>"), " element"},

			el.Code("string"),
			"Add as a child to the label",

			el.Code("b.Disabled"),
			"Disable the checkbox input",

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
	"Bulma examples", "https://bulma.io/documentation/form/checkbox/",
	c.Example(
		`b.Checkbox("Remember me")`,
		b.Checkbox("Remember me"),
	),
	c.Example(
		`b.Checkbox("I agree to the ", b.AHref("#", "terms and conditions"))`,
		b.Checkbox("I agree to the ", b.AHref("#", "terms and conditions")),
	),
	c.Example(
		`b.Checkbox(b.Disabled, "Save my preferences")`,
		b.Checkbox(b.Disabled, "Save my preferences"),
	),
)
