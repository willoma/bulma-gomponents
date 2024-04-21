package bulma

import (
	"io"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type button struct {
	fn       func(...gomponents.Node) gomponents.Node
	children []any
}

func (b *button) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			b.With(c...)
		default:
			b.children = append(b.children, c)
		}
	}

	return b
}

func (b *button) Render(w io.Writer) error {
	return Elem(b.fn, Class("button"), elemOptionSpanAroundNonIconsIfHasIcons, b.children).Render(w)
}

// Button creates a button.
//
// In order to disable the button, use html.Disabled().
//
// If the button contains an icon, all other elements are automatically wrapped
// in spans.
//
// The following modifiers change the button behaviour:
//   - Text: display as an underlined text
//   - Ghost: display as a blue underlined link
//   - Responsive: responsive size
//   - FullWidth: take the whole width
//   - Outlined: outline style
//   - Inverted: inverted style
//   - Rounded: rounded button
//   - Hovered: apply the hovered style
//   - Focused: apply the focused style
//   - Active: apply the active style
//   - Loading: replace the content with a loading spinner
//   - Static: make the button non-interactive
//   - html.Disabled(): disable the button
//   - Selected: in a list of attached buttons (Buttons with Addons), make sure
//     this button is above the other buttons
//
// The following modifiers change the button size:
//   - Small
//   - Normal
//   - Medium
//   - Large
//
// The following modifiers change the button color:
//   - White
//   - Light
//   - Dark
//   - Black
//   - Primary
//   - Link
//   - Info
//   - Success
//   - Warning
//   - Danger
//   - PrimaryLight
//   - LinkLight
//   - InfoLight
//   - SuccessLight
//   - WarningLight
//   - DangerLight
//   - PrimaryDark
//   - LinkDark
//   - InfoDark
//   - SuccessDark
//   - WarningDark
//   - DangerDark
func Button(children ...any) Element {
	return (&button{fn: html.Button}).With(children...)
}

// ButtonA creates a button-looking link.
//
// See the documentation on the Button function for modifiers details.
func ButtonA(children ...any) Element {
	return (&button{fn: html.A}).With(children...)
}

// ButtonSubmit creates a submit button.
//
// See the documentation on the Button function for modifiers details.
func ButtonSubmit(children ...any) Element {
	return (&button{fn: html.Button}).With(html.Type("submit"), children)
}

// ButtonInputSubmit creates an input of type submit.
//
// See the documentation on the Button function for modifiers details.
func ButtonInputSubmit(value string, children ...any) Element {
	return (&button{fn: html.Input}).With(
		html.Type("submit"), html.Value(value), children,
	)
}

// ButtonInputReset creates an input of type reset.
//
// See the documentation on the Button function for modifiers details.
func ButtonInputReset(value string, children ...any) Element {
	return (&button{fn: html.Input}).With(
		html.Type("reset"), html.Value(value), children,
	)
}

// Buttons creates a list of buttons.
//
// The following modifiers change the list of buttons behaviour:
//   - Addons: attach the buttons together
//   - Centered: center the buttons
//   - Right: align the buttons to the right
//
// The following modifiers change the size of all buttons in the list:
//   - Small
//   - Medium
//   - Large
func Buttons(children ...any) Element {
	return new(buttons).With(children...)
}

type buttons struct {
	children []any
}

func (b *buttons) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			b.children = append(b.children, changeSizePrefix("are-", c))
		case []any:
			b.With(c...)
		default:
			b.children = append(b.children, c)
		}
	}

	return b
}

func (b *buttons) Render(w io.Writer) error {
	return Elem(html.Div, Class("buttons"), b.children).Render(w)
}
