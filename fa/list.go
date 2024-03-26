package fa

import (
	"io"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

// UList replaces bullets with icons in ul lists.
func UList(children ...any) b.Element {
	return el.Ul(b.Class("fa-ul"), children)
}

// OList replaces numbers with icons in ol lists.
func OList(children ...any) b.Element {
	return el.Ol(b.Class("fa-ul"), children)
}

// Li returns a list element with custom bullet in a UList or a OList,
// with the provided style and name (without the "fa-" prefix).
//   - when a child is a Class, it is added to the icon
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the icon (in order to apply a text color to the entire line, you must
//     provide the Text() variant directly)
//   - when a child is a rotation or animation, it is applied to the icon
//   - all other children types are added to the li
func Li(style Style, name string, children ...any) b.Element {
	return (&li{style: style, name: name}).With(children...)
}

type li struct {
	style        Style
	name         string
	spanChildren []any
	faChildren   []any
	liChildren   []any
}

func (l *li) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class, rotateOrFlip, Rotate, Animation:
			l.faChildren = append(l.faChildren, c)
		case b.ColorClass:
			l.spanChildren = append(l.spanChildren, c.Text())
		case []any:
			l.With(c...)
		default:
			l.liChildren = append(l.liChildren, c)
		}
	}
	return l
}

func (l *li) Render(w io.Writer) error {
	return el.Li(
		el.Span(
			b.Class("fa-li"),
			l.spanChildren,
			FA(l.style, l.name, l.faChildren...),
		),
		l.liChildren,
	).Render(w)
}
