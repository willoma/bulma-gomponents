package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var formCheckbox = c.NewPage(
	"Checkbox", "Checkbox", "/form/checkbox",
	"https://bulma.io/documentation/form/checkbox/",
	c.Example(
		`b.Checkbox("Remember me")`,
		b.Checkbox("Remember me"),
	),
	c.Example(
		`b.Checkbox("I agree to the ", b.AHref("#", "terms and conditions"))`,
		b.Checkbox("I agree to the ", b.AHref("#", "terms and conditions")),
	),
	c.Example(
		`b.CheckboxDisabled("Save my preferences")`,
		b.CheckboxDisabled("Save my preferences"),
	),
)
