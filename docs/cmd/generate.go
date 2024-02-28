package main

import (
	_ "embed"
	"io"
	"os"
	"path/filepath"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/docs"
	"github.com/willoma/bulma-gomponents/fa"
)

const (
	staticDir = "bulma-gomponents"
)

func generate() {
	generateAssets()
	generatePages()
}

func generateAssets() {
	if err := os.Mkdir(staticDir, 0o755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(staticDir, "bulma.css"), b.CSS, 0o644); err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(staticDir, "htmx.min.js"), htmxJS, 0o644); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(staticDir, "fa"), 0o755); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(staticDir, "fa", "css"), 0o755); err != nil {
		panic(err)
	}

	if err := os.Mkdir(filepath.Join(staticDir, "fa", "webfonts"), 0o755); err != nil {
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

		dst, err := os.Create(filepath.Join(staticDir, "fa", "css", finfo.Name()))
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

		dst, err := os.Create(filepath.Join(staticDir, "fa", "webfonts", finfo.Name()))
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

func generatePages() {
	for _, section := range docs.Sections {
		for _, page := range section.Pages {
			page.Path = page.Path + ".html"
		}
	}
	for _, section := range docs.Sections {
		for _, page := range section.Pages {
			dir := filepath.Dir(page.Path)
			if dir != "/" {
				if err := os.MkdirAll(filepath.Join(staticDir, dir), 0o755); err != nil {
					panic(err)
				}
			}
			f, err := os.Create(filepath.Join(staticDir, page.Path))
			if err != nil {
				panic(err)
			}
			if err := docs.Layout(page).Render(f); err != nil {
				panic(err)
			}
			if err := f.Close(); err != nil {
				panic(err)
			}
		}
	}
}
