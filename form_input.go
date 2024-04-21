package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

type input struct {
	inputType string
	children  []any
}

func (i *input) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			i.children = append(i.children, c...)
		case Class:
			switch c {
			case Disabled:
				i.children = append(i.children, html.Disabled())
			default:
				i.children = append(i.children, c)
			}
		case []any:
			i.With(c...)
		default:
			i.children = append(i.children, c)
		}
	}
	return i
}

func (i *input) Render(w io.Writer) error {
	return Elem(html.Input, Class("input"), html.Type(i.inputType), i.children).Render(w)
}

// InputText creates an input element of type text.
//
// The following modifiers change the input behaviour:
//   - Rounded: rounded inputs
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - html.Disabled(): disable the input
//   - html.ReadOnly(): forbid modifications
//   - Static: remove specific styling but maintain vertical spacing
//
// The following modifiers change the input color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the file input size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func InputText(children ...any) Element {
	return (&input{inputType: "text"}).With(children...)
}

// InputPassword creates an input element of type password.
//
// The following modifiers change the input behaviour:
//   - Rounded: rounded inputs
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - html.Disabled(): disable the input
//   - html.ReadOnly(): forbid modifications
//   - Static: remove specific styling but maintain vertical spacing
//
// The following modifiers change the input color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the file input size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func InputPassword(children ...any) Element {
	return (&input{inputType: "password"}).With(children...)
}

// InputEmail creates an input element of type email.
//
// The following modifiers change the input behaviour:
//   - Rounded: rounded inputs
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - html.Disabled(): disable the input
//   - html.ReadOnly(): forbid modifications
//   - Static: remove specific styling but maintain vertical spacing
//
// The following modifiers change the input color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the file input size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func InputEmail(children ...any) Element {
	return (&input{inputType: "email"}).With(children...)
}

// InputTel creates an input element of type tel.
//
// The following modifiers change the input behaviour:
//   - Rounded: rounded input
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Loading: add a a loading spinner to the right of the input
//   - html.Disabled(): disable the input
//   - html.ReadOnly(): forbid modifications
//   - Static: remove specific styling but maintain vertical spacing
//
// The following modifiers change the input color:
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//
// The following modifiers change the input size:
//   - Small
//   - Normal
//   - Medium
//   - Large
func InputTel(children ...any) Element {
	return (&input{inputType: "tel"}).With(children...)
}
