package bulma

import (
	"io"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
)

// Level creates a level element.
//
// http://willoma.github.io/bulma-gomponents/level.html
func Level(children ...any) e.Element {
	l := &level{level: e.Nav(e.Class("level"))}
	l.With(children...)
	return l
}

type level struct {
	level e.Element
	left  e.Element
	right e.Element
}

func (l *level) addToLeft(children ...any) {
	if l.left == nil {
		l.left = e.Div(e.Class("level-left"))
	}
	flattenLevelSection(l.left, children)
}

func (l *level) addToRight(children ...any) {
	if l.right == nil {
		l.right = e.Div(e.Class("level-right"))
	}
	flattenLevelSection(l.right, children)
}

func (l *level) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case levelLeft:
			l.addToLeft(c...)
		case levelRight:
			l.addToRight(c...)
		case e.Element:
			c.With(e.Class("level-item"))
			l.level.With(c)
		case string:
			l.level.With(e.P(e.Class("level-item"), c))
		case gomponents.Node:
			if e.IsAttribute(c) {
				l.level.With(c)
			} else {
				l.level.With(e.Div(e.Class("level-item"), c))
			}
		case []any:
			l.With(c...)
		default:
			l.level.With(c)
		}
	}

	return l
}

func (l *level) Render(w io.Writer) error {
	level := l.level.Clone()

	if l.left != nil {
		level.With(l.left)
	}
	if l.right != nil {
		level.With(l.right)
	}

	return level.Render(w)
}

func (l *level) Clone() e.Element {
	return &level{
		level: l.level.Clone(),
		left:  l.left.Clone(),
		right: l.right.Clone(),
	}
}

func LevelLeft(children ...any) levelLeft {
	return levelLeft(children)
}

type levelLeft []any

func LevelRight(children ...any) levelRight {
	return levelRight(children)
}

type levelRight []any

func flattenLevelSection(target e.Element, children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case e.Element:
			c.With(e.Class("level-item"))
			target.With(c)
		case string:
			target.With(e.P(e.Class("level-item"), c))
		case gomponents.Node:
			if !e.IsAttribute(c) {
				target.With(e.Div(e.Class("level-item"), c))
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
