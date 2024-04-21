package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var formFile = c.NewPage(
	"File", "File upload", "/form/file",
	"",

	b.Content(
		el.P("The ", el.Code("b.File"), " constructor creates a file input. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnCTA(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<span class="file-cta">`), " element"},

			el.Code("b.OnDiv(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="file">`), " element"},

			el.Code("b.OnInput(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<input>"), " element"},

			el.Code("string"),
			"Define the call-to-action label",

			el.Code("b.FileName"),
			"Define the content of the file-name element",

			el.Code("b.FileNameAutoUpdate"),
			"Define the content of the file-name element, which changes when a file is selected",

			el.Code("b.Right"),
			"Move the call-to-action to the right side, align the file input to the right",

			el.Code("b.FullWidth"),
			"Expand the name to fill up the space",

			el.Code("b.Boxed"),
			"Make the input a boxed block",

			el.Code("b.Centered"),
			"Alight the file input to the center",

			el.Code("b.White"),
			"Set file input color to white",

			el.Code("b.Black"),
			"Set file input color to black",

			el.Code("b.Light"),
			"Set file input color to light",

			el.Code("b.Dark"),
			"Set file input color to dark",

			el.Code("b.Primary"),
			"Set file input color to primary",

			el.Code("b.Link"),
			"Set file input color to link",

			el.Code("b.Info"),
			"Set file input color to info",

			el.Code("b.Success"),
			"Set file input color to success",

			el.Code("b.Warning"),
			"Set file input color to warning",

			el.Code("b.Danger"),
			"Set file input color to danger",

			el.Code("b.Small"),
			"Set file input size to small",

			el.Code("b.Normal"),
			"Set file input size to normal",

			el.Code("b.Medium"),
			"Set file input size to medium",

			el.Code("b.Large"),
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
	b.Content(el.P("In order to automatically change the file name when the user has selected a file, use ", el.Code("b.FileNameAutoUpdate"), " with a placeholder:")),
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
