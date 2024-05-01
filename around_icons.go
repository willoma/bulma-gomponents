package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

type elemOptionSpanAroundNonIconsIfHasIcons struct {
	elemFn func(...gomponents.Node) gomponents.Node

	hasIcons bool
	children []any
}

func (el *elemOptionSpanAroundNonIconsIfHasIcons) With(children ...any) e.Element {
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

func (el *elemOptionSpanAroundNonIconsIfHasIcons) Render(w io.Writer) error {
	if el.hasIcons {
		elem := e.Elem(el.elemFn)

		for _, c := range el.children {
			switch c := c.(type) {
			case IconElem:
				elem.With(c)
			case e.Element:
				elem.With(e.Span(c))
			default:
				elem.With(c)
			}
		}

		return elem.Render(w)
	}

	return e.Elem(el.elemFn, el.children...).Render(w)
}

func (el *elemOptionSpanAroundNonIconsIfHasIcons) Clone() e.Element {
	children := make([]any, len(el.children))

	for i, c := range el.children {
		switch c := c.(type) {
		case e.Element:
			children[i] = c.Clone()
		default:
			children[i] = c
		}
	}

	return &elemOptionSpanAroundNonIconsIfHasIcons{
		elemFn:   el.elemFn,
		hasIcons: el.hasIcons,
		children: el.children,
	}
}

type elemOptionSpanAroundNonIcons struct {
	elemFn func(...gomponents.Node) gomponents.Node

	children []any
}

func (el *elemOptionSpanAroundNonIcons) With(children ...any) e.Element {
	el.children = append(el.children, children...)

	return el
}

func (el *elemOptionSpanAroundNonIcons) Render(w io.Writer) error {
	elem := e.Elem(el.elemFn)

	for _, c := range el.children {
		switch c := c.(type) {
		case IconElem:
			elem.With(c)
		case e.Element:
			elem.With(e.Span(c))
		default:
			elem.With(c)
		}
	}

	return elem.Render(w)
}

func (el *elemOptionSpanAroundNonIcons) Clone() e.Element {
	children := make([]any, len(el.children))

	for i, c := range el.children {
		switch c := c.(type) {
		case e.Element:
			children[i] = c.Clone()
		default:
			children[i] = c
		}
	}

	return &elemOptionSpanAroundNonIconsIfHasIcons{
		elemFn:   el.elemFn,
		children: el.children,
	}
}
