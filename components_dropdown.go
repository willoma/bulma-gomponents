package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Dropdown creates a dropdown.
//
// The triggerButton argument must be created with Button. The menu argument
// must be created with DropdownMenu.
//
// The following modifiers change the dropdown behaviour:
//   - Active: open the menu
//   - Hoverable: make it so the menu opens when the cursor hovers the button
func Dropdown(triggerButton, menu Element, children ...any) Element {
	return Elem(
		html.Div,
		Class("dropdown"),
		children,
		Elem(html.Div, Class("dropdown-trigger"), triggerButton),
		menu,
	)
}

// Dropup creates a dropdown which opens to the top.
//
// The triggerButton argument must be created with Button. The menu argument
// must be created with DropdownMenu.
//
// The following modifiers change the dropdown behaviour:
//   - Active: open the menu
//   - Hoverable: make it so the menu opens when the cursor hovers the button
func Dropup(triggerButton, menu Element, children ...any) Element {
	return Dropdown(triggerButton, menu, Class("is-up"), children)
}

// DropdownMenu creates a dropdown menu.
//   - when a child is marked with b.Inner, it is forcibly applied to the <div class="dropdown-content"> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <div class="dropdown-menu"> element
//
// Each child is automatically wrapped in a new dropdown-item element if:
//   - it is an Element but not generated with DropdownItem or DropdownHref or DropdownDivider
//   - it is a gomponents.Node with type gomponents.ElementType
func DropdownMenu(children ...any) Element {
	return new(dropdownMenu).With(children...)
}

type dropdownMenu struct {
	contentChildren []any
	menuChildren    []any
}

func (dm *dropdownMenu) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *ApplyToInner:
			dm.contentChildren = append(dm.contentChildren, c.Child)
		case *ApplyToOuter:
			dm.menuChildren = append(dm.menuChildren, c.Child)
		case *dropdownDivider, *dropdownItem, *dropdownAhref:
			dm.contentChildren = append(dm.contentChildren, c)
		case Element:
			dm.contentChildren = append(dm.contentChildren, DropdownItem(c))
		case gomponents.Node:
			if IsAttribute(c) {
				dm.menuChildren = append(dm.menuChildren, c)
			} else {
				dm.contentChildren = append(dm.contentChildren, c)
			}
		case []any:
			dm.With(c...)
		default:
			dm.menuChildren = append(dm.menuChildren, c)
		}
	}

	return dm
}

func (dm *dropdownMenu) Render(w io.Writer) error {
	return Elem(
		html.Div,
		Class("dropdown-menu"),
		html.Role("menu"),
		dm.menuChildren,
		Elem(html.Div, Class("dropdown-content"), dm.contentChildren),
	).Render(w)
}

// DropdownItem creates a div which is a dropdown item.
func DropdownItem(children ...any) Element {
	return &dropdownItem{Elem(html.Div, Class("dropdown-item"), children)}
}

type dropdownItem struct {
	Element
}

// DropdownAHref creates an AHref element which is a dropdown item.
func DropdownAHref(href string, children ...any) Element {
	return &dropdownAhref{AHref(href, Class("dropdown-item"), children)}
}

type dropdownAhref struct {
	Element
}

// DropdownDivider creates a dropdown divider.
func DropdownDivider() Element {
	return &dropdownDivider{Elem(html.Hr, Class("dropdown-divider"))}
}

type dropdownDivider struct {
	Element
}
