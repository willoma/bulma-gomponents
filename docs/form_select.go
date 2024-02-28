package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var formSelect = c.NewPage(
	"Select", "Select", "/form/select",
	"https://bulma.io/documentation/form/select/",
	c.Example(
		`b.Select(
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
).Section(
	"Multiple select",
	"https://bulma.io/documentation/form/select/#multiple-select",
	c.Example(
		`b.SelectMultiple(
	b.Size(8),
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
)`,
		b.SelectMultiple(
			b.Size(8),
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
).Section(
	"Colors",
	"https://bulma.io/documentation/form/select/#colors",
	c.Example(
		`b.Select(
	b.Primary,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Primary,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Link,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Link,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Info,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Info,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Success,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Success,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Warning,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Warning,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Danger,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Danger,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
).Section(
	"Styles",
	"https://bulma.io/documentation/form/select/#styles",
	c.Example(
		`b.Select(
	b.Rounded,
	b.Option("", "Rounded dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Rounded,
			b.Option("", "Rounded dropdown"),
			b.Option("", "With options"),
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/form/select/#sizes",
	c.Example(
		`b.Select(
	b.Small,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Small,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Normal,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Normal,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Medium,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Medium,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Large,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Large,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
).Section(
	"States",
	"https://bulma.io/documentation/form/select/#states",
	c.Example(
		`b.Select(
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Hovered,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Hovered,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Focused,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Focused,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
	c.Example(
		`b.Select(
	b.Loading,
	b.Option("", "Select dropdown"),
	b.Option("", "With options"),
)`,
		b.Select(
			b.Loading,
			b.Option("", "Select dropdown"),
			b.Option("", "With options"),
		),
	),
).Section(
	"With icons",
	"https://bulma.io/documentation/form/select/#with-icons",
	c.Example(
		`b.Control(
	b.IconsLeft,
	b.Select(
		b.OptionSelected("", "Country"),
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
	fa.Icon(fa.Solid, "globe", b.Small, b.Left),
)`,
		b.Control(
			b.IconsLeft,
			b.Select(
				b.OptionSelected("", "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Small, b.Left),
		),
	),
	c.Example(
		`b.Control(
	b.IconsLeft,
	b.Select(
		b.Small,
		b.OptionSelected("", "Country"),
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
	fa.Icon(fa.Solid, "globe", b.Small, b.Left),
)`,
		b.Control(
			b.IconsLeft,
			b.Select(
				b.Small,
				b.OptionSelected("", "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Small, b.Left),
		),
	),
	c.Example(
		`b.Control(
	b.IconsLeft,
	b.Select(
		b.OptionSelected("", "Country"),
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
	fa.Icon(fa.Solid, "globe", b.Left),
)`,
		b.Control(
			b.IconsLeft,
			b.Select(
				b.OptionSelected("", "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Left),
		),
	),
	c.Example(
		`b.Control(
	b.IconsLeft,
	b.Select(
		b.Medium,
		b.OptionSelected("", "Country"),
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
	fa.Icon(fa.Solid, "globe", b.Medium, b.Left),
)`,
		b.Control(
			b.IconsLeft,
			b.Select(
				b.Medium,
				b.OptionSelected("", "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Medium, b.Left),
		),
	),
	c.Example(
		`b.Control(
	b.IconsLeft,
	b.Select(
		b.Large,
		b.OptionSelected("", "Country"),
		b.Option("", "Select dropdown"),
		b.Option("", "With options"),
	),
	fa.Icon(fa.Solid, "globe", b.Large, b.Left),
)`,
		b.Control(
			b.IconsLeft,
			b.Select(
				b.Large,
				b.OptionSelected("", "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Large, b.Left),
		),
	),
)
