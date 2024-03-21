package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Tag creates a tag element.
//
// The following modifiers change the tag color:
//   - Black
//   - Dark
//   - Light
//   - White
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
//
// The following modifiers change the tag size:
//   - Normal
//   - Medium
//   - Large
func Tag(children ...any) Element {
	return Elem(html.Span, Class("tag"), children)
}

// Tags create a list of tags element.
//
// The following modifiers change the tags list behaviour:
//   - Addons: attach the contained tags together
func Tags(children ...any) Element {
	return new(tags).With(children...)
}

type tags struct {
	children []any
}

func (t *tags) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			t.children = append(t.children, changeSizePrefix("are-", c))
		case []any:
			t.With(c...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}

func (t *tags) Render(w io.Writer) error {
	return Elem(html.Div, Class("tags"), t.children).Render(w)
}

// DeleteTag creates a tag which is a delete button-looking a element.
func DeleteTag(children ...any) Element {
	return Elem(html.A, Class("tag"), Class("is-delete"), children)
}
