package docs

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var elements = c.NewPage(
	"Elements", "Elements", "/elements",
	"",

	b.Content(
		el.P(el.Em("Bulma-Gomponents"), " elements satisfy the ", el.Code("b.Element"), " interface, which is compatible with ", el.Code("gomponents.Node"), ", making them directly usable as children of other ", el.Em("Gomponents"), " nodes. This interface also provides the ", el.Code("With(...any) b.Element"), " function, which may be used to extend the element after its creation."),
	),
).Section(
	"Document root", "",

	b.Content(
		el.P("In order to generate a document root, use ", el.Code("b.HTML(...any)"), ". The allowed options are:"),
		b.DList(
			el.Code("b.CSSPath(string)"),
			[]any{"path to the CSS file, automatically applied in a ", el.Code(`<link rel="stylesheet">`), " node (if it is not defined, the default CDN is used)"},
			el.Code("b.Language(string)"),
			"define the document language",
			el.Code("b.HTitle(string)"),
			"define the document title",
			el.Code("b.Description(string)"),
			"define the document description",
			el.Code("b.Head(...gomponents.Node)"),
			[]any{"add one or multiple children to the ", el.Code("<head>")},
			"any other value",
			[]any{"add the node to the ", el.Code("<body>")},
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
	b.Head(html.Meta(html.Charset("utf-8"))),
	b.Content(
		el.H1("Hello"),
		el.P("Hello world"),
	),
)`),
	),
).Section(
	"Creating elements", "",
	b.Content(
		el.P("Most functions in ", el.Em("Bulma-Gomponents"), " return objects which satisfy the ", el.Code("b.Element"), " interface. These objects may be used as children of each other, or children of the result of ", el.Code("b.HTML"), ", or even children of ", el.Em("gomponents.Node"), " elements."),
	),
	c.Example(
		`import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)
[...]
h.Div(
	h.Class("someClass otherClass"),
	b.Box(
		b.Title3("Here it is"),
		b.IconText(
			fa.Icon(fa.Solid, "lightbulb"),
			"The library",
		),
		" that makes it easier to create web GUIs, based on ",
		el.A(html.Href("https://www.gomponents.com/"), "Gomponents"),
		" and ",
		html.A(h.Href("https://bulma.io/"), gomponents.Text("Bulma")),
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
				el.A(html.Href("https://www.gomponents.com/"), "Gomponents"),
				" and ",
				html.A(html.Href("https://bulma.io/"), gomponents.Text("Bulma")),
				"!",
			),
		),
	),
	b.Content(
		el.P("The ", el.Code("b.ParentModifier"), " interface describes an element and additionally allows modifying its parent in the rare situation where it is needed (for instance, in order to add a class to it)."),
		el.P("When you need to create an element from scratch, you can also use the ", el.Code("b.Elem"), " function, which returns a blank element of the provided type (for example,", el.Code(`b.Elem(html.Div, "Hello world")`), "), however the ", el.Code("el"), " package already provides such helpers for all known HTML elements (for example, ", el.Code(`el.Div("Hello world")`), ")."),
	),
).Section(
	"Providing children to elements", "",
	b.Content(
		el.P("In order to provide children to a ", el.Em("Bulma-Gomponents"), " element, the following methods can be used:"),
		el.Ul(
			el.Li("pass the children as direct arguments to the element constructor function"),
			el.Li("pass a slice of children (as a ", el.Code("[]any"), ") as an argument to the constructor function (this method works recursively)"),
			el.Li("pass the children as arguments of the ", el.Code("With"), " function of the element"),
			el.Li("pass a slice of children (as a ", el.Code("[]any"), ") as an argument to the ", el.Code("With"), " function of the element"),
		),
		el.P("The following examples are all equivalent:"),
	),
	c.Example(
		`el.Div(
	el.Strong("Hello"), " ",
	b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
	"...",
),
el.Div(
	[]any{
		el.Strong("Hello"), " ",
		b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
	},
	"...",
),
el.Div(el.Strong("Hello"), " ").With(
	b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
	"...",
),
el.Div(el.Strong("Hello"), " ").
	With(b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe"))).
	With("..."),
el.Div(
	[]any{el.Strong("Hello"), " "},
	b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
).With("..."),
el.Div().
	With(el.Strong("Hello"), " ").
	With(b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe"))).
	With("..."),
el.Div().
	With(
		[]any{el.Strong("Hello"), " "},
		b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
	).
	With("...")`,
		el.Div(
			el.Strong("Hello"), " ",
			b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
			"...",
		),
		el.Div(
			[]any{
				el.Strong("Hello"), " ",
				b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
			},
			"...",
		),
		el.Div(el.Strong("Hello"), " ").With(
			b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
			"...",
		),
		el.Div(el.Strong("Hello"), " ").
			With(b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe"))).
			With("..."),
		el.Div(
			[]any{el.Strong("Hello"), " "},
			b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
		).With("..."),
		el.Div().
			With(el.Strong("Hello"), " ").
			With(b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe"))).
			With("..."),
		el.Div().
			With(
				[]any{el.Strong("Hello"), " "},
				b.IconText(el.Em("world"), fa.Icon(fa.Solid, "globe")),
			).
			With("..."),
	),
).Section(
	"Recognized children", "",
	b.Content(
		el.P("All ", el.Em("Bulma-Gomponents"), " elements recognize the following children types, either as constructor arguments or arguments to the ", el.Code("With"), " function :"),
		b.DList(
			el.Code("func(...gomponents.Node) gomponents.Node"),
			"Use this constructor function to create the element",

			el.Code("b.Class"),
			"Apply the class",

			el.Code("b.Classer"),
			[]any{"Apply the class returned by the child's ", el.Code("Class()"), " function"},

			el.Code("b.Classeser"),
			[]any{"Apply the classes returned by the child's ", el.Code("Classes()"), " function"},

			[]any{"Result of ", el.Code("b.Style")},
			"Apply the styles",

			el.Code("b.ID"),
			"Define the element's ID in the HTML DOM",

			el.Code("string"),
			"Append the string as a text node",

			[]any{el.Code("b.IconElem"), " (ie. result of ", el.Code("b.Icon"), " or ", el.Code("fa.Icon"), ")"},
			"Append the icon",

			el.Code("b.ParentModifier"),
			[]any{"Modify the current element by calling the ", el.Code("ModifyParent()"), " function on the child element, and append the element"},

			el.Code("b.Element"),
			"append the element",

			[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
			"apply the attribute",

			[]any{"Other ", el.Code("gomponents.Node")},
			"append the element",

			el.Code("[]any"),
			"apply or append each of the contained objects (recursively)",

			el.Code("fmt.Stringer"),
			"append the string as a text node",
		),
		el.P("Specific elements may also recognize other children types, see the other documentation pages. Any other child type is ignored."),
	),
).Section(
	"Cloning ans preparing elements", "",

	b.Content(
		el.P("If you need to clone an element generated by ", el.Em("Bulma-Gomponents"), ", you can call the ", el.Code("Clone()"), " method on that element. It ensures any modification made to the clone is independent from the original element. It may be useful to generate multiple components that look the same and have only a few different parameters. Moreover, some elements have attributes that are generated on render-time, depending on the children you have already provided. Therefore, if you need to render an element and modify it later in order to re-render it with more content, simply clone it before rendering and modify its clone."),
		el.P("On the other hand, if you have a component that will never be cloned and that will be re-used multiple times, you can prepare its rendering with the ", el.Code("Prepare()"), " method, which returns an object which implements ", el.Code("gomponents.Node"), " and has its full content already pre-rendered in memory: next calls to ", el.Code("Render(io.Writer)"), " will simply copy the pre-rendered content to the provided writer. It may be useful for content you reuse often, like icons or static menus."),
	),
)
