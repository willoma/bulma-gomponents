package bulma

import (
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
	return Dropdown(triggerButton, menu, children, Class("is-up"))
}

// DropdownItem creates a div which is a dropdown item.
func DropdownItem(children ...any) Element {
	return Elem(html.Div, Class("dropdown-item"), children)
}

// DropdownAHref creates an AHref element which is a dropdown item.
func DropdownAHref(href string, children ...any) Element {
	return AHref(href, Class("dropdown-item"), children)
}

// DropdownDivider creates a dropdown divider.
func DropdownDivider() Element {
	return Elem(html.Hr, Class("dropdown-divider"))
}

// DropdownMenu creates a dropdown menu.
//
// Each child is automatically wrapped in a new dropdown-item element if:
//   - it is an Element but not generated with DropdownItem or DropdownHref or DropdownDivider
//   - it is a gomponents.Node with type gomponents.ElementType
func DropdownMenu(children ...any) Element {
	dm := &dropdownMenu{}
	dm.addChildren(children)
	return dm.elem()
}

type dropdownMenu struct {
	contentChildren []any
	menuChildren    []any
}

func (dm *dropdownMenu) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			if c.hasClass("dropdown-item") || c.hasClass("dropdown-divider") {
				dm.contentChildren = append(dm.contentChildren, c)
			} else {
				dm.contentChildren = append(dm.contentChildren, DropdownItem(c))
			}
		case gomponents.Node:
			if IsAttribute(c) {
				dm.menuChildren = append(dm.menuChildren, c)
			} else {
				dm.contentChildren = append(dm.contentChildren, c)
			}
		default:
			dm.menuChildren = append(dm.menuChildren, c)
		}
	}
}

func (dm *dropdownMenu) elem() Element {
	return Elem(
		html.Div,
		Class("dropdown-menu"),
		html.Role("menu"),
		dm.menuChildren,
		Elem(html.Div, Class("dropdown-content"), dm.contentChildren),
	)
}
