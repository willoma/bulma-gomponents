package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

func demoNavbar(white bool, child any) *b.Element {
	var logoSrc string
	if white {
		logoSrc = "https://bulma.io/images/bulma-logo-white.png"
	} else {
		logoSrc = "https://bulma.io/images/bulma-logo.png"
	}
	return b.Navbar(
		child,
		b.NavbarBrand(
			b.NavbarAHref(
				"https://bulma.io",
				el.Img(
					html.Src(logoSrc),
					html.Alt("Bulma: a modern CSS framework based on Flexbox"),
					html.Width("112"), html.Height("28"),
				),
			),
		),
		b.NavbarStart(
			b.NavbarAHref("https://bulma.io/", "Home"),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownHoverable,
				easy.NavbarDropdownBoxed,
				b.NavbarAHref("https://bulma.io/documentation/overview/start/", "Overview"),
				b.NavbarAHref("https://bulma.io/documentation/overview/modifiers/", "Modifiers"),
				b.NavbarAHref("https://bulma.io/documentation/columns/basics/", "Columns"),
				b.NavbarAHref("https://bulma.io/documentation/layout/container/", "Layout"),
				b.NavbarAHref("https://bulma.io/documentation/form/general/", "Form"),
				b.NavbarDivider(),
				b.NavbarAHref("https://bulma.io/documentation/elements/box/", "Elements"),
				b.NavbarAHref("https://bulma.io/documentation/breadcrumb/", b.Active, "Components"),
			),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Field(
					b.Grouped,
					b.Control(
						html.P,
						b.ButtonA(
							html.DataAttr("social-network", "Twitter"),
							html.DataAttr("social-action", "tweet"),
							html.DataAttr("social-target", "https://bulma.io"),
							html.Target("_blank"),
							html.Href("https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=https://bulma.io&amp;via=jgthms"),
							fa.Icon(fa.Brand, "twitter"),
							"Tweet",
						),
					),
					b.Control(
						html.P,
						b.ButtonA(
							b.Primary,
							html.Href("https://github.com/jgthms/bulma/releases/download/0.9.4/bulma-0.9.4.zip"),
							fa.Icon(fa.Solid, "download"),
							"Download",
						),
					),
				),
			),
		),
	)
}

var navbar = c.NewPage(
	"Navbar", "Navbar", "/navbar",
	"https://bulma.io/documentation/components/navbar/",
).Section(
	"Basic Navbar",
	"https://bulma.io/documentation/components/navbar/#basic-navbar",
	c.HorizontalExample(
		`b.Navbar(
	html.Role("navigation"),
	html.Aria("label", "main navigation"),
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			b.ImgSrc(
				"https://bulma.io/images/bulma-logo.png",
				html.Width("112"), html.Height("28"),
			),
		),
	),
	b.NavbarStart(
		b.NavbarAHref("#", "Home"),
		b.NavbarAHref("#", "Documentation"),
		easy.NavbarDropdown(
			"More",
			b.NavbarDropdownHoverable,
			b.NavbarAHref("#", "About"),
			b.NavbarAHref("#", "Jobs"),
			b.NavbarAHref("#", "Contact"),
			b.NavbarDivider(),
			b.NavbarAHref("#", "Report an issue"),
		),
	),
	b.NavbarEnd(
		b.NavbarItem(
			b.Buttons(
				b.ButtonA(
					b.Primary,
					el.Strong("Sign up"),
				),
				b.ButtonA(
					b.Light,
					"Log in",
				),
			),
		),
	),
)`,
		b.Navbar(
			html.Role("navigation"),
			html.Aria("label", "main navigation"),
			b.NavbarBrand(
				b.NavbarAHref(
					"https://bulma.io",
					b.ImgSrc(
						"https://bulma.io/images/bulma-logo.png",
						html.Width("112"), html.Height("28"),
					),
				),
			),
			b.NavbarStart(
				b.NavbarAHref("#", "Home"),
				b.NavbarAHref("#", "Documentation"),
				easy.NavbarDropdown(
					"More",
					easy.NavbarDropdownHoverable,
					b.NavbarAHref("#", "About"),
					b.NavbarAHref("#", "Jobs"),
					b.NavbarAHref("#", "Contact"),
					b.NavbarDivider(),
					b.NavbarAHref("#", "Report an issue"),
				),
			),
			b.NavbarEnd(
				b.NavbarItem(
					b.Buttons(
						b.ButtonA(
							b.Primary,
							el.Strong("Sign up"),
						),
						b.ButtonA(
							b.Light,
							"Log in",
						),
					),
				),
			),
		),
	),
).Section(
	"Navbar brand",
	"https://bulma.io/documentation/components/navbar/#navbar-brand",
	c.Example(
		`b.Navbar(
	html.Role("navigation"),
	html.Aria("label", "main navigation"),
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			b.ImgSrc(
				"https://bulma.io/images/bulma-logo.png",
				html.Width("112"), html.Height("28"),
			),
		),
	),
)`,
		b.Navbar(
			html.Role("navigation"),
			html.Aria("label", "main navigation"),
			b.NavbarBrand(
				b.NavbarAHref(
					"https://bulma.io",
					b.ImgSrc(
						"https://bulma.io/images/bulma-logo.png",
						html.Width("112"), html.Height("28"),
					),
				),
			),
		),
	),
).Section(
	"Navbar burger",
	"https://bulma.io/documentation/components/navbar/#navbar-burger",
	b.Content("Nothing specific here. The burger is automatically generated."),
).Section(
	"Navbar menu",
	"https://bulma.io/documentation/components/navbar/#navbar-menu",
	b.Content(
		el.P("JavaScript is already included, no addition needed for the base use-case. If you need to trigger the change, you can add an ID with ", el.Code("html.ID"), " and use the following:"),
		el.Pre(`const navbar = document.getElementById("your-id")
navbar.getElementsByClassName("navbar-burger")[0].classList.toggle("is-active")
navbar.getElementsByClassName("navbar-menu")[0].classList.toggle("is-active")`),
	),
).Section(
	"Navbar start and navbar end",
	"https://bulma.io/documentation/components/navbar/#navbar-start-and-navbar-end",
	b.Content(
		b.UList(
			[]any{"left section: ", el.Code("b.NavbarStart")},
			[]any{"right section: ", el.Code("b.NavbarEnd")},
		),
	),
).Section(
	"Navbar item",
	"https://bulma.io/documentation/components/navbar/#navbar-item",
	b.Content(
		el.Ul(
			el.Li(
				"a navigation ", el.Strong("link"),
				el.Pre(`b.NavbarAHref("#", "Home")`),
			),
			el.Li(
				"a container for the ", el.Strong("brand logo"),
				el.Pre(`b.NavbarAHref(
	"#",
	b.ImgSrc(
		"https://bulma.io/images/bulma-logo.png",
		html.Width("112"), html.Height("28"),
		html.Alt("Bulma"),
	),
)`),
			),
			el.Li(
				"the ", el.Strong("parent"), " of a dropdown menu",
				el.Pre(`b.NavbarItem(
	b.HasDropdown,
	b.NavbarLink("Docs"),
	b.NavbarDropdown(
		// Other navbar items
	),
),`),
				"or use the ", el.Code("easy.NavbarDropdown"), " helper",
			),
			el.Li(
				"a child of a ", el.Strong("navbar dropdown"),
				el.Pre(`b.NavbarDropdown(
	b.NavbarAHref("#", "Overview"),
)`),
			),
			el.Li(
				"a container for almost ", el.Strong("anything"), " you want, like a ", el.Code("field"),
				el.Pre(`b.NavbarItem(
	b.Field(
		b.Grouped,
		b.Control(
			html.P,
			b.ButtonA(
				fa.Icon(fa.Brand, "twitter"),
				"Tweet",
			),
		),
		b.Control(
			b.ButtonA(
				b.Primary,
				fa.Icon(fa.Solid, "download"),
				"Download"
			),
		),
	),
)`),
			),
		),
		el.P("You can add the following ", el.Strong("modifier"), " classes:"),
		el.Ul(
			el.Li(el.Code("b.Expanded"), " to turn it into a full-width element"),
			el.Li(el.Code("b.Tab"), " to add a bottom border on hover and show the bottom border using ", el.Code("b.Active")),
		),
	),
).Section(
	"Transparent navbar",
	"https://bulma.io/documentation/components/navbar/#transparent-navbar",
	c.HorizontalExample(
		`b.Navbar(
	children,
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			el.Img(
				html.Src("https://bulma.io/images/bulma-logo.png"),
				html.Alt("Bulma: a modern CSS framework based on Flexbox"),
				html.Width("112"), html.Height("28"),
			),
		),
	),
	b.NavbarStart(
		b.NavbarAHref("https://bulma.io/", "Home"),
		easy.NavbarDropdown(
			"Docs",
			b.NavbarDropdownHoverable,
			b.NavbarDropdownBoxed,
			b.NavbarAHref("https://bulma.io/documentation/overview/start/", "Overview"),
			b.NavbarAHref("https://bulma.io/documentation/overview/modifiers/", "Modifiers"),
			b.NavbarAHref("https://bulma.io/documentation/columns/basics/", "Columns"),
			b.NavbarAHref("https://bulma.io/documentation/layout/container/", "Layout"),
			b.NavbarAHref("https://bulma.io/documentation/form/general/", "Form"),
			b.NavbarDivider(),
			b.NavbarAHref("https://bulma.io/documentation/elements/box/", "Elements"),
			b.NavbarAHref("https://bulma.io/documentation/breadcrumb/", b.Active, "Components"),
		),
	),
	b.NavbarEnd(
		b.NavbarItem(
			b.Field(
				b.Grouped,
				b.Control(
					html.P,
					b.ButtonA(
						html.DataAttr("social-network", "Twitter"),
						html.DataAttr("social-action", "tweet"),
						html.DataAttr("social-target", "https://bulma.io"),
						html.Target("_blank"),
						html.Href("https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=https://bulma.io&amp;via=jgthms"),
						fa.Icon(fa.Brand, "twitter"),
						"Tweet",
					),
				),
				b.Control(
					html.P,
					b.ButtonA(
						b.Primary,
						html.Href("https://github.com/jgthms/bulma/releases/download/0.9.4/bulma-0.9.4.zip"),
						fa.Icon(fa.Solid, "download"),
						"Download",
					),
				),
			),
		),
	),
)`,
		demoNavbar(false, b.Transparent),
	),
).Section(
	"Fixed navbar",
	"https://bulma.io/documentation/components/navbar/#fixed-navbar",
	b.Content(
		el.P("You may manually add ", el.Code("b.FixedTop"), " or ", el.Code("b.FixedBottom"), " to ", el.Code("b.Navbar"), ", and add ", el.Code("b.NavbarFixedTop"), " or ", el.Code("b.NavbarFixedBottom"), " to ", el.Code("b.HTML"), "."),
		el.P("However, it is easier to use ", el.Code("b.TopNavbar"), " or ", el.Code("b.BottomNavbar"), " as a ", el.Strong("direct child"), " of ", el.Code("b.HTML"), " instead of ", el.Code("b.Navbar"), ", in which case the corresponding class is automatically added to the body."),
	),
).Section(
	"Dropdown menu",
	"https://bulma.io/documentation/components/navbar/#dropdown-menu",
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	easy.NavbarDropdown(
		"Docs",
		b.NavbarDropdownHoverable,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownHoverable,
				b.NavbarAHref("#", "Overview"),
				b.NavbarAHref("#", "Elements"),
				b.NavbarAHref("#", "Components"),
				b.NavbarDivider(),
				b.NavbarItem("Version 0.9.4"),
			),
		),
	),
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	easy.NavbarDropdown(
		"Docs",
		b.NavbarDropdownClickable,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownClickable,
				b.NavbarAHref("#", "Overview"),
				b.NavbarAHref("#", "Elements"),
				b.NavbarAHref("#", "Components"),
				b.NavbarDivider(),
				b.NavbarItem("Version 0.9.4"),
			),
		),
	),
	b.Content(
		el.H3("Right dropdown"),
		el.P("If your parent ", el.Code("b.NavbarItem"), " is on the right side, you can position the dropdown to start from the ", el.Strong("right"), " with the ", el.Code("b.Right"), " modifier."),
	),
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarStart(
		easy.NavbarDropdown(
			"Left",
			b.NavbarDropdownClickable,
			b.NavbarAHref("#", "Overview"),
			b.NavbarAHref("#", "Elements"),
			b.NavbarAHref("#", "Components"),
			b.NavbarDivider(),
			b.NavbarItem("Version 0.9.4"),
		),
	),
	b.NavbarEnd(
		easy.NavbarDropdown(
			"Right",
			b.NavbarDropdownClickable,
			b.Right,
			b.NavbarAHref("#", "Overview"),
			b.NavbarAHref("#", "Elements"),
			b.NavbarAHref("#", "Components"),
			b.NavbarDivider(),
			b.NavbarItem("Version 0.9.4"),
		),
	),
),
b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", el.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarStart(
				easy.NavbarDropdown(
					"Left",
					easy.NavbarDropdownClickable,
					b.NavbarAHref("#", "Overview"),
					b.NavbarAHref("#", "Elements"),
					b.NavbarAHref("#", "Components"),
					b.NavbarDivider(),
					b.NavbarItem("Version 0.9.4"),
				),
			),
			b.NavbarEnd(
				easy.NavbarDropdown(
					"Right",
					easy.NavbarDropdownClickable,
					b.Right,
					b.NavbarAHref("#", "Overview"),
					b.NavbarAHref("#", "Elements"),
					b.NavbarAHref("#", "Components"),
					b.NavbarDivider(),
					b.NavbarItem("Version 0.9.4"),
				),
			),
		),
		b.Hero(
			b.Primary,
			b.Title(html.P, "Documentation"),
			b.Subtitle(
				html.P,
				"Everything you need to ", el.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		el.H3("Dropup"),
		el.P("If you're using a navbar at the bottom, like the ", b.AHref("http://localhost:8080/navbar#fixed-navbar", "fixed bottom navbar"), ", you might want to use a ", el.Strong("dropup menu"), ". Simply add the ", el.Code("b.NavbarDropup"), " modifier to the parent ", el.Code("b.NavbarItem"), ":"),
	),
	c.Example(
		`b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", el.Strong("create a website"), " with Bulma",
	),
),
b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarStart(
		easy.NavbarDropdown(
			"Dropup",
			b.NavbarDropdownClickable,
			b.NavbarDropup,
			b.NavbarAHref("#", "Overview"),
			b.NavbarAHref("#", "Elements"),
			b.NavbarAHref("#", "Components"),
			b.NavbarDivider(),
			b.NavbarItem("Version 0.9.4"),
		),
	),
)`,
		b.Hero(
			b.Primary,
			b.Title(html.P, "Documentation"),
			b.Subtitle(
				html.P,
				"Everything you need to ", el.Strong("create a website"), " with Bulma",
			),
		),
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarStart(
				easy.NavbarDropdown(
					"Dropup",
					easy.NavbarDropdownClickable,
					easy.NavbarDropup,
					b.NavbarAHref("#", "Overview"),
					b.NavbarAHref("#", "Elements"),
					b.NavbarAHref("#", "Components"),
					b.NavbarDivider(),
					b.NavbarItem("Version 0.9.4"),
				),
			),
		),
	),
	b.Content(
		el.H3("Dropdown without arrow"),
		el.P("You can remove the arrow in the items of the navbar by adding the ", el.Code("b.Arrowless"), " modifier to the ", el.Code("b.NavbarLink"), "element, or the ", el.Code("b.NavbarDropdownArrowless"), " option to the ", el.Code("easy.NavbarDropdown"), " element."),
	),
	c.Example(
		`easy.NavbarDropdown(
	"Link without arrow",
	b.NavbarDropdownHoverable,
	b.NavbarDropdownArrowless,
	b.NavbarAHref("#", "Overview"),
	b.NavbarAHref("#", "Elements"),
	b.NavbarAHref("#", "Components"),
	b.NavbarDivider(),
	b.NavbarItem("Version 0.9.4"),
)`,
		easy.NavbarDropdown(
			"Link without arrow",
			easy.NavbarDropdownHoverable,
			easy.NavbarDropdownArrowless,
			b.NavbarAHref("#", "Overview"),
			b.NavbarAHref("#", "Elements"),
			b.NavbarAHref("#", "Components"),
			b.NavbarDivider(),
			b.NavbarItem("Version 0.9.4"),
		),
	),
	b.Content(
		el.H3("Styles for the dropdown menu"),
	),
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarAHref(
		"#",
		b.ImgSrc(
			"https://bulma.io/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	easy.NavbarDropdown(
		"Docs",
		b.NavbarDropdownActive,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
),
b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", el.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarAHref(
				"#",
				b.ImgSrc(
					"https://bulma.io/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownActive,
				b.NavbarAHref("#", "Overview"),
				b.NavbarAHref("#", "Elements"),
				b.NavbarAHref("#", "Components"),
				b.NavbarDivider(),
				b.NavbarItem("Version 0.9.4"),
			),
		),
		b.Hero(
			b.Primary,
			b.Title(html.P, "Documentation"),
			b.Subtitle(
				html.P,
				"Everything you need to ", el.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		el.P("When having a ", b.AHref("http://localhost:8080/navbar#transparent-navbar", "transparent navbar"), ", it is preferable to use the boxed version of the dropdown, by using the ", el.Code("b.NavbarDropdownBoxed"), " option."),
	),
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarAHref(
		"#",
		b.ImgSrc(
			"https://bulma.io/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	easy.NavbarDropdown(
		"Docs",
		b.NavbarDropdownActive,
		b.NavbarDropdownBoxed,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
),
b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", el.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarAHref(
				"#",
				b.ImgSrc(
					"https://bulma.io/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownActive,
				easy.NavbarDropdownBoxed,
				b.NavbarAHref("#", "Overview"),
				b.NavbarAHref("#", "Elements"),
				b.NavbarAHref("#", "Components"),
				b.NavbarDivider(),
				b.NavbarItem("Version 0.9.4"),
			),
		),
		b.Hero(
			b.Primary,
			b.Title(html.P, "Documentation"),
			b.Subtitle(
				html.P,
				"Everything you need to ", el.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(el.H3("Active dropdown navbar item")),
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarAHref(
		"#",
		b.ImgSrc(
			"https://bulma.io/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	easy.NavbarDropdown(
		"Docs",
		b.NavbarDropdownActive,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", b.Active, "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
),
b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", el.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarAHref(
				"#",
				b.ImgSrc(
					"https://bulma.io/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			easy.NavbarDropdown(
				"Docs",
				easy.NavbarDropdownActive,
				b.NavbarAHref("#", "Overview"),
				b.NavbarAHref("#", b.Active, "Elements"),
				b.NavbarAHref("#", "Components"),
				b.NavbarDivider(),
				b.NavbarItem("Version 0.9.4"),
			),
		),
		b.Hero(
			b.Primary,
			b.Title(html.P, "Documentation"),
			b.Subtitle(
				html.P,
				"Everything you need to ", el.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		el.H3(" Dropdown divider"),
		el.Pre("b.NavbarDivider()"),
	),
).Section(
	"Colors",
	"https://bulma.io/documentation/components/navbar/#colors",
	b.Content(
		el.P("In order to change the background color of the navbarn simply, apply one of the color modifiers to the ", el.Code("b.Navbar"), " element:"),
		b.UList(
			el.Code("b.Primary"),
			el.Code("b.Link"),
			el.Code("b.Info"),
			el.Code("b.Success"),
			el.Code("b.Warning"),
			el.Code("b.Danger"),
			el.Code("b.Black"),
			el.Code("b.Dark"),
			el.Code("b.Light"),
			el.Code("b.White"),
		),
		b.Block(demoNavbar(true, b.Primary)),
		b.Block(demoNavbar(true, b.Link)),
		b.Block(demoNavbar(true, b.Info)),
		b.Block(demoNavbar(false, b.Success)),
		b.Block(demoNavbar(false, b.Warning)),
		b.Block(demoNavbar(true, b.Danger)),
		b.Block(demoNavbar(true, b.Black)),
		b.Block(demoNavbar(true, b.Dark)),
		b.Block(demoNavbar(false, b.Light)),
		b.Block(demoNavbar(false, b.White)),
	),
).Section(
	"Navbar Helper Classes",
	"https://bulma.io/documentation/components/navbar/#navbar-helper-classes",
	b.Content(
		el.Ul(
			el.Li(
				el.Code("b.Spaced"), " sets ", el.Strong("Top"), " and ", el.Strong("Bottom"), " paddings with ", el.Strong("1rem"), ", ", el.Strong("Left"), " and ", el.Strong("Right"), " paddings with ", el.Strong("2rem"),
			),
			el.Li(
				el.Code("b.Shadow"), " adds a small amount of box-shadow around the navbar",
			),
		),
	),
)
