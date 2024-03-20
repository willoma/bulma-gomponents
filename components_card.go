package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Card creates a card.
func Card(children ...any) *Element {
	return Elem(html.Div, Class("card"), children)
}

// CardContent creates a card content.
func CardContent(children ...any) *Element {
	return Elem(html.Div, Class("card-content"), children)
}

// CardFooter creates a card footer.
//
// It add the "card-footer-item" class to all its *Element children.
//
// Values in a child of type []any are added as if they were direct children.
//
// Children may be provided in []any arguments, recursively if needed.
func CardFooter(children ...any) *Element {
	cf := &cardFooter{}
	cf.addChildren(children)
	return cf.elem()
}

type cardFooter struct {
	children []any
}

func (cf *cardFooter) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			cf.children = append(cf.children, c.With(Class("card-footer-item")))
		case []any:
			cf.addChildren(c)
		default:
			cf.children = append(cf.children, c)
		}
	}
}

func (cf *cardFooter) elem() *Element {
	return Elem(html.Footer, Class("card-footer"), cf.children)
}

// CardHeader creates a card header.
//
// If a child is a string, it is wrapped into a CardHeaderTitle element.
//
// If a child is an *Element with class "icon", it is wrapped into a
// CardHeaderIcon element.
func CardHeader(children ...any) *Element {
	ch := &cardHeader{}
	ch.addChildren(children)
	return ch.elem()
}

type cardHeader struct {
	children []any
}

func (ch *cardHeader) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("icon") {
				ch.children = append(
					ch.children,
					Elem(
						html.Button,
						Class("card-header-icon"),
						c,
					),
				)
			} else {
				ch.children = append(ch.children, c)
			}
		case string:
			ch.children = append(
				ch.children,
				Elem(
					html.P,
					Class("card-header-title"),
					c,
				),
			)
		case []any:
			ch.addChildren(c)
		default:
			ch.children = append(ch.children, c)
		}
	}
}

func (ch *cardHeader) elem() *Element {
	return Elem(html.Header, Class("card-header"), ch.children)
}

// CardHeaderIcon creates an icon for a card header.
//
// If you provide an *Element with class "icon" to the CardHeader function (ie.
// / the result of the Icon function), you don't need CardHeaderIcon.
func CardHeaderIcon(children ...any) *Element {
	return Elem(html.Button, Class("card-header-icon"), children)
}

// CardHeaderTitle creates a title for a card header.
//
// If you provide a string to the CardHeader function, you don't need
// CardHeaderTitle.
func CardHeaderTitle(children ...any) *Element {
	return Elem(html.P, Class("card-header-title"), children)
}

// CardImage creates a card image.
//
// See the Image function documentation for the list of allowed modifiers.
func CardImage(children ...any) *Element {
	return Elem(html.Div, Class("card-image"), Image(children...))
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// See the ImageImg function documentation for the list of allowed modifiers.
func CardImageImg(src string, children ...any) *Element {
	return Elem(html.Div, Class("card-image"), ImageImg(src, children...))
}
