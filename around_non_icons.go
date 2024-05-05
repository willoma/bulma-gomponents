package bulma

import (
	"fmt"
	"io"

	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

type spanAroundNonIconsIfHasIcons struct {
	elemFn func(...gomponents.Node) gomponents.Node

	hasIcons bool
	children []any
}

func (el *spanAroundNonIconsIfHasIcons) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			el.hasIcons = true
			el.children = append(el.children, c)
		case []any:
			el.With(c...)
		default:
			el.children = append(el.children, c)
		}
	}

	return el
}

func (el *spanAroundNonIconsIfHasIcons) Render(w io.Writer) error {
	if el.hasIcons {
		elem := &spanAroundNonIcons{e.Elem(el.elemFn)}
		elem.With(el.children...)
		return elem.Render(w)
	}

	return e.Elem(el.elemFn, el.children...).Render(w)
}

func (el *spanAroundNonIconsIfHasIcons) Clone() e.Element {
	children := make([]any, len(el.children))

	for i, c := range el.children {
		switch c := c.(type) {
		case e.Element:
			children[i] = c.Clone()
		default:
			children[i] = c
		}
	}

	return &spanAroundNonIconsIfHasIcons{
		elemFn:   el.elemFn,
		hasIcons: el.hasIcons,
		children: el.children,
	}
}

type spanAroundNonIcons struct {
	e.Element
}

func (el *spanAroundNonIcons) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem, e.Class:
			// Classes must be explicitly added, because e.Class is a string
			el.Element.With(c)
		case e.ParentModifierAndNode, gomponents.Node:
			// Must be before fmt.Stringer because gomponents.Node is a stringer
			if e.IsAttribute(c) {
				el.Element.With(c)
			} else {
				el.Element.With(e.Span(c))
			}
		case string, fmt.Stringer:
			el.Element.With(e.Span(c))
		case []any:
			el.With(c...)
		default:
			el.Element.With(c)
		}
	}

	return el
}
