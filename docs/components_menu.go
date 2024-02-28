package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
)

var menu = c.NewPage(
	"Menu", "Menu", "/menu",
	"https://bulma.io/documentation/components/menu/",
	c.Example(
		`b.Menu(
	b.MenuLabel("General"),
	b.MenuList(
		b.MenuEntry(el.A("Dashboard")),
		b.MenuEntry(el.A("Customers")),
	),
	b.MenuLabel("Administration"),
	b.MenuList(
		b.MenuEntry(el.A("Team Settings")),
		b.MenuEntry(
			el.A(b.Active, "Manage Your Team"),
			b.MenuSublist(
				b.MenuEntry(el.A("Members")),
				b.MenuEntry(el.A("Plugins")),
				b.MenuEntry(el.A("Add a member")),
			),
		),
		b.MenuEntry(el.A("Invitations")),
		b.MenuEntry(el.A("Cloud Storage Environment Settings")),
		b.MenuEntry(el.A("Authentication")),
	),
	b.MenuLabel("Transactions"),
	b.MenuList(
		b.MenuEntry(el.A("Payments")),
		b.MenuEntry(el.A("Transfers")),
		b.MenuEntry(el.A("Balance")),
	),
)`,
		b.Menu(
			b.MenuLabel("General"),
			b.MenuList(
				b.MenuEntry(el.A("Dashboard")),
				b.MenuEntry(el.A("Customers")),
			),
			b.MenuLabel("Administration"),
			b.MenuList(
				b.MenuEntry(el.A("Team Settings")),
				b.MenuEntry(
					el.A(b.Active, "Manage Your Team"),
					b.MenuSublist(
						b.MenuEntry(el.A("Members")),
						b.MenuEntry(el.A("Plugins")),
						b.MenuEntry(el.A("Add a member")),
					),
				),
				b.MenuEntry(el.A("Invitations")),
				b.MenuEntry(el.A("Cloud Storage Environment Settings")),
				b.MenuEntry(el.A("Authentication")),
			),
			b.MenuLabel("Transactions"),
			b.MenuList(
				b.MenuEntry(el.A("Payments")),
				b.MenuEntry(el.A("Transfers")),
				b.MenuEntry(el.A("Balance")),
			),
		),
	),
	b.Content(
		el.P(el.Em("Bulma-Gomponents"), " provides the ", el.Code("easy.Menu"), " helper, which can be used when there is no specific need for elements modifiers:"),
	),
	c.Example(
		`easy.Menu(
	"General",
	el.A("Dashboard"),
	el.A("Customers"),
	"Administration",
	el.A("Team Settings"),
	easy.MenuEntryWithSublist(
		el.A(b.Active, "Manage Your Team"),
		el.A("Members"),
		el.A("Plugins"),
		el.A("Add a member"),
	),
	el.A("Invitations"),
	el.A("Cloud Storage Environment Settings"),
	el.A("Authentication"),
	"Transactions",
	el.A("Payments"),
	el.A("Transfers"),
	el.A("Balance"),
)`,
		easy.Menu(
			"General",
			el.A("Dashboard"),
			el.A("Customers"),
			"Administration",
			el.A("Team Settings"),
			easy.MenuEntryWithSublist(
				el.A(b.Active, "Manage Your Team"),
				el.A("Members"),
				el.A("Plugins"),
				el.A("Add a member"),
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
