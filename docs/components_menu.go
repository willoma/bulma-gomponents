package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var menu = c.NewPage(
	"Menu", "Menu", "/menu",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Menu"), " constructor creates a menu. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnMenu(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<aside class="menu">`), " e.Element"},

			e.Code("string"),
			[]any{"Add the string as a menu label"},

			e.Code("b.MenuLabel(...)"),
			[]any{"Add a menu label"},

			e.Code("b.MenuEntry(...)"),
			[]any{"Add a menu entry to the current menu list, creating it if it does not exist"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the menu",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the current menu list, creating it if it does not exist",

			e.Code("e.Element"),
			"Add this e.Element to the current menu list, creating it if it does not exist",
		),
		e.P("Other children are added to the menu e.Element."),
		e.P("The ", e.Code("b.MenuLabel"), " constructor creates a menu label."),
		e.P("The ", e.Code("b.MenuList"), " constructor creates a menu list."),
		e.P("The ", e.Code("b.MenuEntry"), " constructor creates a menu entry. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnLi(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<li>"), " e.Element"},

			e.Code("b.MenuEntry(...)"),
			[]any{"Add a menu entry to the current submenu, creating it if it does not exist"},

			e.Code("b.MenuSublist(...)"),
			[]any{"Add a submenu to the current menu entry"},
		),
		e.P(
			"The ", e.Code("b.MenuAHref"), " constructor creates an entry for a menu list with a link wrapped inside it. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnLi(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<li>"), " e.Element"},

			e.Code("b.OnA(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<a>"), " e.Element"},

			e.Code("b.Active"),
			"Mark the link as active",

			"Any other class or style",
			"Apply to the entry e.Element",

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the link",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the entry",
		),
		e.P("The ", e.Code("b.MenuSublist"), " constructor creates a submenu."),
		e.P("Other children are added to the menu entry e.Element."),
		e.P(
			"Add ", e.Code("b.Active"), " to a ", e.Code("<a>"), " link in a ", e.Code("b.MenuEntry"), " e.Element in order to make it display as active.",
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
