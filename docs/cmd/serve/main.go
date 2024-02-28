package main

import (
	"log"
	"net/http"
	"strings"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/docs"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/fa"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /bulma.css", b.CSSHandlerFunc)
	mux.HandleFunc("GET /htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(docs.HtmxJS)
	})

	mux.Handle("GET /fa/", http.StripPrefix("/fa/", http.FileServer(http.FS(fa.Assets))))

	for _, section := range docs.Sections {
		for _, page := range section.Pages {
			page.Path = strings.Replace(page.Path, "/index", "", 1)
			route(mux, page)
		}
	}

	log.Print("Listening on localhost:8080...")
	http.ListenAndServe("localhost:8080", mux)
}

func route(mux *http.ServeMux, p *c.Page) {
	mux.HandleFunc("GET "+p.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		docs.Layout(p).Render(w)
	})
}
