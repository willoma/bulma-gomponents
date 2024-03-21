package bulma

import (
	"io"

	"github.com/maragudk/gomponents/html"
)

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
func File(children ...any) Element {
	return new(file).With(children...)
}

type file struct {
	fileName      string
	divChildren   []any
	ctaChildren   []any
	inputChildren []any
}

func (f *file) With(children ...any) Element {
	for _, c := range children {
		switch c := c.(type) {
		case Class, ColorClass:
			f.divChildren = append(f.divChildren, c)
		case string:
			f.ctaChildren = append(
				f.ctaChildren,
				Elem(html.Span, Class("file-label"), c),
			)
		case FileName:
			f.divChildren = append(f.divChildren, Class("has-name"))
			f.fileName = string(c)
		case FileNameAutoUpdate:
			f.divChildren = append(f.divChildren, Class("has-name"))
			f.fileName = string(c)
			f.inputChildren = append(f.inputChildren, On("change", fileNameAutoUpdateScript))
		case *icon:
			c.iconClass = Class("file-icon")
			f.ctaChildren = append(f.ctaChildren, c)
		case Element:
			f.ctaChildren = append(f.ctaChildren, c)
		case []any:
			f.With(c...)
		default:
			f.inputChildren = append(f.inputChildren, c)
		}
	}

	return f
}

func (f *file) Render(w io.Writer) error {
	label := Elem(
		html.Label,
		Class("file-label"),
		Elem(html.Input, Class("file-input"), html.Type("file"), f.inputChildren),
		Elem(html.Span, Class("file-cta"), f.ctaChildren),
	)

	if f.fileName != "" {
		label.With(Elem(html.Span, Class("file-name"), f.fileName))
	}

	return Elem(html.Div, Class("file"), f.divChildren, label).Render(w)
}
