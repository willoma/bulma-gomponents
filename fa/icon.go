package fa

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

type (
	Style string
	Class string
)

func (c Class) Class() b.Class {
	return b.Class(c)
}

// Styles
const (
	Brand        = Style("fa-brands")
	Duotone      = Style("fa-duotone")
	Light        = Style("fa-light")
	SharpLight   = Style("fa-sharp fa-light")
	Regular      = Style("fa-regular")
	SharpRegular = Style("fa-sharp fa-regular")
	Solid        = Style("fa-solid")
	SharpSolid   = Style("fa-sharp fa-solid")
	Thin         = Style("fa-thin")
	SharpThin    = Style("fa-sharp fa-thin")
)

// Variations
const (
	FixedWidth = Class("fa-fw")
)

// FA returns a Font Awesome icon, in an i element, with the provided style and
// name (without the "fa-" prefix).
//
// // https://willoma.github.io/bulma-gomponents/icon.html#font-awesome
func FA(style Style, name string, children ...any) b.Element {
	f := &fa{Element: b.Elem(html.I, b.Class(style), b.Class("fa-"+name))}
	f.With(children...)
	return f
}

type fa struct {
	b.Element

	rotateOrFlips []any

	rendered sync.Once
}

func (f *fa) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			f.Element.With(c)
		case rotateOrFlip, Rotate, Animation:
			f.rotateOrFlips = append(f.rotateOrFlips, c)
		case b.Color:
			f.Element.With(c.Text())
		case []any:
			f.With(c...)
		default:
			f.Element.With(c)
		}
	}

	return f
}

func (f *fa) Render(w io.Writer) error {
	f.rendered.Do(func() {
		if len(f.rotateOrFlips) > 0 {
			f.Element.With(f.rotateOrFlips[len(f.rotateOrFlips)-1])
		}
	})

	elem := f.Element

	for i := len(f.rotateOrFlips) - 2; i >= 0; i-- {
		elem = el.Span(b.Style("display", "inline-block"), f.rotateOrFlips[i], elem)
	}

	return elem.Render(w)
}

// Icon returns a Font Awesome icon, in an i element within a b.Icon element,
// with the provided style and name (without the "fa-" prefix).
//
// // https://willoma.github.io/bulma-gomponents/icon.html#font-awesome
func Icon(style Style, name string, children ...any) b.Element {
	fa := FA(style, name)
	i := &icon{
		Element: b.Icon(fa),
		fa:      fa,
	}
	i.With(children...)
	return i
}

type icon struct {
	b.Element
	fa b.Element
}

func (i *icon) SetIconClass(c b.Class) {
	i.Element.(b.IconElem).SetIconClass(c)
}

func (i *icon) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onFA:
			i.fa.With(c...)
		case onIcon:
			i.Element.With(c...)
		case Class, rotateOrFlip, Rotate, Animation:
			i.fa.With(c)
		case b.Color:
			i.Element.With(c.Text())
		case []any:
			i.With(c...)
		default:
			i.Element.With(c)
		}
	}

	return i
}

func OnFA(children ...any) onFA {
	return onFA(children)
}

type onFA []any

func OnIcon(children ...any) onIcon {
	return onIcon(children)
}

type onIcon []any
