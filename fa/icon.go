package fa

import (
	"fmt"

	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

type (
	faStyle string
	Class   string
)

// Styles
const (
	Brand        = faStyle("fa-brands")
	Duotone      = faStyle("fa-duotone")
	Light        = faStyle("fa-light")
	SharpLight   = faStyle("fa-sharp fa-light")
	Regular      = faStyle("fa-regular")
	SharpRegular = faStyle("fa-sharp fa-regular")
	Solid        = faStyle("fa-solid")
	SharpSolid   = faStyle("fa-sharp fa-solid")
	Thin         = faStyle("fa-thin")
	SharpThin    = faStyle("fa-sharp fa-thin")
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
//   - when a child id a FaClass, it is added as a class to the i element
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the i element
//   - when a child is a Rotate, the provided rotation (in degree) is applied to
//     the icon
//   - all other children types are added as-is to the i element
//
// The rotating+flipping combination is supported and the needed span element
// is automatically created when needed.
func FA(style faStyle, name string, children ...any) *b.Element {
	var rotateClass Class
	var rotateAngle Rotate

	e := b.Elem(html.I).With(b.Class(style)).With(b.Class("fa-" + name))
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			if c == Rotate90 || c == Rotate180 || c == Rotate270 || c == FlipHorizontal || c == FlipVertical || c == FlipBoth {
				rotateClass = c
			} else {
				e.With(b.Class(c))
			}
		case b.ColorClass:
			e.With(c.Text())
		case Rotate:
			rotateAngle = c
		case Animation:
			class, styles := c.attrs()
			e.With(class)
			if len(styles) > 0 {
				e.With(styles)
			}
		default:
			e.With(c)
		}
	}

	switch {
	case rotateClass != "" && rotateAngle != 0:
		e.With(b.Class("fa-rotate-by"))
		e.With(b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", rotateAngle)))
		return el.Span(
			b.Class(rotateClass),
			b.Style("display", "inline-block"),
			e,
		)
	case rotateAngle != 0:
		e.With(b.Class("fa-rotate-by"))
		e.With(b.Style("--fa-rotate-angle", fmt.Sprintf("%vdeg", rotateAngle)))
	case rotateClass != "":
		e.With(b.Class(rotateClass))
	}
	return e
}

// Icon returns a Font-Awesome icon, in an i element within a b.Icon element,
// with the provided style and name (without the "fa-" prefix).
//   - when a child id a FaClass, it is added to the i element
//   - when a child is a b.ColorClass, its Text() variant is added as a class to
//     the b.Icon
//   - all other children types are added as-is to the b.Icon
//
// See the FA documentation for more information.
func Icon(style faStyle, name string, children ...any) *b.Element {
	ic := b.Icon()

	var faChildren []any

	for _, c := range children {
		switch c := c.(type) {
		case Class:
			faChildren = append(faChildren, c)
		case b.ColorClass:
			ic.With(c.Text())
		default:
			ic.With(c)
		}
	}
	return ic.With(FA(style, name, faChildren...))
}
