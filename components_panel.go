package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Panel creates a panel element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func Panel(children ...any) Element {
	return Elem(html.Nav, Class("panel"), children)
}

// PanelHeading creates a panel heading element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelHeading(children ...any) Element {
	return Elem(html.P, Class("panel-heading"), children)
}

// PanelBlock creates a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelBlock(children ...any) Element {
	return Elem(html.Div, Class("panel-block"), children)
}

// PanelLink creates a link which is a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelLink(children ...any) Element {
	p := &panelLink{Elem(html.A, Class("panel-block"))}
	p.With(children...)
	return p
}

type panelLink struct {
	Element
}

func (p *panelLink) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			c.SetIconClass(Class("panel-icon"))
			p.Element.With(c)
		case []any:
			p.With(c...)
		default:
			p.Element.With(c)
		}
	}

	return p
}

func (p *panelLink) Clone() Element {
	return &panelLink{p.Element.Clone()}
}

// PanelAHref creates a link which is a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelAHref(href string, children ...any) Element {
	return PanelLink(html.Href(href), children)
}

// PanelTabs creates a panel tabs element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelTabs(children ...any) Element {
	return Elem(html.P, Class("panel-tabs"), children)
}

// PanelCheckbox creates a label panel block element containing a checkbox.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelCheckbox(children ...any) Element {
	input := Elem(html.Input, html.Type("checkbox"))
	p := &panelCheckbox{
		Element: Elem(html.Label, Class("panel-block"), input),
		input:   input,
	}
	p.With(children...)
	return p
}

type panelCheckbox struct {
	Element
	input Element
}

func (p *panelCheckbox) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onLabel:
			p.Element.With(c...)
		case onInput:
			p.input.With(c...)
		case gomponents.Node:
			if isAttribute(c) {
				p.input.With(c)
			} else {
				p.Element.With(c)
			}
		case []any:
			p.With(c...)
		default:
			p.Element.With(c)
		}
	}

	return p
}

func (p *panelCheckbox) Clone() Element {
	return &panelCheckbox{
		Element: p.Element.Clone(),
		input:   p.input.Clone(),
	}
}
