package bulma

import "github.com/maragudk/gomponents/html"

// Tabs create a tabs container.
//
// Its arguments must include TabsLink. The ul element is automatically created.
//
// If a child is the Container function or one of its derivative (Container*),
// this function is executed and its result is used as an intermediate container.
//
// The following modifiers change the tabs behaviour:
//   - Centered: center the tabs
//   - Right: align the tabs to the right
//   - Boxed: draw boxed tabs
//   - Toggle: button-looking tabs
//   - ToggleRounded: rounded button-looking tabs
//   - FullWidth: take the whole width
//
// The following modifiers change the tabs size:
//   - Small
//   - Medium
//   - Large
func Tabs(children ...any) *Element {
	tabs := Elem(html.Div).
		With(Class("tabs"))
	content := Elem(html.Ul)

	var intermediateContainer *Element

	for _, c := range children {
		switch c := c.(type) {
		case Class:
			tabs.With(c)
		case func(children ...any) container:
			intermediateContainer = c()
		default:
			content.With(c)
		}
	}

	if intermediateContainer != nil {
		return tabs.With(intermediateContainer.With(content))
	}
	return tabs.With(content)
}

// TabsLink creates a tab entry which is a link. Use html.Href as an argument
// to define a link target if needed.
func TabsLink(children ...any) *Element {
	li := Elem(html.Li)

	a := Elem(html.A)
	a.spanAroundNonIconsIfHasIcons = true

	for _, c := range children {
		switch c := c.(type) {
		case Class:
			if c == Active {
				li.With(Active)
			} else {
				a.With(c)
			}
		default:
			a.With(c)
		}
	}

	return li.With(a)
}
