package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

// Tile creates a tile element.
//
// If a child is an Element without the "tile" class, the "tile" and the
// "is-child" class are added to the child and the "is-parent" class is added to
// the tile element. This way, the parent-child relation is automatically
// configured without the need to add specific modifiers.
//
// When a Tile is added as a child of a non-tile Element, the "is-ancestor"
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
func Tile(children ...any) Element {
	return (&tile{}).With(children...)
}

type tile struct {
	children []any
}

func (t *tile) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case *tile:
			t.children = append(t.children, c)
		case Element:
			c.With(Class("tile"), Class("is-child"))
			t.children = append(t.children, Class("is-parent"), c)
		case []any:
			t.With(c...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}

func (t *tile) Render(w io.Writer) error {
	return Elem(html.Div, Class("tile"), t.children).Render(w)
}

// VTile creates a tile element with the "is-vertical" class.
func VTile(children ...any) Element {
	return Tile(children, Class("is-vertical"))
}
