package docs

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var formGeneral = c.NewPage(
	"General", "Form controls", "/form",
	"",

	b.Content(
		e.P("The ", e.Code("b.Field"), " constructor creates a field."),
		c.Modifiers(
			c.Row("b.Addons", "Attach the child controls together"),
			c.Row("b.AddonsCentered", "Attach the child controls together and center them"),
			c.Row("b.AddonsRight", "Attach the child controls together and align them to the right"),
			c.Row("b.Grouped", "Group child controls together"),
			c.Row("b.GroupedCentered", "Group child controls together and center them"),
			c.Row("b.GroupedRight", "Group child controls together and align them to the right"),
			c.Row("b.GroupedMultiline", "Group child controls together and allow them to fill up multiple lines"),
			c.Row("b.Horizontal", "Make the field horizontal"),
		),

		e.P("The ", e.Code("b.Label"), " constructor creates a label."),

		e.P("The ", e.Code("b.Control"), " constructor creates a control element."),
		c.Modifiers(
			c.Row("b.IconsLeft", "Leave space on the left side of the input or select child for an icon - the ", e.Code("b.Icon"), " element must be a direct child of ", e.Code("b.Control"), " and have the ", e.Code("b.Left"), " modifier"),
			c.Row("b.IconsRight", "Leave space on the right side of the input or select child for an icon - the ", e.Code("b.Icon"), " element must be a direct child of ", e.Code("b.Control"), " and have the ", e.Code("b.Right"), " modifier"),
			c.Row("b.Expanded", "Expand the control to fill the available width - to apply this style to a ", e.Code("<select>"), " element, you must also add the ", e.Code("b.FullWidth"), " modifier to the ", e.Code("b.Select"), " constructor"),
		),
		e.P("The ", e.Code("b.Help"), " constructor creates a help element, to be used as a child of a ", e.Code("b.Field"), "."),
		e.P("The ", e.Code("b.FieldHorizontal"), " constructor creates a horizontal field."),
		c.Children(
			c.Row("b.OnField(...any)", "Apply children to the ", e.Code(`<div class="field">`), " element"),
			c.Row("b.OnLabel(...any)", "Apply children to the ", e.Code(`<div class="field-label">`), " element"),
			c.Row("b.OnBody(...any)", "Apply children to the ", e.Code(`<div class="field-body">`), " element"),
			c.Row("b.Label(...any)", "Add label to the ", e.Code(`<div class="field-label">`), " element"),
			c.Row("e.Element", "Add element to the ", e.Code(`<div class="field-body">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="field">`), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/general/",
).Subsection(
	"Complete form example",
	"https://bulma.io/documentation/form/general/#complete-form-example",
	c.Example(
		`b.Field(
	b.Label("Name"),
	b.Control(
		b.InputText(e.Placeholder("Text input")),
	),
),
b.Field(
	b.Label("Username"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputText(
			b.Success,
			e.Placeholder("Text input"),
			e.Value("bulma"),
		),
		fa.Icon(fa.Solid, "user", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
	b.Help(b.Success, "This username is available"),
),
b.Field(
	b.Label("Email"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(
			b.Danger,
			e.Placeholder("Email input"),
			e.Value("hello@"),
		),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "exclamation-triangle", b.Small, b.Right),
	),
	b.Help(b.Danger, "This email is invalid"),
),
b.Field(
	b.Label("Subject"),
	b.Control(
		b.Select(
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
),
b.Field(
	b.Label("Message"),
	b.Control(
		b.Textarea(e.Placeholder("Textarea")),
	),
),
b.Field(
	b.Control(
		b.Label(
			"I agree to the ", e.AHref("#", "terms and conditions"),
		),
	),
),
b.Field(
	b.Control(
		b.Radio(
			e.Name("question"),
			"Yes",
		),
		b.Radio(
			e.Name("question"),
			"No",
		),
	),
),
b.Field(
	b.Grouped,
	b.Control(
		b.Button(b.Link, "Submit"),
	),
	b.Control(
		b.Button(b.Link, b.Light, "Cancel"),
	),
),`,
		b.Field(
			b.Label("Name"),
			b.Control(
				b.InputText(e.Placeholder("Text input")),
			),
		),
		b.Field(
			b.Label("Username"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputText(
					b.Success,
					e.Placeholder("Text input"),
					e.Value("bulma"),
				),
				fa.Icon(fa.Solid, "user", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
			b.Help(b.Success, "This username is available"),
		),
		b.Field(
			b.Label("Email"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(
					b.Danger,
					e.Placeholder("Email input"),
					e.Value("hello@"),
				),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "exclamation-triangle", b.Small, b.Right),
			),
			b.Help(b.Danger, "This email is invalid"),
		),
		b.Field(
			b.Label("Subject"),
			b.Control(
				b.Select(
					b.Option("", "Select dropdown"),
					b.Option("", "With options"),
				),
			),
		),
		b.Field(
			b.Label("Message"),
			b.Control(
				b.Textarea(e.Placeholder("Textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Label(
					"I agree to the ", e.AHref("#", "terms and conditions"),
				),
			),
		),
		b.Field(
			b.Control(
				b.Radio(
					e.Name("question"),
					"Yes",
				),
				b.Radio(
					e.Name("question"),
					"No",
				),
			),
		),
		b.Field(
			b.Grouped,
			b.Control(
				b.Button(b.Link, "Submit"),
			),
			b.Control(
				b.Button(b.Link, b.Light, "Cancel"),
			),
		),
	),
).Subsection(
	"Form field",
	"https://bulma.io/documentation/form/general/#form-field",
	c.Example(
		`b.Field(
	b.Label("Label"),
	b.Control(
		b.InputText(e.Placeholder("Text input")),
	),
	b.Help("This is a help text"),
)`,
		b.Field(
			b.Label("Label"),
			b.Control(
				b.InputText(e.Placeholder("Text input")),
			),
			b.Help("This is a help text"),
		),
	),
	c.Example(
		`b.Field(
	b.Label("Name"),
	b.Control(
		b.InputText(e.Placeholder("e.g Alex Smith")),
	),
),
b.Field(
	b.Label("Email"),
	b.Control(
		b.InputEmail(e.Placeholder("e.g. alexsmith@gmail.com")),
	),
)`,
		b.Field(
			b.Label("Name"),
			b.Control(
				b.InputText(e.Placeholder("e.g Alex Smith")),
			),
		),
		b.Field(
			b.Label("Email"),
			b.Control(
				b.InputEmail(e.Placeholder("e.g. alexsmith@gmail.com")),
			),
		),
	),
).Subsection(
	"Form control",
	"https://bulma.io/documentation/form/general/#form-control",
	c.Example(
		`b.Control(
	b.InputText(e.Placeholder("Text input")),
)`,
		b.Control(
			b.InputText(e.Placeholder("Text input")),
		),
	),
	c.Example(
		`b.Control(
	b.Select(
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
)`,
		b.Control(
			b.Select(
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
		),
	),
	c.Example(
		`b.Control(
	b.Button(b.Primary, "Submit"),
)`,
		b.Control(
			b.Button(b.Primary, "Submit"),
		),
	),
).Subsection(
	"With icons",
	"https://bulma.io/documentation/form/general/#with-icons",
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
),
b.Field(
	b.Control(
		b.Button(b.Success, "Login"),
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
		b.Field(
			b.Control(
				b.Button(b.Success, "Login"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.IconsLeft,
		b.Select(
			b.OptionSelected("", "Country"),
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
		fa.Icon(fa.Solid, "globe", b.Small, b.Left),
	),
)`,
		b.Field(
			b.Control(
				b.IconsLeft,
				b.Select(
					b.Option("", b.Selected, "Country"),
					b.Option("", "Select dropdown"),
					b.Option("", "With options"),
				),
				fa.Icon(fa.Solid, "globe", b.Small, b.Left),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Label(b.Small, "Small input"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Small, e.Placeholder("Normal")),
		fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
		fa.Icon(fa.Solid, "check", b.Small, b.Right),
	),
)`,
		b.Field(
			b.Label(b.Small, "Small input"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Small, e.Placeholder("Normal")),
				fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
				fa.Icon(fa.Solid, "check", b.Small, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Label("Normal input"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(e.Placeholder("Extra small")),
		fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(e.Placeholder("Normal")),
		fa.Icon(fa.Solid, "envelope", b.Left),
		fa.Icon(fa.Solid, "check", b.Right),
	),
)`,
		b.Field(
			b.Label("Normal input"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(e.Placeholder("Extra small")),
				fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(e.Placeholder("Normal")),
				fa.Icon(fa.Solid, "envelope", b.Left),
				fa.Icon(fa.Solid, "check", b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Label(b.Medium, "Medium input"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Medium, e.Placeholder("Extra small")),
		fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Medium, e.Placeholder("Small")),
		fa.Icon(fa.Solid, "envelope", fa.SizeSm, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeSm, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Medium, e.Placeholder("Normal")),
		fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
		fa.Icon(fa.Solid, "check", b.Medium, b.Right),
	),
)`,
		b.Field(
			b.Label(b.Medium, "Medium input"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Medium, e.Placeholder("Extra small")),
				fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Medium, e.Placeholder("Small")),
				fa.Icon(fa.Solid, "envelope", fa.SizeSm, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeSm, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Medium, e.Placeholder("Normal")),
				fa.Icon(fa.Solid, "envelope", b.Medium, b.Left),
				fa.Icon(fa.Solid, "check", b.Medium, b.Right),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Label(b.Large, "Large input"),
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, e.Placeholder("Extra small")),
		fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, e.Placeholder("Small")),
		fa.Icon(fa.Solid, "envelope", fa.SizeSm, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeSm, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, e.Placeholder("Normal")),
		fa.Icon(fa.Solid, "envelope", b.Large, b.Left),
		fa.Icon(fa.Solid, "check", b.Large, b.Right),
	),
),
b.Field(
	b.Control(
		b.IconsLeft, b.IconsRight,
		b.InputEmail(b.Large, e.Placeholder("Large")),
		fa.Icon(fa.Solid, "envelope", fa.SizeLg, b.Medium, b.Left),
		fa.Icon(fa.Solid, "check", fa.SizeLg, b.Medium, b.Right),
	),
)`,
		b.Field(
			b.Label(b.Large, "Large input"),
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, e.Placeholder("Extra small")),
				fa.Icon(fa.Solid, "envelope", fa.SizeXs, b.Small, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeXs, b.Small, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, e.Placeholder("Small")),
				fa.Icon(fa.Solid, "envelope", fa.SizeSm, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeSm, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, e.Placeholder("Normal")),
				fa.Icon(fa.Solid, "envelope", b.Large, b.Left),
				fa.Icon(fa.Solid, "check", b.Large, b.Right),
			),
		),
		b.Field(
			b.Control(
				b.IconsLeft, b.IconsRight,
				b.InputEmail(b.Large, e.Placeholder("Large")),
				fa.Icon(fa.Solid, "envelope", fa.SizeLg, b.Medium, b.Left),
				fa.Icon(fa.Solid, "check", fa.SizeLg, b.Medium, b.Right),
			),
		),
	),
).Subsection(
	"Form addons",
	"https://bulma.io/documentation/form/general/#form-addons",
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		b.InputText(e.Placeholder("Find a repository")),
	),
	b.Control(
		b.ButtonA(b.Info, "Search"),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				b.InputText(e.Placeholder("Find a repository")),
			),
			b.Control(
				b.ButtonA(b.Info, "Search"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		b.InputText(e.Placeholder("Your email")),
	),
	b.Control(
		b.ButtonA(b.Static, "@gmail.com"),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				b.InputText(e.Placeholder("Your email")),
			),
			b.Control(
				b.ButtonA(b.Static, "@gmail.com"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		b.Select(
			html.Span,
			b.Option("", "$"),
			b.Option("", "£"),
			b.Option("", "€"),
		),
	),
	b.Control(
		b.InputText(e.Placeholder("Amount of money")),
	),
	b.Control(
		b.ButtonA("Transfer"),
	),
),
b.Field(
	b.Addons,
	b.Control(
		b.Select(
			html.Span,
			b.Option("", "$"),
			b.Option("", "£"),
			b.Option("", "€"),
		),
	),
	b.Control(
		b.Expanded,
		b.InputText(e.Placeholder("Amount of money")),
	),
	b.Control(
		b.ButtonA("Transfer"),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				b.Select(
					html.Span,
					b.Option("", "$"),
					b.Option("", "£"),
					b.Option("", "€"),
				),
			),
			b.Control(
				b.InputText(e.Placeholder("Amount of money")),
			),
			b.Control(
				b.ButtonA("Transfer"),
			),
		),
		b.Field(
			b.Addons,
			b.Control(
				b.Select(
					html.Span,
					b.Option("", "$"),
					b.Option("", "£"),
					b.Option("", "€"),
				),
			),
			b.Control(
				b.Expanded,
				b.InputText(e.Placeholder("Amount of money")),
			),
			b.Control(
				b.ButtonA("Transfer"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		b.Expanded,
		b.Select(
			b.FullWidth,
			e.Name("country"),
			b.Option("Argentina", "Argentina"),
			b.Option("Bolivia", "Bolivia"),
			b.Option("Brazil", "Brazil"),
			b.Option("Chile", "Chile"),
			b.Option("Colombia", "Colombia"),
			b.Option("Ecuador", "Ecuador"),
			b.Option("Guyana", "Guyana"),
			b.Option("Paraguay", "Paraguay"),
			b.Option("Peru", "Peru"),
			b.Option("Suriname", "Suriname"),
			b.Option("Uruguay", "Uruguay"),
			b.Option("Venezuela", "Venezuela"),
		),
	),
	b.Control(
		b.ButtonSubmit(b.Primary, "Choose"),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				b.Expanded,
				b.Select(
					b.FullWidth,
					e.Name("country"),
					b.Option("Argentina", "Argentina"),
					b.Option("Bolivia", "Bolivia"),
					b.Option("Brazil", "Brazil"),
					b.Option("Chile", "Chile"),
					b.Option("Colombia", "Colombia"),
					b.Option("Ecuador", "Ecuador"),
					b.Option("Guyana", "Guyana"),
					b.Option("Paraguay", "Paraguay"),
					b.Option("Peru", "Peru"),
					b.Option("Suriname", "Suriname"),
					b.Option("Uruguay", "Uruguay"),
					b.Option("Venezuela", "Venezuela"),
				),
			),
			b.Control(
				b.ButtonSubmit(b.Primary, "Choose"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.AddonsCentered,
	b.Control(
		b.Select(
			html.Span,
			b.Option("", "$"),
			b.Option("", "£"),
			b.Option("", "€"),
		),
	),
	b.Control(
		b.InputText(e.Placeholder("Amount of money")),
	),
	b.Control(
		b.ButtonA(b.Primary, "Transfer"),
	),
)`,
		b.Field(
			b.AddonsCentered,
			b.Control(
				b.Select(
					html.Span,
					b.Option("", "$"),
					b.Option("", "£"),
					b.Option("", "€"),
				),
			),
			b.Control(
				b.InputText(e.Placeholder("Amount of money")),
			),
			b.Control(
				b.ButtonA(b.Primary, "Transfer"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.AddonsRight,
	b.Control(
		b.Select(
			html.Span,
			b.Option("", "$"),
			b.Option("", "£"),
			b.Option("", "€"),
		),
	),
	b.Control(
		b.InputText(e.Placeholder("Amount of money")),
	),
	b.Control(
		b.ButtonA(b.Primary, "Transfer"),
	),
)`,
		b.Field(
			b.AddonsRight,
			b.Control(
				b.Select(
					html.Span,
					b.Option("", "$"),
					b.Option("", "£"),
					b.Option("", "€"),
				),
			),
			b.Control(
				b.InputText(e.Placeholder("Amount of money")),
			),
			b.Control(
				b.ButtonA(b.Primary, "Transfer"),
			),
		),
	),
).Subsection(
	"Form group",
	"https://bulma.io/documentation/form/general/#form-group",
	c.Example(
		`b.Field(
	b.Grouped,
	b.Control(
		b.ButtonA(b.Primary, "Submit"),
	),
	b.Control(
		b.ButtonA(b.Light, "Cancel"),
	),
)`,
		b.Field(
			b.Grouped,
			b.Control(
				b.ButtonA(b.Primary, "Submit"),
			),
			b.Control(
				b.ButtonA(b.Light, "Cancel"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Grouped,
	b.GroupedCentered,
	b.Control(
		b.ButtonA(b.Primary, "Submit"),
	),
	b.Control(
		b.ButtonA(b.Light, "Cancel"),
	),
)`,
		b.Field(
			b.GroupedCentered,
			b.Control(
				b.ButtonA(b.Primary, "Submit"),
			),
			b.Control(
				b.ButtonA(b.Light, "Cancel"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Grouped,
	b.GroupedRight,
	b.Control(
		b.ButtonA(b.Primary, "Submit"),
	),
	b.Control(
		b.ButtonA(b.Light, "Cancel"),
	),
)`,
		b.Field(
			b.GroupedRight,
			b.Control(
				b.ButtonA(b.Primary, "Submit"),
			),
			b.Control(
				b.ButtonA(b.Light, "Cancel"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.Grouped,
	b.Control(
		b.Expanded,
		b.InputText(e.Placeholder("Find a repository")),
	),
	b.Control(
		b.ButtonA(b.Info, "Search"),
	),
)`,
		b.Field(
			b.Grouped,
			b.Control(
				b.Expanded,
				b.InputText(e.Placeholder("Find a repository")),
			),
			b.Control(
				b.ButtonA(b.Info, "Search"),
			),
		),
	),
	c.Example(
		`b.Field(
	b.GroupedMultiline,
	b.Control(b.ButtonA("One")),
	b.Control(b.ButtonA("Two")),
	b.Control(b.ButtonA("Three")),
	b.Control(b.ButtonA("Four")),
	b.Control(b.ButtonA("Five")),
	b.Control(b.ButtonA("Six")),
	b.Control(b.ButtonA("Seven")),
	b.Control(b.ButtonA("Eight")),
	b.Control(b.ButtonA("Nine")),
	b.Control(b.ButtonA("Ten")),
	b.Control(b.ButtonA("Eleven")),
	b.Control(b.ButtonA("Twelve")),
	b.Control(b.ButtonA("Thirteen")),
)`,
		b.Field(
			b.GroupedMultiline,
			b.Control(b.ButtonA("One")),
			b.Control(b.ButtonA("Two")),
			b.Control(b.ButtonA("Three")),
			b.Control(b.ButtonA("Four")),
			b.Control(b.ButtonA("Five")),
			b.Control(b.ButtonA("Six")),
			b.Control(b.ButtonA("Seven")),
			b.Control(b.ButtonA("Eight")),
			b.Control(b.ButtonA("Nine")),
			b.Control(b.ButtonA("Ten")),
			b.Control(b.ButtonA("Eleven")),
			b.Control(b.ButtonA("Twelve")),
			b.Control(b.ButtonA("Thirteen")),
		),
	),
).Subsection(
	"Horizontal form",
	"https://bulma.io/documentation/form/general/#horizontal-form",
	c.HorizontalExample(
		`b.FieldHorizontal(
	b.FieldLabel(b.Normal, b.Label("From")),
	b.Field(
		b.Control(
			html.P,
			b.Expanded,
			b.IconsLeft,
			b.InputText(e.Placeholder("Name")),
			fa.Icon(fa.Solid, "user", b.Small, b.Left),
		),
	),
	b.Field(
		b.Control(
			html.P,
			b.Expanded,
			b.IconsLeft,
			b.IconsRight,
			b.InputEmail(b.Success, e.Placeholder("Email"), e.Value("alex@smith.com")),
			fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
			fa.Icon(fa.Solid, "check", b.Small, b.Right),
		),
	),
),
b.FieldHorizontal(
	b.FieldLabel(),
	b.Field(
		b.Expanded,
		b.Field(
			b.Addons,
			b.Control(html.P, b.ButtonA(b.Static, "+44")),
			b.Control(html.P, b.Expanded, b.InputTel(e.Placeholder("Your phone number")))),
		b.Help("Do not enter the first zero"),
	),
),
b.FieldHorizontal(
	b.FieldLabel(b.Normal, b.Label("Department")),
	b.Field(
		b.Narrow,
		b.Control(b.Select(
			b.FullWidth,
			b.Option("", "Business development"),
			b.Option("", "Marketing"),
			b.Option("", "Sales"),
		)),
	),
),
b.FieldHorizontal(
	b.FieldLabel(b.Label("Already a member?")),
	b.Field(
		b.Narrow,
		b.Control(
			b.Radio(e.Name("member"), "Yes"),
			b.Radio(e.Name("member"), "No"),
		),
	),
),
b.FieldHorizontal(
	b.FieldLabel(b.Normal, b.Label("Subject")),
	b.Field(
		b.Control(b.InputText(
			b.Danger,
			e.Placeholder("e.g. Partnership opportunity"),
		)),
		b.Help(b.Danger, "This field is required"),
	),
),
b.FieldHorizontal(
	b.FieldLabel(b.Normal, b.Label("Question")),
	b.Field(
		b.Control(b.Textarea(e.Placeholder("Explain how we can help you"))),
	),
),
b.FieldHorizontal(
	b.FieldLabel(),
	b.Field(
		b.Control(b.Button(b.Primary, "Send message")),
	),
)`,
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("From"),
			b.Field(
				b.Control(
					html.P,
					b.Expanded,
					b.IconsLeft,
					b.InputText(e.Placeholder("Name")),
					fa.Icon(fa.Solid, "user", b.Small, b.Left),
				),
			),
			b.Field(
				b.Control(
					html.P,
					b.Expanded,
					b.IconsLeft,
					b.IconsRight,
					b.InputEmail(b.Success, e.Placeholder("Email"), e.Value("alex@smith.com")),
					fa.Icon(fa.Solid, "envelope", b.Small, b.Left),
					fa.Icon(fa.Solid, "check", b.Small, b.Right),
				),
			),
		),
		b.FieldHorizontal(
			b.Field(
				b.Expanded,
				b.Field(
					b.Addons,
					b.Control(html.P, b.ButtonA(b.Static, "+44")),
					b.Control(html.P, b.Expanded, b.InputTel(e.Placeholder("Your phone number")))),
				b.Help("Do not enter the first zero"),
			),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("Department"),
			b.Field(
				b.Narrow,
				b.Control(b.Select(
					b.FullWidth,
					b.Option("", "Business development"),
					b.Option("", "Marketing"),
					b.Option("", "Sales"),
				)),
			),
		),
		b.FieldHorizontal(
			b.Label("Already a member?"),
			b.Field(
				b.Narrow,
				b.Control(
					b.Radio(e.Name("member"), "Yes"),
					b.Radio(e.Name("member"), "No"),
				),
			),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("Subject"),
			b.Field(
				b.Control(b.InputText(
					b.Danger,
					e.Placeholder("e.g. Partnership opportunity"),
				)),
				b.Help(b.Danger, "This field is required"),
			),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("Question"),
			b.Field(
				b.Control(b.Textarea(e.Placeholder("Explain how we can help you"))),
			),
		),
		b.FieldHorizontal(
			b.Field(
				b.Control(b.Button(b.Primary, "Send message")),
			),
		),
	),
	c.HorizontalExample(
		`b.FieldHorizontal(
	b.FieldLabel(b.Label("No padding")),
	b.Field(b.Control(b.Checkbox("Checkbox"))),
),
b.FieldHorizontal(
	b.FieldLabel(b.Small, b.Label("Small padding")),
	b.Field(b.Control(b.InputText(b.Small, e.Placeholder("Small sized input")))),
),
b.FieldHorizontal(
	b.FieldLabel(b.Normal, b.Label("Normal label")),
	b.Field(b.Control(b.InputText(e.Placeholder("Normal sized input")))),
),
b.FieldHorizontal(
	b.FieldLabel(b.Medium, b.Label("Medium label")),
	b.Field(b.Control(b.InputText(b.Medium, e.Placeholder("Medium sized input")))),
),
b.FieldHorizontal(
	b.FieldLabel(b.Large, b.Label("Large label")),
	b.Field(b.Control(b.InputText(b.Large, e.Placeholder("Large sized input")))),
)`,
		b.FieldHorizontal(
			b.Label("No padding"),
			b.Field(b.Control(b.Checkbox("Checkbox"))),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Small),
			b.Label("Small padding"),
			b.Field(b.Control(b.InputText(b.Small, e.Placeholder("Small sized input")))),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Normal),
			b.Label("Normal label"),
			b.Field(b.Control(b.InputText(e.Placeholder("Normal sized input")))),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Medium),
			b.Label("Medium label"),
			b.Field(b.Control(b.InputText(b.Medium, e.Placeholder("Medium sized input")))),
		),
		b.FieldHorizontal(
			b.OnLabel(b.Large),
			b.Label("Large label"),
			b.Field(b.Control(b.InputText(b.Large, e.Placeholder("Large sized input")))),
		),
	),
).Subsection(
	"Disabled form",
	"https://bulma.io/documentation/form/general/#disabled-form",
	c.Example(
		`e.FieldSet(
	e.Disabled(),
	b.Field(
		b.Label("Name"),
		b.Control(
			b.InputText(e.Placeholder("e.g Alex Smith")),
		),
	),
	b.Field(
		b.Label("Email"),
		b.Control(
			b.InputEmail(e.Placeholder("e.g. alexsmith@gmail.com")),
		),
	),
)`,
		e.FieldSet(
			e.Disabled(),
			b.Field(
				b.Label("Name"),
				b.Control(
					b.InputText(e.Placeholder("e.g Alex Smith")),
				),
			),
			b.Field(
				b.Label("Email"),
				b.Control(
					b.InputEmail(e.Placeholder("e.g. alexsmith@gmail.com")),
				),
			),
		),
	),
)
