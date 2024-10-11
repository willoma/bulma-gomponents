package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var extending = c.NewPage(
	"Extending", "Extending Bulma-Gomponents", "/extending",
	"",

	b.Content(
		e.P("Here are some developer notes useful when there is a need to extend ", e.Em("Bulma-Gomponents"), "."),
	),
).Section(
	"Implementing icon element", "",

	b.Content(
		e.P("If you need to provide new icon elements to ", e.Em("Bulma-Gomponents"), " functions, you must implement the ", e.Code("b.IconElem"), " interface. This interface is implemented by the elements returned by ", e.Code("b.Icon"), " and ", e.Code("fa.Icon"), "."),
	),
).Section(
	"Implementing classes and styles", "",

	b.Content(
		e.P("If you need to provide new classes and/or styles to ", e.Em("Bulma-Gomponents"), " functions, the following possibilities arise:"),
		e.Ul(
			e.Li("Create a simple constant or variable of type ", e.Code("e.Class"), " - for instance: ", e.Code(`const myClass = e.Class("my-class")`)),
			e.Li("Create a simple constant or variable of type ", e.Code("b.ResponsiveClass"), " if this class accepts the responsive suffixes"),
			e.Li("Implement the ", e.Code("e.Classer"), " interface to return a single class"),
			e.Li("Implement the ", e.Code("e.Classeser"), " interface to return multiple classes"),
			e.Li("Provide ", e.Code("e.Styles"), " to generate additional styles"),
		),
	),
).Section(
	"Implementing e.Element", "",

	b.Content(
		e.P("If you need to create new components for ", e.Em("Bulma-Gomponents"), ", you may prefer following one of these examples:"),
		c.ExamplePre(`func MyElement(children ...any) e.Element {
	m := &myElement{e.Div(e.Class("my-element"))}
	m.With(children...)
	return m
}

type myElement struct {
	e.Element
}

func (m *myElement) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case b.Color:
			m.Element.With(c.Background())
		case []any:
			m.With(c...)
		default:
			m.Element.With(c)
		}
	}
	return m
}

func (m *myElement) Clone() e.Element {
	return &myElement{m.Element.Clone()}
}
`),
		c.ExamplePre(`type weirdOption string

func MyWeirdElement(children ...any) e.Element {
	other := e.Span()
	m := &myWeirdElement{
		container: e.Div(other),
		other: other,
	}
	m.With(children...)
	return m
}

type myWeirdElement struct {
	container e.Element
	other e.Element

	weirdOption weirdOption
}

func (m *myWeirdElement) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class, e.Classer, e.Classeser, e.Styles:
			// Apply classes and styles to content
			m.container.With(c)
		case weirdOption:
			m.weirdOption = weirdOption
		case e.Element:
			// Add element to content
			m.container.With(c)
		case gomponents.Node:
			if b.IsAttribute(c) {
				// Apply gomponents attribute nodes to content
				m.container.With(c)
			} else {
				// Apply other gomponents nodes (ie. elements) to other content
				m.other.With(c)
			}
		case []any:
			m.With(c...)
		default:
			// Apply all other children to some other content
			m.other.With(c)
		}
	}

	return m
}

func (m *myWeirdElement) Render(w io.Writer) error {
	container := m.container.Clone()
	
	if m.weirdOption != "" {
		doSomethingWeird(container, m.other)
	}

	return m.container.Render(w)
}

func (m *myWeirdElement) Clone() e.Element {
	return &myWeirdElement{
		container: m.container.Clone(),
		other: m.other.Clone(),

		weirdOption: m.weirdOption,
	}
}
`),
	),
)
