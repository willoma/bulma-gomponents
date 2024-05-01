package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var formFile = c.NewPage(
	"File", "File upload", "/form/file",
	"",

	b.Content(
		e.P("The ", e.Code("b.File"), " constructor creates a file input. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnCTA(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<span class="file-cta">`), " e.Element"},

			e.Code("b.OnDiv(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="file">`), " e.Element"},

			e.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<input>"), " e.Element"},

			e.Code("string"),
			"Define the call-to-action label",

			e.Code("b.FileName"),
			"Define the content of the file-name e.Element",

			e.Code("b.FileNameAutoUpdate"),
			"Define the content of the file-name e.Element, which changes when a file is selected",

			e.Code("b.Right"),
			"Move the call-to-action to the right side, align the file input to the right",

			e.Code("b.FullWidth"),
			"Expand the name to fill up the space",

			e.Code("b.Boxed"),
			"Make the input a boxed block",

			e.Code("b.Centered"),
			"Alight the file input to the center",

			e.Code("b.White"),
			"Set file input color to white",

			e.Code("b.Black"),
			"Set file input color to black",

			e.Code("b.Light"),
			"Set file input color to light",

			e.Code("b.Dark"),
			"Set file input color to dark",

			e.Code("b.Primary"),
			"Set file input color to primary",

			e.Code("b.Link"),
			"Set file input color to link",

			e.Code("b.Info"),
			"Set file input color to info",

			e.Code("b.Success"),
			"Set file input color to success",

			e.Code("b.Warning"),
			"Set file input color to warning",

			e.Code("b.Danger"),
			"Set file input color to danger",

			e.Code("b.Small"),
			"Set file input size to small",

			e.Code("b.Normal"),
			"Set file input size to normal",

			e.Code("b.Medium"),
			"Set file input size to medium",

			e.Code("b.Large"),
			"Set file input size to large",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/form/file/",
	c.Example(
		`b.File(
	html.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
)`,
		b.File(
			html.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
		),
	),
).Subsection(
	"Modifiers",
	"https://bulma.io/documentation/form/file/#modifiers",
	c.Example(
		`b.File(
	html.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Right,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Right,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.HorizontalExample(
		`b.File(
	html.Name("resume"),
	b.FullWidth,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.FullWidth,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
)`,
		b.File(
			html.Name("resume"),
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
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
	html.Name("resume"),
	b.Primary,
	fa.Icon(fa.Solid, "upload"),
	"Primary file…",
)`,
		b.File(
			html.Name("resume"),
			b.Primary,
			fa.Icon(fa.Solid, "upload"),
			"Primary file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Info,
	fa.Icon(fa.Solid, "upload"),
	"Info file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Info,
			fa.Icon(fa.Solid, "upload"),
			"Info file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Warning,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Warning file…",
)`,
		b.File(
			html.Name("resume"),
			b.Warning,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Warning file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Danger,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Danger file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
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
	html.Name("resume"),
	b.Small,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
)`,
		b.File(
			html.Name("resume"),
			b.Small,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Normal,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
)`,
		b.File(
			html.Name("resume"),
			b.Normal,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Medium,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
)`,
		b.File(
			html.Name("resume"),
			b.Medium,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Large,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
)`,
		b.File(
			html.Name("resume"),
			b.Large,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Small,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Small,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Normal,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Normal,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Medium,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Medium,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Large,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Large,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Small,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
)`,
		b.File(
			html.Name("resume"),
			b.Small,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Normal,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
)`,
		b.File(
			html.Name("resume"),
			b.Normal,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Medium,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
)`,
		b.File(
			html.Name("resume"),
			b.Medium,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Large,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
)`,
		b.File(
			html.Name("resume"),
			b.Large,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Large file…",
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Small,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Small file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Small,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Small file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Normal,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Normal file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Normal,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Normal file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Medium,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Medium file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
			b.Medium,
			b.Boxed,
			fa.Icon(fa.Solid, "upload"),
			"Medium file…",
			b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
		),
	),
	c.Example(
		`b.File(
	html.Name("resume"),
	b.Large,
	b.Boxed,
	fa.Icon(fa.Solid, "upload"),
	"Large file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
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
	html.Name("resume"),
	b.Centered,
	b.Boxed,
	b.Success,
	fa.Icon(fa.Solid, "upload"),
	"Centered file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
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
	html.Name("resume"),
	b.Right,
	b.Info,
	fa.Icon(fa.Solid, "upload"),
	"Right file…",
	b.FileName("Screen Shot 2017-07-29 at 15.54.25.png"),
)`,
		b.File(
			html.Name("resume"),
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
	html.Name("resume"),
	fa.Icon(fa.Solid, "upload"),
	"Choose a file…",
	b.FileNameAutoUpdate("No file uploaded"),
)`,
		b.File(
			html.Name("resume"),
			fa.Icon(fa.Solid, "upload"),
			"Choose a file…",
			b.FileNameAutoUpdate("No file uploaded"),
		),
	),
)
