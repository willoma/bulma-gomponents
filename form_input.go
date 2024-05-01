package bulma

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"
)

func newInput(inputType string, children []any) e.Element {
	i := &input{e.Input(e.Class("input"), html.Type(inputType))}
	i.With(children...)
	return i
}

type input struct {
	e.Element
}

func (i *input) With(children ...any) e.Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			i.Element.With(c)
		case e.Class:
			switch c {
			case Disabled:
				i.Element.With(html.Disabled())
			default:
				i.Element.With(c)
			}
		case []any:
			i.With(c...)
		default:
			i.Element.With(c)
		}
	}
	return i
}

func (i *input) Clone() e.Element {
	return &input{Element: i.Element.Clone()}
}

// InputText creates an input element of type text.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputText(children ...any) e.Element {
	return newInput("text", children)
}

// InputPassword creates an input element of type password.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputPassword(children ...any) e.Element {
	return newInput("password", children)
}

// InputEmail creates an input element of type email.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputEmail(children ...any) e.Element {
	return newInput("email", children)
}

// InputTel creates an input element of type tel.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputTel(children ...any) e.Element {
	return newInput("tel", children)
}
