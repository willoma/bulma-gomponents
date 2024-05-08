package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var intro = c.NewPage(
	"Introduction", "Bulma-Gomponents", "/index",
	"",

	b.Content(
		e.P("This ", e.Em("Go"), " library makes it easier to use the Bulma CSS framework with the Gomponents library."),
	),
).Section(
	"Importing packages", "",
	b.Content(
		e.P("It is suggested to import ", e.Em("Bulma-Gomponents"), " and ", e.Em("Gomplements"), " with the following aliases:"),
		c.ExamplePre(`import (
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
)`),
		e.P("You may also want to import ", e.Code("github.com/willoma/bulma-gomponents/fa"), " to generate ", e.Em("Font Awesome"), " icons."),
	),
).Section(
	"Bulma examples", "",
	b.Content(
		e.P("This documentation contains examples corresponding to the ", e.AHref("https://bulma.io/documentation/", "official Bulma documentation"), " examples."),
		e.P("This documentation is entirely written with ", e.Em("Bulma-Gomponents"), ", its source code is in ", e.AHref("https://github.com/willoma/bulma-gomponents/tree/docs", "the library repository"), ". Don't hesitate to refer to it if needed."),
	),
)
