package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var formTextarea = c.NewPage(
	"Textarea", "Textarea", "/form/textarea",
	"",

	b.Content(
		e.P("The ", e.Code("b.Textarea"), " constructor creates a text area."),
		c.Modifiers(
			c.Row("b.Rows(int)", "Set the text area height, in number of rows"),
			c.Row("b.Hovered", "Apply the hovered style"),
			c.Row("b.Focused", "Apply the focused style"),
			c.Row("b.Loading", "Add a loading spinner to the right of the text area"),
			c.Row("b.Disabled", "Disable the text area"),
			c.Row("e.ReadOnly()", "Read only text area"),
			c.Row("b.FixedSize", "Disable the text area resizing capability"),
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
			c.Row("b.OnTextarea(...any)", "Apply children to the ", e.Code("<textarea>"), " element - only useful with the ", e.Code("b.Disabled"), " modifier if you want to apply the ", e.Code("is-disabled"), " class to the text area instead of using the ", e.Code("disabled"), " attribute"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/textarea/",
	c.Example(
		`b.Textarea(e.Placeholder("e.g. Hello world"))`,
		b.Textarea(e.Placeholder("e.g. Hello world")),
	),
	c.Example(
		`b.Textarea(e.Placeholder("10 lines of textarea"), b.Rows(10))`,
		b.Textarea(e.Placeholder("10 lines of textarea"), b.Rows(10)),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/form/textarea/#colors",
	c.Example(
		`b.Textarea(b.Primary, e.Placeholder("Primary textarea"))`,
		b.Textarea(b.Primary, e.Placeholder("Primary textarea")),
	),
	c.Example(
		`b.Textarea(b.Link, e.Placeholder("Link textarea"))`,
		b.Textarea(b.Link, e.Placeholder("Link textarea")),
	),
	c.Example(
		`b.Textarea(b.Info, e.Placeholder("Info textarea"))`,
		b.Textarea(b.Info, e.Placeholder("Info textarea")),
	),
	c.Example(
		`b.Textarea(b.Success, e.Placeholder("Success textarea"))`,
		b.Textarea(b.Success, e.Placeholder("Success textarea")),
	),
	c.Example(
		`b.Textarea(b.Warning, e.Placeholder("Warning textarea"))`,
		b.Textarea(b.Warning, e.Placeholder("Warning textarea")),
	),
	c.Example(
		`b.Textarea(b.Danger, e.Placeholder("Danger textarea"))`,
		b.Textarea(b.Danger, e.Placeholder("Danger textarea")),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/form/textarea/#sizes",
	c.Example(
		`b.Field(
	b.Control(
		b.Textarea(b.Small, e.Placeholder("Small textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(e.Placeholder("Normal textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(b.Medium, e.Placeholder("Medium textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(b.Large, e.Placeholder("Large textarea")),
	),
)`,
		b.Field(
			b.Control(
				b.Textarea(b.Small, e.Placeholder("Small textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(e.Placeholder("Normal textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(b.Medium, e.Placeholder("Medium textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(b.Large, e.Placeholder("Large textarea")),
			),
		),
	),
).Subsection(
	"States",
	"https://bulma.io/documentation/form/textarea/#states",
	c.Example(
		`b.Control(
	b.Textarea(e.Placeholder("Normal textarea")),
)`,
		b.Control(
			b.Textarea(e.Placeholder("Normal textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.Hovered, e.Placeholder("Hovered textarea")),
)`,
		b.Control(
			b.Textarea(b.Hovered, e.Placeholder("Hovered textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.Focused, e.Placeholder("Focused textarea")),
)`,
		b.Control(
			b.Textarea(b.Focused, e.Placeholder("Focused textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Loading,
	b.Textarea(e.Placeholder("Loading textarea")),
)`,
		b.Control(
			b.Loading,
			b.Textarea(e.Placeholder("Loading textarea")),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.Small,
		b.Loading,
		b.Textarea(b.Small, e.Placeholder("Small loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Loading,
		b.Textarea(e.Placeholder("Normal loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Medium,
		b.Loading,
		b.Textarea(b.Medium, e.Placeholder("Medium loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Large,
		b.Loading,
		b.Textarea(b.Large, e.Placeholder("Large loading textarea")),
	),
)`,
		b.Field(
			b.Control(
				b.Small,
				b.Loading,
				b.Textarea(b.Small, e.Placeholder("Small loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Loading,
				b.Textarea(e.Placeholder("Normal loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Medium,
				b.Loading,
				b.Textarea(b.Medium, e.Placeholder("Medium loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Large,
				b.Loading,
				b.Textarea(b.Large, e.Placeholder("Large loading textarea")),
			),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(e.Placeholder("Disabled textarea"), b.Disabled),
)`,
		b.Control(
			b.Textarea(e.Placeholder("Disabled textarea"), b.Disabled),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(e.ReadOnly(), "This content is readonly"),
)`,
		b.Control(
			b.Textarea(e.ReadOnly(), "This content is readonly"),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.FixedSize, e.Placeholder("Fixed size textarea")),
)`,
		b.Control(
			b.Textarea(b.FixedSize, e.Placeholder("Fixed size textarea")),
		),
	),
)
