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
		e.P("The ", e.Code("b.Tabs"), " constructor creates a tabs section. Its link children must be created with the ", e.Code("b.TabLink"), " constructor. The following children have a special meaning:"),
		b.DList(
			e.Code("b.OnTabs(...)"),
			[]any{"Force childen to be applied to the ", e.Code(`<div class="tabs">`), " e.Element"},

			e.Code("b.OnUl(...)"),
			[]any{"Force childen to be applied to the ", e.Code("<ul>"), " e.Element"},

			[]any{"one of the class or style types defined in package ", e.Code("b")},
			[]any{"Apply the class or style to the ", e.Code(`<div class="tabs">`), " e.Element"},

			e.Code("b.Container()"),
			"Use this e.Element as an intermediate container",

			e.Code("b.Centered"),
			"Center the tabs",

			e.Code("b.Right"),
			"Align the tabs to the right",

			e.Code("b.Boxed"),
			"Draw boxed tabs",

			e.Code("b.Toggle"),
			"Button-looking tabs",

			e.Code("b.ToggleRounded"),
			"Rounded button-looking tabs",

			e.Code("b.FullWidth"),
			"Take the whole width",

			e.Code("b.Small"),
			"Set tabs size to small",

			e.Code("b.Medium"),
			"Set tabs size to medium",

			e.Code("b.Large"),
			"Set tabs size to large",
		),
		e.P("Other children are added to the ", e.Code("<ul>"), " e.Element."),
		e.P("The ", e.Code("b.TabLink"), " and ", e.Code("b.TabAHref"), " constructors create tab entries which are links. The following children have a special meaning:"),
		b.DList(
			e.Code("b.Active"),
			[]any{"Apply the ", e.Code("is-active"), " class to the ", e.Code("<li>"), " e.Element"},
		),
		e.P("Other children are added to the ", e.Code("<a>"), " e.Element."),
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
