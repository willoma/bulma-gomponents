package bulma

import (
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Textarea creates a textarea element.
//
// https://willoma.github.io/bulma-gomponents/form/textarea.html
func Textarea(children ...any) Element {
	t := &textarea{Elem(html.Textarea, Class("textarea"))}
	t.With(children...)
	return t
}

type textarea struct {
	Element
}

func (t *textarea) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			switch c {
			case Disabled:
				t.Element.With(html.Disabled())
			default:
				t.Element.With(c)
			}
		case []any:
			t.With(c...)
		default:
			t.Element.With(c)
		}
	}

	return t
}

// Rows changes a textarea height.
//
// https://willoma.github.io/bulma-gomponents/form/textarea.html
func Rows(rows int) gomponents.Node {
	return html.Rows(strconv.Itoa(rows))
}
