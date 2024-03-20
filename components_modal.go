package bulma

import (
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
// The id is needed at least in order to open the modal from another element.
//
// JSOpen may be used to get a javascript code to execute in order to open
// the modal, for instance:
//
//	b.Button(
//		b.OnClick(b.JSOpen("myModal"))
//	),
//	b.Modal("myModal", [...])
func Modal(id string, children ...any) Element {
	return Elem(
		html.Div,
		Class("modal"),
		html.ID(id),
		modalBackground(),
		Elem(html.Div, Class("modal-content"), children),
		modalClose(),
	)
}

// ModalCard creates a modal card.
//
// Wrap children in ModalCardHead in order to add them to the card header. Wrap
// children with ModalCardFoot in order to add them to the card footer. Any
// unwrapped child is added to the card body.
func ModalCard(id string, children ...any) Element {
	mc := &modalCard{id: id}
	mc.addChildren(children)
	return mc.elem()
}

type modalCard struct {
	id           string
	headChildren []any
	footChildren []any
	bodyChildren []any
}

func (mc *modalCard) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case modalCardHead:
			mc.headChildren = append(mc.headChildren, c...)
		case modalCardFoot:
			mc.footChildren = append(mc.footChildren, c...)
		case []any:
			mc.addChildren(c)
		default:
			mc.bodyChildren = append(mc.bodyChildren, c)
		}
	}
}

func (mc *modalCard) elem() Element {
	return Elem(
		html.Div,
		Class("modal"),
		html.ID(mc.id),
		modalBackground(),
		Elem(
			html.Div,
			Class("modal-card"),
			Elem(html.Div, Class("modal-card-head"), mc.headChildren),
			Elem(html.Div, Class("modal-card-body"), mc.bodyChildren),
			Elem(html.Div, Class("modal-card-foot"), mc.footChildren),
		),
		modalClose(),
	)
}

type modalCardHead []any

// ModalCardHead designates children to be part of the card head.
func ModalCardHead(children ...any) modalCardHead {
	return modalCardHead(children)
}

// ModalCardTitle creates a title for a card head.
func ModalCardTitle(children ...any) Element {
	return Elem(html.P, Class("modal-card-title"), children)
}

type modalCardFoot []any

// ModalCardFoot designates children to be part of the card head.
func ModalCardFoot(children ...any) modalCardFoot {
	return modalCardFoot(children)
}
