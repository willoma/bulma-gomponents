package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Hero creates a hero element.
//   - when a child is a return value of HeroHead, it is added in the head part
//     of the hero element
//   - when a child is a return value of HeroFoot, it is added in the foot part
//     of the hero element
//   - when a child is an Element, it is added in the body part of the hero
//     element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as a direct child of to the hero element
//   - when a child is a gomponents.Node with another type, it is added in the
//     body part of the hero element
//   - other children types are added as direct children of to the hero element
//
// Each of the head, body and foot parts is only included if it has content.
//
// The following modifiers change the hero color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the hero size:
//   - Small
//   - Medium
//   - Large
//   - HalfHeight
//   - FullHeight
//   - FullHeightWithNavbar
func Hero(children ...any) Element {
	return new(hero).With(children...)
}

type hero struct {
	elemChildren []any
	headChildren []any
	bodyChildren []any
	footChildren []any
}

func (h *hero) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onBody:
			h.bodyChildren = append(h.bodyChildren, c...)
		case onSection:
			h.elemChildren = append(h.elemChildren, c...)
		case heroHead:
			h.headChildren = append(h.headChildren, c...)
		case heroFoot:
			h.footChildren = append(h.footChildren, c...)
		case Element:
			h.bodyChildren = append(h.bodyChildren, c)
		case gomponents.Node:
			if IsAttribute(c) {
				h.elemChildren = append(h.elemChildren, c)
			} else {
				h.bodyChildren = append(h.bodyChildren, c)
			}
		case []any:
			h.With(c...)
		default:
			h.elemChildren = append(h.elemChildren, c)
		}
	}

	return h
}

func (h *hero) Render(w io.Writer) error {
	e := Elem(html.Section, Class("hero"), h.elemChildren)

	if len(h.headChildren) > 0 {
		e.With(
			Elem(html.Div, Class("hero-head"), h.headChildren),
		)
	}

	if len(h.bodyChildren) > 0 {
		e.With(
			Elem(html.Div, Class("hero-body"), h.bodyChildren),
		)
	}

	if len(h.footChildren) > 0 {
		e.With(
			Elem(html.Div, Class("hero-foot"), h.footChildren),
		)
	}

	return e.Render(w)
}

// HeroHead marks children as belonging to the head part of a hero element.
func HeroHead(children ...any) heroHead {
	return heroHead(children)
}

type heroHead []any

// HeroFoot marks children as belonging to the foot part of a hero element.
func HeroFoot(children ...any) heroFoot {
	return heroFoot(children)
}

type heroFoot []any
