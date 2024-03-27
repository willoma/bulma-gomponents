package easy

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

type easyDropdown struct {
	hover            bool
	up               bool
	buttonText       string
	dropdownChildren []any
	menuChildren     []any
}

func (d *easyDropdown) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case *b.ApplyToInner:
			d.menuChildren = append(d.menuChildren, c.Child)
		case *b.ApplyToOuter:
			d.dropdownChildren = append(d.dropdownChildren, c.Child)
		case b.Class, b.ColorClass, b.ExternalClass, b.ExternalClassesAndStyles, b.MultiClass, b.Styles:
			d.dropdownChildren = append(d.dropdownChildren, c)
		case gomponents.Node:
			if b.IsAttribute(c) {
				d.dropdownChildren = append(d.dropdownChildren, c)
			} else {
				d.menuChildren = append(d.menuChildren, c)
			}
		case []any:
			d.With(c...)
		default:
			d.menuChildren = append(d.menuChildren, c)
		}
	}

	return d
}

func (d *easyDropdown) Render(w io.Writer) error {
	var (
		dropdown   func(triggerButton, menu b.Element, children ...any) b.Element
		faIconName string
	)
	if d.up {
		dropdown = b.Dropup
		faIconName = "angle-up"
	} else {
		dropdown = b.Dropdown
		faIconName = "angle-down"
	}

	button := b.Button(
		html.Aria("haspopup", "true"),
		d.buttonText,
		fa.Icon(fa.Solid, faIconName, b.Small, html.Aria("hidden", "true")),
	)

	elem := dropdown(button, b.DropdownMenu(d.menuChildren...))

	if d.hover {
		elem.With(b.Hoverable)
	} else {
		elem.With(b.OnClick(b.JSToggleMe))
		button.With(b.On("blur", b.JSCloseThisDropdown))
	}

	return elem.With(d.dropdownChildren...).Render(w)
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
func ClickableDropdown(buttonText string, children ...any) b.Element {
	return (&easyDropdown{buttonText: buttonText}).With(children...)
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
func HoverableDropdown(buttonText string, children ...any) b.Element {
	return (&easyDropdown{hover: true, buttonText: buttonText}).With(children...)
}

// ClickableDropup creates a dropdown which opens the menu to the top when
// clicking on the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func ClickableDropup(buttonText string, children ...any) b.Element {
	return (&easyDropdown{up: true, buttonText: buttonText}).With(children...)
}

// HoverableDropup creates a dropdown which opens the menu to the top when
// hovering the button with the given text.
//   - when a child is a b.Class, b.ColorClass or b.Styles,  it is added to the
//     menu element
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added to the dropdown element
//   - when a child is a gomponents.Node with another type, it is added to the
//     menu element
//   - other children types are added to the menu element
func HoverableDropup(buttonText string, children ...any) b.Element {
	return (&easyDropdown{hover: true, up: true, buttonText: buttonText}).With(children...)
}
