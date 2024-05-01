package bulma

import (
	e "github.com/willoma/gomplements"
)

// Tag creates a tag element.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func Tag(children ...any) e.Element {
	return e.Span(e.Class("tag"), children)
}

// Tags create a list of tags element.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func Tags(children ...any) e.Element {
	t := &tags{e.Div(e.Class("tags"))}
	t.With(children...)
	return t
}

type tags struct {
	e.Element
}

func (t *tags) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class:
			t.Element.With(changeSizePrefix("are-", c))
		case []any:
			t.With(c...)
		default:
			t.Element.With(c)
		}
	}

	return t
}

func (t *tags) Clone() e.Element {
	return &tags{t.Element.Clone()}
}

// DeleteTag creates a delete button-looking tag.
//
// https://willoma.github.io/bulma-gomponents/tag.html
func DeleteTag(children ...any) e.Element {
	return e.A(e.Class("tag"), e.Class("is-delete"), children)
}
