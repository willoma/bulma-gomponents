package bulma

import (
	"io"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Level creates a level element.
//
// http://willoma.github.io/bulma-gomponents/level.html
func Level(children ...any) Element {
	l := &level{Element: Elem(html.Nav, Class("level"))}
	l.With(children...)
	return l
}

type level struct {
	Element
	left  Element
	right Element

	rendered sync.Once
}

func (l *level) addToLeft(children ...any) {
	if l.left == nil {
		l.left = Elem(html.Div, Class("level-left"))
	}
	flattenLevelSection(l.left, children)
}

func (l *level) addToRight(children ...any) {
	if l.right == nil {
		l.right = Elem(html.Div, Class("level-right"))
	}
	flattenLevelSection(l.right, children)
}

func (l *level) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case levelLeft:
			l.addToLeft(c...)
		case levelRight:
			l.addToRight(c...)
		case Element:
			c.With(Class("level-item"))
			l.Element.With(c)
		case string:
			l.Element.With(Elem(html.P, Class("level-item"), c))
		case gomponents.Node:
			if isAttribute(c) {
				l.Element.With(c)
			} else {
				l.Element.With(Elem(html.Div, Class("level-item"), c))
			}
		case []any:
			l.With(c...)
		default:
			l.Element.With(c)
		}
	}

	return l
}

func (l *level) Render(w io.Writer) error {
	l.rendered.Do(func() {
		if l.left != nil {
			l.Element.With(l.left)
		}
		if l.right != nil {
			l.Element.With(l.right)
		}
	})

	return l.Element.Render(w)
}

func LevelLeft(children ...any) levelLeft {
	return levelLeft(children)
}

type levelLeft []any

func LevelRight(children ...any) levelRight {
	return levelRight(children)
}

type levelRight []any

func flattenLevelSection(target Element, children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			c.With(Class("level-item"))
			target.With(c)
		case string:
			target.With(Elem(html.P, Class("level-item"), c))
		case gomponents.Node:
			if !isAttribute(c) {
				target.With(Elem(html.Div, Class("level-item"), c))
			} else {
				target.With(c)
			}
		case []any:
			flattenLevelSection(target, c)
		default:
			target.With(c)
		}
	}
}
