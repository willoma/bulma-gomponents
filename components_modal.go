package bulma

import (
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
// https://willoma.github.io/bulma-gomponents/modal.html
func Modal(children ...any) Element {
	content := Elem(html.Div, Class("modal-content"))
	m := &modal{
		Element: Elem(
			html.Div, Class("modal"),
			modalBackground(),
			content,
			modalClose(),
		),
		content: content,
	}
	m.With(children...)
	return m
}

type modal struct {
	Element
	content Element
}

func (m *modal) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onModal:
			m.Element.With(c...)
		case onContent:
			m.content.With(c...)
		case gomponents.Node:
			if isAttribute(c) {
				m.Element.With(c)
			} else {
				m.content.With(c)
			}
		case []any:
			m.With(c...)
		default:
			m.content.With(c)
		}
	}
	return m
}

func (m *modal) Clone() Element {
	return &modal{
		Element: m.Element.Clone(),
		content: m.content.Clone(),
	}
}

// ModalCard creates a modal card.
//
// Wrap children in ModalCardHead in order to add them to the card header. Wrap
// children with ModalCardFoot in order to add them to the card footer. Any
// unwrapped child is added to the card body.
//
// https://willoma.github.io/bulma-gomponents/modal.html
func ModalCard(children ...any) Element {
	head := Elem(html.Div, Class("modal-card-head"))
	body := Elem(html.Div, Class("modal-card-body"))
	foot := Elem(html.Div, Class("modal-card-foot"))

	card := Elem(html.Div, Class("modal-card"), head, body, foot)
	m := &modalCard{
		Element: Elem(
			html.Div, Class("modal"),
			modalBackground(),
			card,
			modalClose(),
		),
		card: card,
		head: head,
		body: body,
		foot: foot,
	}
	m.With(children...)
	return m
}

type modalCard struct {
	Element
	card Element
	head Element
	body Element
	foot Element
}

func (m *modalCard) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onModal:
			m.Element.With(c...)
		case onCard:
			m.card.With(c...)
		case modalCardHead:
			m.head.With(c...)
		case *modalCardTitle:
			m.head.With(c)
		case modalCardFoot:
			m.foot.With(c...)
		case gomponents.Node:
			if isAttribute(c) {
				m.Element.With(c)
			} else {
				m.body.With(c)
			}
		case []any:
			m.With(c...)
		default:
			m.body.With(c)
		}
	}

	return m
}

func (m *modalCard) Clone() Element {
	return &modalCard{
		Element: m.Element.Clone(),
		card:    m.card.Clone(),
		head:    m.head.Clone(),
		body:    m.body.Clone(),
		foot:    m.foot.Clone(),
	}
}

// ModalCardHead designates children to be part of the card head.
//
// https://willoma.github.io/bulma-gomponents/modal.html
func ModalCardHead(children ...any) modalCardHead {
	return modalCardHead(children)
}

type modalCardHead []any

// ModalCardTitle creates a title for a card head.
//
// https://willoma.github.io/bulma-gomponents/modal.html
func ModalCardTitle(children ...any) Element {
	m := &modalCardTitle{Elem(html.P, Class("modal-card-title"))}
	m.With(children...)
	return m
}

type modalCardTitle struct {
	Element
}

func (m *modalCardTitle) Clone() Element {
	return &modalCardTitle{m.Element.Clone()}
}

// ModalCardTitleWithClose creates a card head with a text title and a close
// button.
//
// https://willoma.github.io/bulma-gomponents/modal.html
func ModalCardTitleWithClose(title string) modalCardHead {
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
// https://willoma.github.io/bulma-gomponents/modal.html
func ModalCardFoot(children ...any) modalCardFoot {
	return modalCardFoot(children)
}

type modalCardFoot []any
