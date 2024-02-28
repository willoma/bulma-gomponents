package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type heroPart struct {
	foot     bool
	children []any
}

// Hero creates a hero element.
//   - when a child is a return value of HeroHead, it is added in the head part
//     of the hero element
//   - when a child is a return value of HeroFoot, it is added in the foot part
//     of the hero element
//   - when a child is an *Element, it is added in the body part of the hero
//     element
//   - when a child is a return value of Container*, it is added in the body
//     part of the hero element
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
func Hero(children ...any) *Element {
	e := Elem(html.Section).
		With(Class("hero"))

	var head, body, foot []any
	for _, c := range children {
		switch c := c.(type) {
		case heroPart:
			if c.foot {
				foot = append(foot, c.children...)
			} else {
				head = append(head, c.children...)
			}
		case *Element, container:
			body = append(body, c)
		case gomponents.Node:
			if IsAttribute(c) {
				e.With(c)
			} else {
				body = append(body, c)
			}
		default:
			e.With(c)
		}
	}

	if len(head) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("hero-head")).
				Withs(head),
		)
	}

	if len(body) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("hero-body")).
				Withs(body),
		)
	}

	if len(foot) > 0 {
		e.With(
			Elem(html.Div).
				With(Class("hero-foot")).
				Withs(foot),
		)
	}

	return e
}

// HeroHead marks children as belonging to the head part of a hero element.
func HeroHead(children ...any) heroPart {
	return heroPart{false, children}
}

// HeroFoot marks children as belonging to the foot part of a hero element.
func HeroFoot(children ...any) heroPart {
	return heroPart{true, children}
}
