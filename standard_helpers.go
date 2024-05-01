package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Abbr creates an abbr element, with the provided title.
func Abbr(title string, children ...any) e.Element {
	return e.Abbr(html.TitleAttr(title), children)
}

// DList creates a dl element, with the provided children as alternatively
// dt and dd elements.
//
// TODO work with "With"
// TODO accept classes
func DList(dtDds ...any) e.Element {
	var children []any

	for i := 0; i < len(dtDds)-len(dtDds)%2; i += 2 {
		switch c := dtDds[i].(type) {
		case []any:
			children = append(children, e.Dt(c...))
		default:
			children = append(children, e.Dt(c))
		}
		switch c := dtDds[i+1].(type) {
		case []any:
			children = append(children, e.Dd(c...))
		default:
			children = append(children, e.Dd(c))
		}
	}
	return e.Dl(children...)
}

// OList creates an ol element, with the provided children wrapped in li
// elements.
//
// TODO work with "With"
func OList(children ...any) e.Element {
	wrappedChildren := make([]any, len(children))
	for i, c := range children {
		wrappedChildren[i] = e.Li(c)
	}
	return e.Ol(wrappedChildren...)
}

// UList creates an ul element, with the provided children wrapped in li
// elements.
//
// TODO work with "With"
func UList(children ...any) e.Element {
	el := e.Ul()
	for _, c := range children {
		switch c := c.(type) {
		case e.Class, e.Classer, e.Classeser, e.Styles:
			el.With(c)
		case gomponents.Node:
			if e.IsAttribute(c) {
				el.With(c)
			} else {
				el.With(e.Li(c))
			}
		default:
			el.With(e.Li(c))
		}
	}
	return el
}
