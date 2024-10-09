package bulma

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// TitleLevel is used for thtle level, it corresponds to h1 to h6
type TitleLevel int

const (
	TitleLevel1 TitleLevel = 1
	TitleLevel2 TitleLevel = 2
	TitleLevel3 TitleLevel = 3
	TitleLevel4 TitleLevel = 4
	TitleLevel5 TitleLevel = 5
	TitleLevel6 TitleLevel = 6
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
		case TitleLevel:
			switch c {
			case TitleLevel1:
				t.Element.With(html.H1, e.Class("is-1"))
			case TitleLevel2:
				t.Element.With(html.H2, e.Class("is-2"))
			case TitleLevel3:
				t.Element.With(html.H3, e.Class("is-3"))
			case TitleLevel4:
				t.Element.With(html.H4, e.Class("is-4"))
			case TitleLevel5:
				t.Element.With(html.H5, e.Class("is-5"))
			case TitleLevel6:
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
