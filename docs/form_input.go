package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var formInput = c.NewPage(
	"Input", "Input", "/form/input",
	"https://bulma.io/documentation/form/input/",
	c.Example(
		`b.InputText(html.Placeholder("Text input"))`,
		b.InputText(html.Placeholder("Text input")),
	),
).Section(
	"Colors",
	"https://bulma.io/documentation/form/input/#colors",
	c.Example(
		`b.InputText(b.Primary, html.Placeholder("Primary input"))`,
		b.InputText(b.Primary, html.Placeholder("Primary input")),
	),
	c.Example(
		`b.InputText(b.Link, html.Placeholder("Link input"))`,
		b.InputText(b.Link, html.Placeholder("Link input")),
	),
	c.Example(
		`b.InputText(b.Info, html.Placeholder("Info input"))`,
		b.InputText(b.Info, html.Placeholder("Info input")),
	),
	c.Example(
		`b.InputText(b.Success, html.Placeholder("Success input"))`,
		b.InputText(b.Success, html.Placeholder("Success input")),
	),
	c.Example(
		`b.InputText(b.Warning, html.Placeholder("Warning input"))`,
		b.InputText(b.Warning, html.Placeholder("Warning input")),
	),
	c.Example(
		`b.InputText(b.Danger, html.Placeholder("Danger input"))`,
		b.InputText(b.Danger, html.Placeholder("Danger input")),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/form/input/#sizes",
	c.Example(
		`b.InputText(b.Small, html.Placeholder("Small input"))`,
		b.InputText(b.Small, html.Placeholder("Small input")),
	),
	c.Example(
		`b.InputText(b.Normal, html.Placeholder("Normal input"))`,
		b.InputText(b.Normal, html.Placeholder("Normal input")),
	),
	c.Example(
		`b.InputText(b.Medium, html.Placeholder("Medium input"))`,
		b.InputText(b.Medium, html.Placeholder("Medium input")),
	),
	c.Example(
		`b.InputText(b.Large, html.Placeholder("Large input"))`,
		b.InputText(b.Large, html.Placeholder("Large input")),
	),
).Section(
	"Styles",
	"https://bulma.io/documentation/form/input/#styles",
	c.Example(
		`b.InputText(b.Rounded, html.Placeholder("Rounded input"))`,
		b.InputText(b.Rounded, html.Placeholder("Rounded input")),
	),
).Section(
	"States",
	"https://bulma.io/documentation/form/input/#states",
	c.Example(
		`b.Control(
	b.InputText(html.Placeholder("Normal input")),
)`,
		b.Control(
			b.InputText(html.Placeholder("Normal input")),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(b.Hovered, html.Placeholder("Hovered input")),
)`,
		b.Control(
			b.InputText(b.Hovered, html.Placeholder("Hovered input")),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(b.Focused, html.Placeholder("Focused input")),
)`,
		b.Control(
			b.InputText(b.Focused, html.Placeholder("Focused input")),
		),
	),
	c.Example(
		`b.Control(
	b.Loading,
	b.InputText(html.Placeholder("Loading input")),
)`,
		b.Control(
			b.Loading,
			b.InputText(html.Placeholder("Loading input")),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.Small, b.Loading,
		b.InputText(b.Small, html.Placeholder("Small loading input")),
	),
),
b.Field(
	b.Control(
		b.Loading,
		b.InputText(html.Placeholder("Normal loading input")),
	),
),
b.Field(
	b.Control(
		b.Medium, b.Loading,
		b.InputText(b.Medium, html.Placeholder("Medium loading input")),
	),
),
b.Field(
	b.Control(
		b.Large, b.Loading,
		b.InputText(b.Large, html.Placeholder("Large loading input")),
	),
)`,
		b.Field(
			b.Control(
				b.Small, b.Loading,
				b.InputText(b.Small, html.Placeholder("Small loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Loading,
				b.InputText(html.Placeholder("Normal loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Medium, b.Loading,
				b.InputText(b.Medium, html.Placeholder("Medium loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Large, b.Loading,
				b.InputText(b.Large, html.Placeholder("Large loading input")),
			),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(html.Placeholder("Disabled input"), html.Disabled()),
)`,
		b.Control(
			b.InputText(html.Placeholder("Disabled input"), html.Disabled()),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(html.Value("This text is readonly"), html.ReadOnly()),
)`,
		b.Control(
			b.InputText(html.Value("This text is readonly"), html.ReadOnly()),
		),
	),
	c.Example(
		`b.Field(
	b.Horizontal,
	b.FieldLabel(
		b.Normal,
		b.Label("From"),
	),
	b.FieldBody(
		b.Field(
			b.Control(
				b.InputEmail(
					b.Static,
					html.Value("me@example.com"),
					html.ReadOnly(),
				),
			),
		),
	),
),
b.Field(
	b.Horizontal,
	b.FieldLabel(
		b.Normal,
		b.Label("To"),
	),
	b.FieldBody(
		b.Field(
			b.Control(
				b.InputEmail(html.Placeholder("Recipient email")),
			),
		),
	),
)`,
		b.Field(
			b.Horizontal,
			b.FieldLabel(
				b.Normal,
				b.Label("From"),
			),
			b.Field(
				b.Control(
					b.InputEmail(
						b.Static,
						html.Value("me@example.com"),
						html.ReadOnly(),
					),
				),
			),
		),
		b.Field(
			b.Horizontal,
			b.FieldLabel(
				b.Normal,
				b.Label("To"),
			),
			b.Field(
				b.Control(
					b.InputEmail(html.Placeholder("Recipient email")),
				),
			),
		),
	),
).Section(
	"With Font Awesome icons",
	"https://bulma.io/documentation/form/input/#with-font-awesome-icons",
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(html.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft,
		b.InputPassword(html.Placeholder("Password")),
		fa.Icon(fa.Solid, "lock", b.Small, b.Left),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(html.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft,
				b.InputPassword(html.Placeholder("Password")),
				fa.Icon(fa.Solid, "lock", b.Small, b.Left),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Small, html.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Small, html.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(html.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(html.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Medium, html.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Left),
		fa.Icon(fa.Solid, "check", b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Medium, html.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Left),
				fa.Icon(fa.Solid, "check", b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, html.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
		fa.Icon(fa.Solid, "check", b.Medium, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, html.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
				fa.Icon(fa.Solid, "check", b.Medium, b.Right),
			),
		),
	),
)
