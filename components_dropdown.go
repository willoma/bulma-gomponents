package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Dropdown creates a dropdown.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func Dropdown(children ...any) Element {
	trigger := Elem(html.Div, Class("dropdown-trigger"))
	content := Elem(html.Div, Class("dropdown-content"))
	menu := Elem(html.Div, Class("dropdown-menu"), html.Role("menu"), content)

	d := &dropdown{
		Element: Elem(
			html.Div, Class("dropdown"),
			trigger,
			menu,
		),
		trigger: trigger,
		menu:    menu,
		content: content,
	}
	d.With(children...)
	return d
}

type dropdown struct {
	Element
	trigger Element
	menu    Element
	content Element

	up        bool
	clickable bool
	button    Element
	rendered  sync.Once
}

func (d *dropdown) With(childran ...any) Element {
	for _, c := range childran {
		switch c := c.(type) {
		case onDropdown:
			d.Element.With(c...)
		case onTrigger:
			for _, c2 := range c {
				if btn, ok := c2.(*button); ok {
					d.button = btn
				}
				d.trigger.With(c2)
			}
		case onMenu:
			d.menu.With(c...)
		case onContent:
			d.content.With(c...)
		case *dropdownButton:
			d.button = c
			d.trigger.With(c)
		case *dropdownDivider, *dropdownItem, *dropdownAhref:
			d.content.With(c)
		case ID:
			d.menu.With(c)
		case Class:
			switch c {
			case Clickable:
				d.Element.With(OnClick(JSToggleMe))
				d.clickable = true
			case Up:
				d.up = true
				d.Element.With(c)
			default:
				d.Element.With(c)
			}
		case string:
			d.button = DropdownButton(c)
			d.trigger.With(d.button)
		case Element:
			d.content.With(DropdownItem(c))
		case gomponents.Node:
			if isAttribute(c) {
				d.Element.With(c)
			} else {
				d.content.With(DropdownItem(c))
			}
		case []any:
			d.With(c...)
		default:
			d.Element.With(c)
		}
	}
	return d
}

func (d *dropdown) Render(w io.Writer) error {
	d.rendered.Do(func() {
		if d.button != nil {
			if d.up {
				if ddb, ok := d.button.(*dropdownButton); ok {
					ddb.up = true
				}
			}

			if d.clickable {
				d.button.With(On("blur", JSCloseThisDropdown))
			}
		}
	})

	return d.Element.Render(w)
}

// DropdownButton creates a button to be used as a dropdown trigger. It automatically adds a Font Awesome icon to the right if no icon is provided.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownButton(children ...any) Element {
	d := &dropdownButton{
		Element: Button(html.Aria("haspopup", "true")),
	}
	d.With(children...)
	return d
}

type dropdownButton struct {
	Element
	hasIcon bool
	up      bool

	rendered sync.Once
}

func (d *dropdownButton) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			d.hasIcon = true
			d.Element.With(c)
		case []any:
			d.With(c...)
		default:
			d.Element.With(c)
		}
	}

	return d
}

func (d *dropdownButton) Render(w io.Writer) error {
	d.rendered.Do(func() {
		if !d.hasIcon {
			var iconName string
			if d.up {
				iconName = "fa-angle-up"
			} else {
				iconName = "fa-angle-down"
			}
			d.Element.With(
				Icon(
					Elem(html.I, Small, Class("fa-solid"), Class(iconName)),
				),
			)
		}
	})

	return d.Element.Render(w)
}

// DropdownItem creates a div which is a dropdown item.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownItem(children ...any) Element {
	d := &dropdownItem{Elem(html.Div, Class("dropdown-item"))}
	d.With(children...)
	return d
}

type dropdownItem struct {
	Element
}

// DropdownAHref creates an AHref element which is a dropdown item.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownAHref(href string, children ...any) Element {
	d := &dropdownAhref{AHref(href, Class("dropdown-item"))}
	d.With(children...)
	return d
}

type dropdownAhref struct {
	Element
}

// DropdownDivider creates a dropdown divider.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownDivider(children ...any) Element {
	d := &dropdownDivider{Elem(html.Div, Class("dropdown-divider"))}
	d.With(children...)
	return d
}

type dropdownDivider struct {
	Element
}
