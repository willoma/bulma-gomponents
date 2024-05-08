package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var formInput = c.NewPage(
	"Input", "Input", "/form/input",
	"",

	b.Content(
		e.P("The ", e.Code("b.InputText"), " constructor creates a text input. The ", e.Code("b.InputPassword"), " constructor creates a password input. The ", e.Code("b.InputEmail"), " constructor creates an email input. The ", e.Code("b.InputTel"), " constructor creates a telephone number input."),
		c.Modifiers(
			c.Row("b.Rounded", "Make the input rounded"),
			c.Row("b.Hovered", "Apply the hovered style"),
			c.Row("b.Focused", "Apply the focused style"),
			c.Row("b.Loading", "Add a loading spinner to the right of the input"),
			c.Row("b.Static", "Remove specific styling but maintain vertical spacing"),
			c.Row("b.Disabled", "Disable the input (apply the ", e.Code("disabled"), " attribute)"),
			c.Row("e.ReadOnly()", "Read only input"),
			c.Row("e.Placeholder(string)", "Add a placeholder to the input"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Normal", "Set size to normal"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),
		c.Children(
			c.Row("b.OnInput(...any)", "Apply children to the ", e.Code("<input>"), " element - only useful with the ", e.Code("b.Disabled"), " modifier if you want to apply the ", e.Code("is-disabled"), " class to the input instead of using the ", e.Code("disabled"), " attribute"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/input/",
	c.Example(
		`b.InputText(e.Placeholder("Text input"))`,
		b.InputText(e.Placeholder("Text input")),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/form/input/#colors",
	c.Example(
		`b.InputText(b.Primary, e.Placeholder("Primary input"))`,
		b.InputText(b.Primary, e.Placeholder("Primary input")),
	),
	c.Example(
		`b.InputText(b.Link, e.Placeholder("Link input"))`,
		b.InputText(b.Link, e.Placeholder("Link input")),
	),
	c.Example(
		`b.InputText(b.Info, e.Placeholder("Info input"))`,
		b.InputText(b.Info, e.Placeholder("Info input")),
	),
	c.Example(
		`b.InputText(b.Success, e.Placeholder("Success input"))`,
		b.InputText(b.Success, e.Placeholder("Success input")),
	),
	c.Example(
		`b.InputText(b.Warning, e.Placeholder("Warning input"))`,
		b.InputText(b.Warning, e.Placeholder("Warning input")),
	),
	c.Example(
		`b.InputText(b.Danger, e.Placeholder("Danger input"))`,
		b.InputText(b.Danger, e.Placeholder("Danger input")),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/form/input/#sizes",
	c.Example(
		`b.InputText(b.Small, e.Placeholder("Small input"))`,
		b.InputText(b.Small, e.Placeholder("Small input")),
	),
	c.Example(
		`b.InputText(b.Normal, e.Placeholder("Normal input"))`,
		b.InputText(b.Normal, e.Placeholder("Normal input")),
	),
	c.Example(
		`b.InputText(b.Medium, e.Placeholder("Medium input"))`,
		b.InputText(b.Medium, e.Placeholder("Medium input")),
	),
	c.Example(
		`b.InputText(b.Large, e.Placeholder("Large input"))`,
		b.InputText(b.Large, e.Placeholder("Large input")),
	),
).Subsection(
	"Styles",
	"https://bulma.io/documentation/form/input/#styles",
	c.Example(
		`b.InputText(b.Rounded, e.Placeholder("Rounded input"))`,
		b.InputText(b.Rounded, e.Placeholder("Rounded input")),
	),
).Subsection(
	"States",
	"https://bulma.io/documentation/form/input/#states",
	c.Example(
		`b.Control(
	b.InputText(e.Placeholder("Normal input")),
)`,
		b.Control(
			b.InputText(e.Placeholder("Normal input")),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(b.Hovered, e.Placeholder("Hovered input")),
)`,
		b.Control(
			b.InputText(b.Hovered, e.Placeholder("Hovered input")),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(b.Focused, e.Placeholder("Focused input")),
)`,
		b.Control(
			b.InputText(b.Focused, e.Placeholder("Focused input")),
		),
	),
	c.Example(
		`b.Control(
	b.Loading,
	b.InputText(e.Placeholder("Loading input")),
)`,
		b.Control(
			b.Loading,
			b.InputText(e.Placeholder("Loading input")),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.Small, b.Loading,
		b.InputText(b.Small, e.Placeholder("Small loading input")),
	),
),
b.Field(
	b.Control(
		b.Loading,
		b.InputText(e.Placeholder("Normal loading input")),
	),
),
b.Field(
	b.Control(
		b.Medium, b.Loading,
		b.InputText(b.Medium, e.Placeholder("Medium loading input")),
	),
),
b.Field(
	b.Control(
		b.Large, b.Loading,
		b.InputText(b.Large, e.Placeholder("Large loading input")),
	),
)`,
		b.Field(
			b.Control(
				b.Small, b.Loading,
				b.InputText(b.Small, e.Placeholder("Small loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Loading,
				b.InputText(e.Placeholder("Normal loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Medium, b.Loading,
				b.InputText(b.Medium, e.Placeholder("Medium loading input")),
			),
		),
		b.Field(
			b.Control(
				b.Large, b.Loading,
				b.InputText(b.Large, e.Placeholder("Large loading input")),
			),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(e.Placeholder("Disabled input"), b.Disabled),
)`,
		b.Control(
			b.InputText(e.Placeholder("Disabled input"), b.Disabled),
		),
	),
	c.Example(
		`b.Control(
	b.InputText(e.Value("This text is readonly"), e.ReadOnly()),
)`,
		b.Control(
			b.InputText(e.Value("This text is readonly"), e.ReadOnly()),
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
					e.Value("me@example.com"),
					e.ReadOnly(),
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
				b.InputEmail(e.Placeholder("Recipient email")),
			),
		),
	),
)`,
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("From"),
			b.Field(
				b.Control(
					b.InputEmail(
						b.Static,
						e.Value("me@example.com"),
						e.ReadOnly(),
					),
				),
			),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("To"),
			b.Field(
				b.Control(
					b.InputEmail(e.Placeholder("Recipient email")),
				),
			),
		),
	),
).Subsection(
	"With Font Awesome icons",
	"https://bulma.io/documentation/form/input/#with-font-awesome-icons",
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(e.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft,
		b.InputPassword(e.Placeholder("Password")),
		fa.Icon(fa.Solid, "lock", b.Small, b.Left),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(e.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft,
				b.InputPassword(e.Placeholder("Password")),
				fa.Icon(fa.Solid, "lock", b.Small, b.Left),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Small, e.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Small, e.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(e.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(e.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Medium, e.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Left),
		fa.Icon(fa.Solid, "check", b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Medium, e.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Left),
				fa.Icon(fa.Solid, "check", b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, e.Placeholder("Email")),
		fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
		fa.Icon(fa.Solid, "check", b.Medium, b.Right),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, e.Placeholder("Email")),
				fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
				fa.Icon(fa.Solid, "check", b.Medium, b.Right),
			),
		),
	),
)
