package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var tabs = c.NewPage(
	"Tabs", "Tabs", "/tabs",
	"",

	b.Content(
		el.P("The ", el.Code("b.Tabs"), " constructor creates a tabs section. Its link children must be created with the ", el.Code("b.TabLink"), " constructor. The following children have a special meaning:"),
		b.DList(
			el.Code("b.OnTabs(...)"),
			[]any{"Force childen to be applied to the ", el.Code(`<div class="tabs">`), " element"},

			el.Code("b.OnUl(...)"),
			[]any{"Force childen to be applied to the ", el.Code("<ul>"), " element"},

			[]any{"one of the class or style types defined in package ", el.Code("b")},
			[]any{"Apply the class or style to the ", el.Code(`<div class="tabs">`), " element"},

			el.Code("b.Container()"),
			"Use this element as an intermediate container",

			el.Code("b.Centered"),
			"Center the tabs",

			el.Code("b.Right"),
			"Align the tabs to the right",

			el.Code("b.Boxed"),
			"Draw boxed tabs",

			el.Code("b.Toggle"),
			"Button-looking tabs",

			el.Code("b.ToggleRounded"),
			"Rounded button-looking tabs",

			el.Code("b.FullWidth"),
			"Take the whole width",

			el.Code("b.Small"),
			"Set tabs size to small",

			el.Code("b.Medium"),
			"Set tabs size to medium",

			el.Code("b.Large"),
			"Set tabs size to large",
		),
		el.P("Other children are added to the ", el.Code("<ul>"), " element."),
		el.P("The ", el.Code("b.TabLink"), " and ", el.Code("b.TabAHref"), " constructors create tab entries which are links. The following children have a special meaning:"),
		b.DList(
			el.Code("b.Active"),
			[]any{"Apply the ", el.Code("is-active"), " class to the ", el.Code("<li>"), " element"},
		),
		el.P("Other children are added to the ", el.Code("<a>"), " element."),
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
