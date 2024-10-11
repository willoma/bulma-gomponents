package bulma

import (
	"io"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// Icon creates an icon span.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func Icon(children ...any) e.Element {
	i := &icon{icon: e.Span(), iconClass: "icon"}
	i.With(children...)
	return i
}

type icon struct {
	icon e.Element

	iconClass e.Class
}

func (i *icon) SetIconClass(c e.Class) {
	i.iconClass = c
}

func (i *icon) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Color:
			i.icon.With(c.Text())
		case []any:
			i.With(c...)
		default:
			i.icon.With(c)
		}
	}

	return i
}

func (i *icon) Render(w io.Writer) error {
	return i.icon.Clone().With(i.iconClass).Render(w)
}

func (i *icon) Clone() e.Element {
	return &icon{
		icon:      i.icon.Clone(),
		iconClass: i.iconClass,
	}
}

type IconElem interface {
	e.Element
	SetIconClass(e.Class)
}

// IconText creates an icon-text span and embed all its non-icons children into
// spans.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func IconText(children ...any) e.Element {
	return newIconText(html.Span, children...)
}

// FlexIconText creates a flex icon-text span and embed all its non-icons children
// into spans.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func FlexIconText(children ...any) e.Element {
	return newIconText(html.Div, children...)
}

func newIconText(fn func(...gomponents.Node) gomponents.Node, children ...any) e.Element {
	i := &spanAroundNonIcons{e.Elem(fn, e.Class("icon-text"))}
	i.With(children...)
	return i
}
