package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// Title creates a h1 title
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title(children ...any) *Element {
	return Elem(html.H1).
		With(Class("title")).
		Withs(children)
}

// Subtitle creates a h2 subtitle
func Subtitle(children ...any) *Element {
	return Elem(html.H2).
		With(Class("subtitle")).
		Withs(children)
}

// Title1 creates a h1 title of size 1
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title1(children ...any) *Element {
	return Elem(html.H1).
		With(Class("title")).
		With(Class("is-1")).
		Withs(children)
}

// Title2 creates a h2 title of size 2
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title2(children ...any) *Element {
	return Elem(html.H2).
		With(Class("title")).
		With(Class("is-2")).
		Withs(children)
}

// Title3 creates a h3 title of size 3
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title3(children ...any) *Element {
	return Elem(html.H3).
		With(Class("title")).
		With(Class("is-3")).
		Withs(children)
}

// Title4 creates a h4 title of size 4
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title4(children ...any) *Element {
	return Elem(html.H4).
		With(Class("title")).
		With(Class("is-4")).
		Withs(children)
}

// Title5 creates a h5 title of size 5
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title5(children ...any) *Element {
	return Elem(html.H5).
		With(Class("title")).
		With(Class("is-5")).
		Withs(children)
}

// Title6 creates a h6 title of size 6
//
// The following modifiers change the title behaviour:
//   - Spaced: maintain normal spacing between title and next subtitle
func Title6(children ...any) *Element {
	return Elem(html.H6).
		With(Class("title")).
		With(Class("is-6")).
		Withs(children)
}

// Subitle1 creates a h1 subtitle of size 1
func Subtitle1(children ...any) *Element {
	return Elem(html.H1).
		With(Class("subtitle")).
		With(Class("is-1")).
		Withs(children)
}

// Subitle2 creates a h2 subtitle of size 2
func Subtitle2(children ...any) *Element {
	return Elem(html.H2).
		With(Class("subtitle")).
		With(Class("is-2")).
		Withs(children)
}

// Subitle3 creates a h3 subtitle of size 3
func Subtitle3(children ...any) *Element {
	return Elem(html.H3).
		With(Class("subtitle")).
		With(Class("is-3")).
		Withs(children)
}

// Subitle4 creates a h4 subtitle of size 4
func Subtitle4(children ...any) *Element {
	return Elem(html.H4).
		With(Class("subtitle")).
		With(Class("is-4")).
		Withs(children)
}

// Subitle5 creates a h5 subtitle of size 5
func Subtitle5(children ...any) *Element {
	return Elem(html.H5).
		With(Class("subtitle")).
		With(Class("is-5")).
		Withs(children)
}

// Subitle6 creates a h6 subtitle of size 6
func Subtitle6(children ...any) *Element {
	return Elem(html.H6).
		With(Class("subtitle")).
		With(Class("is-6")).
		Withs(children)
}
