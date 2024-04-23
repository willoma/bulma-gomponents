package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func newButton(fn func(...gomponents.Node) gomponents.Node, children ...any) Element {
	b := &button{Elem(fn, Class("button"), elemOptionSpanAroundNonIconsIfHasIcons)}
	b.With(children...)
	return b
}

type button struct {
	Element
}

// Button creates a button.
//
// https://willoma.github.io/bulma-gomponents/button.html
func Button(children ...any) Element {
	return newButton(html.Button, children...)
}

// ButtonA creates a button-looking link.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonA(children ...any) Element {
	return newButton(html.A, children...)
}

// ButtonSubmit creates a submit button.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonSubmit(children ...any) Element {
	return newButton(html.Button, html.Type("submit"), children)
}

// ButtonInputSubmit creates an input of type submit.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonInputSubmit(value string, children ...any) Element {
	return newButton(html.Input, html.Type("submit"), html.Value(value), children)
}

// ButtonInputReset creates an input of type reset.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonInputReset(value string, children ...any) Element {
	return newButton(html.Input, html.Type("reset"), html.Value(value), children)
}

// Buttons creates a list of buttons.
//
// https://willoma.github.io/bulma-gomponents/button.html
func Buttons(children ...any) Element {
	b := &buttons{Elem(html.Div, Class("buttons"))}
	b.With(children...)
	return b
}

type buttons struct {
	Element
}

func (b *buttons) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class:
			b.Element.With(changeSizePrefix("are-", c))
		case []any:
			b.With(c...)
		default:
			b.Element.With(c)
		}
	}

	return b
}
