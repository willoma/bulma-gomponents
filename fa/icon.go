package fa

import (
	"fmt"
	"io"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

type (
	Style string
	Class string
)

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
	Border     = Class("fa-border")
	FixedWidth = Class("fa-fw")
	Inverse    = Class("fa-inverse")
	Pulse      = Class("fa-pulse")
)

// FA returns a Font Awesome icon, in an i element, with the provided style and
// name (without the "fa-" prefix).
//   - when a child id a Class, it is added as a class to the i element
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the i element
//   - when a child is a Rotate, the provided rotation (in degree) is applied to
//     the icon
//   - all other children types are added as-is to the i element
//
// The rotating+flipping combination is supported and the needed span element
// is automatically created when needed.
func FA(style Style, name string, children ...any) b.Element {
	return (&fa{style: style, name: name}).With(children...)
}

type fa struct {
	style       Style
	name        string
	rotateClass Class
	rotateAngle Rotate
	children    []any
}

func (f *fa) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			if c == Rotate90 || c == Rotate180 || c == Rotate270 || c == FlipHorizontal || c == FlipVertical || c == FlipBoth {
				f.rotateClass = c
			} else {
				f.children = append(f.children, b.Class(c))
			}
		case b.ColorClass:
			f.children = append(f.children, c.Text())
		case Rotate:
			f.rotateAngle = c
		case Animation:
			class, styles := c.attrs()
			f.children = append(f.children, class)
			if len(styles) > 0 {
				f.children = append(f.children, styles)
			}
		case []any:
			f.With(c...)
		default:
			f.children = append(f.children, c)
		}
	}

	return f
}

func (f *fa) Render(w io.Writer) error {
	e := b.Elem(html.I, b.Class(f.style), b.Class("fa-"+f.name))

	switch {
	case f.rotateClass != "" && f.rotateAngle != 0:
		e.With(
			b.Class("fa-rotate-by"),
			b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", f.rotateAngle)),
		)
		return el.Span(
			b.Class(f.rotateClass),
			b.Style("display", "inline-block"),
			e,
		).Render(w)
	case f.rotateAngle != 0:
		e.With(
			b.Class("fa-rotate-by"),
			b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", f.rotateAngle)),
		)
	case f.rotateClass != "":
		e.With(b.Class(f.rotateClass))
	}

	return e.With(f.children...).Render(w)
}

// Icon returns a Font Awesome icon, in an i element within a b.Icon element,
// with the provided style and name (without the "fa-" prefix).
//   - when a child id a Class, it is added to the i element
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the b.Icon
//   - all other children types are added as-is to the b.Icon
//
// See the FA documentation for more information.
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
		case Class:
			i.faChildren = append(i.faChildren, c)
		case b.ColorClass:
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
