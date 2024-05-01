package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

// Card creates a card.
//
// https://willoma.github.io/bulma-gomponents/card.html
func Card(children ...any) e.Element {
	c := &card{card: e.Div(e.Class("card"))}
	c.With(children...)
	return c
}

type card struct {
	card    e.Element
	header  e.Element
	content []any
	footer  e.Element

	currentContent e.Element
	rendered       sync.Once
}

func (c *card) addToCurrentContent(children ...any) {
	if c.currentContent == nil {
		c.currentContent = e.Div(e.Class("card-content"))
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
		c.header = e.Header(e.Class("card-header"))
	}
	c.header.With(children...)
}

func (c *card) addToFooter(children ...any) {
	if c.footer == nil {
		c.footer = e.Footer(e.Class("card-footer"))
	}
	c.footer.With(children...)
}

func (c *card) With(children ...any) e.Element {
	var currentContent []any

	for _, ch := range children {
		switch ch := ch.(type) {
		case onCard:
			c.card.With(ch...)
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
		case e.Class, e.Classer, e.Classeser, e.Styles:
			c.card.With(ch)
		case *cardImage, *cardImageImg:
			c.flushCurrentContent()
			c.content = append(c.content, ch)
		case e.Element:
			c.addToCurrentContent(ch)
		case gomponents.Node:
			if e.IsAttribute(c) {
				c.card.With(ch)
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
			c.card.With(c.header)
		}

		if c.content != nil {
			c.card.With(c.content...)
		}

		if c.footer != nil {
			c.card.With(c.footer)
		}
	})

	return c.card.Render(w)
}

func (c *card) Clone() e.Element {
	content := make([]any, len(c.content))
	for i, c := range c.content {
		switch c := c.(type) {
		case e.Element:
			content[i] = c.Clone()
		default:
			content[i] = c
		}
	}

	return &card{
		card:    c.card.Clone(),
		header:  c.header.Clone(),
		content: content,
		footer:  c.footer.Clone(),

		currentContent: c.currentContent.Clone(),
	}
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
func CardHeaderIcon(children ...any) e.Element {
	c := &cardHeaderIcon{e.Button(e.Class("card-header-icon"))}
	c.With(children...)
	return c
}

type cardHeaderIcon struct {
	e.Element
}

func (c *cardHeaderIcon) Clone() e.Element {
	return &cardHeaderIcon{c.Element.Clone()}
}

// CardHeaderTitle creates a title for a card header.
//
// https://willoma.github.io/bulma-gomponents/card.html
func CardHeaderTitle(children ...any) e.Element {
	c := &cardHeaderTitle{e.P(e.Class("card-header-title"))}
	c.With(children...)
	return c
}

type cardHeaderTitle struct {
	e.Element
}

func (c *cardHeaderTitle) Clone() e.Element {
	return &cardHeaderTitle{c.Element.Clone()}
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
		case e.Element:
			cf = append(cf, c.With(e.Class("card-footer-item")))
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
func CardImage(children ...any) e.Element {
	image := Image()
	c := &cardImage{
		Element: e.Div(e.Class("card-image"), image),
		image:   image,
	}
	c.With(children...)
	return c
}

type cardImage struct {
	e.Element
	image e.Element
}

func (ci *cardImage) With(children ...any) e.Element {
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

func (ci *cardImage) Clone() e.Element {
	return &cardImage{
		Element: ci.Element.Clone(),
		image:   ci.image.Clone(),
	}
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// https://willoma.github.io/bulma-gomponents/card.html
func CardImageImg(src string, children ...any) e.Element {
	imageImg := ImageImg(src)
	c := &cardImageImg{
		Element:  e.Div(e.Class("card-image"), imageImg),
		imageImg: imageImg,
	}
	c.With(children...)
	return c
}

type cardImageImg struct {
	e.Element
	imageImg e.Element
}

func (ci *cardImageImg) With(children ...any) e.Element {
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

func (ci *cardImageImg) Clone() e.Element {
	return &cardImageImg{
		Element:  ci.Element.Clone(),
		imageImg: ci.imageImg.Clone(),
	}
}
