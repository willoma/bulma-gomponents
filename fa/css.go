package fa

import (
	"embed"
	"net/http"

	"maragu.dev/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

//go:embed css webfonts
var Assets embed.FS

// CSSHandler is a http handler function for Font Awesome assets.
var CSSHandler = http.FileServer(http.FS(Assets))

// CSSHead returns head section elements for the Font Awesome CSS at the
// provided base path, for inclusion as a child of b.HTML.
func CSSHead(basePath string) any {
	return b.Head(
		html.Link(
			html.Rel("stylesheet"),
			html.Href(basePath+"/css/brands.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(basePath+"/css/fontawesome.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(basePath+"/css/regular.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(basePath+"/css/solid.min.css"),
		),
	)
}
