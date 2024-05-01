package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var progress = c.NewPage(
	"Progress bars", "Progress Bar", "/progress",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Progress"), " constructor creates a progress bar. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.White"),
			"Set progress bar color to white",

			e.Code("b.Black"),
			"Set progress bar color to black",

			e.Code("b.Light"),
			"Set progress bar color to light",

			e.Code("b.Dark"),
			"Set progress bar color to dark",

			e.Code("b.Primary"),
			"Set progress bar color to primary",

			e.Code("b.Link"),
			"Set progress bar color to link",

			e.Code("b.Info"),
			"Set progress bar color to info",

			e.Code("b.Success"),
			"Set progress bar color to success",

			e.Code("b.Warning"),
			"Set progress bar color to warning",

			e.Code("b.Danger"),
			"Set progress bar color to danger",

			e.Code("b.Small"),
			"Set progress bar size to small",

			e.Code("b.Normal"),
			"Set progress bar size to normal",

			e.Code("b.Medium"),
			"Set progress bar size to medium",

			e.Code("b.Large"),
			"Set progress bar size to large",
		),
		e.P(
			"The ", e.Code("b.ProgressIndeterminate"), " constructor creates an indeterminate progress bar. It accepts the same values as ", e.Code("b.Progress"), ".",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/progress/",

	c.Example(
		`b.Progress(15, 100, "15%")`,
		b.Progress(15, 100, "15%"),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/progress/#colors",
	c.Example(
		`b.Progress(15, 100, b.Primary, "15%")`,
		b.Progress(15, 100, b.Primary, "15%"),
	),
	c.Example(
		`b.Progress(30, 100, b.Link, "30%")`,
		b.Progress(30, 100, b.Link, "30%"),
	),
	c.Example(
		`b.Progress(45, 100, b.Info, "45%")`,
		b.Progress(45, 100, b.Info, "45%"),
	),
	c.Example(
		`b.Progress(60, 100, b.Success, "60%")`,
		b.Progress(60, 100, b.Success, "60%"),
	),
	c.Example(
		`b.Progress(75, 100, b.Warning, "75%")`,
		b.Progress(75, 100, b.Warning, "75%"),
	),
	c.Example(
		`b.Progress(90, 100, b.Danger, "90%")`,
		b.Progress(90, 100, b.Danger, "90%"),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/progress/#sizes",
	c.Example(
		`b.Progress(20, 100, b.Small, "20%")`,
		b.Progress(20, 100, b.Small, "20%"),
	),
	c.Example(
		`b.Progress(40, 100, b.Normal, "40%")`,
		b.Progress(40, 100, b.Normal, "40%"),
	),
	c.Example(
		`b.Progress(60, 100, b.Medium, "60%")`,
		b.Progress(60, 100, b.Medium, "60%"),
	),
	c.Example(
		`b.Progress(80, 100, b.Large, "80%")`,
		b.Progress(80, 100, b.Large, "80%"),
	),
).Subsection(
	"Indeterminate",
	"https://bulma.io/documentation/elements/progress/#indeterminate",
	c.Example(
		`b.ProgressIndeterminate(b.Small, b.Primary, "15%"),
b.ProgressIndeterminate(b.Danger, "30%"),
b.ProgressIndeterminate(b.Medium, b.Dark, "45%"),
b.ProgressIndeterminate(b.Large, b.Info, "60%")`,
		b.ProgressIndeterminate(b.Small, b.Primary, "15%"),
		b.ProgressIndeterminate(b.Danger, "30%"),
		b.ProgressIndeterminate(b.Medium, b.Dark, "45%"),
		b.ProgressIndeterminate(b.Large, b.Info, "60%"),
	),
)
