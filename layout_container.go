package bulma

import "github.com/maragudk/gomponents/html"

type container *Element

// Container creates a container element.
func Container(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), children),
	)
}

// ContainerWidescreen creates a container element which is fullwidth up to the
// widescreen breakpoint.
func ContainerWidescreen(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), Class("is-widescreen"), children),
	)
}

// ContainerFullHD creates a container element which is fullwidth up to the
// full HD breakpoint.
func ContainerFullHD(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), Class("is-fullhd"), children),
	)
}

// ContainerMaxDesktop creates a container element which has a max width of
// the desktop breakpoint.
func ContainerMaxDesktop(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), Class("is-max-desktop"), children),
	)
}

// ContainerMaxWidescreen creates a container element which has a max width of
// the widescreen breakpoint.
func ContainerMaxWidescreen(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), Class("is-max-widescreen"), children),
	)
}

// ContainerFluid creates a container element without a maximum width, but which
// keeps the left and right margins.
func ContainerFluid(children ...any) container {
	return container(
		Elem(html.Div, Class("container"), Class("is-fluid"), children),
	)
}
