package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var footer = c.NewPage(
	"Footer", "Footer", "/footer",
	"https://bulma.io/documentation/layout/footer/",
	c.Example(
		`b.Footer(
	b.Content(
		b.TextCentered,
		el.P(
			el.Strong("Bulma"), " by ", b.AHref("https://jgthms.com", "Jeremy Thomas"), ". The source code is licensed ", b.AHref("http://opensource.org/licenses/mit-license.php", "MIT"), ". The website content is licensed ", b.AHref("http://creativecommons.org/licenses/by-nc-sa/4.0/", "CC BY NC SA 4.0"), ".",
		),
	),
)`,
		b.Footer(
			b.Content(
				b.TextCentered,
				el.P(
					el.Strong("Bulma"), " by ", b.AHref("https://jgthms.com", "Jeremy Thomas"), ". The source code is licensed ", b.AHref("http://opensource.org/licenses/mit-license.php", "MIT"), ". The website content is licensed ", b.AHref("http://creativecommons.org/licenses/by-nc-sa/4.0/", "CC BY NC SA 4.0"), ".",
				),
			),
		),
	),
)
