package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var block = c.NewPage(
	"Block", "Block", "/block",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Block"), " constructor returns a container that ensures siblings have a consistent margin.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/block/",
	c.Example(
		`b.Block("This text is within a ", e.Strong("block"), "."),
b.Block("This text is within a ", e.Strong("second block"), ". Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
b.Block("This text is within a ", e.Strong("third block"), ". This block has no margin at the bottom.")`,
		b.Block("This text is within a ", e.Strong("block"), "."),
		b.Block("This text is within a ", e.Strong("second block"), ". Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
		b.Block("This text is within a ", e.Strong("third block"), ". This block has no margin at the bottom."),
	),
	c.Example(
		`e.Div("This text is ", e.Em("not"), " within a ", e.Strong("block"), "."),
e.Div("This text ", e.Em("isn't"), " within a ", e.Strong("block"), " either. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
e.Div("This text is also ", e.Em("not"), " within a ", e.Strong("block"), ".")`,
		e.Div("This text is ", e.Em("not"), " within a ", e.Strong("block"), "."),
		e.Div("This text ", e.Em("isn't"), " within a ", e.Strong("block"), " either. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
		e.Div("This text is also ", e.Em("not"), " within a ", e.Strong("block"), "."),
	),
)
