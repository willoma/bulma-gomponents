package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var menu = c.NewPage(
	"Menu", "Menu", "/menu",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Menu"), " constructor creates a menu.",
		),
		c.Children(
			c.Row("b.OnMenu(...any)", "Apply children to the ", e.Code(`<aside class="menu">`), " element"),
			c.Row("string", "Add string as a menu label"),
			c.Row("b.MenuLabel(...any)", "Add a menu label"),
			c.Row("b.MenuEntry(...any)", "Add a menu entry to the current menu list, creating it if it does not exist"),
			c.RowNodeAttribute("Apply attribute to the menu"),
			c.RowNodeElement("Add element to the current menu list, creating it if it does not exist"),
			c.Row("e.Element", "Add element to the current menu list, creating it if it does not exist"),
			c.RowDefault("Apply child to the ", e.Code(`<aside class="menu">`), " element"),
		),
		e.P("The ", e.Code("b.MenuLabel"), " constructor creates a menu label."),
		e.P("The ", e.Code("b.MenuList"), " constructor creates a menu list."),
		e.P("The ", e.Code("b.MenuEntry"), " constructor creates a menu entry."),
		c.Children(
			c.Row("b.OnLi(...any)", "Apply children to the ", e.Code("<li>"), " element"),
			c.Row("b.MenuEntry(...any)", "Add a menu entry to the current submenu, creating it if it does not exist"),
			c.Row("b.MenuSublist(...any)", "Add a submenu to the current menu entry"),
			c.RowDefault("Apply child to the ", e.Code("<li>"), " element"),
		),
		e.P(
			"The ", e.Code("b.MenuAHref"), " constructor creates an entry for a menu list with a link wrapped inside it.",
		),
		c.Modifiers(
			c.Row("b.Active", "Mark the link as active"),
		),
		c.Children(
			c.Row("b.OnLi(...any)", "Apply children to the ", e.Code("<li>"), " element"),
			c.Row("b.OnA(...any)", "Apply children to the ", e.Code("<a>"), " element"),
			c.RowClassesStyles("Apply child to the ", e.Code("<li>"), " element"),
			c.RowNodeAttribute("Apply attribute to the link"),
			c.RowNodeElement("Add element to the entry"),
			c.RowDefault("Apply child to the ", e.Code("<li>"), " element"),
		),
		e.P("The ", e.Code("b.MenuSublist"), " constructor creates a submenu."),
		e.P(
			"When manually adding a link inside a menu entry, add ", e.Code("b.Active"), " to it in order to mark it as active.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/menu/",
	c.Example(
		`b.Menu(
	"General",
	e.A("Dashboard"),
	e.A("Customers"),
	"Administration",
	e.A("Team Settings"),
	b.MenuEntry(
		e.A(b.Active, "Manage Your Team"),
		b.MenuEntry(e.A("Members")),
		b.MenuEntry(e.A("Plugins")),
		b.MenuEntry(e.A("Add a member")),
	),
	e.A("Invitations"),
	e.A("Cloud Storage Environment Settings"),
	e.A("Authentication"),
	"Transactions",
	e.A("Payments"),
	e.A("Transfers"),
	e.A("Balance"),
)`,
		b.Menu(
			"General",
			e.A("Dashboard"),
			e.A("Customers"),
			"Administration",
			e.A("Team Settings"),
			b.MenuEntry(
				e.A(b.Active, "Manage Your Team"),
				b.MenuEntry(e.A("Members")),
				b.MenuEntry(e.A("Plugins")),
				b.MenuEntry(e.A("Add a member")),
			),
			e.A("Invitations"),
			e.A("Cloud Storage Environment Settings"),
			e.A("Authentication"),
			"Transactions",
			e.A("Payments"),
			e.A("Transfers"),
			e.A("Balance"),
		),
	),
)
