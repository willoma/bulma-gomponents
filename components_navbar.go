package docs

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

func demoNavbar(white bool, color b.Color) e.Element {
	var logoSrc string
	if white {
		logoSrc = "https://bulma.io/assets/brand/Bulma%20Logo%20White.svg"
	} else {
		logoSrc = "https://bulma.io/assets/brand/Bulma%20Logo%20Black.svg"
	}
	return b.Navbar(
		e.Styles{"margin-bottom": "20rem"},
		color,
		b.NavbarBrand(
			b.NavbarAHref("https://bulma.io", e.ImgSrc(logoSrc)),
		),
		b.NavbarStart(
			b.NavbarAHref("https://bulma.io/", "Home"),
			b.NavbarDropdown(
				b.OnLink(
					e.A,
					e.Href("https://bulma.io/documentation/overview/start/"),
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
							e.DataAttr("social-network", "Twitter"),
							e.DataAttr("social-action", "tweet"),
							e.DataAttr("social-target", "https://bulma.io"),
							e.Target("_blank"),
							e.Href("https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=https://bulma.io&amp;via=jgthms"),
							fa.Icon(fa.Brand, "twitter"),
							"Tweet",
						),
					),
					b.Control(
						html.P,
						b.ButtonA(
							color.Light(),
							e.Href("https://github.com/jgthms/bulma/releases/download/0.9.4/bulma-0.9.4.zip"),
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
		e.P("The ", e.Code("b.Navbar"), " constructor creates a navbar. The brand section is added only if at least one brand, start or end child is provided. The section part and the burger button are added only if at least one start or end child is provided. The click event on the burger, which toggles the navbar menu on mobile devices, is already included."),
		c.Modifiers(
			c.Row("b.FixedTop", "Fix the navbar to the top of the page - the ", e.Code("has-navbar-fixed-top"), " class is automatically added to the parent element (which should be the body for this to work)"),
			c.Row("b.FixedBottom", "Fix the navbar to the bottom of the page - the ", e.Code("has-navbar-fixed-bottom"), " class is automatically added to the parent element (which should be the body for this to work)"),
			c.Row("b.Transparent", "Remove hover and active backgrounds from the items"),
			c.Row("b.Spaced", "Increase padding"),
			c.Row("b.Shadow", "Add a small shadow around the navbar"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.White", "Set color to white"),
		),
		c.Children(
			c.Row("b.NavbarBrand(...any)", "Apply children to the brand section (left side, always visible)"),
			c.Row("b.NavbarStart(...any)", "Apply children to the start section (left part of the menu)"),
			c.Row("b.NavbarEnd(...any)", "Apply children to the end section (right part of the menu)"),
			c.Row("b.Container(...any)", "Use element as an intermediate container"),
			c.RowDefault("Apply children to the ", e.Code(`<nav class="navbar">`), " element"),
		),
		e.P("The ", e.Code("b.NavbarItem"), " constructor creates a navbar item for a navbar brand, start or end section, or for a ", e.Code("b.NavbarDropdown"), "."),
		c.Modifiers(
			c.Row("b.Expanded", "Turn the item into a full-width element"),
			c.Row("b.Tab", "Add a bottom border on hover, always show the bottom border when adding ", e.Code("b.Active")),
			c.Row("b.HasDropdown", "Make the item a link+dropdown container"),
			c.Row("b.HasDropup", "Make the item a dropdown container, with the dropdown opening above the link"),
			c.Row("b.Hoverable", "Make the included dropdown automatically show on hover"),
			c.Row("b.Active", "Force the dropdown to be open"),
		),

		e.P("The ", e.Code("b.NavbarAHref"), " constructor creates a link item for a navbar brand, start or end section, or to a ", e.Code("b.NavbarDropdown"), "."),
		c.Modifiers(
			c.Row("b.Expanded", "Turn the item into a full-width element"),
			c.Row("b.Tab", "Add a bottom border on hover, always show the bottom border when adding ", e.Code("b.Active")),
		),
		e.P("The ", e.Code("b.NavbarDropdown"), " constructor creates navbar item with a dropdown menu."),
		c.Modifiers(
			c.Row("b.Hoverable", "Make the menu open when the cursor hovers the trigger"),
			c.Row("b.Clickable", "Make the menu open when the button is clicked (javascript is automatically added)"),
			c.Row("b.Up", "Make the dropdown open to the top"),
			c.Row("b.Arrowless", "Remove the arrow in the item"),
			c.Row("b.Boxed", "Show the dropdown as a boxed element, without the top grey border"),
			c.Row("b.Active", "Force the dropdown to be open"),
			c.Row("b.Right", "Used when a dropdown is on the right side (end section), in order to open its content aligned to the right of the link"),
		),
		c.Children(
			c.Row("b.OnItem(...any)", "Apply children to the ", e.Code(`<div class="navbar-item">`), " element"),
			c.Row("b.OnLink(...any)", "Apply children to the ", e.Code(`<a class="navbar-link">`), " element"),
			c.Row("b.OnDropdown(...any)", "Apply children to the ", e.Code(`<div class="navbar-dropdown">`), " element"),
			c.Row("e.Class", "Apply class to the ", e.Code(`<div class="navbar-dropdown">`), " element"),
			c.Row("string", "Add string to the ", e.Code(`<a class="navbar-link">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="navbar-dropdown">`), " element"),
		),
		e.P("The ", e.Code("b.NavbarLink"), " constructor creates a link element, to include in a ", e.Code("b.NavbarItem"), " with the ", e.Code("b.HasDropdown"), " modifier."),
		c.Modifiers(
			c.Row("b.Arrowless", "Remove the arrow in the item"),
		),
		e.P("The ", e.Code("b.NavbarDivider"), " constructor creates a divider element, to include in a ", e.Code("b.NavbarDropdown"), "."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/navbar/",
).Subsection(
	"Basic Navbar",
	"https://bulma.io/documentation/components/navbar/#basic-navbar",
	c.HorizontalExample(
		`b.Navbar(
	e.AriaNavigation,
	e.AriaLabel("main navigation"),
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			e.ImgSrc(
				"https://bulma.io/assets/images/bulma-logo.png",
				e.Width("112"), e.Height("28"),
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
					e.Strong("Sign up"),
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
			e.AriaNavigation,
			e.AriaLabel("main navigation"),
			b.NavbarBrand(
				b.NavbarAHref(
					"https://bulma.io",
					e.ImgSrc(
						"https://bulma.io/assets/images/bulma-logo.png",
						e.Width("112"), e.Height("28"),
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
							e.Strong("Sign up"),
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
	e.AriaNavigation,
	e.AriaLabel("main navigation"),
	b.NavbarBrand(
		b.NavbarAHref(
			"https://bulma.io",
			e.ImgSrc(
				"https://bulma.io/assets/images/bulma-logo.png",
				e.Width("112"), e.Height("28"),
			),
		),
	),
)`,
		b.Navbar(
			e.AriaNavigation,
			e.AriaLabel("main navigation"),
			b.NavbarBrand(
				b.NavbarAHref(
					"https://bulma.io",
					e.ImgSrc(
						"https://bulma.io/assets/images/bulma-logo.png",
						e.Width("112"), e.Height("28"),
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
		e.P("JavaScript is already included, no addition needed for the base use-case. If you need to trigger the change, you can add an ID with ", e.Code("e.ID"), " and use the following:"),
		e.Pre(`const navbar = document.getElementById("your-id")
navbar.getElementsByClassName("navbar-burger")[0].classList.toggle("is-active")
navbar.getElementsByClassName("navbar-menu")[0].classList.toggle("is-active")`),
	),
).Subsection(
	"Navbar start and navbar end",
	"https://bulma.io/documentation/components/navbar/#navbar-start-and-navbar-end",
	b.Content(
		b.UList(
			[]any{"left section: ", e.Code("b.NavbarStart")},
			[]any{"right section: ", e.Code("b.NavbarEnd")},
		),
	),
).Subsection(
	"Navbar item",
	"https://bulma.io/documentation/components/navbar/#navbar-item",
	b.Content(
		e.Ul(
			e.Li(
				"a navigation ", e.Strong("link"),
				e.Pre(`b.NavbarAHref("#", "Home")`),
			),
			e.Li(
				"a container for the ", e.Strong("brand logo"),
				e.Pre(`b.NavbarAHref(
	"#",
	e.ImgSrc(
		"https://bulma.io/assets/images/bulma-logo.png",
		e.Width("112"), e.Height("28"),
		e.Alt("Bulma"),
	),
)`),
			),
			e.Li(
				"a child of a ", e.Strong("navbar dropdown"),
				e.Pre(`b.NavbarDropdown(
	b.NavbarAHref("#", "Overview"),
)`),
			),
			e.Li(
				"a container for almost ", e.Strong("anything"), " you want, like a ", e.Code("field"),
				e.Pre(`b.NavbarItem(
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
		e.P("You can add the following ", e.Strong("modifier"), " classes:"),
		e.Ul(
			e.Li(e.Code("b.Expanded"), " to turn it into a full-width element"),
			e.Li(e.Code("b.Tab"), " to add a bottom border on hover and show the bottom border using ", e.Code("b.Active")),
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
			e.ImgSrc("https://bulma.io/assets/images/bulma-logo.svg"),
		),
	),
	b.NavbarStart(

		b.NavbarAHref("https://bulma.io/", "Home"),
		b.NavbarDropdown(
			b.OnLink(
				e.A,
				e.Href("https://bulma.io/documentation/overview/start/"),
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
						e.DataAttr("social-network", "Twitter"),
						e.DataAttr("social-action", "tweet"),
						e.DataAttr("social-target", "https://bulma.io"),
						e.Target("_blank"),
						e.Href("https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=https://bulma.io&amp;via=jgthms"),
						fa.Icon(fa.Brand, "twitter"),
						"Tweet",
					),
				),
				b.Control(
					html.P,
					b.ButtonA(
						b.Primary,
						e.Href("https://github.com/jgthms/bulma/releases/download/0.9.4/bulma-0.9.4.zip"),
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
					e.ImgSrc("https://bulma.io/assets/images/bulma-logo.svg"),
				),
			),
			b.NavbarStart(

				b.NavbarAHref("https://bulma.io/", "Home"),
				b.NavbarDropdown(
					b.OnLink(
						e.A,
						e.Href("https://bulma.io/documentation/overview/start/"),
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
								e.DataAttr("social-network", "Twitter"),
								e.DataAttr("social-action", "tweet"),
								e.DataAttr("social-target", "https://bulma.io"),
								e.Target("_blank"),
								e.Href("https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=https://bulma.io&amp;via=jgthms"),
								fa.Icon(fa.Brand, "twitter"),
								"Tweet",
							),
						),
						b.Control(
							html.P,
							b.ButtonA(
								b.Primary,
								e.Href("https://github.com/jgthms/bulma/releases/download/0.9.4/bulma-0.9.4.zip"),
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
		e.P("Add ", e.Code("b.FixedTop"), " or ", e.Code("b.FixedBottom"), " to the ", e.Code("b.Navbar"), " arguments."),
	),
).Subsection(
	"Dropdown menu",
	"https://bulma.io/documentation/components/navbar/#dropdown-menu",
	c.Example(
		`b.Navbar(
	e.AriaLabel("dropdown navigation"),
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
			e.AriaLabel("dropdown navigation"),
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
	e.AriaLabel("dropdown navigation"),
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
			e.AriaLabel("dropdown navigation"),
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
		e.H3("Right dropdown"),
		e.P("If your parent ", e.Code("b.NavbarItem"), " is on the right side, you can position the dropdown to start from the ", e.Strong("right"), " with the ", e.Code("b.Right"), " modifier."),
	),
	c.Example(
		`b.Navbar(
	e.AriaLabel("dropdown navigation"),
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
		"Everything you need to ", e.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			e.AriaLabel("dropdown navigation"),
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
				"Everything you need to ", e.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		e.H3("Dropup"),
		e.P("If you're using a navbar at the bottom, like the ", e.AHref("http://localhost:8080/navbar#fixed-navbar", "fixed bottom navbar"), ", you might want to use a ", e.Strong("dropup menu"), ". Simply add the ", e.Code("b.NavbarDropup"), " modifier to the parent ", e.Code("b.NavbarItem"), ":"),
	),
	c.Example(
		`b.Hero(
	b.Primary,
	b.Title(html.P, "Documentation"),
	b.Subtitle(
		html.P,
		"Everything you need to ", e.Strong("create a website"), " with Bulma",
	),
),
b.Navbar(
	e.AriaLabel("dropdown navigation"),
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
				"Everything you need to ", e.Strong("create a website"), " with Bulma",
			),
		),
		b.Navbar(
			e.AriaLabel("dropdown navigation"),
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
		e.H3("Dropdown without arrow"),
		e.P("You can remove the arrow in the items of the navbar by adding the ", e.Code("b.Arrowless"), " modifier to the ", e.Code("b.NavbarDropdown"), " element."),
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
		e.H3("Styles for the dropdown menu"),
	),
	c.Example(
		`b.Navbar(
	e.AriaLabel("dropdown navigation"),
	b.NavbarAHref(
		"#",
		e.ImgSrc(
			"https://bulma.io/assets/images/bulma-logo.png",
			e.Width("112"), e.Height("28"),
			e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
		"Everything you need to ", e.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			e.AriaLabel("dropdown navigation"),
			b.NavbarAHref(
				"#",
				e.ImgSrc(
					"https://bulma.io/assets/images/bulma-logo.png",
					e.Width("112"), e.Height("28"),
					e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
				"Everything you need to ", e.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		e.P("When having a ", e.AHref("http://localhost:8080/navbar#transparent-navbar", "transparent navbar"), ", it is preferable to use the boxed version of the dropdown, by using the ", e.Code("b.NavbarDropdownBoxed"), " option."),
	),
	c.Example(
		`b.Navbar(
	e.AriaLabel("dropdown navigation"),
	b.NavbarAHref(
		"#",
		e.ImgSrc(
			"https://bulma.io/assets/images/bulma-logo.png",
			e.Width("112"), e.Height("28"),
			e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
		"Everything you need to ", e.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			e.AriaLabel("dropdown navigation"),
			b.NavbarAHref(
				"#",
				e.ImgSrc(
					"https://bulma.io/assets/images/bulma-logo.png",
					e.Width("112"), e.Height("28"),
					e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
				"Everything you need to ", e.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(e.H3("Active dropdown navbar item")),
	c.Example(
		`b.Navbar(
	e.AriaLabel("dropdown navigation"),
	b.NavbarAHref(
		"#",
		e.ImgSrc(
			"https://bulma.io/assets/images/bulma-logo.png",
			e.Width("112"), e.Height("28"),
			e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
		"Everything you need to ", e.Strong("create a website"), " with Bulma",
	),
)`,
		b.Navbar(
			e.AriaLabel("dropdown navigation"),
			b.NavbarAHref(
				"#",
				e.ImgSrc(
					"https://bulma.io/assets/images/bulma-logo.png",
					e.Width("112"), e.Height("28"),
					e.Alt("Bulma: Free, open source, and modern CSS framework based on Flexbox"),
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
				"Everything you need to ", e.Strong("create a website"), " with Bulma",
			),
		),
	),
	b.Content(
		e.H3(" Dropdown divider"),
		e.Pre("b.NavbarDivider()"),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/components/navbar/#colors",
	b.Content(
		e.P("In order to change the background color of the navbarn simply, apply one of the color modifiers to the ", e.Code("b.Navbar"), " Element:"),
		b.UList(
			e.Code("b.Primary"),
			e.Code("b.Link"),
			e.Code("b.Info"),
			e.Code("b.Success"),
			e.Code("b.Warning"),
			e.Code("b.Danger"),
			e.Code("b.Black"),
			e.Code("b.Dark"),
			e.Code("b.Light"),
			e.Code("b.White"),
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
		e.Ul(
			e.Li(
				e.Code("b.Spaced"), " sets ", e.Strong("Top"), " and ", e.Strong("Bottom"), " paddings with ", e.Strong("1rem"), ", ", e.Strong("Left"), " and ", e.Strong("Right"), " paddings with ", e.Strong("2rem"),
			),
			e.Li(
				e.Code("b.Shadow"), " adds a small amount of box-shadow around the navbar",
			),
		),
	),
)
