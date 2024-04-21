package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Card creates a card.
func Card(children ...any) Element {
	return new(card).With(children...)
}

type card struct {
	children []any
	header   []any
	footer   []any
}

func (c *card) With(children ...any) Element {
	var currentContent []any

	for _, ch := range children {
		switch ch := ch.(type) {
		case onCard:
			c.children = append(c.children, ch...)
		case onContent:
			currentContent = append(currentContent, ch...)
		case splitContent:
			if currentContent != nil {
				c.children = append(c.children, Elem(html.Div, Class("card-content"), currentContent))
				currentContent = nil
			}
		case cardHeader:
			c.header = append(c.header, ch...)
		case *cardHeaderIcon:
			c.header = append(c.header, ch)
		case *cardHeaderTitle:
			c.header = append(c.header, ch)
		case cardFooter:
			c.footer = append(c.footer, ch...)
		case Class, Classer, Classeser, ExternalClassesAndStyles, MultiClass, Styles:
			c.children = append(c.children, ch)
		case *cardImage, *cardImageImg:
			if currentContent != nil {
				c.children = append(c.children, Elem(html.Div, Class("card-content"), currentContent))
				currentContent = nil
			}
			c.children = append(c.children, ch)
		case Element:
			currentContent = append(currentContent, ch)
		case gomponents.Node:
			if IsAttribute(c) {
				c.children = append(c.children, ch)
			} else {
				currentContent = append(currentContent, ch)
			}
		case []any:
			currentContent = append(currentContent, ch...)

		default:
			currentContent = append(currentContent, ch)
		}
	}

	if currentContent != nil {
		c.children = append(c.children, Elem(html.Div, Class("card-content"), currentContent))
	}
	return c
}

func (c *card) Render(w io.Writer) error {
	e := Elem(html.Div, Class("card"))
	if c.header != nil {
		e.With(Elem(html.Header, Class("card-header"), c.header))
	}
	if c.children != nil {
		e.With(c.children...)
	}
	if c.footer != nil {
		e.With(Elem(html.Footer, Class("card-footer"), c.footer))
	}
	return e.Render(w)
}

func SplitContent() splitContent {
	return splitContent{}
}

type splitContent struct{}

// CardHeader marks its children as being part of a card header.
//
// If a child is a string, it is wrapped into a CardHeaderTitle element.
//
// If a child is an Element with class "icon", it is wrapped into a
// CardHeaderIcon element.
func CardHeader(children ...any) cardHeader {
	return cardHeader{}.with(children...)
}

type cardHeader []any

func (ch cardHeader) with(children ...any) cardHeader {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			ch = append(ch, CardHeaderIcon(c))
		case string:
			ch = append(ch, CardHeaderTitle(c))
		case []any:
			ch.with(c...)
		default:
			ch = append(ch, c)
		}
	}

	return ch
}

// CardHeaderIcon creates an icon for a card header.
//
// If you provide an icon to the CardHeader function (for instance the result of the Icon function), you don't need CardHeaderIcon.
func CardHeaderIcon(children ...any) Element {
	return new(cardHeaderIcon).With(children...)
}

type cardHeaderIcon struct {
	children []any
}

func (chi *cardHeaderIcon) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			chi.With(c...)
		default:
			chi.children = append(chi.children, c)
		}
	}
	return chi
}

func (c *cardHeaderIcon) Render(w io.Writer) error {
	return Elem(html.Button, Class("card-header-icon"), c.children).Render(w)
}

// CardHeaderTitle creates a title for a card header.
//
// The following modifiers change the title behaviour:
//   - Centered: center the title
//
// If you provide a string to the CardHeader function, you don't need
// CardHeaderTitle.
func CardHeaderTitle(children ...any) Element {
	return new(cardHeaderTitle).With(children...)
}

type cardHeaderTitle struct {
	children []any
}

func (cht *cardHeaderTitle) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			cht.With(c...)
		default:
			cht.children = append(cht.children, c)
		}
	}
	return cht
}

func (cht *cardHeaderTitle) Render(w io.Writer) error {
	return Elem(html.P, Class("card-header-title"), cht.children).Render(w)
}

// CardFooter marks its children as being part of a card footer.
//
// It add the "card-footer-item" class to all its Element children.
func CardFooter(children ...any) cardFooter {
	return cardFooter{}.with(children...)
}

type cardFooter []any

func (cf cardFooter) with(children ...any) cardFooter {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			cf = append(cf, c.With(Class("card-footer-item")))
		case []any:
			cf.with(c...)
		default:
			cf = append(cf, c)
		}
	}

	return cf
}

// CardImage creates a card image.
func CardImage(children ...any) Element {
	return new(cardImage).With(children...)
}

type cardImage struct {
	children    []any
	imgChildren []any
}

func (ci *cardImage) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCardImage:
			ci.children = append(ci.children, c)
		case []any:
			ci.With(c...)
		default:
			ci.imgChildren = append(ci.imgChildren, c)
		}
	}

	return ci
}

func (ci *cardImage) Render(w io.Writer) error {
	return Elem(html.Div, Class("card-image"), Image(ci.imgChildren...), ci.children).Render(w)
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// See the ImageImg function documentation for the list of allowed modifiers.
func CardImageImg(src string, children ...any) Element {
	return (&cardImageImg{src: src}).With(children...)
}

type cardImageImg struct {
	src         string
	children    []any
	imgChildren []any
}

func (ci *cardImageImg) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCardImage:
			ci.children = append(ci.children, c)
		case []any:
			ci.With(c...)
		default:
			ci.imgChildren = append(ci.imgChildren, c)
		}
	}

	return ci
}

func (ci *cardImageImg) Render(w io.Writer) error {
	return Elem(html.Div, Class("card-image"), ImageImg(ci.src, ci.imgChildren...), ci.children).Render(w)
}
