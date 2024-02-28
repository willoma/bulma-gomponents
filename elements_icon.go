package bulma

import (
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
func Icon(children ...any) *Element {
	e := Elem(html.Span).
		With(Class("icon"))

	for _, c := range children {
		switch c := c.(type) {
		case ColorClass:
			e.With(c.Text())
		default:
			e.With(c)
		}
	}

	return e
}

func iconText(
	el func(children ...gomponents.Node) gomponents.Node,
	children []any,
) *Element {
	e := Elem(el).
		With(Class("icon-text"))
	e.spanAroundNonIconsAlways = true

	for _, c := range children {
		switch c := c.(type) {
		case ColorClass:
			e.With(c.Text())
		default:
			e.With(c)
		}
	}

	return e
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
func IconText(children ...any) *Element {
	return iconText(html.Span, children)
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
func FlexIconText(children ...any) *Element {
	return iconText(html.Div, children)
}
