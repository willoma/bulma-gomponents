package bulma

import (
	"io"
	"sync"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"
)

// Navbar creates a navbar.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func Navbar(children ...any) e.Element {
	n := &navbar{navbar: e.Nav(e.Class("navbar"))}
	n.With(children...)
	return n
}

type navbar struct {
	navbar                e.Element
	brand                 e.Element
	start                 e.Element
	end                   e.Element
	intermediateContainer e.Element

	hasMenu  bool
	fixed    e.Class
	rendered sync.Once
}

func (n *navbar) ModifyParent(parent e.Element) {
	switch n.fixed {
	case FixedTop:
		parent.With(NavbarFixedTop)
	case FixedBottom:
		parent.With(NavbarFixedBottom)
	}
}

func (n *navbar) addToBrand(children ...any) {
	if n.brand == nil {
		n.brand = e.Div(e.Class("navbar-brand"))
	}
	n.brand.With(children...)
}

func (n *navbar) addToStart(children ...any) {
	if n.start == nil {
		n.hasMenu = true
		n.start = e.Div(e.Class("navbar-start"))
	}
	n.start.With(children...)
}

func (n *navbar) addToEnd(children ...any) {
	if n.end == nil {
		n.hasMenu = true
		n.end = e.Div(e.Class("navbar-end"))
	}
	n.end.With(children...)
}

func (n *navbar) With(children ...any) e.Element {
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
		case e.Class:
			switch c {
			case FixedTop:
				n.fixed = FixedTop
			case FixedBottom:
				n.fixed = FixedBottom
			}
			n.navbar.With(c)
		case []any:
			n.With(c...)
		default:
			n.navbar.With(c)
		}
	}

	return n
}

func (n *navbar) Render(w io.Writer) error {
	n.rendered.Do(func() {
		var target e.Element
		if n.intermediateContainer != nil {
			target = n.intermediateContainer
			n.navbar.With(target)
		} else {
			target = n.navbar
		}

		if n.brand != nil {
			target.With(n.brand)
		}

		if n.hasMenu {
			n.addToBrand(
				e.A(
					e.AriaButton,
					e.Class("navbar-burger"),
					e.AriaLabel("menu"),
					e.AriaExpandedFalse,
					e.Span(e.AriaHiddenTrue),
					e.Span(e.AriaHiddenTrue),
					e.Span(e.AriaHiddenTrue),
					e.Span(e.AriaHiddenTrue),
					e.OnClick("this.classList.toggle('is-active');this.closest('.navbar').getElementsByClassName('navbar-menu')[0].classList.toggle('is-active')"),
				),
			)

			menu := e.Div(e.Class("navbar-menu"))

			if n.start != nil {
				menu.With(n.start)
			}

			if n.end != nil {
				menu.With(n.end)
			}

			target.With(menu)
		}
	})

	return n.navbar.Render(w)
}

func (n *navbar) Clone() e.Element {
	return &navbar{
		navbar:                n.navbar.Clone(),
		brand:                 n.brand.Clone(),
		start:                 n.start.Clone(),
		end:                   n.end.Clone(),
		intermediateContainer: n.intermediateContainer.Clone(),

		hasMenu: n.hasMenu,
		fixed:   n.fixed,
	}
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
func NavbarItem(children ...any) e.Element {
	return e.Div(e.Class("navbar-item"), children)
}

// NavbarAHref creates a link item to add to a navbar brand, start or end
// section, or to a NavbarDropdown.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarAHref(href string, children ...any) e.Element {
	return e.A(e.Class("navbar-item"), html.Href(href), children)
}

// NavbarDropdown creates a nav bar item which includes a dropdown menu.
func NavbarDropdown(children ...any) e.Element {
	link := NavbarLink()
	dropdown := e.Div(e.Class("navbar-dropdown"))

	n := &navbarDropdown{
		Element:  NavbarItem(e.Class("has-dropdown"), link, dropdown),
		link:     link,
		dropdown: dropdown,
	}
	n.With(children...)
	return n
}

type navbarDropdown struct {
	e.Element
	link     e.Element
	dropdown e.Element
}

func (n *navbarDropdown) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onItem:
			n.Element.With(c...)
		case onLink:
			n.link.With(c...)
		case onDropdown:
			n.dropdown.With(c...)
		case e.Class:
			switch c {
			case Hoverable:
				n.Element.With(c)
			case Clickable:
				n.link.With(e.OnClick(JSToggleThisNavbarItem))
			case Up:
				n.Element.With(e.Class("has-dropdown-up"))
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

func (n *navbarDropdown) Clone() e.Element {
	return &navbarDropdown{
		Element:  n.Element.Clone(),
		link:     n.link.Clone(),
		dropdown: n.dropdown.Clone(),
	}
}

// NavbarLink creates a link element, to include in a NavbarItem with the
// HasDropdown modifier.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarLink(children ...any) e.Element {
	return e.A(e.Class("navbar-link"), children)
}

// NavbarDivider creates a divider element, to include in a NavbarDropdown.
//
// https://willoma.github.io/bulma-gomponents/navbar.html
func NavbarDivider() e.Element {
	return e.Hr(e.Class("navbar-divider"))
}
