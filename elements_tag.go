package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Tag creates a tag element.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func Tag(children ...any) Element {
	return Elem(html.Span, Class("tag"), children)
}

// Tags create a list of tags element.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func Tags(children ...any) Element {
	t := &tags{Elem(html.Div, Class("tags"))}
	t.With(children...)
	return t
}

type tags struct {
	Element
}

func (t *tags) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			t.Element.With(changeSizePrefix("are-", c))
		case []any:
			t.With(c...)
		default:
			t.Element.With(c)
		}
	}

	return t
}

// DeleteTag creates a delete button-looking tag.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func DeleteTag(children ...any) Element {
	return Elem(html.A, Class("tag"), Class("is-delete"), children)
}
