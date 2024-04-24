package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Card creates a card.
//
// https://willoma.github.io/bulma-gomponents/card.html
func Card(children ...any) Element {
	c := &card{Element: Elem(html.Div, Class("card"))}
	c.With(children...)
	return c
}

type card struct {
	Element
	header  Element
	content []any
	footer  Element

	currentContent Element
	rendered       sync.Once
}

func (c *card) addToCurrentContent(children ...any) {
	if c.currentContent == nil {
		c.currentContent = Elem(html.Div, Class("card-content"))
	}
	c.currentContent.With(children...)
}

func (c *card) flushCurrentContent() {
	if c.currentContent != nil {
		c.content = append(c.content, c.currentContent)
		c.currentContent = nil
	}
}

func (c *card) addToHeader(children ...any) {
	if c.header == nil {
		c.header = Elem(html.Header, Class("card-header"))
	}
	c.header.With(children...)
}

func (c *card) addToFooter(children ...any) {
	if c.footer == nil {
		c.footer = Elem(html.Footer, Class("card-footer"))
	}
	c.footer.With(children...)
}

func (c *card) With(children ...any) Element {
	var currentContent []any

	for _, ch := range children {
		switch ch := ch.(type) {
		case onCard:
			c.Element.With(ch...)
		case onContent:
			c.addToCurrentContent(ch...)
		case splitContent:
			c.flushCurrentContent()
		case cardHeader:
			c.addToHeader(ch...)
		case *cardHeaderIcon:
			c.addToHeader(ch)
		case *cardHeaderTitle:
			c.addToHeader(ch)
		case cardFooter:
			c.addToFooter(ch...)
		case Class, Classer, Classeser, ExternalClassesAndStyles, Styles:
			c.Element.With(ch)
		case *cardImage, *cardImageImg:
			c.flushCurrentContent()
			c.content = append(c.content, ch)
		case Element:
			c.addToCurrentContent(ch)
		case gomponents.Node:
			if isAttribute(c) {
				c.Element.With(ch)
			} else {
				c.addToCurrentContent(ch)
			}
		case []any:
			currentContent = append(currentContent, ch...)

		default:
			c.addToCurrentContent(ch)
		}
	}

	return c
}

func (c *card) Render(w io.Writer) error {
	c.rendered.Do(func() {
		c.flushCurrentContent()

		if c.header != nil {
			c.Element.With(c.header)
		}

		if c.content != nil {
			c.Element.With(c.content...)
		}

		if c.footer != nil {
			c.Element.With(c.footer)
		}
	})

	return c.Element.Render(w)
}

// SplitContent instructs Card to close the current card-content element.
func SplitContent() splitContent {
	return splitContent{}
}

type splitContent struct{}

// CardHeader marks its children as being part of a card header.
//
// https://willoma.github.io/bulma-gomponents/card.html
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
// https://willoma.github.io/bulma-gomponents/card.html
func CardHeaderIcon(children ...any) Element {
	c := &cardHeaderIcon{Elem(html.Button, Class("card-header-icon"))}
	c.With(children...)
	return c
}

type cardHeaderIcon struct {
	Element
}

// CardHeaderTitle creates a title for a card header.
//
// https://willoma.github.io/bulma-gomponents/card.html
func CardHeaderTitle(children ...any) Element {
	c := &cardHeaderTitle{Elem(html.P, Class("card-header-title"))}
	c.With(children...)
	return c
}

type cardHeaderTitle struct {
	Element
}

// CardFooter marks its children as being part of a card footer.
//
// https://willoma.github.io/bulma-gomponents/card.html
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
//
// https://willoma.github.io/bulma-gomponents/card.html
func CardImage(children ...any) Element {
	image := Image()
	c := &cardImage{
		Element: Elem(html.Div, Class("card-image"), image),
		image:   image,
	}
	c.With(children...)
	return c
}

type cardImage struct {
	Element
	image Element
}

func (ci *cardImage) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCardImage:
			ci.Element.With(c...)
		case []any:
			ci.With(c...)
		default:
			ci.image.With(c)
		}
	}

	return ci
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// https://willoma.github.io/bulma-gomponents/card.html
func CardImageImg(src string, children ...any) Element {
	imageImg := ImageImg(src)
	c := &cardImageImg{
		Element:  Elem(html.Div, Class("card-image"), imageImg),
		imageImg: imageImg,
	}
	c.With(children...)
	return c
}

type cardImageImg struct {
	Element
	imageImg Element
}

func (ci *cardImageImg) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onCardImage:
			ci.Element.With(c...)
		case []any:
			ci.With(c...)
		default:
			ci.imageImg.With(c)
		}
	}

	return ci
}
