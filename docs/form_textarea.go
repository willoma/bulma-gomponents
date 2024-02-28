package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var formTextarea = c.NewPage(
	"Textarea", "Textarea", "/form/textarea",
	"https://bulma.io/documentation/form/textarea/",
	c.Example(
		`b.Textarea(html.Placeholder("e.g. Hello world"))`,
		b.Textarea(html.Placeholder("e.g. Hello world")),
	),
	c.Example(
		`b.Textarea(html.Placeholder("10 lines of textarea"), b.Rows(10))`,
		b.Textarea(html.Placeholder("10 lines of textarea"), b.Rows(10)),
	),
).Section(
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
).Section(
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
).Section(
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
	b.Textarea(html.Placeholder("Disabled textarea"), html.Disabled()),
)`,
		b.Control(
			b.Textarea(html.Placeholder("Disabled textarea"), html.Disabled()),
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
