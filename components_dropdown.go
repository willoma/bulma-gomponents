package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Dropdown creates a dropdown.
//
// The following children have a special meaning:
//   - OnDropdown: apply the children to the dropdown
//   - OnTrigger: apply the children to the trigger
//   - OnMenu: apply the children to the menu
//   - OnContent: apply the children to the menu content
//   - DropdownTrigger: define the trigger button
//   - DropdownItem: add an item to the menu
//   - DropdownAHref: add a link item to the menu
//   - DropdownDivider: add a divider to the menu
//
// The following modifiers change the dropdown behaviour:
//   - Active: open the menu
//   - Clickable: make it so the menu opens when the button is clicked (javascript is automatically added)
//   - Hoverable: make it so the menu opens when the cursor hovers the button
//   - Up: make the dropdown open to the top
func Dropdown(children ...any) Element {
	return new(dropdown).With(children...)
}

type dropdown struct {
	up              bool
	clickable       bool
	children        []any
	triggerChildren []any
	contentChildren []any
	menuChildren    []any
	button          Element
}

func (d *dropdown) With(childran ...any) Element {
	for _, c := range childran {
		switch c := c.(type) {
		case onDropdown:
			d.children = append(d.children, c...)
		case onTrigger:
			d.triggerChildren = append(d.triggerChildren, c...)
		case onMenu:
			d.menuChildren = append(d.menuChildren, c...)
		case onContent:
			d.contentChildren = append(d.contentChildren, c...)
		case dropdownTrigger:
			for _, c2 := range c {
				if btn, ok := c2.(*button); ok {
					d.button = btn
				}
				d.triggerChildren = append(d.triggerChildren, c2)
			}
		case *dropdownButton:
			d.button = c
			d.triggerChildren = append(d.triggerChildren, c)
		case *dropdownDivider, *dropdownItem, *dropdownAhref:
			d.contentChildren = append(d.contentChildren, c)
		case ID:
			d.menuChildren = append(d.menuChildren, c)
		case Class:
			switch c {
			case Clickable:
				d.clickable = true
			case Up:
				d.up = true
				d.children = append(d.children, c)
			default:
				d.children = append(d.children, c)
			}
		case Element:
			d.contentChildren = append(d.contentChildren, DropdownItem(c))
		case gomponents.Node:
			if IsAttribute(c) {
				d.children = append(d.children, c)
			} else {
				d.contentChildren = append(d.contentChildren, DropdownItem(c))
			}
		case []any:
			d.With(c...)
		default:
			d.children = append(d.children, c)
		}
	}
	return d
}

func (d *dropdown) Render(w io.Writer) error {
	if d.up && d.button != nil {
		if ddb, ok := d.button.(*dropdownButton); ok {
			ddb.up = true
		}
	}

	if d.clickable {
		d.children = append(d.children, OnClick(JSToggleMe))
		if d.button != nil {
			d.button.With(On("blur", JSCloseThisDropdown))
		}
	}

	return Elem(
		html.Div, Class("dropdown"),
		d.children,
		Elem(html.Div, Class("dropdown-trigger"), d.triggerChildren),
		Elem(
			html.Div, Class("dropdown-menu"), html.Role("menu"),
			d.menuChildren, Elem(html.Div, Class("dropdown-content"), d.contentChildren),
		),
	).Render(w)
}

// DropdownButton creates a button to be used as a dropdown trigger. It automatically adds a Font Awesome icon to the right if no icon is provided.
func DropdownButton(children ...any) Element {
	return new(dropdownButton).With(children...)
}

type dropdownButton struct {
	up       bool
	hasIcon  bool
	children []any
}

func (d *dropdownButton) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case IconElem:
			d.hasIcon = true
			d.children = append(d.children, c)
		case []any:
			d.With(c...)
		default:
			d.children = append(d.children, c)
		}
	}

	return d
}

func (d *dropdownButton) Render(w io.Writer) error {
	b := Button(d.children...).With(
		html.Aria("haspopup", "true"),
	)

	if !d.hasIcon {
		var iconName string
		if d.up {
			iconName = "fa-angle-up"
		} else {
			iconName = "fa-angle-down"
		}
		b.With(
			Icon(
				Elem(html.I, Small, Class("fa-solid"), Class(iconName)),
			),
		)
	}

	return b.Render(w)
}

func DropdownTrigger(children ...any) dropdownTrigger {
	return dropdownTrigger(children)
}

type dropdownTrigger []any

// DropdownItem creates a div which is a dropdown item.
func DropdownItem(children ...any) Element {
	return new(dropdownItem).With(children...)
}

type dropdownItem struct {
	children []any
}

func (d *dropdownItem) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			d.With(c...)
		default:
			d.children = append(d.children, c)
		}
	}

	return d
}

func (d *dropdownItem) Render(w io.Writer) error {
	return Elem(html.Div, Class("dropdown-item"), d.children).Render(w)
}

// DropdownAHref creates an AHref element which is a dropdown item.
func DropdownAHref(href string, children ...any) Element {
	return (&dropdownAhref{href: href}).With(children...)
}

type dropdownAhref struct {
	href     string
	children []any
}

func (d *dropdownAhref) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			d.With(c...)
		default:
			d.children = append(d.children, c)
		}
	}

	return d
}

func (d *dropdownAhref) Render(w io.Writer) error {
	return AHref(d.href, Class("dropdown-item"), d.children).Render(w)
}

// DropdownDivider creates a dropdown divider.
func DropdownDivider(children ...any) Element {
	return new(dropdownDivider).With(children...)
}

type dropdownDivider struct {
	children []any
}

func (d *dropdownDivider) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			d.With(c...)
		default:
			d.children = append(d.children, c)
		}
	}
	return d
}

func (d *dropdownDivider) Render(w io.Writer) error {
	return Elem(html.Div, Class("dropdown-divider"), d.children).Render(w)
}
