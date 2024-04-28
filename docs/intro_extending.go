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
		el.P("If you need to provide new icon elements to ", el.Em("Bulma-Gomponents"), " functions, you must implement the ", el.Code("b.IconElem"), " interface. This interface is implemented by the elements returned by ", el.Code("b.Icon"), " and ", el.Code("fa.Icon"), "."),
	),
).Section(
	"Implementing classes and styles", "",

	b.Content(
		el.P("If you need to provide new classes and/or styles to ", el.Em("Bulma-Gomponents"), " functions, the following possibilities arise:"),
		el.Ul(
			el.Li("Create a simple constant or variable of type ", el.Code("b.Class"), " - for instance: ", el.Code(`const myClass = b.Class("my-class")`)),
			el.Li("Create a simple constant or variable of type ", el.Code("b.ResponsiveClass"), " if this class accepts the responsive suffixes"),
			el.Li("Implement the ", el.Code("b.Classer"), " interface to return a single class"),
			el.Li("Implement the ", el.Code("b.Classeser"), " interface to return multiple classes"),
			el.Li("Execute the ", el.Code("b.Style"), " function to generate additional styles"),
		),
	),
).Section(
	"Implementing Element", "",

	b.Content(
		el.P("If you need to create new components for ", el.Em("Bulma-Gomponents"), ", you may prefer following one of these examples:"),
		c.ExamplePre(`func MyElement(children ...any) b.Element {
	m := &myElement{el.Div(b.Class("my-element"))}
	m.With(children...)
	return m
}

type myElement struct {
	b.Element
}

func (m *myElement) With(children ...any) Element {
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

func (m *myElement) Clone() Element {
	return &myElement{m.Element.Clone()}
}
`),
		c.ExamplePre(`type weirdOption string

func MyWeirdElement(children ...any) b.Element {
	other := b.Elem(html.Span)
	m := &myWeirdElement{
		container: b.Elem(html.Div, other),
		other: other,
	}
	m.With(children...)
	return m
}

type myWeirdElement struct {
	container b.Element
	other b.Element

	weirdOption weirdOption

	rendered sync.Once
}

func (m *myWeirdElement) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case b.Class, b.Classer, b.Classeser, b.Styles:
			// Apply classes and styles to content
			m.container.With(c)
		case weirdOption:
			b.weirdOption = weirdOption
		case b.Element:
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
	m.rendered.Do(func() {
		if m.weirdOption != "" {
			doSomethingWeird(m.container, m.other)
		}
	})

	return m.container.Render(w)
}

func (m *myWeirdElement) Clone() Element {
	return &myWeirdElement{
		container: m.container.Clone(),
		other: m.other.Clone(),

		weirdOption: m.weirdOption,
	}
}
`),
	),
)
