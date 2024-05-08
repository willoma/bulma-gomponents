package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var formFile = c.NewPage(
	"File", "File upload", "/form/file",
	"",

	b.Content(
		e.P("The ", e.Code("b.File"), " constructor creates a file input."),
		c.Modifiers(
			c.Row("b.Right", "Move the call-to-action to the right side, align the file input to the right"),
			c.Row("b.FullWidth", "Expand the name to fill up the space"),
			c.Row("b.Boxed", "Make the input a boxed block"),
			c.Row("b.Centered", "Alight the file input to the center"),
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
		c.Children(
			c.Row("b.OnCTA(...any)", "Apply children to the ", e.Code(`<span class="file-cta">`), " element"),
			c.Row("b.OnDiv(...any)", "Apply children to the ", e.Code(`<div class="file">`), " element"),
			c.Row("b.OnInput(...any)", "Apply children to the ", e.Code("<input>"), " element"),
			c.Row("string", "Define the call-to-action label"),
			c.Row("b.FileName", "Define the content of the ", e.Code(`<span class="file-name">`), " element"),
			c.Row("b.FileNameAutoUpdate", "Define the content of the ", e.Code(`<span class="file-name">`), " element, which changes when a file is selected"),
			c.Row("e.Element", "Add element to the ", e.Code(`<span class="file-cta">`), " element"),
			c.RowDefault("Apply child to the ", e.Code("<input>"), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/file/",
	c.Example(
		`b.File(
	e.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
)`,
		b.File(
			e.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
		),
	),
).Subsection(
	"Modifiers",
	"https://bulma.io/documentation/form/file/#modifiers",
	c.Example(
		`b.File(
	e.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Right,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Right,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.HorizontalExample(
		`b.File(
	e.Name("resume"),
	b.FullWidth,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.FullWidth,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
)`,
		b.File(
			e.Name("resume"),
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/form/file/#colors",
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Primary,
	fa.Icon(fa.Solid, "upload"),
	"Primary file…",
)`,
		b.File(
			e.Name("resume"),
			b.Primary,
			fa.Icon(fa.Solid, "upload"),
			"Primary file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Info,
	fa.Icon(fa.Solid, "upload"),
	"Info file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Info,
			fa.Icon(fa.Solid, "upload"),
			"Info file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Warning,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Warning file…",
)`,
		b.File(
			e.Name("resume"),
			b.Warning,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Warning file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Danger,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Danger file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Danger,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Danger file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/form/file/#sizes",
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Small,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
)`,
		b.File(
			e.Name("resume"),
			b.Small,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Normal,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
)`,
		b.File(
			e.Name("resume"),
			b.Normal,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Medium,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
)`,
		b.File(
			e.Name("resume"),
			b.Medium,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Large,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
)`,
		b.File(
			e.Name("resume"),
			b.Large,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Small,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Small,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Normal,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Normal,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Medium,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Medium,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Large,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Large,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Small,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
)`,
		b.File(
			e.Name("resume"),
			b.Small,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Normal,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
)`,
		b.File(
			e.Name("resume"),
			b.Normal,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Medium,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
)`,
		b.File(
			e.Name("resume"),
			b.Medium,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Large,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
)`,
		b.File(
			e.Name("resume"),
			b.Large,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Small,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Small,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Normal,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Normal,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Medium,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Medium,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Large,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Large,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
).Subsection(
	"Alignment",
	"https://bulma.io/documentation/form/file/#alignment",
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Centered,
	b.Boxed,
	b.Success,
	fa.Icon(fa.Solid, "upload"),
	"Centered file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Centered,
			b.Boxed,
			b.Success,
			fa.Icon(fa.Solid, "upload"),
			"Centered file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	e.Name("resume"),
	b.Right,
	b.Info,
	fa.Icon(fa.Solid, "upload"),
	"Right file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			e.Name("resume"),
			b.Right,
			b.Info,
			fa.Icon(fa.Solid, "upload"),
			"Right file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
).Subsection(
	"JavaScript",
	"https://bulma.io/documentation/form/file/#javascript",
	b.Content(e.P("In order to automatically change the file name when the user has selected a file, use ", e.Code("b.FileNameAutoUpdate"), " with a placeholder:")),
	c.Example(
		`b.File(
	e.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileNameAutoUpdate("No file uploaded"),
)`,
		b.File(
			e.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileNameAutoUpdate("No file uploaded"),
		),
	),
)
