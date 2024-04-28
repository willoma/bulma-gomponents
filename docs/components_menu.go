package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var menu = c.NewPage(
	"Menu", "Menu", "/menu",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Menu"), " constructor creates a menu. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnMenu(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<aside class="menu">`), " element"},

			el.Code("string"),
			[]any{"Add the string as a menu label"},

			el.Code("b.MenuLabel(...)"),
			[]any{"Add a menu label"},

			el.Code("b.MenuEntry(...)"),
			[]any{"Add a menu entry to the current menu list, creating it if it does not exist"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the menu",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the current menu list, creating it if it does not exist",

			el.Code("b.Element"),
			"Add this element to the current menu list, creating it if it does not exist",
		),
		el.P("Other children are added to the menu element."),
		el.P("The ", el.Code("b.MenuLabel"), " constructor creates a menu label."),
		el.P("The ", el.Code("b.MenuList"), " constructor creates a menu list."),
		el.P("The ", el.Code("b.MenuEntry"), " constructor creates a menu entry. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnLi(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<li>"), " element"},

			el.Code("b.MenuEntry(...)"),
			[]any{"Add a menu entry to the current submenu, creating it if it does not exist"},

			el.Code("b.MenuSublist(...)"),
			[]any{"Add a submenu to the current menu entry"},
		),
		el.P(
			"The ", el.Code("b.MenuAHref"), " constructor creates an entry for a menu list with a link wrapped inside it. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.OnLi(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<li>"), " element"},

			el.Code("b.OnA(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<a>"), " element"},

			el.Code("b.Active"),
			"Mark the link as active",

			"Any other class or style",
			"Apply to the entry element",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the link",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the entry",
		),
		el.P("The ", el.Code("b.MenuSublist"), " constructor creates a submenu."),
		el.P("Other children are added to the menu entry element."),
		el.P(
			"Add ", el.Code("b.Active"), " to a ", el.Code("<a>"), " link in a ", el.Code("b.MenuEntry"), " element in order to make it display as active.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/menu/",
	c.Example(
		`b.Menu(
	"General",
	el.A("Dashboard"),
	el.A("Customers"),
	"Administration",
	el.A("Team Settings"),
	b.MenuEntry(
		el.A(b.Active, "Manage Your Team"),
		b.MenuEntry(el.A("Members")),
		b.MenuEntry(el.A("Plugins")),
		b.MenuEntry(el.A("Add a member")),
	),
	el.A("Invitations"),
	el.A("Cloud Storage Environment Settings"),
	el.A("Authentication"),
	"Transactions",
	el.A("Payments"),
	el.A("Transfers"),
	el.A("Balance"),
)`,
		b.Menu(
			"General",
			el.A("Dashboard"),
			el.A("Customers"),
			"Administration",
			el.A("Team Settings"),
			b.MenuEntry(
				el.A(b.Active, "Manage Your Team"),
				b.MenuEntry(el.A("Members")),
				b.MenuEntry(el.A("Plugins")),
				b.MenuEntry(el.A("Add a member")),
			),
			el.A("Invitations"),
			el.A("Cloud Storage Environment Settings"),
			el.A("Authentication"),
			"Transactions",
			el.A("Payments"),
			el.A("Transfers"),
			el.A("Balance"),
		),
	),
)
