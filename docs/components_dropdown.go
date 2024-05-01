package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var dropdown = c.NewPage(
	"Dropdown", "Dropdown", "/dropdown",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Dropdown"), " constructor creates a dropdown menu. In order to make the dropdown menu open to the top of the button, use ", e.Code("b.Dropup"), " instead. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnDropdown(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="dropdown">`), " e.Element"},

			e.Code("b.OnTrigger(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="dropdown-trigger">`), " e.Element"},

			e.Code("b.OnMenu(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="dropdown-menu">`), " e.Element"},

			e.Code("b.OnContent(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="dropdown-content">`), " e.Element"},

			e.Code("b.DropdownButton(...)"),
			"Make this button (part of) the trigger",

			[]any{e.Code("b.DropdownItem(...)"), ", ", e.Code("b.DropdownAHref(...)"), " and ", e.Code("b.DropdownDivider()")},
			"Add the child to the content",

			e.Code("b.ID"),
			[]any{"Set the ID of the ", e.Code(`<div class="dropdown-menu">`), " e.Element"},

			e.Code("b.Clickable"),
			"Make the menu open when the button is clicked (javascript is automatically added)",

			e.Code("e.Element"),
			[]any{"Wrap the child with the ", e.Code("b.DropdownItem"), " constructor and add it to the content"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the dropdown",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the content",

			e.Code("b.Active"),
			"Open the menu",

			e.Code("b.Hoverable"),
			"Make the menu open when the cursor hovers the button",

			e.Code("b.Up"),
			"Make the dropdown open to the top",

			e.Code("string"),
			[]any{"Create a dropdown button with this label"},
		),
		e.P("The following constructors allow providing content to the dropdown menu:"),
		b.DList(
			e.Code("b.DropdownItem(children ...any)"),
			"Wrap the children in a dropdown-item div",

			e.Code("b.DropdownAHref(href string, children ...any)"),
			"Create a dropdown-item link",

			e.Code("b.DropdownDivider()"),
			"Create a dropdown divider",

			e.Code(`b.DropdownButton(children ...any)`),
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
			e.AriaHasPopupTrue,
			e.AriaControlsID("dropdown-menu"),
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
					e.AriaHasPopupTrue,
					e.AriaControlsID("dropdown-menu"),
					"Dropdown button",
					fa.Icon(fa.Solid, "angle-down", b.Small),
				),
			),
			e.ID("dropdown-menu"),
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
	b.Content(e.P(e.Em("Bulma-Gomponents"), " automatically embeds any content that is not a legit dropdown item or divider into a ", e.Code("b.DropdownItem"), ":")),
	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.DropdownButton(
		e.AriaControlsID("dropdown-menu2"),
		"Content",
	),
	e.ID("dropdown-menu2"),
	e.P("You can insert ", e.Strong("any type of content"), " within the dropdown menu."),
	b.DropdownDivider(),
	e.P("You simply need to use a ", e.Code("b.DropdownItem"), ", or even no specific container, instead."),
	b.DropdownDivider(),
	b.DropdownAHref("#", "This is a link"),
	b.Active,
)`,
		b.Dropdown(
			b.Clickable,
			b.DropdownButton(
				e.AriaControlsID("dropdown-menu2"),
				"Content",
			),
			e.ID("dropdown-menu2"),
			e.P("You can insert ", e.Strong("any type of content"), " within the dropdown menu."),
			b.DropdownDivider(),
			e.P("You simply need to use a ", e.Code("b.DropdownItem"), ", or even no specific container, instead."),
			b.DropdownDivider(),
			b.DropdownAHref("#", "This is a link"),
			b.Active,
		),
	),
).Subsection(
	"Hoverable or Toggable",
	"https://bulma.io/documentation/components/dropdown/#hoverable-or-toggable",
	b.Content(e.P(e.Em("Bulma-Gomponents"), " provides the", e.Code("easy.ClickableDropdown"), " and ", e.Code("easy.HoverableDropdown"), "variants:")),
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
	e.P("You can insert ", e.Strong("any type of content"), " within the dropdown menu."),
)`,
		b.Dropdown(
			b.Hoverable,
			"Hover me",
			e.P("You can insert ", e.Strong("any type of content"), " within the dropdown menu."),
		),
	),
).Subsection(
	"Right aligned",
	"https://bulma.io/documentation/components/dropdown/#right-aligned",

	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.DropdownButton("Left aligned"),
	e.P("The dropdown is ", e.Strong("left-aligned"), " by default."),
)`,
		b.Dropdown(
			b.Clickable,
			b.DropdownButton("Left aligned"),
			e.P("The dropdown is ", e.Strong("left-aligned"), " by default."),
		),
	),
	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.Right,
	b.DropdownButton("Right aligned"),
	e.P("Add the ", e.Code("b.Right"), " modifier for a ", e.Strong("right-aligned"), " dropdown."),
)`,
		b.Dropdown(
			b.Clickable,
			b.Right,
			b.DropdownButton("Right aligned"),
			e.P("Add the ", e.Code("b.Right"), " modifier for a ", e.Strong("right-aligned"), " dropdown."),
		),
	),
).Subsection(
	"Dropup",
	"https://bulma.io/documentation/components/dropdown/#dropup",

	c.Example(
		`b.Dropdown(
	b.Clickable,
	b.Up,
	e.P("You can use the ", e.Code("b.Up"), " modifier to have a dropdown menu that appears above the dropdown button."),
)`,
		b.Dropdown(
			b.Clickable,
			b.Up,
			b.DropdownButton("Dropup button"),
			e.P("You can use the ", e.Code("b.Up"), " modifier to have a dropdown menu that appears above the dropdown button."),
		),
	),
)
