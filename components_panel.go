package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

// Panel creates a panel element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func Panel(children ...any) e.Element {
	return e.Nav(e.Class("panel"), children)
}

// PanelHeading creates a panel heading element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelHeading(children ...any) e.Element {
	return e.P(e.Class("panel-heading"), children)
}

// PanelBlock creates a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelBlock(children ...any) e.Element {
	return e.Div(e.Class("panel-block"), children)
}

// PanelLink creates a link which is a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelLink(children ...any) e.Element {
	p := &panelLink{e.A(e.Class("panel-block"))}
	p.With(children...)
	return p
}

type panelLink struct {
	e.Element
}

func (p *panelLink) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			c.SetIconClass(e.Class("panel-icon"))
			p.Element.With(c)
		case []any:
			p.With(c...)
		default:
			p.Element.With(c)
		}
	}

	return p
}

func (p *panelLink) Clone() e.Element {
	return &panelLink{p.Element.Clone()}
}

// PanelAHref creates a link which is a panel block element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelAHref(href string, children ...any) e.Element {
	return PanelLink(html.Href(href), children)
}

// PanelTabs creates a panel tabs element.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelTabs(children ...any) e.Element {
	return e.P(e.Class("panel-tabs"), children)
}

// PanelCheckbox creates a label panel block element containing a checkbox.
//
// https://willoma.github.io/bulma-gomponents/panel.html
func PanelCheckbox(children ...any) e.Element {
	input := e.Input(html.Type("checkbox"))
	p := &panelCheckbox{
		Element: e.Label(e.Class("panel-block"), input),
		input:   input,
	}
	p.With(children...)
	return p
}

type panelCheckbox struct {
	e.Element
	input e.Element
}

func (p *panelCheckbox) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onLabel:
			p.Element.With(c...)
		case onInput:
			p.input.With(c...)
		case gomponents.Node:
			if e.IsAttribute(c) {
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

func (p *panelCheckbox) Clone() e.Element {
	return &panelCheckbox{
		Element: p.Element.Clone(),
		input:   p.input.Clone(),
	}
}
