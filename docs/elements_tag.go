package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var tag = c.NewPage(
	"Tag", "Tags", "/tag",
	"https://bulma.io/documentation/elements/tag/",

	c.Example(
		`b.Tag("Tag label")`,
		b.Tag("Tag label"),
	),
).Section(
	"Colors",
	"https://bulma.io/documentation/elements/tag/#colors",
	c.Example(
		`b.Tag(b.Black, "Black"),
b.Tag(b.Dark, "Dark"),
b.Tag(b.Light, "Light"),
b.Tag(b.White, "White"),
b.Tag(b.Primary, "Primary"),
b.Tag(b.Link, "Link"),
b.Tag(b.Info, "Info"),
b.Tag(b.Success, "Success"),
b.Tag(b.Warning, "Warning"),
b.Tag(b.Danger, "Danger")`,
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Black, "Black")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Dark, "Dark")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Light, "Light")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.White, "White")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Primary, "Primary")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Link, "Link")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Info, "Info")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Success, "Success")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Warning, "Warning")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Danger, "Danger")),
	),
	c.Example(
		`b.Tag(b.PrimaryLight, "Primary"),
b.Tag(b.LinkLight, "Link"),
b.Tag(b.InfoLight, "Info"),
b.Tag(b.SuccessLight, "Success"),
b.Tag(b.WarningLight, "Warning"),
b.Tag(b.DangerLight, "Danger")`,
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.PrimaryLight, "Primary")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.LinkLight, "Link")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.InfoLight, "Info")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.SuccessLight, "Success")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.WarningLight, "Warning")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.DangerLight, "Danger")),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/elements/tag/#sizes",
	c.Example(
		`b.Tag(b.Link, b.Normal, "Normal"),
b.Tag(b.Primary, b.Medium, "Normal"),
b.Tag(b.Info, b.Large, "Normal")`,
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Link, b.Normal, "Normal")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Primary, b.Medium, "Medium")),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(b.Info, b.Large, "Large")),
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
).Section(
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
).Section(
	"Combinations",
	"https://bulma.io/documentation/elements/tag/#combinations",
	c.Example(
		`b.Tag(
	b.Success,
	"Bar",
	b.Delete(b.Small),
),
b.Tag(
	b.Warning, b.Medium,
	"Hello",
	b.Delete(b.Small),
),
b.Tag(
	b.Danger, b.Large,
	"World",
	b.Delete(),
)`,
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(
			b.Success,
			"Bar",
			b.Delete(b.Small),
		)),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(
			b.Warning, b.Medium,
			"Hello",
			b.Delete(b.Small),
		)),
		el.Div(b.MarginBottom(b.Spacing1), b.Tag(
			b.Danger, b.Large,
			"World",
			b.Delete(),
		)),
	),
).Section(
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
).Section(
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
