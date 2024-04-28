package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Icon creates an icon span.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func Icon(children ...any) Element {
	i := &icon{icon: Elem(html.Span), iconClass: "icon"}
	i.With(children...)
	return i
}

type icon struct {
	icon Element

	iconClass Class
	rendered  sync.Once
}

func (i *icon) SetIconClass(c Class) {
	i.iconClass = c
}

func (i *icon) With(children ...any) Element {
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
	i.rendered.Do(func() {
		i.With(i.iconClass)
	})

	return i.icon.Render(w)
}

func (i *icon) Clone() Element {
	return &icon{
		icon:      i.icon.Clone(),
		iconClass: i.iconClass,
	}
}

type IconElem interface {
	Element
	SetIconClass(Class)
}

// IconText creates an icon-text span and embed all its non-icons children into
// spans.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func IconText(children ...any) Element {
	return newIconText(html.Span, children...)
}

// FlexIconText creates a flex icon-text span and embed all its non-icons children
// into spans.
//
// https://willoma.github.io/bulma-gomponents/icon.html
func FlexIconText(children ...any) Element {
	return newIconText(html.Div, children...)
}

func newIconText(fn func(...gomponents.Node) gomponents.Node, children ...any) Element {
	i := &iconText{Elem(fn, Class("icon-text"), elemOptionSpanAroundNonIconsAlways)}
	i.With(children...)
	return i
}

type iconText struct {
	Element
}

func (i *iconText) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Color:
			i.Element.With(c.Text())
		case []any:
			i.With(c...)
		default:
			i.Element.With(c)
		}
	}

	return i
}

func (i *iconText) Clone() Element {
	return &iconText{i.Element.Clone()}
}
