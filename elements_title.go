package bulma

import (
	e "github.com/willoma/gomplements"
)

// Title creates a h1 title.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title(children ...any) e.Element {
	return e.H1(e.Class("title"), children)
}

// Subtitle creates a h2 subtitle.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle(children ...any) e.Element {
	return e.H2(e.Class("subtitle"), children)
}

// Title1 creates a h1 title of size 1.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title1(children ...any) e.Element {
	return e.H1(e.Class("title"), e.Class("is-1"), children)
}

// Title2 creates a h2 title of size 2.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title2(children ...any) e.Element {
	return e.H2(e.Class("title"), e.Class("is-2"), children)
}

// Title3 creates a h3 title of size 3.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title3(children ...any) e.Element {
	return e.H3(e.Class("title"), e.Class("is-3"), children)
}

// Title4 creates a h4 title of size 4.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title4(children ...any) e.Element {
	return e.H4(e.Class("title"), e.Class("is-4"), children)
}

// Title5 creates a h5 title of size 5.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title5(children ...any) e.Element {
	return e.H5(e.Class("title"), e.Class("is-5"), children)
}

// Title6 creates a h6 title of size 6.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Title6(children ...any) e.Element {
	return e.H6(e.Class("title"), e.Class("is-6"), children)
}

// Subitle1 creates a h1 subtitle of size 1.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle1(children ...any) e.Element {
	return e.H1(e.Class("subtitle"), e.Class("is-1"), children)
}

// Subitle2 creates a h2 subtitle of size 2.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle2(children ...any) e.Element {
	return e.H2(e.Class("subtitle"), e.Class("is-2"), children)
}

// Subitle3 creates a h3 subtitle of size 3.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle3(children ...any) e.Element {
	return e.H3(e.Class("subtitle"), e.Class("is-3"), children)
}

// Subitle4 creates a h4 subtitle of size 4.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle4(children ...any) e.Element {
	return e.H4(e.Class("subtitle"), e.Class("is-4"), children)
}

// Subitle5 creates a h5 subtitle of size 5.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle5(children ...any) e.Element {
	return e.H5(e.Class("subtitle"), e.Class("is-5"), children)
}

// Subitle6 creates a h6 subtitle of size 6.
//
// https://willoma.github.io/bulma-gomponents/title.html
func Subtitle6(children ...any) e.Element {
	return e.H6(e.Class("subtitle"), e.Class("is-6"), children)
}
