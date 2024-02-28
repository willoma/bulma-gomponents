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
	head := Elem(html.Div).
		With(Class("modal-card-head"))

	body := Elem(html.Div).
		With(Class("modal-card-body"))

	foot := Elem(html.Div).
		With(Class("modal-card-foot"))

	for _, c := range children {
		switch c := c.(type) {
		case modalCardHead:
			head.Withs([]any(c))
		case modalCardFoot:
			foot.Withs([]any(c))
		default:
			body.With(c)
		}
	}

	return Elem(html.Div).
		With(Class("modal")).
		With(html.ID(id)).
		With(modalBackground()).
		With(
			Elem(html.Div).
				With(Class("modal-card")).
				With(head).
				With(body).
				With(foot),
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
