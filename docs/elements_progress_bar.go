package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var progress = c.NewPage(
	"Progress bars", "Progress Bar", "/progress",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Progress"), " constructor creates a progress bar. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.White"),
			"Set progress bar color to white",

			el.Code("b.Black"),
			"Set progress bar color to black",

			el.Code("b.Light"),
			"Set progress bar color to light",

			el.Code("b.Dark"),
			"Set progress bar color to dark",

			el.Code("b.Primary"),
			"Set progress bar color to primary",

			el.Code("b.Link"),
			"Set progress bar color to link",

			el.Code("b.Info"),
			"Set progress bar color to info",

			el.Code("b.Success"),
			"Set progress bar color to success",

			el.Code("b.Warning"),
			"Set progress bar color to warning",

			el.Code("b.Danger"),
			"Set progress bar color to danger",

			el.Code("b.Small"),
			"Set progress bar size to small",

			el.Code("b.Normal"),
			"Set progress bar size to normal",

			el.Code("b.Medium"),
			"Set progress bar size to medium",

			el.Code("b.Large"),
			"Set progress bar size to large",
		),
		el.P(
			"The ", el.Code("b.ProgressIndeterminate"), " constructor creates an indeterminate progress bar. It accepts the same values as ", el.Code("b.Progress"), ".",
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
