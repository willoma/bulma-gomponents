package docs

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var hero = c.NewPage(
	"Hero", "Hero", "/hero",
	"",

	b.Content(
		e.P("The ", e.Code("b.Hero"), " constructor creates a hero."),
		c.Modifiers(
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
			c.Row("b.HalfHeight", "Set size to 50% of the height of the viewport"),
			c.Row("b.FullHeight", "Set size to 100% of the height of the viewport"),
			c.Row("b.FullHeightWithNavbar", "Set size to 100% of the height of the viewport minus the navbar height"),
		),
		c.Children(
			c.Row("b.OnSection(...any)", "Apply children to the ", e.Code(`<section class="hero">`), " element"),
			c.Row("b.OnBody(...any)", "Apply children to the ", e.Code(`<div class="hero-body">`), " element"),
			c.Row("b.HeroHead(...any)", "Apply children to the ", e.Code(`<div class="hero-head">`), " element"),
			c.Row("b.HeroFoot(...any)", "Apply children to the ", e.Code(`<div class="hero-foot">`), " element"),
			c.Row("e.Element", "Add element to the ", e.Code(`<div class="hero-body">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<section class="hero">`), " element"),
			c.RowNodeElement("Add element to the ", e.Code(`<div class="hero-body">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<section class="hero">`), " element"),
		),
		e.P("The head, body and foot elements are created only if at least one child has been provided for the part."),
	),
).Section(
	"Bulma examples",
	"https://bulma.io/documentation/layout/hero/",

	b.Content(e.P("The ", e.Em("hero body"), " is implied when you provide elements as children to ", e.Code("b.Hero"), ":")),
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
).Subsection(
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
).Subsection(
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
	e.Div(
		b.Title(html.P, "Half height hero"),
		b.Subtitle(html.P, "Half height subtitle"),
	),
)`,
		b.Hero(
			b.HalfHeight, b.Success,
			e.Div(
				b.Title(html.P, "Half height hero"),
				b.Subtitle(html.P, "Half height subtitle"),
			),
		),
	),
	c.Example(
		`b.Hero(
	b.FullHeight, b.Danger,
	e.Div(
		b.Title(html.P, "Full height hero"),
		b.Subtitle(html.P, "Full height subtitle"),
	),
)`,
		b.Hero(
			b.FullHeight, b.Danger,
			e.Div(
				b.Title(html.P, "Full height hero"),
				b.Subtitle(html.P, "Full height subtitle"),
			),
		),
	),
).Subsection(
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
).Subsection(
	"Fullheight hero in 3 parts",
	"https://bulma.io/documentation/layout/hero/#fullheight-hero-in-3-parts",
	b.Content(
		e.P("Use ", e.Code("b.HeroHead"), " and ", e.Code("b.HeroFoot"), " to define body header and footer. Please note they can be used multiple times, in which case the children will be aggregated."),
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
					e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
			b.TabLink(b.Active, "Overview"),
			b.TabLink("Modifiers"),
			b.TabLink("Grid"),
			b.TabLink("Elements"),
			b.TabLink("Components"),
			b.TabLink("Layout"),
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
							e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
					b.TabLink(b.Active, "Overview"),
					b.TabLink("Modifiers"),
					b.TabLink("Grid"),
					b.TabLink("Elements"),
					b.TabLink("Components"),
					b.TabLink("Layout"),
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
					e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
			b.TabLink(b.Active, "Overview"),
			b.TabLink("Modifiers"),
			b.TabLink("Grid"),
			b.TabLink("Elements"),
			b.TabLink("Components"),
			b.TabLink("Layout"),
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
							e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
					b.TabLink(b.Active, "Overview"),
					b.TabLink("Modifiers"),
					b.TabLink("Grid"),
					b.TabLink("Elements"),
					b.TabLink("Components"),
					b.TabLink("Layout"),
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
					e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
			b.TabLink(b.Active, "Overview"),
			b.TabLink("Modifiers"),
			b.TabLink("Grid"),
			b.TabLink("Elements"),
			b.TabLink("Components"),
			b.TabLink("Layout"),
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
							e.ImgSrc("https://bulma.io/images/bulma-type-white.png", e.Alt("Logo")),
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
					b.TabLink(b.Active, "Overview"),
					b.TabLink("Modifiers"),
					b.TabLink("Grid"),
					b.TabLink("Elements"),
					b.TabLink("Components"),
					b.TabLink("Layout"),
				),
			),
		),
	),
)
