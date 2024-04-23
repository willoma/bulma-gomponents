package bulma

import (
	"io"
	"strings"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Abbr creates an abbr element, with the provided title.
func Abbr(title string, children ...any) Element {
	return Elem(html.Abbr, html.TitleAttr(title), children)
}

// AHref creates an a element, with the provided href.
func AHref(href string, children ...any) Element {
	return Elem(html.A, html.Href(href), children)
}

// DList creates a dl element, with the provided children as alternatively
// dt and dd elements.
//
// TODO work with "With"
// TODO accept classes
func DList(dtDds ...any) Element {
	var children []any

	for i := 0; i < len(dtDds)-len(dtDds)%2; i += 2 {
		switch c := dtDds[i].(type) {
		case []any:
			children = append(children, Elem(html.Dt, c...))
		default:
			children = append(children, Elem(html.Dt, c))
		}
		switch c := dtDds[i+1].(type) {
		case []any:
			children = append(children, Elem(html.Dd, c...))
		default:
			children = append(children, Elem(html.Dd, c))
		}
	}
	return Elem(html.Dl, children...)
}

// ImgSrc creates an img element, with the provided src.
func ImgSrc(src string, children ...any) Element {
	return Elem(html.Img, html.Src(src), children)
}

// OList creates an ol element, with the provided children wrapped in li
// elements.
//
// TODO work with "With"
func OList(children ...any) Element {
	wrappedChildren := make([]any, len(children))
	for i, c := range children {
		wrappedChildren[i] = Elem(html.Li, c)
	}
	return Elem(html.Ol, wrappedChildren...)
}

// On adds a "on<event>" attribute to a gomponents.Node element.
func On(event, script string) gomponents.Node {
	return &eventAttr{event: event, script: script}
}

// OnClick adds a "onclick" attribute to a gomponents.Node element.
func OnClick(script string) gomponents.Node {
	return &eventAttr{event: "click", script: script}
}

type eventAttr struct {
	event  string
	script string
}

func (a *eventAttr) Render(w io.Writer) error {
	_, err := w.Write([]byte(" on" + a.event + `="` + strings.ReplaceAll(a.script, `"`, "&#34;") + `"`))
	return err
}

func (a *eventAttr) Type() gomponents.NodeType {
	return gomponents.AttributeType
}

// UList creates an ul element, with the provided children wrapped in li
// elements.
//
// TODO work with "With"
func UList(children ...any) Element {
	e := Elem(html.Ul)
	for _, c := range children {
		switch c := c.(type) {
		case Class, Classer, Classeser, ExternalClassesAndStyles, MultiClass, Styles:
			e.With(c)
		case gomponents.Node:
			if isAttribute(c) {
				e.With(c)
			} else {
				e.With(Elem(html.Li, c))
			}
		default:
			e.With(Elem(html.Li, c))
		}
	}
	return e
}

type ID string
