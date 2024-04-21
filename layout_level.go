package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Level creates a level element.
//   - when a child is the result of LevelLeft(...), its children are added to
//     the level left section, with following rules applied
//   - when a child is the result of LevelRight(...), its children are added to
//     the level right section, with following rules applied
//   - when a child is an Element, the "level-item" class is added to it and it
//     is added as a child to the level
//   - when a child is a string, it is wrapped in an Element with the
//     "level-item" class which is added as a child to the level
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     an Element with the "level-item" class and added as a child to the level
//   - other children types are added as children to the level
//
// For children included provided with LevelLeft or LevelRight:
//   - when a child is an Element, the "level-item" class is added to it and it
//     is added as a child to the level section
//   - when a child is a string, it is wrapped in an Element with the
//     "level-item" class which is added as a child to the level section
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level section
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     an Element with the "level-item" class and added as a child to the level
//     section
//   - other children types are added as children to the level section
func Level(children ...any) Element {
	return new(level).With(children...)
}

type level struct {
	leftChildren  []any
	rightChildren []any
	children      []any
}

func (l *level) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case levelLeft:
			l.leftChildren = append(l.leftChildren, flattenLevelSection(c)...)
		case levelRight:
			l.rightChildren = append(l.rightChildren, flattenLevelSection(c)...)
		case Element:
			c.With(Class("level-item"))
			l.children = append(l.children, c)
		case string:
			l.children = append(l.children, Elem(html.P, Class("level-item"), c))
		case gomponents.Node:
			if !IsAttribute(c) {
				l.children = append(l.children, Elem(html.Div, Class("level-item"), c))
			} else {
				l.children = append(l.children, c)
			}
		case []any:
			l.children = append(l.children, c...)
		default:
			l.children = append(l.children, c)
		}
	}

	return l
}

func (l *level) Render(w io.Writer) error {
	elem := Elem(html.Nav, Class("level"), l.children)
	if len(l.leftChildren) > 0 {
		elem.With(Elem(html.Div, Class("level-left"), l.leftChildren))
	}

	if len(l.rightChildren) > 0 {
		elem.With(Elem(html.Div, Class("level-right"), l.rightChildren))
	}

	return elem.Render(w)
}

func LevelLeft(children ...any) levelLeft {
	return levelLeft(children)
}

type levelLeft []any

func LevelRight(children ...any) levelRight {
	return levelRight(children)
}

type levelRight []any

func flattenLevelSection(children []any) []any {
	var result []any

	for _, c := range children {
		switch c := c.(type) {
		case Element:
			c.With(Class("level-item"))
			result = append(result, c)
		case string:
			result = append(result, Elem(html.P, Class("level-item"), c))
		case gomponents.Node:
			if !IsAttribute(c) {
				result = append(result, Elem(html.Div, Class("level-item"), c))
			} else {
				result = append(result, c)
			}
		case []any:
			result = append(result, flattenLevelSection(c)...)
		default:
			result = append(result, c)
		}
	}

	return result
}
