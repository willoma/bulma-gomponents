package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Card creates a card.
func Card(children ...any) *Element {
	return Elem(html.Div).
		With(Class("card")).
		Withs(children)
}

// CardContent creates a card content.
func CardContent(children ...any) *Element {
	return Elem(html.Div).
		With(Class("card-content")).
		Withs(children)
}

// CardFooter creates a card footer.
//
// It add the "card-footer-item" class to all its *Element children.
func CardFooter(children ...any) *Element {
	e := Elem(html.Footer).
		With(Class("card-footer"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			e.With(c.With(Class("card-footer-item")))
		default:
			e.With(c)
		}
	}

	return e
}

// CardHeader creates a card header.
//
// If a child is a string, it is wrapped into a CardHeaderTitle element.
//
// If a child is an *Element with class "icon", it is wrapped into a
// CardHeaderIcon element.
func CardHeader(children ...any) *Element {
	e := Elem(html.Header).
		With(Class("card-header"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("icon") {
				e.With(
					Elem(html.Button).
						With(Class("card-header-icon")).
						With(c),
				)
			} else {
				e.With(c)
			}
		case string:
			e.With(
				Elem(html.P).
					With(Class("card-header-title")).
					With(c),
			)
		default:
			e.With(c)
		}
	}

	return e
}

// CardHeaderIcon creates an icon for a card header.
//
// If you provide an *Element with class "icon" to the CardHeader function (ie.
// / the result of the Icon function), you don't need CardHeaderIcon.
func CardHeaderIcon(children ...any) *Element {
	return Elem(html.Button).
		With(Class("card-header-icon")).
		Withs(children)
}

// CardHeaderTitle creates a title for a card header.
//
// If you provide a string to the CardHeader function, you don't need
// CardHeaderTitle.
func CardHeaderTitle(children ...any) *Element {
	return Elem(html.P).
		With(Class("card-header-title")).
		Withs(children)
}

// CardImage creates a card image.
//
// See the Image function documentation for the list of allowed modifiers.
func CardImage(children ...any) *Element {
	return Elem(html.Div).
		With(Class("card-image")).
		With(Image(children...))
}

// CardImageImg creates a card image which contains an img element with the
// provided src.
//
// See the ImageImg function documentation for the list of allowed modifiers.
func CardImageImg(src string, children ...any) *Element {
	return Elem(html.Div).
		With(Class("card-image")).
		With(ImageImg(src, children...))
}
