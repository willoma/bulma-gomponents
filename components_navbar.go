package bulma

import (
	"io"

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
// The FixedTop modifier fixes the navbar to the top of the window; when it is
// used, the NavbarFixedTop modifier must be added to the body. It is
// automatically done when using NavbarTop.
//
// The FixedBottom modifier fixes the navbar to the bottom of the window; when
// it is used, the NavbarFixedBottom modifier must be added to the body. It is
// automatically done when using NavbarBottom.
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
func Navbar(children ...any) Element {
	return new(navbar).With(children...)
}

type navbar struct {
	intermediateContainer *container
	brandChildren         []any
	startChildren         []any
	endChildren           []any
	navbarChildren        []any

	fixed Class
}

func (n *navbar) ModifyParent(parent Element) {
	switch n.fixed {
	case FixedTop:
		parent.With(NavbarFixedTop)
	case FixedBottom:
		parent.With(NavbarFixedBottom)
	}
}

func (n *navbar) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case navbarBrand:
			n.brandChildren = append(n.brandChildren, c...)
		case navbarStart:
			n.startChildren = append(n.startChildren, c...)
		case navbarEnd:
			n.endChildren = append(n.endChildren, c...)
		case *container:
			n.intermediateContainer = c
		case Class:
			switch c {
			case FixedTop:
				n.fixed = FixedTop
			case FixedBottom:
				n.fixed = FixedBottom
			}
			n.navbarChildren = append(n.navbarChildren, c)
		case []any:
			n.With(c...)
		default:
			n.navbarChildren = append(n.navbarChildren, c)
		}
	}

	return n
}

func (n *navbar) Render(w io.Writer) error {
	navbar := Elem(html.Nav, Class("navbar"), n.navbarChildren)

	var target Element
	if n.intermediateContainer != nil {
		target = n.intermediateContainer
		navbar.With(target)
	} else {
		target = navbar
	}

	if len(n.brandChildren) > 0 || len(n.startChildren) > 0 || len(n.endChildren) > 0 {
		brandElem := Elem(html.Div, Class("navbar-brand"))

		if len(n.brandChildren) > 0 {
			brandElem.With(n.brandChildren...)
		}

		if len(n.startChildren) > 0 || len(n.endChildren) > 0 {
			brandElem.With(
				Elem(
					html.A,
					html.Role("button"),
					Class("navbar-burger"),
					html.Aria("label", "menu"),
					html.Aria("expanded", "false"),
					Elem(html.Span, html.Aria("hidden", "true")),
					Elem(html.Span, html.Aria("hidden", "true")),
					Elem(html.Span, html.Aria("hidden", "true")),
					Elem(html.Span, html.Aria("hidden", "true")),
					OnClick("this.classList.toggle('is-active');this.closest('.navbar').getElementsByClassName('navbar-menu')[0].classList.toggle('is-active')"),
				),
			)
		}

		target.With(brandElem)
	}

	if len(n.startChildren) > 0 || len(n.endChildren) > 0 {
		menu := Elem(html.Div, Class("navbar-menu"))

		if len(n.startChildren) > 0 {
			menu.With(
				Elem(html.Div, Class("navbar-start"), n.startChildren),
			)
		}

		if len(n.endChildren) > 0 {
			menu.With(
				Elem(html.Div, Class("navbar-end"), n.endChildren),
			)
		}

		target.With(menu)
	}

	return navbar.Render(w)
}

// NavbarBrand designates children to be part of the navbar brand section.
//
// It cannot be used as a node by itself.
func NavbarBrand(children ...any) navbarBrand {
	return navbarBrand(children)
}

type navbarBrand []any

// NavbarStart designates children to be part of the navbar start section.
//
// It cannot be used as a node by itself.
func NavbarStart(children ...any) navbarStart {
	return navbarStart(children)
}

type navbarStart []any

// NavbarEnd designates children to be part of the navbar end section.
//
// It cannot be used as a node by itself.
func NavbarEnd(children ...any) navbarEnd {
	return navbarEnd(children)
}

type navbarEnd []any

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
func NavbarItem(children ...any) Element {
	return Elem(html.Div, Class("navbar-item"), children)
}

// NavbarAHref creates a link item to add to a navbar brand, start or end
// section, or to a NavbarDropdown or easy.NavbarDropdown.
//
// The following modifiers change the link item behaviour:
//   - Expanded: turn the item into a full-width element
//   - Tab: add a bottom border on hover, always show the bottom border when
//     adding Active
func NavbarAHref(href string, children ...any) Element {
	return Elem(html.A, Class("navbar-item"), html.Href(href), children)
}

// NavbarDropdown creates a nav bar item which includes a dropdown menu.
func NavbarDropdown(children ...any) Element {
	link := NavbarLink()
	dropdown := Elem(html.Div, Class("navbar-dropdown"))

	n := &navbarDropdown{
		Element:  NavbarItem(Class("has-dropdown"), link, dropdown),
		link:     link,
		dropdown: dropdown,
	}
	n.With(children...)
	return n
}

type navbarDropdown struct {
	Element
	link     Element
	dropdown Element
}

func (n *navbarDropdown) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onItem:
			n.Element.With(c...)
		case onLink:
			n.link.With(c...)
		case onDropdown:
			n.dropdown.With(c...)
		case Class:
			switch c {
			case Hoverable:
				n.Element.With(c)
			case Clickable:
				n.link.With(OnClick(JSToggleThisNavbarItem))
			case Up:
				n.Element.With(Class("has-dropdown-up"))
			case Arrowless:
				n.link.With(c)
			case Active:
				n.Element.With(c)
			default:
				n.dropdown.With(c)
			}
		case string:
			n.link.With(c)
		case []any:
			n.With(c...)
		default:
			n.dropdown.With(c)
		}
	}
	return n
}

// NavbarLink creates a link element, to include in a NavbarItem with the
// HasDropdown modifier.
//
// The following modifiers change the link behaviour:
//   - Arrowless: remove the arrow in the item
//
// Creating a navbar dropdown is made easier with `easy.NavbarDropdown`, in
// which case this function is not needed.
func NavbarLink(children ...any) Element {
	return Elem(html.A, Class("navbar-link"), children)
}

// NavbarDivider creates a divider element, to include in a NavbarDropdown or
// an easy.NavbarDropdown.
func NavbarDivider() Element {
	return Elem(html.Hr, Class("navbar-divider"))
}
