package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var formSelect = c.NewPage(
	"Select", "Select", "/form/select",
	"",

	b.Content(
		e.P("The ", e.Code("b.Select"), " constructor creates a dropdown select. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnDiv(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="select">`), " e.Element"},

			e.Code("b.OnSelect(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<select>"), " e.Element"},

			e.Code("b.Size(int)"),
			"For multiple selects, set the select size to the specified number of options",

			e.Code("b.Rounded"),
			"Make the select rounded",

			e.Code("b.Hovered"),
			"Apply the hovered style",

			e.Code("b.Focused"),
			"Apply the focused style",

			e.Code("b.Multiple"),
			"Make the select a multiple select",

			e.Code("b.Disabled"),
			"Disable the select",

			e.Code("b.Loading"),
			"Add a loading spinner to the right of the select",

			e.Code("b.Primary"),
			"Set select color to primary",

			e.Code("b.Link"),
			"Set select color to link",

			e.Code("b.Info"),
			"Set select color to info",

			e.Code("b.Success"),
			"Set select color to success",

			e.Code("b.Warning"),
			"Set select color to warning",

			e.Code("b.Danger"),
			"Set select color to danger",

			e.Code("b.Small"),
			"Set select size to small",

			e.Code("b.Normal"),
			"Set select size to normal",

			e.Code("b.Medium"),
			"Set select size to medium",

			e.Code("b.Large"),
			"Set select size to large",
		),

		e.P("The ", e.Code("b.Option"), " constructor creates an option e.Element to be used as a child of ", e.Code("b.Select"), ". The ", e.Code("b.OptionSelected"), " constructor creates a selected option e.Element."),
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
