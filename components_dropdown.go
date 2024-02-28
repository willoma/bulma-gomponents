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
func Dropdown(triggerButton, menu *Element, children ...any) *Element {
	return Elem(html.Div).
		With(Class("dropdown")).
		Withs(children).
		With(
			Elem(html.Div).
				With(Class("dropdown-trigger")).
				With(triggerButton),
		).
		With(menu)
}

// Dropup creates a dropdown which opens to the top.
//
// The triggerButton argument must be created with Button. The menu argument
// must be created with DropdownMenu.
//
// The following modifiers change the dropdown behaviour:
//   - Active: open the menu
//   - Hoverable: make it so the menu opens when the cursor hovers the button
func Dropup(triggerButton, menu *Element, children ...any) *Element {
	return Elem(html.Div).
		With(Class("dropdown")).
		With(Class("is-up")).
		Withs(children)
}

// DropdownItem creates a div which is a dropdown item.
func DropdownItem(children ...any) *Element {
	return Elem(html.Div).
		With(Class("dropdown-item")).
		Withs(children)
}

// DropdownAHref creates an AHref element which is a dropdown item.
func DropdownAHref(href string, children ...any) *Element {
	return AHref(href, Class("dropdown-item")).Withs(children)
}

// DropdownDivider creates a dropdown divider.
func DropdownDivider() *Element {
	return Elem(html.Hr).
		With(Class("dropdown-divider"))
}

// DropdownMenu creates a dropdown menu.
//
// Each child is automatically wrapped in a new dropdown-item element if:
//   - it is an *Element but not generated with DropdownItem or DropdownHref or DropdownDivider
//   - it is a gomponents.Node with type gomponents.ElementType
func DropdownMenu(children ...any) *Element {
	menu := Elem(html.Div).
		With(Class("dropdown-menu")).
		With(html.Role("menu"))

	content := Elem(html.Div).
		With(Class("dropdown-content"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("dropdown-item") || c.hasClass("dropdown-divider") {
				content.With(c)
			} else {
				content.With(DropdownItem(c))
			}
		case gomponents.Node:
			if IsAttribute(c) {
				menu.With(c)
			} else {
				content.With(c)
			}
		default:
			menu.With(c)
		}
	}

	return menu.With(content)
}
