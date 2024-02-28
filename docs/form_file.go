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
	"https://bulma.io/documentation/form/file/",
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
).Section(
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
).Section(
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
).Section(
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
).Section(
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
).Section(
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
