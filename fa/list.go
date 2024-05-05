package fa

import (
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
)

// UList replaces bullets with icons in ul lists.
func UList(children ...any) e.Element {
	return e.Ul(e.Class("fa-ul"), children)
}

// OList replaces numbers with icons in ol lists.
func OList(children ...any) e.Element {
	return e.Ol(e.Class("fa-ul"), children)
}

// Li returns a list element with custom bullet in a UList or a OList,
// with the provided style and name (without the "fa-" prefix).
//   - when a child is a Class, it is added to the icon
//   - when a child is a color, its text variant is applied to the icon (in
//     order ti apply the color to the entire line, you must use the Text*
//     variant)
//   - when a child is a rotation or animation, it is applied to the icon
//   - all other children types are added to the li
func Li(style Style, name string, children ...any) e.Element {
	fa := FA(style, name)
	span := e.Span(e.Class("fa-li"), fa)
	l := &li{
		Element: e.Li(span),
		fa:      fa,
		span:    span,
	}
	l.With(children...)
	return l
}

type li struct {
	e.Element
	fa   e.Element
	span e.Element
}

func (l *li) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class, rotateOrFlip, Rotate, animationI:
			l.fa.With(c)
		case b.Color:
			l.span.With(c.Text())
		case []any:
			l.With(c...)
		default:
			l.Element.With(c)
		}
	}
	return l
}
