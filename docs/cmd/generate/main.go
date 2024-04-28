package main

import (
	_ "embed"
	"flag"
	"io"
	"os"
	"path/filepath"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/docs"
	"github.com/willoma/bulma-gomponents/fa"
)

func main() {
	baseurl := flag.String("baseurl", "", "Doc base URL")
	destination := flag.String("destination", "public", "Destination directory")
	flag.Parse()

	generateAssets(*destination)
	generatePages(*destination, *baseurl)
}

func generateAssets(destination string) {
	if err := os.Mkdir(destination, 0o755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(destination, "bulma.css"), b.CSS, 0o644); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(destination, "htmx.min.js"), docs.HtmxJS, 0o644); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(destination, "fa"), 0o755); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(destination, "fa", "css"), 0o755); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(destination, "fa", "webfonts"), 0o755); err != nil {
		panic(err)
	}

	entries, err := fa.Assets.ReadDir("css")
	if err != nil {
		panic(err)
	}
	for _, finfo := range entries {
		src, err := fa.Assets.Open(filepath.Join("css", finfo.Name()))
		if err != nil {
			panic(err)
		}

		dst, err := os.Create(filepath.Join(destination, "fa", "css", finfo.Name()))
		if err != nil {
			src.Close()
			panic(err)
		}

		if _, err := io.Copy(dst, src); err != nil {
			src.Close()
			panic(err)
		}

		src.Close()
	}

	entries, err = fa.Assets.ReadDir("webfonts")
	if err != nil {
		panic(err)
	}
	for _, finfo := range entries {
		src, err := fa.Assets.Open(filepath.Join("webfonts", finfo.Name()))
		if err != nil {
			panic(err)
		}

		dst, err := os.Create(filepath.Join(destination, "fa", "webfonts", finfo.Name()))
		if err != nil {
			src.Close()
			panic(err)
		}

		if _, err := io.Copy(dst, src); err != nil {
			src.Close()
			panic(err)
		}

		src.Close()
	}
}

func generatePages(destination, baseurl string) {
	for _, section := range docs.Sections {
		for _, page := range section.Pages {
			page.BaseURL = baseurl
			page.Path = page.Path + ".html"
		}
	}
	for _, section := range docs.Sections {
		for _, page := range section.Pages {
			dir := filepath.Dir(page.Path)
			if dir != "/" {
				if err := os.MkdirAll(filepath.Join(destination, dir), 0o755); err != nil {
					panic(err)
				}
			}
			f, err := os.Create(filepath.Join(destination, page.Path))
			if err != nil {
				panic(err)
			}
			if err := page.Prepare(docs.Sections).Render(f); err != nil {
				panic(err)
			}
			if err := f.Close(); err != nil {
				panic(err)
			}
		}
	}
}
