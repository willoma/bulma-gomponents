package bulma

import (
	e "github.com/willoma/gomplements"
)

// Footer creates a footer element.
//
// http://willoma.github.io/bulma-gomponents/footer.html
func Footer(children ...any) e.Element {
	return e.Footer(e.Class("footer"), children)
}
