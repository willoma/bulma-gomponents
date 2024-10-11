package bulma

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"
)

// ImgAlt is a modifier that, when applied to ImageImg, adds an alt attribute.
type ImgAlt string

// Image creates a figure element with class "image".
//
// https://willoma.github.io/bulma-gomponents/image.html
func Image(children ...any) e.Element {
	return e.Figure(e.Class("image"), children)
}

// ImageImg creates a figure element with class "image" and the inner img
// element, with the provided src.
//
// https://willoma.github.io/bulma-gomponents/image.html
func ImageImg(src string, children ...any) e.Element {
	img := e.Img(html.Src(src))
	i := &imageImg{
		Element: e.Figure(e.Class("image"), img),
		img:     img,
	}
	i.With(children...)
	return i
}

type imageImg struct {
	e.Element
	img e.Element
}

func (i *imageImg) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onImg:
			i.img.With(c...)
		case onFigure:
			i.Element.With(c...)
		case e.Class:
			if c == Rounded {
				i.img.With(c)
			} else {
				i.Element.With(c)
			}
		case ImgAlt:
			i.img.With(html.Alt(string(c)))
		case []any:
			i.With(c...)
		default:
			i.Element.With(c)
		}
	}

	return i
}

func (i *imageImg) Clone() e.Element {
	return &imageImg{
		Element: i.Element.Clone(),
		img:     i.img.Clone(),
	}
}
