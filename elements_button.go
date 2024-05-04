package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

func newButton(fn func(...gomponents.Node) gomponents.Node, children ...any) e.Element {
	b := &spanAroundNonIconsIfHasIcons{elemFn: fn}
	b.With(e.Class("button")).With(children...)

	return b
}

type button struct {
	e.Element
}

func (b *button) Clone() e.Element {
	return &button{b.Element.Clone()}
}

// Button creates a button.
//
// https://willoma.github.io/bulma-gomponents/button.html
func Button(children ...any) e.Element {
	return newButton(html.Button, children...)
}

// ButtonA creates a button-looking link.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonA(children ...any) e.Element {
	return newButton(html.A, children...)
}

// ButtonAHref creates a button-looking link, with the provided href.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonAHref(href string, children ...any) e.Element {
	return newButton(html.A, html.Href(href), children)
}

// ButtonSubmit creates a submit button.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonSubmit(children ...any) e.Element {
	return newButton(html.Button, html.Type("submit"), children)
}

// ButtonInputSubmit creates an input of type submit.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonInputSubmit(value string, children ...any) e.Element {
	return newButton(html.Input, html.Type("submit"), html.Value(value), children)
}

// ButtonInputReset creates an input of type reset.
//
// https://willoma.github.io/bulma-gomponents/button.html
func ButtonInputReset(value string, children ...any) e.Element {
	return newButton(html.Input, html.Type("reset"), html.Value(value), children)
}

// Buttons creates a list of buttons.
//
// https://willoma.github.io/bulma-gomponents/button.html
func Buttons(children ...any) e.Element {
	b := &buttons{e.Div(e.Class("buttons"))}
	b.With(children...)
	return b
}

type buttons struct {
	e.Element
}

func (b *buttons) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case e.Class:
			b.Element.With(changeSizePrefix("are-", c))
		case []any:
			b.With(c...)
		default:
			b.Element.With(c)
		}
	}

	return b
}

func (b *buttons) Clone() e.Element {
	return &buttons{b.Element.Clone()}
}
