package bulma

import "github.com/maragudk/gomponents/html"

// Tile creates a tile element.
//
// If a child is an *Element without the "tile" class, the "tile" and the
// "is-child" class are added to the child and the "is-parent" class is added to
// the tile element. This way, the parent-child relation is automatically
// configured without the need to add specific modifiers.
//
// When a Tile is added as a child of a non-tile *Element, the "is-ancestor"
// class is added to the Tile.
//
// The following modifiers change its width:
//   - Size12: 100%
//   - Size11: 91.66%
//   - Size10: 83.33%
//   - Size9: 75%
//   - Size8: 66.66%
//   - Size7: 58.33%
//   - Size6: 50%
//   - Size5: 41.66%
//   - Size4: 33.33%
//   - Size3: 25%
//   - Size2: 16.66%
//   - Size1: 8.33%
func Tile(children ...any) *Element {
	e := Elem(html.Div).
		With(Class("tile"))

	for _, c := range children {
		switch c := c.(type) {
		case *Element:
			if !c.hasClass("tile") {
				c.With(Class("tile"))
				c.With(Class("is-child"))
				e.With(Class("is-parent"))
			}
			e.With(c)
		default:
			e.With(c)
		}
	}

	return e
}

// VTile creates a tile element with the "is-vertical" class.
func VTile(children ...any) *Element {
	return Tile(children...).
		With(Class("is-vertical"))
}
