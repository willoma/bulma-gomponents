package bulma

import (
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
func Tag(children ...any) *Element {
	return Elem(html.Span).
		With(Class("tag")).
		Withs(children)
}

// Tags create a list of tags element.
//
// The following modifiers change the tags list behaviour:
//   - Addons: attach the contained tags together
func Tags(children ...any) *Element {
	e := Elem(html.Div).
		With(Class("tags"))
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			e.With(changeSizePrefix("are-", c))
		default:
			e.With(c)
		}
	}
	return e
}

// DeleteTag creates a tag which is a delete button-looking a element.
func DeleteTag(children ...any) *Element {
	return Elem(html.A).
		With(Class("tag")).
		With(Class("is-delete")).
		Withs(children)
}
