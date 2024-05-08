package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var progress = c.NewPage(
	"Progress bars", "Progress Bar", "/progress",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Progress"), " constructor creates a progress bar.",
		),
		c.Modifiers(
			c.Row("b.White", "Set color to white"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
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
		e.P(
			"The ", e.Code("b.ProgressIndeterminate"), " constructor creates an indeterminate progress bar. It accepts the same modifiers as ", e.Code("b.Progress"), ".",
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
