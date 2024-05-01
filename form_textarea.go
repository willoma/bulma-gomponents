package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var formTextarea = c.NewPage(
	"Textarea", "Textarea", "/form/textarea",
	"",

	b.Content(
		e.P("The ", e.Code("b.Textarea"), " constructor creates a text area. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Rows(int)"),
			"Set the text area height, in number of rows",

			e.Code("b.Hovered"),
			"Apply the hovered style",

			e.Code("b.Focused"),
			"Apply the focused style",

			e.Code("b.Loading"),
			"Add a loading spinner to the right of the text area",

			e.Code("b.Disabled"),
			"Disable the text area",

			e.Code("html.ReadOnly()"),
			"Read only text area",

			e.Code("b.FixedSize"),
			"Disable the text area resizing capability",

			e.Code("b.Primary"),
			"Set text area color to primary",

			e.Code("b.Link"),
			"Set text area color to link",

			e.Code("b.Info"),
			"Set text area color to info",

			e.Code("b.Success"),
			"Set text area color to success",

			e.Code("b.Warning"),
			"Set text area color to warning",

			e.Code("b.Danger"),
			"Set text area color to danger",

			e.Code("b.Small"),
			"Set text area size to small",

			e.Code("b.Normal"),
			"Set text area size to normal",

			e.Code("b.Medium"),
			"Set text area size to medium",

			e.Code("b.Large"),
			"Set text area size to large",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/textarea/",
	c.Example(
		`b.Textarea(html.Placeholder("e.g. Hello world"))`,
		b.Textarea(html.Placeholder("e.g. Hello world")),
	),
	c.Example(
		`b.Textarea(html.Placeholder("10 lines of textarea"), b.Rows(10))`,
		b.Textarea(html.Placeholder("10 lines of textarea"), b.Rows(10)),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/form/textarea/#colors",
	c.Example(
		`b.Textarea(b.Primary, html.Placeholder("Primary textarea"))`,
		b.Textarea(b.Primary, html.Placeholder("Primary textarea")),
	),
	c.Example(
		`b.Textarea(b.Link, html.Placeholder("Link textarea"))`,
		b.Textarea(b.Link, html.Placeholder("Link textarea")),
	),
	c.Example(
		`b.Textarea(b.Info, html.Placeholder("Info textarea"))`,
		b.Textarea(b.Info, html.Placeholder("Info textarea")),
	),
	c.Example(
		`b.Textarea(b.Success, html.Placeholder("Success textarea"))`,
		b.Textarea(b.Success, html.Placeholder("Success textarea")),
	),
	c.Example(
		`b.Textarea(b.Warning, html.Placeholder("Warning textarea"))`,
		b.Textarea(b.Warning, html.Placeholder("Warning textarea")),
	),
	c.Example(
		`b.Textarea(b.Danger, html.Placeholder("Danger textarea"))`,
		b.Textarea(b.Danger, html.Placeholder("Danger textarea")),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/form/textarea/#sizes",
	c.Example(
		`b.Field(
	b.Control(
		b.Textarea(b.Small, html.Placeholder("Small textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(html.Placeholder("Normal textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(b.Medium, html.Placeholder("Medium textarea")),
	),
),
b.Field(
	b.Control(
		b.Textarea(b.Large, html.Placeholder("Large textarea")),
	),
)`,
		b.Field(
			b.Control(
				b.Textarea(b.Small, html.Placeholder("Small textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(html.Placeholder("Normal textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(b.Medium, html.Placeholder("Medium textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Textarea(b.Large, html.Placeholder("Large textarea")),
			),
		),
	),
).Subsection(
	"States",
	"https://bulma.io/documentation/form/textarea/#states",
	c.Example(
		`b.Control(
	b.Textarea(html.Placeholder("Normal textarea")),
)`,
		b.Control(
			b.Textarea(html.Placeholder("Normal textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.Hovered, html.Placeholder("Hovered textarea")),
)`,
		b.Control(
			b.Textarea(b.Hovered, html.Placeholder("Hovered textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.Focused, html.Placeholder("Focused textarea")),
)`,
		b.Control(
			b.Textarea(b.Focused, html.Placeholder("Focused textarea")),
		),
	),
	c.Example(
		`b.Control(
	b.Loading,
	b.Textarea(html.Placeholder("Loading textarea")),
)`,
		b.Control(
			b.Loading,
			b.Textarea(html.Placeholder("Loading textarea")),
		),
	),
	c.Example(
		`b.Field(
	b.Control(
		b.Small,
		b.Loading,
		b.Textarea(b.Small, html.Placeholder("Small loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Loading,
		b.Textarea(html.Placeholder("Normal loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Medium,
		b.Loading,
		b.Textarea(b.Medium, html.Placeholder("Medium loading textarea")),
	),
),
b.Field(
	b.Control(
		b.Large,
		b.Loading,
		b.Textarea(b.Large, html.Placeholder("Large loading textarea")),
	),
)`,
		b.Field(
			b.Control(
				b.Small,
				b.Loading,
				b.Textarea(b.Small, html.Placeholder("Small loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Loading,
				b.Textarea(html.Placeholder("Normal loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Medium,
				b.Loading,
				b.Textarea(b.Medium, html.Placeholder("Medium loading textarea")),
			),
		),
		b.Field(
			b.Control(
				b.Large,
				b.Loading,
				b.Textarea(b.Large, html.Placeholder("Large loading textarea")),
			),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(html.Placeholder("Disabled textarea"), b.Disabled),
)`,
		b.Control(
			b.Textarea(html.Placeholder("Disabled textarea"), b.Disabled),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(html.ReadOnly(), "This content is readonly"),
)`,
		b.Control(
			b.Textarea(html.ReadOnly(), "This content is readonly"),
		),
	),
	c.Example(
		`b.Control(
	b.Textarea(b.FixedSize, html.Placeholder("Fixed size textarea")),
)`,
		b.Control(
			b.Textarea(b.FixedSize, html.Placeholder("Fixed size textarea")),
		),
	),
)
