package fa

import (
	"embed"
	"net/http"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

//go:embed css webfonts
var Assets embed.FS

// CSSHandler is a http handler function for Font Awesome assets
var CSSHandler = http.FileServer(http.FS(Assets))

func CSSPath(path string) []gomponents.Node {
	return b.Head(
		html.Link(
			html.Rel("stylesheet"),
			html.Href(path+"/css/brands.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(path+"/css/fontawesome.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(path+"/css/regular.min.css"),
		),
		html.Link(
			html.Rel("stylesheet"),
			html.Href(path+"/css/solid.min.css"),
		),
	)
}
