package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var block = c.NewPage(
	"Block", "Block", "/block",
	"https://bulma.io/documentation/elements/block/",

	c.Example(
		`b.Block("This text is within a ", el.Strong("block"), "."),
b.Block("This text is within a ", el.Strong("second block"), ". Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
b.Block("This text is within a ", el.Strong("third block"), ". This block has no margin at the bottom.")`,
		b.Block("This text is within a ", el.Strong("block"), "."),
		b.Block("This text is within a ", el.Strong("second block"), ". Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
		b.Block("This text is within a ", el.Strong("third block"), ". This block has no margin at the bottom."),
	),
	c.Example(
		`el.Div("This text is ", el.Em("not"), " within a ", el.Strong("block"), "."),
el.Div("This text ", el.Em("isn't"), " within a ", el.Strong("block"), " either. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
el.Div("This text is also ", el.Em("not"), " within a ", el.Strong("block"), ".")`,
		el.Div("This text is ", el.Em("not"), " within a ", el.Strong("block"), "."),
		el.Div("This text ", el.Em("isn't"), " within a ", el.Strong("block"), " either. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis."),
		el.Div("This text is also ", el.Em("not"), " within a ", el.Strong("block"), "."),
	),
)
