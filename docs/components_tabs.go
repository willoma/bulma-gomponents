package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

var tabs = c.NewPage(
	"Tabs", "Tabs", "/tabs",
	"https://bulma.io/documentation/components/tabs/",
	c.Example(
		`b.Tabs(
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
).Section(
	"Alignment",
	"https://bulma.io/documentation/components/tabs/#alignment",
	c.Example(
		`b.Tabs(
	b.Centered,
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Right,
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.Right,
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
).Section(
	"Icons",
	"https://bulma.io/documentation/components/tabs/#icons",
	c.Example(
		`b.Tabs(
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/components/tabs/#sizes",
	c.Example(
		`b.Tabs(
	b.Small,
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.Small,
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Medium,
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.Medium,
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Large,
	b.TabsLink(b.Active, "Pictures"),
	b.TabsLink("Music"),
	b.TabsLink("Videos"),
	b.TabsLink("Documents"),
)`,
		b.Tabs(
			b.Large,
			b.TabsLink(b.Active, "Pictures"),
			b.TabsLink("Music"),
			b.TabsLink("Videos"),
			b.TabsLink("Documents"),
		),
	),
).Section(
	"Styles",
	"https://bulma.io/documentation/components/tabs/#styles",
	c.Example(
		`b.Tabs(
	b.Boxed,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Boxed,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.ToggleRounded,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.ToggleRounded,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.FullWidth,
	b.TabsLink(fa.Icon(fa.Solid, "angle-left"), "Left"),
	b.TabsLink(fa.Icon(fa.Solid, "angle-up"), "Up"),
	b.TabsLink(fa.Icon(fa.Solid, "angle-right"), "Right"),
)`,
		b.Tabs(
			b.FullWidth,
			b.TabsLink(fa.Icon(fa.Solid, "angle-left"), "Left"),
			b.TabsLink(fa.Icon(fa.Solid, "angle-up"), "Up"),
			b.TabsLink(fa.Icon(fa.Solid, "angle-right"), "Right"),
		),
	),
).Section(
	"Combining",
	"https://bulma.io/documentation/components/tabs/#combining",
	c.Example(
		`b.Tabs(
	b.Centered,
	b.Boxed,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.Boxed,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.FullWidth,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.FullWidth,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Centered,
	b.Boxed,
	b.Medium,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.Boxed,
			b.Medium,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.FullWidth,
	b.Large,
	b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.FullWidth,
			b.Large,
			b.TabsLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabsLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabsLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabsLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
)
