package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var hero = c.NewPage(
	"Hero", "Hero", "/hero",
	"https://bulma.io/documentation/layout/hero/",
	b.Content(el.P("The ", el.Em("hero body"), " is implied when you provide elements as children to ", el.Code("b.Hero"), ":")),
	c.Example(
		`b.Hero(
	b.Title(html.P, "Hero title"),
	b.Subtitle(html.P ,"Hero subtitle"),
)`,
		b.Hero(
			b.Title(html.P, "Hero title"),
			b.Subtitle(html.P, "Hero subtitle"),
		),
	),
).Section(
	"Colors",
	"https://bulma.io/documentation/layout/hero/#colors",
	c.Example(
		`b.Hero(
	b.Primary,
	b.Title(html.P, "Primary hero"),
	b.Subtitle(html.P ,"Primary subtitle"),
)`,
		b.Hero(
			b.Primary,
			b.Title(html.P, "Primary hero"),
			b.Subtitle(html.P, "Primary subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Link,
	b.Title(html.P, "Link hero"),
	b.Subtitle(html.P ,"Link subtitle"),
)`,
		b.Hero(
			b.Link,
			b.Title(html.P, "Link hero"),
			b.Subtitle(html.P, "Link subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Info,
	b.Title(html.P, "Info hero"),
	b.Subtitle(html.P ,"Info subtitle"),
)`,
		b.Hero(
			b.Info,
			b.Title(html.P, "Info hero"),
			b.Subtitle(html.P, "Info subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Success,
	b.Title(html.P, "Success hero"),
	b.Subtitle(html.P ,"Success subtitle"),
)`,
		b.Hero(
			b.Success,
			b.Title(html.P, "Success hero"),
			b.Subtitle(html.P, "Success subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Warning,
	b.Title(html.P, "Warning hero"),
	b.Subtitle(html.P ,"Warning subtitle"),
)`,
		b.Hero(
			b.Warning,
			b.Title(html.P, "Warning hero"),
			b.Subtitle(html.P, "Warning subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Danger,
	b.Title(html.P, "Danger hero"),
	b.Subtitle(html.P ,"Danger subtitle"),
)`,
		b.Hero(
			b.Danger,
			b.Title(html.P, "Danger hero"),
			b.Subtitle(html.P, "Danger subtitle"),
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/layout/hero/#sizes",
	c.Example(
		`b.Hero(
	b.Small, b.Primary,
	b.Title(html.P, "Small hero"),
	b.Subtitle(html.P, "Small subtitle"),
)`,
		b.Hero(
			b.Small, b.Primary,
			b.Title(html.P, "Small hero"),
			b.Subtitle(html.P, "Small subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Medium, b.Link,
	b.Title(html.P, "Medium hero"),
	b.Subtitle(html.P, "Medium subtitle"),
)`,
		b.Hero(
			b.Medium, b.Link,
			b.Title(html.P, "Medium hero"),
			b.Subtitle(html.P, "Medium subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.Large, b.Info,
	b.Title(html.P, "Large hero"),
	b.Subtitle(html.P, "Large subtitle"),
)`,
		b.Hero(
			b.Large, b.Info,
			b.Title(html.P, "Large hero"),
			b.Subtitle(html.P, "Large subtitle"),
		),
	),
	c.Example(
		`b.Hero(
	b.HalfHeight, b.Success,
	el.Div(
		b.Title(html.P, "Half height hero"),
		b.Subtitle(html.P, "Half height subtitle"),
	),
)`,
		b.Hero(
			b.HalfHeight, b.Success,
			el.Div(
				b.Title(html.P, "Half height hero"),
				b.Subtitle(html.P, "Half height subtitle"),
			),
		),
	),
	c.Example(
		`b.Hero(
	b.FullHeight, b.Danger,
	el.Div(
		b.Title(html.P, "Full height hero"),
		b.Subtitle(html.P, "Full height subtitle"),
	),
)`,
		b.Hero(
			b.FullHeight, b.Danger,
			el.Div(
				b.Title(html.P, "Full height hero"),
				b.Subtitle(html.P, "Full height subtitle"),
			),
		),
	),
).Section(
	"Fullheight with navbar",
	"https://bulma.io/documentation/layout/hero/#fullheight-with-navbar",
	c.HorizontalExample(
		`b.Navbar(
	b.NavbarStart(
		b.NavbarAHref("#", "Home"),
		b.NavbarAHref("#", "Documentation"),
	),
	b.NavbarEnd(
		b.NavbarItem(
			b.Buttons(
				b.ButtonA(b.Dark, "Github"),
				b.ButtonA(b.Link, "Download"),
			),
		),
	),
),
b.Hero(
	b.Link, b.FullHeightWithNavbar,
	b.Title(html.P, "Fullheight hero with navbar"),
)`,
		b.Navbar(
			b.NavbarStart(
				b.NavbarAHref("#", "Home"),
				b.NavbarAHref("#", "Documentation"),
			),
			b.NavbarEnd(
				b.NavbarItem(
					b.Buttons(
						b.ButtonA(b.Dark, "Github"),
						b.ButtonA(b.Link, "Download"),
					),
				),
			),
		),
		b.Hero(
			b.Link, b.FullHeightWithNavbar,
			b.Title(html.P, "Fullheight hero with navbar"),
		),
	),
).Section(
	"Fullheight hero in 3 parts",
	"https://bulma.io/documentation/layout/hero/#fullheight-hero-in-3-parts",
	b.Content(
		el.P("Use ", el.Code("b.HeroHead"), " and ", el.Code("b.HeroFoot"), " to define body header and footer. Please note they can be used multiple times, in which case the children will be aggregated."),
	),
	c.HorizontalExample(
		`b.Hero(
	b.Primary, b.Medium,
	// Hero head: will stick at the top
	b.HeroHead(
		b.Navbar(
			b.NavbarBrand(
				b.NavbarAHref(
					"#",
					b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
				),
			),
			b.NavbarEnd(
				b.NavbarAHref("#", b.Active, "Home"),
				b.NavbarAHref("#", "Examples"),
				b.NavbarAHref("#", "Documentation"),
				b.NavbarItem(
					html.Span,
					b.ButtonA(
						b.Primary, b.Inverted,
						fa.Icon(fa.Brand, "github"),
						"Download",
					),
				),
			),
		),
	),

	// Hero content: will be in the middle
	b.Container(
		b.TextCentered,
		b.Title(html.P, "Title"),
		b.Subtitle(html.P, "Subtitle"),
	),

	// Hero footer: will stick at the bottom
	b.HeroFoot(
		b.Tabs(
			b.Container,
			b.TabsLink(b.Active, "Overview"),
			b.TabsLink("Modifiers"),
			b.TabsLink("Grid"),
			b.TabsLink("Elements"),
			b.TabsLink("Components"),
			b.TabsLink("Layout"),
		),
	),
)`,
		b.Hero(
			b.Primary, b.Medium,
			// Hero head: will stick at the top
			b.HeroHead(
				b.Navbar(
					b.Container,
					b.NavbarBrand(
						b.NavbarAHref(
							"#",
							b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
						),
					),
					b.NavbarEnd(
						b.NavbarAHref("#", b.Active, "Home"),
						b.NavbarAHref("#", "Examples"),
						b.NavbarAHref("#", "Documentation"),
						b.NavbarItem(
							html.Span,
							b.ButtonA(
								b.Primary, b.Inverted,
								fa.Icon(fa.Brand, "github"),
								"Download",
							),
						),
					),
				),
			),

			// Hero content: will be in the middle
			b.Container(
				b.TextCentered,
				b.Title(html.P, "Title"),
				b.Subtitle(html.P, "Subtitle"),
			),

			// Hero footer: will stick at the bottom
			b.HeroFoot(
				b.Tabs(
					b.Container,
					b.TabsLink(b.Active, "Overview"),
					b.TabsLink("Modifiers"),
					b.TabsLink("Grid"),
					b.TabsLink("Elements"),
					b.TabsLink("Components"),
					b.TabsLink("Layout"),
				),
			),
		),
	),
	c.HorizontalExample(
		`b.Hero(
	b.Info, b.Large,
	b.HeroHead(
		b.Navbar(
			b.Container,
			b.NavbarBrand(
				b.NavbarAHref(
					"#",
					b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
				),
			),
			b.NavbarEnd(
				b.NavbarAHref("#", b.Active, "Home"),
				b.NavbarAHref("#", "Examples"),
				b.NavbarAHref("#", "Documentation"),
				b.NavbarItem(
					html.Span,
					b.ButtonA(
						b.Primary, b.Inverted,
						fa.Icon(fa.Brand, "github"),
						"Download",
					),
				),
			),
		),
	),
	b.Container(
		b.TextCentered,
		b.Title(html.P, "Title"),
		b.Subtitle(html.P, "Subtitle"),
	),
	b.HeroFoot(
		b.Tabs(
			b.Boxed, b.FullWidth,
			b.Container,
			b.TabsLink(b.Active, "Overview"),
			b.TabsLink("Modifiers"),
			b.TabsLink("Grid"),
			b.TabsLink("Elements"),
			b.TabsLink("Components"),
			b.TabsLink("Layout"),
		),
	),
)`,
		b.Hero(
			b.Info, b.Large,
			b.HeroHead(
				b.Navbar(
					b.Container,
					b.NavbarBrand(
						b.NavbarAHref(
							"#",
							b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
						),
					),
					b.NavbarEnd(
						b.NavbarAHref("#", b.Active, "Home"),
						b.NavbarAHref("#", "Examples"),
						b.NavbarAHref("#", "Documentation"),
						b.NavbarItem(
							html.Span,
							b.ButtonA(
								b.Primary, b.Inverted,
								fa.Icon(fa.Brand, "github"),
								"Download",
							),
						),
					),
				),
			),
			b.Container(
				b.TextCentered,
				b.Title(html.P, "Title"),
				b.Subtitle(html.P, "Subtitle"),
			),
			b.HeroFoot(
				b.Tabs(
					b.Boxed, b.FullWidth,
					b.Container,
					b.TabsLink(b.Active, "Overview"),
					b.TabsLink("Modifiers"),
					b.TabsLink("Grid"),
					b.TabsLink("Elements"),
					b.TabsLink("Components"),
					b.TabsLink("Layout"),
				),
			),
		),
	),
	c.HorizontalExample(
		`b.Hero(
	b.Success, b.FullHeight,
	b.HeroHead(
		b.Navbar(
			b.Container,
			b.NavbarBrand(
				b.NavbarAHref(
					"#",
					b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
				),
			),
			b.NavbarEnd(
				b.NavbarAHref("#", b.Active, "Home"),
				b.NavbarAHref("#", "Examples"),
				b.NavbarAHref("#", "Documentation"),
				b.NavbarItem(
					html.Span,
					b.ButtonA(
						b.Primary, b.Inverted,
						fa.Icon(fa.Brand, "github"),
						"Download",
					),
				),
			),
		),
	),
	b.Container(
		b.TextCentered,
		b.Title(html.P, "Title"),
		b.Subtitle(html.P, "Subtitle"),
	),
	b.HeroFoot(
		b.Tabs(
			b.Boxed, b.FullWidth,
			b.Container,
			b.TabsLink(b.Active, "Overview"),
			b.TabsLink("Modifiers"),
			b.TabsLink("Grid"),
			b.TabsLink("Elements"),
			b.TabsLink("Components"),
			b.TabsLink("Layout"),
		),
	),
)`,
		b.Hero(
			b.Success, b.FullHeight,
			b.HeroHead(
				b.Navbar(
					b.Container,
					b.NavbarBrand(
						b.NavbarAHref(
							"#",
							b.ImgSrc("https://bulma.io/images/bulma-type-white.png", html.Alt("Logo")),
						),
					),
					b.NavbarEnd(
						b.NavbarAHref("#", b.Active, "Home"),
						b.NavbarAHref("#", "Examples"),
						b.NavbarAHref("#", "Documentation"),
						b.NavbarItem(
							html.Span,
							b.ButtonA(
								b.Primary, b.Inverted,
								fa.Icon(fa.Brand, "github"),
								"Download",
							),
						),
					),
				),
			),
			b.Container(
				b.TextCentered,
				b.Title(html.P, "Title"),
				b.Subtitle(html.P, "Subtitle"),
			),
			b.HeroFoot(
				b.Tabs(
					b.Boxed, b.FullWidth,
					b.Container,
					b.TabsLink(b.Active, "Overview"),
					b.TabsLink("Modifiers"),
					b.TabsLink("Grid"),
					b.TabsLink("Elements"),
					b.TabsLink("Components"),
					b.TabsLink("Layout"),
				),
			),
		),
	),
)
