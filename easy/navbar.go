package easy

import (
	"io"

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

func NavbarDropdown(label string, children ...any) b.Element {
	return (&navbarDropdown{label: label}).With(children...)
}

type navbarDropdown struct {
	label            string
	dropdownChildren []any
	itemChildren     []any
	linkChildren     []any
}

func (n *navbarDropdown) With(children ...any) b.Element {
	for _, c := range children {
		switch c := c.(type) {
		case navbarDropdownOption:
			switch c {
			case NavbarDropdownHoverable:
				n.itemChildren = append(n.itemChildren, b.Hoverable)
			case NavbarDropdownClickable:
				n.linkChildren = append(n.linkChildren, b.OnClick(b.JSToggleThisNavbarItem))
			case NavbarDropup:
				n.itemChildren = append(n.itemChildren, b.Class("has-dropdown-up"))
			case NavbarDropdownArrowless:
				n.linkChildren = append(n.linkChildren, b.Arrowless)
			case NavbarDropdownBoxed:
				n.dropdownChildren = append(n.dropdownChildren, b.Boxed)
			case NavbarDropdownActive:
				n.itemChildren = append(n.itemChildren, b.Active)
			}
		case []any:
			n.With(c...)
		default:
			n.dropdownChildren = append(n.dropdownChildren, c)
		}
	}

	return n
}

func (n *navbarDropdown) Render(w io.Writer) error {
	return b.NavbarItem(
		b.HasDropdown,
		n.itemChildren,
		b.NavbarLink(n.label, n.linkChildren),
		b.NavbarDropdown(n.dropdownChildren...),
	).Render(w)
}
