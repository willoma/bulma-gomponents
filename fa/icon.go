package fa

import (
	"io"

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
//   - when a child id a Class, it is added as a class to the i element
//   - when a child is a color, its text variant is applied to the i element
//   - when a child is a rotation or animation, it is applied to the icon
//   - all other children types are added as-is to the i element
//
// The rotating+flipping combination is supported and the needed span element
// is automatically created when needed.
func FA(style Style, name string, children ...any) b.Element {
	return (&fa{style: style, name: name}).With(children...)
}

type fa struct {
	style         Style
	name          string
	rotateOrFlips []any
	children      []any
}

func (f *fa) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			f.children = append(f.children, c)
		case rotateOrFlip, Rotate, Animation:
			f.rotateOrFlips = append(f.rotateOrFlips, c)
		case b.Color:
			f.children = append(f.children, c.Text())
		case []any:
			f.With(c...)
		default:
			f.children = append(f.children, c)
		}
	}

	return f
}

func (f *fa) Render(w io.Writer) error {
	if len(f.rotateOrFlips) == 0 {
		return b.Elem(html.I, b.Class(f.style), b.Class("fa-"+f.name)).With(f.children...).Render(w)
	}

	e := b.Elem(html.I, b.Class(f.style), b.Class("fa-"+f.name), f.rotateOrFlips[len(f.rotateOrFlips)-1]).With(f.children...)

	for i := len(f.rotateOrFlips) - 2; i >= 0; i-- {
		e = el.Span(b.Style("display", "inline-block"), f.rotateOrFlips[i], e)
	}

	return e.Render(w)
}

// Icon returns a Font Awesome icon, in an i element within a b.Icon element,
// with the provided style and name (without the "fa-" prefix).
//   - when a child is marked with b.Inner, it is forcibly applied to the <i> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <span> element
//   - when a child id a Class, it is added to the i element
//   - when a child is a color, its text variant is applied to the b.Icon
//   - when a child is a rotation or animation, it is applied to the icon
//   - all other children types are added as-is to the b.Icon
func Icon(style Style, name string, children ...any) b.Element {
	return (&icon{style: style, name: name}).With(children...)
}

type icon struct {
	iconClass    b.Class
	style        Style
	name         string
	iconChildren []any
	faChildren   []any
}

func (i *icon) SetIconClass(c b.Class) {
	i.iconClass = c
}

func (i *icon) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onFA:
			i.faChildren = append(i.faChildren, c...)
		case onSpan:
			i.iconChildren = append(i.iconChildren, c...)
		case Class, rotateOrFlip, Rotate, Animation:
			i.faChildren = append(i.faChildren, c)
		case b.Color:
			i.iconChildren = append(i.iconChildren, c.Text())
		case []any:
			i.With(c...)
		default:
			i.iconChildren = append(i.iconChildren, c)
		}
	}

	return i
}

func (i *icon) Render(w io.Writer) error {
	ic := b.Icon(i.iconChildren, FA(i.style, i.name, i.faChildren...))
	if i.iconClass != "" {
		ic.(b.IconElem).SetIconClass(i.iconClass)
	}
	return ic.Render(w)
}

func OnFA(children ...any) onFA {
	return onFA(children)
}

type onFA []any

func OnSpan(children ...any) onSpan {
	return onSpan(children)
}

type onSpan []any
