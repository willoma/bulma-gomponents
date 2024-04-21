package bulma

import (
	"io"
	"strconv"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Textarea creates a textarea element.
//
// The following modifiers change the textarea behaviour:
//   - Rows(int): set the textarea heights, in number of rows
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - b.Disabled: disable the input
//   - html.ReadOnly(): forbid modifications
//   - FixedSize: disable the textarea resizing possibility
//
// The following modifiers change the textarea color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the textarea size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func Textarea(children ...any) Element {
	return new(textarea).With(children...)
}

type textarea struct {
	children []any
}

func (t *textarea) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			switch c {
			case Disabled:
				t.children = append(t.children, html.Disabled())
			default:
				t.children = append(t.children, c)
			}
		case []any:
			t.With(c...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}

func (t *textarea) Render(w io.Writer) error {
	return Elem(html.Textarea, Class("textarea"), t.children).Render(w)
}

// Rows changes a textarea height.
func Rows(rows int) gomponents.Node {
	return html.Rows(strconv.Itoa(rows))
}
