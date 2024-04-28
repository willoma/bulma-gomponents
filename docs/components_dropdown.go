package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var dropdown = c.NewPage(
	"Dropdown", "Dropdown", "/dropdown",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Dropdown"), " constructor creates a dropdown menu. In order to make the dropdown menu open to the top of the button, use ", el.Code("b.Dropup"), " instead. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnDropdown(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="dropdown">`), " element"},

			el.Code("b.OnTrigger(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="dropdown-trigger">`), " element"},

			el.Code("b.OnMenu(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="dropdown-menu">`), " element"},

			el.Code("b.OnContent(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="dropdown-content">`), " element"},

			el.Code("b.DropdownButton(...)"),
			"Make this button (part of) the trigger",

			[]any{el.Code("b.DropdownItem(...)"), ", ", el.Code("b.DropdownAHref(...)"), " and ", el.Code("b.DropdownDivider()")},
			"Add the child to the content",

			el.Code("b.ID"),
			[]any{"Set the ID of the ", el.Code(`<div class="dropdown-menu">`), " element"},

			el.Code("b.Clickable"),
			"Make the menu open when the button is clicked (javascript is automatically added)",

			el.Code("b.Element"),
			[]any{"Wrap the child with the ", el.Code("b.DropdownItem"), " constructor and add it to the content"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the dropdown",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the content",

			el.Code("b.Active"),
			"Open the menu",

			el.Code("b.Hoverable"),
			"Make the menu open when the cursor hovers the button",

			el.Code("b.Up"),
			"Make the dropdown open to the top",

			el.Code("string"),
			[]any{"Create a dropdown button with this label"},
		),
		el.P("The following constructors allow providing content to the dropdown menu:"),
		b.DList(
			el.Code("b.DropdownItem(children ...any)"),
			"Wrap the children in a dropdown-item div",

			el.Code("b.DropdownAHref(href string, children ...any)"),
			"Create a dropdown-item link",

			el.Code("b.DropdownDivider()"),
			"Create a dropdown divider",

			el.Code(`b.DropdownButton(children ...any)`),
			"Create a dropdown button with the provided content - if no icon is provided, a default icon will be added: angle down or up",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/dropdown/",

	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.OnTrigger(
		b.Button(
			html.Aria("haspopup", "true"),
			html.Aria("controls", "dropdown-menu"),
			"Dropdown button",
			fa.Icon(fa.Solid, "angle-down", b.Small),
		),
	),
	b.ID("dropdown-menu"),
	b.DropdownAHref("#", "Dropdown item"),
	b.DropdownItem(html.A, "Other dropdown item"),
	b.DropdownAHref("#", b.Active, "Active dropdown item"),
	b.DropdownAHref("#", "Other dropdown item"),
	b.DropdownDivider(),
	b.DropdownAHref("#", "With a divider"),
	b.Active,
)`,
		b.Dropdown(
			b.Clickable,
			b.OnTrigger(
				b.Button(
					html.Aria("haspopup", "true"),
					html.Aria("controls", "dropdown-menu"),
					"Dropdown button",
					fa.Icon(fa.Solid, "angle-down", b.Small),
				),
			),
			b.ID("dropdown-menu"),
			b.DropdownAHref("#", "Dropdown item"),
			b.DropdownItem(html.A, "Other dropdown item"),
			b.DropdownAHref("#", b.Active, "Active dropdown item"),
			b.DropdownAHref("#", "Other dropdown item"),
			b.DropdownDivider(),
			b.DropdownAHref("#", "With a divider"),
			b.Active,
		),
	),
).Subsection(
	"Dropdown content",
	"https://bulma.io/documentation/components/dropdown/#dropdown-content",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " automatically embeds any content that is not a legit dropdown item or divider into a ", el.Code("b.DropdownItem"), ":")),
	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.DropdownButton(
		html.Aria("controls", "dropdown-menu2"),
		"Content",
	),
	html.ID("dropdown-menu2"),
	el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
	b.DropdownDivider(),
	el.P("You simply need to use a ", el.Code("b.DropdownItem"), ", or even no specific container, instead."),
	b.DropdownDivider(),
	b.DropdownAHref("#", "This is a link"),
	b.Active,
)`,
		b.Dropdown(
			b.Clickable,
			b.DropdownButton(
				html.Aria("controls", "dropdown-menu2"),
				"Content",
			),
			html.ID("dropdown-menu2"),
			el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
			b.DropdownDivider(),
			el.P("You simply need to use a ", el.Code("b.DropdownItem"), ", or even no specific container, instead."),
			b.DropdownDivider(),
			b.DropdownAHref("#", "This is a link"),
			b.Active,
		),
	),
).Subsection(
	"Hoverable or Toggable",
	"https://bulma.io/documentation/components/dropdown/#hoverable-or-toggable",
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the", el.Code("easy.ClickableDropdown"), " and ", el.Code("easy.HoverableDropdown"), "variants:")),
	c.Example(
		`b.Dropdown(
	b.Clickable,
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
		b.Dropdown(
			b.Clickable,
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
		`b.Dropdown(
	b.Hoverable,
	"Hover me",
	el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
)`,
		b.Dropdown(
			b.Hoverable,
			"Hover me",
			el.P("You can insert ", el.Strong("any type of content"), " within the dropdown menu."),
		),
	),
).Subsection(
	"Right aligned",
	"https://bulma.io/documentation/components/dropdown/#right-aligned",

	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.DropdownButton("Left aligned"),
	el.P("The dropdown is ", el.Strong("left-aligned"), " by default."),
)`,
		b.Dropdown(
			b.Clickable,
			b.DropdownButton("Left aligned"),
			el.P("The dropdown is ", el.Strong("left-aligned"), " by default."),
		),
	),
	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.Right,
	b.DropdownButton("Right aligned"),
	el.P("Add the ", el.Code("b.Right"), " modifier for a ", el.Strong("right-aligned"), " dropdown."),
)`,
		b.Dropdown(
			b.Clickable,
			b.Right,
			b.DropdownButton("Right aligned"),
			el.P("Add the ", el.Code("b.Right"), " modifier for a ", el.Strong("right-aligned"), " dropdown."),
		),
	),
).Subsection(
	"Dropup",
	"https://bulma.io/documentation/components/dropdown/#dropup",

	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.Up,
	el.P("You can use the ", el.Code("b.Up"), " modifier to have a dropdown menu that appears above the dropdown button."),
)`,
		b.Dropdown(
			b.Clickable,
			b.Up,
			b.DropdownButton("Dropup button"),
			el.P("You can use the ", el.Code("b.Up"), " modifier to have a dropdown menu that appears above the dropdown button."),
		),
	),
)
