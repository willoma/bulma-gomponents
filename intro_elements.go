package docs

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var elements = c.NewPage(
	"Elements", "Elements", "/elements",
	"",

	b.Content(
		e.P(e.Em("Bulma-Gomponents"), " relies on ", e.Em("Gomplements"), " to provide elements satisfying the ", e.Code("e.Element"), " interface, which is compatible with ", e.Code("gomponents.Node"), ", making them directly usable as children of other ", e.Em("Gomponents"), " nodes. This interface also provides the ", e.Code("With(...any) e.Element"), " function, which may be used to extend the element after its creation."),
	),
).Section(
	"Document root", "",

	b.Content(
		e.P("In order to generate a document root, use ", e.Code("b.HTML(...any)"), "."),
		c.Modifiers(
			c.Row("b.CSSPath(string)", "Define the path to the CSS file (automatically applied in a ", e.Code(`<link rel="stylesheet">`), "node) - if not defined, the default CDN is used"),
			c.Row("b.Language(string)", "Define the document language"),
			c.Row("b.HTitle(string)", "Define the document title"),
			c.Row("b.Description(string)", "Define the document description"),
			c.Row("b.Head(...gomponents.Node)", "Add one or multiple children to the ", e.Code("<head>")),
			c.RowDefault("Apply child to the ", e.Code("<body>")),
		),
	),
).Subsection(
	"Example", "",
	b.Content(
		c.ExamplePre(`b.HTML(
	b.Script("/htmx.min.js"),
	b.HTitle(p.Title+" | Bulma-Gomponents"),
	b.CSSPath("/bulma.css"),
	fa.CSSHead("/fa"),
	b.Language("en"),
	b.Head(e.Meta(e.Charset("utf-8"))),
	b.Content(
		e.H1("Hello"),
		e.P("Hello world"),
	),
)`),
	),
).Section(
	"Creating elements", "",
	b.Content(
		e.P("Most functions in ", e.Em("Bulma-Gomponents"), " return objects which satisfy the ", e.Code("e.Element"), " interface. These objects may be used as children of each other, or children of the result of ", e.Code("b.HTML"), ", or even children of ", e.Em("gomponents.Node"), " elements."),
	),
	c.Example(
		`import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)
[...]
html.Div(
	html.Class("someClass otherClass"),
	b.Box(
		b.Title3("Here it is"),
		b.IconText(
			fa.Icon(fa.Solid, "lightbulb"),
			"The library",
		),
		" that makes it easier to create web GUIs, based on ",
		e.A(e.Href("https://www.gomponents.com/"), "Gomponents"),
		" and ",
		html.A(html.Href("https://bulma.io/"), gomponents.Text("Bulma")),
		"!",
	),
)`,
		html.Div(
			html.Class("someClass otherClass"),
			b.Box(
				b.Title3("Here it is"),
				b.IconText(
					fa.Icon(fa.Solid, "lightbulb"),
					"The library",
				),
				" that makes it easier to create web GUIs, based on ",
				e.A(e.Href("https://www.gomponents.com/"), "Gomponents"),
				" and ",
				html.A(html.Href("https://bulma.io/"), gomponents.Text("Bulma")),
				"!",
			),
		),
	),
	b.Content(e.P("See ", e.AHref("https://pkg.go.dev/github.com/willoma/gomplements", "the ", e.Em("Gomplements"), " documentation"), " to learn more about the ", e.Code("e.Element"), " interface.")),
).Section(
	"Providing children to elements", "",
	b.Content(
		e.P("In order to provide children to a ", e.Em("Bulma-Gomponents"), " element, the following methods can be used:"),
		e.Ul(
			e.Li("pass the children as direct arguments to the element constructor function"),
			e.Li("pass a slice of children (as a ", e.Code("[]any"), ") as an argument to the constructor function (this method works recursively)"),
			e.Li("pass the children as arguments of the ", e.Code("With"), " function of the element"),
			e.Li("pass a slice of children (as a ", e.Code("[]any"), ") as an argument to the ", e.Code("With"), " function of the element"),
		),
		e.P("The following examples are all equivalent:"),
	),
	c.Example(
		`e.Div(
	e.Strong("Hello"), " ",
	b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
	"...",
),
e.Div(
	[]any{
		e.Strong("Hello"), " ",
		b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
	},
	"...",
),
e.Div(e.Strong("Hello"), " ").With(
	b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
	"...",
),
e.Div(e.Strong("Hello"), " ").
	With(b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe"))).
	With("..."),
e.Div(
	[]any{e.Strong("Hello"), " "},
	b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
).With("..."),
e.Div().
	With(e.Strong("Hello"), " ").
	With(b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe"))).
	With("..."),
e.Div().
	With(
		[]any{e.Strong("Hello"), " "},
		b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
	).
	With("...")`,
		e.Div(
			e.Strong("Hello"), " ",
			b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
			"...",
		),
		e.Div(
			[]any{
				e.Strong("Hello"), " ",
				b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
			},
			"...",
		),
		e.Div(e.Strong("Hello"), " ").With(
			b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
			"...",
		),
		e.Div(e.Strong("Hello"), " ").
			With(b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe"))).
			With("..."),
		e.Div(
			[]any{e.Strong("Hello"), " "},
			b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
		).With("..."),
		e.Div().
			With(e.Strong("Hello"), " ").
			With(b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe"))).
			With("..."),
		e.Div().
			With(
				[]any{e.Strong("Hello"), " "},
				b.IconText(e.Em("world"), fa.Icon(fa.Solid, "globe")),
			).
			With("..."),
	),
).Section(
	"Cloning and preparing elements", "",

	b.Content(
		e.P("If you need to clone an element generated by ", e.Em("Bulma-Gomponents"), ", you can call the ", e.Code("Clone()"), " method on that element. It ensures any modification made to the clone is independent from the original element. It may be useful to generate multiple components that look the same and have only a few different parameters. Moreover, some elements have attributes that are generated on render-time, depending on the children you have already provided. Therefore, if you need to render an element and modify it later in order to re-render it with more content, simply clone it before rendering and modify its clone."),
		e.P("On the other hand, if you have a component that will never be cloned and that will be re-used multiple times, you can prepare its rendering with the ", e.Code("e.Prepare(e.Element)"), " function, which returns an object which implements ", e.Code("gomponents.Node"), " with the element's full content already pre-rendered in memory: further calls to ", e.Code("Render(io.Writer)"), " on the prepared element will simply copy the pre-rendered content to the provided writer. It may be useful for content you reuse often, like icons or static menus."),
	),
)
