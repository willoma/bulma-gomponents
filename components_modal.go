package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func modalBackground() Element {
	return Elem(html.Div, Class("modal-background"), OnClick(JSCloseThisModal))
}

func modalClose() Element {
	return Elem(
		html.Button,
		Class("modal-close"),
		Large,
		html.Aria("label", "close"),
		OnClick(JSCloseThisModal),
	)
}

// Modal creates a modal. The modal background and the modal close button are automatically added. The children are added to the modal content.
//
// Provide an ID with html.ID in order to identify the modal and, for example,
// open it from another element.
//
// JSOpen may be used to get a javascript code to execute in order to open
// the modal, for instance:
//
//	b.Button(
//		b.OnClick(b.JSOpen("myModal"))
//	),
//	b.Modal(html.ID("myModal"), [...])
func Modal(children ...any) Element {
	return new(modal).With(children...)
}

type modal struct {
	ownChildren     []any
	contentChildren []any
}

func (m *modal) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onModal:
			m.ownChildren = append(m.ownChildren, c...)
		case gomponents.Node:
			if IsAttribute(c) {
				m.ownChildren = append(m.ownChildren, c)
			} else {
				m.contentChildren = append(m.contentChildren, c)
			}
		case []any:
			m.With(c...)
		default:
			m.contentChildren = append(m.contentChildren, c)
		}
	}
	return m
}

func (m *modal) Render(w io.Writer) error {
	return Elem(
		html.Div,
		Class("modal"),
		modalBackground(),
		m.ownChildren,
		Elem(html.Div, Class("modal-content"), m.contentChildren),
		modalClose(),
	).Render(w)
}

// ModalCard creates a modal card.
//
// Wrap children in ModalCardHead in order to add them to the card header. Wrap
// children with ModalCardFoot in order to add them to the card footer. Any
// unwrapped child is added to the card body.
func ModalCard(children ...any) Element {
	return new(modalCard).With(children...)
}

type modalCard struct {
	ownChildren  []any
	cardChildren []any
	headChildren []any
	footChildren []any
	bodyChildren []any
}

func (mc *modalCard) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onModal:
			mc.ownChildren = append(mc.ownChildren, c...)
		case onCard:
			mc.cardChildren = append(mc.cardChildren, c...)
		case modalCardHead:
			mc.headChildren = append(mc.headChildren, c...)
		case modalCardTitle:
			mc.headChildren = append(mc.headChildren, c)
		case modalCardFoot:
			mc.footChildren = append(mc.footChildren, c...)
		case gomponents.Node:
			if IsAttribute(c) {
				mc.ownChildren = append(mc.ownChildren, c)
			} else {
				mc.bodyChildren = append(mc.bodyChildren, c)
			}
		case []any:
			mc.With(c...)
		default:
			mc.bodyChildren = append(mc.bodyChildren, c)
		}
	}

	return mc
}

func (mc *modalCard) Render(w io.Writer) error {
	return Elem(
		html.Div,
		Class("modal"),
		mc.ownChildren,
		modalBackground(),
		Elem(
			html.Div,
			Class("modal-card"),
			mc.cardChildren,
			Elem(html.Div, Class("modal-card-head"), mc.headChildren),
			Elem(html.Div, Class("modal-card-body"), mc.bodyChildren),
			Elem(html.Div, Class("modal-card-foot"), mc.footChildren),
		),
		modalClose(),
	).Render(w)
}

// ModalCardHead designates children to be part of the card head.
//
// It cannot be used as a node by itself.
func ModalCardHead(children ...any) modalCardHead {
	return modalCardHead(children)
}

type modalCardHead []any

// ModalCardTitle creates a title for a card head.
func ModalCardTitle(children ...any) Element {
	return new(modalCardTitle).With(children...)
}

type modalCardTitle struct {
	children []any
}

func (mct *modalCardTitle) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			mct.With(c...)
		default:
			mct.children = append(mct.children, c)
		}
	}

	return mct
}

func (mct *modalCardTitle) Render(w io.Writer) error {
	return Elem(html.P, Class("modal-card-title"), mct.children).Render(w)
}

// ModalCardHeadTitle creates a card head with a text title and a close button.
func ModalCardTitleWithClose(title string) any {
	return ModalCardHead(
		ModalCardTitle(title),
		Delete(
			OnClick(JSCloseThisModal),
			html.Aria("label", "close"),
		),
	)
}

// ModalCardFoot designates children to be part of the card head.
//
// It cannot be used as a node by itself.
func ModalCardFoot(children ...any) modalCardFoot {
	return modalCardFoot(children)
}

type modalCardFoot []any
