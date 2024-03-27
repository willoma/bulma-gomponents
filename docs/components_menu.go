package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
)

var menu = c.NewPage(
	"Menu", "Menu", "/menu",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Menu"), " constructor creates a menu. It accepts the standard set of children types, and needs a set of ", el.Code("b.MenuLabel"), " and ", el.Code("b.MenuList"), " children in order to build its content.",
		),
		el.P(
			"The ", el.Code("b.MenuList"), " constructor creates a list in a menu. It accepts the standard set of children types, and needs a set of ", el.Code("b.MenuEntry"), " or ", el.Code("b.MenuAHref"), " children in order to build its content. The ", el.Code("b.MenuSublist"), " constructor creates a sublist in a menu entry, and  needs a set of ", el.Code("b.MenuEntry"), " children in order to build its content.",
		),
		el.P(
			"Add ", el.Code("b.Active"), " to a ", el.Code("<a>"), " link in a ", el.Code("b.MenuEntry"), " element in order to make it display as active.",
		),
		el.P(
			"The ", el.Code("b.MenuAHref"), " constructor creates an entry for a menu list with a link wrapped inside it. It accepts the following values additionally to the standard set of children types:",
		),
		b.DList(
			el.Code("b.Inner(any)"),
			[]any{"forcibly apply the child to the link element"},

			el.Code("b.Outer(any)"),
			[]any{"forcibly apply the child to the entry element"},

			el.Code("b.Active"),
			"Mark the link as active",

			"Any other class or style",
			"Apply to the entry element",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the link",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the entry",
		),
	),
).Section(
	"Easy helper", "",

	b.Content(
		el.P(
			"The ", el.Code("easy.Menu"), " constructor creates a menu. It accepts the following values:",
		),
		b.DList(
			el.Code("string"),
			"Append the string as a menu label",

			el.Code("[]any"),
			[]any{"Wrap the slice content in a single ", el.Code("b.MenuEntry"), " and append it"},

			"Any other value",
			[]any{"Wrap the child in a ", el.Code("b.MenuEntry"), " and append it"},
		),
		el.P(
			"Do not provide a ", el.Code("b.MenuEntry"), " as a child to ", el.Code("easy.Menu"), ", otherwise it would be embedded into another ", el.Code("b.MenuEntry"), ".",
		),
		el.P(
			"The ", el.Code("easy.MenyEntryWithSublist"), " function returns the appropriate content for an entry with a sublist, to be provided to ", el.Code("easy.Menu"), ".",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/menu/",
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
