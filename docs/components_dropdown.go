package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var dropdown = c.NewPage(
	"Dropdown", "Dropdown", "/dropdown",
	"https://bulma.io/documentation/components/dropdown/",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the ", el.Code(`b.DropdownAHref`), " helper.")),
	c.Example(
		`b.Dropdown(
	b.Button(
		html.Aria("haspopup", "true"),
		html.Aria("controls", "dropdown-menu"),
		"Dropdown button",
		fa.Icon(fa.Solid, "angle-down", b.Small),
	),
	b.DropdownMenu(
		html.ID("dropdown-menu"),
		b.DropdownAHref("#", "Dropdown item"),
		b.DropdownItem(html.A, "Other dropdown item"),
		b.DropdownAHref("#", b.Active, "Active dropdown item"),
		b.DropdownAHref("#", "Other dropdown item"),
		b.DropdownDivider(),
		b.DropdownAHref("#", "With a divider"),
	),
	b.Active,
)`,
		b.Dropdown(
			b.Button(
				html.Aria("haspopup", "true"),
				html.Aria("controls", "dropdown-menu"),
				"Dropdown button",
				fa.Icon(fa.Solid, "angle-down", b.Small),
			),
			b.DropdownMenu(
				html.ID("dropdown-menu"),
				b.DropdownAHref("#", "Dropdown item"),
				b.DropdownItem(html.A, "Other dropdown item"),
				b.DropdownAHref("#", b.Active, "Active dropdown item"),
				b.DropdownAHref("#", "Other dropdown item"),
				b.DropdownDivider(),
				b.DropdownAHref("#", "With a divider"),
			),
			b.Active,
			b.OnClick(`this.classList.toggle("is-active")`),
		),
	),
).Section(
	"Dropdown content",
	"https://bulma.io/documentation/components/dropdown/#dropdown-content",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " automatically embeds any content that is not a legit dropdown item or divider into a ", el.Code("b.DropdownItem"), ":")),
	c.Example(
		`b.Dropdown(
	b.Button(
		html.Aria("haspopup", "true"),
		html.Aria("controls", "dropdown-menu2"),
		"Content",
		fa.Icon(
			fa.Solid, "angle-down",
			b.Small,
		),
	),
	b.DropdownMenu(
		html.ID("dropdown-menu2"),
		el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
		b.DropdownDivider(),
		el.P("You simply need to use a ", el.Code("b.DropdownItem"), ", or even no specific container, instead."),
		b.DropdownDivider(),
		b.DropdownAHref("#", "This is a link"),
	),
	b.Active,
)`,
		b.Dropdown(
			b.Button(
				html.Aria("haspopup", "true"),
				html.Aria("controls", "dropdown-menu2"),
				"Content",
				fa.Icon(
					fa.Solid, "angle-down",
					b.Small,
				),
			),
			b.DropdownMenu(
				html.ID("dropdown-menu2"),
				el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
				b.DropdownDivider(),
				el.P("You simply need to use a ", el.Code("b.DropdownItem"), ", or even no specific container, instead."),
				b.DropdownDivider(),
				b.DropdownAHref("#", "This is a link"),
			),
			b.Active,
			b.OnClick(`this.classList.toggle("is-active")`),
		),
	),
).Section(
	"Hoverable or Toggable",
	"https://bulma.io/documentation/components/dropdown/#hoverable-or-toggable",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the", el.Code("easy.ClickableDropdown"), " and ", el.Code("easy.HoverableDropdown"), "variants:")),
	c.Example(
		`easy.ClickableDropdown(
	"Click me",
	b.DropdownAHref("#", "Overview"),
	b.DropdownAHref("#", "Modifiers"),
	b.DropdownAHref("#", "Grid"),
	b.DropdownAHref("#", "Form"),
	b.DropdownAHref("#", "Elements"),
	b.DropdownAHref("#", "Components"),
	b.DropdownAHref("#", "Layout"),
	b.DropdownDivider(),
	b.DropdownAHref("#", "More"),
)`,
		easy.ClickableDropdown(
			"Click me",
			b.DropdownAHref("#", "Overview"),
			b.DropdownAHref("#", "Modifiers"),
			b.DropdownAHref("#", "Grid"),
			b.DropdownAHref("#", "Form"),
			b.DropdownAHref("#", "Elements"),
			b.DropdownAHref("#", "Components"),
			b.DropdownAHref("#", "Layout"),
			b.DropdownDivider(),
			b.DropdownAHref("#", "More"),
		),
	),
	c.Example(
		`easy.HoverableDropdown(
	"Hover me",
	el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
)`,
		easy.HoverableDropdown(
			"Hover me",
			el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
		),
	),
).Section(
	"Right aligned",
	"https://bulma.io/documentation/components/dropdown/#right-aligned",
	c.Example(
		`easy.ClickableDropdown(
	"Left aligned",
	el.P("The dropdown is ", el.Strong("left-aligned"), " by default."),
)`,
		easy.ClickableDropdown(
			"Left aligned",
			el.P("The dropdown is ", el.Strong("left-aligned"), " by default."),
		),
	),
	c.Example(
		`easy.ClickableDropdown(
	"Right aligned",
	b.Right,
	el.P("Add the ", el.Code("b.Right"), " modifier for a ", el.Strong("right-aligned"), " dropdown."),
)`,
		easy.ClickableDropdown(
			"Right aligned",
			b.Right,
			el.P("Add the ", el.Code("b.Right"), " modifier for a ", el.Strong("right-aligned"), " dropdown."),
		),
	),
).Section(
	"Dropup",
	"https://bulma.io/documentation/components/dropdown/#dropup",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the ", el.Code("b.Dropup"), ", ", el.Code("easy.ClickableDropup"), " and ", el.Code("easy.HoverableDropup"), " variants:")),
	c.Example(
		`easy.ClickableDropup(
	"Dropup button",
	el.P("You can use a ", el.Code("Dropup"), " variant to have a dropdown menu that appears above the dropdown button."),
)`,
		easy.ClickableDropup(
			"Dropup button",
			el.P("You can use a ", el.Code("*Dropup"), " variant to have a dropdown menu that appears above the dropdown button."),
		),
	),
)
