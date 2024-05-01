package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var footer = c.NewPage(
	"Footer", "Footer", "/footer",
	"",

	b.Content(
		e.P("The ", e.Code("b.Footer"), " constructor creates a page footer."),
	),
).Section(
	"Bulma example",
	"https://bulma.io/documentation/layout/footer/",
	c.Example(
		`b.Footer(
	b.Content(
		b.TextCentered,
		e.P(
			e.Strong("Bulma"), " by ", e.AHref("https://jgthms.com", "Jeremy Thomas"), ". The source code is licensed ", e.AHref("http://opensource.org/licenses/mit-license.php", "MIT"), ". The website content is licensed ", e.AHref("http://creativecommons.org/licenses/by-nc-sa/4.0/", "CC BY NC SA 4.0"), ".",
		),
	),
)`,
		b.Footer(
			b.Content(
				b.TextCentered,
				e.P(
					e.Strong("Bulma"), " by ", e.AHref("https://jgthms.com", "Jeremy Thomas"), ". The source code is licensed ", e.AHref("http://opensource.org/licenses/mit-license.php", "MIT"), ". The website content is licensed ", e.AHref("http://creativecommons.org/licenses/by-nc-sa/4.0/", "CC BY NC SA 4.0"), ".",
				),
			),
		),
	),
)
