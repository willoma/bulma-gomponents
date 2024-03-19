package bulma

import (
	"github.com/maragudk/gomponents/html"
)

func modalBackground() *Element {
	return Elem(html.Div).
		With(Class("modal-background")).
		With(OnClick(JSCloseThisModal))
}

func modalClose() *Element {
	return Elem(html.Button).
		With(Class("modal-close")).
		With(Large).
		With(html.Aria("label", "close")).
		With(OnClick(JSCloseThisModal))
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
func Modal(id string, children ...any) *Element {
	return Elem(html.Div).
		With(Class("modal")).
		With(html.ID(id)).
		With(modalBackground()).
		With(
			Elem(html.Div).
				With(Class("modal-content")).
				Withs(children),
		).
		With(modalClose())
}

// ModalCard creates a modal card.
//
// Wrap children in ModalCardHead in order to add them to the card header. Wrap
// children with ModalCardFoot in order to add them to the card footer. Any
// unwrapped child is added to the card body.
func ModalCard(id string, children ...any) *Element {
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

func (mc *modalCard) elem() *Element {
	return Elem(html.Div).
		With(Class("modal")).
		With(html.ID(mc.id)).
		With(modalBackground()).
		With(
			Elem(html.Div).
				With(Class("modal-card")).
				With(
					Elem(html.Div).
						With(Class("modal-card-head")).
						Withs(mc.headChildren),
				).
				With(
					Elem(html.Div).
						With(Class("modal-card-body")).
						Withs(mc.bodyChildren),
				).
				With(
					Elem(html.Div).
						With(Class("modal-card-foot")).
						Withs(mc.footChildren),
				),
		).
		With(modalClose())
}

type modalCardHead []any

// ModalCardHead designates children to be part of the card head.
func ModalCardHead(children ...any) modalCardHead {
	return modalCardHead(children)
}

// ModalCardTitle creates a title for a card head.
func ModalCardTitle(children ...any) *Element {
	return Elem(html.P).
		With(Class("modal-card-title")).
		Withs(children)
}

type modalCardFoot []any

// ModalCardFoot designates children to be part of the card head.
func ModalCardFoot(children ...any) modalCardFoot {
	return modalCardFoot(children)
}
