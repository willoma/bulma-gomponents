package easy

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

func simpleDropdown(hover, up bool, buttonText string, children []any) *b.Element {
	var (
		dropdownChildren []any
		menuChildren     []any
		dropdown         func(triggerButton, menu *b.Element, children ...any) *b.Element
		faIconName       string
	)
	if up {
		dropdown = b.Dropup
		faIconName = "angle-up"
	} else {
		dropdown = b.Dropdown
		faIconName = "angle-down"
	}

	button := b.Button(
		html.Aria("haspopup", "true"),
		buttonText,
		fa.Icon(fa.Solid, faIconName, b.Small, html.Aria("hidden", "true")),
	)

	if hover {
		dropdownChildren = append(dropdownChildren, b.Hoverable)
	} else {
		dropdownChildren = append(
			dropdownChildren,
			b.OnClick(b.JSToggleMe),
		)
		button.With(
			b.On("blur", b.JSCloseThisDropdown),
		)
	}

	for _, c := range children {
		switch c := c.(type) {
		case b.Class, b.ColorClass, b.Styles:
			dropdownChildren = append(dropdownChildren, c)
		case gomponents.Node:
			if b.IsAttribute(c) {
				dropdownChildren = append(dropdownChildren, c)
			} else {
				menuChildren = append(menuChildren, c)
			}
		default:
			menuChildren = append(menuChildren, c)
		}
	}

	return dropdown(button, b.DropdownMenu(menuChildren...), dropdownChildren...)
}

// ClickableDropdown creates a dropdown which opens the menu when clicking on
// the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func ClickableDropdown(buttonText string, children ...any) *b.Element {
	return simpleDropdown(false, false, buttonText, children)
}

// HoverableDropdown creates a dropdown which opens the menu when hovering
// the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func HoverableDropdown(buttonText string, children ...any) *b.Element {
	return simpleDropdown(true, false, buttonText, children)
}

// ClickableDropdown creates a dropdown which opens the menu to the top when
// clicking on the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func ClickableDropup(buttonText string, children ...any) *b.Element {
	return simpleDropdown(false, true, buttonText, children)
}

// ClickableDropdown creates a dropdown which opens the menu to the top when
// hovering the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func HoverableDropup(buttonText string, children ...any) *b.Element {
	return simpleDropdown(true, true, buttonText, children)
}
