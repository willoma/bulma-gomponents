package bulma

import (
	"github.com/maragudk/gomponents/html"
)

// ImgAlt is a modifier that, when applied to ImageImg, adds an alt attribute.
type ImgAlt string

// Image creates a figure element with class "image".
//
// https://willoma.github.io/bulma-gomponents/image.html
func Image(children ...any) Element {
	return Elem(html.Figure, Class("image"), children)
}

// ImageImg creates a figure element with class "image" and the inner img
// element, with the provided src.
//
// https://willoma.github.io/bulma-gomponents/image.html
func ImageImg(src string, children ...any) Element {
	img := Elem(html.Img, html.Src(src))
	i := &imageImg{
		Element: Elem(html.Figure, Class("image"), img),
		img:     img,
	}
	i.With(children...)
	return i
}

type imageImg struct {
	Element
	img Element
}

func (i *imageImg) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onImg:
			i.img.With(c...)
		case onFigure:
			i.Element.With(c...)
		case Class:
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

func (i *imageImg) Clone() Element {
	return &imageImg{
		Element: i.Element.Clone(),
		img:     i.img.Clone(),
	}
}
