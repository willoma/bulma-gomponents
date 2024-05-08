package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var formSelect = c.NewPage(
	"Select", "Select", "/form/select",
	"",

	b.Content(
		e.P("The ", e.Code("b.Select"), " constructor creates a dropdown select."),
		c.Modifiers(
			c.Row("b.Size(int)", "For multiple selects, set size to the specified number of options"),
			c.Row("b.Rounded", "Make the select rounded"),
			c.Row("b.Hovered", "Apply the hovered style"),
			c.Row("b.Focused", "Apply the focused style"),
			c.Row("b.Multiple", "Make the select a multiple select"),
			c.Row("b.Disabled", "Disable the select"),
			c.Row("b.Loading", "Add a loading spinner to the right of the select"),
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
			c.Row("b.OnSelect(...any)", "Apply children to the ", e.Code("<select>"), " element"),
			c.Row("b.OnDiv(...any)", "Apply children to the ", e.Code(`<div class="select">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="select">`), " element"),
		),

		e.P("The ", e.Code("b.Option"), " constructor creates an option element to be used as a child of ", e.Code("b.Select"), "."),
		c.Modifiers(
			c.Row("b.Selected", "Mark the option as selected (for multiple selects)"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/select/",
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
).Subsection(
	"Multiple select",
	"https://bulma.io/documentation/form/select/#multiple-select",
	c.Example(
		`b.Select(
	b.Multiple,
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
		b.Select(
			b.Multiple,
			b.SelectSize(8),
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
).Subsection(
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
).Subsection(
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
).Subsection(
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
).Subsection(
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
).Subsection(
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
				b.Option("", b.Selected, "Country"),
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
				b.Option("", b.Selected, "Country"),
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
				b.Option("", b.Selected, "Country"),
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
				b.Option("", b.Selected, "Country"),
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
				b.Option("", b.Selected, "Country"),
				b.Option("", "Select dropdown"),
				b.Option("", "With options"),
			),
			fa.Icon(fa.Solid, "globe", b.Large, b.Left),
		),
	),
)
