package bulma

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
)

func modalBackground() e.Element {
	return e.Div(e.Class("modal-background"), e.OnClick(JSCloseThisModal))
}

func modalClose() e.Element {
	return e.Button(
		e.Class("modal-close"),
		Large,
		e.AriaLabel("close"),
		e.OnClick(JSCloseThisModal),
	)
}

// Modal creates a modal. The modal background and the modal close button are automatically added. The children are added to the modal content.
//
// https://willoma.github.io/bulma-gomponents/modal.html
func Modal(children ...any) e.Element {
	content := e.Div(e.Class("modal-content"))
	m := &modal{
		Element: e.Div(
			e.Class("modal"),
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
	e.Element
	content e.Element
}

func (m *modal) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onModal:
			m.Element.With(c...)
		case onContent:
			m.content.With(c...)
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (m *modal) Clone() e.Element {
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
func ModalCard(children ...any) e.Element {
	head := e.Div(e.Class("modal-card-head"))
	body := e.Div(e.Class("modal-card-body"))
	foot := e.Div(e.Class("modal-card-foot"))

	card := e.Div(e.Class("modal-card"), head, body, foot)
	m := &modalCard{
		Element: e.Div(
			e.Class("modal"),
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
	e.Element
	card e.Element
	head e.Element
	body e.Element
	foot e.Element
}

func (m *modalCard) With(children ...any) e.Element {
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
			if e.IsAttribute(c) {
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

func (m *modalCard) Clone() e.Element {
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
func ModalCardTitle(children ...any) e.Element {
	m := &modalCardTitle{e.P(e.Class("modal-card-title"))}
	m.With(children...)
	return m
}

type modalCardTitle struct {
	e.Element
}

func (m *modalCardTitle) Clone() e.Element {
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
			e.OnClick(JSCloseThisModal),
			e.AriaLabel("close"),
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
