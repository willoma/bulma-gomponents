package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var panel = c.NewPage(
	"Panel", "Panel", "/panel",
	"",

	b.Content(
		el.P("The ", el.Code("b.Panel"), " constructor creates a panel. The following children have a special meaning:"),
		b.DList(
			el.Code("b.Primary"),
			"Set panel color to primary",

			el.Code("b.Link"),
			"Set panel color to link",

			el.Code("b.Info"),
			"Set panel color to info",

			el.Code("b.Success"),
			"Set panel color to success",

			el.Code("b.Warning"),
			"Set panel color to warning",

			el.Code("b.Danger"),
			"Set panel color to danger",
		),
		el.P("The ", el.Code("b.PanelHeading"), " constructor creates a panel heading."),
		el.P("The ", el.Code("b.PanelBlock"), " constructor creates a panel block."),
		el.P("The ", el.Code("b.PanelLink"), " and ", el.Code("b.PanelAHref"), " constructors create panel blocks which are ", el.Code("<a>"), " elements."),
		el.P("The ", el.Code("b.PanelLabel"), " constructor creates a panel block which is a ", el.Code("<label>"), " element."),
		el.P("The ", el.Code("b.PanelCheckbox"), " constructor creates a panel block which is a ", el.Code("<label>"), " element, containing a checkbox. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnLabel(...)"),
			[]any{"Force children to be applied to the ", el.Code(`<label class="panel-block">`), " element"},

			el.Code("b.OnInput(...)"),
			[]any{"Force children to be applied to the ", el.Code("<input>"), " element"},

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"Apply the attribute to the input element",

			[]any{"Other ", el.Code("gomponents.Node")},
			"Add this element to the block element",
		),
		el.P("Other children are added to the ", el.Code("<input>"), " element."),
		el.P("The ", el.Code("b.PanelTabs"), " constructor creates a panel tabs section. Its children must be ", el.Code("<a>"), " elements (for instance ", el.Code("b.AHref"), "). Add the ", el.Code("b.Active"), "modifier to a link to mark it as the active tab."),
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
			b.InputText(html.Placeholder("Search")),
			fa.Icon(fa.Solid, "search", b.Left),
		),
	),
	b.PanelTabs(
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
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
					b.InputText(html.Placeholder("Search")),
					fa.Icon(fa.Solid, "search", b.Left),
				),
			),
			b.PanelTabs(
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
		el.A(b.Active, "All"),
		el.A("Public"),
		el.A("Private"),
		el.A("Sources"),
		el.A("Forks"),
	),
	b.PanelBlock(
		b.Control(
			html.P,
			b.IconsLeft,
			b.InputText(html.Placeholder("Search")),
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
				el.A(b.Active, "All"),
				el.A("Public"),
				el.A("Private"),
				el.A("Sources"),
				el.A("Forks"),
			),
			b.PanelBlock(
				b.Control(
					html.P,
					b.IconsLeft,
					b.InputText(html.Placeholder("Search")),
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
