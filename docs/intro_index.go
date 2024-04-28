package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var intro = c.NewPage(
	"Introduction", "Bulma-Gomponents", "/index",
	"",

	b.Content(
		el.P("This ", el.Em("Go"), " library makes it easier to use the Bulma CSS framework with the Gomponents library."),
	),
).Section(
	"Importing packages", "",
	b.Content(
		el.P("It is suggested to import ", el.Em("Bulma-Gomponents"), " with the ", el.Code("b"), " alias:"),
		c.ExamplePre(`import b "github.com/willoma/bulma-gomponents"`),
		el.P(el.Em("Bulma-Gomponents"), " also provides the following packages:"),
		b.DList(
			el.Code("github.com/willoma/bulma-gomponents/easy"),
			"Easy-to-use opiniated helpers to generate some components.",
			el.Code("github.com/willoma/bulma-gomponents/el"),
			[]any{"Base HTML elements, implemented as ", el.Code("b.Element"), " objects, as alternatives to functions in ", el.Code("github.com/maragudk/gomponents/html"), "."},
			el.Code("github.com/willoma/bulma-gomponents/fa"),
			[]any{"Helpers to generate ", el.Em("Font Awesome"), " icons as ", el.Code("b.Element"), "objects."},
		),
	),
).Section(
	"Bulma examples", "",
	b.Content(
		el.P("This documentation contains examples corresponding to the ", b.AHref("https://bulma.io/documentation/", "official Bulma documentation"), " examples."),
		el.P("This documentation is entirely written with ", el.Em("Bulma-Gomponents"), ", its source code is in ", b.AHref("https://github.com/willoma/bulma-gomponents/tree/main/docs", "the library repository"), ". Don't hesitate to refer to it if needed."),
	),
)
