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
func Content(children ...any) Element {
	ct := &content{}
	ct.addChildren(children)
	return ct.elem()
}

type content struct {
	children []any
}

func (ct *content) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case string:
			ct.children = append(ct.children, Elem(html.P, c))
		case []any:
			ct.addChildren(c)
		default:
			ct.children = append(ct.children, c)
		}
	}
}

func (ct *content) elem() Element {
	return Elem(html.Div, Class("content"), ct.children)
}
