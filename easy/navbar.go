package easy

import (
	b "github.com/willoma/bulma-gomponents"
)

type navbarDropdownOption int

const (
	NavbarDropdownHoverable navbarDropdownOption = 1 << iota
	NavbarDropdownClickable
	NavbarDropup
	NavbarDropdownArrowless
	NavbarDropdownBoxed
	NavbarDropdownActive
)

func NavbarDropdown(label string, children ...any) *b.Element {
	item := b.NavbarItem(b.HasDropdown)

	link := b.NavbarLink(label)

	dropdown := b.NavbarDropdown()

	for _, c := range children {
		switch c := c.(type) {
		case navbarDropdownOption:
			switch c {
			case NavbarDropdownHoverable:
				item.With(b.Hoverable)
			case NavbarDropdownClickable:
				link.With(b.OnClick(b.JSToggleThisNavbarItem))
			case NavbarDropup:
				item.With(b.Class("has-dropdown-up"))
			case NavbarDropdownArrowless:
				link.With(b.Arrowless)
			case NavbarDropdownBoxed:
				dropdown.With(b.Boxed)
			case NavbarDropdownActive:
				item.With(b.Active)
			}
		default:
			dropdown.With(c)
		}
	}

	return item.
		With(link).
		With(dropdown)
}
