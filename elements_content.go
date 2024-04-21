package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

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
	return new(content).With(children...)
}

type content struct {
	children []any
}

func (ct *content) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case noP:
			ct.withNoP(c)
		case string:
			ct.children = append(ct.children, Elem(html.P, c))
		case []any:
			ct.With(c...)
		default:
			ct.children = append(ct.children, c)
		}
	}

	return ct
}

func (ct *content) withNoP(children []any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			ct.With(c...)
		default:
			ct.children = append(ct.children, c)
		}
	}

	return ct
}

func (ct *content) Render(w io.Writer) error {
	return Elem(html.Div, Class("content"), ct.children).Render(w)
}

func NoP(children ...any) noP {
	return noP(children)
}

type noP []any
