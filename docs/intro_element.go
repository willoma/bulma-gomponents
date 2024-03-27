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
		el.P(el.Em("Bulma-Gomponents"), " elements satisfy the ", el.Code("b.Element"), " interface, which is compatible with ", el.Code("gomponents.Node"), ", making them directly usable as children of other ", el.Em("Gomponents"), " nodes. This interface also provides the ", el.Code("With(...any) Element"), " function, which may be used to extend the element after its creation."),
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
).Subsection(
	"Inner and Outer markers", "",

	b.Content(
		el.P("Some constructors recognize the inner or outer markers in order to force children to apply to a specific element generated by the constructor. This distinction is detailed in the constructors' documentations. In order to mark a child, use the ", el.Code("b.Inner"), " or ", el.Code("b.Outer"), "functions."),
	),
).Section(
	"Recognized children", "",
	b.Content(
		el.P("All ", el.Em("Bulma-Gomponents"), " elements recognize the following children types, either as constructor arguments or arguments to the ", el.Code("With"), " function :"),
		b.DList(
			el.Code("b.Class"),
			"apply the class",

			el.Code("b.ColorClass"),
			"apply the class, as its light or dark variant when required",

			el.Code("b.ExternalClass"),
			"apply the class",

			el.Code("b.ExternalClassesAndStyles"),
			"apply the classes and styles",

			el.Code("b.MultiClass"),
			"apply the classes",

			[]any{"Result of ", el.Code("b.Style")},
			"apply the styles",

			el.Code("string"),
			"append the string as a text node",

			[]any{el.Code("b.IconElem"), " (ie. result of ", el.Code("b.Icon"), " or ", el.Code("fa.Icon"), ")"},
			"append the icon",

			[]any{"Result of ", el.Code("b.TopNavbar"), " or ", el.Code("b.BottomNavbar")},
			"append the navbar and mark the element as having a top or bottom navbar (useful only directly on the body)",

			[]any{"Result of ", el.Code("b.Tile")},
			"append the tile and mark it as an ancestor if the current element is not a tile",

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
)
