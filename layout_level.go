package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Level creates a level element, which should contain either LevelLeft and/or
// LevelRight children, or LevelItem children.
//   - when a child is the return value of LevelItem, LevelLeft or LevelRight,
//     it is added as a child to the level
//   - when a child is any other Element, it is wrapped into a LevelItem and
//     added as a child to the level
//   - when a child is a string, it is wrapped into a LevelItem and added as a
//     child to the level
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     a LevelItem and added as a child to the level
//   - other children types are added as children to the level
//
// The following modifiers change the level behaviour:
//   - Mobile: show the level horizontally on mobile too (otherwise, the level
//     items are placed vertically)
func Level(children ...any) Element {
	l := &level{}
	l.addChildren(children)
	return l.elem()
}

type level struct {
	children []any
}

func (l *level) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			if c.hasClass("level-left") || c.hasClass("level-right") || c.hasClass("level-item") {
				l.children = append(l.children, c)
			} else {
				l.children = append(l.children, LevelItem(c))
			}
		case string:
			l.children = append(l.children, LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				l.children = append(l.children, LevelItem(c))
			} else {
				l.children = append(l.children, c)
			}
		case []any:
			l.addChildren(c)
		default:
			l.children = append(l.children, c)
		}
	}
}

func (l *level) elem() Element {
	return Elem(html.Nav, Class("level"), l.children)
}

// LevelItem creates a level item, to be used as a child for LevelLeft,
// LevelRight or Level elements.
func LevelItem(children ...any) Element {
	return Elem(html.Div, Class("level-item"), children)
}

// LevelLeft creates the left section of a level.
//   - when a child is the return value of LevelItem, it is added as a child to
//     the level section
//   - when a child is any other Element, it is wrapped into a LevelItem and
//     added as a child to the level section
//   - when a child is a string, it is wrapped into a LevelItem and added as a
//     child to the level section
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level section
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     a LevelItem and added as a child to the level section
//   - other children types are added as children to the level section
func LevelLeft(children ...any) Element {
	l := &levelSection{}
	l.addChildren(children)
	return l.elem()
}

// LevelLeft creates the right section of a level.
//   - when a child is the return value of LevelItem, it is added as a child to
//     the level section
//   - when a child is any other Element, it is wrapped into a LevelItem and
//     added as a child to the level section
//   - when a child is a string, it is wrapped into a LevelItem and added as a
//     child to the level section
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level section
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     a LevelItem and added as a child to the level section
//   - other children types are added as children to the level section
func LevelRight(children ...any) Element {
	l := &levelSection{right: true}
	l.addChildren(children)
	return l.elem()
}

type levelSection struct {
	right    bool
	children []any
}

func (l *levelSection) addChildren(children []any) {
	for _, c := range children {
		switch c := c.(type) {
		case Element:
			if c.hasClass("level-item") {
				l.children = append(l.children, c)
			} else {
				l.children = append(l.children, LevelItem(c))
			}
		case string:
			l.children = append(l.children, LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				l.children = append(l.children, LevelItem(c))
			} else {
				l.children = append(l.children, c)
			}
		case []any:
			l.addChildren(c)
		default:
			l.children = append(l.children, c)
		}
	}
}

func (l *levelSection) elem() Element {
	e := Elem(html.Div)

	if l.right {
		e.With(Class("level-right"))
	} else {
		e.With(Class("level-left"))
	}

	return e.With(l.children...)
}
