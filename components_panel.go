package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var panel = c.NewPage(
	"Panel", "Panel", "/panel",
	"",

	b.Content(
		e.P("The ", e.Code("b.Panel"), " constructor creates a panel. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Primary"),
			"Set panel color to primary",

			e.Code("b.Link"),
			"Set panel color to link",

			e.Code("b.Info"),
			"Set panel color to info",

			e.Code("b.Success"),
			"Set panel color to success",

			e.Code("b.Warning"),
			"Set panel color to warning",

			e.Code("b.Danger"),
			"Set panel color to danger",
		),
		e.P("The ", e.Code("b.PanelHeading"), " constructor creates a panel heading."),
		e.P("The ", e.Code("b.PanelBlock"), " constructor creates a panel block."),
		e.P("The ", e.Code("b.PanelLink"), " and ", e.Code("b.PanelAHref"), " constructors create panel blocks which are ", e.Code("<a>"), " elements."),
		e.P("The ", e.Code("b.PanelLabel"), " constructor creates a panel block which is a ", e.Code("<label>"), " e.Element."),
		e.P("The ", e.Code("b.PanelCheckbox"), " constructor creates a panel block which is a ", e.Code("<label>"), " e.Element, containing a checkbox. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnLabel(...)"),
			[]any{"Force children to be applied to the ", e.Code(`<label class="panel-block">`), " e.Element"},

			e.Code("b.OnInput(...)"),
			[]any{"Force children to be applied to the ", e.Code("<input>"), " e.Element"},

			[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
			"Apply the attribute to the input e.Element",

			[]any{"Other ", e.Code("gomponents.Node")},
			"Add this e.Element to the block e.Element",
		),
		e.P("Other children are added to the ", e.Code("<input>"), " e.Element."),
		e.P("The ", e.Code("b.PanelTabs"), " constructor creates a panel tabs section. Its children must be ", e.Code("<a>"), " elements (for instance ", e.Code("e.AHref"), "). Add the ", e.Code("b.Active"), "modifier to a link to mark it as the active tab."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/panel/",
	c.Example(
		`b.Panel(
	b.PanelHeading("Repositories"),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "code-branch"),
		"daniellowtw/infboard",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "code-branch"),
		"mojs",
	),
	b.PanelCheckbox("remember me"),
	),
	b.PanelBlock(
		b.Button(
			b.Link,
			b.Outlined,
			b.FullWidth,
			"Reset all filters",
		),
	),
)`,
		b.Panel(
			b.PanelHeading("Repositories"),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "code-branch"),
				"daniellowtw/infboard",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "code-branch"),
				"mojs",
			),
			b.PanelCheckbox("remember me"),
			b.PanelBlock(
				b.Button(
					b.Link,
					b.Outlined,
					b.FullWidth,
					"Reset all filters",
				),
			),
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/components/panel/#colors",
	c.Example(
		`b.Panel(
	b.Primary,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Primary,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
	c.Example(
		`b.Panel(
	b.Link,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Link,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
	c.Example(
		`b.Panel(
	b.Info,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Info,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
	c.Example(
		`b.Panel(
	b.Success,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Success,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
	c.Example(
		`b.Panel(
	b.Warning,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Warning,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
	c.Example(
		`b.Panel(
	b.Danger,
	b.PanelHeading("Repositories"),
	b.PanelTabs(
		e.A(b.Active, "All"),
		e.A("Public"),
		e.A("Private"),
		e.A("Sources"),
		e.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(e.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelLink(
		b.Active,
		fa.Icon(fa.Solid, "book"),
		"bulma",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"marksheet",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"minireset.css",
	),
	b.PanelLink(
		fa.Icon(fa.Solid, "book"),
		"jgthms.github.io",
	),
)`,
		b.Panel(
			b.Danger,
			b.PanelHeading("Repositories"),
			b.PanelTabs(
				e.A(b.Active, "All"),
				e.A("Public"),
				e.A("Private"),
				e.A("Sources"),
				e.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(e.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelLink(
				b.Active,
				fa.Icon(fa.Solid, "book"),
				"bulma",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"marksheet",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"minireset.css",
			),
			b.PanelLink(
				fa.Icon(fa.Solid, "book"),
				"jgthms.github.io",
			),
		),
	),
)
