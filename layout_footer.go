package bulma

import "github.com/maragudk/gomponents/html"

// Footer creates a footer element.
//
// http://willoma.github.io/bulma-gomponents/footer.html
func Footer(children ...any) Element {
	return Elem(html.Footer, Class("footer"), children)
}
