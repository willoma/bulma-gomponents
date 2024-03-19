package bulma

import (
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
func Panel(children ...any) *Element {
	return Elem(html.Nav).
		With(Class("panel")).
		Withs(children)
}

// PanelHeading creates a panel heading element.
func PanelHeading(children ...any) *Element {
	return Elem(html.P).
		With(Class("panel-heading")).
		Withs(children)
}

// PanelBlock creates a panel block element.
func PanelBlock(children ...any) *Element {
	return Elem(html.Div).
		With(Class("panel-block")).
		Withs(children)
}

// PanelTabs creates a panel tabs element.
//
// Its children must be "a" html elements (for instance AHref). Add the Active
// modifier to a link to mark it as the active tab.
func PanelTabs(children ...any) *Element {
	return Elem(html.P).
		With(Class("panel-tabs")).
		Withs(children)
}

// PanelLink creates a link which is a panel block element.
func PanelLink(children ...any) *Element {
	p := panelLink{}
	p.addChildren(children)
	return p.elem()
}

type panelLink struct {
	children []any
}

func (p *panelLink) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("icon") {
				c.With(Class("panel-icon"))
				c.classes["icon"] = false
			}
			p.children = append(p.children, c)
		case []any:
			p.addChildren(c)
		default:
			p.children = append(p.children, c)
		}
	}
}

func (p *panelLink) elem() *Element {
	return Elem(html.A).
		With(Class("panel-block")).
		Withs(p.children)
}

// PanelLabel creates a label which is a panel block element, which must contain
// a Checkbox.
func PanelLabel(children ...any) *Element {
	return Elem(html.Label).
		With(Class("panel-block")).
		Withs(children)
}
