package bulma

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

const cdnPath = "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"

// CSSPath, when provided to HTML, sets the path to the Bulma CSS
type CSSPath string

// Language, when provided to HTML, sets the page lang attribute
type Language string

// HTitle, when provided to HTML, sets the page title
type HTitle string

// Description, when provided to HTML, sets the description metadata
type Description string

type head []gomponents.Node

type htmlElem struct {
	cssPath     string
	language    string
	title       string
	description string
	head        []gomponents.Node
	body        *Element
}

// Head identifies children as part of the head section
func Head(children ...gomponents.Node) head {
	return head(children)
}

// HTML returns a gomponents.Node which represents a whole HTML page.
//   - when a child is of type CSSPath, the path to the Bulma CSS is set to the
//     provided path (if there is no child with this type, the path is set to
//     the Bulma CDN)
//   - when a child is of type HTitle, it is used as the page title
//   - when a child is of type Language, it is used as the lang attribute
//   - when a child is of type Description, it is used as the description meta
//   - when a child is wrapped with Head, it is used as a part of the head section
//   - other children are used as children of the body element
func HTML(children ...any) gomponents.Node {
	def := &htmlElem{
		body: Elem(html.Body),
	}

	for _, c := range children {
		switch c := c.(type) {
		case CSSPath:
			def.cssPath = string(c)
		case HTitle:
			def.title = string(c)
		case Language:
			def.language = string(c)
		case Description:
			def.description = string(c)
		case head:
			def.head = append(def.head, c...)
		default:
			def.body.With(c)
		}
	}

	if def.cssPath == "" {
		def.cssPath = cdnPath
	}

	return components.HTML5(components.HTML5Props{
		Title:       def.title,
		Description: def.description,
		Language:    def.language,
		Head: append(
			def.head,
			html.Link(
				html.Rel("stylesheet"),
				html.Href(def.cssPath),
			),
		),
		Body: def.body.getChildren(),
	})
}

// Stylesheet returns a head node which makes the browser load the provided path
// as a CSS stylesheet.
func Stylesheet(path string) head {
	return head([]gomponents.Node{
		html.Link(
			html.Rel("stylesheet"),
			html.Href(path),
		),
	})
}

// Script returns a head node which makes the browser load the provided path
// as a JS script.
func Script(path string) head {
	return head([]gomponents.Node{
		html.Script(html.Src(path)),
	})
}
