package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Icon creates an icon span.
//
// The following modifiers change the icon color:
//   - White
//   - Black
//   - Light
//   - Dark
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - BlackBis
//   - BlackTer
//   - GreyDarker
//   - GreyDark
//   - Grey
//   - GreyLight
//   - GreyLighter
//   - WhiteTer
//   - WhiteBis
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
//   - PrimaryDark
//   - LinkDark
//   - InfoDark
//   - SuccessDark
//   - WarningDark
//   - DangerDark
//
// The following modifiers change the icon size:
//   - Small
//   - Medium
//   - Large
func Icon(children ...any) Element {
	return (&icon{iconClass: Class("icon")}).With(children...)
}

type icon struct {
	iconClass Class
	children  []any
}

func (i *icon) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case ColorClass:
			i.children = append(i.children, c.Text())
		case []any:
			i.With(c...)
		default:
			i.children = append(i.children, c)
		}
	}

	return i
}

func (i *icon) Render(w io.Writer) error {
	return Elem(html.Span, i.iconClass, i.children).Render(w)
}

// IconText creates an icon-text span and embed all its non-icons children into
// spans.
//
// The following modifiers change the icon and text color:
//   - White
//   - Black
//   - Light
//   - Dark
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - BlackBis
//   - BlackTer
//   - GreyDarker
//   - GreyDark
//   - Grey
//   - GreyLight
//   - GreyLighter
//   - WhiteTer
//   - WhiteBis
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
//   - PrimaryDark
//   - LinkDark
//   - InfoDark
//   - SuccessDark
//   - WarningDark
//   - DangerDark
func IconText(children ...any) Element {
	i := &iconText{el: html.Span}
	i.addChildren(children)
	return i.elem()
}

// IconText creates a flex icon-text span and embed all its non-icons children
// into spans.
//
// The following modifiers change the icon and text color:
//   - White
//   - Black
//   - Light
//   - Dark
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - BlackBis
//   - BlackTer
//   - GreyDarker
//   - GreyDark
//   - Grey
//   - GreyLight
//   - GreyLighter
//   - WhiteTer
//   - WhiteBis
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
//   - PrimaryDark
//   - LinkDark
//   - InfoDark
//   - SuccessDark
//   - WarningDark
//   - DangerDark
func FlexIconText(children ...any) Element {
	i := &iconText{el: html.Div}
	i.addChildren(children)
	return i.elem()
}

type iconText struct {
	el       func(children ...gomponents.Node) gomponents.Node
	children []any
}

func (i *iconText) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case ColorClass:
			i.children = append(i.children, c.Text())
		case []any:
			i.addChildren(c)
		default:
			i.children = append(i.children, c)
		}
	}
}

func (i *iconText) elem() Element {
	return Elem(i.el, Class("icon-text"), elemOptionSpanAroundNonIconsAlways, i.children)
}
