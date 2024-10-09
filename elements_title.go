package bulma

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"
)

// Title creates a h1 title.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title(children ...any) e.Element {
	t := &title{e.H1(e.Class("title"))}

	t.With(children...)

	return t
}

// Subtitle creates a h2 subtitle.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle(children ...any) e.Element {
	t := &title{e.H2(e.Class("subtitle"))}

	t.With(children...)

	return t
}

type title struct {
	e.Element
}

func (t *title) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case int:
			switch c {
			case 1:
				t.Element.With(html.H1, e.Class("is-1"))
			case 2:
				t.Element.With(html.H2, e.Class("is-2"))
			case 3:
				t.Element.With(html.H3, e.Class("is-3"))
			case 4:
				t.Element.With(html.H4, e.Class("is-4"))
			case 5:
				t.Element.With(html.H5, e.Class("is-5"))
			case 6:
				t.Element.With(html.H6, e.Class("is-6"))
			}
		case []any:
			t.With(c...)
		default:
			t.Element.With(c)
		}
	}
	return t
}

func (t *title) Clone() e.Element {
	return &title{t.Element.Clone()}
}
