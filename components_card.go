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
	header   Element
	children []any
	footer   Element
}

func (c *card) With(children ...any) Element {
	for _, ch := range children {
		switch ch := ch.(type) {
		case *cardHeader:
			c.header = ch
		case *cardFooter:
			c.footer = ch
		default:
			c.children = append(c.children, ch)
		}
	}
	return c
}

func (c *card) Render(w io.Writer) error {
	e := Elem(html.Div, Class("card"))
	if c.header != nil {
		e.With(c.header)
	}
	var currentContent Element
	for _, ch := range c.children {
		switch ch := ch.(type) {
		case Class, ColorClass, ExternalClass, ExternalClassesAndStyles, MultiClass, Styles:
			e.With(ch)
		case *cardImage, *cardImageImg:
			if currentContent != nil {
				e.With(currentContent)
				currentContent = nil
			}
			e.With(ch)
		case Element:
			if currentContent == nil {
				currentContent = Elem(html.Div, Class("card-content"))
			}
			currentContent.With(ch)
		case gomponents.Node:
			if IsAttribute(c) {
				e.With(ch)
			} else {
				if currentContent == nil {
					currentContent = Elem(html.Div, Class("card-content"))
				}
				currentContent.With(ch)
			}
		default:
			if currentContent == nil {
				currentContent = Elem(html.Div, Class("card-content"))
			}
			currentContent.With(ch)
		}
	}
	if currentContent != nil {
		e.With(currentContent)
	}
	if c.footer != nil {
		e.With(c.footer)
	}
	return e.Render(w)
}

// CardFooter creates a card footer.
//
// It add the "card-footer-item" class to all its Element children.
func CardFooter(children ...any) Element {
	return new(cardFooter).With(children...)
}

type cardFooter struct {
	children []any
}

func (cf *cardFooter) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			cf.children = append(cf.children, c.With(Class("card-footer-item")))
		case []any:
			cf.With(c...)
		default:
			cf.children = append(cf.children, c)
		}
	}

	return cf
}

func (cf *cardFooter) Render(w io.Writer) error {
	return Elem(html.Footer, Class("card-footer"), cf.children).Render(w)
}

// CardHeader creates a card header.
//
// If a child is a string, it is wrapped into a CardHeaderTitle element.
//
// If a child is an Element with class "icon", it is wrapped into a
// CardHeaderIcon element.
func CardHeader(children ...any) Element {
	return new(cardHeader).With(children...)
}

type cardHeader struct {
	children []any
}

func (ch *cardHeader) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			ch.children = append(ch.children, CardHeaderIcon(c))
		case string:
			ch.children = append(ch.children, CardHeaderTitle(c))
		case []any:
			ch.With(c...)
		default:
			ch.children = append(ch.children, c)
		}
	}

	return ch
}

func (ch *cardHeader) Render(w io.Writer) error {
	return Elem(html.Header, Class("card-header"), ch.children).Render(w)
}

// CardHeaderIcon creates an icon for a card header.
//
// If you provide an icon to the CardHeader function (for instance the result of the Icon function), you don't need CardHeaderIcon.
func CardHeaderIcon(children ...any) Element {
	return Elem(html.Button, Class("card-header-icon"), children)
}

// CardHeaderTitle creates a title for a card header.
//
// The following modifiers change the title behaviour:
//   - Centered: center the title
//
// If you provide a string to the CardHeader function, you don't need
// CardHeaderTitle.
func CardHeaderTitle(children ...any) Element {
	return Elem(html.P, Class("card-header-title"), children)
}

// CardImage creates a card image.
func CardImage(children ...any) Element {
	return new(cardImage).With(children...)
}

type cardImage struct {
	children []any
}

func (ci *cardImage) With(children ...any) Element {
	ci.children = append(ci.children, children...)
	return ci
}

func (ci *cardImage) Render(w io.Writer) error {
	return Elem(html.Div, Class("card-image"), Image(ci.children...)).Render(w)
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// See the ImageImg function documentation for the list of allowed modifiers.
func CardImageImg(src string, children ...any) Element {
	return (&cardImageImg{src: src}).With(children...)
}

type cardImageImg struct {
	src      string
	children []any
}

func (ci *cardImageImg) With(children ...any) Element {
	ci.children = append(ci.children, children...)
	return ci
}

func (ci *cardImageImg) Render(w io.Writer) error {
	return Elem(html.Div, Class("card-image"), ImageImg(ci.src, ci.children...)).Render(w)
}
