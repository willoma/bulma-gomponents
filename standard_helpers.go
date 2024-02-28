package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Abbr creates an abbr element, with the provided title.
func Abbr(title string, children ...any) *Element {
	return Elem(html.Abbr).
		With(html.TitleAttr(title)).
		Withs(children)
}

// AHref creates an a element, with the provided href.
func AHref(href string, children ...any) *Element {
	return Elem(html.A).
		With(html.Href(href)).
		Withs(children)
}

// DList creates a dl element, with the provided children as alternatively
// dt and dd elements.
func DList(dtDds ...any) *Element {
	var limit int
	if len(dtDds)%2 == 0 {
		limit = len(dtDds)
	} else {
		limit = len(dtDds) - 1
	}

	e := Elem(html.Dl)

	for i := 0; i < limit; i += 2 {
		switch c := dtDds[i].(type) {
		case []any:
			e.With(Elem(html.Dt).Withs(c))
		default:
			e.With(Elem(html.Dt).With(c))
		}
		switch c := dtDds[i+1].(type) {
		case []any:
			e.With(Elem(html.Dd).Withs(c))
		default:
			e.With(Elem(html.Dd).With(c))
		}
	}
	return e
}

// ImgSrc creates an img element, with the provided src.
func ImgSrc(src string, children ...any) *Element {
	return Elem(html.Img).
		With(html.Src(src)).
		Withs(children)
}

// OList creates an ol element, with the provided children wrapped in li
// elements. A child may be of type []any in order to add multiple elements in
// a li.
func OList(children ...any) *Element {
	e := Elem(html.Ol)
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			e.With(Elem(html.Li).Withs(c))
		default:
			e.With(Elem(html.Li).With(c))
		}
	}
	return e
}

// On adds a "on<event>" attribute to a gomponents.Node element.
func On(event, script string) gomponents.Node {
	return gomponents.Attr("on"+event, script)
}

// OnClick adds a "onclick" attribute to a gomponents.Node element.
func OnClick(script string) gomponents.Node {
	return gomponents.Attr("onclick", script)
}

// UList creates an ul element, with the provided children wrapped in li
// elements. A child may be of type []any in order to add multiple elements in
// a li.
func UList(children ...any) *Element {
	e := Elem(html.Ul)
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			e.With(Elem(html.Li).Withs(c))
		default:
			e.With(Elem(html.Li).With(c))
		}
	}
	return e
}
