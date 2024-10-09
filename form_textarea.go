package bulma

import (
	"strconv"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// Textarea creates a textarea element.
//
// https://willoma.github.io/bulma-gomponents/form/textarea.html
func Textarea(children ...any) e.Element {
	t := &textarea{e.Textarea(e.Class("textarea"))}
	t.With(children...)
	return t
}

type textarea struct {
	e.Element
}

func (t *textarea) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onTextarea:
			t.Element.With(c)
		case e.Class:
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

func (t *textarea) Clone() e.Element {
	return &textarea{t.Element.Clone()}
}

// Rows changes a textarea height.
//
// https://willoma.github.io/bulma-gomponents/form/textarea.html
func Rows(rows int) gomponents.Node {
	return html.Rows(strconv.Itoa(rows))
}
