package fa

import (
	"fmt"

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

// FA returns a Font-Awesome icon, in an i element, with the provided style and
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
func FA(style Style, name string, children ...any) *b.Element {
	f := &fa{style: style, name: name}
	f.addChildren(children)
	return f.elem()
}

type fa struct {
	style       Style
	name        string
	rotateClass Class
	rotateAngle Rotate
	children    []any
}

func (f *fa) addChildren(children []any) {
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
			f.addChildren(c)
		default:
			f.children = append(f.children, c)
		}
	}
}

func (f *fa) elem() *b.Element {
	e := b.Elem(html.I).With(b.Class(f.style)).With(b.Class("fa-" + f.name))

	switch {
	case f.rotateClass != "" && f.rotateAngle != 0:
		e.With(b.Class("fa-rotate-by"))
		e.With(b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", f.rotateAngle)))
		return el.Span(
			b.Class(f.rotateClass),
			b.Style("display", "inline-block"),
			e,
		)
	case f.rotateAngle != 0:
		e.With(b.Class("fa-rotate-by"))
		e.With(b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", f.rotateAngle)))
	case f.rotateClass != "":
		e.With(b.Class(f.rotateClass))
	}

	return e
}

// Icon returns a Font-Awesome icon, in an i element within a b.Icon element,
// with the provided style and name (without the "fa-" prefix).
//   - when a child id a Class, it is added to the i element
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the b.Icon
//   - all other children types are added as-is to the b.Icon
//
// See the FA documentation for more information.
func Icon(style Style, name string, children ...any) *b.Element {
	i := &icon{style: style, name: name}
	i.addChildren(children)
	return i.elem()
}

type icon struct {
	style        Style
	name         string
	iconChildren []any
	faChildren   []any
}

func (i *icon) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			i.faChildren = append(i.faChildren, c)
		case b.ColorClass:
			i.iconChildren = append(i.iconChildren, c.Text())
		case []any:
			i.addChildren(c)
		default:
			i.iconChildren = append(i.iconChildren, c)
		}
	}
}

func (i *icon) elem() *b.Element {
	return b.Icon(i.iconChildren...).With(FA(i.style, i.name, i.faChildren...))
}
