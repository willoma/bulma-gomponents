package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var button = c.NewPage(
	"Button", "Button", "/button",
	"https://bulma.io/documentation/elements/button/",

	c.Example(
		`b.Button("Button")`,
		b.Button("Button"),
	),
	c.Example(
		`b.ButtonA("Anchor"),
b.Button("Button"),
b.ButtonInputSubmit("Submit input"),
b.ButtonInputReset("Reset input")`,
		b.Buttons(
			b.ButtonA("Anchor"),
			b.Button("Button"),
			b.ButtonInputSubmit("Submit input"),
			b.ButtonInputReset("Reset input"),
		),
	),
).Section(
	"Colors",
	"https://bulma.io/documentation/elements/button/#colors",
	c.Example(
		`b.Button(b.White, "White"),
b.Button(b.Light, "Light"),
b.Button(b.Dark, "Dark"),
b.Button(b.Black, "Black"),
b.Button(b.Text, "Text"),
b.Button(b.Ghost, "Ghost")`,
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
).Section(
	"Sizes",
	"https://bulma.io/documentation/elements/button/#sizes",
	c.Example(
		`b.Button(b.Small, "Small"),
b.Button("Default"),
b.Button(b.Normal, "Normal"),
b.Button(b.Medium, "Medium"),
b.Button(b.Large, "Large")`,
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
).Section(
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
).Section(
	"Displays",
	"https://bulma.io/documentation/elements/button/#displays",
	c.Example(
		`b.Button(b.Small, b.FullWidth, "Small"),
b.Button(b.FullWidth, "Normal"),
b.Button(b.Medium, b.FullWidth, "Medium"),
b.Button(b.Large, b.FullWidth, "Large")`,
		b.Buttons(
			b.Button(b.Small, b.FullWidth, "Small"),
			b.Button(b.FullWidth, "Normal"),
			b.Button(b.Medium, b.FullWidth, "Medium"),
			b.Button(b.Large, b.FullWidth, "Large"),
		),
	),
).Section(
	"Styles",
	"https://bulma.io/documentation/elements/button/#styles",
	c.Example(
		`b.Button(b.Outlined, "Outlined"),
b.Button(b.Primary, b.Outlined, "Outlined"),
b.Button(b.Link, b.Outlined, "Outlined"),
b.Button(b.Info, b.Outlined, "Outlined"),
b.Button(b.Success, b.Outlined, "Outlined"),
b.Button(b.Danger, b.Outlined, "Outlined")`,
		b.Buttons(
			b.Button(b.Outlined, "Outlined"),
			b.Button(b.Primary, b.Outlined, "Outlined"),
			b.Button(b.Link, b.Outlined, "Outlined"),
			b.Button(b.Info, b.Outlined, "Outlined"),
			b.Button(b.Success, b.Outlined, "Outlined"),
			b.Button(b.Danger, b.Outlined, "Outlined"),
		),
	),
	c.Example(
		`b.Button(b.Primary, b.Inverted, "Inverted"),
b.Button(b.Link, b.Inverted, "Inverted"),
b.Button(b.Info, b.Inverted, "Inverted"),
b.Button(b.Success, b.Inverted, "Inverted"),
b.Button(b.Danger, b.Inverted, "Inverted")`,
		b.Buttons(
			b.Button(b.Primary, b.Inverted, "Inverted"),
			b.Button(b.Link, b.Inverted, "Inverted"),
			b.Button(b.Info, b.Inverted, "Inverted"),
			b.Button(b.Success, b.Inverted, "Inverted"),
			b.Button(b.Danger, b.Inverted, "Inverted"),
		),
	),
	c.Example(
		`b.Button(b.Primary, b.Inverted, b.Outlined, "Invert outlined"),
b.Button(b.Link, b.Inverted, b.Outlined, "Invert outlined"),
b.Button(b.Info, b.Inverted, b.Outlined, "Invert outlined"),
b.Button(b.Success, b.Inverted, b.Outlined, "Invert outlined"),
b.Button(b.Danger, b.Inverted, b.Outlined, "Invert outlined")`,
		b.Buttons(
			b.Button(b.Primary, b.Inverted, b.Outlined, "Invert outlined"),
			b.Button(b.Link, b.Inverted, b.Outlined, "Invert outlined"),
			b.Button(b.Info, b.Inverted, b.Outlined, "Invert outlined"),
			b.Button(b.Success, b.Inverted, b.Outlined, "Invert outlined"),
			b.Button(b.Danger, b.Inverted, b.Outlined, "Invert outlined"),
		),
	),
	c.Example(
		`b.Button(b.Rounded, "Rounded"),
b.Button(b.Primary, b.Rounded, "Rounded"),
b.Button(b.Link, b.Rounded, "Rounded"),
b.Button(b.Info, b.Rounded, "Rounded"),
b.Button(b.Success, b.Rounded, "Rounded"),
b.Button(b.Danger, b.Rounded, "Rounded")`,
		b.Buttons(
			b.Button(b.Rounded, "Rounded"),
			b.Button(b.Primary, b.Rounded, "Rounded"),
			b.Button(b.Link, b.Rounded, "Rounded"),
			b.Button(b.Info, b.Rounded, "Rounded"),
			b.Button(b.Success, b.Rounded, "Rounded"),
			b.Button(b.Danger, b.Rounded, "Rounded"),
		),
	),
).Section(
	"States",
	"https://bulma.io/documentation/elements/button/#states",
	c.Example(
		`b.Button("Normal"),
b.Button(b.Primary, "Normal"),
b.Button(b.Link, "Normal"),
b.Button(b.Info, "Normal"),
b.Button(b.Success, "Normal"),
b.Button(b.Danger, "Normal")`,
		b.Buttons(
			b.Button("Normal"),
			b.Button(b.Primary, "Normal"),
			b.Button(b.Link, "Normal"),
			b.Button(b.Info, "Normal"),
			b.Button(b.Success, "Normal"),
			b.Button(b.Danger, "Normal"),
		),
	),
	c.Example(
		`b.Button(b.Hovered, "Hover"),
b.Button(b.Primary, b.Hovered, "Hover"),
b.Button(b.Link, b.Hovered, "Hover"),
b.Button(b.Info, b.Hovered, "Hover"),
b.Button(b.Success, b.Hovered, "Hover"),
b.Button(b.Danger, b.Hovered, "Hover")`,
		b.Buttons(
			b.Button(b.Hovered, "Hover"),
			b.Button(b.Primary, b.Hovered, "Hover"),
			b.Button(b.Link, b.Hovered, "Hover"),
			b.Button(b.Info, b.Hovered, "Hover"),
			b.Button(b.Success, b.Hovered, "Hover"),
			b.Button(b.Danger, b.Hovered, "Hover"),
		),
	),
	c.Example(
		`b.Button(b.Focused, "Focus"),
b.Button(b.Primary, b.Focused, "Focus"),
b.Button(b.Link, b.Focused, "Focus"),
b.Button(b.Info, b.Focused, "Focus"),
b.Button(b.Success, b.Focused, "Focus"),
b.Button(b.Danger, b.Focused, "Focus")`,
		b.Buttons(
			b.Button(b.Focused, "Focus"),
			b.Button(b.Primary, b.Focused, "Focus"),
			b.Button(b.Link, b.Focused, "Focus"),
			b.Button(b.Info, b.Focused, "Focus"),
			b.Button(b.Success, b.Focused, "Focus"),
			b.Button(b.Danger, b.Focused, "Focus"),
		),
	),
	c.Example(
		`b.Button(b.Active, "Active"),
b.Button(b.Primary, b.Active, "Active"),
b.Button(b.Link, b.Active, "Active"),
b.Button(b.Info, b.Active, "Active"),
b.Button(b.Success, b.Active, "Active"),
b.Button(b.Danger, b.Active, "Active")`,

		b.Buttons(
			b.Button(b.Active, "Active"),
			b.Button(b.Primary, b.Active, "Active"),
			b.Button(b.Link, b.Active, "Active"),
			b.Button(b.Info, b.Active, "Active"),
			b.Button(b.Success, b.Active, "Active"),
			b.Button(b.Danger, b.Active, "Active"),
		),
	),
	c.Example(
		`b.Button(b.Loading, "Loading"),
b.Button(b.Primary, b.Loading, "Loading"),
b.Button(b.Link, b.Loading, "Loading"),
b.Button(b.Info, b.Loading, "Loading"),
b.Button(b.Success, b.Loading, "Loading"),
b.Button(b.Danger, b.Loading, "Loading")`,
		b.Buttons(
			b.Button(b.Loading, "Loading"),
			b.Button(b.Primary, b.Loading, "Loading"),
			b.Button(b.Link, b.Loading, "Loading"),
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
		`b.Button(
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
	b.Content(el.P("The ", el.Code("bulma/fa"), " package includes helper functions and values for using ", el.Em("Font Awesome"), " icons with ", el.Em("Bulma-Gomponents"), ".")),
	c.Example(
		`b.Buttons(
	b.Button(fa.Icon(fa.Solid, "bold")),
	b.Button(fa.Icon(fa.Solid, "italic")),
	b.Button(fa.Icon(fa.Solid, "underline")),
),
b.Buttons(
	b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Primary, fa.Icon(fa.Brand, "twitter"), "@jgthms"),
	b.Button(b.Success, fa.Icon(fa.Solid, "check"), "Save"),
	b.Button(b.Danger, b.Outlined, "Delete", fa.Icon(fa.Solid, "times", b.Small)),
),
b.Buttons(
	b.Button(b.Small, fa.Icon(fa.Brand, "github", b.Small), "GitHub"),
	b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Medium, fa.Icon(fa.Brand, "github"), "GitHub"),
	b.Button(b.Large, fa.Icon(fa.Brand, "github", b.Medium), "GitHub"),
)`,
		b.Buttons(
			b.Button(fa.Icon(fa.Solid, "bold")),
			b.Button(fa.Icon(fa.Solid, "italic")),
			b.Button(fa.Icon(fa.Solid, "underline")),
		),
		b.Buttons(
			b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Primary, fa.Icon(fa.Brand, "twitter"), "@jgthms"),
			b.Button(b.Success, fa.Icon(fa.Solid, "check"), "Save"),
			b.Button(b.Danger, b.Outlined, "Delete", fa.Icon(fa.Solid, "times", b.Small)),
		),
		b.Buttons(
			b.Button(b.Small, fa.Icon(fa.Brand, "github", b.Small), "GitHub"),
			b.Button(fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Medium, fa.Icon(fa.Brand, "github"), "GitHub"),
			b.Button(b.Large, fa.Icon(fa.Brand, "github", b.Medium), "GitHub"),
		),
	),
).Section(
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
).Section(
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
).Section(
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
).Section(
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
