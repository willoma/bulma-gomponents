package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var tag = c.NewPage(
	"Tag", "Tags", "/tag",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Tag"), " constructor creates a tag. The ", e.Code("b.DeleteTag"), " constructor creates a delete button-looking tag.",
		),
		c.Modifiers(
			c.Row("b.White", "Set color to white"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
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
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Normal", "Set size to normal"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
			c.Row("b.Rounded", "Make a rounded tag"),
		),
		e.P(
			"The ", e.Code("b.Tags"), " constructor creates a list of tags.",
		),
		c.Modifiers(
			c.Row("b.Addons", "Attach the contained tags together"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/tag/",

	c.Example(
		`b.Tag("Tag label")`,
		b.Tag("Tag label"),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/tag/#colors",
	c.Example(
		`e.Div(b.MarginBottom(1), b.Tag(b.Black, "Black")),
e.Div(b.MarginBottom(1), b.Tag(b.Dark, "Dark")),
e.Div(b.MarginBottom(1), b.Tag(b.Light, "Light")),
e.Div(b.MarginBottom(1), b.Tag(b.White, "White")),
e.Div(b.MarginBottom(1), b.Tag(b.Primary, "Primary")),
e.Div(b.MarginBottom(1), b.Tag(b.Link, "Link")),
e.Div(b.MarginBottom(1), b.Tag(b.Info, "Info")),
e.Div(b.MarginBottom(1), b.Tag(b.Success, "Success")),
e.Div(b.MarginBottom(1), b.Tag(b.Warning, "Warning")),
e.Div(b.MarginBottom(1), b.Tag(b.Danger, "Danger")),`,
		e.Div(b.MarginBottom(1), b.Tag(b.Black, "Black")),
		e.Div(b.MarginBottom(1), b.Tag(b.Dark, "Dark")),
		e.Div(b.MarginBottom(1), b.Tag(b.Light, "Light")),
		e.Div(b.MarginBottom(1), b.Tag(b.White, "White")),
		e.Div(b.MarginBottom(1), b.Tag(b.Primary, "Primary")),
		e.Div(b.MarginBottom(1), b.Tag(b.Link, "Link")),
		e.Div(b.MarginBottom(1), b.Tag(b.Info, "Info")),
		e.Div(b.MarginBottom(1), b.Tag(b.Success, "Success")),
		e.Div(b.MarginBottom(1), b.Tag(b.Warning, "Warning")),
		e.Div(b.MarginBottom(1), b.Tag(b.Danger, "Danger")),
	),
	c.Example(
		`e.Div(b.MarginBottom(1), b.Tag(b.PrimaryLight, "Primary")),
e.Div(b.MarginBottom(1), b.Tag(b.LinkLight, "Link")),
e.Div(b.MarginBottom(1), b.Tag(b.InfoLight, "Info")),
e.Div(b.MarginBottom(1), b.Tag(b.SuccessLight, "Success")),
e.Div(b.MarginBottom(1), b.Tag(b.WarningLight, "Warning")),
e.Div(b.MarginBottom(1), b.Tag(b.DangerLight, "Danger"))`,
		e.Div(b.MarginBottom(1), b.Tag(b.PrimaryLight, "Primary")),
		e.Div(b.MarginBottom(1), b.Tag(b.LinkLight, "Link")),
		e.Div(b.MarginBottom(1), b.Tag(b.InfoLight, "Info")),
		e.Div(b.MarginBottom(1), b.Tag(b.SuccessLight, "Success")),
		e.Div(b.MarginBottom(1), b.Tag(b.WarningLight, "Warning")),
		e.Div(b.MarginBottom(1), b.Tag(b.DangerLight, "Danger")),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/tag/#sizes",
	c.Example(
		`e.Div(b.MarginBottom(1), b.Tag(b.Link, b.Normal, "Normal")),
e.Div(b.MarginBottom(1), b.Tag(b.Primary, b.Medium, "Medium")),
e.Div(b.MarginBottom(1), b.Tag(b.Info, b.Large, "Large"))`,
		e.Div(b.MarginBottom(1), b.Tag(b.Link, b.Normal, "Normal")),
		e.Div(b.MarginBottom(1), b.Tag(b.Primary, b.Medium, "Medium")),
		e.Div(b.MarginBottom(1), b.Tag(b.Info, b.Large, "Large")),
	),
	c.Example(
		`b.Tags(
	b.Medium,
	b.Tag("All"),
	b.Tag("Medium"),
	b.Tag("Size"),
)`,
		b.Tags(
			b.Medium,
			b.Tag("All"),
			b.Tag("Medium"),
			b.Tag("Size"),
		),
	),
	c.Example(
		`b.Tags(
	b.Large,
	b.Tag("All"),
	b.Tag("Large"),
	b.Tag("Size"),
)`,
		b.Tags(
			b.Large,
			b.Tag("All"),
			b.Tag("Large"),
			b.Tag("Size"),
		),
	),
	c.Example(
		`b.Tags(
	b.Medium,
	b.Tag("Medium"),
	b.Tag(b.Normal, "Normal"),
	b.Tag("Medium"),
	b.Tag(b.Large, "Large"),
	b.Tag("Medium"),
)`,
		b.Tags(
			b.Medium,
			b.Tag("Medium"),
			b.Tag(b.Normal, "Normal"),
			b.Tag("Medium"),
			b.Tag(b.Large, "Large"),
			b.Tag("Medium"),
		),
	),
).Subsection(
	"Modifiers",
	"https://bulma.io/documentation/elements/tag/#modifiers",
	c.Example(
		`b.Tag(b.Rounded, "Rounded")`,
		b.Tag(b.Rounded, "Rounded"),
	),
	c.Example(
		`b.DeleteTag()`,
		b.DeleteTag(),
	),
).Subsection(
	"Combinations",
	"https://bulma.io/documentation/elements/tag/#combinations",
	c.Example(
		`e.Div(b.MarginBottom(1), b.Tag(
	b.Success,
	"Bar",
	b.Delete(b.Small),
)),
e.Div(b.MarginBottom(1), b.Tag(
	b.Warning, b.Medium,
	"Hello",
	b.Delete(b.Small),
)),
e.Div(b.MarginBottom(1), b.Tag(
	b.Danger, b.Large,
	"World",
	b.Delete(),
))`,
		e.Div(b.MarginBottom(1), b.Tag(
			b.Success,
			"Bar",
			b.Delete(b.Small),
		)),
		e.Div(b.MarginBottom(1), b.Tag(
			b.Warning, b.Medium,
			"Hello",
			b.Delete(b.Small),
		)),
		e.Div(b.MarginBottom(1), b.Tag(
			b.Danger, b.Large,
			"World",
			b.Delete(),
		)),
	),
).Subsection(
	"List of tags",
	"https://bulma.io/documentation/elements/tag/#list-of-tags",
	c.Example(
		`b.Tags(
	b.Tag("One"),
	b.Tag("Two"),
	b.Tag("Three"),
)`,
		b.Tags(
			b.Tag("One"),
			b.Tag("Two"),
			b.Tag("Three"),
		),
	),
	c.Example(
		`b.Tags(
	b.Tag("One"),
	b.Tag("Two"),
	b.Tag("Three"),
	b.Tag("Four"),
	b.Tag("Five"),
	b.Tag("Six"),
	b.Tag("Seven"),
	b.Tag("Eight"),
	b.Tag("Nine"),
	b.Tag("Ten"),
	b.Tag("Eleven"),
	b.Tag("Twelve"),
	b.Tag("Thirteen"),
	b.Tag("Fourteen"),
	b.Tag("Fifteen"),
	b.Tag("Sixteen"),
	b.Tag("Seventeen"),
	b.Tag("Eighteen"),
	b.Tag("Nineteen"),
	b.Tag("Twenty"),
)`,
		b.Tags(
			b.Tag("One"),
			b.Tag("Two"),
			b.Tag("Three"),
			b.Tag("Four"),
			b.Tag("Five"),
			b.Tag("Six"),
			b.Tag("Seven"),
			b.Tag("Eight"),
			b.Tag("Nine"),
			b.Tag("Ten"),
			b.Tag("Eleven"),
			b.Tag("Twelve"),
			b.Tag("Thirteen"),
			b.Tag("Fourteen"),
			b.Tag("Fifteen"),
			b.Tag("Sixteen"),
			b.Tag("Seventeen"),
			b.Tag("Eighteen"),
			b.Tag("Nineteen"),
			b.Tag("Twenty"),
		),
	),
).Subsection(
	"Tag addons",
	"https://bulma.io/documentation/elements/tag/#tag-addons",
	c.Example(
		`b.Tags(
	b.Addons,
	b.Tag("Package"),
	b.Tag(b.Primary, "Bulma"),
)`,
		b.Tags(
			b.Addons,
			b.Tag("Package"),
			b.Tag(b.Primary, "Bulma"),
		),
	),
	c.Example(
		`b.Tags(
	b.Addons,
	b.Tag(b.Danger, "Alex Smith"),
	b.DeleteTag(),
)`,
		b.Tags(
			b.Addons,
			b.Tag(b.Danger, "Alex Smith"),
			b.DeleteTag(),
		),
	),
	c.Example(
		`b.Field(
	b.GroupedMultiline,
	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Dark, "npm"),
			b.Tag(b.Info, "0.9.4"),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Dark, "build"),
			b.Tag(b.Success, "passing"),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Dark, "chat"),
			b.Tag(b.Primary, "on gitter"),
		),
	),
)`,
		b.Field(
			b.GroupedMultiline,
			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Dark, "npm"),
					b.Tag(b.Info, "0.9.4"),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Dark, "build"),
					b.Tag(b.Success, "passing"),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Dark, "chat"),
					b.Tag(b.Primary, "on gitter"),
				),
			),
		),
	),
	c.Example(
		`b.Field(
	b.GroupedMultiline,
	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Technology"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "CSS"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Flexbox"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Web Design"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Open Source"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Community"),
			b.DeleteTag(),
		),
	),

	b.Control(
		b.Tags(
			b.Addons,
			b.Tag(b.Link, "Documentation"),
			b.DeleteTag(),
		),
	),
)`,
		b.Field(
			b.GroupedMultiline,
			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Technology"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "CSS"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Flexbox"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Web Design"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Open Source"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Community"),
					b.DeleteTag(),
				),
			),

			b.Control(
				b.Tags(
					b.Addons,
					b.Tag(b.Link, "Documentation"),
					b.DeleteTag(),
				),
			),
		),
	),
)
