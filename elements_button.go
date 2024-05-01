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
			"The ", e.Code("b.Button"), ", ", e.Code("b.ButtonA"), ", ", e.Code("b.ButtonAHref"), ", ", e.Code("b.ButtonSubmit"), ", ", e.Code("b.ButtonInputSubmit"), " and ", e.Code("b.ButtonInputReset"), " constructors create buttons. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.Responsive"),
			"Responsive size",

			e.Code("b.FullWidth"),
			"Take the whole width",

			e.Code("b.Outlined"),
			"Outline style",

			e.Code("b.Inverted"),
			"Inverted style",

			e.Code("b.Rounded"),
			"Rounded button",

			e.Code("b.Hovered"),
			"Apply the hovered style",

			e.Code("b.Focused"),
			"Apply the focused style",

			e.Code("b.Active"),
			"Apply the active style",

			e.Code("b.Loading"),
			"Replace the content with a loading spinner",

			e.Code("b.Static"),
			"Make the button non-interactive",

			e.Code("html.Disabled()"),
			"Disable the button",

			e.Code("b.Selected"),
			"In a list of attached buttons (Buttons with Addons), make sure this button is above the other buttons",

			e.Code("b.Small"),
			"Set button size to small",

			e.Code("b.Normal"),
			"Set button size to normal",

			e.Code("b.Medium"),
			"Set button size to medium",

			e.Code("b.Large"),
			"Set button size to large",

			e.Code("b.White"),
			"Set button color to white",

			e.Code("b.Light"),
			"Set button color to light",

			e.Code("b.Dark"),
			"Set button color to dark",

			e.Code("b.Black"),
			"Set button color to black",

			e.Code("b.Text"),
			"Set button style to underlined text",

			e.Code("b.Ghost"),
			"Set button style to link-looking blue text",

			e.Code("b.Primary"),
			"Set button color to primary",

			e.Code("b.Link"),
			"Set button color to link",

			e.Code("b.Info"),
			"Set button color to info",

			e.Code("b.Success"),
			"Set button color to success",

			e.Code("b.Warning"),
			"Set button color to warning",

			e.Code("b.Danger"),
			"Set button color to danger",

			e.Code("b.PrimaryLight"),
			"Set button color to primary light",

			e.Code("b.LinkLight"),
			"Set button color to link light",

			e.Code("b.InfoLight"),
			"Set button color to info light",

			e.Code("b.SuccessLight"),
			"Set button color to success light",

			e.Code("b.WarningLight"),
			"Set button color to warning light",

			e.Code("b.DangerLight"),
			"Set button color to danger light",

			e.Code("b.PrimaryDark"),
			"Set button color to primary dark",

			e.Code("b.LinkDark"),
			"Set button color to link dark",

			e.Code("b.InfoDark"),
			"Set button color to info dark",

			e.Code("b.SuccessDark"),
			"Set button color to success dark",

			e.Code("b.WarningDark"),
			"Set button color to warning dark",

			e.Code("b.DangerDark"),
			"Set button color to danger dark",
		),
		e.P("The ", e.Code("b.Buttons"), " constructor creates a button. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Addons"),
			"Attach the buttons together",

			e.Code("b.Centered"),
			"Center the buttons",

			e.Code("b.Right"),
			"Align the buttons to the right",

			e.Code("b.Small"),
			"Set buttons size to small",

			e.Code("b.Medium"),
			"Set buttons size to medium",

			e.Code("b.Large"),
			"Set buttons size to large",
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
		html.TitleAttr("Disabled button"),
		html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Primary, html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Link, html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Info, html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Success, html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Warning, html.Disabled(), "Disabled",
	),
	b.Button(
		html.TitleAttr("Disabled button"),
		b.Danger, html.Disabled(), "Disabled",
	),
)`,
		b.Buttons(
			b.Button(
				html.TitleAttr("Disabled button"),
				html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Primary, html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Link, html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Info, html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Success, html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Warning, html.Disabled(), "Disabled",
			),
			b.Button(
				html.TitleAttr("Disabled button"),
				b.Danger, html.Disabled(), "Disabled",
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
			b.Button(b.Medium, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size2x)),
		),
		b.Buttons(
			html.P,
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Small)),
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", fa.SizeLg)),
			b.Button(b.Large, fa.Icon(fa.Solid, "heading", b.Medium, fa.Size2x)),
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
