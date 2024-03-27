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
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Dropdown"), " constructor creates a dropdown menu. Its first child must be a button element and its second child must be a ", el.Code("b.DromdownMenu"), " element. In order to make the dropdown menu open to the top of the button, use ", el.Code("b.Dropup"), ". These constructors accept the following values additionally to the standard set of children types:",
		),
		b.DList(
			el.Code("b.Active"),
			"Open the menu",
			el.Code("b.Hoverable"),
			"Make the menu open when the cursor hovers the button",
		),
		el.P(
			"The ", el.Code("b.DropdownMenu"), " constructor creates the menu part of a dropdown. It accepts the following values additionally to the standard set of children types:",
		),
		b.DList(
			el.Code("b.Inner(any)"),
			[]any{"forcibly apply the child to the ", el.Code(`<div class="dropdown-content">`), " element"},

			el.Code("b.Outer(any)"),
			[]any{"forcibly apply the child to the ", el.Code(`<div class="dropdown-menu">`), " element"},

			[]any{el.Code("b.DropdownItem(...)"), ", ", el.Code("b.DropdownAHref(...)"), " and ", el.Code("b.DropdownDivider()")},
			"Add the child to the content",

			el.Code("b.Element"),
			[]any{"Wrap the child with the ", el.Code("b.DropdownItem"), " constructor and add it to the content"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the menu",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the content",
		),
		el.P("The following constructors allow providing content to the menu:"),
		b.DList(
			el.Code("b.DropdownItem(children ...any)"),
			"Wrap the children in a dropdown-item div",

			el.Code("b.DropdownAHref(href string, children ...any)"),
			"Create a dropdown-item link",

			el.Code("b.DropdownDivider()"),
			"Create a dropdown divider",
		),
	),
).Section(
	"Easy helper", "",

	b.Content(
		el.P("The ", el.Code("easy.ClickableDropdown"), ", ", el.Code("easy.HoverableDropdown"), ", ", el.Code("easy.ClickableDropup"), " and ", el.Code("easy.HoverableDropup"), " constructors build dropdown menus with the provided test on the trigger button and accept the following values additionally to the standard set of children types:"),
		b.DList(
			el.Code("b.Inner(any)"),
			[]any{"Forcibly apply the child to the ", el.Code(`<div class="dropdown-menu">`), " element"},

			el.Code("b.Outer(any)"),
			[]any{"Forcibly apply the child to the ", el.Code(`<div class="dropdown">`), " element"},

			"classes and styles",
			"Apply to the dropdown element",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the dropdown element",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the menu element",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/dropdown/",

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
).Subsection(
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
).Subsection(
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
).Subsection(
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
).Subsection(
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
