package bulma

import (
	_ "embed"
	"net/http"
)

// CSS is the default Bulma CSS content
//
//go:embed bulma.min.css
var CSS []byte

// CSSHandlerFunc is a http handler function that writes the Bulma CSS
func CSSHandlerFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(CSS)
}
