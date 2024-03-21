package bulma

import (
	"io"

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
	return new(level).With(children...)
}

type level struct {
	children []any
}

func (l *level) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *levelItem, *levelLeft, *levelRight:
			l.children = append(l.children, c)
		case Element, string:
			l.children = append(l.children, LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				l.children = append(l.children, LevelItem(c))
			} else {
				l.children = append(l.children, c)
			}
		case []any:
			l.With(c...)
		default:
			l.children = append(l.children, c)
		}
	}

	return l
}

func (l *level) Render(w io.Writer) error {
	return Elem(html.Nav, Class("level"), l.children).Render(w)
}

// LevelItem creates a level item, to be used as a child for LevelLeft,
// LevelRight or Level elements.
func LevelItem(children ...any) *levelItem {
	return &levelItem{Elem(html.Div, Class("level-item"), children)}
}

type levelItem struct {
	Element
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
func LevelLeft(children ...any) *levelLeft {
	return &levelLeft{
		(&levelSection{positionClass: Class("level-left")}).With(children),
	}
}

type levelLeft struct {
	Element
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
func LevelRight(children ...any) *levelRight {
	return &levelRight{
		(&levelSection{positionClass: Class("level-right")}).With(children),
	}
}

type levelRight struct {
	Element
}

type levelSection struct {
	positionClass Class
	children      []any
}

func (l *levelSection) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *levelItem:
			l.children = append(l.children, c)
		case Element:
			l.children = append(l.children, LevelItem(c))
		case string:
			l.children = append(l.children, LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				l.children = append(l.children, LevelItem(c))
			} else {
				l.children = append(l.children, c)
			}
		case []any:
			l.With(c...)
		default:
			l.children = append(l.children, c)
		}
	}

	return l
}

func (l *levelSection) Render(w io.Writer) error {
	return Elem(html.Div, l.positionClass, l.children).Render(w)
}
