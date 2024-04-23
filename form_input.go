package bulma

import (
	"github.com/maragudk/gomponents/html"
)

func newInput(inputType string, children []any) Element {
	i := &input{Elem(html.Input, Class("input"), html.Type(inputType))}
	i.With(children...)
	return i
}

type input struct {
	Element
}

func (i *input) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case onInput:
			i.Element.With(c)
		case Class:
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

// InputText creates an input element of type text.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputText(children ...any) Element {
	return newInput("text", children)
}

// InputPassword creates an input element of type password.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputPassword(children ...any) Element {
	return newInput("password", children)
}

// InputEmail creates an input element of type email.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputEmail(children ...any) Element {
	return newInput("email", children)
}

// InputTel creates an input element of type tel.
//
// https://willoma.github.io/bulma-gomponents/form/input.html
func InputTel(children ...any) Element {
	return newInput("tel", children)
}
