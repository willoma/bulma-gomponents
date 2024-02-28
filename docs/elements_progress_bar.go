package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var progress = c.NewPage(
	"Progress bars", "Progress Bar", "/progress",
	"https://bulma.io/documentation/elements/progress/",

	c.Example(
		`b.Progress(15, 100, "15%")`,
		b.Progress(15, 100, "15%"),
	),
).Section(
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
).Section(
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
).Section(
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
