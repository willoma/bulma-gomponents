package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents/html"
)

// Navbar creates a navbar.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func Navbar(children ...any) Element {
	n := &navbar{Element: Elem(html.Nav, Class("navbar"))}
	n.With(children...)
	return n
}

type navbar struct {
	Element
	brand                 Element
	start                 Element
	end                   Element
	intermediateContainer Element

	hasMenu  bool
	fixed    Class
	rendered sync.Once
}

func (n *navbar) ModifyParent(parent Element) {
	switch n.fixed {
	case FixedTop:
		parent.With(NavbarFixedTop)
	case FixedBottom:
		parent.With(NavbarFixedBottom)
	}
}

func (n *navbar) addToBrand(children ...any) {
	if n.brand == nil {
		n.brand = Elem(html.Div, Class("navbar-brand"))
	}
	n.brand.With(children...)
}

func (n *navbar) addToStart(children ...any) {
	if n.start == nil {
		n.hasMenu = true
		n.start = Elem(html.Div, Class("navbar-start"))
	}
	n.start.With(children...)
}

func (n *navbar) addToEnd(children ...any) {
	if n.end == nil {
		n.hasMenu = true
		n.end = Elem(html.Div, Class("navbar-end"))
	}
	n.end.With(children...)
}

func (n *navbar) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case navbarBrand:
			n.addToBrand(c...)
		case navbarStart:
			n.addToStart(c...)
		case navbarEnd:
			n.addToEnd(c...)
		case *container:
			n.intermediateContainer = c
		case Class:
			switch c {
			case FixedTop:
				n.fixed = FixedTop
			case FixedBottom:
				n.fixed = FixedBottom
			}
			n.Element.With(c)
		case []any:
			n.With(c...)
		default:
			n.Element.With(c)
		}
	}

	return n
}

func (n *navbar) Render(w io.Writer) error {
	n.rendered.Do(func() {
		var target Element
		if n.intermediateContainer != nil {
			target = n.intermediateContainer
			n.Element.With(target)
		} else {
			target = n.Element
		}

		if n.brand != nil {
			target.With(n.brand)
		}

		if n.hasMenu {
			n.addToBrand(
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

			menu := Elem(html.Div, Class("navbar-menu"))

			if n.start != nil {
				menu.With(n.start)
			}

			if n.end != nil {
				menu.With(n.end)
			}

			target.With(menu)
		}
	})

	return n.Element.Render(w)
}

// NavbarBrand designates children to be part of the navbar brand section.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarBrand(children ...any) navbarBrand {
	return navbarBrand(children)
}

type navbarBrand []any

// NavbarStart designates children to be part of the navbar start section.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarStart(children ...any) navbarStart {
	return navbarStart(children)
}

type navbarStart []any

// NavbarEnd designates children to be part of the navbar end section.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarEnd(children ...any) navbarEnd {
	return navbarEnd(children)
}

type navbarEnd []any

// NavbarItem creates an item to add to a navbar brand, start or end section, or
// to a NavbarDropdown.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarItem(children ...any) Element {
	return Elem(html.Div, Class("navbar-item"), children)
}

// NavbarAHref creates a link item to add to a navbar brand, start or end
// section, or to a NavbarDropdown.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
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
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarLink(children ...any) Element {
	return Elem(html.A, Class("navbar-link"), children)
}

// NavbarDivider creates a divider element, to include in a NavbarDropdown.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarDivider() Element {
	return Elem(html.Hr, Class("navbar-divider"))
}
