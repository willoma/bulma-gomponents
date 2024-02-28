package bulma

import "github.com/maragudk/gomponents/html"

type container *Element

// Container creates a container element.
func Container(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			Withs(children),
	)
}

// ContainerWidescreen creates a container element which is fullwidth up to the
// widescreen breakpoint.
func ContainerWidescreen(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			With(Class("is-widescreen")).
			Withs(children),
	)
}

// ContainerFullHD creates a container element which is fullwidth up to the
// full HD breakpoint.
func ContainerFullHD(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			With(Class("is-fullhd")).
			Withs(children),
	)
}

// ContainerMaxDesktop creates a container element which has a max width of
// the desktop breakpoint.
func ContainerMaxDesktop(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			With(Class("is-max-desktop")).
			Withs(children),
	)
}

// ContainerMaxWidescreen creates a container element which has a max width of
// the widescreen breakpoint.
func ContainerMaxWidescreen(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			With(Class("is-max-widescreen")).
			Withs(children),
	)
}

// ContainerFluid creates a container element without a maximum width, but which
// keeps the left and right margins.
func ContainerFluid(children ...any) container {
	return container(
		Elem(html.Div).
			With(Class("container")).
			With(Class("is-fluid")).
			Withs(children),
	)
}
