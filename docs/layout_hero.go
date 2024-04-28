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
	"",

	b.Content(
		el.P("The ", el.Code("b.Hero"), " constructor creates a hero. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnBody(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="hero-body">`), " element"},

			el.Code("b.On(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<section class="hero">`), " element"},

			el.Code("b.HeroHead(...)"),
			"Add the children to the head part of the hero",

			el.Code("b.HeroFoot(...)"),
			"Add the children to the foot part of the hero",

			el.Code("b.Element"),
			"Add this element to the body part of the hero",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the hero section",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the body part",

			el.Code("b.Primary"),
			"Set hero color to primary",

			el.Code("b.Link"),
			"Set hero color to link",

			el.Code("b.Info"),
			"Set hero color to info",

			el.Code("b.Success"),
			"Set hero color to success",

			el.Code("b.Warning"),
			"Set hero color to warning",

			el.Code("b.Danger"),
			"Set hero color to danger",

			el.Code("b.Small"),
			"Set hero size to small",

			el.Code("b.Medium"),
			"Set hero size to medium",

			el.Code("b.Large"),
			"Set hero size to large",

			el.Code("b.HalfHeight"),
			"Set hero size to 50% of the height of the viewport",

			el.Code("b.FullHeight"),
			"Set hero size to 100% of the height of the viewport",

			el.Code("b.FullHeightWithNavbar"),
			"Set hero size to 100% of the height of the viewport minus the navbar height",
		),
		el.P("Other children are added to the ", el.Code(`<section class="hero">`), " element. The head, body and foot elements are created only if at least one child has been provided for the part."),
	),
).Section(
	"Bulma examples",
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
