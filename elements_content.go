package bulma

import "github.com/maragudk/gomponents/html"

// Content creates a content element.
//
// The following modifiers change the text size in the content:
//   - Small
//   - Normal
//   - Medium
//   - Large
//
// Add one of the the following modifiers to Ol elements in order to change the
// list style: OlLowerAlpha, OlLowerRoman, OlUpperAlpha, OlUpperRoman.
func Content(children ...any) *Element {
	e := Elem(html.Div).
		With(Class("content"))

	for _, c := range children {
		switch c := c.(type) {
		case string:
			e.With(Elem(html.P).With(c))
		default:
			e.With(c)
		}
	}

	return e
}
