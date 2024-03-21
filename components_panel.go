package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Panel creates a panel element.
//
// The following modifiers change the panel color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
func Panel(children ...any) Element {
	return Elem(html.Nav, Class("panel"), children)
}

// PanelHeading creates a panel heading element.
func PanelHeading(children ...any) Element {
	return Elem(html.P, Class("panel-heading"), children)
}

// PanelBlock creates a panel block element.
func PanelBlock(children ...any) Element {
	return Elem(html.Div, Class("panel-block"), children)
}

// PanelTabs creates a panel tabs element.
//
// Its children must be "a" html elements (for instance AHref). Add the Active
// modifier to a link to mark it as the active tab.
func PanelTabs(children ...any) Element {
	return Elem(html.P, Class("panel-tabs"), children)
}

// PanelLink creates a link which is a panel block element.
func PanelLink(children ...any) Element {
	return new(panelLink).With(children...)
}

type panelLink struct {
	children []any
}

func (p *panelLink) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *icon:
			c.iconClass = Class("panel-icon")
			p.children = append(p.children, c)
		case Element:
			p.children = append(p.children, c)
		case []any:
			p.With(c...)
		default:
			p.children = append(p.children, c)
		}
	}

	return p
}

func (p *panelLink) Render(w io.Writer) error {
	return Elem(html.A, Class("panel-block"), p.children).Render(w)
}

// PanelLabel creates a label which is a panel block element, which must contain
// a Checkbox.
func PanelLabel(children ...any) Element {
	return Elem(html.Label, Class("panel-block"), children)
}
