package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var extending = c.NewPage(
	"Extending", "Extending Bulma-Gomponents", "/extending",
	"",

	b.Content(
		el.P("Here are some developer notes useful when there is a need to extend ", el.Em("Bulma-Gomponents"), "."),
	),
).Section(
	"Implementing icon element", "",

	b.Content(
		el.P("If you need to provide new icon elements to ", el.Em("Bulma-Gomponents"), " functinos, you must implement the ", el.Code("b.IconElem"), " interface. This interface is implemented by the elements returned by ", el.Code("b.Icon"), " and ", el.Code("fa.Icon"), "."),
	),
).Section(
	"Implementing classes and styles", "",

	b.Content(
		el.P("If you need to provide new classes and/or styles to ", el.Em("Bulma-Gomponents"), " functions, the following possibilities arise:"),
		el.Ul(
			el.Li("Create a simple constant or variable with type ", el.Code("b.Class"), " - for instance: ", el.Code(`const myClass = b.Class("my-class")`)),
			el.Li("Create a collection of classes, of which some can be made responsibe, with type ", el.Code("b.MultiClass")),
			el.Li("Implement the ", el.Code("b.ExternalClass"), " interface to return a single class"),
			el.Li("Implement the ", el.Code("b.ExternalClassesAndStyles"), " interface to return multiple classes and/or styles"),
			el.Li("Execute the ", el.Code("b.Style"), " function to generate additional styles"),
		),
	),
).Section(
	"Implementing Element", "",

	b.Content(
		el.P("If you need to create new components for ", el.Em("Bulma-Gomponents"), ", you may prefer following this example:"),
		el.Pre(`func MyElement(children ...any) Element {
	return new(myElement).With(children...)
}

type myElement struct {
	someContent      []any
	someOtherContent []any
}

func (m *myElement) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case b.Class, b.ColorClass, b.ExternalClass, b.ExternalClassesAndStyles, b.MultiClass, b.Styles:
			// Apply classes and styles to some content
		case b.Element:
			// Add element to some content
		case gomponents.Node:
			if b.IsAttribute(c) {
				// Apply gomponents attribute nodes to some content
			} else {
				// Apply other gomponents nodes (ie. elements) to some other content
			}
		case []any:
			m.With(c...)
		default:
			// Apply all other children to some other content
		}
	}

	return m
}

func (m *myElement) Render(w io.Writer) error {
	// call b.Elem with the appropriate arguments according to the component needs
	e := b.Elem(html.Div) // ...

	// If needed, process children of myElement and apply them wherever needed

	return e.Render(w)
}
`),
	),
)
