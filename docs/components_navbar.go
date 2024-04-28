package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

func demoNavbar(white bool, color b.Color) b.Element {
	var logoSrc string
	if white {
		logoSrc = "https://bulma.io/assets/brand/Bulma%20Logo%20White.svg"
	} else {
		logoSrc = "https://bulma.io/assets/brand/Bulma%20Logo%20Black.svg"
	}
	return b.Navbar(
		b.Style("margin-bottom", "20rem"),
		color,
		b.NavbarBrand(
			b.NavbarAHref("https://bulma.io", b.ImgSrc(logoSrc)),
		),
		b.NavbarStart(
			b.NavbarAHref("https://bulma.io/", "Home"),
			b.NavbarDropdown(
				b.OnLink(
					el.A,
					html.Href("https://bulma.io/documentation/overview/start/"),
				),
				"Docs",
				b.Active,
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
							color.Soft(),
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
							color.Light(),
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
	"",

	b.Content(
		el.P("The ", el.Code("b.Navbar"), " constructor creates a navbar. The brand section is added only if at least one brand, start or end child is provided. The section part and the burger button are added only if at least one start or end child is provided. The click event on the burger, which toggles the navbar menu on mobile devices, is already included. The following children have a special meaning:"),
		b.DList(
			el.Code("b.NavbarBrand(...)"),
			"Add children to the brand section (left side, always visible)",

			el.Code("b.NavbarStart(...)"),
			"Add children to the start section (left part of the menu)",

			el.Code("b.NavbarEnd(...)"),
			"Add children to the end section (right part of the menu)",

			el.Code("b.Container()"),
			"Use this element as an intermediate container",

			el.Code("b.FixedTop"),
			[]any{"Fix the navbar to the top of the page - the ", el.Code("has-navbar-fixed-top"), " class is automatically added to the parent element (which should be the body for this to work)"},

			el.Code("b.FixedBottom"),
			[]any{"Fix the navbar to the bottom of the page - the ", el.Code("has-navbar-fixed-bottom"), " class is automatically added to the parent element (which should be the body for this to work)"},

			el.Code("b.Transparent"),
			"Remove hover and active backgrounds from the items",

			el.Code("b.Spaced"),
			"Increase padding",

			el.Code("b.Shadow"),
			"Add a small shadow around the navbar",

			el.Code("b.Primary"),
			"Set icon color to primary",

			el.Code("b.Link"),
			"Set icon color to link",

			el.Code("b.Info"),
			"Set icon color to info",

			el.Code("b.Success"),
			"Set icon color to success",

			el.Code("b.Warning"),
			"Set icon color to warning",

			el.Code("b.Danger"),
			"Set icon color to danger",

			el.Code("b.Black"),
			"Set icon color to black",

			el.Code("b.Light"),
			"Set icon color to light",

			el.Code("b.Dark"),
			"Set icon color to dark",

			el.Code("b.White"),
			"Set icon color to white",
		),

		el.P("Other children are added to the ", el.Code(`<nav class="navbar">`), " element."),

		el.P("The ", el.Code("b.NavbarItem"), " constructor creates a navbar item for a navbar brand, start or end section, or for a ", el.Code("b.NavbarDropdown"), ". The following children have a special meaning:"),
		b.DList(
			el.Code("b.Expanded"),
			"Turn the item into a full-width element",

			el.Code("b.Tab"),
			[]any{"Add a bottom border on hover, always show the bottom border when adding ", el.Code("b.Active")},

			el.Code("b.HasDropdown"),
			"Make the item a link+dropdown container",

			el.Code("b.HasDropup"),
			"Make the item a dropdown container, with the dropdown opening above the link",

			el.Code("b.Hoverable"),
			"Make the included dropdown automatically show on hover",

			el.Code("b.Active"),
			"Force the dropdown to be open",
		),

		el.P("The ", el.Code("b.NavbarAHref"), " constructor creates a link item for a navbar brand, start or end section, or to a ", el.Code("b.NavbarDropdown"), ". The following children have a special meaning:"),
		b.DList(
			el.Code("b.Expanded"),
			"Turn the item into a full-width element",

			el.Code("b.Tab"),
			[]any{"Add a bottom border on hover, always show the bottom border when adding ", el.Code("b.Active")},
		),

		el.P("The ", el.Code("b.NavbarDropdown"), " constructor creates navbar item with a dropdown menu. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnItem(...)"),
			[]any{"Force childen to be applied to the item"},

			el.Code("b.OnLink(...)"),
			[]any{"Force childen to be applied to the trigger link"},

			el.Code("b.OnDropdown(...)"),
			[]any{"Force childen to be applied to the dropdown menu"},

			el.Code("b.Hoverable"),
			"Make the menu open when the cursor hovers the trigger",

			el.Code("b.Clickable"),
			"Make the menu open when the button is clicked (javascript is automatically added)",

			el.Code("b.Up"),
			"Make the dropdown open to the top",

			el.Code("b.Arrowless"),
			"Remove the arrow in the item",

			el.Code("b.Boxed"),
			"Show the dropdown as a boxed element, without the top grey border",

			el.Code("b.Active"),
			"Force the dropdown to be open",

			el.Code("b.Right"),
			"Used when a dropdown is on the right side (end section), in order to open its content aligned to the right of the link",

			el.Code("string"),
			"Use this string as the content of the trigger",
		),
		el.P("Other children are added to the dropdown menu."),

		el.P("The ", el.Code("b.NavbarLink"), " constructor creates a link element, to include in a ", el.Code("b.NavbarItem"), " with the ", el.Code("b.HasDropdown"), " modifier. The following children have a special meaning:"),
		b.DList(
			el.Code("b.Arrowless"),
			"Remove the arrow in the item",
		),

		el.P("The ", el.Code("b.NavbarDivider"), " constructor creates a divider element, to include in a ", el.Code("b.NavbarDropdown"), "."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/navbar/",
).Subsection(
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
				"https://bulma.io/assets/images/bulma-logo.png",
				html.Width("112"), html.Height("28"),
			),
		),
	),
	b.NavbarStart(
		b.NavbarAHref("#", "Home"),
		b.NavbarAHref("#", "Documentation"),
		b.NavbarDropdown(
			"More",
			b.Hoverable,
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
						"https://bulma.io/assets/images/bulma-logo.png",
						html.Width("112"), html.Height("28"),
					),
				),
			),
			b.NavbarStart(
				b.NavbarAHref("#", "Home"),
				b.NavbarAHref("#", "Documentation"),
				b.NavbarDropdown(
					"More",
					b.Hoverable,
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
).Subsection(
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
				"https://bulma.io/assets/images/bulma-logo.png",
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
						"https://bulma.io/assets/images/bulma-logo.png",
						html.Width("112"), html.Height("28"),
					),
				),
			),
		),
	),
).Subsection(
	"Navbar burger",
	"https://bulma.io/documentation/components/navbar/#navbar-burger",
	b.Content("Nothing specific here. The burger is automatically generated."),
).Subsection(
	"Navbar menu",
	"https://bulma.io/documentation/components/navbar/#navbar-menu",
	b.Content(
		el.P("JavaScript is already included, no addition needed for the base use-case. If you need to trigger the change, you can add an ID with ", el.Code("html.ID"), " and use the following:"),
		el.Pre(`const navbar = document.getElementById("your-id")
navbar.getElementsByClassName("navbar-burger")[0].classList.toggle("is-active")
navbar.getElementsByClassName("navbar-menu")[0].classList.toggle("is-active")`),
	),
).Subsection(
	"Navbar start and navbar end",
	"https://bulma.io/documentation/components/navbar/#navbar-start-and-navbar-end",
	b.Content(
		b.UList(
			[]any{"left section: ", el.Code("b.NavbarStart")},
			[]any{"right section: ", el.Code("b.NavbarEnd")},
		),
	),
).Subsection(
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
		"https://bulma.io/assets/images/bulma-logo.png",
		html.Width("112"), html.Height("28"),
		html.Alt("Bulma"),
	),
)`),
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
).Subsection(
	"Transparent navbar",
	"https://bulma.io/documentation/components/navbar/#transparent-navbar",
	c.HorizontalExample(
		`b.Navbar(
	b.Transparent,
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			b.ImgSrc("https://bulma.io/assets/images/bulma-logo.svg"),
		),
	),
	b.NavbarStart(

		b.NavbarAHref("https://bulma.io/", "Home"),
		b.NavbarDropdown(
			b.OnLink(
				el.A,
				html.Href("https://bulma.io/documentation/overview/start/"),
				"Docs",
			),
			b.Active,
			b.Boxed,
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
		b.Navbar(
			b.Transparent,
			b.NavbarBrand(
				b.NavbarAHref(
					"https://bulma.io",
					b.ImgSrc("https://bulma.io/assets/images/bulma-logo.svg"),
				),
			),
			b.NavbarStart(

				b.NavbarAHref("https://bulma.io/", "Home"),
				b.NavbarDropdown(
					b.OnLink(
						el.A,
						html.Href("https://bulma.io/documentation/overview/start/"),
						"Docs",
					),
					b.Active,
					b.Boxed,
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
		),
	),
).Subsection(
	"Fixed navbar",
	"https://bulma.io/documentation/components/navbar/#fixed-navbar",
	b.Content(
		el.P("Add ", el.Code("b.FixedTop"), " or ", el.Code("b.FixedBottom"), " to the ", el.Code("b.Navbar"), " arguments."),
	),
).Subsection(
	"Dropdown menu",
	"https://bulma.io/documentation/components/navbar/#dropdown-menu",
	c.Example(
		`b.Navbar(
	html.Aria("label", "dropdown navigation"),
	b.NavbarDropdown(
		"Docs",
		b.Hoverable,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarDropdown(
				"Docs",
				b.Hoverable,
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
	b.NavbarDropdown(
		"Docs",
		b.Clickable,
		b.NavbarAHref("#", "Overview"),
		b.NavbarAHref("#", "Elements"),
		b.NavbarAHref("#", "Components"),
		b.NavbarDivider(),
		b.NavbarItem("Version 0.9.4"),
	),
)`,
		b.Navbar(
			html.Aria("label", "dropdown navigation"),
			b.NavbarDropdown(
				"Docs",
				b.Clickable,
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
		b.NavbarDropdown(
			"Left",
			b.Clickable,
			b.NavbarAHref("#", "Overview"),
			b.NavbarAHref("#", "Elements"),
			b.NavbarAHref("#", "Components"),
			b.NavbarDivider(),
			b.NavbarItem("Version 0.9.4"),
		),
	),
	b.NavbarEnd(
		b.NavbarDropdown(
			"Right",
			b.Clickable,
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
				b.NavbarDropdown(
					"Left",
					b.Clickable,
					b.NavbarAHref("#", "Overview"),
					b.NavbarAHref("#", "Elements"),
					b.NavbarAHref("#", "Components"),
					b.NavbarDivider(),
					b.NavbarItem("Version 0.9.4"),
				),
			),
			b.NavbarEnd(
				b.NavbarDropdown(
					"Right",
					b.Clickable,
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
		b.NavbarDropdown(
			"Dropup",
			b.Clickable,
			b.Up,
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
				b.NavbarDropdown(
					"Dropup",
					b.Clickable,
					b.Up,
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
		el.P("You can remove the arrow in the items of the navbar by adding the ", el.Code("b.Arrowless"), " modifier to the ", el.Code("b.NavbarDropdown"), " element."),
	),
	c.Example(
		`b.NavbarDropdown(
	"Link without arrow",
	b.Hoverable,
	b.Arrowless,
	b.NavbarAHref("#", "Overview"),
	b.NavbarAHref("#", "Elements"),
	b.NavbarAHref("#", "Components"),
	b.NavbarDivider(),
	b.NavbarItem("Version 0.9.4"),
)`,
		b.NavbarDropdown(
			"Link without arrow",
			b.Hoverable,
			b.Arrowless,
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
			"https://bulma.io/assets/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	b.NavbarDropdown(
		"Docs",
		b.Active,
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
					"https://bulma.io/assets/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			b.NavbarDropdown(
				"Docs",
				b.Active,
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
			"https://bulma.io/assets/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	b.NavbarDropdown(
		"Docs",
		b.Active,
		b.Boxed,
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
					"https://bulma.io/assets/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			b.NavbarDropdown(
				"Docs",
				b.Active,
				b.Boxed,
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
			"https://bulma.io/assets/images/bulma-logo.png",
			html.Width("112"), html.Height("28"),
			html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
		),
	),
	b.NavbarDropdown(
		"Docs",
		b.Active,
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
					"https://bulma.io/assets/images/bulma-logo.png",
					html.Width("112"), html.Height("28"),
					html.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
				),
			),
			b.NavbarDropdown(
				"Docs",
				b.Active,
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
).Subsection(
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
	),
	demoNavbar(false, b.Primary),
	demoNavbar(true, b.Link),
	demoNavbar(false, b.Info),
	demoNavbar(false, b.Success),
	demoNavbar(false, b.Warning),
	demoNavbar(false, b.Danger),
	demoNavbar(true, b.Black),
	demoNavbar(true, b.Dark),
	demoNavbar(false, b.Light),
	demoNavbar(false, b.White),
).Subsection(
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
