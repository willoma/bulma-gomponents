package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Level creates a level element, which should contain either LevelLeft and/or
// LevelRight children, or LevelItem children.
//   - when a child is the return value of LevelItem, LevelLeft or LevelRight,
//     it is added as a child to the level
//   - when a child is any other *Element, it is wrapped into a LevelItem and
//     added as a child to the level
//   - when a child is a container, it is wrapped into a LevelItem and added as
//     a child to the level
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
func Level(children ...any) *Element {
	e := Elem(html.Nav).
		With(Class("level"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("level-left") || c.hasClass("level-right") || c.hasClass("level-item") {
				e.With(c)
			} else {
				e.With(LevelItem(c))
			}
		case container, string:
			e.With(LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				e.With(LevelItem(c))
			} else {
				e.With(c)
			}
		default:
			e.With(c)
		}
	}

	return e
}

// LevelItem creates a level item, to be used as a child for LevelLeft,
// LevelRight or Level elements.
func LevelItem(children ...any) *Element {
	return Elem(html.Div).
		With(Class("level-item")).
		Withs(children)
}

func levelSection(right bool, children []any) *Element {
	e := Elem(html.Div)
	if right {
		e.With(Class("level-right"))
	} else {
		e.With(Class("level-left"))
	}
	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if c.hasClass("level-item") {
				e.With(c)
			} else {
				e.With(LevelItem(c))
			}
		case container, string:
			e.With(LevelItem(c))
		case gomponents.Node:
			if !IsAttribute(c) {
				e.With(LevelItem(c))
			} else {
				e.With(c)
			}
		default:
			e.With(c)
		}
	}

	return e
}

// LevelLeft creates the left section of a level.
//   - when a child is the return value of LevelItem, it is added as a child to
//     the level section
//   - when a child is any other *Element, it is wrapped into a LevelItem and
//     added as a child to the level section
//   - when a child is a container, it is wrapped into a LevelItem and added as
//     a child to the level section
//   - when a child is a string, it is wrapped into a LevelItem and added as a
//     child to the level section
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level section
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     a LevelItem and added as a child to the level section
//   - other children types are added as children to the level section
func LevelLeft(children ...any) *Element {
	return levelSection(false, children)
}

// LevelLeft creates the right section of a level.
//   - when a child is the return value of LevelItem, it is added as a child to
//     the level section
//   - when a child is any other *Element, it is wrapped into a LevelItem and
//     added as a child to the level section
//   - when a child is a container, it is wrapped into a LevelItem and added as
//     a child to the level section
//   - when a child is a string, it is wrapped into a LevelItem and added as a
//     child to the level section
//   - when a child is a gomponents.Node with type gomponents.AttributeType, it
//     is added as an attribute to the level section
//   - when a child is a gopmponents.Node with another type, it is wrapped into
//     a LevelItem and added as a child to the level section
//   - other children types are added as children to the level section
func LevelRight(children ...any) *Element {
	return levelSection(true, children)
}
