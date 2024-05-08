package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var button = c.NewPage(
	"Button", "Button", "/button",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Button"), ", ", e.Code("b.ButtonA"), ", ", e.Code("b.ButtonAHref"), ", ", e.Code("b.ButtonSubmit"), ", ", e.Code("b.ButtonInputSubmit"), " and ", e.Code("b.ButtonInputReset"), " constructors create buttons.",
		),
		c.Modifiers(
			b.Small,
			c.Row("b.Responsive", "Responsive size"),
			c.Row("b.FullWidth", "Take the whole width"),
			c.Row("b.Outlined", "Outline style"),
			c.Row("b.Inverted", "Inverted style"),
			c.Row("b.Rounded", "Rounded button"),
			c.Row("b.Hovered", "Apply the hovered style"),
			c.Row("b.Focused", "Apply the focused style"),
			c.Row("b.Active", "Apply the active style"),
			c.Row("b.Loading", "Replace the content with a loading spinner"),
			c.Row("b.Static", "Make the button non-interactive"),
			c.Row("e.Disabled()", "Disable the button"),
			c.Row("e.Disabled(bool)", "Disable the button conditionally"),
			c.Row("b.Selected", "In a list of attached buttons (Buttons with Addons), make sure this button is above the other buttons"),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Normal", "Set size to normal"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
			c.Row("b.White", "Set color to white"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Text", "Set style to underlined text"),
			c.Row("b.Ghost", "Set style to link-looking blue text"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.PrimaryLight", "Set color to primary light"),
			c.Row("b.LinkLight", "Set color to link light"),
			c.Row("b.InfoLight", "Set color to info light"),
			c.Row("b.SuccessLight", "Set color to success light"),
			c.Row("b.WarningLight", "Set color to warning light"),
			c.Row("b.DangerLight", "Set color to danger light"),
			c.Row("b.PrimaryDark", "Set color to primary dark"),
			c.Row("b.LinkDark", "Set color to link dark"),
			c.Row("b.InfoDark", "Set color to info dark"),
			c.Row("b.SuccessDark", "Set color to success dark"),
			c.Row("b.WarningDark", "Set color to warning dark"),
			c.Row("b.DangerDark", "Set color to danger dark"),
		),

		e.P("The ", e.Code("b.Buttons"), " constructor creates a button."),
		c.Modifiers(
			c.Row("b.Addons", "Attach the buttons together"),
			c.Row("b.Centered", "Center the buttons"),
			c.Row("b.Right", "Align the buttons to the right"),
			c.Row("b.Small", "Set the buttons size to small"),
			c.Row("b.Medium", "Set the buttons size to medium"),
			c.Row("b.Large", "Set the buttons size to large"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/button/",
	c.Example(
		`b.Button("Button")`,
		b.Button("Button"),
	),
	c.Example(
		`b.Buttons(
	b.ButtonA("Anchor"),
	b.Button("Button"),
	b.ButtonInputSubmit("Submit input"),
	b.ButtonInputReset("Reset input"),
)`,
		b.Buttons(
			b.ButtonA("Anchor"),
			b.Button("Button"),
			b.ButtonInputSubmit("Submit input"),
			b.ButtonInputReset("Reset input"),
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/button/#colors",
	c.Example(
		`b.Buttons(
	b.Button(b.White, "White"),
	b.Button(b.Light, "Light"),
	b.Button(b.Dark, "Dark"),
	b.Button(b.Black, "Black"),
	b.Button(b.Text, "Text"),
	b.Button(b.Ghost, "Ghost"),
)`,
		b.Buttons(
			b.Button(b.White, "White"),
			b.Button(b.Light, "Light"),
			b.Button(b.Dark, "Dark"),
			b.Button(b.Black, "Black"),
			b.Button(b.Text, "Text"),
			b.Button(b.Ghost, "Ghost"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Primary, "Primary"),
	b.Button(b.Link, "Link"),
),
b.Buttons(
	b.Button(b.Info, "Info"),
	b.Button(b.Success, "Success"),
	b.Button(b.Warning, "Warning"),
	b.Button(b.Danger, "Danger"),
)`,
		b.Buttons(
			b.Button(b.Primary, "Primary"),
			b.Button(b.Link, "Link"),
		),
		b.Buttons(
			b.Button(b.Info, "Info"),
			b.Button(b.Success, "Success"),
			b.Button(b.Warning, "Warning"),
			b.Button(b.Danger, "Danger"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.PrimaryLight, "Primary"),
	b.Button(b.LinkLight, "Link"),
),
b.Buttons(
	b.Button(b.InfoLight, "Info"),
	b.Button(b.SuccessLight, "Success"),
	b.Button(b.WarningLight, "Warning"),
	b.Button(b.DangerLight, "Danger"),
)`,
		b.Buttons(
			b.Button(b.PrimaryLight, "Primary"),
			b.Button(b.LinkLight, "Link"),
		),
		b.Buttons(
			b.Button(b.InfoLight, "Info"),
			b.Button(b.SuccessLight, "Success"),
			b.Button(b.WarningLight, "Warning"),
			b.Button(b.DangerLight, "Danger"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.PrimaryDark, "Primary"),
	b.Button(b.LinkDark, "Link"),
),
b.Buttons(
	b.Button(b.InfoDark, "Info"),
	b.Button(b.SuccessDark, "Success"),
	b.Button(b.WarningDark, "Warning"),
	b.Button(b.DangerDark, "Danger"),
)`,
		b.Buttons(
			b.Button(b.PrimaryDark, "Primary"),
			b.Button(b.LinkDark, "Link"),
		),
		b.Buttons(
			b.Button(b.InfoDark, "Info"),
			b.Button(b.SuccessDark, "Success"),
			b.Button(b.WarningDark, "Warning"),
			b.Button(b.DangerDark, "Danger"),
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/button/#sizes",
	c.Example(
		`b.Buttons(
	b.Button(b.Small, "Small"),
	b.Button("Default"),
	b.Button(b.Normal, "Normal"),
	b.Button(b.Medium, "Medium"),
	b.Button(b.Large, "Large"),
)`,
		b.Buttons(
			b.Button(b.Small, "Small"),
			b.Button("Default"),
			b.Button(b.Normal, "Normal"),
			b.Button(b.Medium, "Medium"),
			b.Button(b.Large, "Large"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Medium,
	b.Button("All"),
	b.Button("Medium"),
	b.Button("Size"),
)`,
		b.Buttons(
			b.Medium,
			b.Button("All"),
			b.Button("Medium"),
			b.Button("Size"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Small,
	b.Button("Small"),
	b.Button("Small"),
	b.Button("Small"),
	b.Button(b.Normal, "Normal"),
	b.Button("Small"),
)`,
		b.Buttons(
			b.Small,
			b.Button("Small"),
			b.Button("Small"),
			b.Button("Small"),
			b.Button(b.Normal, "Normal"),
			b.Button("Small"),
		),
	),
).Subsection(
	"Responsive sizes",
	"https://bulma.io/documentation/elements/button/#responsive-sizes",
	c.Example(
		`b.Button(b.Responsive, "Default")`,
		b.Button(b.Responsive, "Default"),
	),
	c.Example(
		`b.Button(b.Small, b.Responsive, "Small")`,
		b.Button(b.Small, b.Responsive, "Small"),
	),
	c.Example(
		`b.Button(b.Normal, b.Responsive, "Normal")`,
		b.Button(b.Normal, b.Responsive, "Normal"),
	),
	c.Example(
		`b.Button(b.Medium, b.Responsive, "Medium")`,
		b.Button(b.Medium, b.Responsive, "Medium"),
	),
	c.Example(
		`b.Button(b.Large, b.Responsive, "Large")`,
		b.Button(b.Large, b.Responsive, "Large"),
	),
).Subsection(
	"Displays",
	"https://bulma.io/documentation/elements/button/#displays",
	c.Example(
		`b.Buttons(
	b.Button(b.Small, b.FullWidth, "Small"),
	b.Button(b.FullWidth, "Normal"),
	b.Button(b.Medium, b.FullWidth, "Medium"),
	b.Button(b.Large, b.FullWidth, "Large"),
)`,
		b.Buttons(
			b.Button(b.Small, b.FullWidth, "Small"),
			b.Button(b.FullWidth, "Normal"),
			b.Button(b.Medium, b.FullWidth, "Medium"),
			b.Button(b.Large, b.FullWidth, "Large"),
		),
	),
).Subsection(
	"Styles",
	"https://bulma.io/documentation/elements/button/#styles",
	c.Example(
		`b.Buttons(
	b.Button(b.Link, b.Outlined, "Outlined"),
	b.Button(b.Primary, b.Outlined, "Outlined"),
	b.Button(b.Info, b.Outlined, "Outlined"),
	b.Button(b.Success, b.Outlined, "Outlined"),
	b.Button(b.Danger, b.Outlined, "Outlined"),
)`,
		b.Buttons(
			b.Button(b.Link, b.Outlined, "Outlined"),
			b.Button(b.Primary, b.Outlined, "Outlined"),
			b.Button(b.Info, b.Outlined, "Outlined"),
			b.Button(b.Success, b.Outlined, "Outlined"),
			b.Button(b.Danger, b.Outlined, "Outlined"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Link, b.Inverted, "Inverted"),
	b.Button(b.Primary, b.Inverted, "Inverted"),
	b.Button(b.Info, b.Inverted, "Inverted"),
	b.Button(b.Success, b.Inverted, "Inverted"),
	b.Button(b.Danger, b.Inverted, "Inverted"),
)`,
		b.Buttons(
			b.Button(b.Link, b.Inverted, "Inverted"),
			b.Button(b.Primary, b.Inverted, "Inverted"),
			b.Button(b.Info, b.Inverted, "Inverted"),
			b.Button(b.Success, b.Inverted, "Inverted"),
			b.Button(b.Danger, b.Inverted, "Inverted"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Rounded, "Rounded"),
	b.Button(b.Link, b.Rounded, "Rounded"),
	b.Button(b.Primary, b.Rounded, "Rounded"),
	b.Button(b.Info, b.Rounded, "Rounded"),
	b.Button(b.Success, b.Rounded, "Rounded"),
	b.Button(b.Danger, b.Rounded, "Rounded"),
)`,
		b.Buttons(
			b.Button(b.Rounded, "Rounded"),
			b.Button(b.Link, b.Rounded, "Rounded"),
			b.Button(b.Primary, b.Rounded, "Rounded"),
			b.Button(b.Info, b.Rounded, "Rounded"),
			b.Button(b.Success, b.Rounded, "Rounded"),
			b.Button(b.Danger, b.Rounded, "Rounded"),
		),
	),
).Subsection(
	"States",
	"https://bulma.io/documentation/elements/button/#states",
	c.Example(
		`b.Buttons(
	b.Button("Normal"),
	b.Button(b.Link, "Normal"),
	b.Button(b.Primary, "Normal"),
	b.Button(b.Info, "Normal"),
	b.Button(b.Success, "Normal"),
	b.Button(b.Danger, "Normal"),
)`,
		b.Buttons(
			b.Button("Normal"),
			b.Button(b.Link, "Normal"),
			b.Button(b.Primary, "Normal"),
			b.Button(b.Info, "Normal"),
			b.Button(b.Success, "Normal"),
			b.Button(b.Danger, "Normal"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Hovered, "Hover"),
	b.Button(b.Link, b.Hovered, "Hover"),
	b.Button(b.Primary, b.Hovered, "Hover"),
	b.Button(b.Info, b.Hovered, "Hover"),
	b.Button(b.Success, b.Hovered, "Hover"),
	b.Button(b.Danger, b.Hovered, "Hover"),
)`,
		b.Buttons(
			b.Button(b.Hovered, "Hover"),
			b.Button(b.Link, b.Hovered, "Hover"),
			b.Button(b.Primary, b.Hovered, "Hover"),
			b.Button(b.Info, b.Hovered, "Hover"),
			b.Button(b.Success, b.Hovered, "Hover"),
			b.Button(b.Danger, b.Hovered, "Hover"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Focused, "Focus"),
	b.Button(b.Link, b.Focused, "Focus"),
	b.Button(b.Primary, b.Focused, "Focus"),
	b.Button(b.Info, b.Focused, "Focus"),
	b.Button(b.Success, b.Focused, "Focus"),
	b.Button(b.Danger, b.Focused, "Focus"),
)`,
		b.Buttons(
			b.Button(b.Focused, "Focus"),
			b.Button(b.Link, b.Focused, "Focus"),
			b.Button(b.Primary, b.Focused, "Focus"),
			b.Button(b.Info, b.Focused, "Focus"),
			b.Button(b.Success, b.Focused, "Focus"),
			b.Button(b.Danger, b.Focused, "Focus"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Active, "Active"),
	b.Button(b.Link, b.Active, "Active"),
	b.Button(b.Primary, b.Active, "Active"),
	b.Button(b.Info, b.Active, "Active"),
	b.Button(b.Success, b.Active, "Active"),
	b.Button(b.Danger, b.Active, "Active"),
)`,

		b.Buttons(
			b.Button(b.Active, "Active"),
			b.Button(b.Link, b.Active, "Active"),
			b.Button(b.Primary, b.Active, "Active"),
			b.Button(b.Info, b.Active, "Active"),
			b.Button(b.Success, b.Active, "Active"),
			b.Button(b.Danger, b.Active, "Active"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button(b.Loading, "Loading"),
	b.Button(b.Link, b.Loading, "Loading"),
	b.Button(b.Primary, b.Loading, "Loading"),
	b.Button(b.Info, b.Loading, "Loading"),
	b.Button(b.Success, b.Loading, "Loading"),
	b.Button(b.Danger, b.Loading, "Loading"),
)`,
		b.Buttons(
			b.Button(b.Loading, "Loading"),
			b.Button(b.Link, b.Loading, "Loading"),
			b.Button(b.Primary, b.Loading, "Loading"),
			b.Button(b.Info, b.Loading, "Loading"),
			b.Button(b.Success, b.Loading, "Loading"),
			b.Button(b.Danger, b.Loading, "Loading"),
		),
	),
	c.Example(
		`b.Button(b.Static, "Static")`,
		b.Button(b.Static, "Static"),
	),
	c.Example(
		`b.Buttons(
	b.Button(
		e.TitleAttr("Disabled button"),
		e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Primary, e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Link, e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Info, e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Success, e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Warning, e.Disabled(), "Disabled",
	),
	b.Button(
		e.TitleAttr("Disabled button"),
		b.Danger, e.Disabled(), "Disabled",
	),
)`,
		b.Buttons(
			b.Button(
				e.TitleAttr("Disabled button"),
				e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Primary, e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Link, e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Info, e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Success, e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Warning, e.Disabled(), "Disabled",
			),
			b.Button(
				e.TitleAttr("Disabled button"),
				b.Danger, e.Disabled(), "Disabled",
			),
		),
	),
	b.Content(e.P("The ", e.Code("fa"), " package includes helper functions and values for using ", e.Em("Font Awesome"), " icons with ", e.Em("Bulma-Gomponents"), ".")),
	c.Example(
		`b.Buttons(
	html.P,
	b.Button(fa.Icon(fa.Solid, "bold")),
	b.Button(fa.Icon(fa.Solid, "italic")),
	b.Button(fa.Icon(fa.Solid, "underline")),
),
b.Buttons(
	html.P,
	b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Primary, fa.Icon(fa.Brand, "twitter"), "@jgthms"),
	b.Button(b.Success, fa.Icon(fa.Solid, "check"), "Save"),
	b.Button(b.Danger, b.Outlined, "Delete", fa.Icon(fa.Solid, "times", b.Small)),
),
b.Buttons(
	html.P,
	b.Button(b.Small, fa.Icon(fa.Brand, "github", b.Small), "GitHub"),
	b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Medium, fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Large, fa.Icon(fa.Brand, "github", b.Medium), "GitHub"),
)`,
		b.Buttons(
			html.P,
			b.Button(fa.Icon(fa.Solid, "bold")),
			b.Button(fa.Icon(fa.Solid, "italic")),
			b.Button(fa.Icon(fa.Solid, "underline")),
		),
		b.Buttons(
			html.P,
			b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Primary, fa.Icon(fa.Brand, "twitter"), "@jgthms"),
			b.Button(b.Success, fa.Icon(fa.Solid, "check"), "Save"),
			b.Button(b.Danger, b.Outlined, "Delete", fa.Icon(fa.Solid, "times", b.Small)),
		),
		b.Buttons(
			html.P,
			b.Button(b.Small, fa.Icon(fa.Brand, "github", b.Small), "GitHub"),
			b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Medium, fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Large, fa.Icon(fa.Brand, "github", b.Medium), "GitHub"),
		),
	),
	c.Example(
		`b.Buttons(
	html.P,
	b.Button(b.Small, fa.Icon(fa.Solid, "heading", b.Small)),
),
b.Buttons(
	html.P,
	b.Button(fa.Icon(fa.Solid, "heading", b.Small)),
	b.Button(fa.Icon(fa.Solid, "heading", fa.SizeLg)),
),
b.Buttons(
	html.P,
	b.Button(b.Medium, fa.Icon(fa.Solid, "heading", b.Small)),
	b.Button(b.Medium, fa.Icon(fa.Solid, "heading", fa.SizeLg)),
	b.Button(b.Medium, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size2x)),
),
b.Buttons(
	html.P,
	b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Small)),
	b.Button(b.Large, fa.Icon(fa.Solid, "heading", fa.SizeLg)),
	b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size2x)),
)`,
		b.Buttons(
			html.P,
			b.Button(b.Small, fa.Icon(fa.Solid, "heading", b.Small)),
		),
		b.Buttons(
			html.P,
			b.Button(fa.Icon(fa.Solid, "heading", b.Small)),
			b.Button(fa.Icon(fa.Solid, "heading", fa.SizeLg)),
		),
		b.Buttons(
			html.P,
			b.Button(b.Medium, fa.Icon(fa.Solid, "heading", b.Small)),
			b.Button(b.Medium, fa.Icon(fa.Solid, "heading", fa.SizeLg)),
			b.Button(b.Medium, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size(2))),
		),
		b.Buttons(
			html.P,
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Small)),
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", fa.SizeLg)),
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size(2))),
		),
	),
).Subsection(
	"Button group",
	"https://bulma.io/documentation/elements/button/#button-group",
	c.Example(
		`b.Field(
	b.Grouped,
	b.Control(
		html.P,
		b.Button(b.Link, "Save changes"),
	),
	b.Control(
		html.P,
		b.Button("Cancel"),
	),
	b.Control(
		html.P,
		b.Button(b.Danger, "Delete post"),
	),
)`,
		b.Field(
			b.Grouped,
			b.Control(
				html.P,
				b.Button(b.Link, "Save changes"),
			),
			b.Control(
				html.P,
				b.Button("Cancel"),
			),
			b.Control(
				html.P,
				b.Button(b.Danger, "Delete post"),
			),
		),
	),
).Subsection(
	"Button addons",
	"https://bulma.io/documentation/elements/button/#button-addons",
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-left", b.Small),
			"Left",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-center", b.Small),
			"Center",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-right", b.Small),
			"Right",
		),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-left", b.Small),
					"Left",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-center", b.Small),
					"Center",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-right", b.Small),
					"Right",
				),
			),
		),
	),
).Subsection(
	"Button group with addons",
	"https://bulma.io/documentation/elements/button/#button-group-with-addons",
	c.Example(
		`b.Field(
	b.Addons,
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "bold", b.Small),
			"Bold",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "italic", b.Small),
			"Italic",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "underline", b.Small),
			"Underline",
		),
	),
),
b.Field(
	b.Addons,
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-left", b.Small),
			"Left",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-center", b.Small),
			"Center",
		),
	),
	b.Control(
		html.P,
		b.Button(
			fa.Icon(fa.Solid, "align-right", b.Small),
			"Right",
		),
	),
)`,
		b.Field(
			b.Addons,
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "bold", b.Small),
					"Bold",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "italic", b.Small),
					"Italic",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "underline", b.Small),
					"Underline",
				),
			),
		),
		b.Field(
			b.Addons,
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-left", b.Small),
					"Left",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-center", b.Small),
					"Center",
				),
			),
			b.Control(
				html.P,
				b.Button(
					fa.Icon(fa.Solid, "align-right", b.Small),
					"Right",
				),
			),
		),
	),
).Subsection(
	"List of buttons",
	"https://bulma.io/documentation/elements/button/#list-of-buttons",
	c.Example(
		`b.Buttons(
	b.Button(b.Success, "Save changes"),
	b.Button(b.Info, "Save and continue"),
	b.Button(b.Danger, "Cancel"),
)`,
		b.Buttons(
			b.Button(b.Success, "Save changes"),
			b.Button(b.Info, "Save and continue"),
			b.Button(b.Danger, "Cancel"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Button("One"),
	b.Button("Two"),
	b.Button("Three"),
	b.Button("Four"),
	b.Button("Five"),
	b.Button("Six"),
	b.Button("Seven"),
	b.Button("Eight"),
	b.Button("Nine"),
	b.Button("Ten"),
	b.Button("Eleven"),
	b.Button("Twelve"),
	b.Button("Thirteen"),
	b.Button("Fourteen"),
	b.Button("Fifteen"),
	b.Button("Sixteen"),
	b.Button("Seventeen"),
	b.Button("Eighteen"),
	b.Button("Nineteen"),
	b.Button("Twenty"),
)`,
		b.Buttons(
			b.Button("One"),
			b.Button("Two"),
			b.Button("Three"),
			b.Button("Four"),
			b.Button("Five"),
			b.Button("Six"),
			b.Button("Seven"),
			b.Button("Eight"),
			b.Button("Nine"),
			b.Button("Ten"),
			b.Button("Eleven"),
			b.Button("Twelve"),
			b.Button("Thirteen"),
			b.Button("Fourteen"),
			b.Button("Fifteen"),
			b.Button("Sixteen"),
			b.Button("Seventeen"),
			b.Button("Eighteen"),
			b.Button("Nineteen"),
			b.Button("Twenty"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Addons,
	b.Button("Yes"),
	b.Button("Maybe"),
	b.Button("No"),
)`,
		b.Buttons(
			b.Addons,
			b.Button("Yes"),
			b.Button("Maybe"),
			b.Button("No"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Addons,
	b.Centered,
	b.Button("Yes"),
	b.Button("Maybe"),
	b.Button("No"),
),
b.Buttons(
	b.Addons,
	b.Right,
	b.Button("Yes"),
	b.Button("Maybe"),
	b.Button("No"),
)`,
		b.Buttons(
			b.Addons,
			b.Centered,
			b.Button("Yes"),
			b.Button("Maybe"),
			b.Button("No"),
		),
		b.Buttons(
			b.Addons,
			b.Right,
			b.Button("Yes"),
			b.Button("Maybe"),
			b.Button("No"),
		),
	),
	c.Example(
		`b.Buttons(
	b.Addons,
	b.Button(b.Success, b.Selected, "Yes"),
	b.Button("Maybe"),
	b.Button("No"),
),
b.Buttons(
	b.Addons,
	b.Button("Yes"),
	b.Button(b.Info, b.Selected, "Maybe"),
	b.Button("No"),
),
b.Buttons(
	b.Addons,
	b.Button("Yes"),
	b.Button("Maybe"),
	b.Button(b.Danger, b.Selected, "No"),
)`,
		b.Buttons(
			b.Addons,
			b.Button(b.Success, b.Selected, "Yes"),
			b.Button("Maybe"),
			b.Button("No"),
		),
		b.Buttons(
			b.Addons,
			b.Button("Yes"),
			b.Button(b.Info, b.Selected, "Maybe"),
			b.Button("No"),
		),
		b.Buttons(
			b.Addons,
			b.Button("Yes"),
			b.Button("Maybe"),
			b.Button(b.Danger, b.Selected, "No"),
		),
	),
)
