package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// ImgAlt is a modifier that, when applied to ImageImg, adds an alt attribute.
type ImgAlt string

// Image creates a figure element with class "image".
//
// The following modifiers change the image behaviour:
//   - ImgSq16: fixed square image, of 16x16px
//   - ImgSq24: fixed square image, of 24x24px
//   - ImgSq32: fixed square image, of 32x32px
//   - ImgSq48: fixed square image, of 48x48px
//   - ImgSq64: fixed square image, of 64x64px
//   - ImgSq96: fixed square image, of 96x96px
//   - ImgSq128 : fixed square image, of 128x128px
//   - ImgSquare: force image ratio to square
//   - Img1By1: force image ratio to 1 by 1
//   - Img5By4: force image ratio to 5 by 4
//   - Img4By3: force image ratio to 4 by 3
//   - Img3By2: force image ratio to 3 by 2
//   - Img5By3: force image ratio to 5 by 3
//   - Img16By9: force image ratio to 16 by 9
//   - Img2By1: force image ratio to 2 by 1
//   - Img3By1: force image ratio to 3 by 1
//   - Img4By5: force image ratio to 4 by 5
//   - Img3By4: force image ratio to 3 by 4
//   - Img2By3: force image ratio to 2 by 3
//   - Img3By5: force image ratio to 3 by 5
//   - Img9By16: force image ratio to 9 by 16
//   - Img1By2: force image ratio to 1 by 2
//   - Img1By3: force image ratio to 1 by 3
//   - FullWidth: make sure the image takes the whole width of its container
//     (usually not needed)
//
// Use the Rounded modifier on the inner image to make it rounded (associate
// with an ImgSq* modifier on the Image element).
//
// Use the Ratio modifier on the inner element (other than image) in order
// to apply it the ratio (associate with an Img*By* modifier on the Image
// element).
func Image(children ...any) Element {
	return Elem(html.Figure, Class("image"), children)
}

// ImageImg creates a figure element with class "image" and the inner img
// element, with the provided src.
//   - when a child is marked with b.Inner, it is forcibly applied to the <img> element
//   - when a child is marked with b.Outer, it is forcibly applied to the <figure> element
//
// Use ImgAlt to define the image alt text.
//
// The following modifiers change the image behaviour:
//   - ImgSq16: fixed square image, of 16x16px
//   - ImgSq24: fixed square image, of 24x24px
//   - ImgSq32: fixed square image, of 32x32px
//   - ImgSq48: fixed square image, of 48x48px
//   - ImgSq64: fixed square image, of 64x64px
//   - ImgSq96: fixed square image, of 96x96px
//   - ImgSq128 : fixed square image, of 128x128px
//   - ImgSquare: force image ratio to square
//   - Img1By1: force image ratio to 1 by 1
//   - Img5By4: force image ratio to 5 by 4
//   - Img4By3: force image ratio to 4 by 3
//   - Img3By2: force image ratio to 3 by 2
//   - Img5By3: force image ratio to 5 by 3
//   - Img16By9: force image ratio to 16 by 9
//   - Img2By1: force image ratio to 2 by 1
//   - Img3By1: force image ratio to 3 by 1
//   - Img4By5: force image ratio to 4 by 5
//   - Img3By4: force image ratio to 3 by 4
//   - Img2By3: force image ratio to 2 by 3
//   - Img3By5: force image ratio to 3 by 5
//   - Img9By16: force image ratio to 9 by 16
//   - Img1By2: force image ratio to 1 by 2
//   - Img1By3: force image ratio to 1 by 3
//   - FullWidth: make sure the image takes the whole width of its container
//     (usually not needed)
//   - Rounded: make the image rounded (associate with an ImgSq* modifier)
func ImageImg(src string, children ...any) Element {
	return (&imageImg{src: src}).With(children...)
}

type imageImg struct {
	src            string
	imgChildren    []any
	figureChildren []any
}

func (i *imageImg) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onImg:
			i.imgChildren = append(i.imgChildren, c...)
		case onFigure:
			i.figureChildren = append(i.figureChildren, c...)
		case Class:
			if c == Rounded {
				i.imgChildren = append(i.imgChildren, c)
			} else {
				i.figureChildren = append(i.figureChildren, c)
			}
		case ImgAlt:
			i.imgChildren = append(i.imgChildren, html.Alt(string(c)))
		case []any:
			i.With(c...)
		default:
			i.figureChildren = append(i.figureChildren, c)
		}
	}

	return i
}

func (i *imageImg) Render(w io.Writer) error {
	return Elem(
		html.Figure,
		Class("image"),
		i.figureChildren,
		Elem(html.Img, html.Src(i.src), i.imgChildren),
	).Render(w)
}
