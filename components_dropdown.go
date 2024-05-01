package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	e "github.com/willoma/gomplements"
)

// Dropdown creates a dropdown.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func Dropdown(children ...any) e.Element {
	trigger := e.Div(e.Class("dropdown-trigger"))
	content := e.Div(e.Class("dropdown-content"))
	menu := e.Div(e.Class("dropdown-menu"), e.AriaMenu, content)

	d := &dropdown{
		dropdown: e.Div(
			e.Class("dropdown"),
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
	dropdown e.Element
	trigger  e.Element
	menu     e.Element
	content  e.Element

	up        bool
	clickable bool
	button    e.Element
	rendered  sync.Once
}

func (d *dropdown) With(childran ...any) e.Element {
	for _, c := range childran {
		switch c := c.(type) {
		case onDropdown:
			d.dropdown.With(c...)
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
		case e.ID:
			d.menu.With(c)
		case e.Class:
			switch c {
			case Clickable:
				d.dropdown.With(e.OnClick(JSToggleMe))
				d.clickable = true
			case Up:
				d.up = true
				d.dropdown.With(c)
			default:
				d.dropdown.With(c)
			}
		case string:
			d.button = DropdownButton(c)
			d.trigger.With(d.button)
		case e.Element:
			d.content.With(DropdownItem(c))
		case gomponents.Node:
			if e.IsAttribute(c) {
				d.dropdown.With(c)
			} else {
				d.content.With(DropdownItem(c))
			}
		case []any:
			d.With(c...)
		default:
			d.dropdown.With(c)
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
				d.button.With(e.On("blur", JSCloseThisDropdown))
			}
		}
	})

	return d.dropdown.Render(w)
}

func (d *dropdown) Clone() e.Element {
	return &dropdown{
		dropdown: d.dropdown.Clone(),
		trigger:  d.trigger.Clone(),
		menu:     d.menu.Clone(),
		content:  d.content.Clone(),

		up:        d.up,
		clickable: d.clickable,
		button:    d.button.Clone(),
	}
}

// DropdownButton creates a button to be used as a dropdown trigger. It automatically adds a Font Awesome icon to the right if no icon is provided.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownButton(children ...any) e.Element {
	d := &dropdownButton{
		button: Button(e.AriaHasPopupTrue),
	}
	d.With(children...)
	return d
}

type dropdownButton struct {
	button  e.Element
	hasIcon bool
	up      bool

	rendered sync.Once
}

func (d *dropdownButton) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			d.hasIcon = true
			d.button.With(c)
		case []any:
			d.With(c...)
		default:
			d.button.With(c)
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
			d.button.With(
				Icon(
					e.I(Small, e.Class("fa-solid"), e.Class(iconName)),
				),
			)
		}
	})

	return d.button.Render(w)
}

func (d *dropdownButton) Clone() e.Element {
	return &dropdownButton{
		button:  d.button.Clone(),
		hasIcon: d.hasIcon,
		up:      d.up,
	}
}

// DropdownItem creates a div which is a dropdown item.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownItem(children ...any) e.Element {
	d := &dropdownItem{e.Div(e.Class("dropdown-item"))}
	d.With(children...)
	return d
}

type dropdownItem struct {
	e.Element
}

func (d *dropdownItem) Clone() e.Element {
	return &dropdownItem{d.Element.Clone()}
}

// DropdownAHref creates an AHref element which is a dropdown item.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownAHref(href string, children ...any) e.Element {
	d := &dropdownAhref{e.AHref(href, e.Class("dropdown-item"))}
	d.With(children...)
	return d
}

type dropdownAhref struct {
	e.Element
}

func (d *dropdownAhref) Clone() e.Element {
	return &dropdownAhref{d.Element.Clone()}
}

// DropdownDivider creates a dropdown divider.
//
// https://willoma.github.io/bulma-gomponents/dropdown.html
func DropdownDivider(children ...any) e.Element {
	d := &dropdownDivider{e.Div(e.Class("dropdown-divider"))}
	d.With(children...)
	return d
}

type dropdownDivider struct {
	e.Element
}

func (d *dropdownDivider) Clone() e.Element {
	return &dropdownDivider{d.Element.Clone()}
}
