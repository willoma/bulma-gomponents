package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var tabs = c.NewPage(
	"Tabs", "Tabs", "/tabs",
	"",

	b.Content(
		e.P("The ", e.Code("b.Tabs"), " constructor creates a tabs container. Its link children must be created with the ", e.Code("b.TabLink"), " constructor."),
		c.Modifiers(
			c.Row("b.Centered", "Center the tabs"),
			c.Row("b.Right", "Align the tabs to the right"),
			c.Row("b.Boxed", "Draw boxed tabs"),
			c.Row("b.Toggle", "Button-looking tabs"),
			c.Row("b.ToggleRounded", "Rounded button-looking tabs"),
			c.Row("b.FullWidth", "Take the whole width"),
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),
		c.Children(
			c.Row("b.OnTabs(...any)", "Apply children to the ", e.Code(`<div class="tabs">`), " element"),
			c.Row("b.OnUl(...any)", "Apply children to the ", e.Code("<ul>"), " element"),
			c.RowClassesStyles("Apply child to the ", e.Code(`<div class="tabs">`), " element"),
			c.Row("b.Container(...any)", "Use element as an intermediate container"),
			c.RowDefault("Apply children to the ", e.Code("<ul>"), " element"),
		),
		e.P("The ", e.Code("b.TabLink"), " and ", e.Code("b.TabAHref"), " constructors create tab entries which are links."),
		c.Modifiers(
			c.Row("b.Active", "Apply the ", e.Code("is-active"), " class to the ", e.Code("<li>"), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/tabs/",
	c.Example(
		`b.Tabs(
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
).Section(
	"Alignment",
	"https://bulma.io/documentation/components/tabs/#alignment",
	c.Example(
		`b.Tabs(
	b.Centered,
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Right,
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.Right,
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
).Section(
	"Icons",
	"https://bulma.io/documentation/components/tabs/#icons",
	c.Example(
		`b.Tabs(
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/components/tabs/#sizes",
	c.Example(
		`b.Tabs(
	b.Small,
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.Small,
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Medium,
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.Medium,
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Large,
	b.TabLink(b.Active, "Pictures"),
	b.TabLink("Music"),
	b.TabLink("Videos"),
	b.TabLink("Documents"),
)`,
		b.Tabs(
			b.Large,
			b.TabLink(b.Active, "Pictures"),
			b.TabLink("Music"),
			b.TabLink("Videos"),
			b.TabLink("Documents"),
		),
	),
).Section(
	"Styles",
	"https://bulma.io/documentation/components/tabs/#styles",
	c.Example(
		`b.Tabs(
	b.Boxed,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Boxed,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.ToggleRounded,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.ToggleRounded,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.FullWidth,
	b.TabLink(fa.Icon(fa.Solid, "angle-left"), "Left"),
	b.TabLink(fa.Icon(fa.Solid, "angle-up"), "Up"),
	b.TabLink(fa.Icon(fa.Solid, "angle-right"), "Right"),
)`,
		b.Tabs(
			b.FullWidth,
			b.TabLink(fa.Icon(fa.Solid, "angle-left"), "Left"),
			b.TabLink(fa.Icon(fa.Solid, "angle-up"), "Up"),
			b.TabLink(fa.Icon(fa.Solid, "angle-right"), "Right"),
		),
	),
).Section(
	"Combining",
	"https://bulma.io/documentation/components/tabs/#combining",
	c.Example(
		`b.Tabs(
	b.Centered,
	b.Boxed,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.Boxed,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.FullWidth,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.FullWidth,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Centered,
	b.Boxed,
	b.Medium,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Centered,
			b.Boxed,
			b.Medium,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
	c.Example(
		`b.Tabs(
	b.Toggle,
	b.FullWidth,
	b.Large,
	b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
	b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
	b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
	b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
)`,
		b.Tabs(
			b.Toggle,
			b.FullWidth,
			b.Large,
			b.TabLink(b.Active, fa.Icon(fa.Solid, "image"), "Pictures"),
			b.TabLink(fa.Icon(fa.Solid, "music"), "Music"),
			b.TabLink(fa.Icon(fa.Solid, "film"), "Videos"),
			b.TabLink(fa.Icon(fa.Solid, "file-alt"), "Documents"),
		),
	),
)
