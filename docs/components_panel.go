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
	"https://bulma.io/documentation/components/panel/",
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
	b.PanelLabel(
		b.Checkbox("remember me"),
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
			b.PanelLabel(
				b.Checkbox("remember me"),
			),
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
).Section(
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
