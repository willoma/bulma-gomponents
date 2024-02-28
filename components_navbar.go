package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Navbar creates a navbar.
//   - use NavbarBrand to declare children placed in the brand subsection
//   - use NavbarStart to declare children placed in the start subsection
//   - use NavbarEnd to declare children placed in the end section
//   - if a child is the Container function or one of its derivative
//     (Container*), this function is executed and its result is used as an
//     intermediate container
//
// The brand section is added only if at least one brand, start or end child is
// provided. The section part and the burger button are added only if at least
// one start or end child is provided. The click event on the burger, which
// toggles the navbar menu on mobile devices, is already included.
//
// The FixedTop modifier change the separator fixes the navbar to the top of
// the window; when it is used, the NavbarFixedTop modifier must be added to the
// body. It is automatically done when using NavbarTop.
//
// The FixedBottom modifier change the separator fixes the navbar to the bottom
// of the window; when it is used, the NavbarFixedBottom modifier must be added
// to the body. It is automatically done when using NavbarBottom.
//
// The following modifiers change the navbar behaviour:
//   - Transparent: remove hover and active backgrounds from the items
//   - Spaced: increase padding
//   - Shadow: add a small shadow around the navbar
//
// The following modifiers change the navbar color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - Black
//   - Dark
//   - Light
//   - White
func Navbar(children ...any) *Element {
	navbar := Elem(html.Nav).
		With(Class("navbar"))

	var brand, start, end []any

	target := navbar

	for _, c := range children {
		switch c := c.(type) {
		case navbarBrand:
			brand = append(brand, c...)
		case navbarStart:
			start = append(start, c...)
		case navbarEnd:
			end = append(end, c...)
		case func(children ...any) container:
			target = c()
			navbar.With(target)
		default:
			navbar.With(c)
		}
	}

	if len(brand) > 0 || len(start) > 0 || len(end) > 0 {
		brandElem := Elem(html.Div).
			With(Class("navbar-brand"))

		if len(brand) > 0 {
			brandElem.Withs(brand)
		}

		if len(start) > 0 || len(end) > 0 {
			brandElem.With(
				Elem(html.A).
					With(html.Role("button")).
					With(Class("navbar-burger")).
					With(html.Aria("label", "menu")).
					With(html.Aria("expanded", "false")).
					With(Elem(html.Span).With(html.Aria("hidden", "true"))).
					With(Elem(html.Span).With(html.Aria("hidden", "true"))).
					With(Elem(html.Span).With(html.Aria("hidden", "true"))).
					With(OnClick(`this.classList.toggle("is-active");this.closest(".navbar").getElementsByClassName("navbar-menu")[0].classList.toggle("is-active")`)),
			)
		}

		target.With(brandElem)
	}

	if len(start) > 0 || len(end) > 0 {
		menu := Elem(html.Div).
			With(Class("navbar-menu"))

		if len(start) > 0 {
			menu.With(
				Elem(html.Div).
					With(Class("navbar-start")).
					Withs(start),
			)
		}

		if len(end) > 0 {
			menu.With(
				Elem(html.Div).
					With(Class("navbar-end")).
					Withs(end),
			)
		}

		target.With(menu)
	}

	return navbar
}

// TopNavbar, which must be used as an argument to HTML (which means it will be
// added to the body), adds a fixed-to-top navbar. It automatically adds the
// "has-navbar-fixed-top" class to the body.
//
// Please refer to the Navbar documentation for description of its behaviour.
func TopNavbar(children ...any) *Element {
	return Navbar(children...).
		With(FixedTop)
}

// BottomNavbar, which must be used as an argument to HTML (which means it will
// be added to the body), adds a fixed-to-bottom navbar. It automatically adds
// the "has-navbar-fixed-bottom" class to the body.
//
// Please refer to the Navbar documentation for description of its behaviour.
func BottomNavbar(children ...any) *Element {
	return Navbar(children...).
		With(FixedBottom)
}

type navbarBrand []any

// NavbarBrand designates children to be part of the navbar brand section.
func NavbarBrand(children ...any) navbarBrand {
	return navbarBrand(children)
}

type navbarStart []any

// NavbarStart designates children to be part of the navbar start section.
func NavbarStart(children ...any) navbarStart {
	return navbarStart(children)
}

type navbarEnd []any

// NavbarEnd designates children to be part of the navbar end section.
func NavbarEnd(children ...any) navbarEnd {
	return navbarEnd(children)
}

// NavbarItem creates an item to add to a navbar brand, start or end section, or
// to a NavbarDropdown or easy.NavbarDropdown.
//
// The following modifiers change the item behaviour:
//   - Expanded: turn the item into a full-width element
//   - Tab: add a bottom border on hover, always show the bottom border when
//     adding Active
//
// The following modifiers are used for dropdowns:
//   - HasDropdown: makes the item a link+dropdown container
//   - HasDropup: makes the item a dropdown container, with the dropdown
//     opening above the link (this modifier automatically applies the
//     "has-dropdown" class along with the "has-dropdown-up" class, therefore
//     HasDropdown is not needed in conjuction with HasDropup)
//   - Hoverable: makes the included dropdown automatically show on hover
//   - Active: force the dropdown to be open
func NavbarItem(children ...any) *Element {
	return Elem(html.Div).
		With(Class("navbar-item")).
		Withs(children)
}

// NavbarAHref creates a link item to add to a navbar brand, start or end
// section, or to a NavbarDropdown or easy.NavbarDropdown.
//
// The following modifiers change the link item behaviour:
//   - Expanded: turn the item into a full-width element
//   - Tab: add a bottom border on hover, always show the bottom border when
//     adding Active
func NavbarAHref(href string, children ...any) *Element {
	return Elem(html.A).
		With(Class("navbar-item")).
		With(html.Href(href)).
		Withs(children)
}

// NavbarDropdown creates a dropdown element, to include in a NavbarItem with
// the HasDropdown modifier.
//
// The following modifiers change the dropdown behaviour:
//   - Boxed: show the dropdown as a boxed element, without the top grey border
//   - Right: used when a dropdown is on the right side (end section), in order
//     to open its content aligned to the right of the link
//
// Creating a navbar dropdown is made easier with `easy.NavbarDropdown`, in
// which case this function is not needed.
func NavbarDropdown(children ...any) *Element {
	return Elem(html.Div).
		With(Class("navbar-dropdown")).
		Withs(children)
}

// NavbarLink creates a link element, to include in a NavbarItem with the
// HasDropdown modifier.
//
// The following modifiers change the link behaviour:
//   - Arrowless: remove the arrow in the item
//
// Creating a navbar dropdown is made easier with `easy.NavbarDropdown`, in
// which case this function is not needed.
func NavbarLink(children ...any) *Element {
	return Elem(html.A).
		With(Class("navbar-link")).
		Withs(children)
}

// NavbarDivider creates a divider element, to include in a NavbarDropdown or
// an easy.NavbarDropdown.
func NavbarDivider() *Element {
	return Elem(html.Hr).
		With(Class("navbar-divider"))
}
