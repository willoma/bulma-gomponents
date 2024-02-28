package bulma

import "github.com/maragudk/gomponents/html"

const (
	fileNameAutoUpdateScript = `if (this.files.length>0) { this.parentNode.getElementsByClassName("file-name")[0].textContent = this.files[0].name }`
)

// FileName is a modifier that, when applied to File, defines the default value
// for the file-name element.
type FileName string

// FileNameAutoUpdate is a modifier that, when applied to File, defines the
// default value for the file-name element, and adds the needed javascript
// event to update the file name.
type FileNameAutoUpdate string

// File creates a file input element.
//   - when a child is a string, it is used as the call-to-action label
//   - when a child is a FileName, it is the content of the file-name element
//   - when a child is a FileNameAutoUpdate, it is the content of the file-name
//     element, which changes when a file is selected
//
// The following modifiers change the file input behaviour:
//   - Right: move the call-to-action to the right side, align the file input
//     to the right
//   - FullWidth: expand the name to fill up the space
//   - Boxed: make the input a boxed block
//   - Centered: align the file input to the center
//
// The following modifiers change the file input color:
//   - White
//   - Black
//   - Light
//   - Dark
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
func File(children ...any) *Element {
	div := Elem(html.Div).
		With(Class("file"))

	label := Elem(html.Label).
		With(Class("file-label"))

	input := Elem(html.Input).
		With(Class("file-input")).
		With(html.Type("file"))

	callToAction := Elem(html.Span).
		With(Class("file-cta"))

	var fileName *Element

	for _, c := range children {
		switch c := c.(type) {
		case Class, ColorClass:
			div.With(c)
		case string:
			callToAction.With(
				Elem(html.Span).
					With(Class("file-label")).
					With(c),
			)
		case FileName:
			div.With(Class("has-name"))
			fileName = Elem(html.Span).
				With(Class("file-name")).
				With(string(c))
		case FileNameAutoUpdate:
			div.With(Class("has-name"))
			fileName = Elem(html.Span).
				With(Class("file-name")).
				With(string(c))
			input.With(
				On("change", fileNameAutoUpdateScript),
			)
		case *Element:
			if c.hasClass("icon") {
				c.With(Class("file-icon"))
				c.classes["icon"] = false
			}
			callToAction.With(c)
		default:
			input.With(c)
		}
	}

	label.
		With(input).
		With(callToAction)

	if fileName != nil {
		label.With(fileName)
	}

	return div.With(label)
}
