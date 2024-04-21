package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Delete creates a delete button.
//
// The following modifiers change the button size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func Delete(children ...any) Element {
	return &delete{Elem(html.Button, Class("delete"), children)}
}

type delete struct {
	Element
}
