package bulma

import "github.com/maragudk/gomponents/html"

func input(inputType string, children []any) Element {
	return Elem(html.Input, Class("input"), html.Type(inputType), children)
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
	return input("text", children)
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
	return input("password", children)
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
	return input("email", children)
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
	return input("tel", children)
}
