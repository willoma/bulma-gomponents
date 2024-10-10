package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var dropdown = c.NewPage(
	"Dropdown", "Dropdown", "/dropdown",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Dropdown"), " constructor creates a dropdown menu.",
		),
		c.Modifiers(
			c.Row("b.Active", "Open the menu"),
			c.Row("b.Clickable", "Make the menu open when the button is clicked (javascript is automatically added)"),
			c.Row("b.Hoverable", "Make the menu open when the cursor hovers the button"),
			c.Row("b.Up", "Make the dropdown open to the top"),
		),
		c.Children(
			c.Row("b.OnDropdown(...any)", "Apply children to the ", e.Code(`<div class="dropdown">`), " element"),
			c.Row("b.OnTrigger(...any)", "Apply children to the ", e.Code(`<div class="dropdown-trigger">`), " element"),
			c.Row("b.OnMenu(...any)", "Apply children to the ", e.Code(`<div class="dropdown-menu">`), " element"),
			c.Row("b.OnContent(...any)", "Apply children to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.Row("b.DropdownButton(...any)", "Make this button (part of) the trigger"),
			c.Row("b.DropdownDivider(...any)", "Apply divider to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.Row("b.DropdownItem(...any)", "Apply item to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.Row("b.DropdownAHref(...any)", "Apply link item to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.Row("b.ID", "Set the ID of the ", e.Code(`<div class="dropdown-menu">`), " element"),
			c.Row("e.Class", "Apply class to the ", e.Code(`<div class="dropdown">`), " element"),
			c.Row("string", "Create a dropdown button with this label"),
			c.Row("e.Element", "Wrap child with the ", e.Code("b.DropdownItem"), " constructor and add it to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.RowNodeAttribute("Apply child to the ", e.Code(`<div class="dropdown">`), " element"),
			c.RowNodeElement("Wrap child with the ", e.Code("b.DropdownItem"), " constructor and add it to the ", e.Code(`<div class="dropdown-content">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="dropdown">`), " element"),
		),
		e.P(
			"The ", e.Code("b.DropdownButton"), " constructor creates a button to be used as a dropdown trigger. It automatically adds a Font Awesome icon to the right if no icon is provided.",
		),
		e.P(
			"The ", e.Code("b.DropdownItem"), " constructor creates a dropdown item.",
		),
		e.P(
			"The ", e.Code("b.DropdownAHref"), " constructor creates a dropdown link item.",
		),
		e.P(
			"The ", e.Code("b.DropdownDivider"), " constructor creates a dropdown divider.",
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
