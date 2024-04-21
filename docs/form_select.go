package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var formSelect = c.NewPage(
	"Select", "Select", "/form/select",
	"",

	b.Content(
		el.P("The ", el.Code("b.Select"), " constructor creates a dropdown select. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnDiv(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="select">`), " element"},

			el.Code("b.OnSelect(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<select>"), " element"},

			el.Code("b.Size(int)"),
			"For multiple selects, set the select size to the specified number of options",

			el.Code("b.Rounded"),
			"Make the select rounded",

			el.Code("b.Hovered"),
			"Apply the hovered style",

			el.Code("b.Focused"),
			"Apply the focused style",

			el.Code("b.Multiple"),
			"Make the select a multiple select",

			el.Code("b.Disabled"),
			"Disable the select",

			el.Code("b.Loading"),
			"Add a loading spinner to the right of the select",

			el.Code("b.Primary"),
			"Set select color to primary",

			el.Code("b.Link"),
			"Set select color to link",

			el.Code("b.Info"),
			"Set select color to info",

			el.Code("b.Success"),
			"Set select color to success",

			el.Code("b.Warning"),
			"Set select color to warning",

			el.Code("b.Danger"),
			"Set select color to danger",

			el.Code("b.Small"),
			"Set select size to small",

			el.Code("b.Normal"),
			"Set select size to normal",

			el.Code("b.Medium"),
			"Set select size to medium",

			el.Code("b.Large"),
			"Set select size to large",
		),

		el.P("The ", el.Code("b.Option"), " constructor creates an option element to be used as a child of ", el.Code("b.Select"), ". The ", el.Code("b.OptionSelected"), " constructor creates a selected option element."),
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
